# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: EigenVector.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11\x45igenVector.proto\x12\x0b\x65igenvector\"\x17\n\x08RealRows\x12\x0b\n\x03row\x18\x01 \x03(\x01\"0\n\rComplexNumber\x12\x0c\n\x04real\x18\x01 \x01(\x01\x12\x11\n\timaginary\x18\x02 \x01(\x01\"6\n\x0b\x43omplexRows\x12\'\n\x03row\x18\x01 \x03(\x0b\x32\x1a.eigenvector.ComplexNumber\"9\n\x12\x45igenVectorRequest\x12#\n\x04rows\x18\x01 \x03(\x0b\x32\x15.eigenvector.RealRows\"=\n\x13\x45igenVectorResponse\x12&\n\x04rows\x18\x01 \x03(\x0b\x32\x18.eigenvector.ComplexRows2f\n\x0b\x45igenVector\x12W\n\x0eGetEigenVector\x12\x1f.eigenvector.EigenVectorRequest\x1a .eigenvector.EigenVectorResponse(\x01\x30\x01\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'EigenVector_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_REALROWS']._serialized_start=34
  _globals['_REALROWS']._serialized_end=57
  _globals['_COMPLEXNUMBER']._serialized_start=59
  _globals['_COMPLEXNUMBER']._serialized_end=107
  _globals['_COMPLEXROWS']._serialized_start=109
  _globals['_COMPLEXROWS']._serialized_end=163
  _globals['_EIGENVECTORREQUEST']._serialized_start=165
  _globals['_EIGENVECTORREQUEST']._serialized_end=222
  _globals['_EIGENVECTORRESPONSE']._serialized_start=224
  _globals['_EIGENVECTORRESPONSE']._serialized_end=285
  _globals['_EIGENVECTOR']._serialized_start=287
  _globals['_EIGENVECTOR']._serialized_end=389
# @@protoc_insertion_point(module_scope)
