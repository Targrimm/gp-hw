package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func terminatableInput(input In, done In) In {
	if done == nil {
		return input
	}

	result := make(Bi)

	go func() {
		defer close(result)

		select {
		case <-done:
			go func() {
				for range input {
				}
			}()
			return
		case value, ok := <-input:
			if !ok {
				return
			}
			result <- value
		}
	}()

	return result
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		out = stage(terminatableInput(out, done))
	}

	return out
}
