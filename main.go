package main 
import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"github.com/gdamore/tcell"
)

func main(){
	// fmt.Println("hello world");

	screen,err = tcell.Newscreen()
	if err != nil{
		log.Fatalf("%+v",err)
	}

	if err = screen.Init(); err != nil {
		log.fatalF("%+v",err)
	}

	

	fmt.Print("done")
}