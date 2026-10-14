package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint8; F2 int16 }
type GenS4 struct { F0 bool }
func genFn0() uint32 {
    var v7 uint32 = 491
    if (!false) {
        v7 = (v7 ^ v7)
    }
    var v8 int = 0
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int16, p1 uint8) int64 {
    if true {
        switch "" {
        case "foo bar":
            v9 := []int{32767, -392}
            _ = v9
        case "\n\t":
            v10 := GenS1{F0: 65535}
            _ = v10
        }
        v11 := []int{42, 42}
        _ = v11
    } else {
        var v12 uint8 = (p1 & p1)
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 float64 = 0.1
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    p1 = p1
    var v15 string = "hello"
    var v16 bool = (0 < 4294967295)
    v17 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    _ = v17
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []uint32{(v1 + (v1 << ((514) & 63))), v1}
    var v3 uint16 = 1
    var v4 int16 = 42
    var v5 uint16 = (v3 / ((v3) | 1))
    var v6 bool = (true || false)
    v18 := GenS1{F0: v1}
    var v19 uint = 2
    var v20 string = "foo bar"
    var v21 float32 = ((float32(3.14159) / float32(3.14159)) / 1.0)
    v22 := GenS2{F0: v20, F1: v1}
    var v23 int8 = 0
    v6 = ((v6 && true) || ((((-2.25 / -2.25) - -2.25) / ((0.1 - 3.14159) + ((1.5 - float64(v21)) + 3.14159))) <= 0.0))
    v24 := GenS2{F0: "a", F1: v1}
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 6; fi26++ {
            v3 -= (v5 << ((v19) & 63))
            var v27 int32 = 1
            _ = fi26
            _ = v27
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v28 := GenS3{F0: (v19 | v19), F1: 1, F2: 32767}
        _ = fi25
        _ = v28
    }
    var v29 uint32 = (^(v1))
    v20 = (v20 + ("foo bar" + v20))
    var v30 uint16 = (v3 << ((42) & 63))
    v31 := GenS4{F0: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
