package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int32; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint; F1 int8; F2 uint64 }

func main() {
    var v1 int64 = 0
    var v2 int = 128
    var v3 float32 = float32((0.1 + 0.0))
    var v4 string = "a"
    var v5 uint8 = 128
    var v6 uint32 = 309
    if false {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int8 = 1
            if ((1 == 42) && (uint32(v5) >= 42)) {
                v5 = (v5 ^ (v5 & v5))
                var v9 uint = 18446744073709551615
                if true {
                    switch fi7 {
                    case 32768:
                        v10 := GenS1{F0: 957, F1: v4}
                        _ = v10
                    case 128:
                        for fi11 := 0; fi11 < 12; fi11++ {
                            v12 := []int{fi11, 170}
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            var v13 bool = true
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    default:
                        var v14 float64 = float64((v3 / v3))
                        _ = v14
                    }
                    if true {
                        if false {
                            v9 -= v9
                        }
                        fmt.Println(v9, v6, v4)
                    }
                    v15 := []int{}
                    _ = v15
                }
                _ = v9
            } else {
                switch ("" + v4) {
                case "hello":
                    var v16 uint = 2
                    _ = v16
                }
                v1 += v1
            }
            _ = fi7
            _ = v8
        }
        v17 := GenS1{F0: 0, F1: (v4 + "a")}
        _ = v17
    }
    if ((uint64(v5) != 788) && false) {
        v1 -= -158
        v18 := GenS1{F0: 4294967295, F1: ""}
        _ = v18
    } else {
        v19 := GenS2{F0: 0, F1: false, F2: 1}
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 int = fi20
        v4 = "foo bar"
        _ = fi20
        _ = v21
    }
    v22 := []int{(v2 * (-(v2))), v2, v2}
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := GenS1{F0: 0, F1: (("a" + v4) + v4)}
        if len(v22) > 0 {
            v22[0] = 128
        }
        var v25 uint16 = 2
        _ = fi23
        _ = v24
        _ = v25
    }
    switch 0 {
    case 878:
        var v26 int8 = int8(v6)
        _ = v26
    case 4294967295:
        if len(v22) > 0 {
            v22[0] = 127
        }
    }
    if (!true) {
        v6 = (v6 + v6)
        v3 += ((v3 / v3) / v3)
    } else {
        v22 = append(v22, (v2 % ((v2) | 1)))
        v27 := GenS1{F0: 929, F1: v4}
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        v1 += v1
        _ = fi28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 float32 = (v3 - 0.1)
        fmt.Println(v6, v4)
        _ = fi29
        _ = v30
    }
    v3 -= (((((v3 - (v3 * (((v3 - (v3 + (v3 + v3))) + ((1e10 / -2.25) - ((v3 * 0.1) * v3))) - ((v3 / v3) - (((v3 / -2.25) / v3) + 1.5))))) * (0.1 - (v3 + v3))) / 3.14159) * 0.0) * 3.14159)
    v31 := GenS2{F0: -804, F1: true, F2: 945}
    for fi32 := 0; fi32 < 10; fi32++ {
        fmt.Println(v5, v1)
        if (4294967295 == 42) {
            v3 = -2.25
            var v33 int16 = 254
            var v34 float64 = (((((1e10 + float64(3.14159)) / (1e10 / float64(-2.25))) + 1e10) - float64(v3)) + float64(((v3 - float32(-2.25)) * v3)))
            _ = v33
            _ = v34
        } else {
            v1 = (v1 * (v1 / ((v1) | 1)))
            v3 = ((-2.25 * ((float32(1e10) - float32(((3.14159 - 0.0) - -2.25))) * v3)) - v3)
        }
        _ = fi32
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 int64 = -197
        v37 := GenS3{F0: 4294967295, F1: 0, F2: uint64(v5)}
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
