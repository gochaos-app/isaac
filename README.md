# Isaac

Isaac is a basic CLI with AI integrated. 


https://github.com/gochaos-app/isaac/assets/47430149/d6d07d26-cd47-4d8d-b1c0-2e4b677826de


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

The first command to run is `isaac init` this sets up the config file `$HOME/.isaac_config.json`, this command can setup custom data or setting defaults values.



In chat mode, Isaac will run commands in the background and can also save the prompts to a json file. Isaac is capable to load, read and make a summary of text files, with the special command `load()`
```
isaac chat

@Isaac: Check the running process in linux    

To check the running processes in Linux, you can use the "top" command. This command provides a real-time view of the running processes and their resource usage, such as CPU, memory, and disk I/O. You can also use the "ps" command to list the currently running processes. Additionally, you can use the "htop" command, which is a more advanced and interactive version of top.

@Isaac: cmd(ps)
    PID TTY          TIME CMD
     52 pts/2    00:00:00 sh
     17 pts/2    00:00:00 isaac
     23 pts/2    00:00:00 ps

@Isaac: load(Makefile)

The text defines a series of targets (dev, prod, install, compile, and move) with corresponding commands. The targets are specified using the syntax: target: command.

The target "dev" runs the command "go build -o isaac .", which builds the Isaac executable and saves it in the isaac file. The target "prod" runs the same command as "dev", but also adds the "-ldflags" flag, which sets the build options for Isaac. The target "install" runs the "move" target, which moves the Isaac executable to the ~/bin directory.

The target "compile" runs a series of commands for each operating system and platform combination, using the GOOS and GOARCH environment variables to specify the target system. The commands build the Isaac executable for each target system, saving it in the bin directory with the naming convention "isaac-[OS]-[ARCH]".

In summary, the text defines a workflow for building and installing the Isaac executable on various operating systems and platforms.

@Isaac: sys.save
Saving...

@Isaac: sys.exit
Goodbye!

```

File (only includes prompts and answers) get saved in working directory and is named `prompts.json`

In prompt mode, you can use isaac to made a simple query and get a response

```
isaac prompt "write a 3 line paragraph about AWS and how can it help my customers"

AWS is a comprehensive suite of cloud computing services that can help my customers reduce costs, improve efficiency, and scale their businesses. 
With AWS, my customers can access a wide range of computing, storage, and networking resources, allowing them to focus on their core businesses rather than on infrastructure management. 
AWS also offers a wide range of tools and features to help my customers manage their infrastructure, including identity and access management, security, and automation.
```


