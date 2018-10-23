# protoc-gen-atlas-query-validate

### Purpose

A [protobuf](https://developers.google.com/protocol-buffers/) compiler plugin 
designed to simplify validation of [Atlas](https://github.com/infobloxopen/atlas-app-toolkit)
[gRPC](https://grpc.io/) List [query](https://github.com/infobloxopen/atlas-app-toolkit/blob/master/query/collection_operators.proto) parameters
by generating .pb.atlas.query.validate.go files with validation rules and functions.
Currently query.Filtering, query.Sorting and query.FieldSelection are supported for validation.

### Prerequisites

#### 1. Protobuf Compiler

The protobuf compiler (protoc) is required.

[Official instructions](https://github.com/google/protobuf#protocol-compiler-installation)

[Abbreviated version](https://github.com/grpc-ecosystem/grpc-gateway#installation)

#### 2. Golang Protobuf Code Generator

Get the golang protobuf code generator:

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

#### 3. Vendored Dependencies

Retrieve and install the vendored dependencies for this project with [dep](https://github.com/golang/dep):

```
dep ensure
```

### Installation

To use this tool, install it from code with `make install`, `go install` directly,
or `go get github.com/infobloxopen/protoc-gen-atlas-query-validate`.

### Usage

Once installed, the `atlas-query-validate_out=.` or `--atlas-query-validate_out=${GOPATH}src`
option can be specified in a protoc command to generate the .pb.atlas.query.validate.go files.

#### Validation rules

Validation rules are generated for each gRPC method containing query.Filtering, query.Sorting or query.FieldSelection in it's request message
based on the message included in the method's response message(will call it *resource message* for the rest of the document).

By default all fields of *resource message` are allowed for filtering/sorting.

The list of allowed filtering operators and filtering value type/condition type depends on the *value_type* of the field,
which is either taken from `(atlas.query.validate).value_type` proto field option or is computed by the
plugin based on the field type if the option is not supplied. Currently *value_type* can be either STRING or NUMBER.
The following table shows what is allowed for each *value_type*:

|                                     | STRING | NUMBER |
|-------------------------------------|--------|--------|
| **Filtering operators**                 | EQ, MATCH, GT, GE, LT, LE, IEQ, IN | EQ, GT, GE, LT, LE, IN | 
| **Filtering value type/condition type** | String, null/StringCondition, NullCondition, StringArray(only for IN)| Number, null/NumberCondition, NullCondition, NumberArray(only for IN)|

The next table shows how *value_type* is computed from a proto field type:

| Proto field type            | value_type |
|-----------------------------|---------------|
| enum                        | STRING        |
| string                      | STRING        |
| double                      | NUMBER        |
| float                       | NUMBER        |
| int32                       | NUMBER        |
| int64                       | NUMBER        |
| sint32                      | NUMBER        |
| sint64                      | NUMBER        |
| uint32                      | NUMBER        |
| uint64                      | NUMBER        |
| google.protobuf.StringValue | STRING        |
| google.protobuf.DoubleValue | NUMBER        |
| google.protobuf.FloatValue  | NUMBER        |
| google.protobuf.Int32Value  | NUMBER        |
| google.protobuf.Int64Value  | NUMBER        |
| google.protobuf.UInt32Value | NUMBER        |
| google.protobuf.UInt64Value | NUMBER        |
| google.protobuf.Timestamp   | STRING        |
| gorm.types.UUID             | STRING        |
| gorm.types.UUIDValue        | STRING        |
| atlas.rpc.Identifier        | STRING        |
| gorm.types.InetValue        | STRING        |

#### Validation functions

Validation functions are the entry points for the generated functionality.
The following validation functions are generated:

```golang
func {Proto_file_name}ValidateFiltering(methodName string, f *query.Filtering) error
```

```golang
func {Proto_file_name}ValidateSorting(methodName string, s *query.Sorting) error
```

```golang
func {Proto_file_name}ValidateFieldSelection(methodName string, s *query.FieldSelection) error
```

Non-nil `error` is returned by the functions if validation is not passed.


### Customization

Currently only field-level proto options are supported as customization means. We're planning to add method-level options which will override
field-level options in order to support different validation rules for List methods having the same *resource message*.

* In order to disable sorting for a field set `(atlas.query.validate).sorting.disable` option to `true`.
```golang
bool on_vacation = 3 [(atlas.query.validate).sorting.disable = true];
```

* In order to customize the list of allowed filtering operators pass either a set of `(atlas.query.validate).filtering.allow` or 
a set of `(atlas.query.validate).filtering.deny` options.
  - In case of using `(atlas.query.validate).filtering.allow` only specified filtering operators are allowed:
    ```golang
    string first_name = 1 [(atlas.query.validate).filtering = {allow: MATCH, allow: EQ}];
    ```
  - In case of using `(atlas.query.validate).filtering.deny` all appropriate(for the field type) filtering operators except specified ones are allowed:
    ```golang
    string first_name = 1 [(atlas.query.validate).filtering = {deny: GT, deny: GE, deny: LT, deny: LE}];
    ```
* In order to change default *value_type* for a field pass `(atlas.query.validate).value_type` option.
```golang
CustomType custom_type_string = 10 [(atlas.query.validate) = {value_type: STRING}];
```
* In order to enable filtering/sorting by nested fields set `(atlas.query.validate).enable_nested_fields` option to true
on the field of a message type.

```golang
message User {
    Address home_address = 11 [(atlas.query.validate) = {enable_nested_fields: true}];
    Address work_address = 12;
}

message Address {
    string city = 1 [(atlas.query.validate) = {filtering: {allow: EQ}, sorting: {disable: true}}];
    string country = 2;
}
```


### Examples

The best way to get started with the plugin is to check out our [example](example/example.proto).
Example .proto files and generated .pb.atlas.query.validate.go demonstrate most of the use cases of the plugin.

Running `make example` will recompile all these test proto files, if you want
to test the effects of changing the options and fields.

### Limitations

This project is currently in development, and is expected to undergo "breaking"
(and fixing) changes
