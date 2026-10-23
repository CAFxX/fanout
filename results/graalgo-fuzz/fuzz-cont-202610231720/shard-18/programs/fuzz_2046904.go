package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
func genFn0() uint16 {
    var v7 int64 = 1
    v7 -= v7
    var v8 bool = false
    var v9 int32 = 676
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 int16 = 985
    var v2 uint64 = 87
    v3 := []int{127, 565}
    var v4 int8 = 127
    var v5 uint64 = v2
    var v6 float64 = 3.14159
    var v10 int16 = v1
    for fi11 := 0; fi11 < 13; fi11++ {
        v4 = (v4 * 36)
        _ = fi11
    }
    v2 = (v5 + (^(v2)))
    if (624 > 0) {
        v12 := GenS1{F0: (!((^(v5)) >= (v2 - (v5 - 1)))), F1: 0}
        _ = v12
    } else {
        if true {
            var v13 float32 = (-2.25 + 0.1)
            _ = v13
        } else {
            v1 = v10
            var v14 int32 = 0
            var v15 float64 = float64((0.0 * ((((-2.25 - 1e10) - 3.14159) - 0.0) + 1.5)))
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            v6 = -2.25
            _ = fi16
        }
        var v17 int16 = (v10 & v1)
        _ = v17
    }
    var v18 uint = 519
    for fi19 := 0; fi19 < 16; fi19++ {
        if false {
            fmt.Println(v5, v6)
            v20 := GenS1{F0: false, F1: 127}
            v3 = append(v3, fi19)
            _ = v20
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi19
    }
    var v21 uint16 = 255
    _ = len(v3)
    _ = len(v3)
    switch -2147483648 {
    case 1:
        v22 := GenS2{F0: 9223372036854775807, F1: (v4 | (-(v4))), F2: (!(!false))}
        _ = v22
    case 171:
        var v23 string = "a"
        _ = v23
    case 2147483647:
        v1 += v10
    default:
        for fi24 := 0; fi24 < 3; fi24++ {
            if (v6 == (v6 + v6)) {
                v4 -= (v4 ^ v4)
            }
            v25 := GenS1{F0: true, F1: v4}
            _ = fi24
            _ = v25
        }
    }
    fmt.Println(v21)
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 float64 = 0.1
        v3 = append(v3, (fi26 * fi26))
        v27 = (0.1 + float64(3.14159))
        _ = fi26
        _ = v27
    }
    var v28 int64 = 142
    var v29 int64 = (v28 - 937)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
