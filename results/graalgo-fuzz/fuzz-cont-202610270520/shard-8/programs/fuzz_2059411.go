package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int }

func main() {
    var v1 int8 = -1
    var v2 uint16 = 303
    var v3 float32 = float32((float64(1.5) * ((0.1 - 0.0) + 1e10)))
    var v4 float64 = 3.14159
    v5 := []uint32{307, 2147483647}
    var v6 uint8 = 248
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        v5 = append(v5, 65535)
        v7 := GenS1{F0: (v6 ^ 152)}
        _ = v7
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint64 = uint64(v6)
            _ = fi8
            _ = v9
        }
    } else {
        var v10 uint8 = 2
        switch 9223372036854775807 {
        case 2147483647:
            fmt.Println(v2, v10)
        case -1:
            v4 -= (0.0 * v4)
        }
        _ = v10
    }
    v11 := GenS1{F0: 128}
    v12 := GenS1{F0: 148}
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS2{F0: 42, F1: 1}
        v1 = (^(v1))
        _ = fi13
        _ = v14
    }
    var v15 uint8 = (v6 / ((v6) | 1))
    v16 := GenS2{F0: 2, F1: 0}
    if ((-128 != (v1 % (((v1 >> ((1) & 63))) | 1))) || true) {
        switch (("hello" + ("hello" + ("hello" + ("hello" + "a")))) + "a") {
        case "a":
            var v17 uint = 520
            _ = v17
        case "hello":
            v5 = append(v5, 819)
        }
    } else {
        var v18 uint8 = (v15 % (((^(v15))) | 1))
        v19 := GenS3{F0: 1}
        _ = v18
        _ = v19
    }
    v1 -= 100
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS4{F0: 1, F1: 42}
        var v22 uint8 = (v15 << ((uint(v2)) & 63))
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
