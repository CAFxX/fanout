package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int16; F1 bool; F2 uint8 }
type GenS6 struct { F0 string }

func main() {
    var v1 uint64 = 0
    var v2 int = 127
    var v3 int64 = 42
    var v4 float32 = (3.14159 + 3.14159)
    v5 := GenS1{F0: (v2 / ((v2) | 1)), F1: 98}
    v6 := []int{}
    fmt.Println(v2, v3)
    v7 := GenS2{F0: 2147483647, F1: int(v3)}
    var v8 uint16 = 32767
    v9 := GenS2{F0: 2147483647, F1: (v2 % ((v2) | 1))}
    switch (v2 % ((32768) | 1)) {
    case -9223372036854775808:
        if (!false) {
            var v10 float64 = float64(v4)
            v11 := GenS3{F0: v1, F1: 2}
            v12 := GenS4{F0: v1, F1: 128}
            _ = v10
            _ = v11
            _ = v12
        } else {
            if false {
                var v13 float64 = 0.1
                _ = v13
            }
            v8 = v8
        }
    default:
        var v14 bool = false
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        v2 = (v2 + (fi15 * 4294967295))
        v16 := GenS5{F0: 356, F1: false, F2: 42}
        v17 := GenS1{F0: fi15, F1: 1}
        _ = fi15
        _ = v16
        _ = v17
    }
    if len(v6) > 0 {
        v6[0] = (v2 - int(v8))
    }
    var v18 uint8 = 2
    var v19 float32 = v4
    if true {
        v20 := GenS6{F0: "a"}
        _ = v20
    }
    var v21 string = ("a" + ("" + ("" + "hello")))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
