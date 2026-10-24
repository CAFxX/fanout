package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 int32) int {
    p0 = p0
    if false {
        v6 := []int{}
        _ = v6
    } else {
        p1 = p1
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        p1 = (p1 | (p1 ^ (p1 ^ (p1 / (((p1 % ((617) | 1))) | 1)))))
        var v8 float64 = ((float64((0.1 + 1e10)) + 1.5) - ((0.0 / 3.14159) - 3.14159))
        _ = fi7
        _ = v8
    }
    p0 -= (-(p0))
    v9 := []int{(^(p0)), (p0 | (^(p0)))}
    _ = p0
    _ = p1
    _ = v9
    return p0
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]string)
    var v3 float32 = (3.14159 - 1.5)
    var v4 string = "hello"
    var v5 float32 = v3
    if false {
        var v10 int8 = 127
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := v2[455]
            _ = len(v2)
            v10 -= v10
            _ = fi11
            _ = v12
        }
        if (0 > 2) {
            v13 := []int{9223372036854775807, 0, 591}
            v14 := GenS1{F0: 977, F1: 65535}
            if ((3.14159 * 0.1) == 1e10) {
                fmt.Println(v3)
                v15 := GenS1{F0: 1, F1: 0}
                _ = v15
            }
            _ = v13
            _ = v14
        }
        _ = v10
    }
    switch v4 {
    case "foo bar":
        var v16 int = 32767
        _ = v16
    case "\n\t":
        var v17 uint16 = 429
        _ = v17
    default:
        v18 := []int{-107}
        _ = v18
    }
    v1 = v1
    var v19 uint = 1
    _ = len(v2)
    var v20 int = -18
    for fi21 := 0; fi21 < 1; fi21++ {
        if false {
            for fi22 := 0; fi22 < 7; fi22++ {
                if false {
                    v19 = v19
                    v19 = 18446744073709551615
                } else {
                    v3 -= (v5 - v5)
                }
                _ = fi22
            }
            var v23 int = (fi21 << (((^(v19))) & 63))
            _ = v23
        }
        v1 -= v1
        _ = fi21
    }
    v20 += (v20 + (v20 - (v20 ^ 32768)))
    var v24 int64 = 42
    v19 = v19
    for fi25 := 0; fi25 < 6; fi25++ {
        fmt.Println(v24)
        for fi26 := 0; fi26 < 9; fi26++ {
            if false {
                var v27 uint16 = 65535
                var v28 int8 = -51
                var v29 int16 = 1000
                _ = v27
                _ = v28
                _ = v29
            } else {
                v5 += (0.0 / ((0.0 - v3) - 1e10))
            }
            _ = fi26
        }
        for fi30 := 0; fi30 < 3; fi30++ {
            var v31 bool = (0 <= (v20 * fi30))
            switch (fi25 ^ (fi30 + (^(fi30)))) {
            case 933:
                v31 = (!(!v31))
            case -1:
                v32 := []int{(fi30 + 32767), 32767, v20}
                _ = v32
            case 2147483647:
                v33 := GenS1{F0: 1000, F1: uint16(v19)}
                _ = v33
            }
            _ = fi30
            _ = v31
        }
        _ = fi25
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        v3 = float32(0.0)
        if (true && false) {
            for fi35 := 0; fi35 < 6; fi35++ {
                if true {
                    v36 := []int{0, (v20 - (v20 & fi34))}
                    _ = v36
                }
                var v37 float32 = float32((1e10 - 3.14159))
                _ = fi35
                _ = v37
            }
            v38 := v2[594]
            _ = v38
        } else {
            v4 = v4
            v39 := GenS1{F0: 1, F1: 322}
            _ = v39
        }
        _ = fi34
    }
    v40 := GenS1{F0: -32768, F1: 628}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
