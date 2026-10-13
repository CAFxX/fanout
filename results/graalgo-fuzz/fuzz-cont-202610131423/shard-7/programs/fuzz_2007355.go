package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint64; F1 int8; F2 string }
type GenS4 struct { F0 uint64; F1 int64; F2 int32 }
type GenS5 struct { F0 uint64; F1 int32; F2 string }
func genFn0() int {
    if false {
        var v5 uint32 = 42
        v6 := []int{645}
        _ = v5
        _ = v6
    }
    var v7 int32 = 42
    for fi8 := 0; fi8 < 19; fi8++ {
        v7 = (^(v7))
        v9 := GenS1{F0: 2, F1: 65535, F2: -9223372036854775808}
        var v10 uint16 = 2
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v7
    return 54
}
func genFn1(p0 string, p1 int64) bool {
    var v11 uint32 = 216
    v12 := GenS1{F0: (v11 % (((^(v11))) | 1)), F1: v11, F2: -232}
    p1 -= (p1 << ((1) & 63))
    switch p0 {
    case "":
        v11 = (v11 ^ (v11 - 2147483648))
    case "a":
        v13 := GenS1{F0: v11, F1: (v11 / (((v11 & v11)) | 1)), F2: -9223372036854775808}
        _ = v13
    case "foo bar":
        v14 := GenS2{F0: -1, F1: 9223372036854775807}
        _ = v14
    }
    v15 := GenS1{F0: (v11 * (^(v11))), F1: v11, F2: -9223372036854775808}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v15
    return false
}

func main() {
    var v1 int16 = 1000
    v2 := []string{("" + ""), "hello"}
    v3 := make(map[uint32]uint32)
    v4 := []string{("" + "a"), ("" + "foo bar")}
    fmt.Println(v1)
    switch 42 {
    case -475:
        var v16 int16 = (v1 ^ (v1 * 42))
        _ = v16
    }
    v1 = -980
    v17 := GenS3{F0: 741, F1: 100, F2: "\n\t"}
    if true {
        v1 += v1
        if true {
            switch 9223372036854775807 {
            case 127:
                v1 = 0
            default:
                _ = len(v3)
            }
        } else {
            v1 += v1
        }
    }
    if (-533 > 0) {
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 13; fi19++ {
                v1 = 42
                _ = fi19
            }
            _ = fi18
        }
    } else {
        v1 = (v1 << ((18446744073709551615) & 63))
        for fi20 := 0; fi20 < 5; fi20++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v3)
            var v21 int32 = -2147483648
            _ = fi20
            _ = v21
        }
    }
    var v22 float32 = (3.14159 - float32(float64(float32(0.1))))
    var v23 float32 = 1e10
    for fi24 := 0; fi24 < 10; fi24++ {
        if (!false) {
            var v25 float64 = -2.25
            _ = v25
        }
        switch ("" + ("a" + "a")) {
        case "foo bar":
            var v26 int64 = 604
            _ = v26
        case "hello":
            if (1e10 > 0.0) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v1 = 1000
            }
        }
        _ = fi24
    }
    if false {
        var v27 int = 1
        var v28 float64 = 1.5
        if (false || true) {
            var v29 uint8 = 6
            fmt.Println(v22, v28)
            _ = v29
        }
        _ = v27
        _ = v28
    } else {
        v30 := GenS4{F0: 800, F1: 1, F2: -1}
        if true {
            var v31 uint32 = 2147483648
            for fi32 := 0; fi32 < 18; fi32++ {
                switch genFn0() {
                case 527:
                    if false {
                        v1 = v1
                        var v33 float32 = float32(1.5)
                        _ = v33
                    } else {
                        switch fi32 {
                        case -817:
                            switch genFn0() {
                            case 2147483647:
                                var v34 uint = 4294967295
                                _ = v34
                            case 1:
                                v35 := GenS4{F0: 567, F1: int64(fi32), F2: 1}
                                _ = v35
                            case 0:
                                v1 = 1
                            }
                        case -364:
                            v31 += v31
                        case 782:
                            var v36 int8 = -128
                            _ = v36
                        default:
                            fmt.Println(v1, v23, v31)
                        }
                        var v37 float32 = ((v22 / v23) + float32(0.0))
                        _ = v37
                    }
                }
                _ = len(v3)
                _ = fi32
            }
            _ = v31
        }
        v38 := GenS5{F0: 466, F1: -2147483648, F2: ("a" + "hello")}
        _ = v30
        _ = v38
    }
    v22 += v23
    v39 := GenS4{F0: 1, F1: -9223372036854775808, F2: -830}
    var v40 uint8 = 28
    if false {
        if (!false) {
            var v41 int = 0
            v1 -= 32767
            v42 := v3[2147483648]
            _ = v41
            _ = v42
        }
    } else {
        fmt.Println(v23, v23, v23)
        v22 = ((((v22 / 1.0) - (v22 / v23)) * 1e10) * v23)
        fmt.Println(v23, v22, v22)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
