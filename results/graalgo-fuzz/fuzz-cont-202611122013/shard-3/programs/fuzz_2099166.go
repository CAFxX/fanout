package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 bool }

func main() {
    var v1 uint32 = 65535
    var v2 uint = 4294967295
    v3 := []int{86}
    var v4 uint = (v2 << ((v2) & 63))
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = (v1 + (v1 >> (((v4 - 65535)) & 63)))
        _ = len(v3)
    } else {
        var v5 int8 = 127
        switch "hello" {
        case "":
            switch "\n\t" {
            case "\n\t":
                v2 += 0
            case "a":
                var v6 int = 113
                _ = v6
            }
        }
        fmt.Println(v5, v4)
        _ = v5
    }
    _ = len(v3)
    var v7 uint = 2
    var v8 float32 = 3.14159
    if len(v3) > 0 {
        v3[0] = -167
    }
    var v9 float64 = 0.0
    var v10 uint64 = 17
    if (0 >= 272) {
        var v11 float32 = float32(0.1)
        v12 := GenS1{F0: (^(v1)), F1: 271, F2: (!false)}
        _ = v11
        _ = v12
    }
    v1 = (v1 & v1)
    if true {
        v1 = (^(v1))
        fmt.Println(v10, v4, v9)
        var v13 bool = (!(false || false))
        _ = v13
    } else {
        var v14 bool = true
        var v15 int8 = -128
        _ = v14
        _ = v15
    }
    var v16 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
