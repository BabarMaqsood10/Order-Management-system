The gateway and the orders are different services.
Gateway is responsible for handling incoming request from the client , mob , customers.
Orders than do buisness logic.
Server in gateway is an Http server.
Gateway handles HTTP traffic and routing.
Server in Orders is an grpc server.
Orders handlers buisness logic and orderdata.

                      GATEWAY 
Gateway is the front door of the system.
Gateway purpose is to receive request from outside, usally via HTTP.
It translate or forward those request to correct interal service.
It returns the response back to the client.

In this project gateway is a layer between the browser/client and the internal
microservices like orders, payments and stocks.

                      Example 
Client sends:
POST /api/customers/123/orders
Gateway receives request
Gateway reads path/body data
Gateway calls the orders gRPC service
orders service does the actual work
Gateway sends back the result as HTTP JSON


                 GATEWAY MAINLY DOES:
routing
HTTP parsing
validation
calling downstream services
formatting responses
.