package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int16; F2 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 uint64; F2 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 int32; F1 int; F2 string }
type GenS11 struct { F0 int64; F1 bool; F2 uint }
type GenS12 struct { F0 uint8; F1 int16; F2 uint8 }
func genFn0() int16 {
    if (!true) {
        var v7 uint = 0
        v8 := GenS1{F0: -1, F1: 220}
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint = v7
            fmt.Println(v7, v7, fi9)
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS2{F0: fi11}
                switch "" {
                case "foo bar":
                    var v13 float32 = 0.0
                    _ = v13
                case "\n\t":
                    v14 := []int{fi9}
                    _ = v14
                case "hello":
                    v15 := []int{fi11}
                    _ = v15
                }
                if false {
                    if false {
                        switch (fi11 >> (((v7 - (v7 + (v10 * (v7 << (((v7 / (((v10 * v10)) | 1))) & 63)))))) & 63)) {
                        case -568:
                            fmt.Println(fi11, v10)
                            if ((!false) || (((1e10 - (1.5 / float32(((3.14159 + 0.1) * -2.25)))) / 3.14159) >= 1e10)) {
                                fmt.Println(v10, fi9, fi9)
                            } else {
                                v16 := GenS3{F0: ("\n\t" + "a")}
                                var v17 string = ((("\n\t" + "hello") + "foo bar") + "\n\t")
                                _ = v16
                                _ = v17
                            }
                        case 132:
                            var v18 int32 = 664
                            _ = v18
                        case 27:
                            switch (("hello" + "foo bar") + "") {
                            case "foo bar":
                                var v19 uint8 = 2
                                _ = v19
                            }
                        }
                        fmt.Println(fi11, v10)
                    } else {
                        v20 := GenS3{F0: "foo bar"}
                        switch (fi11 * (fi9 + 127)) {
                        case -1:
                            var v21 string = "\n\t"
                            _ = v21
                        case 32767:
                            v22 := GenS1{F0: (fi9 - 4294967295), F1: 18446744073709551615}
                            _ = v22
                        case -2147483648:
                            v23 := GenS3{F0: ""}
                            _ = v23
                        default:
                            var v24 string = ("hello" + "foo bar")
                            _ = v24
                        }
                        v25 := GenS4{F0: 897, F1: 1, F2: 1}
                        fmt.Println(v10, fi11, fi9)
                        v10 += 847
                        if true {
                            var v26 int32 = 1
                            v26 = -1
                            v27 := GenS1{F0: (fi9 * -2147483648), F1: 260}
                            if false {
                                v26 = (-(v26))
                            }
                            switch "hello" {
                            case "hello":
                                v10 += v7
                            case "foo bar":
                                v26 = (^(v26))
                            }
                            fmt.Println(v7, v26)
                            v26 -= v26
                            _ = v26
                            _ = v27
                        } else {
                            if false {
                                switch "a" {
                                case "\n\t":
                                    v28 := GenS4{F0: -622, F1: 781, F2: 480}
                                    _ = v28
                                case "foo bar":
                                    v29 := []int{}
                                    _ = v29
                                }
                                switch ("a" + "a") {
                                case "foo bar":
                                    fmt.Println(v7, v10, fi9)
                                default:
                                    var v30 int16 = -1
                                    _ = v30
                                }
                            } else {
                                v7 -= v7
                            }
                        }
                        _ = v20
                        _ = v25
                    }
                    var v31 float32 = (float32(1.5) + (0.1 - 3.14159))
                    if false {
                        v32 := GenS5{F0: 100}
                        var v33 uint8 = uint8(v7)
                        var v34 int16 = 945
                        var v35 uint = v10
                        v35 -= (v10 / (((v10 | v10)) | 1))
                        _ = v32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    _ = v31
                }
                v36 := GenS2{F0: 0}
                fmt.Println(fi11)
                v37 := GenS6{F0: 1, F1: 42, F2: fi11}
                if (((v10 ^ (v7 << (((^(v10))) & 63))) < v10) || (42 == 98)) {
                    v38 := GenS7{F0: 42}
                    switch "a" {
                    case "":
                        if (((1 >= 9223372036854775807) && (1.5 == 1.5)) || false) {
                            var v39 int32 = 846
                            switch "" {
                            case "foo bar":
                                v40 := []int{(^(fi11)), 128}
                                v39 -= -510
                                _ = v40
                            }
                            v7 = (v7 >> ((v10) & 63))
                            _ = v39
                        }
                    }
                    var v41 string = "foo bar"
                    _ = v38
                    _ = v41
                }
                if true {
                    v42 := GenS8{F0: (fi11 & (fi9 | fi11))}
                    _ = v42
                } else {
                    switch fi9 {
                    case 898:
                        v43 := []int{(fi9 - fi9), fi11, fi11}
                        _ = v43
                    case 128:
                        v44 := []int{fi11, fi9, (fi9 - -917)}
                        if len(v44) > 0 {
                            v44[0] = fi11
                        }
                        _ = len(v44)
                        var v45 uint16 = 0
                        _ = v44
                        _ = v45
                    case 32768:
                        var v46 float32 = 0.0
                        _ = v46
                    }
                    var v47 uint16 = uint16(v10)
                    if (185 <= 42) {
                        if (!(100 >= 127)) {
                            var v48 bool = (((!false) && false) || false)
                            fmt.Println(v10, v47)
                            _ = v48
                        }
                        var v49 int32 = -788
                        v7 += (v10 & 65535)
                        _ = v49
                    }
                    _ = v47
                }
                v10 = (^(v7))
                _ = fi11
                _ = v12
                _ = v36
                _ = v37
            }
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    } else {
        if true {
            v50 := GenS4{F0: -1, F1: 42, F2: 32767}
            _ = v50
        }
    }
    var v51 int32 = 2147483647
    v52 := GenS4{F0: v51, F1: -1, F2: 42}
    if true {
        fmt.Println(v51)
        var v53 int16 = -1
        var v54 int8 = 1
        _ = v53
        _ = v54
    } else {
        if (42 < 1) {
            var v55 int16 = 42
            _ = v55
        }
        if true {
            if (true || true) {
                v51 += (v51 << ((255) & 63))
            } else {
                v51 = v51
                var v56 int8 = 100
                _ = v56
            }
            var v57 uint8 = 1
            _ = v57
        }
    }
    _ = v51
    _ = v52
    return -32768
}
func genFn1(p0 int16, p1 uint8, p2 uint) int16 {
    for fi58 := 0; fi58 < 7; fi58++ {
        for fi59 := 0; fi59 < 8; fi59++ {
            switch -506 {
            case 4294967295:
                var v60 float32 = 1.5
                _ = v60
            case -106:
                v61 := []int{fi58, -223, (fi59 * (fi58 ^ (fi58 >> (((p2 + p2)) & 63)))), (fi59 + -149)}
                _ = v61
            }
            var v62 uint32 = 2147483647
            _ = fi59
            _ = v62
        }
        _ = fi58
    }
    var v63 float64 = ((0.0 + 0.0) / -2.25)
    if true {
        var v64 float64 = v63
        v65 := GenS5{F0: -128}
        _ = v64
        _ = v65
    }
    p0 = 252
    v63 += 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v63
    return genFn0()
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 int = 776
    var v5 uint32 = 2
    var v6 bool = false
    switch "" {
    case "":
        var v66 float64 = 1e10
        _ = v66
    case "foo bar":
        fmt.Println(v6)
    }
    v67 := []int{v4, v4, v4}
    fmt.Println(v5, v5)
    switch (v4 / ((-9223372036854775808) | 1)) {
    case -9223372036854775808:
        for fi68 := 0; fi68 < 18; fi68++ {
            for fi69 := 0; fi69 < 6; fi69++ {
                var v70 bool = v6
                var v71 uint16 = 0
                var v72 float64 = 0.1
                _ = fi69
                _ = v70
                _ = v71
                _ = v72
            }
            _ = fi68
        }
    case 481:
        var v73 uint32 = (v5 | v5)
        _ = v73
    case 9223372036854775807:
        var v74 uint8 = 149
        _ = v74
    }
    fmt.Println(v4, v5, v4)
    fmt.Println(v4, v4, v5)
    switch v4 {
    case -492:
        _ = len(v1)
    default:
        if (false && false) {
            v4 = 9223372036854775807
            switch "a" {
            case "hello":
                var v75 int8 = 66
                _ = v75
            case "":
                for fi76 := 0; fi76 < 7; fi76++ {
                    for fi77 := 0; fi77 < 9; fi77++ {
                        v78 := v2[""]
                        v79 := v1[(fi77 + 875)]
                        if false {
                            if v6 {
                                var v80 int32 = 683
                                var v81 uint64 = 18446744073709551615
                                if v6 {
                                    var v82 float32 = (0.1 - (0.0 - 0.1))
                                    v6 = v6
                                    v83 := GenS9{F0: (v79 - v4), F1: v5}
                                    _ = v82
                                    _ = v83
                                }
                                _ = v80
                                _ = v81
                            }
                            v84 := GenS10{F0: 42, F1: (v79 | v79), F2: (v78 + "")}
                            v78 = v78
                            _ = v84
                        }
                        _ = fi77
                        _ = v78
                        _ = v79
                    }
                    _ = fi76
                }
            case "foo bar":
                var v85 int8 = -1
                _ = v85
            default:
                var v86 uint64 = 1
                _ = v86
            }
        } else {
            var v87 uint64 = 186
            fmt.Println(v4, v87)
            _ = v87
        }
    }
    var v88 float64 = (-2.25 + 3.14159)
    if len(v67) > 0 {
        v67[0] = (v4 >> ((1) & 63))
    }
    v89 := GenS2{F0: -2147483648}
    var v90 uint8 = 255
    fmt.Println(v6)
    v67 = append(v67, v4)
    v91 := GenS4{F0: -1, F1: 0, F2: 1}
    v92 := v3[(v4 % ((-454) | 1))]
    switch 32768 {
    case -190:
        for fi93 := 0; fi93 < 11; fi93++ {
            v94 := GenS11{F0: 539, F1: (!v6), F2: uint(v90)}
            v95 := v2[""]
            _ = fi93
            _ = v94
            _ = v95
        }
    default:
        v96 := GenS12{F0: v90, F1: -1, F2: 42}
        _ = v96
    }
    v90 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v88:", v88)
    fmt.Println("v90:", v90)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v88
    _ = v89
    _ = v90
    _ = v91
    _ = v92
}
