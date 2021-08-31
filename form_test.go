package gmif_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/kulak/gmif"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	form := gmif.Form{
		Action: "update profile",
		Error:  "There is 1 error to correct",
		Group: []gmif.Group{
			{
				Title: "Personal Information",
				Field: []gmif.Field{
					{
						Name:      "First Name",
						Value:     "John",
						ValueType: "int",
					},
					{
						Name:  "Last Name",
						Value: "",
						Error: "Last name is required",
					},
				},
			},
			{
				Title: "Address",
				Field: []gmif.Field{
					{
						Name:  "Street 1",
						Value: "",
					},
					{
						Name:  "ZIP",
						Value: 99019,
					},
				},
			},
		},
	}
	var buf bytes.Buffer
	var enc = toml.NewEncoder(&buf)
	err := enc.Encode(&form)
	require.NoError(t, err)
	fmt.Println(buf.String())

	buf.Reset()
	enc.Indent = ""
	err = enc.Encode(&form)
	require.NoError(t, err)
	fmt.Println(buf.String())
}

func TestUnmarshal(t *testing.T) {
	body := []byte(`Action = 'update profile'
	[[Group]]
	Title = 'Personal Information'
	[[Group.Field]]
	Name = 'First Name'
	Value = "John"
	Description = ''
	[[Group.Field]]
	Name = 'Last Name'
	Value = ''
	Description = ''
	
	[[Group]]
	Title = 'Address'
	[[Group.Field]]
	Name = 'Street 1'
	Value = ''
	Description = ''
	[[Group.Field]]
	Name = "ZIP"
	Value = 99019
	Description = ''
	Validator = 'field is not defined'
	`)
	var cnt gmif.Form
	err := toml.Unmarshal(body, &cnt)
	require.NoError(t, err)
	require.Equal(t, "John", cnt.Group[0].Field[0].Value)
	require.Equal(t, int64(99019), cnt.Group[1].Field[1].Value)

	var ok bool
	_, ok = cnt.Group[0].Field[0].Value.(string)
	require.True(t, ok)
	_, ok = cnt.Group[0].Field[0].Value.(int64)
	require.False(t, ok)
	_, ok = cnt.Group[1].Field[1].Value.(int64)
	require.True(t, ok)
	_, ok = cnt.Group[1].Field[1].Value.(string)
	require.False(t, ok)
}
