package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int16; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 uint8 }
func genFn0(p0 int32) int64 {
    if (!true) {
        v6 := GenS1{F0: 696, F1: 42, F2: p0}
        _ = v6
    } else {
        var v7 uint8 = 2
        switch 1 {
        case 128:
            if false {
                if (255 <= 0) {
                    var v8 uint = 18446744073709551615
                    _ = v8
                }
                var v9 string = ""
                _ = v9
            } else {
                var v10 string = "foo bar"
                for fi11 := 0; fi11 < 1; fi11++ {
                    if true {
                        var v12 float32 = 1e10
                        _ = v12
                    }
                    var v13 float32 = -2.25
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v15 := []int{}
                        fmt.Println(fi11)
                        _ = fi14
                        _ = v15
                    }
                    _ = fi11
                    _ = v13
                }
                _ = v10
            }
        case 9223372036854775807:
            if true {
                v16 := GenS1{F0: p0, F1: -895, F2: 42}
                _ = v16
            }
        case -2147483648:
            var v17 bool = ((("" + "a") + "\n\t") == "a")
            _ = v17
        default:
            v18 := []int{482, 1, 999}
            _ = v18
        }
        _ = v7
    }
    v19 := GenS1{F0: -1, F1: 927, F2: 1}
    v20 := []int{4294967295, -9223372036854775808, 32768}
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 uint32 = 757
        var v23 int64 = 0
        v23 = (v23 >> ((255) & 63))
        _ = fi21
        _ = v22
        _ = v23
    }
    p0 = (p0 << ((739) & 63))
    _ = p0
    _ = v19
    _ = v20
    return 1
}

