package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int64; F1 int32; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int64, p1 string) int64 {
    p1 = p1
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v7
    return p0
}
func genFn1(p0 uint8, p1 uint8) uint32 {
    var v8 float32 = 1.5
    p1 += 19
    v8 -= v8
    if true {
        if (!(!false)) {
            var v9 float32 = float32(((1.5 / 1.0) * 0.0))
            switch -9223372036854775808 {
            case 545:
                if (702 >= 615) {
                    switch 97 {
                    case 2147483647:
                        var v10 uint32 = 2147483648
                        _ = v10
                    default:
                        var v11 int16 = -949
                        _ = v11
                    }
                }
            case -235:
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 int64 = 139
                    v14 := []int{}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            default:
                v9 = 0.0
            }
            var v15 uint8 = p0
            _ = v15
            _ = v9
        }
        v8 += (0.1 - (v8 * (float32((1.5 * (float64(0.1) + 3.14159))) / 1.5)))
    } else {
        if false {
            var v16 uint32 = 4294967295
            v17 := []int{0, 42, 9223372036854775807, 192}
            v18 := GenS1{F0: 0, F1: true, F2: 32768}
            _ = v16
            _ = v17
            _ = v18
        }
    }
    var v19 float32 = (v8 + v8)
    _ = p0
    _ = p1
    _ = v19
    _ = v8
    return 2
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []uint32{2147483648}
    var v3 uint16 = 42
    var v4 int8 = 127
    var v5 uint32 = 857
    v6 := make(map[uint32]uint32)
    if len(v2) > 0 {
        v2[0] = (v5 / ((0) | 1))
    }
    v20 := GenS1{F0: 1, F1: false, F2: v1}
    if true {
        v3 = v3
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 uint64 = 311
            _ = fi21
            _ = v22
        }
        v23 := GenS2{F0: 1, F1: -32768}
        _ = v23
    } else {
        v24 := GenS3{F0: 1, F1: 1, F2: 65535}
        v25 := GenS1{F0: int32(v5), F1: (-2.25 < 3.14159), F2: v1}
        v4 += (v4 & (v4 | v4))
        _ = v24
        _ = v25
    }
    if (true || true) {
        if ((725 > 4294967295) || ((v5 > v5) || false)) {
            v4 -= v4
        }
        if (genFn0(-998, "a") >= genFn0(9223372036854775807, "\n\t")) {
            var v26 uint32 = v5
            var v27 int16 = 42
            if (v27 != (v27 << ((1) & 63))) {
                v2 = append(v2, v5)
            }
            _ = v26
            _ = v27
        }
    }
    var v28 uint16 = (v3 | 1)
    var v29 int16 = int16(v3)
    switch v1 {
    case -111:
        if (272 < 946) {
            v29 += (-(v29))
        }
    case 408:
        for fi30 := 0; fi30 < 3; fi30++ {
            var v31 int8 = int8(v5)
            var v32 int16 = (v29 % ((v29) | 1))
            switch (v1 % ((fi30) | 1)) {
            case -140:
                if (!(("a" + "hello") > "\n\t")) {
                    for fi33 := 0; fi33 < 16; fi33++ {
                        v34 := GenS1{F0: -692, F1: (-1 >= 679), F2: fi30}
                        _ = fi33
                        _ = v34
                    }
                } else {
                    var v35 string = "a"
                    switch v35 {
                    case "a":
                        for fi36 := 0; fi36 < 20; fi36++ {
                            if true {
                                var v37 int32 = -2147483648
                                _ = v37
                            }
                            if false {
                                v32 = v32
                                fmt.Println(v1, v35)
                            } else {
                                var v38 bool = false
                                _ = v38
                            }
                            v39 := GenS4{F0: 42}
                            _ = fi36
                            _ = v39
                        }
                    case "":
                        v40 := GenS2{F0: -529, F1: v29}
                        _ = v40
                    default:
                        if (v5 < 147) {
                            var v41 bool = (9223372036854775807 < -1)
                            var v42 uint8 = 0
                            _ = v41
                            _ = v42
                        }
                    }
                    v43 := GenS3{F0: 9223372036854775807, F1: -2147483648, F2: 92}
                    _ = v35
                    _ = v43
                }
            case 9223372036854775807:
                v44 := GenS1{F0: 1, F1: (2 < 2), F2: (^(fi30))}
                _ = v44
            case 128:
                if true {
                    v45 := GenS5{F0: 595}
                    var v46 uint16 = 65535
                    _ = v45
                    _ = v46
                }
            }
            _ = fi30
            _ = v31
            _ = v32
        }
    case 32768:
        var v47 int64 = 0
        _ = v47
    default:
        var v48 string = "\n\t"
        _ = v48
    }
    var v49 uint8 = 128
    _ = len(v6)
    v1 = (v1 | v1)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
