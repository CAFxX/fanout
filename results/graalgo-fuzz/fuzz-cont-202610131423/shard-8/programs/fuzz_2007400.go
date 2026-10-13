package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint; F1 string; F2 string }
func genFn0(p0 uint8) int32 {
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 int = fi4
        var v6 int8 = -124
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 418, F1: 1}
            _ = fi7
            _ = v8
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    v9 := GenS2{F0: 2, F1: 384, F2: 255}
    switch 127 {
    case 544:
        var v10 uint16 = 32767
        _ = v10
    case -2147483648:
        v11 := GenS2{F0: 1, F1: 18446744073709551615, F2: 32767}
        _ = v11
    case 32767:
        p0 += (p0 / (((p0 + p0)) | 1))
    default:
        if ((!(((float32(1.5) - 3.14159) / ((3.14159 / 1e10) - (-2.25 - (((-2.25 - 0.1) / 1.0) * (0.0 * (0.0 / (-2.25 - (0.1 + -2.25)))))))) != 0.0)) || false) {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int8 = -118
                if (!true) {
                    fmt.Println(v13, v13, fi12)
                }
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := []int{242}
                if (!(!(fi14 != (fi14 - 749)))) {
                    fmt.Println(fi14)
                }
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi14
                _ = v15
            }
            p0 = p0
        } else {
            if true {
                for fi16 := 0; fi16 < 8; fi16++ {
                    p0 = (^(p0))
                    v17 := []int{(fi16 / ((128) | 1)), int(p0), 904}
                    _ = fi16
                    _ = v17
                }
                p0 -= p0
            } else {
                v18 := []int{-2147483648}
                _ = len(v18)
                _ = v18
            }
        }
    }
    _ = p0
    _ = v9
    return -844
}
func genFn1() uint32 {
    var v19 int8 = 1
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 float64 = ((1.5 + 1e10) / 0.1)
        _ = fi20
        _ = v21
    }
    _ = v19
    return 1
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint8 = 211
    var v3 uint16 = 208
    for fi22 := 0; fi22 < 4; fi22++ {
        v3 = (v3 + v3)
        var v23 int64 = 1
        _ = fi22
        _ = v23
    }
    v24 := []int{-2147483648, 9223372036854775807, 4294967295, 128}
    v25 := GenS3{F0: 255, F1: "hello", F2: ("a" + ("" + "\n\t"))}
    switch -168 {
    case 2147483647:
        v2 -= 42
    }
    if false {
        var v26 int64 = 9223372036854775807
        if (!(2 <= 65535)) {
            if (v1 > (v1 % (((v1 % ((genFn0(127)) | 1))) | 1))) {
                if len(v24) > 0 {
                    v24[0] = 2147483647
                }
                var v27 int64 = (^(v26))
                var v28 int8 = 1
                _ = v27
                _ = v28
            }
            if (2 < v3) {
                v2 = (^(v2))
            }
        }
        _ = v26
    } else {
        v3 -= (v3 >> ((255) & 63))
        v29 := GenS3{F0: 255, F1: ("hello" + "hello"), F2: ("" + ("hello" + "foo bar"))}
        _ = v29
    }
    var v30 bool = true
    v24 = append(v24, 726)
    var v31 int32 = v1
    var v32 uint64 = 1
    if (!false) {
        fmt.Println(v2, v1, v30)
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
    }
    var v33 string = "foo bar"
    switch v33 {
    case "hello":
        v34 := GenS2{F0: 637, F1: (^(v32)), F2: v3}
        _ = v34
    case "\n\t":
        var v35 uint32 = 208
        _ = v35
    default:
        v3 = 930
    }
    fmt.Println(v32, v2)
    var v36 bool = v30
    fmt.Println(v31, v36, v33)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v36
}
