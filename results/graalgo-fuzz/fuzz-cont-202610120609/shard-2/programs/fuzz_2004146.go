package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int64, p1 int16, p2 string) uint64 {
    var v5 uint64 = 66
    if (((float32((0.1 - 1.5)) == float32((1e10 + 0.1))) && ((0.0 + 3.14159) > -2.25)) && true) {
        if (56 < 367) {
            fmt.Println(p1, p0, p2)
        }
        v6 := []int{2147483647, -32, 220}
        _ = v6
    }
    var v7 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return v5
}

func main() {
    v1 := []int{763, 1}
    var v2 uint32 = 4294967295
    var v3 string = ""
    var v4 uint16 = 983
    for fi8 := 0; fi8 < 18; fi8++ {
        v2 -= (v2 / (((v2 * (v2 << ((18446744073709551615) & 63)))) | 1))
        v3 = ("a" + ((v3 + "a") + v3))
        v4 -= (v4 ^ v4)
        _ = fi8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi9
    }
    var v10 float64 = 3.14159
    v11 := GenS1{F0: (((("" + v3) + (v3 + ((("a" + "foo bar") + (v3 + "foo bar")) + ""))) + (v3 + v3)) + v3)}
    for fi12 := 0; fi12 < 12; fi12++ {
        v10 += v10
        var v13 float32 = ((0.0 - 0.0) * 3.14159)
        var v14 uint = 2
        _ = fi12
        _ = v13
        _ = v14
    }
    v2 -= (^(v2))
    fmt.Println(v3, v10, v4)
    fmt.Println(v2, v2, v3)
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            if true {
                var v17 uint8 = 157
                _ = v17
            } else {
                v10 -= ((((v10 + 1.5) + (v10 - v10)) / 1.0) * v10)
                var v18 int16 = -496
                _ = len(v1)
                _ = v18
            }
            var v19 string = v3
            fmt.Println(v3, v3, v4)
            _ = fi16
            _ = v19
        }
        var v20 uint32 = (^(v2))
        _ = fi15
        _ = v20
    }
    v2 -= v2
    var v21 int32 = 924
    switch 42 {
    case 331:
        for fi22 := 0; fi22 < 13; fi22++ {
            if false {
                v21 = (v21 | v21)
                var v23 uint = 80
                var v24 int8 = -1
                _ = v23
                _ = v24
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi22
        }
    case -718:
        switch "foo bar" {
        case "":
            for fi25 := 0; fi25 < 7; fi25++ {
                v10 = v10
                _ = fi25
            }
        }
    case -668:
        var v26 uint64 = 0
        _ = v26
    default:
        v3 = ("\n\t" + ((v3 + (v3 + (v3 + v3))) + v3))
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        v27 := GenS1{F0: v3}
        v4 -= v4
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
