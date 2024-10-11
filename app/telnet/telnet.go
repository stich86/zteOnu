package telnet

import (
	"fmt"
	"net"
	"strings"
	"strconv"
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
	cmd := []byte(strings.Join(commands, ctrl) + ctrl)
	n, err := t.Conn.Write(cmd)
	if err != nil {
		return err
	}

	if expected, actual := len(cmd), n; expected != actual {
		err := fmt.Errorf("transmission problem: tried sending %d bytes, but actually only sent %d bytes", expected, actual)
		return err
	}

	return nil
}

func (t *Telnet) Reboot() error {
        if err := t.sendCmd("reboot"); err != nil {
                return err
        }

        return nil
}
