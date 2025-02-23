package study_case

import (
	"fmt"
	"sync"
	"time"
)

type FileDownload struct {
	id int
	name string
	size int
	duration time.Duration
}

type DownloadManager struct {
	downloadQueue chan FileDownload
	wg sync.WaitGroup
}

func NewDownloadManager() *DownloadManager {
	return &DownloadManager{
		downloadQueue: make(chan FileDownload, 5),
	}
}

func (dm *DownloadManager) downloadWithGoroutine(workerID int) {
	defer dm.wg.Done()
	for file := range dm.downloadQueue {
		fmt.Printf("[Goroutine-%d] Mulai download %s (%dMB)\n",
		workerID, file.name, file.size)

		time.Sleep(file.duration * time.Second)

		fmt.Printf("[Goroutine-%d] Selesai download %s\n", workerID, file.name)
	}
}

func downloadSequential(files []FileDownload) time.Duration {
	start := time.Now()

	for _, file := range files {
		fmt.Printf("[Sequential] Mulai download %s (%dMB)\n", 
			file.name, file.size)
		
		time.Sleep(file.duration * time.Second)
		
		fmt.Printf("[Sequential] Selesai download %s\n", 
			file.name)
	}

	return time.Since(start)
}

func downloadWithGoroutines(files []FileDownload) time.Duration {
	start := time.Now()
	dm := NewDownloadManager()

	numWorkers := 3
	dm.wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go dm.downloadWithGoroutine(i)
	}

	for _, file := range files {
		dm.downloadQueue <- file
	}

	close(dm.downloadQueue)
	dm.wg.Wait()

	return time.Since(start)
}

func GoRoutinesCases() {
	files := []FileDownload {
		{1, "movie.mp4", 780, 4},
		{2, "taufik.zip", 800, 4},
		{3, "Windows.zip", 800, 4},
	}

	seqTime := downloadSequential(files)

	goroutineTime := downloadWithGoroutines(files)

	fmt.Printf("Waktu sequential: %v\n", seqTime)
	fmt.Printf("Waktu dengan Goroutine: %v\n", goroutineTime)
	fmt.Printf("Selisih Waktu: %v\n", seqTime-goroutineTime)

}