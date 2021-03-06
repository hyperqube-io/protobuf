/**
 * @fileoverview
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var util_pb = require('./util_pb.js');
goog.exportSymbol('proto.CreateNetworkReply', null, global);
goog.exportSymbol('proto.CreateNetworkRequest', null, global);
goog.exportSymbol('proto.GetNetworksReply', null, global);
goog.exportSymbol('proto.GetNetworksRequest', null, global);
goog.exportSymbol('proto.Network', null, global);
goog.exportSymbol('proto.UpdateNetworkFlagRequest', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Network = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.Network.repeatedFields_, null);
};
goog.inherits(proto.Network, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.Network.displayName = 'proto.Network';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.Network.repeatedFields_ = [3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.Network.prototype.toObject = function(opt_includeInstance) {
  return proto.Network.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.Network} msg The msg instance to transform.
 * @return {!Object}
 */
proto.Network.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: msg.getId(),
    name: msg.getName(),
    machineidsList: jspb.Message.getField(msg, 3),
    dhcp: msg.getDhcp(),
    internet: msg.getInternet(),
    xml: msg.getXml()
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Network}
 */
proto.Network.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.Network;
  return proto.Network.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Network} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Network}
 */
proto.Network.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {!Array.<number>} */ (reader.readPackedUint32());
      msg.setMachineidsList(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setDhcp(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setInternet(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setXml(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.Network} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.Network.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Network.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.Network.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = this.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = this.getMachineidsList();
  if (f.length > 0) {
    writer.writePackedUint32(
      3,
      f
    );
  }
  f = this.getDhcp();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = this.getInternet();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
  f = this.getXml();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * Creates a deep clone of this proto. No data is shared with the original.
 * @return {!proto.Network} The clone.
 */
proto.Network.prototype.cloneMessage = function() {
  return /** @type {!proto.Network} */ (jspb.Message.cloneMessage(this));
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.Network.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldProto3(this, 1, ""));
};


/** @param {string} value  */
proto.Network.prototype.setId = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional string name = 2;
 * @return {string}
 */
proto.Network.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldProto3(this, 2, ""));
};


/** @param {string} value  */
proto.Network.prototype.setName = function(value) {
  jspb.Message.setField(this, 2, value);
};


/**
 * repeated uint32 machineIds = 3;
 * If you change this array by adding, removing or replacing elements, or if you
 * replace the array itself, then you must call the setter to update it.
 * @return {!Array.<number>}
 */
proto.Network.prototype.getMachineidsList = function() {
  return /** @type {!Array.<number>} */ (jspb.Message.getField(this, 3));
};


/** @param {Array.<number>} value  */
proto.Network.prototype.setMachineidsList = function(value) {
  jspb.Message.setField(this, 3, value || []);
};


proto.Network.prototype.clearMachineidsList = function() {
  jspb.Message.setField(this, 3, []);
};


/**
 * optional bool dhcp = 4;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.Network.prototype.getDhcp = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldProto3(this, 4, false));
};


/** @param {boolean} value  */
proto.Network.prototype.setDhcp = function(value) {
  jspb.Message.setField(this, 4, value);
};


/**
 * optional bool internet = 5;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.Network.prototype.getInternet = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldProto3(this, 5, false));
};


/** @param {boolean} value  */
proto.Network.prototype.setInternet = function(value) {
  jspb.Message.setField(this, 5, value);
};


/**
 * optional string xml = 6;
 * @return {string}
 */
proto.Network.prototype.getXml = function() {
  return /** @type {string} */ (jspb.Message.getFieldProto3(this, 6, ""));
};


/** @param {string} value  */
proto.Network.prototype.setXml = function(value) {
  jspb.Message.setField(this, 6, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.GetNetworksRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.GetNetworksRequest.repeatedFields_, null);
};
goog.inherits(proto.GetNetworksRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.GetNetworksRequest.displayName = 'proto.GetNetworksRequest';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.GetNetworksRequest.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.GetNetworksRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.GetNetworksRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.GetNetworksRequest} msg The msg instance to transform.
 * @return {!Object}
 */
proto.GetNetworksRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    idsList: jspb.Message.getField(msg, 1)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.GetNetworksRequest}
 */
