# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: median.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0cmedian.proto\x12\x06median\"\x19\n\x07Request\x12\x0e\n\x06number\x18\x01 \x01(\x05\"\x1a\n\x08Response\x12\x0e\n\x06median\x18\x01 \x01(\x01\x32\x46\n\x11\x46indMedianService\x12\x31\n\nget_median\x12\x0f.median.Request\x1a\x10.median.Response(\x01\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'median_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_REQUEST']._serialized_start=24
  _globals['_REQUEST']._serialized_end=49
  _globals['_RESPONSE']._serialized_start=51
  _globals['_RESPONSE']._serialized_end=77
  _globals['_FINDMEDIANSERVICE']._serialized_start=79
  _globals['_FINDMEDIANSERVICE']._serialized_end=149
# @@protoc_insertion_point(module_scope)