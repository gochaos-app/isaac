# Isaac

Isaac is a basic CLI with AI integrated. 

https://github.com/gochaos-app/isaac/assets/47430149/fd934c52-ef6b-453b-8bab-0572c8613b4e

## Introduction
Isaac has two modes, chat and prompt mode.

```
NAME:
   Isaac -  CLI that can help you on various tasks

USAGE:
   Isaac [global options] command [command options] [arguments...]

COMMANDS:
   init, i    Initialize Isaac
   chat, c    Chat with Isaac
   prompt, p  make a simple prompt, prompt should be enclosed in quotes
   help, h    Shows a list of commands or help for one command
```

Special commands in chat mode: 

* **command**:    User can input what it wnats to do and Isaac will return a posssible command to use as well as ask for confirmation. 

* **kubernetes**: Ask kubernetes related questions and the response will be a kubectl command as well with a brief explanation of the command.

* **file**: load a file and ask for a summary or review general written code.
* **save**: save the prompts in a file, default name `prompts.jsonl`.

* **uploadS3**:  Upload prompts file to an s3 specified in init file.


## Usage

````markdown
isaac chat

@Isaac →  Check the running process in linux    

To check the running processes in Linux, you can use the "top" command. This command provides a real-time view of the running processes and their resource usage, such as CPU,
memory, and disk I/O. You can also use the "ps" command to list the currently running processes. Additionally, you can use the "htop" command, which is a more advanced and
interactive version of top.

@Isaac → command: check the running process in linux

```ps```
Execute command? Only yes is accepted: 
   ps   
yes 

PID TTY          TIME CMD
52 pts/2    00:00:00 sh
17 pts/2    00:00:00 isaac
23 pts/2    00:00:00 ps

@Isaac → kubernetes: check for pods in default namespace

```
kubectl get pods -n default

```

This command will list all of the pods in the`default`namespace. The`-n`flag allows you to specify a specific namespace, and the`get pods`command lists all of the pods in that namespace.

@Isaac → file:Makefile make a summary of this makefile
File does exist

The provided Makefile consists of a set of targets, each representing a specific action. The targets are organized into sections, denoted by labels like "dev", "prod", "install", and "compile".

Targets starting with ".PHONY" are recognized as phony targets. Phony targets are used to indicate that a command or actions should be taken, rather than actually performing a specific task.

Targets starting with "dev" are used to compile and build the "isaac" program. The "go build" command is used to compile the source code and build an executable binary, which is then placed in "~/bin/isaac".

Targets starting with "prod" are used to compile and build the "isaac" program with additional options, specifically the "-ldflags" option, which is used to specify additional flags for the linker. The resulting binary is then placed in "~/bin/isaac".

Targets starting with "install" are used to perform the installation and deployment of the "isaac" program. The "install" target first performs the "prod" target, and then performs the "move" target, which moves the compiled binary


@Isaac: sys.exit
Goodbye!

````

### Prompt mode

```
isaac prompt "write a 3 line paragraph about AWS and how can it help my customers"

AWS is a comprehensive suite of cloud computing services that can help my customers reduce costs, improve efficiency, and scale their businesses. 
With AWS, my customers can access a wide range of computing, storage, and networking resources, allowing them to focus on their core businesses rather than on infrastructure management. 
AWS also offers a wide range of tools and features to help my customers manage their infrastructure, including identity and access management, security, and automation.
```


