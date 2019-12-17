// Code generated by tool/convert
// DO NOT EDIT!
package data

type container struct {
	s map[string]interface{}
}

func (c *container) Update(name string, st interface{}) error {
	switch name {

	case "Queue":
		tmp := st.(*Queue)
		UpdateQueue(tmp)

	case "Global":
		tmp := st.(*Global)
		UpdateGlobal(tmp)

	default:
		panic("unsupported struct " + name)
	}
	return nil
}

func (c *container) GetStructs() map[string]interface{} {
	return c.s
}

var ObjectsContainer = &container{
	s: map[string]interface{}{

		"Queue": Queue{},

		"Global": Global{},
	},
}
