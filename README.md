# chatapp
Steps to run program
FOR THE FIRST TIME RUNNING THE PROGRAM, DO POINT 5 FIRST AND REMEMBER TO RUN MYSQL IN services(search services in windows)
1.) Create and start docker container [ONLY NEED TO RUN STEP 1 ONCE, FOR FUTURE RUN OF THE PROGRAM YOU JUST NEED TO START DOCKER AND RUN THE chatapp container]
    a.) Open command prompt
    b.) cd into chatapp folder
    c.) Create and Start the docker using the command
        docker-compose up -d 
    d.) Subsequent running of program can just run the image (chatapp) from docker
        
2.) Open consumer.go in visual studio code window.
    go run consumer.go

3.) Open server folder in a new visual studio code window.
    open main.go in visual studio code
    go run .

4.) Go to browser 
   http://localhost:5221
   
You only need to do point 5 Once! 
5.) Create user account database on mysql(not docker) to store user account information 
	a.) In mysql workbench, create an account name user1 with all administrative privilege
	b.) Use the following sql script to create the database and table
    		CREATE database mysql;
    		USE MYSTOREDBFOODPANDA;
    		CREATE TABLE Users (UserName VARCHAR(30) NOT NULL PRIMARY KEY, Password VARCHAR(256), FirstName VARCHAR(30), LastName VARCHAR(30), Language VARCHAR(30));

6.) NSQ admin
   http://localhost:4171
/***********************************************************************************************************************************************************************/

Meeting 3 Links 2 Sep 2021:
1.) https://bluzelle.com/blog/things-you-should-know-about-database-caching

2.) https://www.sohamkamani.com/golang/rsa-encryption/

3.) https://stackoverflow.com/questions/6476945/how-do-i-run-redis-on-windows
https://crypto.stackexchange.com/questions/35530/where-and-how-to-store-private-keys-in-web-applications-for-private-messaging-wi/52488

/***********************************************************************************************************************************************************************/
Meeting 5 Links 7 Oct 2021:
1.) https://github.com/gorilla/websocket/tree/master/examples/chat

2.) https://github.com/gorilla/websocket/blob/master/examples/chat/client.go

3.) https://stackoverflow.com/questions/10668028/message-queues-vs-sockets

4.) https://study.com/academy/lesson/synchronous-asynchronous-networks-in-wan.html#:~:text=An%20asynchronous%20network%20is%20the,also%20known%20as%20half%2Dduplex.

5.) https://medium.com/@fzambia/building-real-time-messaging-server-in-go-5661c0a45248

6.) https://medium.com/@jawadahmadd/nsq-with-go-77ca1b69c4ec
