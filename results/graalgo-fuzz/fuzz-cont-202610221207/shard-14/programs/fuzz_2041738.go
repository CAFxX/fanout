package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint; F2 int32 }
type GenS4 struct { F0 uint16; F1 int; F2 uint8 }

func main() {
    var v1 int8 = 100
    v2 := []uint32{0}
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    v5 := []string{"hello"}
    if true {
        v6 := GenS1{F0: 82, F1: "\n\t"}
        var v7 int16 = -32768
        _ = v6
        _ = v7
    }
    v1 = v1
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := v4[fi8]
        fmt.Println(fi8)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        v1 -= v1
        v11 := GenS1{F0: 42, F1: ""}
        var v12 int32 = 42
        _ = fi10
        _ = v11
        _ = v12
    }
    switch ("hello" + "a") {
    case "foo bar":
        v1 += v1
    default:
        var v13 bool = true
        _ = v13
    }
    switch 11 {
    case 9223372036854775807:
        v14 := GenS2{F0: 0}
        _ = v14
    }
    if false {
        var v15 uint32 = 4294967295
        if false {
            v1 += v1
            v15 = v15
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v15
    }
    v1 = (^(v1))
    var v16 float64 = 3.14159
    fmt.Println(v1)
    _ = len(v2)
    _ = len(v3)
    if (100 != (v1 >> ((18446744073709551615) & 63))) {
        v17 := GenS3{F0: 4294967295, F1: 255, F2: 0}
        v16 += (v16 / v16)
        v16 -= v16
        _ = v17
    } else {
        v18 := GenS1{F0: 9223372036854775807, F1: ("" + "foo bar")}
        v1 = (v1 ^ (v1 & (v1 - (-(v1)))))
        var v19 int = -2147483648
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        v16 = ((((v16 - -2.25) / (3.14159 / 1.0)) / v16) + 3.14159)
        _ = fi20
    }
    var v21 int = -2147483648
    v16 -= ((v16 - (v16 * -2.25)) * v16)
    fmt.Println(v1, v21)
    v22 := GenS4{F0: 0, F1: -41, F2: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
