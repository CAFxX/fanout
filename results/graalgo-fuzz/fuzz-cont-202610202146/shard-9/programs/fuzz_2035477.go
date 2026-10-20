package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int32 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint8 }

func main() {
    v1 := []uint32{610}
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    var v4 float64 = (1.5 / (0.0 + 1e10))
    v4 = ((float64((0.1 * 1e10)) / 1.0) - (-2.25 + (((v4 / 3.14159) / (v4 / float64(1e10))) + ((0.0 + float64(1.5)) * (((3.14159 / 0.1) * (((float64(0.1) * 1e10) / (v4 - 0.0)) / v4)) * v4)))))
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 float64 = (v4 / 1.5)
        _ = fi5
        _ = v6
    }
    if (false || (-32768 < -46)) {
        v3[2147483647] = true
        v7 := GenS1{F0: ("a" + "a"), F1: 42, F2: -235}
        fmt.Println(v4)
        _ = v7
    } else {
        switch "foo bar" {
        case "hello":
            v4 += 1e10
        default:
            fmt.Println(v4)
        }
    }
    var v8 int64 = 894
    v4 -= ((1.5 - ((0.0 - (3.14159 - 1.5)) / v4)) + v4)
    if (-18 == 127) {
        if false {
            var v9 string = "a"
            _ = v9
        }
        v10 := GenS2{F0: 238, F1: ("" < ("foo bar" + "\n\t")), F2: 2}
        v11 := GenS2{F0: 2147483648, F1: (false && true), F2: 128}
        _ = v10
        _ = v11
    }
    var v12 uint16 = 255
    var v13 uint8 = 167
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 uint32 = 4294967295
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint32 = 4294967295
        _ = fi16
        _ = v17
    }
    var v18 uint = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
