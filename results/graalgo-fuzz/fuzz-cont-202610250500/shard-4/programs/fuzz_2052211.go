package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32; F1 int; F2 string }
func genFn0(p0 uint32, p1 int32, p2 uint16) int64 {
    if true {
        v4 := GenS1{F0: uint16(p0), F1: 42}
        for fi5 := 0; fi5 < 1; fi5++ {
            v6 := []int{}
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 uint32 = 949
                var v9 uint64 = 18446744073709551615
                v8 = (p0 << ((4294967295) & 63))
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 uint16 = (p2 % ((32767) | 1))
            _ = fi5
            _ = v10
            _ = v6
        }
        _ = v4
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 string = ""
        var v13 string = "foo bar"
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return -681
}

func main() {
    var v1 int32 = 1
    v2 := make(map[string]int)
    var v3 string = "hello"
    if (0.0 >= ((1e10 / 3.14159) + (0.0 / 1.0))) {
        fmt.Println(v3)
        v1 = (v1 / ((0) | 1))
    }
    v14 := []int{}
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 bool = false
        _ = fi15
        _ = v16
    }
    switch ((((v3 + (((v3 + v3) + v3) + ("\n\t" + v3))) + v3) + "") + v3) {
    case "a":
        var v17 int16 = 587
        _ = v17
    case "":
        switch v3 {
        case "foo bar":
            _ = len(v2)
        case "":
            v14 = append(v14, 0)
        default:
            if (891 > 0) {
                var v18 uint = 4294967295
                v19 := GenS1{F0: 694, F1: 9223372036854775807}
                v18 = (v18 ^ v18)
                _ = v18
                _ = v19
            }
        }
    case "foo bar":
        var v20 uint8 = 128
        _ = v20
    }
    var v21 int = 128
    if true {
        var v22 float64 = (((float64(float32((-2.25 * (1e10 + (((1e10 / ((-2.25 - (0.0 / 1.5)) / 3.14159)) - (-2.25 + 0.1)) / 3.14159))))) - 1e10) + 3.14159) * 0.0)
        _ = len(v14)
        v1 -= -1
        _ = v22
    } else {
        v23 := v2[(v3 + ("" + ""))]
        v2[v3] = (v23 % (((v23 * (v21 / ((731) | 1)))) | 1))
        v24 := GenS2{F0: -889, F1: v21, F2: v3}
        _ = v23
        _ = v24
    }
    switch 593 {
    case 127:
        var v25 int = v21
        _ = v25
    case 4294967295:
        v3 = "foo bar"
    case 209:
        v26 := v2[v3]
        _ = v26
    default:
        v14 = append(v14, v21)
    }
    _ = len(v14)
    fmt.Println(v1, v1, v3)
    v21 -= 124
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
}
