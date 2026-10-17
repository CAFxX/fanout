package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }

func main() {
    var v1 float64 = (((1e10 + (((((0.1 / (0.1 - 1.5)) - (0.1 * float64(-2.25))) + (1e10 - 1e10)) + 3.14159) / ((0.0 / 1.5) - 3.14159))) + 0.0) - (3.14159 + (1.5 * (0.0 - 1e10))))
    v2 := []uint32{806, 42}
    v3 := make(map[string]uint32)
    var v4 int32 = 0
    var v5 uint8 = 127
    if (!((v1 / v1) >= (((1e10 - (v1 * v1)) * v1) * (v1 * v1)))) {
        for fi6 := 0; fi6 < 5; fi6++ {
            v4 = (-(v4))
            _ = fi6
        }
        if false {
            var v7 int32 = (v4 % ((v4) | 1))
            v8 := GenS1{F0: 209, F1: true, F2: 0}
            if false {
                _ = len(v3)
            } else {
                var v9 int32 = v7
                _ = v9
            }
            _ = v7
            _ = v8
        } else {
            v4 += (v4 << ((4294967295) & 63))
            v5 = (v5 ^ 84)
        }
        var v10 float64 = 0.1
        _ = v10
    }
    v5 += 1
    v2 = append(v2, 2147483647)
    switch ("foo bar" + "foo bar") {
    case "":
        var v11 bool = (-609 < 42)
        _ = v11
    case "\n\t":
        var v12 float64 = v1
        _ = v12
    case "a":
        var v13 uint8 = v5
        _ = v13
    default:
        _ = len(v2)
    }
    var v14 float64 = v1
    var v15 float32 = float32((float64(1.5) - v1))
    var v16 float64 = v1
    var v17 uint16 = 624
    v14 += v14
    fmt.Println(v1, v15)
    switch "" {
    case "\n\t":
        delete(v3, "hello")
    }
    var v18 uint32 = 106
    for fi19 := 0; fi19 < 14; fi19++ {
        if true {
            for fi20 := 0; fi20 < 7; fi20++ {
                v18 += v18
                v21 := GenS2{F0: 65535, F1: 18446744073709551615, F2: false}
                _ = fi20
                _ = v21
            }
        }
        _ = fi19
    }
    var v22 int16 = -1
    if (((v5 ^ v5) < 128) && true) {
        v18 = (v18 | v18)
    } else {
        var v23 string = "\n\t"
        _ = v23
    }
    fmt.Println(v14, v22)
    for fi24 := 0; fi24 < 7; fi24++ {
        if (false || (527 < (v22 + 824))) {
            var v25 int32 = -1
            v15 += v15
            v5 += v5
            _ = v25
        } else {
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 string = ("\n\t" + ((("a" + ("hello" + ("hello" + ((("foo bar" + "a") + ("hello" + ("\n\t" + ""))) + "hello")))) + ("hello" + ("foo bar" + "foo bar"))) + "foo bar"))
                _ = fi26
                _ = v27
            }
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
