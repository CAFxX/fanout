package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
func genFn0(p0 uint16) string {
    switch -1 {
    case 969:
        if ((!(0 < 32767)) && true) {
            p0 = 0
        } else {
            p0 -= (p0 / (((p0 % (((p0 << ((255) & 63))) | 1))) | 1))
            p0 = p0
        }
    case 2147483647:
        fmt.Println(p0)
    default:
        v6 := GenS1{F0: (2 > 246), F1: 1}
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint = 479
        _ = fi7
        _ = v8
    }
    var v9 int32 = -2147483648
    _ = p0
    _ = v9
    return "hello"
}
func genFn1() uint {
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    switch 4294967295 {
    case 416:
        var v12 float64 = (0.0 * 3.14159)
        _ = v12
    case 127:
        v13 := GenS1{F0: (!(-2.25 < -2.25)), F1: 42}
        _ = v13
    }
    v14 := []int{127}
    var v15 string = "\n\t"
    _ = v14
    _ = v15
    return 224
}

func main() {
    var v1 int8 = -43
    v2 := []int{1}
    v3 := make(map[uint32]int)
    var v4 uint64 = 18446744073709551615
    v5 := make(map[string]int)
    var v16 int32 = 0
    _ = len(v3)
    v16 = 42
    v17 := GenS1{F0: (!(-2.25 != 0.0)), F1: -1}
    switch -1 {
    case -154:
        v18 := GenS1{F0: false, F1: 42}
        _ = v18
    case -2147483648:
        v2 = append(v2, -925)
    case -198:
        if false {
            if ((((("foo bar" == "") || false) && (676 < 0)) || ((1.5 * -2.25) >= 0.1)) && true) {
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 uint32 = 1
                    var v21 int16 = 0
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                var v22 int = 2147483647
                _ = v22
            }
        }
    default:
        v1 -= (v1 + v1)
    }
    _ = len(v2)
    if (!true) {
        for fi23 := 0; fi23 < 10; fi23++ {
            _ = len(v3)
            _ = fi23
        }
        _ = len(v2)
        _ = len(v2)
    } else {
        v4 -= (^(v4))
        fmt.Println(v4)
    }
    if (false && (!(1 != 384))) {
        var v24 int32 = 0
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS1{F0: true, F1: int64(v4)}
        _ = fi25
        _ = v26
    }
    fmt.Println(v4)
    v27 := GenS2{F0: 472, F1: "foo bar", F2: 2}
    v4 -= (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
