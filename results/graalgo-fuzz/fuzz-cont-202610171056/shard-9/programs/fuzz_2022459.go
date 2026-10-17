package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint8, p2 uint8) uint64 {
    v4 := []int{}
    var v5 int64 = 0
    if true {
        p0 -= p0
        var v6 float64 = 0.0
        v5 += (-(v5))
        _ = v6
    } else {
        v7 := GenS1{F0: 0, F1: 804, F2: 712}
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(v5)
            if true {
                for fi9 := 0; fi9 < 12; fi9++ {
                    switch ("" + "") {
                    case "":
                        if (65535 >= 65535) {
                            var v10 int = 910
                            if (!((false || (42 < 1)) || true)) {
                                p2 = (p2 * p2)
                                var v11 int32 = -752
                                switch v10 {
                                case 128:
                                    v12 := GenS1{F0: 18446744073709551615, F1: (p0 / (((p0 - p0)) | 1)), F2: 42}
                                    _ = v12
                                case 32768:
                                    _ = len(v4)
                                }
                                _ = v11
                            }
                            _ = v10
                        } else {
                            var v13 uint = 1
                            var v14 uint8 = p1
                            _ = v13
                            _ = v14
                        }
                    case "a":
                        v5 = v5
                    default:
                        if false {
                            p2 += (p1 / ((p2) | 1))
                            _ = len(v4)
                            p1 = (p2 >> ((214) & 63))
                        }
                    }
                    _ = fi9
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v15 int = fi8
                _ = v15
            }
            _ = fi8
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 72
    var v3 int32 = 0
    var v16 int32 = v3
    v17 := GenS1{F0: 4294967295, F1: 552, F2: 42}
    var v18 int32 = v3
    fmt.Println(v16)
    if false {
        var v19 bool = true
        _ = v19
    }
    var v20 uint32 = 0
    if true {
        v21 := GenS1{F0: genFn0(239, 128, 2), F1: 0, F2: 42}
        v3 -= v18
        _ = v21
    } else {
        v18 -= v16
    }
    var v22 uint64 = 2
    v23 := GenS2{F0: 65535, F1: (v20 >> ((uint(v20)) & 63)), F2: (v22 - (v22 % ((v22) | 1)))}
    fmt.Println(v2, v20, v20)
    v2 = v2
    v2 += v2
    v24 := GenS1{F0: v22, F1: 83, F2: v22}
    v25 := GenS1{F0: (v22 - genFn0(uint16(v20), 108, 42)), F1: 0, F2: v22}
    if false {
        fmt.Println(v2)
    } else {
        var v26 uint16 = 217
        _ = v26
    }
    switch 495 {
    case 0:
        for fi27 := 0; fi27 < 2; fi27++ {
            v22 += 2
            if (!true) {
                v28 := GenS2{F0: 65535, F1: v20, F2: (v22 / (((^(v22))) | 1))}
                _ = v28
            }
            _ = fi27
        }
    case -859:
        switch ("a" + "\n\t") {
        case "hello":
            if ((v18 + v18) <= v3) {
                v29 := GenS3{F0: v22, F1: (^(v20))}
                _ = v29
            } else {
                switch "hello" {
                case "":
                    fmt.Println(v2, v22)
                }
                var v30 float32 = 3.14159
                var v31 uint32 = (v20 << ((18446744073709551615) & 63))
                _ = v30
                _ = v31
            }
        default:
            var v32 int = 128
            _ = v32
        }
    default:
        v20 -= 2147483647
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        v34 := []int{(fi33 + (fi33 * fi33))}
        if (v22 == (v22 & v22)) {
            fmt.Println(fi33, fi33, v16)
            var v35 int16 = 337
            var v36 uint16 = 0
            _ = v35
            _ = v36
        } else {
            v37 := GenS4{F0: -2147483648}
            _ = v37
        }
        if len(v34) > 0 {
            v34[0] = (fi33 / (((fi33 ^ fi33)) | 1))
        }
        _ = fi33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
