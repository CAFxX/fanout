package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int32; F1 int; F2 int32 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint) uint16 {
    p0 += 2
    v4 := []int{127, -539, 127, 32768}
    v5 := GenS1{F0: p0, F1: 18}
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    _ = len(v4)
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 int64 = 0
    var v2 int64 = (v1 - 42)
    v3 := make(map[string]string)
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := []int{-2147483648, (fi6 / ((-806) | 1)), 9223372036854775807}
        for fi8 := 0; fi8 < 7; fi8++ {
            v2 = v1
            _ = fi8
        }
        v1 = (v2 - v2)
        _ = fi6
        _ = v7
    }
    if true {
        var v9 uint64 = 0
        _ = v9
    } else {
        v1 = (v2 - v2)
        v1 = (v2 << ((255) & 63))
        fmt.Println(v2)
    }
    v1 -= (v1 | (v2 - v1))
    v2 = v2
    fmt.Println(v2, v1)
    switch "a" {
    case "\n\t":
        var v10 int32 = 1
        _ = v10
    default:
        _ = len(v3)
    }
    if true {
        v11 := []int{128, 9223372036854775807}
        _ = len(v11)
        for fi12 := 0; fi12 < 8; fi12++ {
            v3[("\n\t" + "\n\t")] = "\n\t"
            _ = fi12
        }
        _ = v11
    }
    if true {
        if false {
            var v13 float32 = (0.0 * (-2.25 - (3.14159 + (0.1 / (-2.25 + float32(float64(float32(float64(float32(-2.25))))))))))
            v13 -= 1.5
            var v14 int64 = v2
            _ = v13
            _ = v14
        } else {
            v1 = v1
            v15 := GenS2{F0: -2147483648, F1: 128, F2: 2147483647}
            _ = v15
        }
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := GenS3{F0: 2147483647, F1: -2147483648}
            v18 := v3["hello"]
            _ = fi16
            _ = v17
            _ = v18
        }
        switch -585 {
        case 2147483647:
            if false {
                v19 := GenS4{F0: 32, F1: 32767}
                var v20 float32 = 3.14159
                _ = v19
                _ = v20
            } else {
                v21 := GenS2{F0: 714, F1: 9223372036854775807, F2: -2147483648}
                v22 := []int{32768, -856}
                _ = v21
                _ = v22
            }
        case 1:
            for fi23 := 0; fi23 < 6; fi23++ {
                if ((!((true && false) || true)) && (!true)) {
                    if true {
                        for fi24 := 0; fi24 < 13; fi24++ {
                            var v25 uint64 = 4294967295
                            _ = fi24
                            _ = v25
                        }
                    } else {
                        v26 := GenS5{F0: 847, F1: 42, F2: 255}
                        v1 += (v1 & (-(v1)))
                        v27 := GenS4{F0: 18446744073709551615, F1: 1}
                        _ = v26
                        _ = v27
                    }
                    v2 -= v1
                } else {
                    var v28 int64 = 1
                    v28 -= v28
                    _ = v28
                }
                _ = fi23
            }
        case 32767:
            for fi29 := 0; fi29 < 16; fi29++ {
                fmt.Println(v1)
                if false {
                    var v30 string = "hello"
                    var v31 string = ""
                    v32 := []int{-683, fi29}
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 8; fi33++ {
                        if (v2 == v2) {
                            var v34 int16 = 0
                            switch 32767 {
                            case 4294967295:
                                if (v34 < (v34 >> ((65535) & 63))) {
                                    var v35 uint = 65535
                                    var v36 float32 = 0.1
                                    v37 := []int{}
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                } else {
                                    v38 := GenS6{F0: 2}
                                    v39 := []int{(fi33 ^ fi29), fi33, fi33, fi33}
                                    fmt.Println(fi33)
                                    _ = v38
                                    _ = v39
                                }
                            case -2147483648:
                                v40 := []int{}
                                _ = v40
                            }
                            var v41 int64 = -9223372036854775808
                            v42 := []int{fi33}
                            _ = v34
                            _ = v41
                            _ = v42
                        }
                        _ = fi33
                    }
                }
                var v43 bool = true
                _ = fi29
                _ = v43
            }
        default:
            fmt.Println(v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}
