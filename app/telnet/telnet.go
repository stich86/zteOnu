package telnet

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func New(user string, pass string, ip string, port int) (*Telnet, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil, err
	}

	t := &Telnet{
		user: user,
		pass: pass,
		Conn: conn,
	}

	return t, nil
}

func (t *Telnet) PermTelnet(SecLvl int) error {
	if err := t.loginTelnet(); err != nil {
		return err
	}

	if err := t.modifyDB(SecLvl); err != nil {
		return err
	}

	if err := t.modifyFW(); err != nil {
		return err
	}

	return nil
}

func (t *Telnet) loginTelnet() error {
	return t.sendCmd(t.user, t.pass)
}

func (t *Telnet) modifyDB(SecLvl int) error {
	// set DB data
	prefix := "sendcmd 1 DB set TelnetCfg 0 "
	tsEnable := prefix + "TS_Enable 1 > /dev/null"
	lanEnable := prefix + "Lan_Enable 1  > /dev/null"
	tsLanUser := prefix + "TSLan_UName root > /dev/null"
	tsLanPwd := prefix + "TSLan_UPwd Zte521 > /dev/null"
	tsUser := prefix + "TS_UName root > /dev/null"
	tsPwd := prefix + "TS_UPwd Zte521 > /dev/null"
	maxConn := prefix + "Max_Con_Num 3 > /dev/null"
	initSecLvl := prefix + "InitSecLvl " + strconv.Itoa(SecLvl) + " > /dev/null"

	// save DB
	save := "sendcmd 1 DB save"

	if err := t.sendCmd(tsEnable, lanEnable, tsLanUser, tsLanPwd, tsUser, tsPwd, maxConn, initSecLvl, save); err != nil {
		return err
	}

	return nil
}

func (t *Telnet) modifyFW() error {
	// set DB data
	addrow := "sendcmd 1 DB addr FWSC 0  > /dev/null"
        prefix := "sendcmd 1 DB set FWSC 0 "
	viewName := prefix + "ViewName IGD.FWSc.FWSC1 > /dev/null"
	enable := prefix + "Enable 1 > /dev/null"
	intName := prefix + "INCName LAN > /dev/null"
	intViewName := prefix + "INCViewName IGD.LD1 > /dev/null"
	service := prefix + "Servise 8 > /dev/null"
	filter := prefix + "FilterTarget 1 > /dev/null"

	// save DB
	save := "sendcmd 1 DB save"


	if err := t.sendCmd(addrow, viewName, enable, intName, intViewName, service, filter, save); err != nil {
		return err
	}

	return nil
}

func (t *Telnet) sendCmd(commands ...string) error {
        for _, command := range commands {
                cmd := []byte(command + ctrl)

                actual, err := t.Conn.Write(cmd)
                if err != nil {
                        return fmt.Errorf("failed to send command %s: %v", command, err)
                }

                expected := len(cmd)
                if expected != actual {
                        return fmt.Errorf("transmission problem: tried sending %d bytes, but actually only sent %d bytes for command %s", expected, actual, command)
                }

                time.Sleep(200 * time.Millisecond)
        }

        return nil
}

func (t *Telnet) Reboot() error {
        if err := t.sendCmd("reboot"); err != nil {
                return err
        }

        return nil
}
