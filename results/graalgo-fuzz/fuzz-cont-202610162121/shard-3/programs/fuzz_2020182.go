package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int16; F1 int16; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 string }
func genFn0(p0 int, p1 int, p2 int) int32 {
    var v5 int64 = 42
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}
func genFn1(p0 int64, p1 int8) int64 {
    v6 := GenS1{F0: 1}
    var v7 int16 = 548
    v8 := []int{625}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 65535
    v2 := []uint32{2147483648, 2147483648, 1}
    var v3 uint64 = 0
    var v4 uint64 = v3
    for fi9 := 0; fi9 < 8; fi9++ {
        v4 -= v4
        _ = len(v2)
        _ = fi9
    }
    var v10 bool = (128 < -1)
    v2 = append(v2, 4294967295)
    v11 := GenS1{F0: 1}
    var v12 int = -2147483648
    v1 += v1
    var v13 string = ""
    var v14 float64 = (0.1 + (1.5 - 1e10))
    v15 := GenS1{F0: (v3 & 594)}
    for fi16 := 0; fi16 < 3; fi16++ {
        if len(v2) > 0 {
            v2[0] = 736
        }
        var v17 int64 = 9223372036854775807
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(v14, v3, v12)
            var v19 int32 = int32(v4)
            var v20 int64 = v17
            v20 -= v17
            var v21 int32 = v19
            switch 1 {
            case 42:
                v22 := GenS1{F0: (v4 ^ (v3 >> ((227) & 63)))}
                _ = v22
            case 2147483647:
                var v23 int16 = 1000
                _ = v23
            default:
                v17 = int64(fi16)
            }
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 1; fi25++ {
            fmt.Println(v10)
            var v26 bool = false
            if ((float64((1e10 - (float32(float64(0.1)) * (0.0 - ((float32(float64(0.1)) / (-2.25 - 0.0)) * 0.0))))) * 1e10) == (v14 - 3.14159)) {
                v2 = append(v2, 0)
                switch fi25 {
                case 4294967295:
                    v10 = false
                default:
                    if ((^(v4)) == (v4 % ((v4) | 1))) {
                        var v27 string = v13
                        if (v27 >= (v13 + v27)) {
                            v28 := GenS1{F0: v3}
                            _ = v28
                        }
                        fmt.Println(v3)
                        _ = v27
                    } else {
                        if (630 == 0) {
                            v4 += v3
                            v10 = (!(!(((^(v4)) >= 1) && v26)))
                            var v29 int32 = 0
                            _ = v29
                        } else {
                            if v26 {
                                v10 = (v14 != 1e10)
                            } else {
                                v12 += v12
                                v14 -= (0.1 / v14)
                                var v30 float64 = 1.5
                                _ = v30
                            }
                            switch v12 {
                            case 1:
                                v13 = ((v13 + "a") + "\n\t")
                            case 234:
                                var v31 int64 = 0
                                v31 += 9223372036854775807
                                v4 -= v4
                                _ = v31
                            }
                            v10 = v10
                        }
                    }
                }
            } else {
                var v32 float64 = 1e10
                _ = v32
            }
            _ = fi25
            _ = v26
        }
        var v33 int16 = 42
        switch 32767 {
        case 32767:
            var v34 uint64 = 0
            _ = v34
        default:
            v10 = (!(((0.1 - (float64(0.0) / (v14 * v14))) + v14) < v14))
        }
        _ = fi24
        _ = v33
    }
    if v10 {
        fmt.Println(v13, v13, v14)
        switch v13 {
        case "a":
            for fi35 := 0; fi35 < 1; fi35++ {
                v12 = 2147483647
                var v36 float32 = (((0.0 * 3.14159) - (((1e10 - (3.14159 + 0.1)) + 0.1) + 0.1)) * 1.5)
                switch int(v1) {
                case -134:
                    var v37 int64 = 42
                    _ = v37
                default:
                    if true {
                        v14 += (0.0 * v14)
                        v3 -= 383
                        if true {
                            fmt.Println(v36, fi35)
                            v38 := GenS2{F0: genFn0((fi35 - fi35), fi35, -2147483648), F1: 69}
                            if v10 {
                                fmt.Println(v12, v14)
                                v12 = (-(v12))
                            } else {
                                v2 = append(v2, 4294967295)
                            }
                            _ = v38
                        }
                    } else {
                        if ((((v13 + v13) > v13) || v10) || (-390 > 629)) {
                            v2 = append(v2, uint32(v1))
                        }
                        v4 += 42
                    }
                }
                _ = fi35
                _ = v36
            }
        case "hello":
            if len(v2) > 0 {
                v2[0] = 0
            }
        default:
            for fi39 := 0; fi39 < 12; fi39++ {
                if (float32(3.14159) > 3.14159) {
                    if (((v4 * (^(v4))) > (v4 | v3)) && v10) {
                        if (!(9223372036854775807 >= (-(v12)))) {
                            var v40 float32 = 3.14159
                            _ = v40
                        }
                        var v41 uint16 = (v1 & v1)
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                        _ = v41
                    } else {
                        v42 := GenS2{F0: -498, F1: -1}
                        v43 := GenS3{F0: -32768, F1: 42, F2: -128}
                        _ = v42
                        _ = v43
                    }
                    v44 := GenS4{F0: -57}
                    _ = v44
                } else {
                    var v45 float64 = v14
                    _ = v45
                }
                fmt.Println(v13)
                _ = fi39
            }
        }
        v2 = append(v2, 619)
    } else {
        v10 = ((^(v1)) <= v1)
        v14 = v14
    }
    for fi46 := 0; fi46 < 13; fi46++ {
        v47 := GenS5{F0: 425, F1: v13}
        var v48 uint8 = 124
        _ = fi46
        _ = v47
        _ = v48
    }
    v10 = (v4 > 0)
    v14 += v14
    fmt.Println(v10, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
