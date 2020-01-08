# HellJumper
A short program written in go to SSH into a jumpserver and prepare it for use.

# What it does:
Currently it will login to your server with a private key, update & upgrade, modify the sshd_config to allow ssh port forwarding on any port and finally it installs certbot - but does not configure it.

# Notes:
It currently lacks support for passwords, I use public key based authentication. As digital ocean allows you to spin up a droplet with your key ready - that's what I've gone for. 
