package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int16 }
func genFn0(p0 uint, p1 int) int8 {
    for fi4 := 0; fi4 < 18; fi4++ {
        if (-1 == 42) {
            p0 += (p0 % ((p0) | 1))
            v5 := GenS1{F0: 9223372036854775807, F1: 0, F2: -32768}
            _ = v5
        }
        _ = fi4
    }
    if (2 > 32767) {
        v6 := GenS1{F0: 0, F1: -1, F2: -1}
        _ = v6
    } else {
        var v7 int = p1
        _ = v7
    }
    p1 -= (p1 << ((42) & 63))
    _ = p0
    _ = p1
    return 87
}
func genFn1(p0 uint, p1 uint32) uint {
    v8 := GenS1{F0: -9223372036854775808, F1: -90, F2: 612}
    p1 = (p1 + (p1 + p1))
    if false {
        var v9 uint = 2
        _ = v9
    }
    var v10 int64 = -1
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 484
    var v2 string = "a"
    var v3 uint32 = 4294967295
    var v11 uint16 = v1
    fmt.Println(v1, v3, v11)
    for fi12 := 0; fi12 < 11; fi12++ {
        switch ("" + "a") {
        case "hello":
            v3 += (v3 >> ((255) & 63))
        case "a":
            v11 += (v1 ^ (v11 + 722))
        }
        switch fi12 {
        case -394:
            var v13 int8 = -1
            _ = v13
        }
        _ = fi12
    }
    v2 = v2
    var v14 int = 42
    var v15 uint8 = 128
    v16 := GenS1{F0: 42, F1: -2147483648, F2: 443}
    var v17 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
