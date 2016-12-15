package main

import "flag"

// TODO
// security??
// have some sort of chan of percentage
// trigger start, either via a handler or something
// get results, either via a handler or something

var (
	port        = flag.Int("port", 0, "port for the server")
	cpuInterval = flag.Int("cpuInterval", 0, "Sampling interval, if 0 is given it will compare the current cpu times against the last call")
	perCPU      = flag.Bool("perCpu", false, "Percent calculates the percentage of cpu used either per CPU or combined")
)

func main() {
	flag.Parse()
	//http.HandleFunc("/start", start)         // set router
	//err := http.ListenAndServe(":9090", nil) // set listen port
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}
