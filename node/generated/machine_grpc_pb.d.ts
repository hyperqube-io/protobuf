// GENERATED CODE -- DO NOT EDIT!

// package: 
// file: machine.proto

import * as machine_pb from "./machine_pb";
import * as grpc from "@grpc/grpc-js";

interface IMachineServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  getMany: grpc.MethodDefinition<machine_pb.GetMachinesRequest, machine_pb.GetMachinesReply>;
}

export const MachineServiceService: IMachineServiceService;

export interface IMachineServiceServer extends grpc.UntypedServiceImplementation {
  getMany: grpc.handleUnaryCall<machine_pb.GetMachinesRequest, machine_pb.GetMachinesReply>;
}

export class MachineServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  getMany(argument: machine_pb.GetMachinesRequest, callback: grpc.requestCallback<machine_pb.GetMachinesReply>): grpc.ClientUnaryCall;
  getMany(argument: machine_pb.GetMachinesRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<machine_pb.GetMachinesReply>): grpc.ClientUnaryCall;
  getMany(argument: machine_pb.GetMachinesRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<machine_pb.GetMachinesReply>): grpc.ClientUnaryCall;
}
