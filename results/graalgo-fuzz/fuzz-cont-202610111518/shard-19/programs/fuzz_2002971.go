package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 uint16; F2 uint32 }
func genFn0() uint64 {
    v7 := []int{}
    var v8 uint = 2
    switch 241 {
    case 547:
        v8 += (^(v8))
    }
    fmt.Println(v8)
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int32 = 0
        _ = fi9
        _ = v10
    }
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    var v3 int8 = -18
    v4 := make(map[string]string)
    v5 := []uint32{42, 1}
    var v6 uint = 255
    fmt.Println(v3, v3)
    v6 = v6
    delete(v4, "hello")
    var v11 bool = false
    v6 -= (v6 - 2)
    v11 = ((32767 <= 875) && (v11 || ((167 <= v6) && true)))
    v11 = true
    for fi12 := 0; fi12 < 3; fi12++ {
        switch ("" + "\n\t") {
        case "a":
            if (((0.0 * 1e10) / (float32(0.0) + 3.14159)) == 1e10) {
                _ = len(v1)
            } else {
                fmt.Println(v11, v11, v3)
            }
        case "hello":
            v13 := GenS1{F0: int64(v6), F1: 128, F2: ((fi12 ^ 286) > (-(fi12)))}
            _ = v13
        default:
            var v14 int8 = 0
            _ = v14
        }
        var v15 int64 = -9223372036854775808
        if (97 <= 128) {
            v16 := GenS1{F0: v15, F1: 33, F2: v11}
            v3 = v3
            _ = v16
        } else {
            for fi17 := 0; fi17 < 7; fi17++ {
                fmt.Println(v15, v11, v6)
                v6 += (v6 ^ (v6 * v6))
                v6 = (v6 + (v6 / ((v6) | 1)))
                _ = fi17
            }
            v18 := GenS1{F0: -173, F1: 42, F2: (float64((1e10 * -2.25)) != ((-2.25 + 3.14159) / (float64(float32(0.1)) + (float64(-2.25) - 0.1))))}
            _ = v18
        }
        _ = fi12
        _ = v15
    }
    v3 = v3
    _ = len(v5)
    var v19 int16 = -1
    for fi20 := 0; fi20 < 2; fi20++ {
        if false {
            if v11 {
                var v21 int = (fi20 - fi20)
                _ = v21
            } else {
                v11 = v11
            }
            v3 += v3
        } else {
            v3 = 100
            var v22 int8 = (-(v3))
            _ = v22
        }
        _ = fi20
    }
    v5 = append(v5, 0)
    var v23 uint64 = 709
    var v24 uint = (^(v6))
    v25 := v2[-9223372036854775808]
    if v11 {
        v26 := GenS1{F0: 0, F1: 10, F2: (v25 != (^(v25)))}
        switch -2147483648 {
        case 1:
            var v27 int64 = -1
            _ = v27
        case 778:
            for fi28 := 0; fi28 < 2; fi28++ {
                switch "\n\t" {
                case "hello":
                    var v29 int16 = v19
                    _ = v29
                case "a":
                    v23 = (v23 + v23)
                }
                switch (fi28 - fi28) {
                case 1:
                    if v11 {
                        v30 := GenS2{F0: 1, F1: 709, F2: v25}
                        switch fi28 {
                        case 128:
                            for fi31 := 0; fi31 < 3; fi31++ {
                                v32 := GenS2{F0: 2147483647, F1: 32767, F2: v25}
                                _ = fi31
                                _ = v32
                            }
                        case 275:
                            if false {
                                var v33 float32 = 1e10
                                v11 = v11
                                _ = v33
                            }
                        }
                        _ = v30
                    }
                }
                delete(v4, "")
                _ = fi28
            }
        default:
            v3 = int8(v24)
        }
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
