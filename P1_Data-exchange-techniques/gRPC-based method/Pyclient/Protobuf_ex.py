import grpc
from concurrent import futures
import time
import service_pb2
import service_pb2_grpc


class CalculatorServicer(service_pb2_grpc.CalculatorServicer):
    def Add(self,request,context):
        response = service_pb2.AddResponse()
        response.result = request.a + request.b
        return response
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service_pb2_grpc.add_CalculatorServicer_to_server(CalculatorServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Python gRPC server started on 50051")
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(grace=5)


if __name__ == '__main__':
    serve()

