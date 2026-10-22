package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int32; F2 bool }

func main() {
    var v1 string = "a"
    var v2 int16 = 1
    var v3 int32 = 1
    var v4 uint32 = 2147483647
    var v5 int8 = -128
    switch v1 {
    case "a":
        fmt.Println(v5, v4, v2)
    case "\n\t":
        var v6 uint8 = 176
        _ = v6
    default:
        v7 := GenS1{F0: v4, F1: 0}
        _ = v7
    }
    var v8 int8 = (^(v5))
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v1)
        _ = fi9
    }
    fmt.Println(v5, v5)
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(v2, v3, v2)
        if false {
            v4 -= 791
            var v11 uint8 = 255
            _ = v11
        }
        _ = fi10
    }
    v12 := GenS1{F0: 782, F1: 255}
    if false {
        if true {
            v2 = 42
            v13 := []int{}
            v2 -= 32767
            _ = v13
        }
        if false {
            if (42 == 18446744073709551615) {
                v14 := []int{798, 42}
                for fi15 := 0; fi15 < 16; fi15++ {
                    v4 = v4
                    _ = fi15
                }
                var v16 uint16 = 1
                _ = v14
                _ = v16
            }
        }
    }
    v17 := []int{-756, -349, 9223372036854775807}
    v18 := GenS2{F0: 42, F1: v3, F2: false}
    var v19 uint16 = 1
    v17 = append(v17, 284)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
