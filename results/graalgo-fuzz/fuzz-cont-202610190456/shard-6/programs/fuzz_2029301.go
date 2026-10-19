package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 int16) int8 {
    v4 := []int{2147483647, 42}
    p0 += 693
    var v5 uint32 = 1
    fmt.Println(v5)
    p0 -= p0
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1() int {
    v6 := []int{671, 32767, -2147483648, 42}
    var v7 uint64 = 42
    if (-1 != 842) {
        var v8 uint32 = 377
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 string = "a"
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 32767, F1: (v7 % ((18446744073709551615) | 1)), F2: 1}
        _ = v11
        _ = v8
    } else {
        v7 -= 18446744073709551615
        v12 := GenS1{F0: 42, F1: (^(v7)), F2: 113}
        var v13 float32 = float32(1e10)
        _ = v12
        _ = v13
    }
    _ = v6
    _ = v7
    return 795
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []string{}
    v3 := make(map[int]bool)
    switch 212 {
    case 127:
        v14 := GenS1{F0: 1000, F1: 18446744073709551615, F2: 125}
        _ = v14
    case -746:
        fmt.Println(v1)
    case -1:
        if false {
            switch -9223372036854775808 {
            case 128:
                for fi15 := 0; fi15 < 7; fi15++ {
                    v3[int(v1)] = (1 != 1)
                    var v16 uint8 = 255
                    _ = fi15
                    _ = v16
                }
            }
        }
    default:
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 uint64 = 11
            _ = fi17
            _ = v18
        }
    }
    var v19 float64 = ((1.5 + 1.5) / ((1.5 + 0.0) - 0.0))
    var v20 float64 = v19
    for fi21 := 0; fi21 < 13; fi21++ {
        v19 = v20
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi21
    }
    switch -1 {
    case -9223372036854775808:
        for fi22 := 0; fi22 < 18; fi22++ {
            for fi23 := 0; fi23 < 6; fi23++ {
                fmt.Println(v20, v1)
                _ = fi23
            }
            _ = fi22
        }
    case 127:
        for fi24 := 0; fi24 < 5; fi24++ {
            v1 += v1
            v19 = (float64(0.0) - v19)
            v25 := GenS1{F0: 1000, F1: 0, F2: uint8(v1)}
            _ = fi24
            _ = v25
        }
    case -326:
        switch "" {
        case "":
            if (!true) {
                var v26 string = "foo bar"
                var v27 int8 = -56
                _ = v26
                _ = v27
            } else {
                var v28 int64 = 42
                v29 := GenS2{F0: 0, F1: 18446744073709551615}
                for fi30 := 0; fi30 < 15; fi30++ {
                    v31 := v3[-579]
                    _ = fi30
                    _ = v31
                }
                _ = v28
                _ = v29
            }
        case "hello":
            var v32 int32 = 0
            _ = v32
        }
    }
    v33 := GenS2{F0: 42, F1: 287}
    fmt.Println(v20, v20, v20)
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 int8 = -1
        var v36 string = ("hello" + "")
        _ = fi34
        _ = v35
        _ = v36
    }
    if false {
        if ((v1 >> ((2) & 63)) != v1) {
            v20 = v20
        }
        v37 := GenS2{F0: -99, F1: 543}
        _ = v37
    }
    fmt.Println(v19)
    var v38 int64 = 619
    v20 = (0.1 / (v20 / float64((1e10 * 1e10))))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v38
}
