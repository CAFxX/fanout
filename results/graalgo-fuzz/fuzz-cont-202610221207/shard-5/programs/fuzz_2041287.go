package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int32; F1 int64 }
func genFn0(p0 int32, p1 int16, p2 int8) uint64 {
    var v6 float32 = (0.0 / (3.14159 - ((-2.25 / 0.1) - 0.1)))
    var v7 int = 128
    v8 := GenS1{F0: "\n\t", F1: 1, F2: ("hello" + ("\n\t" + ("" + ("foo bar" + "\n\t"))))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint = 795
    v2 := make(map[int]string)
    v3 := []uint32{}
    var v4 bool = false
    var v5 int32 = 781
    fmt.Println(v1, v4, v4)
    fmt.Println(v1)
    fmt.Println(v5, v1, v4)
    if (148 > 42) {
        v1 -= v1
    } else {
        var v9 float64 = 3.14159
        _ = v9
    }
    v10 := GenS2{F0: -2147483648}
    v5 = int32(v1)
    var v11 int64 = 414
    if (!v4) {
        if (!v4) {
            v4 = ((3.14159 * 1.5) != 0.0)
        } else {
            v12 := GenS1{F0: ("hello" + ("hello" + "a")), F1: 722, F2: (((("hello" + ("\n\t" + "foo bar")) + "\n\t") + "foo bar") + "")}
            v13 := GenS1{F0: ("foo bar" + "\n\t"), F1: 32767, F2: "\n\t"}
            _ = v12
            _ = v13
        }
        var v14 uint64 = 531
        _ = v14
    } else {
        var v15 uint16 = 32767
        v16 := GenS3{F0: 2, F1: 255}
        fmt.Println(v11, v11, v5)
        _ = v15
        _ = v16
    }
    v5 -= v5
    v17 := v2[-435]
    var v18 int32 = (v5 | 191)
    var v19 int8 = 113
    switch (v17 + "foo bar") {
    case "a":
        v18 -= v5
    case "foo bar":
        v3 = append(v3, 65535)
    default:
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 uint16 = 0
            var v22 bool = (v5 < (v5 >> ((v1) & 63)))
            v23 := GenS4{F0: (v5 ^ -201), F1: (v11 % (((v11 & (v11 ^ v11))) | 1))}
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5, v11, v4)
    v1 += (^(v1))
    var v24 uint64 = 107
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
