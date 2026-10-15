package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int }

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 964
    var v3 int32 = 11
    var v4 int32 = v3
    var v5 bool = true
    v2 = v2
    var v6 int16 = 1000
    v2 = (^(v2))
    v2 -= 4294967295
    v5 = false
    v7 := GenS1{F0: v3, F1: 32767, F2: 32767}
    var v8 bool = true
    v9 := []int{-2147483648, -9223372036854775808, 128}
    v2 -= (v2 / ((v2) | 1))
    if true {
        for fi10 := 0; fi10 < 1; fi10++ {
            switch 42 {
            case 9223372036854775807:
                if false {
                    var v11 uint8 = 162
                    fmt.Println(v5, v11, v6)
                    fmt.Println(v5)
                    _ = v11
                }
            case 4294967295:
                var v12 uint = (v2 << (((v2 / ((539) | 1))) & 63))
                _ = v12
            default:
                var v13 uint8 = 127
                _ = v13
            }
            fmt.Println(v6, v6)
            fmt.Println(v2)
            _ = fi10
        }
        switch 9223372036854775807 {
        case 128:
            if len(v9) > 0 {
                v9[0] = 42
            }
        case -1:
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int8 = -128
                v8 = v8
                _ = fi14
                _ = v15
            }
        }
    }
    switch "hello" {
    case "a":
        if len(v9) > 0 {
            v9[0] = -89
        }
    case "\n\t":
        fmt.Println(v4, v3, v2)
    case "":
        fmt.Println(v6)
    }
    v5 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
