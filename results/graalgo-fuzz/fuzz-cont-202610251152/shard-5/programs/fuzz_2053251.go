package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool; F1 string }
func genFn0() int8 {
    var v6 uint32 = 234
    v6 = v6
    fmt.Println(v6)
    _ = v6
    return -1
}

func main() {
    var v1 uint64 = 552
    v2 := []uint32{2147483647, 2, 42, uint32(v1)}
    var v3 float64 = (0.1 / (1.5 / 1e10))
    var v4 uint64 = (v1 ^ v1)
    var v5 uint = 320
    v7 := GenS1{F0: 32767, F1: 372}
    switch (("foo bar" + ("foo bar" + "")) + "\n\t") {
    case "":
        var v8 int64 = -9223372036854775808
        _ = v8
    case "hello":
        var v9 int = 127
        _ = v9
    case "foo bar":
        if ((false || (-114 == genFn0())) && true) {
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 uint32 = 0
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: -2147483648, F1: 32767}
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 uint8 = 196
        _ = fi13
        _ = v14
    }
    var v15 int32 = 1
    if false {
        if true {
            v4 -= v4
            var v16 string = ("a" + "foo bar")
            v3 = v3
            _ = v16
        }
        var v17 bool = (!((3.14159 <= ((v3 * (v3 / (float64(float32((((0.0 / (((v3 + v3) / 1.5) + float64(3.14159))) + (v3 / v3)) - v3))) * v3))) + ((float64(((1.5 - 1.5) / float32(-2.25))) - 3.14159) * v3))) || (-9223372036854775808 == -1)))
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v1 += (v4 * (v1 - (v4 * 1)))
        switch "foo bar" {
        case "foo bar":
            v19 := GenS1{F0: 913, F1: 42}
            _ = v19
        }
        _ = fi18
    }
    if true {
        if len(v2) > 0 {
            v2[0] = 885
        }
    } else {
        fmt.Println(v1, v15)
        v4 = v1
        var v20 int32 = 864
        _ = v20
    }
    var v21 string = "foo bar"
    v5 = (v5 + (v5 - 2))
    v4 = (v4 % (((^(v4))) | 1))
    var v22 uint64 = (v1 + uint64(v5))
    switch 42 {
    case 32768:
        v1 -= v22
    case -16:
        v1 = v4
    case -9223372036854775808:
        var v23 uint = (^(v5))
        _ = v23
    default:
        v24 := GenS2{F0: false, F1: ("\n\t" + "hello")}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
