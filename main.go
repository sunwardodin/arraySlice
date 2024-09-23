package main

import "fmt"

func main() {

	/*
								// Array literal
								a := [3]int{42, 43, 44}
								fmt.Println(a)

								b := [...]string{"Hello", "Gophers"}
								fmt.Println(b)

								var c [2]int
								c[0] = 7
								c[1] = 8
								fmt.Println(c)

							// This is a slice
							xs := []string{"hello", "world"}
							for _, v := range xs {
								fmt.Printf("%v\n", v)
							}
							fmt.Println(len(xs))
							fmt.Println("--------------------")

						// This is appending to the slice
						xs = append(xs, "!")
						fmt.Println(xs)

					// This is slicing a slice
					xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
					fmt.Printf("xi - %#v\n", xi)
					fmt.Println("--------------------")

					// [inclusive:exclusive]
					fmt.Printf("xi - %#v\n", xi[1:5])
					fmt.Println("--------------------")

					// [inclusive:]
					fmt.Printf("xi - %#v\n", xi[2:])
					fmt.Println("--------------------")

					// [:exclusive]
					fmt.Printf("xi - %#v\n", xi[:5])
					fmt.Println("--------------------")

				// Deleting from a slice
				xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				fmt.Printf("xi - %#v\n", xi)
				fmt.Println("--------------------")

				xi = append(xi[:4], xi[5:]...)
				fmt.Printf("xi - %#v\n", xi)
				fmt.Println("--------------------")

			// make function
			xi2 := make([]int, 0, 10)
			fmt.Println(xi2)
			fmt.Println(len(xi2))
			fmt.Println(cap(xi2))
			xi2 = append(xi2, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
			fmt.Println(xi2)
			fmt.Println("--------------------")
			xi2 = append(xi2, 10, 11, 12, 13)
			fmt.Println(xi2)
			fmt.Println(len(xi2))
			fmt.Println(cap(xi2))

		// Multi dimensional slice
		jb := []string{"J", "B", "M", "C"}
		jm := []string{"M", "P", "G", "V"}
		fmt.Println(jb, jm)

		xxs := [][]string{jb, jm}
		fmt.Println(xxs)
	*/

	a := []int{0, 1, 2, 3, 4, 5}
	b := a

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("---------------")

	a[0] = 7

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("---------------")

	/*
		{
			var ni [7]int
			ni[0] = 42
			fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

			// Array literal
			ni2 := [4]int{55, 56, 57, 58}
			fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

			ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
			fmt.Printf("%#v \t %T\n", ns, ns)

			fmt.Println(len(ni))
			fmt.Println(len(ni2))
			fmt.Println(len(ns))
		}
	*/
}
