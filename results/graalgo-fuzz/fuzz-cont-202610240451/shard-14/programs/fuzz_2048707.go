package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 bool; F1 uint; F2 int32 }
type GenS3 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint, p1 uint32, p2 uint16) int {
    var v7 uint16 = (p2 + (^(p2)))
    p2 = (p2 / (((p2 | (^(v7)))) | 1))
    v7 = (p2 % ((32767) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1() bool {
    var v8 int64 = 0
    var v9 float32 = 1.5
    var v10 uint64 = 78
    _ = v10
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := []uint32{2147483647, 2147483648, 431, 42}
    var v2 float64 = 0.0
    v3 := make(map[int]string)
    var v4 float32 = float32((1e10 - 0.1))
    var v5 uint = 18446744073709551615
    var v6 int64 = 1
    v11 := GenS1{F0: false, F1: 1000}
    delete(v3, 2147483647)
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 float64 = 1.5
        v13 -= float64(1e10)
        _ = fi12
        _ = v13
    }
    var v14 int32 = 0
    v14 = (v14 | (v14 << (((v5 + v5)) & 63)))
    for fi15 := 0; fi15 < 18; fi15++ {
        switch ("a" + "") {
        case "":
            for fi16 := 0; fi16 < 3; fi16++ {
                v17 := GenS2{F0: false, F1: v5, F2: v14}
                v18 := GenS1{F0: true, F1: int16(v5)}
                if (v6 <= 1) {
                    v14 = -173
                    if false {
                        var v19 int8 = 100
                        _ = v19
                    }
                } else {
                    v2 = v2
                    var v20 uint32 = 2147483647
                    _ = v20
                }
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        _ = fi15
    }
    fmt.Println(v2, v4)
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 int8 = 42
        for fi23 := 0; fi23 < 14; fi23++ {
            v2 += 0.0
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    v5 += 1
    v24 := v3[genFn0(18446744073709551615, 42, 415)]
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 uint16 = 65535
        v5 = v5
        _ = fi25
        _ = v26
    }
    if (-1 < -9223372036854775808) {
        var v27 int8 = 1
        v2 = v2
        _ = v27
    }
    _ = len(v1)
    var v28 uint = v5
    for fi29 := 0; fi29 < 6; fi29++ {
        if (0 >= (v14 - (v14 | v14))) {
            v6 -= v6
            var v30 int8 = -79
            v28 = (^(v28))
            _ = v30
        } else {
            var v31 uint8 = 42
            _ = v31
        }
        for fi32 := 0; fi32 < 3; fi32++ {
            switch genFn0(974, 42, 2) {
            case 42:
                fmt.Println(v4, v6, v2)
            default:
                if ((177 > 128) || (2147483648 >= 311)) {
                    var v33 int64 = v6
                    _ = v33
                }
                fmt.Println(fi29, v5)
            }
            _ = fi32
        }
        v34 := GenS3{F0: v5, F1: 1}
        _ = fi29
        _ = v34
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        var v36 uint32 = 65535
        for fi37 := 0; fi37 < 12; fi37++ {
            switch fi37 {
            case -949:
                if ((458 <= v5) && (false || true)) {
                    var v38 int = 2147483647
                    _ = v38
                }
            case 1:
                v4 = v4
            case 32767:
                if (-524 != 32767) {
                    if len(v1) > 0 {
                        v1[0] = v36
                    }
                    var v39 uint8 = 0
                    _ = v39
                }
            }
            _ = fi37
        }
        if true {
            v3[(fi35 & (fi35 / (((fi35 & fi35)) | 1)))] = v24
            for fi40 := 0; fi40 < 5; fi40++ {
                v41 := GenS2{F0: true, F1: 19, F2: -2147483648}
                if true {
                    v6 += v6
                    v1 = append(v1, 0)
                }
                _ = fi40
                _ = v41
            }
        } else {
            v36 = (v36 * 323)
        }
        _ = fi35
        _ = v36
    }
    if (false || (!(false || (false || false)))) {
        var v42 uint32 = uint32(v5)
        _ = len(v3)
        _ = v42
    } else {
        if false {
            v3[42] = v24
            v14 = 519
        }
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
