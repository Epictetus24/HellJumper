# HellJumper
A short program written in go to SSH into a jumpserver and prepare it for use.

# What it does:
Currently it will; Login to your (presumed Ubuntu) server with a private key, update & upgrade, modify the sshd_config to allow ssh port forwarding on any port and installs certbot - but does not configure it.

I plan to finish this script so that it also configures the ssl certificate with a domain provided as a flag. I would also one day like to use the digital ocean api so that this spins up a brand new droplet on each run.

Note:
It currently lacks support for passwords, I use public key based authentication. As digital ocean allows you to spin up a droplet with your key ready - that's what I've gone for. 

