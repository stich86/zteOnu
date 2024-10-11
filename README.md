# zteOnu

This is a fork from original [project](https://github.com/Septrum101/zteOnu) with some changes.

Please note that if you have firmware that doesn't open telnet, you can try `--new` flag and set mac-address of your NIC to `00:07:29:55:35:57`

Current supported options:

```
./zteOnu -h

Flags:
  -h, --help          help for zteOnu
  -i, --ip string     ONU ip address (default "192.168.1.1")
      --new           Use new method to open telnet, MAC address must set to 00:07:29:55:35:57
  -u, --user string   Factory mode auth username (If not provided, a known list will be used)
  -p, --pass string   Factory mode auth password (If not provided, a known list will be used)
      --port int      ONU http port (default 80)
      --seclvl int    Security level for telnet access, if you got "Permission Denied", try 3.
                      Use with --telnet flag (default 2)
      --telnet        Enable permanent telnet (user: root, pass: Zte521)
      --tp int        ONU telnet port (default 23)
```

# What's different from original one

- Added all known user/password combinations in a loop; the binary will attempt all of them to enable Telnet.
- Added the --seclvl parameter (default: 2) to change the Telnet access level and avoid the "Access Denied" error.
- Added firewall configuration when enabling permanent Telnet access.
- Modify login retries up to 5 attempts.
- Changed to use the default HTTP port 80 instead of 8080.

# Tested ONTs

| ONT     | Firmware                | Result                                             | Issues                                        |
|---------|-------------------------|----------------------------------------------------|-----------------------------------------------|
| F601V6  | V6.0.10P6N7 (OpenFiber) | Open Telnet (with known OF credentials)            | Permanent Telnet doesn't work with the tool   |
| F601V6  | V6.0.10N40 (TIM)        | Open Telnet                                        | Permanent Telnet doesn't work with the tool   |
| F601V7  | V7.0.10P6N7 (OpenFiber) | Open Telnet (with known OF credentials)            | Permanent Telnet doesn't have full privileges |
| F601V9  | V9.0.10P2N1 (OpenFiber) | Open Telnet (with known OF credentials)            | Permanent Telnet doesn't have full privileges |
| F6005V3 | V3.0.10P3N2 (OpenFiber) | Open Telnet (with known OF credentials)            |                                               |
| F6005V3 | V3.0.10N06 (TIM)        | Open Telnet with `--new` flag and mac-addr changed |                                               |
