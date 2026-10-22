package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 468
    var v2 string = "a"
    v3 := make(map[string]uint32)
    fmt.Println(v1, v2)
    v2 = v2
    if false {
        if false {
            if (1 > 903) {
                switch ("" + v2) {
                case "a":
                    var v4 int64 = 42
                    _ = v4
                }
                v1 += (^(v1))
                fmt.Println(v1)
            }
        }
        switch v2 {
        case "hello":
            v5 := GenS1{F0: 65535, F1: 1, F2: 29}
            _ = v5
        case "a":
            v1 -= 1
        case "foo bar":
            v2 = v2
        }
    }
    var v6 uint8 = 1
    var v7 float64 = 1.5
    v7 = v7
    v8 := GenS1{F0: 0, F1: 4294967295, F2: v6}
    var v9 int = 0
    v2 = v2
    if ((v1 * 695) > (v1 << ((489) & 63))) {
        v10 := GenS1{F0: 32767, F1: 18446744073709551615, F2: (v6 ^ v6)}
        _ = v10
    }
    v2 = (("\n\t" + "hello") + ((((v2 + v2) + (v2 + "hello")) + v2) + v2))
    if ((!(!((0.0 / (3.14159 + -2.25)) > (1e10 * -2.25)))) && true) {
        v11 := GenS1{F0: 65535, F1: 2, F2: (v6 * v6)}
        var v12 uint32 = uint32(v1)
        _ = v11
        _ = v12
    } else {
        v6 += uint8(v1)
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 int16 = 42
            _ = fi13
            _ = v14
        }
        var v15 float32 = 1.5
        _ = v15
    }
    var v16 string = v2
    for fi17 := 0; fi17 < 9; fi17++ {
        switch 32768 {
        case -2147483648:
            v7 = (float64((3.14159 - float32(v7))) * ((0.1 - 1e10) - v7))
        case 42:
            var v18 uint = (^(v1))
            _ = v18
        case 127:
            v19 := GenS2{F0: 255}
            _ = v19
        default:
            v20 := GenS2{F0: v1}
            _ = v20
        }
        v2 = v2
        _ = fi17
    }
    fmt.Println(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