proto.GetNetworksRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.GetNetworksRequest;
  return proto.GetNetworksRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.GetNetworksRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.GetNetworksRequest}
 */
proto.GetNetworksRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.getIdsList().push(value);
      msg.setIdsList(msg.getIdsList());
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.GetNetworksRequest} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.GetNetworksRequest.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.GetNetworksRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.GetNetworksRequest.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getIdsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
};


/**
 * Creates a deep clone of this proto. No data is shared with the original.
 * @return {!proto.GetNetworksRequest} The clone.
 */
proto.GetNetworksRequest.prototype.cloneMessage = function() {
  return /** @type {!proto.GetNetworksRequest} */ (jspb.Message.cloneMessage(this));
};


/**
 * repeated string ids = 1;
 * If you change this array by adding, removing or replacing elements, or if you
 * replace the array itself, then you must call the setter to update it.
 * @return {!Array.<string>}
 */
proto.GetNetworksRequest.prototype.getIdsList = function() {
  return /** @type {!Array.<string>} */ (jspb.Message.getField(this, 1));
};


/** @param {Array.<string>} value  */
proto.GetNetworksRequest.prototype.setIdsList = function(value) {
  jspb.Message.setField(this, 1, value || []);
};


proto.GetNetworksRequest.prototype.clearIdsList = function() {
  jspb.Message.setField(this, 1, []);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.GetNetworksReply = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.GetNetworksReply.repeatedFields_, null);
};
goog.inherits(proto.GetNetworksReply, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.GetNetworksReply.displayName = 'proto.GetNetworksReply';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.GetNetworksReply.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.GetNetworksReply.prototype.toObject = function(opt_includeInstance) {
  return proto.GetNetworksReply.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.GetNetworksReply} msg The msg instance to transform.
 * @return {!Object}
 */
proto.GetNetworksReply.toObject = function(includeInstance, msg) {
  var f, obj = {
    networksList: jspb.Message.toObjectList(msg.getNetworksList(),
    proto.Network.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.GetNetworksReply}
 */
proto.GetNetworksReply.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.GetNetworksReply;
  return proto.GetNetworksReply.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.GetNetworksReply} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.GetNetworksReply}
 */
proto.GetNetworksReply.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.Network;
      reader.readMessage(value,proto.Network.deserializeBinaryFromReader);
      msg.getNetworksList().push(value);
      msg.setNetworksList(msg.getNetworksList());
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.GetNetworksReply} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.GetNetworksReply.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.GetNetworksReply.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.GetNetworksReply.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getNetworksList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.Network.serializeBinaryToWriter
    );
  }
};


/**
 * Creates a deep clone of this proto. No data is shared with the original.
 * @return {!proto.GetNetworksReply} The clone.
 */
proto.GetNetworksReply.prototype.cloneMessage = function() {
  return /** @type {!proto.GetNetworksReply} */ (jspb.Message.cloneMessage(this));
};


/**
 * repeated Network networks = 1;
 * If you change this array by adding, removing or replacing elements, or if you
 * replace the array itself, then you must call the setter to update it.
 * @return {!Array.<!proto.Network>}
 */
proto.GetNetworksReply.prototype.getNetworksList = function() {
  return /** @type{!Array.<!proto.Network>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.Network, 1));
};


/** @param {Array.<!proto.Network>} value  */
proto.GetNetworksReply.prototype.setNetworksList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


proto.GetNetworksReply.prototype.clearNetworksList = function() {
  this.setNetworksList([]);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.CreateNetworkRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.CreateNetworkRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.CreateNetworkRequest.displayName = 'proto.CreateNetworkRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.CreateNetworkRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.CreateNetworkRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.CreateNetworkRequest} msg The msg instance to transform.
 * @return {!Object}
 */
proto.CreateNetworkRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: msg.getName(),
    xml: msg.getXml()
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.CreateNetworkRequest}
 */
proto.CreateNetworkRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.CreateNetworkRequest;
  return proto.CreateNetworkRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.CreateNetworkRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.CreateNetworkRequest}
 */
proto.CreateNetworkRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setXml(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.CreateNetworkRequest} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.CreateNetworkRequest.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.CreateNetworkRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.CreateNetworkRequest.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = this.getXml();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * Creates a deep clone of this proto. No data is shared with the original.
 * @return {!proto.CreateNetworkRequest} The clone.
 */
