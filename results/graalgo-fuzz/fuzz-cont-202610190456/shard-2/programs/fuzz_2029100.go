package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 int16; F1 uint; F2 int8 }
func genFn0(p0 int, p1 bool, p2 uint) uint32 {
    p0 -= (p0 >> (((p2 >> (((p2 << (((p2 ^ (^(p2)))) & 63))) & 63))) & 63))
    switch "hello" {
    case "hello":
        if true {
            v4 := []int{p0}
            if (((p0 ^ p0) > p0) && (p1 || (2 <= 1))) {
                switch "" {
                case "\n\t":
                    fmt.Println(p1, p2, p0)
                default:
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(p0, p0, p0)
                    default:
                        var v5 uint64 = 1
                        _ = v5
                    }
                }
                fmt.Println(p2, p0, p1)
                p0 += p0
            } else {
                p0 += -1
            }
            var v6 int16 = 0
            _ = v4
            _ = v6
        } else {
            var v7 int32 = -2147483648
            switch "a" {
            case "\n\t":
                fmt.Println(p1, p1, v7)
            case "":
                for fi8 := 0; fi8 < 18; fi8++ {
                    p2 = p2
                    _ = fi8
                }
            case "hello":
                v7 += (v7 - v7)
            default:
                fmt.Println(v7, p1, p2)
            }
            var v9 int32 = -2147483648
            _ = v7
            _ = v9
        }
    case "foo bar":
        v10 := []int{-9223372036854775808, (p0 & p0), 4294967295}
        _ = v10
    default:
        v11 := []int{893, 0}
        _ = v11
    }
    v12 := []int{p0, 2147483647, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 748
}
func genFn1(p0 bool, p1 uint16) uint16 {
    switch 2147483647 {
    case 32768:
        v13 := GenS1{F0: 628, F1: (p1 - p1)}
        _ = v13
    case -9223372036854775808:
        if p0 {
            v14 := []int{4294967295}
            for fi15 := 0; fi15 < 9; fi15++ {
                for fi16 := 0; fi16 < 5; fi16++ {
                    v14 = append(v14, fi15)
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    _ = len(v14)
                    if p0 {
                        if len(v14) > 0 {
                            v14[0] = (fi15 | -333)
                        }
                        switch "foo bar" {
                        case "foo bar":
                            p1 -= 1
                            if p0 {
                                var v17 uint32 = genFn0(fi15, (0 >= 0), 65535)
                                _ = v17
                            } else {
                                if true {
                                    p1 = (p1 * (p1 % ((p1) | 1)))
                                }
                                var v18 int8 = 42
                                v19 := GenS1{F0: 1, F1: 2}
                                _ = v18
                                _ = v19
                            }
                        case "a":
                            if (2147483648 < 4294967295) {
                                v20 := GenS2{F0: 1, F1: "a"}
                                var v21 int32 = -806
                                _ = v20
                                _ = v21
                            }
                        case "hello":
                            var v22 uint64 = 757
                            _ = v22
                        default:
                            var v23 float64 = 3.14159
                            _ = v23
                        }
                        p1 = p1
                    } else {
                        switch fi15 {
                        case 0:
                            var v24 int16 = 1
                            _ = v24
                        case 128:
                            var v25 int = (fi15 / ((fi15) | 1))
                            _ = v25
                        case 32767:
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        default:
                            if p0 {
                                if (1 <= 42) {
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                    if ("a" <= "\n\t") {
                                        fmt.Println(p1, p1, p1)
                                    } else {
                                        if true {
                                            var v26 uint16 = (p1 ^ p1)
                                            var v27 int64 = -9223372036854775808
                                            _ = v26
                                            _ = v27
                                        } else {
                                            var v28 bool = p0
                                            p0 = v28
                                            _ = v28
                                        }
                                        p0 = true
                                        v29 := GenS2{F0: p1, F1: "\n\t"}
                                        _ = v29
                                    }
                                }
                                var v30 float32 = 0.0
                                _ = v30
                            }
                        }
                    }
                    if (!(p0 || (p1 != (p1 >> ((4294967295) & 63))))) {
                        if len(v14) > 0 {
                            v14[0] = 492
                        }
                        v31 := GenS3{F0: p1}
                        _ = v31
                    } else {
                        if len(v14) > 0 {
                            v14[0] = (fi16 << ((536) & 63))
                        }
                    }
                    _ = fi16
                }
                v32 := GenS3{F0: (p1 - 32767)}
                _ = fi15
                _ = v32
            }
            var v33 uint16 = 720
            _ = v14
            _ = v33
        }
    case 731:
        v34 := GenS4{F0: -1, F1: 42, F2: 2147483648}
        _ = v34
    }
    v35 := []int{42, 127, -746}
    p1 = p1
    _ = p0
    _ = p1
    _ = v35
    return p1
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v36 uint8 = 53
    v36 = (v36 | v36)
    v37 := GenS4{F0: -394, F1: 65535, F2: 2}
    fmt.Println(v1, v1)
    for fi38 := 0; fi38 < 11; fi38++ {
        switch fi38 {
        case 128:
            v39 := []int{128, 1}
            _ = v39
        }
        var v40 int8 = 0
        _ = fi38
        _ = v40
    }
    if true {
        var v41 int32 = -1
        fmt.Println(v36, v41, v41)
        v1 = float64((1.5 + 3.14159))
        _ = v41
    } else {
        v36 = v36
        v42 := GenS5{F0: true, F1: 547, F2: 32767}
        switch "foo bar" {
        case "hello":
            fmt.Println(v36, v36)
        default:
            fmt.Println(v36, v36)
        }
        _ = v42
    }
    _ = len(v2)
    if (-1 != 965) {
        fmt.Println(v36, v1)
        for fi43 := 0; fi43 < 2; fi43++ {
            var v44 float32 = 3.14159
            v45 := GenS6{F0: -128, F1: 2147483647}
            var v46 int8 = -128
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
    }
    var v47 uint32 = 2147483648
    for fi48 := 0; fi48 < 15; fi48++ {
        v49 := []int{(^(fi48)), fi48}
        if (372 > 1) {
            v49 = append(v49, 127)
            if (("hello" + "\n\t") == "foo bar") {
                v1 -= -2.25
            } else {
                if false {
                    v50 := GenS2{F0: uint16(v47), F1: ("\n\t" + ("foo bar" + ("a" + ("hello" + ""))))}
                    if ((v47 % ((2147483647) | 1)) >= v47) {
                        v51 := GenS7{F0: 32767, F1: 4294967295, F2: -1}
                        fmt.Println(v47, v1)
                        _ = v51
                    } else {
                        var v52 uint = 729
                        if (!(-262 <= 2147483647)) {
                            v53 := GenS7{F0: 0, F1: v52, F2: 127}
                            _ = v53
                        }
                        var v54 uint64 = 4294967295
                        _ = v52
                        _ = v54
                    }
                    v36 -= (^(v36))
                    _ = v50
                } else {
                    fmt.Println(v1, fi48)
                }
                v55 := GenS7{F0: 1000, F1: 750, F2: 1}
                _ = v55
            }
        } else {
            delete(v2, "")
        }
        var v56 int8 = 1
        _ = fi48
        _ = v49
        _ = v56
    }
    v57 := []int{-1}
    for fi58 := 0; fi58 < 2; fi58++ {
        var v59 uint16 = 42
        _ = fi58
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v57:", len(v57), v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v47
    _ = v57
}
