package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint = 296
    var v2 uint8 = 128
    var v3 uint8 = 0
    var v4 string = "hello"
    v5 := make(map[int]string)
    delete(v5, 128)
    v3 += v2
    var v6 float64 = 0.0
    v4 = "hello"
    if true {
        if false {
            v7 := GenS1{F0: -241}
            _ = v7
        }
        fmt.Println(v4, v3)
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 string = ("\n\t" + v4)
            _ = fi8
            _ = v9
        }
    } else {
        if (v1 == v1) {
            fmt.Println(v1)
        }
    }
    switch v4 {
    case "":
        if true {
            var v10 string = (v4 + v4)
            var v11 uint8 = (v2 | 15)
            _ = v10
            _ = v11
        } else {
            v12 := []int{128, 913, -151}
            _ = v12
        }
    case "\n\t":
        var v13 int8 = -128
        _ = v13
    case "a":
        v14 := GenS2{F0: -302}
        _ = v14
    }
    var v15 uint8 = (v3 + (v2 >> (((v1 - v1)) & 63)))
    v16 := []int{}
    fmt.Println(v6, v4, v4)
    fmt.Println(v4, v4, v3)
    var v17 uint8 = 1
    var v18 int16 = 194
    var v19 int16 = v18
    switch -2147483648 {
    case 1:
        v1 = 4294967295
    }
    v2 = 2
    var v20 uint = (v1 << (((v1 / (((v1 / ((42) | 1))) | 1))) & 63))
    if false {
        if true {
            v16 = append(v16, 42)
        }
    } else {
        var v21 uint64 = 2
        var v22 uint8 = 127
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