proto.CreateNetworkRequest.prototype.cloneMessage = function() {
  return /** @type {!proto.CreateNetworkRequest} */ (jspb.Message.cloneMessage(this));
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.CreateNetworkRequest.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldProto3(this, 1, ""));
};


/** @param {string} value  */
proto.CreateNetworkRequest.prototype.setName = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional string xml = 2;
 * @return {string}
 */
proto.CreateNetworkRequest.prototype.getXml = function() {
  return /** @type {string} */ (jspb.Message.getFieldProto3(this, 2, ""));
};


/** @param {string} value  */
proto.CreateNetworkRequest.prototype.setXml = function(value) {
  jspb.Message.setField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.CreateNetworkReply = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.CreateNetworkReply, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.CreateNetworkReply.displayName = 'proto.CreateNetworkReply';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.CreateNetworkReply.prototype.toObject = function(opt_includeInstance) {
  return proto.CreateNetworkReply.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.CreateNetworkReply} msg The msg instance to transform.
 * @return {!Object}
 */
proto.CreateNetworkReply.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: msg.getId()
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.CreateNetworkReply}
 */
proto.CreateNetworkReply.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.CreateNetworkReply;
  return proto.CreateNetworkReply.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.CreateNetworkReply} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.CreateNetworkReply}
 */
proto.CreateNetworkReply.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.CreateNetworkReply} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.CreateNetworkReply.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.CreateNetworkReply.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.CreateNetworkReply.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * Creates a deep clone of this proto. No data is shared with the original.
 * @return {!proto.CreateNetworkReply} The clone.
 */
proto.CreateNetworkReply.prototype.cloneMessage = function() {
  return /** @type {!proto.CreateNetworkReply} */ (jspb.Message.cloneMessage(this));
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.CreateNetworkReply.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldProto3(this, 1, ""));
};


/** @param {string} value  */
proto.CreateNetworkReply.prototype.setId = function(value) {
  jspb.Message.setField(this, 1, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.UpdateNetworkFlagRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.UpdateNetworkFlagRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.UpdateNetworkFlagRequest.displayName = 'proto.UpdateNetworkFlagRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.UpdateNetworkFlagRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.UpdateNetworkFlagRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.UpdateNetworkFlagRequest} msg The msg instance to transform.
 * @return {!Object}
 */
proto.UpdateNetworkFlagRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    networkid: msg.getNetworkid(),
    value: msg.getValue()
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.UpdateNetworkFlagRequest}
 */
proto.UpdateNetworkFlagRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.UpdateNetworkFlagRequest;
  return proto.UpdateNetworkFlagRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.UpdateNetworkFlagRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.UpdateNetworkFlagRequest}
 */
proto.UpdateNetworkFlagRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setNetworkid(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setValue(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.UpdateNetworkFlagRequest} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.UpdateNetworkFlagRequest.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.UpdateNetworkFlagRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.UpdateNetworkFlagRequest.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getNetworkid();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = this.getValue();
  if (f) {
    writer.writeBool(
      2,
      f
    );
  }
};


/**
 * Creates a deep clone of this proto. No data is shared with the original.
 * @return {!proto.UpdateNetworkFlagRequest} The clone.
 */
proto.UpdateNetworkFlagRequest.prototype.cloneMessage = function() {
  return /** @type {!proto.UpdateNetworkFlagRequest} */ (jspb.Message.cloneMessage(this));
};


/**
 * optional string networkId = 1;
 * @return {string}
 */
proto.UpdateNetworkFlagRequest.prototype.getNetworkid = function() {
  return /** @type {string} */ (jspb.Message.getFieldProto3(this, 1, ""));
};


/** @param {string} value  */
proto.UpdateNetworkFlagRequest.prototype.setNetworkid = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional bool value = 2;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.UpdateNetworkFlagRequest.prototype.getValue = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldProto3(this, 2, false));
};


/** @param {boolean} value  */
proto.UpdateNetworkFlagRequest.prototype.setValue = function(value) {
  jspb.Message.setField(this, 2, value);
};


goog.object.extend(exports, proto);
