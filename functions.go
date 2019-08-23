package main

func sa(A, B *Stack) {
	A.swap()
}

func sb(A, B *Stack) {
	B.swap()
}

func ss(A, B *Stack) {
	A.swap()
	B.swap()
}

func pa(A, B *Stack) {
	A.push(B.pop())
}

func pb(A, B *Stack) {
	B.push(A.pop())
}

func ra(A, B *Stack) {
	A.rotateUp()
}

func rb(A, B *Stack) {
	B.rotateUp()
}

func rr(A, B *Stack) {
	A.rotateUp()
	B.rotateUp()
}

func rra(A, B *Stack) {
	A.rotateDown()
}

func rrb(A, B *Stack) {
	B.rotateDown()
}

func rrr(A, B *Stack) {
	A.rotateDown()
	B.rotateDown()
}

func check(A, B *Stack) {

}
