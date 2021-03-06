// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var machine_pb = require('./machine_pb.js');
var network_pb = require('./network_pb.js');
var util_pb = require('./util_pb.js');

function serialize_ActionReply(arg) {
  if (!(arg instanceof util_pb.ActionReply)) {
    throw new Error('Expected argument of type ActionReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ActionReply(buffer_arg) {
  return util_pb.ActionReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CreateMachineReply(arg) {
  if (!(arg instanceof machine_pb.CreateMachineReply)) {
    throw new Error('Expected argument of type CreateMachineReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreateMachineReply(buffer_arg) {
  return machine_pb.CreateMachineReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CreateMachineRequest(arg) {
  if (!(arg instanceof machine_pb.CreateMachineRequest)) {
    throw new Error('Expected argument of type CreateMachineRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreateMachineRequest(buffer_arg) {
  return machine_pb.CreateMachineRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetMachineRequest(arg) {
  if (!(arg instanceof machine_pb.GetMachineRequest)) {
    throw new Error('Expected argument of type GetMachineRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetMachineRequest(buffer_arg) {
  return machine_pb.GetMachineRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetMachinesReply(arg) {
  if (!(arg instanceof machine_pb.GetMachinesReply)) {
    throw new Error('Expected argument of type GetMachinesReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetMachinesReply(buffer_arg) {
  return machine_pb.GetMachinesReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetMachinesRequest(arg) {
  if (!(arg instanceof machine_pb.GetMachinesRequest)) {
    throw new Error('Expected argument of type GetMachinesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetMachinesRequest(buffer_arg) {
  return machine_pb.GetMachinesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetVncPortReply(arg) {
  if (!(arg instanceof machine_pb.GetVncPortReply)) {
    throw new Error('Expected argument of type GetVncPortReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetVncPortReply(buffer_arg) {
  return machine_pb.GetVncPortReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_LinkNetworkReply(arg) {
  if (!(arg instanceof machine_pb.LinkNetworkReply)) {
    throw new Error('Expected argument of type LinkNetworkReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_LinkNetworkReply(buffer_arg) {
  return machine_pb.LinkNetworkReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_MachineNetworkRequest(arg) {
  if (!(arg instanceof machine_pb.MachineNetworkRequest)) {
    throw new Error('Expected argument of type MachineNetworkRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_MachineNetworkRequest(buffer_arg) {
  return machine_pb.MachineNetworkRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_StopMachinesRequest(arg) {
  if (!(arg instanceof machine_pb.StopMachinesRequest)) {
    throw new Error('Expected argument of type StopMachinesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_StopMachinesRequest(buffer_arg) {
  return machine_pb.StopMachinesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateMachineBootTypeRequest(arg) {
  if (!(arg instanceof machine_pb.UpdateMachineBootTypeRequest)) {
    throw new Error('Expected argument of type UpdateMachineBootTypeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdateMachineBootTypeRequest(buffer_arg) {
  return machine_pb.UpdateMachineBootTypeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateMachineCpusRequest(arg) {
  if (!(arg instanceof machine_pb.UpdateMachineCpusRequest)) {
    throw new Error('Expected argument of type UpdateMachineCpusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdateMachineCpusRequest(buffer_arg) {
  return machine_pb.UpdateMachineCpusRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateMachineDisksRequest(arg) {
  if (!(arg instanceof machine_pb.UpdateMachineDisksRequest)) {
    throw new Error('Expected argument of type UpdateMachineDisksRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdateMachineDisksRequest(buffer_arg) {
  return machine_pb.UpdateMachineDisksRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateMachineMemoryRequest(arg) {
  if (!(arg instanceof machine_pb.UpdateMachineMemoryRequest)) {
    throw new Error('Expected argument of type UpdateMachineMemoryRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdateMachineMemoryRequest(buffer_arg) {
  return machine_pb.UpdateMachineMemoryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var MachineServiceService = exports.MachineServiceService = {
  getVncPort: {
    path: '/MachineService/GetVncPort',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.GetMachineRequest,
    responseType: machine_pb.GetVncPortReply,
    requestSerialize: serialize_GetMachineRequest,
    requestDeserialize: deserialize_GetMachineRequest,
    responseSerialize: serialize_GetVncPortReply,
    responseDeserialize: deserialize_GetVncPortReply,
  },
  getMany: {
    path: '/MachineService/GetMany',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.GetMachinesRequest,
    responseType: machine_pb.GetMachinesReply,
    requestSerialize: serialize_GetMachinesRequest,
    requestDeserialize: deserialize_GetMachinesRequest,
    responseSerialize: serialize_GetMachinesReply,
    responseDeserialize: deserialize_GetMachinesReply,
  },
  create: {
    path: '/MachineService/Create',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.CreateMachineRequest,
    responseType: machine_pb.CreateMachineReply,
    requestSerialize: serialize_CreateMachineRequest,
    requestDeserialize: deserialize_CreateMachineRequest,
    responseSerialize: serialize_CreateMachineReply,
    responseDeserialize: deserialize_CreateMachineReply,
  },
  delete: {
    path: '/MachineService/Delete',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.GetMachinesRequest,
    responseType: util_pb.ActionReply,
    requestSerialize: serialize_GetMachinesRequest,
    requestDeserialize: deserialize_GetMachinesRequest,
    responseSerialize: serialize_ActionReply,
    responseDeserialize: deserialize_ActionReply,
  },
  start: {
    path: '/MachineService/Start',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.GetMachinesRequest,
    responseType: util_pb.ActionReply,
    requestSerialize: serialize_GetMachinesRequest,
    requestDeserialize: deserialize_GetMachinesRequest,
    responseSerialize: serialize_ActionReply,
    responseDeserialize: deserialize_ActionReply,
  },
  stop: {
    path: '/MachineService/Stop',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.StopMachinesRequest,
    responseType: util_pb.ActionReply,
    requestSerialize: serialize_StopMachinesRequest,
    requestDeserialize: deserialize_StopMachinesRequest,
    responseSerialize: serialize_ActionReply,
    responseDeserialize: deserialize_ActionReply,
  },
  linkNetwork: {
    path: '/MachineService/LinkNetwork',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.MachineNetworkRequest,
    responseType: machine_pb.LinkNetworkReply,
    requestSerialize: serialize_MachineNetworkRequest,
    requestDeserialize: deserialize_MachineNetworkRequest,
    responseSerialize: serialize_LinkNetworkReply,
    responseDeserialize: deserialize_LinkNetworkReply,
  },
  unlinkNetwork: {
    path: '/MachineService/UnlinkNetwork',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.MachineNetworkRequest,
    responseType: util_pb.ActionReply,
    requestSerialize: serialize_MachineNetworkRequest,
    requestDeserialize: deserialize_MachineNetworkRequest,
    responseSerialize: serialize_ActionReply,
    responseDeserialize: deserialize_ActionReply,
  },
  updateBootType: {
    path: '/MachineService/UpdateBootType',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.UpdateMachineBootTypeRequest,
    responseType: util_pb.ActionReply,
    requestSerialize: serialize_UpdateMachineBootTypeRequest,
    requestDeserialize: deserialize_UpdateMachineBootTypeRequest,
    responseSerialize: serialize_ActionReply,
    responseDeserialize: deserialize_ActionReply,
  },
  updateCpus: {
    path: '/MachineService/UpdateCpus',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.UpdateMachineCpusRequest,
    responseType: util_pb.ActionReply,
    requestSerialize: serialize_UpdateMachineCpusRequest,
    requestDeserialize: deserialize_UpdateMachineCpusRequest,
    responseSerialize: serialize_ActionReply,
    responseDeserialize: deserialize_ActionReply,
  },
  updateDisks: {
    path: '/MachineService/UpdateDisks',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.UpdateMachineDisksRequest,
    responseType: util_pb.ActionReply,
    requestSerialize: serialize_UpdateMachineDisksRequest,
    requestDeserialize: deserialize_UpdateMachineDisksRequest,
    responseSerialize: serialize_ActionReply,
    responseDeserialize: deserialize_ActionReply,
  },
  updateMemory: {
    path: '/MachineService/UpdateMemory',
    requestStream: false,
    responseStream: false,
    requestType: machine_pb.UpdateMachineMemoryRequest,
    responseType: util_pb.ActionReply,
    requestSerialize: serialize_UpdateMachineMemoryRequest,
    requestDeserialize: deserialize_UpdateMachineMemoryRequest,
    responseSerialize: serialize_ActionReply,
    responseDeserialize: deserialize_ActionReply,
  },
};

exports.MachineServiceClient = grpc.makeGenericClientConstructor(MachineServiceService);
