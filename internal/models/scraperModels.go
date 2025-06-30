package models

type JobListing struct {
	Fields []string `json:"fields"`
}

func (j *JobListing) Hash() string {
	hashVal := ""
	for i, v := range j.Fields {
		hashVal += v
		if i == len(j.Fields)-1 || i == 2 {
			break
		}
	}
	return hashVal
}

func (j *JobListing) String() string {
	formatted := ""
	for i, v := range j.Fields {
		if i == len(j.Fields)-1 {
			formatted += v
		} else {
			formatted += (v + " - ")
		}
	}

	return formatted
}

func (j *JobListing) String_NameOnly() string {
	return j.Fields[0]
}
