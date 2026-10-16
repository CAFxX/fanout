package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{-2147483648, 32767}
    v3 := []int{-9223372036854775808, 736}
    v4 := []string{}
    var v5 float32 = -2.25
    var v6 float32 = v5
    v7 := GenS1{F0: 0, F1: -128, F2: 1}
    var v8 string = "\n\t"
    var v9 uint64 = 140
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 int8 = 127
            fmt.Println(v5)
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    var v13 string = ""
    var v14 int64 = v1
    var v15 uint8 = 0
    switch "hello" {
    case "":
        if len(v2) > 0 {
            v2[0] = 373
        }
    case "hello":
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 bool = true
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v15 = v15
        v8 = ""
        _ = fi18
    }
    fmt.Println(v6, v8, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
