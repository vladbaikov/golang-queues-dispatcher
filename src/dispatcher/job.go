package dispatcher

import (
	"fmt"
	"log"
)

type Job struct {
	Data string
}

func (job *Job) Handler() error {
	// Add your handler here

	if len(job.Data) == 0 {
		return fmt.Errorf("Data is empty")
	}

	log.Println("Data:", job.Data)
	return nil
}
