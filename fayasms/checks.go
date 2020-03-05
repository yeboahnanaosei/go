package fayasms

import (
	"fmt"
)


// checkMandatoryFields checks to ensure that the mandatory fields are set
func (f *FayaSMS) checkMandatoryFields(mandatoryFields []string) error {
	for _, field := range mandatoryFields {
		if f.payload.Get(field) == "" {
			return fmt.Errorf("fayasms: a mandatory field has not been set. please supply all mandatory fields which are: %v", mandatoryFields)
		}
	}

	return nil
}


// checkContingentFields checks that all contingent fields required by endpoint are set
func (f *FayaSMS) checkContingentFields(endpoint string, contingentFields map[string][]map[string]string) error {
	fields, ok := contingentFields[endpoint]

	// Some endpoints do not have any contingent fields
	if !ok {
		return nil
	}

	for _, field := range fields {
		if f.payload.Get(field["name"]) == "" {
			return fmt.Errorf("fayasms: %v", field["errMsg"])
		}
	}

	return nil
}

