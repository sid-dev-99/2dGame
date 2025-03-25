package main 
import (
	"fmt"
	"log"
	"github.com/gdamore/tcell"
)

func main(){

	screen,err := tcell.NewScreen()
	if err != nil{
		log.Fatalf("%+v",err)
	}

	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v",err)
	}

	//global var

	

	// running loop 

	isRunning := true


	for isRunning{


	screen.Clear()

	//draw something here

	screen.SetContent(10,10,'@',nil,tcell.StyleDefault)

	screen.Show()



		ev := screen.PollEvent()
		switch ev := ev.(type){
		case *tcell.EventKey:
			switch ev.Rune(){
			case 'q':
				isRunning = false
			case 'a':
				Sprite.x++

			}
		}

	}

	


	w,h := screen.Size()
	fmt.Println("size is %d\n%d", w ,h)
	//screen.Clear()


	// draw

	//screen.setContext()

	//
	//screen.show()


	fmt.Print("done")
}