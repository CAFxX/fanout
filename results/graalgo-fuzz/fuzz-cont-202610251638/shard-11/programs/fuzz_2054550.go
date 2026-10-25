package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint }
type GenS2 struct { F0 bool; F1 int8; F2 bool }
func genFn0(p0 int64, p1 string) int64 {
    p1 = p1
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: "a", F1: 866, F2: 2}
        fmt.Println(fi6, p1)
        fmt.Println(p0)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: "", F1: -2147483648, F2: 2}
    v9 := []int{0, 760}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (p0 % ((42) | 1))
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{"\n\t", "a"}
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v5 uint16 = 224
    v5 -= 42
    v10 := GenS1{F0: ("foo bar" + ("\n\t" + ("a" + "hello"))), F1: 0, F2: 255}
    v5 = v5
    for fi11 := 0; fi11 < 18; fi11++ {
        v5 += 587
        switch 128 {
        case -86:
            v12 := GenS2{F0: (!(0 < 0)), F1: 0, F2: ((false || false) && false)}
            _ = v12
        case -2147483648:
            var v13 float64 = (-2.25 + 0.1)
            _ = v13
        default:
            if true {
                for fi14 := 0; fi14 < 19; fi14++ {
                    v5 = 2
                    var v15 uint8 = 200
                    var v16 uint64 = 18446744073709551615
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                v1 = v1
                if true {
                    v17 := GenS2{F0: false, F1: 100, F2: (!(!true))}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v17
                }
            }
        }
        _ = fi11
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 float64 = 0.0
        var v20 uint16 = 42
        if ((fi18 != fi18) || false) {
            var v21 float64 = v19
            _ = v21
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    if ("foo bar" <= "a") {
        v1 -= (v1 * v1)
    }
    fmt.Println(v5, v1)
    var v22 int = 112
    var v23 int32 = 42
    fmt.Println(v22)
    if (915 != 2) {
        var v24 uint8 = 128
        _ = v24
    }
    v25 := GenS2{F0: false, F1: 0, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
