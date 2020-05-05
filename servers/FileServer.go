package servers

import "io/ioutil"

// LoadFile loads the HTML file requested by user
func LoadFile(fileName string) (string, error) {
	bytes,err := ioutil.ReadFile(fileName)
	if err != nil {
		return "",err
	}
	return string(bytes),nil
}