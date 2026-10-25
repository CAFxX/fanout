package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint16; F1 string; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint16, p1 string, p2 uint64) uint {
    v6 := []int{897, 128, 1, 14}
    if true {
        var v7 uint32 = 4294967295
        p2 -= p2
        _ = v7
    } else {
        fmt.Println(p0, p2, p1)
        v6 = append(v6, 128)
    }
    v6 = append(v6, 127)
    if (!((83 >= -99) || (false || false))) {
        if false {
            if (true || true) {
                var v8 int32 = -2147483648
                _ = len(v6)
                if (((729 < 0) && false) || (false || false)) {
                    var v9 uint = uint(p2)
                    _ = v9
                }
                _ = v8
            } else {
                for fi10 := 0; fi10 < 9; fi10++ {
                    v11 := GenS1{F0: p2}
                    _ = fi10
                    _ = v11
                }
                var v12 uint16 = (p0 % ((p0) | 1))
                _ = v12
            }
            if ((1e10 / -2.25) >= (1e10 - ((0.0 * 0.0) + 1e10))) {
                if false {
                    v13 := GenS1{F0: (p2 % (((p2 & p2)) | 1))}
                    fmt.Println(p0, p1, p0)
                    _ = v13
                }
                var v14 int32 = 0
                _ = v14
            } else {
                var v15 int16 = -699
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v15
            }
        } else {
            switch 506 {
            case 32767:
                if len(v6) > 0 {
                    v6[0] = -939
                }
            case 9223372036854775807:
                p2 = 378
            case -2147483648:
                if len(v6) > 0 {
                    v6[0] = 9223372036854775807
                }
            }
            _ = len(v6)
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 int32 = 343
                var v18 int8 = 42
                v19 := GenS2{F0: 204, F1: 127, F2: 1}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        }
        v6 = append(v6, 192)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}
func genFn1(p0 int32) int64 {
    p0 = p0
    v20 := []int{128, -224, 850, 218}
    if true {
        switch ("foo bar" + ("a" + "\n\t")) {
        case "foo bar":
            for fi21 := 0; fi21 < 19; fi21++ {
                fmt.Println(fi21)
                _ = fi21
            }
        case "\n\t":
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS1{F0: 4294967295}
                _ = fi22
                _ = v23
            }
        case "a":
            v24 := GenS1{F0: 18446744073709551615}
            _ = v24
        default:
            if true {
                if ((3.14159 / float32(1e10)) == 1.5) {
                    if len(v20) > 0 {
                        v20[0] = -85
                    }
                    v25 := GenS1{F0: 1}
                    _ = v25
                }
            }
        }
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        var v26 int8 = 1
        _ = v26
    } else {
        var v27 uint16 = 255
        _ = v27
    }
    var v28 int = 0
    _ = p0
    _ = v20
    _ = v28
    return 833
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{2147483647, 458}
    var v3 float32 = (-2.25 - float32((1e10 * (0.0 * 0.1))))
    var v4 float64 = 3.14159
    v5 := make(map[string]uint32)
    v29 := GenS2{F0: 611, F1: 169, F2: 9223372036854775807}
    v30 := GenS3{F0: 0, F1: "hello", F2: "hello"}
    v5[("" + "")] = 85
    var v31 int64 = 42
    v32 := GenS4{F0: "a"}
    for fi33 := 0; fi33 < 4; fi33++ {
        v34 := GenS5{F0: 1}
        _ = fi33
        _ = v34
    }
    v2 = append(v2, -9223372036854775808)
    var v35 float32 = v3
    switch "a" {
    case "":
        fmt.Println(v35, v35)
    case "a":
        v35 += v35
    case "foo bar":
        if (-940 != (v31 & (v31 / (((^(v31))) | 1)))) {
            if false {
                if false {
                    var v36 uint32 = 776
                    v37 := GenS4{F0: "\n\t"}
                    switch ("\n\t" + "a") {
                    case "hello":
                        var v38 uint64 = uint64(v36)
                        _ = v38
                    default:
                        var v39 uint8 = 0
                        _ = v39
                    }
                    _ = v36
                    _ = v37
                }
            }
            v40 := GenS5{F0: 0}
            fmt.Println(v3)
            _ = v40
        }
    }
    if len(v2) > 0 {
        v2[0] = 295
    }
    var v41 uint8 = 2
    v31 = v31
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v41
    _ = v5
}