func main() {
    v1 := []int{2147483647}
    v2 := []uint32{42, 603, 0, 2147483647}
    v3 := make(map[int]int)
    var v4 uint = 813
    v5 := make(map[uint32]uint32)
    for fi24 := 0; fi24 < 15; fi24++ {
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint = v4
            var v27 uint32 = 4294967295
            v28 := GenS1{F0: 1, F1: fi25, F2: 896}
            if true {
                v27 -= (v27 << ((v4) & 63))
                if true {
                    if false {
                        var v29 uint32 = v27
                        v26 = 2
                        v30 := GenS1{F0: 42, F1: (fi24 * (fi25 | fi24)), F2: 2147483647}
                        _ = v29
                        _ = v30
                    } else {
                        v31 := GenS1{F0: 42, F1: fi24, F2: 42}
                        _ = v31
                    }
                } else {
                    _ = len(v5)
                }
            } else {
                v32 := GenS2{F0: (!(!true)), F1: -128}
                v4 -= (v26 % ((v26) | 1))
                _ = v32
            }
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        if (42 <= 573) {
            var v33 int8 = -14
            var v34 int8 = v33
            _ = v33
            _ = v34
        } else {
            fmt.Println(fi24, fi24)
            if true {
                if true {
                    for fi35 := 0; fi35 < 15; fi35++ {
                        v4 = (^(v4))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v36 := GenS3{F0: (!((^(fi24)) != (fi24 << ((v4) & 63)))), F1: 42}
                        _ = fi35
                        _ = v36
                    }
                    v37 := v3[fi24]
                    v37 += v37
                    _ = v37
                } else {
                    fmt.Println(fi24)
                }
                switch (^(fi24)) {
                case -9223372036854775808:
                    v38 := GenS4{F0: -928, F1: 2, F2: 32767}
                    _ = v38
                default:
                    var v39 int32 = 42
                    _ = v39
                }
            } else {
                for fi40 := 0; fi40 < 10; fi40++ {
                    fmt.Println(v4, fi24, fi40)
                    _ = fi40
                }
            }
        }
        if true {
            v4 -= v4
            var v41 float64 = -2.25
            _ = len(v5)
            _ = v41
        } else {
            v2 = append(v2, 65535)
            _ = len(v2)
        }
        _ = fi24
    }
    v4 = (v4 << ((v4) & 63))
    fmt.Println(v4)
    fmt.Println(v4)
    var v42 string = "a"
    switch v42 {
    case "":
        if ("hello" > "") {
            var v43 uint16 = 0
            for fi44 := 0; fi44 < 5; fi44++ {
                _ = len(v2)
                v45 := GenS2{F0: true, F1: 0}
                switch -2147483648 {
                case -487:
                    switch "a" {
                    case "":
                        fmt.Println(v43, v42, v4)
                    }
                }
                _ = fi44
                _ = v45
            }
            switch -1 {
            case 32768:
                var v46 bool = (-2.25 > 3.14159)
                _ = v46
            case 1:
                fmt.Println(v4, v43, v4)
            }
            _ = v43
        } else {
            switch -679 {
            case 476:
                v47 := GenS1{F0: -2147483648, F1: 9223372036854775807, F2: -82}
                _ = v47
            case 32767:
                _ = len(v1)
            default:
                switch v42 {
                case "hello":
                    v48 := GenS1{F0: 0, F1: 127, F2: -1}
                    _ = v48
                case "\n\t":
                    v49 := GenS5{F0: 18446744073709551615}
                    _ = v49
                default:
                    var v50 uint64 = 4294967295
                    _ = v50
                }
            }
        }
    }
    switch v42 {
    case "hello":
        for fi51 := 0; fi51 < 9; fi51++ {
            v42 = v42
            switch (fi51 * (fi51 & fi51)) {
            case 2147483647:
                v52 := GenS3{F0: false, F1: -32768}
                _ = v52
            case -2147483648:
                for fi53 := 0; fi53 < 19; fi53++ {
                    fmt.Println(fi53, fi53)
                    v1 = append(v1, (fi51 >> (((v4 * (v4 / (((^(v4))) | 1)))) & 63)))
                    _ = fi53
                }
            case 127:
                for fi54 := 0; fi54 < 8; fi54++ {
                    if ((("" + (("" + "hello") + "foo bar")) + v42) == "foo bar") {
                        fmt.Println(fi54, fi54)
                        switch (fi51 * (fi54 * (fi51 >> ((v4) & 63)))) {
                        case 1:
                            fmt.Println(v42, fi51, fi51)
                        case 561:
                            v42 = v42
                        case -266:
                            var v55 float64 = 0.0
                            _ = v55
                        }
                        var v56 float64 = -2.25
                        _ = len(v2)
                        _ = v56
                    }
                    _ = fi54
                }
            default:
                for fi57 := 0; fi57 < 7; fi57++ {
                    v58 := GenS6{F0: v4}
                    switch (fi57 >> (((^(v4))) & 63)) {
                    case 762:
                        var v59 uint64 = 853
                        _ = v59
                    case 42:
                        switch (fi51 / ((fi57) | 1)) {
                        case 4294967295:
                            var v60 float32 = ((0.0 * (float32((3.14159 - 3.14159)) / 1.0)) - 0.1)
                            _ = len(v2)
                            fmt.Println(fi51)
                            _ = v60
                        default:
                            if (-128 < 42) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                    }
                    _ = fi57
                    _ = v58
                }
            }
            _ = fi51
        }
    default:
        var v61 uint = v4
        _ = v61
    }
    switch 270 {
    case 835:
        v4 -= 2
    default:
        fmt.Println(v4)
    }
    v62 := GenS7{F0: 100, F1: 4294967295}
    var v63 float32 = (3.14159 - ((((float32(0.0) * ((((0.1 - -2.25) / -2.25) - float32(((3.14159 - ((0.0 - -2.25) - 1e10)) / 0.1))) / (0.1 - (0.0 * 0.1)))) - 1.5) / 1.0) / -2.25))
    var v64 int64 = 1
    var v65 uint32 = 0
    switch v42 {
    case "hello":
        var v66 uint8 = 0
        _ = v66
    }
    if false {
        var v67 uint8 = 0
        v1 = append(v1, -451)
        for fi68 := 0; fi68 < 8; fi68++ {
            for fi69 := 0; fi69 < 17; fi69++ {
                var v70 int16 = 32767
                _ = fi69
                _ = v70
            }
            _ = fi68
        }
        _ = v67
    }
    fmt.Println(v42, v64, v65)
    v71 := GenS8{F0: 1}
    for fi72 := 0; fi72 < 11; fi72++ {
        var v73 uint64 = 4294967295
        if (int64(fi72) <= (^(v64))) {
            v73 -= (^(v73))
            fmt.Println(v42)
            v42 = "\n\t"
        } else {
            v74 := GenS3{F0: false, F1: 810}
            if false {
                v4 -= (v4 % ((uint(v65)) | 1))
            }
            _ = v74
        }
        for fi75 := 0; fi75 < 19; fi75++ {
            var v76 int16 = 67
            var v77 uint32 = 42
            var v78 uint8 = 166
            _ = fi75
            _ = v76
            _ = v77
            _ = v78
        }
        _ = fi72
        _ = v73
    }
    v42 = ("" + "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v71
}
