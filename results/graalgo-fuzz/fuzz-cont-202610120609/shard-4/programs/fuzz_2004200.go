package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }

func main() {
    var v1 uint16 = 255
    var v2 uint8 = 26
    var v3 bool = false
    v3 = false
    v4 := GenS1{F0: 705, F1: -855, F2: -589}
    fmt.Println(v3)
    if (!v3) {
        if true {
            v5 := GenS1{F0: 2147483647, F1: 1000, F2: 42}
            var v6 uint16 = 65535
            _ = v5
            _ = v6
        } else {
            v7 := []int{978, 2147483647, 128, 870}
            _ = v7
        }
    }
    var v8 string = ("foo bar" + "")
    v2 = uint8(v1)
    v9 := []int{}
    v8 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
