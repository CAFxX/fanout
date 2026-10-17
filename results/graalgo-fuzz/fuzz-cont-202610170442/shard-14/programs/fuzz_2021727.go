package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int16; F2 int16 }
func genFn0(p0 string) int64 {
    switch "" {
    case "hello":
        v5 := []int{4294967295, 206}
        _ = v5
    case "foo bar":
        if (false && false) {
            if (true && false) {
                var v6 uint8 = 127
                var v7 float64 = 0.1
                _ = v6
                _ = v7
            }
        }
    case "\n\t":
        var v8 float64 = 0.1
        _ = v8
    }
    v9 := []int{128, 42}
    v10 := GenS1{F0: 42, F1: 176, F2: 0}
    p0 = p0
    var v11 uint64 = 18446744073709551615
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    v1 := []string{}
    var v2 uint64 = 1
    v3 := make(map[int]int)
    var v4 int16 = -128
    v12 := GenS1{F0: 833, F1: (^(v2)), F2: v2}
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: -9223372036854775808, F1: v2, F2: (v2 << ((255) & 63))}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v16 := GenS1{F0: 0, F1: (v2 * (v2 | 2)), F2: 871}
    v2 = v2
    v4 = v4
    var v17 int16 = v4
    var v18 float64 = ((((float64((float32(((3.14159 * 0.0) + -2.25)) - ((((3.14159 * 1e10) * 0.1) + (3.14159 + ((0.0 + 1.5) * 1.5))) * -2.25))) + (0.0 - ((3.14159 / 1.0) * 1.5))) * 1.5) + (1e10 + 0.1)) - 1e10)
    fmt.Println(v2, v2, v2)
    v19 := GenS2{F0: v2, F1: v4, F2: 859}
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 float32 = 1.5
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        v18 -= v18
        v23 := v3[(fi22 * 32767)]
        _ = fi22
        _ = v23
    }
    switch ("hello" + "foo bar") {
    case "foo bar":
        if false {
            v1 = append(v1, "foo bar")
            var v24 int16 = (^(v17))
            _ = v24
        }
    case "\n\t":
        if (!(2147483648 < 914)) {
            if (!(true && (2 < 42))) {
                if true {
                    v25 := GenS1{F0: genFn0(""), F1: (^(v2)), F2: (v2 / ((1) | 1))}
                    var v26 uint16 = 0
                    for fi27 := 0; fi27 < 2; fi27++ {
                        var v28 float64 = 1e10
                        _ = fi27
                        _ = v28
                    }
                    _ = v25
                    _ = v26
                }
                v4 = (v17 / (((v4 & (v17 ^ v17))) | 1))
            } else {
                for fi29 := 0; fi29 < 10; fi29++ {
                    switch fi29 {
                    case 84:
                        var v30 bool = (((float64(0.0) - (((v18 + (3.14159 * 0.1)) / v18) * float64(float32(v18)))) - v18) > (v18 - ((3.14159 + float64((0.1 / ((((0.1 + (1e10 / -2.25)) - (1.5 * (0.1 + 1.5))) - 1.5) / -2.25)))) - float64(1e10))))
                        _ = v30
                    }
                    switch ("" + "") {
                    case "":
                        v4 = (v17 - -20)
                    default:
                        fmt.Println(v2)
                    }
                    _ = fi29
                }
            }
        }
    case "a":
        if true {
            fmt.Println(v2, v17, v2)
            v17 = (v4 | v17)
            v31 := v3[128]
            _ = v31
        } else {
            var v32 int32 = 42
            v33 := GenS2{F0: (v2 | (v2 >> ((299) & 63))), F1: 1000, F2: (^(v17))}
            v34 := GenS2{F0: v2, F1: (v4 % ((593) | 1)), F2: v4}
            _ = v32
            _ = v33
            _ = v34
        }
    }
    _ = len(v1)
    for fi35 := 0; fi35 < 4; fi35++ {
        v36 := GenS2{F0: v2, F1: v4, F2: v4}
        _ = fi35
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
