package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int16; F2 int8 }
func genFn0() uint {
    v4 := GenS1{F0: false, F1: 946, F2: -673}
    var v5 bool = (187 == 925)
    switch -1 {
    case -9223372036854775808:
        v6 := GenS1{F0: (v5 || false), F1: 1, F2: 116}
        _ = v6
    case 42:
        switch "" {
        case "hello":
            v7 := GenS2{F0: 65535, F1: 42}
            _ = v7
        }
    default:
        v5 = false
    }
    v8 := GenS1{F0: v5, F1: 42, F2: 246}
    _ = v4
    _ = v5
    _ = v8
    return 536
}
func genFn1(p0 uint8) uint32 {
    v9 := []int{}
    _ = len(v9)
    _ = p0
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint8 = 65
    var v2 uint32 = uint32(v1)
    var v3 int16 = 975
    if (!(uint32(v1) >= genFn1(255))) {
        switch -1 {
        case 1:
            switch "a" {
            case "\n\t":
                var v10 float32 = -2.25
                _ = v10
            default:
                for fi11 := 0; fi11 < 4; fi11++ {
                    v12 := GenS2{F0: 42, F1: 127}
                    _ = fi11
                    _ = v12
                }
            }
        }
        v13 := GenS3{F0: (v2 ^ (v2 / ((genFn1((v1 / ((1) | 1)))) | 1))), F1: "hello"}
        v2 = v2
        _ = v13
    }
    var v14 int16 = -86
    if (!false) {
        if false {
            switch -437 {
            case 0:
                v1 += 0
            case 32768:
                v2 += v2
            case 128:
                v2 += v2
            }
            fmt.Println(v14, v2, v1)
        }
        switch "hello" {
        case "a":
            v15 := []int{1, int(v1), 4294967295}
            _ = v15
        case "hello":
            v16 := []int{}
            _ = v16
        default:
            if true {
                var v17 int32 = 1
                v18 := GenS1{F0: (((1.5 * (0.0 - 0.0)) != 3.14159) && ((v14 == 0) || false)), F1: 42, F2: 42}
                v3 -= (v3 >> ((630) & 63))
                _ = v17
                _ = v18
            }
        }
    } else {
        var v19 int64 = int64(v2)
        _ = v19
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v3, v2)
    case "\n\t":
        var v20 uint64 = 405
        _ = v20
    case "foo bar":
        v21 := []int{1, 1, 952, 2147483647}
        _ = v21
    default:
        if true {
            if true {
                for fi22 := 0; fi22 < 3; fi22++ {
                    fmt.Println(fi22, v3, v2)
                    v23 := GenS4{F0: 724}
                    _ = fi22
                    _ = v23
                }
            } else {
                var v24 int64 = 42
                switch 614 {
                case 128:
                    for fi25 := 0; fi25 < 4; fi25++ {
                        switch fi25 {
                        case 416:
                            fmt.Println(v24)
                        }
                        _ = fi25
                    }
                }
                var v26 int64 = (v24 / ((v24) | 1))
                _ = v24
                _ = v26
            }
            if (v3 <= 42) {
                v27 := []int{0, 128, 9223372036854775807}
                var v28 uint = genFn0()
                _ = v27
                _ = v28
            }
        } else {
            v29 := []int{}
            v30 := GenS3{F0: (^(v2)), F1: ("foo bar" + (("\n\t" + "") + "a"))}
            _ = v29
            _ = v30
        }
    }
    if true {
        v3 = v3
    }
    fmt.Println(v2)
    if (false || true) {
        fmt.Println(v14, v1)
        var v31 int = 42
        _ = v31
    } else {
        var v32 float64 = float64(0.1)
        v33 := []int{32767}
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        fmt.Println(fi34)
        v35 := []int{}
        _ = fi34
        _ = v35
    }
    v36 := []int{622, 42, int(v1)}
    var v37 int32 = -1
    v38 := GenS5{F0: "hello", F1: (^(v3)), F2: 1}
    _ = len(v36)
    v14 = v14
    if (824 != v37) {
        if ((42 != -128) || false) {
            v1 = (^(v1))
        }
        for fi39 := 0; fi39 < 4; fi39++ {
            switch -9223372036854775808 {
            case 0:
                if ((v37 >= v37) && true) {
                    var v40 bool = false
                    _ = v40
                } else {
                    fmt.Println(v3, v3, v37)
                    fmt.Println(v14)
                }
            default:
                _ = len(v36)
            }
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
            _ = fi39
        }
        fmt.Println(v2, v3)
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(v2)
    case "a":
        fmt.Println(v2, v37, v1)
    default:
        var v41 int16 = (^(v14))
        _ = v41
    }
    if true {
        switch 432 {
        case 468:
            if len(v36) > 0 {
                v36[0] = 32768
            }
        case 2147483647:
            if false {
                fmt.Println(v14, v1, v37)
            } else {
                var v42 uint64 = 0
                _ = v42
            }
        }
        var v43 float64 = 0.0
        _ = v43
    } else {
        if false {
            _ = len(v36)
        }
        for fi44 := 0; fi44 < 15; fi44++ {
            if (126 >= 1) {
                for fi45 := 0; fi45 < 20; fi45++ {
                    var v46 int = (fi44 | (fi44 - (fi45 + fi45)))
                    _ = fi45
                    _ = v46
                }
            } else {
                v47 := GenS5{F0: "hello", F1: (v14 + v3), F2: 40}
                _ = v47
            }
            _ = fi44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
