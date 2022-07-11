# Simple Server Project

## Simple description:
I will develop this project as a tutorial for myself. 
But in order not to write empty code, I can simultaneously solve routine tasks that I often encounter 
in my daily work, for example, deploying a server with a basic web setup.

This process includes creating users for the service and executing code, 
deploying a node, nginx, php, git, and a gitlab runner. At the same time, 
everything should work like clockwork and be configured literally in one command.

## Future functionality:

1. Create "Sudo" user
2. Create "Execution" user
3. Add basic commands before any installation (apt update && apt upgrade)
4. Install and configure Nginx
5. Install and configure Php
6. Install and configure Gitlab Runner
7. Install MySQL
8. Install Node
9. Install Git

## Install

### **Attention! This is an experimental program. Something may be working incorrectly. All problems can be sent to the "Problems" section.**



Go home dir:
```bash
cd ~
```

Download binary file:
```bash
wget https://github.com/Sinersis/simple-server/releases/download/0.0.1/easy-server
```

Set chmod:
```bash
chmod +x ./easy-server
```

And just run:
```bash
./easy-server start
```