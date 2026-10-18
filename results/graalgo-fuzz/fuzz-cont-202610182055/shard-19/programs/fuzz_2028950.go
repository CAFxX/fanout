package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int16; F2 uint }
type GenS4 struct { F0 uint }
func genFn0() uint64 {
    v6 := []int{32767, 128}
    switch 308 {
    case 9223372036854775807:
        var v7 bool = (true || false)
        _ = v7
    case 128:
        var v8 int64 = 881
        _ = v8
    }
    var v9 float64 = (1e10 / 3.14159)
    v10 := GenS1{F0: 42, F1: 127, F2: 958}
    var v11 bool = false
    _ = v10
    _ = v11
    _ = v6
    _ = v9
    return 4294967295
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]bool)
    var v3 float32 = (0.1 / 1.5)
    var v4 int8 = -124
    var v5 int16 = 788
    for fi12 := 0; fi12 < 18; fi12++ {
        v1 = (float64(v3) / v1)
        var v13 uint64 = genFn0()
        if (-2.25 < (v3 / 0.1)) {
            switch fi12 {
            case -1:
                fmt.Println(v5, v4)
            case 2147483647:
                if (42 > -9223372036854775808) {
                    v3 += (v3 / v3)
                } else {
                    v14 := GenS2{F0: 878}
                    switch fi12 {
                    case 9223372036854775807:
                        v5 -= v5
                    case 127:
                        if ((((float64((v3 - v3)) / (v1 - v1)) - v1) / v1) >= v1) {
                            v15 := GenS1{F0: 4294967295, F1: v4, F2: uint16(v13)}
                            var v16 int32 = -2147483648
                            _ = v15
                            _ = v16
                        }
                    }
                    v17 := GenS3{F0: (v13 >> ((18446744073709551615) & 63)), F1: v5, F2: 659}
                    _ = v14
                    _ = v17
                }
            default:
                if (42 != 127) {
                    v4 = 0
                    v18 := GenS1{F0: (v13 - v13), F1: v4, F2: 65535}
                    _ = v18
                } else {
                    if (false && (2 < 1)) {
                        v19 := GenS3{F0: v13, F1: -639, F2: 4294967295}
                        var v20 int32 = int32(v5)
                        if (!((!true) || true)) {
                            var v21 float32 = (1e10 / 1.0)
                            fmt.Println(v21, fi12)
                            _ = v21
                        }
                        _ = v19
                        _ = v20
                    } else {
                        var v22 float64 = v1
                        _ = v22
                    }
                    v4 = v4
                    for fi23 := 0; fi23 < 5; fi23++ {
                        if true {
                            if (!true) {
                                var v24 int = (fi12 + fi12)
                                if (true || true) {
                                    switch fi12 {
                                    case -9223372036854775808:
                                        v3 -= v3
                                    case 2147483647:
                                        var v25 int = fi23
                                        _ = v25
                                    case 127:
                                        var v26 bool = (85 >= 840)
                                        _ = v26
                                    }
                                    var v27 float64 = (float64(v3) + (v1 + ((0.1 + (-2.25 + ((v1 - ((v1 - 1.5) * (v1 - 1.5))) * float64(((((((float32(v1) + -2.25) + v3) / 1.0) * (3.14159 - (1.5 - v3))) * v3) + v3))))) / 1e10)))
                                    var v28 uint8 = 18
                                    fmt.Println(v3, v13, v28)
                                    v29 := GenS4{F0: 872}
                                    fmt.Println(v4, v1, v5)
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                _ = v24
                            } else {
                                v3 = (1e10 + v3)
                            }
                            v5 -= v5
                        } else {
                            var v30 int16 = v5
                            _ = v30
                        }
                        v31 := GenS2{F0: 815}
                        v13 += (v13 * (v13 | (v13 * (v13 ^ v13))))
                        _ = fi23
                        _ = v31
                    }
                }
            }
            fmt.Println(fi12, v3)
            v32 := []int{(fi12 + (fi12 % (((fi12 * fi12)) | 1)))}
            _ = v32
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v1)
    v2[4294967295] = (3.14159 > float64(3.14159))
    v33 := GenS1{F0: 2, F1: v4, F2: 65535}
    fmt.Println(v5, v4, v1)
    v2[65535] = false
    var v34 uint64 = 235
    fmt.Println(v3, v4)
    if (!false) {
        fmt.Println(v5)
        for fi35 := 0; fi35 < 4; fi35++ {
            var v36 uint16 = 895
            v3 += (((v3 * (0.1 - v3)) + (v3 * (1e10 + v3))) / 0.1)
            v37 := GenS1{F0: (v34 ^ genFn0()), F1: v4, F2: 2}
            _ = fi35
            _ = v36
            _ = v37
        }
    } else {
        v38 := GenS1{F0: (v34 & v34), F1: (v4 * 100), F2: 2}
        v39 := GenS1{F0: (v34 * genFn0()), F1: int8(v5), F2: 590}
        v5 = v5
        _ = v38
        _ = v39
    }
    v40 := v2[2]
    var v41 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
