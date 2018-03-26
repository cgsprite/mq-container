/*
© Copyright IBM Corporation 2018

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"os"
	"path"
	"text/template"
)

// processTemplateFile takes a Go templateFile, and processes it with the
// supplied data, writing to destFile
func processTemplateFile(templateFile, destFile string, data interface{}) error {
	// Re-configure channel if app password not set
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Error(err)
		return err
	}
	dir := path.Dir(destFile)
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dir, 0660)
			// TODO: Lookup value for MQM user here?
			err = os.Chown(dir, 999, 999)
			if err != nil {
				log.Error(err)
				return err
			}
		} else {
			return err
		}
	}
	f, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, 0660)
	defer f.Close()
	err = t.Execute(f, data)
	if err != nil {
		log.Error(err)
		return err
	}
	// TODO: Lookup value for MQM user here?
	err = os.Chown(destFile, 999, 999)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
