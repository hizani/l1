// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {

	// Получаем количество воркеров
	if len(os.Args) < 2 {
		fmt.Println("no arguments")
		return
	}
	workerCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// инициализируем отменяемый контектст для корректной обработки программы
	ctx, cancel := context.WithCancel(context.Background())
	// инициализируем wg
	wg := sync.WaitGroup{}
	// создаем канал для данных
	jobs := make(chan int64)
	// увеличиваем счетчик wg на кол-во работников
	wg.Add(workerCount)

	// Запускаем работников
	for i := 0; i < workerCount; i++ {
		go worker(ctx, &wg, i, jobs)
	}

	// Обработка сигнала завершения программы
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, os.Interrupt)

	var count int64
	for {
		select {
		// Закрываем контекст и завершаем работу
		case <-sigint:
			close(jobs)
			cancel()
			wg.Wait()
			return
		// Передаем данные в канал
		default:
			jobs <- count
			count++
		}
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, idx int, jobs <-chan int64) {
	for {
		select {
		// Завершаем работу воркера, если поступил сигнал отмены контектса
		case <-ctx.Done():
			fmt.Println("Close worker", idx)
			wg.Done()
			return
		// выводим данные из канала
		default:
			fmt.Printf("Worker %d: %d\n", idx, <-jobs)
		}
	}
}
