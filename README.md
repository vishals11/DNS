# DNS
Drone Navigation System


Develop a Drone Navigation System (DNS) to locate a databank to upload a data keeping in mind following things  
● each observed sector of the galaxy has unique numeric SectorID assigned to it  
● each sector will have at least one DNS deployed  
● each sector has different number of drones deployed at any given moment  


A. Dependencies

Golang IDE   
Docker

Dependency management using go module - go mod init  

B.  
Run docker daemon on host OS  
Build docker image using following command:  
docker build -t "drone_navigation_system:1.0" .  

Run docker container using following command:  
docker run -p 8000:8000 --env SECTOR_ID=1 drone_navigation_system:1.0  


C. Sample DNS Request  

set  Content-type = application/json  

POST http://localhost:8000/v1/dns  

{  
"x": "123.12",  
"y": "456.56",  
"z": "789.89",  
"vel": "20.0"  
}  

Sample DNS Response  
{  
    "loc": 1389.5700000000002  
}  
