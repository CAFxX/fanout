package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
func genFn0(p0 uint64, p1 int16, p2 int32) int {
    v6 := []int{0}
    var v7 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint = 805
    var v2 uint32 = 870
    var v3 int32 = 1
    v4 := []int{32767, 166}
    v5 := make(map[uint32]int)
    var v8 float32 = 1e10
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 int16 = -1
            var v12 float64 = 1.5
            fmt.Println(v12)
            _ = fi10
            _ = v11
            _ = v12
        }
        v8 = (0.1 + v8)
        v8 += float32(((0.1 - ((-2.25 / (1.5 / 1.0)) - 0.0)) + 1.5))
        _ = fi9
    }
    var v13 int16 = 830
    if (((-679 != 843) || false) && true) {
        var v14 int = 682
        v1 += 4294967295
        _ = v14
    } else {
        switch 127 {
        case -553:
            v13 = (-(v13))
        default:
            var v15 int8 = 100
            _ = v15
        }
    }
    if false {
        v13 = (v13 + (v13 | (v13 * -32768)))
    }
    v4 = append(v4, 42)
    var v16 float32 = 0.1
    switch 128 {
    case -2147483648:
        fmt.Println(v1, v16, v16)
    case 123:
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := GenS1{F0: v2, F1: (("" + (("foo bar" + "") + "")) + "hello"), F2: "foo bar"}
            _ = fi17
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(v2)
            _ = fi19
        }
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        v5[v2] = (fi20 & (fi20 >> (((^(v1))) & 63)))
        v21 := GenS1{F0: v2, F1: "", F2: ""}
        v16 = v16
        _ = fi20
        _ = v21
    }
    fmt.Println(v1, v13, v8)
    var v22 uint16 = 981
    var v23 uint = v1
    fmt.Println(v23, v16, v23)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
