(main (print "Hello, World!") (add 1 2))

; the above code is not right
; the correct code is:

(file "add.lisp" 
    (package main)
    (import "fmt")
    (imports "fmt" "strconv")

    (def add (a int b int) (int)
        (fmt.Println (strconv.Itoa (+ a b)))
    )
    (def main ()
        
    )
)

; and heres a complex example

(file "complex.lisp"
    (package "main")
    (imports "fmt" "strconv" "math/rand")

    ; Define a recursive function to compute factorial
    (defun factorial (n int) -> (int)
        (if (<= n 1)
            1
            (* n (factorial (- n 1)))
        )
    )

    ; Define a function using a map and array
    (defun processItems ()
        (let ((scores (map string int ((1 90) (2 80) (3 70) (5 50) (8 30)))))
            (let ((values (array 1 2 3 5 8)))
                (foreach value values
                    (if (haskey scores value)
                        (fmt.Println "Found: " (get scores value))
                        (fmt.Println "Not found: " value)
                    )
                )
            )
        )
    )

    ; Use pattern matching to handle different types
    (defun typeCheck (item interface{}) 
        (match (type item)
            (int (fmt.Println "Integer: " item))
            (string (fmt.Println "String: " item))
            (default (fmt.Println "Unknown type"))
        )
    )

    ; pattern match with value
    (defun valueCheck (item interface{}) 
        (match item
            (1 (fmt.Println "One"))
            (2 (fmt.Println "Two"))
            (3 (fmt.Println "Three"))
            (default (fmt.Println "Unknown value"))
        )
    )

    ; making arrays
    ; (array type values)
    (defun makeArray (array (array int 1 2 3 5 8))
        (do 
            (fmt.Println "Array vals:")
            (foreach val array
                (fmt.Println val)
            )
        )
    )

    ; Main function to call other functions
    (defun main ()
        (do
            (println "This is the main function.")
            (let ((result (factorial 5)))
                (println "Factorial calculated:" result)
            )
            (println "Main function completed.")
        )
    )
)
