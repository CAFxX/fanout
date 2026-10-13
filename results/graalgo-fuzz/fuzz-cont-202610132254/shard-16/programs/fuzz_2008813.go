package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }

func main() {
    var v1 uint8 = 42
    var v2 uint8 = 2
    var v3 int16 = -424
    switch -689 {
    case 9223372036854775807:
        if false {
            v4 := GenS1{F0: 1, F1: (v3 & -1), F2: 266}
            _ = v4
        }
    }
    fmt.Println(v2, v3, v3)
    fmt.Println(v3, v3)
    v3 = 32767
    v5 := []int{-9223372036854775808, -1}
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    v3 = (v3 % ((v3) | 1))
    fmt.Println(v1, v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 bool = (!(false && (42 != (v2 << ((18446744073709551615) & 63)))))
    var v7 int16 = (^(v3))
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
