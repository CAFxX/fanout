package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint64 }
func genFn0() uint16 {
    var v5 float64 = float64(0.1)
    v5 -= ((-2.25 + v5) - v5)
    _ = v5
    return 0
}

func main() {
    v1 := []string{}
    var v2 string = "a"
    var v3 float32 = 1e10
    v4 := []uint32{42, 878, 138, 4294967295}
    if (-1 > -78) {
        for fi6 := 0; fi6 < 6; fi6++ {
            v2 = "foo bar"
            v2 = ("hello" + v2)
            _ = fi6
        }
        if ((!false) && (((511 > 2) && true) || true)) {
            var v7 uint8 = 0
            _ = v7
        }
        v1 = append(v1, v2)
    } else {
        if (227 == 485) {
            v4 = append(v4, 2)
        }
        var v8 int32 = -1
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        if len(v1) > 0 {
            v1[0] = v2
        }
        _ = fi9
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v10 int16 = -211
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        _ = len(v1)
        v11 := GenS1{F0: 138, F1: -122, F2: 670}
        var v12 int = -232
        _ = v11
        _ = v12
    }
    v10 -= 1000
    if false {
        for fi13 := 0; fi13 < 14; fi13++ {
            if false {
                v3 -= v3
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 uint8 = 127
                    var v16 bool = (v15 <= v15)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            _ = fi13
        }
        v2 = ("hello" + v2)
    } else {
        v10 += v10
        v2 = (v2 + "hello")
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 bool = (!(true || (-524 <= -65)))
            v10 = (-(v10))
            var v19 int8 = -71
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    var v20 int64 = -592
    var v21 int = 174
    var v22 int16 = (v10 >> ((866) & 63))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, 672)
    v1 = append(v1, v2)
    var v23 uint64 = 42
    v1 = append(v1, (v2 + v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
