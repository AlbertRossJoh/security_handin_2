# How to run
I have some requirements that needs to be met to be able to run this project.
You need:
- docker installed, you can get it from [here](https://www.docker.com)
- a unix shell, use wsl on windows.
- make


This project uses docker compose to automatically scale the amount of clients, though if you want to change this you're going to have to update the code.

You just need to update the amount of replicates in the docker compose file, the project has only been tested with 3 peers, so no guarantees can be made for more peers.

Given that docker is installed on your machine, and you have a Unix shell open you can run

```sh
sh launch.sh
```
or
```sh
chmod +x launch.sh
./launch.sh
```

Please be aware that the containers waits a bit before performing the protocol.

