package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"bytes"

	"golang.org/x/crypto/ssh"
)

func runSSH(userkey, remoteserver, cmd string) {

        key, err := ioutil.ReadFile(userkey)
        if err != nil {
            log.Fatalf("unable to read private key: %v", err)
        }

        // Create the Signer for this private key.
        signer, err := ssh.ParsePrivateKey(key)
        if err != nil {
            log.Fatalf("unable to parse private key: %v", err)
        }

        config := &ssh.ClientConfig{
            User: "root",
            Auth: []ssh.AuthMethod{
                // Use the PublicKeys method for remote authentication.
                ssh.PublicKeys(signer),
            },
            //HostKeyCallback: ssh.FixedHostKey(hostKey),
            HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        }
	// Connect to the remote server and perform the SSH handshake.
        client, err := ssh.Dial("tcp", remoteserver, config)
        if err != nil {
            log.Fatalf("unable to connect: %v", err)
        }

	update, err := client.NewSession()
        if err != nil {
                log.Fatal("Failed to create session: ", err)
        }

        defer update.Close()

        var b bytes.Buffer
        update.Stdout = &b
        if err := update.Run(cmd); err != nil {
                log.Fatal("Failed to run: " + err.Error())
        }
        fmt.Println(b.String())
}

func main() {
	if len(os.Args) < 3 {
                log.Fatalf("Usage: %s [sshprivatekey] [sshserver] [certbot_domain]",os.Args[0])
	}
	//var hostKey ssh.PublicKey
	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	flagkey := os.Args[1]
	flagserver := os.Args[2]
	flagserver += ":22"
	fmt.Printf("Remote server: %s\n", flagserver)

        fmt.Println("[+] Updating your server")
	updater := "/usr/bin/apt update"
	runSSH(flagkey,flagserver,updater)

        fmt.Println("[+] Updating complete, now moving to upgrade")

        upgrader := "/usr/bin/apt upgrade -y"
	runSSH(flagkey,flagserver,upgrader)

        //add clientspecified yes to ssh config

        fmt.Println("[+] Upgrade completed. Modifying SSH config.")

        sshconfcmd := "echo c2VkIC1pICcvR2F0ZXdheVBvcnRzL2NcR2F0ZXdheVBvcnRzIGNsaWVudHNwZWNpZmllZCcgL2V0Yy9zc2gvc3NoZF9jb25maWcK | base64 -d | /bin/bash"
	runSSH(flagkey,flagserver,sshconfcmd)

        //install certbot by adding repo and then installing the client
        fmt.Println("[+] Installing Certbot Repo and client")
        certbotSSH := "add-apt-repository ppa:certbot/certbot && apt install python-certbot-apache -y"
	runSSH(flagkey,flagserver,certbotSSH)

	//start apache2 and setup stuff to then install certifcate
	fmt.Println("[+] Ensuring Apache2 is started")
	apache2 := "service apache2 start"
	runSSH(flagkey,flagserver,apache2)

	fmt.Printf("[+] Allowing port 80,443 through firewall.\n[+] Grabbing certbot certificate with non-interactive, no email ;'certonly --apache'.\n")
	certbotEX := "ufw allow 80,443/tcp && certbot certonly --apache --register-unsafely-without-email -n --domains "
	certbotEX += os.Args[3]
	runSSH(flagkey,flagserver,certbotEX)

}
