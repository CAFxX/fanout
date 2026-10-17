package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int32; F1 uint; F2 uint }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 uint64; F1 int8; F2 int }
type GenS10 struct { F0 uint16; F1 bool }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]uint32)
    var v3 uint8 = 255
    var v4 int16 = v1
    v5 := make(map[int]string)
    var v6 bool = (!(-1 >= 1))
    var v7 uint64 = 471
    if v6 {
        if (42 != -542) {
            if false {
                switch 32768 {
                case -668:
                    var v8 int = 32768
                    _ = v8
                default:
                    v9 := v5[9223372036854775807]
                    _ = v9
                }
                for fi10 := 0; fi10 < 5; fi10++ {
                    v1 -= v4
                    _ = fi10
                }
            } else {
                fmt.Println(v4, v4)
                switch -877 {
                case -1:
                    v11 := GenS1{F0: true, F1: 255, F2: 198}
                    _ = v11
                }
            }
            if (2 < 561) {
                v12 := GenS1{F0: false, F1: 2, F2: 2}
                var v13 int64 = -9223372036854775808
                _ = v12
                _ = v13
            }
        } else {
            var v14 uint16 = 209
            v15 := GenS2{F0: 0, F1: v4}
            v16 := GenS3{F0: "foo bar", F1: v1}
            _ = v14
            _ = v15
            _ = v16
        }
    } else {
        if v6 {
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 int8 = 1
                _ = fi17
                _ = v18
            }
            var v19 int16 = v1
            v20 := []int{127, -129}
            _ = v19
            _ = v20
        } else {
            if v6 {
                v21 := GenS4{F0: 42, F1: 18446744073709551615, F2: 1}
                _ = v21
            } else {
                var v22 int64 = 0
                var v23 uint32 = 558
                _ = v22
                _ = v23
            }
            fmt.Println(v4)
        }
        var v24 int64 = -1
        _ = v24
    }
    v25 := GenS2{F0: -128, F1: (v4 - (-(v1)))}
    switch 585 {
    case 56:
        for fi26 := 0; fi26 < 9; fi26++ {
            v1 += v4
            var v27 uint16 = 2
            var v28 uint64 = uint64(v3)
            _ = fi26
            _ = v27
            _ = v28
        }
    case 872:
        v29 := GenS1{F0: (1 != 42), F1: 678, F2: 104}
        _ = v29
    case 9223372036854775807:
        v30 := GenS5{F0: 618, F1: 1}
        _ = v30
    }
    var v31 int64 = 618
    fmt.Println(v4, v7, v1)
    v32 := []int{1, 242, 32767, -225}
    switch -501 {
    case 842:
        v31 = (v31 * (v31 * (v31 & v31)))
    default:
        var v33 float64 = -2.25
        _ = v33
    }
    v1 += (v4 | int16(v7))
    v7 = v7
    var v34 int64 = 1
    _ = len(v32)
    v6 = (1.5 >= 1e10)
    v7 = (v7 - (v7 << ((150) & 63)))
    if false {
        if true {
            v1 += -32768
            fmt.Println(v7, v31)
        } else {
            v35 := GenS6{F0: (v1 - 0)}
            switch 9223372036854775807 {
            case -9223372036854775808:
                var v36 float32 = 1e10
                _ = v36
            case 127:
                var v37 int64 = v31
                _ = v37
            default:
                for fi38 := 0; fi38 < 16; fi38++ {
                    for fi39 := 0; fi39 < 20; fi39++ {
                        switch ("" + "") {
                        case "foo bar":
                            v40 := GenS5{F0: (^(v34)), F1: (v34 << ((1) & 63))}
                            _ = v40
                        case "hello":
                            if true {
                                if len(v32) > 0 {
                                    v32[0] = fi39
                                }
                                if (0 <= -116) {
                                    if v6 {
                                        if len(v32) > 0 {
                                            v32[0] = fi38
                                        }
                                        var v41 uint16 = 42
                                        fmt.Println(v34, v34, v6)
                                        var v42 uint16 = v41
                                        _ = v41
                                        _ = v42
                                    } else {
                                        var v43 int = fi39
                                        _ = v43
                                    }
                                    v44 := GenS7{F0: 4294967295}
                                    var v45 string = "\n\t"
                                    v46 := GenS3{F0: v45, F1: v1}
                                    switch "foo bar" {
                                    case "":
                                        switch v45 {
                                        case "a":
                                            v4 += (v1 + 1000)
                                        }
                                    }
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    switch fi39 {
                                    case 2147483647:
                                        var v47 bool = (345 <= 42)
                                        _ = v47
                                    case -2147483648:
                                        v48 := GenS4{F0: -2147483648, F1: 65535, F2: 141}
                                        _ = v48
                                    default:
                                        v6 = false
                                        if len(v32) > 0 {
                                            v32[0] = -392
                                        }
                                    }
                                    if len(v32) >= 2 {
                                        v32 = v32[1:len(v32)]
                                    }
                                    _ = len(v32)
                                    if (4294967295 >= 0) {
                                        switch "foo bar" {
                                        case "hello":
                                            var v49 uint16 = 42
                                            _ = v49
                                        case "\n\t":
                                            v1 -= (v4 * v4)
                                        case "":
                                            _ = len(v32)
                                        default:
                                            if v6 {
                                                if len(v32) >= 2 {
                                                    v32 = v32[1:len(v32)]
                                                }
                                                _ = len(v5)
                                                switch (-(fi39)) {
                                                case -9223372036854775808:
                                                    var v50 string = "a"
                                                    _ = v50
                                                }
                                            } else {
                                                if ((1 <= 420) && true) {
                                                    fmt.Println(v7)
                                                    if v6 {
                                                        var v51 float32 = 1.5
                                                        _ = v51
                                                    } else {
                                                        v32 = append(v32, fi38)
                                                        var v52 uint16 = 32767
                                                        _ = v52
                                                    }
                                                    switch (fi38 * (fi38 | (-(fi38)))) {
                                                    case 0:
                                                        fmt.Println(fi38, v4, v6)
                                                    case 170:
                                                        if v6 {
                                                            v1 -= v4
                                                            v3 = v3
                                                            v4 -= (^(v4))
                                                        }
                                                    }
                                                }
                                                _ = len(v5)
                                                fmt.Println(v3)
                                                if len(v32) > 0 {
                                                    v32[0] = (fi38 / ((fi39) | 1))
                                                }
                                                v53 := GenS8{F0: "\n\t", F1: 1}
                                                var v54 uint64 = v7
                                                v55 := GenS1{F0: v6, F1: 286, F2: 274}
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                        }
                                    } else {
                                        fmt.Println(v34, fi38)
                                    }
                                    if len(v32) >= 2 {
                                        v32 = v32[1:len(v32)]
                                    }
                                }
                                v34 += v31
                                v56 := GenS8{F0: "", F1: 42}
                                if v6 {
                                    var v57 int64 = (v31 << ((0) & 63))
                                    v1 += v4
                                    var v58 uint64 = (v7 / (((^(v7))) | 1))
                                    _ = v57
                                    _ = v58
                                } else {
                                    v59 := GenS9{F0: (^(v7)), F1: 42, F2: (^(fi39))}
                                    var v60 uint64 = v7
                                    _ = len(v5)
                                    fmt.Println(fi38, v31, v34)
                                    var v61 uint8 = v3
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                _ = v56
                            }
                        default:
                            if (false && v6) {
                                v31 -= (v34 & v34)
                            } else {
                                v3 += v3
                            }
                        }
                        switch (fi38 ^ fi38) {
                        case -9223372036854775808:
                            v7 -= 42
                        case -2147483648:
                            v31 -= (v34 % (((v34 | 42)) | 1))
                        default:
                            var v62 string = "a"
                            v63 := GenS10{F0: 931, F1: v6}
                            switch (fi39 | (fi38 / (((fi39 >> ((0) & 63))) | 1))) {
                            case 128:
                                if false {
                                    var v64 int = (-(fi39))
                                    fmt.Println(v64, fi38)
                                    v62 = v62
                                    if v6 {
                                        _ = len(v32)
                                    }
                                    var v65 float64 = (0.0 - 3.14159)
                                    v32 = append(v32, int(v3))
                                    var v66 int64 = -567
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                            default:
                                fmt.Println(v7, v34, v6)
                            }
                            _ = v62
                            _ = v63
                        }
                        _ = fi39
                    }
                    v31 -= v31
                    _ = fi38
                }
            }
            _ = v35
        }
        for fi67 := 0; fi67 < 15; fi67++ {
            v7 -= 42
            _ = fi67
        }
        v3 = (v3 ^ 255)
    }
    fmt.Println(v34, v7)
    v6 = (v6 || (2 > v7))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
