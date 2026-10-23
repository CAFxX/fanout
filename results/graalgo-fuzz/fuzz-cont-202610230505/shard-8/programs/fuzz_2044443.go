package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
func genFn0(p0 uint, p1 uint8) uint16 {
    var v5 int64 = 9223372036854775807
    var v6 int16 = 1
    v7 := []int{684}
    v5 -= v5
    fmt.Println(p0, v6)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 516
}
func genFn1(p0 uint64) int64 {
    v8 := []int{2147483647}
    if (669 < 0) {
        var v9 int32 = 42
        if (false || (2 == 103)) {
            v10 := GenS1{F0: 42, F1: false}
            _ = v10
        } else {
            if true {
                for fi11 := 0; fi11 < 12; fi11++ {
                    if (!true) {
                        p0 -= p0
                        v9 = -11
                    } else {
                        v12 := GenS1{F0: 0, F1: true}
                        _ = v12
                    }
                    _ = fi11
                }
                v9 = 2147483647
                for fi13 := 0; fi13 < 20; fi13++ {
                    v14 := GenS1{F0: -332, F1: true}
                    fmt.Println(p0, p0, fi13)
                    _ = fi13
                    _ = v14
                }
            } else {
                if len(v8) > 0 {
                    v8[0] = 4294967295
                }
                var v15 uint64 = 552
                _ = v15
            }
        }
        if (2 != 0) {
            switch "" {
            case "hello":
                p0 = 4294967295
            }
            var v16 int64 = -891
            _ = v16
        }
        _ = v9
    }
    p0 = 1
    if true {
        fmt.Println(p0)
    } else {
        fmt.Println(p0)
        p0 = (p0 * p0)
        switch ("foo bar" + "") {
        case "hello":
            var v17 float32 = ((0.0 + 0.0) + 0.0)
            _ = v17
        }
    }
    p0 += (p0 - (p0 >> ((255) & 63)))
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := []int{}
    var v2 int32 = -138
    v3 := []uint32{2, 2, 121}
    var v4 int = -2147483648
    var v18 int64 = 9223372036854775807
    v19 := GenS1{F0: 32767, F1: (((-292 < 1000) && false) && true)}
    var v20 float64 = 0.0
    v21 := GenS1{F0: 42, F1: (!((("foo bar" + "\n\t") + ("hello" + "hello")) >= ("\n\t" + "a")))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += v2
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 int32 = 377
        _ = fi22
        _ = v23
    }
    v20 = float64(1.5)
    v24 := GenS1{F0: 1, F1: true}
    fmt.Println(v2, v4, v2)
    _ = len(v3)
    var v25 int = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
