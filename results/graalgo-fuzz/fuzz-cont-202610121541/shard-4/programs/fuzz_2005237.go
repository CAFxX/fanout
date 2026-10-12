package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 255
    var v3 int32 = -1
    var v4 float64 = -2.25
    v5 := make(map[string]int)
    v3 += 2147483647
    v6 := GenS1{F0: 0, F1: 2}
    var v7 uint16 = (v2 + 65535)
    if (false && (float64(-2.25) != (((((((v4 * -2.25) * v4) - (v4 - -2.25)) - float64(((-2.25 + (0.1 / (0.1 - 1.5))) - 0.0))) - v4) * v4) - v4))) {
        var v8 int32 = 42
        fmt.Println(v3)
        _ = v8
    }
    _ = len(v1)
    for fi9 := 0; fi9 < 19; fi9++ {
        v2 += v7
        v3 = (-(v3))
        fmt.Println(v2)
        _ = fi9
    }
    v10 := []int{0, 1, 806}
    if (!((0 != 863) || true)) {
        var v11 uint32 = 113
        v2 = v7
        fmt.Println(v3, v4)
        _ = v11
    } else {
        switch 9223372036854775807 {
        case 2147483647:
            var v12 int64 = -432
            _ = v12
        }
        v2 += (v7 >> ((4294967295) & 63))
        v7 -= (v2 << ((42) & 63))
    }
    var v13 float64 = -2.25
    var v14 uint32 = 2147483648
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS1{F0: v14, F1: 127}
        _ = fi15
        _ = v16
    }
    if false {
        var v17 int64 = 121
        switch -2147483648 {
        case -2147483648:
            if ((v14 / ((v14) | 1)) < v14) {
                v10 = append(v10, 128)
                v18 := GenS1{F0: v14, F1: 127}
                for fi19 := 0; fi19 < 1; fi19++ {
                    v20 := GenS1{F0: (v14 - 418), F1: 184}
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                v5[((("hello" + "") + "foo bar") + ("foo bar" + "foo bar"))] = 32767
                if false {
                    if false {
                        v21 := GenS2{F0: 2, F1: 255, F2: v14}
                        _ = v21
                    } else {
                        v3 = (v3 - (v3 ^ (v3 ^ v3)))
                    }
                    switch "\n\t" {
                    case "\n\t":
                        v14 -= 2
                    }
                } else {
                    v22 := v5["a"]
                    for fi23 := 0; fi23 < 19; fi23++ {
                        v4 += v4
                        var v24 string = ""
                        var v25 uint16 = 42
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    v13 -= v13
                    _ = v22
                }
                switch -1 {
                case -1:
                    _ = len(v5)
                case 4294967295:
                    _ = len(v10)
                }
            }
        }
        _ = v17
    }
    delete(v5, (("foo bar" + (("foo bar" + ("hello" + ("" + "\n\t"))) + "foo bar")) + ("" + "")))
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 uint8 = 128
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        var v28 int8 = 1
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 float64 = (v13 - (1.5 + 0.1))
            _ = len(v1)
            if false {
                v13 += (v4 - (v31 / v4))
                v13 += 1.5
            }
            _ = fi30
            _ = v31
        }
        v32 := GenS2{F0: 2, F1: 0, F2: v14}
        _ = fi29
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
