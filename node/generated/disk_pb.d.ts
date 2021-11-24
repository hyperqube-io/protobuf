// package: 
// file: disk.proto

import * as jspb from "google-protobuf";
import * as util_pb from "./util_pb";

export class CopyFileRequest extends jspb.Message {
  getFrom(): string;
  setFrom(value: string): void;

  getTo(): string;
  setTo(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CopyFileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CopyFileRequest): CopyFileRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CopyFileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CopyFileRequest;
  static deserializeBinaryFromReader(message: CopyFileRequest, reader: jspb.BinaryReader): CopyFileRequest;
}

export namespace CopyFileRequest {
  export type AsObject = {
    from: string,
    to: string,
  }
}

export class FileRequest extends jspb.Message {
  getPath(): string;
  setPath(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: FileRequest): FileRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: FileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FileRequest;
  static deserializeBinaryFromReader(message: FileRequest, reader: jspb.BinaryReader): FileRequest;
}

export namespace FileRequest {
  export type AsObject = {
    path: string,
  }
}

export class GetFileContentsReply extends jspb.Message {
  getData(): string;
  setData(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFileContentsReply.AsObject;
  static toObject(includeInstance: boolean, msg: GetFileContentsReply): GetFileContentsReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetFileContentsReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFileContentsReply;
  static deserializeBinaryFromReader(message: GetFileContentsReply, reader: jspb.BinaryReader): GetFileContentsReply;
}

export namespace GetFileContentsReply {
  export type AsObject = {
    data: string,
  }
}

export class GetFileExistsReply extends jspb.Message {
  getExists(): boolean;
  setExists(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFileExistsReply.AsObject;
  static toObject(includeInstance: boolean, msg: GetFileExistsReply): GetFileExistsReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetFileExistsReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFileExistsReply;
  static deserializeBinaryFromReader(message: GetFileExistsReply, reader: jspb.BinaryReader): GetFileExistsReply;
}

export namespace GetFileExistsReply {
  export type AsObject = {
    exists: boolean,
  }
}

export class GetFileSizeReply extends jspb.Message {
  getSize(): number;
  setSize(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFileSizeReply.AsObject;
  static toObject(includeInstance: boolean, msg: GetFileSizeReply): GetFileSizeReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetFileSizeReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFileSizeReply;
  static deserializeBinaryFromReader(message: GetFileSizeReply, reader: jspb.BinaryReader): GetFileSizeReply;
}

export namespace GetFileSizeReply {
  export type AsObject = {
    size: number,
  }
}

export class WriteFileRequest extends jspb.Message {
  getPath(): string;
  setPath(value: string): void;

  getData(): string;
  setData(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WriteFileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: WriteFileRequest): WriteFileRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WriteFileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WriteFileRequest;
  static deserializeBinaryFromReader(message: WriteFileRequest, reader: jspb.BinaryReader): WriteFileRequest;
}

export namespace WriteFileRequest {
  export type AsObject = {
    path: string,
    data: string,
  }
}

