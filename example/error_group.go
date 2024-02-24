package example

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func Job(jobID int) error {
	if 1 == jobID {
		return fmt.Errorf("job %v failed", jobID)
	}
	if 2 == jobID || 4 == jobID {
		time.Sleep(time.Second * 6)
	} else {
		time.Sleep(time.Second * 4)
	}

	fmt.Printf("Job %v done.\n", jobID)
	return nil
}

func ErrorGroup() {
	var eg errgroup.Group

	for i := 0; i < 10; i++ {
		jobID := i
		eg.Go(func() error {
			return Job(jobID)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println("Encountered error:", err)
	}
	fmt.Println("Successfully finished.")
}
