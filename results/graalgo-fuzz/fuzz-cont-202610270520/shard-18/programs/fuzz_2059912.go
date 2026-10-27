package main

import (
    "fmt"
)


func main() {
    var v1 string = ("a" + "foo bar")
    var v2 uint16 = 42
    var v3 float64 = 3.14159
    var v4 uint8 = 255
    var v5 int16 = -32768
    fmt.Println(v5, v3)
    switch 32768 {
    case -2147483648:
        fmt.Println(v5)
    case 148:
        fmt.Println(v5, v3, v1)
    default:
        if (4294967295 > 2) {
            v1 = "\n\t"
        } else {
            if true {
                v1 = v1
            }
        }
    }
    fmt.Println(v5, v2)
    var v6 uint32 = 0
    v1 = ((("a" + v1) + "foo bar") + (v1 + (("foo bar" + (v1 + v1)) + v1)))
    v6 -= v6
    var v7 uint8 = (v4 % ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
