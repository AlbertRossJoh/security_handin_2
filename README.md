# How to run
I have some requirements that needs to be met to be able to run this project.
You need:
- docker installed, you can get it from [here](https://www.docker.com)
- a unit shell, use wsl on windows.
- make


This project uses docker compose to automatically scale the amount of clients, though if you want to change this you're going to have to update the code.

You just need to update the amount of replicates in the docker compose file and set the `shareChan` size by the amount of clients minus 1. I use this to make sure that even though the clients run concurrently, we still need some synchronous behavior.

Given that docker is installed on your machine, and you have a Unix shell open you can run
```
sh launch.sh
```
or
```
chmod +x launch.sh
./launch.sh
```