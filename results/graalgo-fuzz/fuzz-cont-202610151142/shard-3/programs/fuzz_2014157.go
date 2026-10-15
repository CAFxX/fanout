package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0(p0 uint8, p1 uint) uint {
    v4 := GenS1{F0: -1, F1: 0}
    if (!(!(((1 > 2) || (((("" + (("foo bar" + "hello") + "")) + (("hello" + (("" + ("\n\t" + "a")) + "a")) + "foo bar")) + "foo bar") >= ((("" + "") + "") + "foo bar"))) && (2147483647 <= 42)))) {
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := []int{}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: 42, F1: 42}
            _ = fi7
            _ = v8
        }
        fmt.Println(p1, p0)
    } else {
        switch "hello" {
        case "":
            var v9 bool = false
            _ = v9
        default:
            v10 := []int{-680, 630}
            _ = v10
        }
        var v11 int8 = -128
        var v12 int32 = 0
        _ = v11
        _ = v12
    }
    var v13 uint32 = 65535
    var v14 float32 = 0.1
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v4
    return (p1 & (p1 ^ p1))
}

func main() {
    var v1 bool = false
    var v2 uint16 = 323
    var v3 int = -9223372036854775808
    v1 = ((!(1 < 0)) || (((1e10 / (0.1 / 1e10)) * 0.1) <= 1e10))
    v2 += v2
    switch "a" {
    case "":
        v15 := []int{v3, (v3 & v3)}
        _ = v15
    case "foo bar":
        var v16 uint8 = 42
        _ = v16
    case "hello":
        fmt.Println(v2, v3, v1)
    default:
        var v17 uint8 = 1
        _ = v17
    }
    v18 := []int{(v3 / ((v3) | 1)), v3}
    for fi19 := 0; fi19 < 13; fi19++ {
        v3 += 918
        switch v3 {
        case -2147483648:
            v20 := GenS2{F0: -766, F1: (!(-8 <= 42)), F2: (v2 + v2)}
            _ = v20
        case 1:
            _ = len(v18)
        default:
            var v21 string = "\n\t"
            _ = v21
        }
        v22 := GenS1{F0: 133, F1: (v2 - 581)}
        _ = fi19
        _ = v22
    }
    _ = len(v18)
    v2 -= 345
    v2 -= (v2 | (v2 ^ (v2 - (v2 | v2))))
    var v23 uint64 = 18446744073709551615
    fmt.Println(v23)
    switch 42 {
    case 2147483647:
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 uint16 = (v2 * v2)
            var v26 float64 = (1.5 - 0.0)
            _ = fi24
            _ = v25
            _ = v26
        }
    default:
        if (862 != v2) {
            _ = len(v18)
        } else {
            var v27 float32 = (1e10 + 0.0)
            _ = v27
        }
    }
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    if v1 {
        v28 := GenS2{F0: -857, F1: (!((v23 % ((42) | 1)) != (v23 | 72))), F2: v2}
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := GenS2{F0: -571, F1: ((v1 && v1) && (2 >= 160)), F2: 554}
            v23 = v23
            var v31 uint64 = 360
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    } else {
        if len(v18) > 0 {
            v18[0] = (v3 << ((0) & 63))
        }
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
