# Gemini Form Specification

## Overview

Gemini protocol is like HTTP GET method.  Gemini does not provide a way to update data on the server.  The goal is to describe protocol that can be used to update server content while staying compliant to origina Gemini + Titan specification combination.

Titan protocol provides a way to replace resource on Gemini server through use of "titan" schema in URL.  Titan allows to upload diferent types of data based on supplied mime type.

The goal is to provide both user and machine friendly data format.  This data format is called GmiForm in the document.

### GmiForm delivery

When server delivers gmiform it marks it with appropriate MIME type. The proposed value for MIME type is "text/gmiform".

Existing document like this one can embed GmiForm into the document using existing method of embedding data into the document via "Preformatted Text".  Example:

```toml gmiform
Action = "update profile"
```

Multiple forms can be included and identified by unique Action value:

```toml gmiform
Action = "upate payment data"
```

It is not a browser responsiblity to validate uniqueness of the form.  If action is omitted its default value is empty string.

### Manual GmiForm Upload Use Case

1. If browser does not support gmiform schema, then user is expected to copy content of the form and trigger Titan upload on the page using browser menu item.

2. User pastes form into upload field.

3. User updates values in the form using editor provided by the browser.

4. User Submits the form.

Unfortunately, browser implies MIME type when uploading data.  Ideally it should submit "text/gmiform" MIME type.

### Machine GmiForm Upload Use Case

GmiForm data provides deterministic model to build UI elements.  The recommended solution is a top down flow of listed components as it will work for all screen widths.

Browser has 3 levels of handling GmiForm:

1. Not handled.
2. Handled with all values assumed strings.
3. Handled with data types support.  For example, data field can be processed by the browser.

Time handling at level 3

Browser must assume that user inputs data in local time and convert it to UTC.  Date Time must be sent as Unix UTC seconds.  Unix UTC acts as an anchor time zone acorss all potential environmetns.  It is the only way to ensure proper handling of dates in the environment that does not provide information about user time zone.

## GmiForm Schema

Schema is best described through GO programming language syntax documented through file form.go in this project.  See Appendix A for copy of GO code of the schema.

### GmiForm Action Property

Action is optional.  Default value is empty string.

It is a free form string that is not required to be unique within a single document.

### Gmiform Error Property

Error is optional.  If omitted default value is empty string.

It is a free form string that is used by the server to report errors at form level.  

If there is error anywhere in the form, this property must be set. It can be the only error property set in the form.

### GmiForm Group Property

Group is optional.  If omitted its default value is empty array.

A set of fields can be grouped for user convenience.  For example, Address information consists of multiple input fields that may be grouped together.

### Group Title Property

Title is optional.  If omitted its default value is empty string.

### Group Error Property

Error is optional.  If omitted its default value is empty string.

Even if error exists in one of the fields group owns the Error field at group level may be not set.  It is provided here as an option to communicate additional instructions if such may exist.

## Group Field Property

Field is optional.  If omitted its default value is empty array.

Field contains a set of fields that describe user input data.

## Field Name Property

Name is required and cannot be empty.

Name is a field label to be displayed to the user.

## Field Value Property

Value is optional.  If omitted it default value is empty string.

Value can hold string or number data type.  It can potentially hold other data types.

At Level 2 browser implementation value is always a string.

At Level 1 and Level 3 browser implementation value type is likely to match type declared in ValueType property.

### Field Description Property

Description is optional.  If omitted its default value is empty string.

Description is meant to provide instructions to the user about the field.

### Field ValueType Property

ValueType is optional.  If omitted its default value is empty string.

Valid value types values:

* '' is empty string and shall be treated as string type
* 'int' is integer
* 'float' is a floating number
* 'double' is a double precision floating number
* 'unixsec' is a date time as a 64-bit number in Unix UTC seconds.

### Field Error Property

Error is optional.  Default value is empty string.

Error provides eror message as it applies to the field.

## Discussion

### Why TOML?

There are currently 2 major human friendly data formats:

* YAML
* TOML

YAML is easier to read, because it requies indentation.  That indentation is also what makes it not a well suitable format for average user that is unaware of indentation impact.

YAML is complicated and due to its complexity error messages are more likely to be not obvious to the user.

TOML is simple and does not require indentation.  Its error messages are supposed to be more clear when there is a syntax problem.

### Why one level of grouping

Infininetly nested groups create UI presentation problem.

### Why no required field?

With no option to validate user input on the client side there is not much point to do partial validation.

### Why validation of data type?

Only Level 3 requires data type validation.  The goal is to help user with data input more than just provide validation rule.

For example, date entry could benefit from Date control.

Reducing user input errors improves overall experience.  

For example, input control can restrict input to just numeric characters when number is expected.  It has nothing to do with final validation, which still has to occur on the server side.

### Why no choice

Choice might be difficult to present for easy human consumption.

### Why no Yes/No Data Type

It probably should be added, but complexity of its values description keeps it out.

It is recommended to use numbers 0 and 1 or String with Y and N values as a replacement.

### Why is Date Time in Unix UTC Seconds?

Because browser does not provide server with local time.  Just about any project out there has to deal with UTC server time and conversion of it to Browser's local time.  

Some deal with it using user profile configuration, but that requires user profile.

If date or date time is passed as a string one has to define which date format to use.

Does it mean that Level 0 requires user to input time in Unix UTC seconds?  No.  It means that until browser can support the format, users will be entering date time as a string and Description proprty will contain instructions which format to use.

64-bit integer is one of the smallest data formats to hold date time within resolution that makes sense for humans.

## Examples

## Example with Indentation

```toml gmiform
Action = "update profile"
Error = "There is 1 error to correct"

[[Group]]
  Title = "Personal Information"

  [[Group.Field]]
    Name = "First Name"
    Value = "John"
    ValueType = "int"

  [[Group.Field]]
    Name = "Last Name"
    Value = ""
    Error = "Last name is required"

[[Group]]
  Title = "Address"

  [[Group.Field]]
    Name = "Street 1"
    Value = ""

  [[Group.Field]]
    Name = "ZIP"
    Value = 99019
```

## Example without indentation

```toml gmiform
Action = "update profile"
Error = "There is 1 error to correct"

[[Group]]
Title = "Personal Information"

[[Group.Field]]
Name = "First Name"
Value = "John"
ValueType = "int"

[[Group.Field]]
Name = "Last Name"
Value = ""
Error = "Last name is required"

[[Group]]
Title = "Address"

[[Group.Field]]
Name = "Street 1"
Value = ""

[[Group.Field]]
Name = "ZIP"
Value = 99019
```

## Appendix A: Go Langauge Schema

```go
package gmif

type Form struct {
  Action string  `toml:",omitempty"`
  Group  []Group `toml:",omitempty"`
  Error  string  `toml:",omitempty"`
}

type Group struct {
  Title string  `toml:",omitempty"`
  Field []Field `toml:",omitempty"`
  Error string  `toml:",omitempty"`
}

type DataType string

const (
  String   DataType = ""
  Int      DataType = "int"
  Float    DataType = "float"
  Double   DataType = "double"
  UnixSecs DataType = "unixsec"
)

type Field struct {
  Name        string
  Value       interface{}
  Description string   `toml:",omitempty"`
  ValueType   DataType `toml:",omitempty"`
  Error       string   `toml:",omitempty"`
}
```