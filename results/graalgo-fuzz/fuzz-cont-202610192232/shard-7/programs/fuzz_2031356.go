package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int64, p1 int) uint {
    v7 := []int{2147483647}
    p0 -= (-(p0))
    p0 -= p0
    p1 -= p1
    for fi8 := 0; fi8 < 6; fi8++ {
        p0 = p0
        var v9 float64 = float64((-2.25 + (1.5 / ((0.1 - 1.5) - (0.1 / 1.0)))))
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := []string{"\n\t"}
    v2 := []uint32{}
    var v3 bool = false
    var v4 string = "\n\t"
    v5 := []int{127, 4294967295, -1, 4294967295}
    v6 := make(map[uint32]string)
    var v10 int32 = -632
    for fi11 := 0; fi11 < 3; fi11++ {
        v3 = (282 != -32768)
        var v12 int8 = -128
        var v13 uint64 = 364
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v3 = false
        v15 := GenS1{F0: 2, F1: (0 <= -795)}
        v16 := GenS1{F0: 909, F1: true}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := v6[4294967295]
    v18 := GenS2{F0: 2, F1: 155}
    if (9223372036854775807 == 1) {
        delete(v6, 2147483648)
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        if v3 {
            v10 = (v10 ^ v10)
            if true {
                fmt.Println(v10, fi19, v17)
                v6[812] = (v4 + v4)
            } else {
                for fi20 := 0; fi20 < 19; fi20++ {
                    if v3 {
                        if (!(v10 >= v10)) {
                            var v21 float64 = 3.14159
                            var v22 string = (v17 + v17)
                            _ = v21
                            _ = v22
                        }
                        var v23 int8 = -84
                        _ = v23
                    }
                    var v24 bool = false
                    _ = len(v5)
                    _ = fi20
                    _ = v24
                }
                switch v4 {
                case "a":
                    v6[846] = v17
                case "":
                    v10 = (v10 & 0)
                }
            }
        }
        _ = fi19
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 bool = (158 > 37)
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 int64 = -1
            v4 = ((((("foo bar" + v17) + "\n\t") + "") + v17) + v17)
            var v29 float64 = 0.0
            _ = fi27
            _ = v28
            _ = v29
        }
        var v30 int = fi25
        _ = fi25
        _ = v26
        _ = v30
    }
    if true {
        switch -9223372036854775808 {
        case -1:
            var v31 uint64 = 18446744073709551615
            _ = v31
        case 127:
            for fi32 := 0; fi32 < 16; fi32++ {
                var v33 bool = (!(-1 < 0))
                _ = fi32
                _ = v33
            }
        default:
            v17 = v4
        }
        var v34 uint32 = 2147483647
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
