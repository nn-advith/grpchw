### Simple gRPC hello world

**Server**
Processes a HelloRequest and responds with a HelloResponse

**Client**
Connects to the server and sends a HelloRequest

**Protofile**

Defines GreetServer with rpc SayHello, which expects a HelloRequest and sends a HelloResponse
HelloRequest -> message with Name field
HelloResponse -> message with Message field
