# HellJumper
A short program written in go to SSH into a jumpserver and prepare it for use. Mostly for learning Go, otherwise a bash script could do basically all of this.

# What it does:
Currently it will; Login to your (presumed Ubuntu) server with a private key, update & upgrade, modify the sshd_config to allow ssh port forwarding on any port and installs certbot with an apache certificate - but does not configure it fully.

I would one day like to use the digital ocean api so that this spins up a brand new droplet on each run.

Note:
It currently lacks support for passwords, I use public key based authentication. As digital ocean allows you to spin up a droplet with your key ready - that's what I've gone for. 

```sh
 go run helljumper.go /root/.ssh/id_rsa sshserver mydomain.com
Remote server: remoteserver:22
[+] Updating your server
Get:1 http://mirrors.digitalocean.com/ubuntu bionic InRelease [242 kB]
Hit:2 http://mirrors.digitalocean.com/ubuntu bionic-updates InRelease
Hit:3 http://mirrors.digitalocean.com/ubuntu bionic-backports InRelease
Hit:4 http://security.ubuntu.com/ubuntu bionic-security InRelease
Hit:5 http://ppa.launchpad.net/certbot/certbot/ubuntu bionic InRelease
Fetched 242 kB in 1s (364 kB/s)
Reading package lists...
Building dependency tree...
Reading state information...
All packages are up to date.

[+] Updating complete, now moving to upgrade
Reading package lists...
Building dependency tree...
Reading state information...
Calculating upgrade...
The following packages were automatically installed and are no longer required:
  grub-pc-bin libdumbnet1
Use 'apt autoremove' to remove them.
0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.

[+] Upgrade completed. Modifying SSH config.

[+] Installing Certbot Repo and client
 This is the PPA for packages prepared by Debian Let's Encrypt Team and backported for Ubuntu.

Note: Packages are only provided for currently supported Ubuntu releases.
 More info: https://launchpad.net/~certbot/+archive/ubuntu/certbot
Get:1 http://mirrors.digitalocean.com/ubuntu bionic InRelease [242 kB]
Hit:2 http://security.ubuntu.com/ubuntu bionic-security InRelease
Hit:3 http://ppa.launchpad.net/certbot/certbot/ubuntu bionic InRelease
Hit:4 http://mirrors.digitalocean.com/ubuntu bionic-updates InRelease
Hit:5 http://mirrors.digitalocean.com/ubuntu bionic-backports InRelease
Fetched 242 kB in 1s (317 kB/s)
Reading package lists...
Reading package lists...
Building dependency tree...
Reading state information...
python-certbot-apache is already the newest version (0.31.0-1+ubuntu18.04.1+certbot+1).
The following packages were automatically installed and are no longer required:
  grub-pc-bin libdumbnet1
Use 'apt autoremove' to remove them.
0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.

[+] Allowing port 80,443 through firewall.
[+] Grabbing certbot certificate with 'certonly --apache'.
Skipping adding existing rule
Skipping adding existing rule (v6)
IMPORTANT NOTES:
 - Congratulations! Your certificate and chain have been saved at:
   /etc/letsencrypt/live/redacted/fullchain.pem
   Your key file has been saved at:
   /etc/letsencrypt/live/redacted/privkey.pem
   Your cert will expire on redacted. To obtain a new or tweaked
   version of this certificate in the future, simply run certbot
   again. To non-interactively renew *all* of your certificates, run
   "certbot renew"
 - If you like Certbot, please consider supporting our work by:

   Donating to ISRG / Let's Encrypt:   https://letsencrypt.org/donate
   Donating to EFF:                    https://eff.org/donate-le

```
