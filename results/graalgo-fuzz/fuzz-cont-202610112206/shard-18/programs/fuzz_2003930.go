package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
func genFn0(p0 uint, p1 int16, p2 int32) string {
    p1 += 42
    v4 := GenS1{F0: 53, F1: 121}
    for fi5 := 0; fi5 < 8; fi5++ {
        if false {
            fmt.Println(p1, fi5)
            p2 += (p2 | (p2 | p2))
            switch fi5 {
            case 0:
                var v6 int8 = 42
                _ = v6
            case -8:
                switch ("hello" + ("foo bar" + "foo bar")) {
                case "\n\t":
                    v7 := GenS1{F0: 2147483648, F1: 1}
                    _ = v7
                }
            case -204:
                var v8 int16 = -688
                _ = v8
            }
        } else {
            p0 += p0
            switch fi5 {
            case 42:
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 int64 = int64(p2)
                    fmt.Println(fi5, v10)
                    v11 := GenS1{F0: 1, F1: 327}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case 128:
                v12 := GenS1{F0: 684, F1: 420}
                _ = v12
            case 32767:
                var v13 int64 = 504
                _ = v13
            default:
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 uint16 = 941
                    var v16 uint8 = 255
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            var v17 uint32 = 4294967295
            _ = v17
        }
        v18 := []int{fi5, (fi5 >> (((p0 * (p0 | p0))) & 63)), fi5}
        var v19 uint32 = 1
        _ = fi5
        _ = v18
        _ = v19
    }
    p0 += p0
    if (false || (65 > 2)) {
        p2 += p2
        v20 := GenS2{F0: 1, F1: p0}
        _ = v20
    } else {
        v21 := []int{439, -2147483648}
        p0 = p0
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ("hello" + "foo bar")
}
func genFn1() uint32 {
    for fi22 := 0; fi22 < 13; fi22++ {
        fmt.Println(fi22)
        _ = fi22
    }
    if false {
        if false {
            var v23 float64 = 0.0
            if false {
                v23 += v23
                v24 := GenS2{F0: 1, F1: 18446744073709551615}
                v23 = v23
                _ = v24
            }
            _ = v23
        }
        var v25 int = 15
        v26 := GenS3{F0: 9223372036854775807, F1: -319}
        _ = v25
        _ = v26
    } else {
        v27 := GenS4{F0: -1}
        _ = v27
    }
    var v28 bool = false
    _ = v28
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (4294967295 < 42)
    v3 := make(map[string]bool)
    v2 = (2 == 827)
    for fi29 := 0; fi29 < 19; fi29++ {
        var v30 int32 = 42
        v31 := GenS5{F0: (fi29 >> ((2) & 63))}
        var v32 float32 = 3.14159
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 int32 = 1
    v1[4294967295] = "\n\t"
    var v34 int32 = -2147483648
    var v35 int32 = v34
    var v36 bool = (v2 && (127 < 255))
    v34 += -769
    v37 := []int{}
    var v38 bool = false
    v39 := GenS3{F0: 744, F1: -206}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
