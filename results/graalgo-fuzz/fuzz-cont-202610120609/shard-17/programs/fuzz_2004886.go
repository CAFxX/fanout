package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int64 }

func main() {
    v1 := []int{32768, 925}
    var v2 int8 = 1
    v3 := make(map[uint32]bool)
    var v4 uint8 = 42
    var v5 uint8 = (^(v4))
    fmt.Println(v2, v4)
    var v6 int16 = 42
    v7 := GenS1{F0: ("foo bar" != "hello")}
    v8 := GenS1{F0: true}
    _ = len(v1)
    fmt.Println(v4, v6, v2)
    v5 += (v5 & (^(v4)))
    if true {
        delete(v3, 4294967295)
        var v9 uint32 = 170
        _ = v9
    } else {
        v5 = v4
    }
    var v10 float32 = float32(3.14159)
    v11 := GenS1{F0: (false || true)}
    fmt.Println(v10, v5, v10)
    var v12 int16 = (v6 + (v6 | v6))
    fmt.Println(v5, v2, v6)
    v13 := GenS2{F0: 1}
    if (-2147483648 < 9223372036854775807) {
        v12 += 1000
    } else {
        v14 := GenS1{F0: (false && false)}
        switch -657 {
        case 9223372036854775807:
            v15 := GenS3{F0: -9223372036854775808, F1: 0}
            _ = v15
        case 985:
            v10 += ((v10 / 0.1) / 1.0)
        case 32767:
            var v16 string = ("hello" + "hello")
            _ = v16
        }
        _ = v14
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS3{F0: 0, F1: 9223372036854775807}
        _ = fi17
        _ = v18
    }
    v3[167] = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
