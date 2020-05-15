package main

import "fmt"

func main() {
	original := Thing{"0"}
	fmt.Printf("no without pointers\n")
	noC(noB(noA(original)))
	fmt.Printf("Now original == %v\n", original.name)

	fmt.Printf("\nReturn pointers in chain\n")
	original = Thing{"0"}
	rpC(*rpB(*rpA(original)))
	fmt.Printf("Now original == %v\n", original.name)

	fmt.Printf("\nTake pointers in chain\n")
	original = Thing{"0"}

	a := tpA(&original)
	b := tpB(&a)
	tpC(&b)
	fmt.Printf("Now original == %v\n", original.name)

	fmt.Printf("\n Both pointers in chain\n")
	original = Thing{"0"}
	x := ppB(ppA(&original))
	x.name = "INJECT"
	tpC(x)
	fmt.Printf("Now original == %v\n", original.name)

	fmt.Printf("\n Take pointers as expected\n")
	original = Thing{"0"}
	tpA(&original)
	b = tpB(&original)
	b.name = "NOTICE_I_DONT_SHOW_UP_NEXT"

	tpC(&original)
	fmt.Printf("Now original == %v\n", original.name)

}

//Thing with a name
type Thing struct {
	name string
}

//No pointers
func (t *Thing) no(str string) Thing {
	fmt.Printf("%v -> %v\n", t.name, str)
	t.name = str
	return *t
}

//Return pointer
func (t *Thing) rp(str string) *Thing {
	fmt.Printf("%v -> %v\n", t.name, str)
	t.name = str
	return t
}

//Take Pointer
func (t *Thing) tp(str string) Thing {
	fmt.Printf("%v -> %v\n", t.name, str)
	t.name = str
	return *t
}

//Both
func pp(t *Thing, str string) *Thing {
	fmt.Printf("%v -> %v\n", t.name, str)
	t.name = str
	return t
}

func noA(t Thing) Thing { return t.no("A") }
func noB(t Thing) Thing { return t.no("B") }
func noC(t Thing) Thing { return t.no("C") }

func rpA(t Thing) *Thing { return t.rp("A") }
func rpB(t Thing) *Thing { return t.rp("B") }
func rpC(t Thing) *Thing { return t.rp("C") }

func tpA(t *Thing) Thing { return t.tp("A") }
func tpB(t *Thing) Thing { return t.tp("B") }
func tpC(t *Thing) Thing { return t.tp("C") }

func ppA(t *Thing) *Thing { return pp(t, "A") }
func ppB(t *Thing) *Thing { return pp(t, "B") }
func ppC(t *Thing) *Thing { return pp(t, "C") }
