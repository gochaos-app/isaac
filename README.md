# Isaac

Isaac is a basic CLI with AI integrated. 

Isaac has two modes, chat and prompt mode.

In chat mode, Isaac will run commands in the background and can also save the prompts to a json file. 
```
isaac chat

@Isaac: Check the running process in linux    

To check the running processes in Linux, you can use the "top" command. This command provides a real-time view of the running processes and their resource usage, such as CPU, memory, and disk I/O. You can also use the "ps" command to list the currently running processes. Additionally, you can use the "htop" command, which is a more advanced and interactive version of top.

@Isaac: sys.save
Saving...

@Isaac: sys.exit
Goodbye!
```

In prompt mode, you can use isaac to made a simple query and get a response

```
isaac prompt "write a 3 line paragraph about AWS and how can it help my customers"

AWS is a comprehensive suite of cloud computing services that can help my customers reduce costs, improve efficiency, and scale their businesses. 
With AWS, my customers can access a wide range of computing, storage, and networking resources, allowing them to focus on their core businesses rather than on infrastructure management. 
AWS also offers a wide range of tools and features to help my customers manage their infrastructure, including identity and access management, security, and automation.
```
