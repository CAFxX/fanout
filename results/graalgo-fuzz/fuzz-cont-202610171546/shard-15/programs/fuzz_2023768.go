package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int16 }
func genFn0() int {
    if (!true) {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 int8 = 0
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 bool = true
                v6 -= (v6 % ((v6) | 1))
                _ = fi7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
        var v9 int8 = -95
        v9 = 41
        _ = v9
    } else {
        var v10 uint32 = 4294967295
        if false {
            v10 = v10
        } else {
            v10 -= v10
            switch (("\n\t" + "a") + "") {
            case "hello":
                v11 := GenS1{F0: 1, F1: (!true)}
                _ = v11
            default:
                v12 := []int{}
                _ = v12
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := []int{(fi13 ^ 32767)}
                if true {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        switch fi13 {
                        case 42:
                            v10 += v10
                        case 657:
                            if false {
                                v10 -= (v10 << ((1) & 63))
                                if (false && ((false && false) && true)) {
                                    var v16 int64 = 9223372036854775807
                                    _ = v16
                                } else {
                                    var v17 int16 = -159
                                    var v18 uint64 = 596
                                    v14 = append(v14, 440)
                                    _ = v17
                                    _ = v18
                                }
                                if (false || false) {
                                    var v19 int8 = -128
                                    v20 := GenS1{F0: uint(v10), F1: (false || (v10 != v10))}
                                    _ = v19
                                    _ = v20
                                }
                            }
                        case 128:
                            _ = len(v14)
                        default:
                            if (!false) {
                                var v21 int64 = -230
                                v10 = (v10 / ((v10) | 1))
                                var v22 float64 = 0.0
                                _ = v21
                                _ = v22
                            } else {
                                v10 = (v10 % ((v10) | 1))
                                v14 = append(v14, fi13)
                                var v23 float64 = 1.5
                                _ = v23
                            }
                        }
                        v10 = (^(v10))
                        var v24 uint32 = (^(v10))
                        if false {
                            var v25 int32 = 1
                            var v26 int8 = 23
                            _ = v25
                            _ = v26
                        }
                        _ = fi15
                        _ = v24
                    }
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    var v27 float32 = ((float32((1e10 / ((0.0 + 1.5) + (0.1 / (((((((3.14159 - 0.1) + float64(float32((1.5 + ((3.14159 + (float64((0.0 / (float32(1.5) + float32((float64((3.14159 + 3.14159)) * (1e10 * -2.25)))))) - (-2.25 + 3.14159))) / (3.14159 + ((((float64(1e10) + float64(1e10)) - float64((float32(3.14159) / 1.0))) + 0.1) / -2.25))))))) * (1e10 / -2.25)) / ((0.1 * 1e10) - float64(((1.5 + (3.14159 + (0.0 - 0.1))) + 1e10)))) + 0.1) + 0.0) + 1e10))))) * -2.25) * (3.14159 * (1e10 - 0.1)))
                    _ = v27
                }
                fmt.Println(fi13, v10)
                _ = fi13
                _ = v14
            }
        }
        _ = v10
    }
    var v28 uint8 = 213
    v29 := []int{127, -906, -972}
    fmt.Println(v28)
    var v30 uint64 = 105
    _ = v28
    _ = v29
    _ = v30
    return 32768
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[int]uint32)
    var v3 uint32 = 1
    var v4 uint64 = 2
    var v31 uint8 = 0
    var v32 bool = false
    if true {
        v33 := []int{}
        v34 := GenS1{F0: 42, F1: v32}
        _ = v33
        _ = v34
    } else {
        switch ("hello" + ("hello" + ("foo bar" + ""))) {
        case "a":
            v35 := v2[v1]
            _ = v35
        case "\n\t":
            v36 := GenS2{F0: 837}
            _ = v36
        case "hello":
            for fi37 := 0; fi37 < 10; fi37++ {
                v31 = 42
                var v38 float32 = 1e10
                _ = fi37
                _ = v38
            }
        }
        var v39 int32 = -880
        _ = v39
    }
    v40 := GenS1{F0: 18446744073709551615, F1: (1 < 0)}
    v31 = uint8(v3)
    var v41 uint16 = 0
    var v42 float32 = -2.25
    v3 += (v3 - v3)
    fmt.Println(v4)
    if ((v31 + v31) > (^(v31))) {
        v1 = (v1 % ((v1) | 1))
    }
    var v43 float32 = (v42 + 1.5)
    var v44 float32 = v43
    if (-2.25 != (float32((float64(float32(-2.25)) - ((float64(v44) / 1.0) + 0.0))) - 1.5)) {
        var v45 string = "a"
        var v46 int = v1
        _ = v45
        _ = v46
    }
    v44 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
