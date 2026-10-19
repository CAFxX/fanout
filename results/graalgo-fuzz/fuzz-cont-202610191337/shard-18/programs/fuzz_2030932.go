package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int64; F1 int32; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool }
func genFn0(p0 int32) uint16 {
    var v7 int16 = -32768
    switch "" {
    case "hello":
        var v8 uint32 = 2147483648
        _ = v8
    case "\n\t":
        if false {
            v9 := GenS1{F0: 104, F1: 42}
            switch ("foo bar" + "") {
            case "a":
                p0 = (p0 / ((0) | 1))
            case "":
                v10 := GenS2{F0: -236, F1: p0, F2: ("" == "hello")}
                _ = v10
            case "foo bar":
                switch 4294967295 {
                case -9223372036854775808:
                    var v11 float64 = float64((-2.25 - (0.0 - 1.5)))
                    _ = v11
                case 32768:
                    v12 := []int{32768, 42, 127, 42}
                    _ = v12
                case -2147483648:
                    var v13 int8 = -128
                    _ = v13
                }
            }
            v14 := GenS1{F0: 32767, F1: 2147483647}
            _ = v14
            _ = v9
        }
    default:
        v7 -= (v7 + 32767)
    }
    if (58 > 127) {
        var v15 int = -2147483648
        var v16 float64 = float64(3.14159)
        v7 -= v7
        _ = v15
        _ = v16
    } else {
        var v17 uint64 = 198
        _ = v17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(fi18, v7)
        var v19 uint = 892
        v20 := []int{(fi18 & fi18), 91, int(p0)}
        _ = fi18
        _ = v19
        _ = v20
    }
    if (!(v7 != (v7 + (v7 / (((v7 | -589)) | 1))))) {
        if true {
            var v21 uint8 = 127
            v21 += (^(v21))
            var v22 uint = 1
            _ = v21
            _ = v22
        } else {
            v23 := []int{-1, 1, -467}
            var v24 float32 = 1.5
            v25 := GenS2{F0: 14, F1: -2147483648, F2: (v24 < (-2.25 + 3.14159))}
            _ = v23
            _ = v24
            _ = v25
        }
        var v26 uint32 = 2147483648
        _ = v26
    } else {
        var v27 uint64 = 1
        v28 := GenS2{F0: 42, F1: (-(p0)), F2: (!true)}
        _ = v27
        _ = v28
    }
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 int = 0
    v2 := []int{32767, (v1 - v1), v1}
    var v3 bool = false
    var v4 int8 = 100
    var v5 string = "a"
    v6 := make(map[int]uint32)
    if v3 {
        for fi29 := 0; fi29 < 16; fi29++ {
            v30 := GenS1{F0: 507, F1: 2}
            v31 := GenS2{F0: 9223372036854775807, F1: int32(v4), F2: (0 != 32)}
            delete(v6, (fi29 / (((fi29 ^ v1)) | 1)))
            _ = fi29
            _ = v30
            _ = v31
        }
        v32 := GenS3{F0: -1}
        _ = v32
    }
    var v33 int = (v1 | v1)
    for fi34 := 0; fi34 < 4; fi34++ {
        var v35 string = v5
        v36 := GenS2{F0: 445, F1: 0, F2: v3}
        var v37 uint = 428
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    fmt.Println(v1, v33, v1)
    var v38 int32 = 734
    switch v5 {
    case "":
        v5 = v5
    case "a":
        v1 += 9223372036854775807
    case "\n\t":
        _ = len(v2)
    default:
        if ((0.0 / float32(1e10)) > float32(1e10)) {
            if (!true) {
                var v39 int32 = (v38 * 0)
                var v40 int16 = -468
                _ = v39
                _ = v40
            }
        } else {
            if len(v2) > 0 {
                v2[0] = (-(v33))
            }
            for fi41 := 0; fi41 < 12; fi41++ {
                v42 := GenS4{F0: v3}
                _ = fi41
                _ = v42
            }
            var v43 int8 = v4
            _ = v43
        }
    }
    var v44 float64 = (1e10 / ((0.0 * float64(((-2.25 * 1e10) * (1e10 + ((3.14159 + -2.25) / (float32(1.5) - (0.0 + 0.1))))))) - 1e10))
    _ = len(v6)
    fmt.Println(v38, v5, v1)
    fmt.Println(v33, v38)
    fmt.Println(v4, v33)
    var v45 uint16 = 1
    var v46 int8 = v4
    v3 = ((4294967295 < 1) && v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
