package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 bool; F2 bool }
type GenS3 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 bool = true
    var v2 uint = 209
    var v3 uint = (v2 * (v2 / ((v2) | 1)))
    var v4 float64 = -2.25
    v5 := []uint32{4294967295}
    if (284 < 42) {
        var v6 string = "\n\t"
        v7 := GenS1{F0: -9223372036854775808}
        _ = v6
        _ = v7
    } else {
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    }
    var v8 int8 = -128
    v9 := GenS2{F0: (((v1 && true) && v1) && false), F1: v1, F2: (((1e10 / 3.14159) + 1.5) >= (float32(1e10) - 1e10))}
    v10 := GenS1{F0: 0}
    for fi11 := 0; fi11 < 10; fi11++ {
        v8 = (v8 & v8)
        _ = fi11
    }
    var v12 int32 = 141
    switch ("foo bar" + "") {
    case "":
        v12 = (v12 - 986)
    case "foo bar":
        var v13 uint16 = 292
        _ = v13
    case "\n\t":
        v12 -= 2147483647
    }
    var v14 uint16 = 42
    fmt.Println(v8, v8, v3)
    v15 := GenS3{F0: 42, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
