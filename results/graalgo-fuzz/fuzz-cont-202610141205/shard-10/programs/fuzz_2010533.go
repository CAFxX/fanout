package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int64; F1 int; F2 bool }
func genFn0(p0 uint16, p1 int8) uint32 {
    switch "" {
    case "a":
        v4 := []int{}
        _ = v4
    case "\n\t":
        v5 := []int{}
        _ = v5
    case "foo bar":
        p0 += (p0 ^ 645)
    }
    switch 882 {
    case 473:
        if true {
            var v6 string = ""
            _ = v6
        } else {
            fmt.Println(p0)
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := GenS1{F0: (fi7 + 127), F1: p1, F2: 100}
                for fi9 := 0; fi9 < 3; fi9++ {
                    p1 = (p1 - p1)
                    p1 = (p1 % (((p1 / ((46) | 1))) | 1))
                    var v10 int = 128
                    switch fi9 {
                    case 32767:
                        var v11 int8 = 0
                        _ = v11
                    case 1:
                        var v12 int = 128
                        _ = v12
                    default:
                        if ((0.0 * float64(3.14159)) <= 3.14159) {
                            if (654 <= 795) {
                                p0 = (p0 + (p0 & (^(p0))))
                                var v13 uint64 = 630
                                _ = v13
                            }
                        } else {
                            v14 := GenS1{F0: 4294967295, F1: p1, F2: 100}
                            var v15 float32 = 0.0
                            var v16 string = ""
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    }
                    var v17 int8 = -86
                    _ = fi9
                    _ = v10
                    _ = v17
                }
                _ = fi7
                _ = v8
            }
        }
    default:
        p1 = (p1 << ((18446744073709551615) & 63))
    }
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{341, 2}
    var v3 string = ("" + "foo bar")
    if (18446744073709551615 == 4294967295) {
        var v18 int8 = -128
        _ = v18
    } else {
        for fi19 := 0; fi19 < 2; fi19++ {
            if false {
                v3 = v3
                v3 = v3
            } else {
                var v20 int8 = 42
                if len(v2) > 0 {
                    v2[0] = 2
                }
                v21 := GenS1{F0: (fi19 & (^(fi19))), F1: v20, F2: 116}
                _ = v20
                _ = v21
            }
            var v22 string = v3
            _ = fi19
            _ = v22
        }
        v3 = ("" + ((("a" + v3) + v3) + v3))
    }
    v3 = "a"
    fmt.Println(v3)
    v23 := GenS1{F0: 2147483647, F1: 0, F2: 127}
    v2 = append(v2, 42)
    var v24 uint16 = 1
    if (false && (42 >= 1)) {
        v3 = v3
    } else {
        var v25 int8 = -73
        v25 = 0
        _ = v25
    }
    switch -762 {
    case -9223372036854775808:
        switch -2147483648 {
        case 576:
            var v26 float64 = 0.1
            _ = v26
        case 0:
            _ = len(v2)
        case -103:
            var v27 int16 = -1
            _ = v27
        default:
            v24 = 1
        }
    case 32768:
        if ((1e10 == 1.5) && true) {
            switch v3 {
            case "hello":
                v28 := GenS1{F0: 2147483647, F1: 100, F2: 127}
                _ = v28
            }
            _ = len(v1)
        } else {
            switch 128 {
            case 32768:
                fmt.Println(v3, v3)
            case -211:
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 uint8 = 26
                    v3 = "\n\t"
                    _ = fi29
                    _ = v30
                }
            case -9223372036854775808:
                fmt.Println(v3, v3)
            default:
                if (0 < 215) {
                    fmt.Println(v3, v3)
                }
            }
            if (38 >= 2) {
                v3 = v3
                for fi31 := 0; fi31 < 14; fi31++ {
                    _ = len(v2)
                    _ = fi31
                }
                v32 := GenS2{F0: genFn0(v24, 100)}
                _ = v32
            }
        }
    case 0:
        switch 910 {
        case 32767:
            switch "a" {
            case "":
                v33 := GenS3{F0: -1, F1: v24, F2: 4294967295}
                _ = v33
            case "hello":
                v34 := v1[("a" + (v3 + (v3 + v3)))]
                _ = v34
            case "\n\t":
                v24 -= 32767
            }
        case -1:
            v3 = v3
        }
    default:
        fmt.Println(v24)
    }
    v35 := GenS3{F0: 42, F1: 391, F2: 2}
    v36 := GenS4{F0: -9223372036854775808, F1: 0, F2: false}
    v37 := GenS3{F0: 98, F1: v24, F2: 1}
    var v38 int = 860
    var v39 int8 = -10
    v1[(((("foo bar" + v3) + v3) + (v3 + v3)) + v3)] = (v3 + v3)
    v3 = v3
    v3 = v3
    v40 := v1[v3]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
