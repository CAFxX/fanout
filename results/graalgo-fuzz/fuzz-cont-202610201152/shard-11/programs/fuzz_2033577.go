package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }
type GenS2 struct { F0 int64; F1 int32; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint8; F2 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 uint; F2 uint }
type GenS8 struct { F0 int32 }
func genFn0() int {
    var v6 int32 = -38
    if true {
        var v7 uint = 443
        v8 := GenS1{F0: 42, F1: 9223372036854775807, F2: ""}
        _ = v7
        _ = v8
    }
    var v9 uint16 = 255
    v10 := []int{224, -1, 2147483647, -48}
    _ = v10
    _ = v6
    _ = v9
    return 4294967295
}
func genFn1(p0 uint64) string {
    for fi11 := 0; fi11 < 4; fi11++ {
        switch "a" {
        case "":
            v12 := []int{(fi11 | genFn0()), (fi11 ^ genFn0())}
            _ = v12
        case "hello":
            p0 += p0
        case "\n\t":
            v13 := GenS1{F0: 9223372036854775807, F1: -996, F2: "\n\t"}
            _ = v13
        default:
            p0 -= (p0 + (p0 + 1))
        }
        v14 := GenS1{F0: 42, F1: 9223372036854775807, F2: "a"}
        _ = fi11
        _ = v14
    }
    if false {
        fmt.Println(p0)
        if false {
            var v15 int64 = -801
            var v16 uint64 = p0
            _ = v15
            _ = v16
        }
        var v17 float64 = (1.5 + float64((-2.25 + 1.5)))
        _ = v17
    }
    p0 = (p0 & 18446744073709551615)
    _ = p0
    return ("\n\t" + "foo bar")
}

func main() {
    var v1 uint8 = 42
    var v2 uint64 = 638
    v3 := make(map[uint32]int)
    var v4 bool = ((0 != 65535) || true)
    var v5 int32 = 42
    for fi18 := 0; fi18 < 7; fi18++ {
        switch "hello" {
        case "hello":
            var v19 float32 = (0.1 + (0.1 / 1.5))
            _ = v19
        case "":
            v20 := GenS1{F0: 0, F1: 42, F2: "a"}
            _ = v20
        case "foo bar":
            for fi21 := 0; fi21 < 16; fi21++ {
                var v22 float64 = 0.0
                switch (fi18 * 32767) {
                case 560:
                    if false {
                        if true {
                            v1 += 42
                            fmt.Println(fi21, v22)
                            delete(v3, 2147483647)
                        }
                        v23 := GenS1{F0: -1, F1: int64(v1), F2: "\n\t"}
                        fmt.Println(v4, v5)
                        _ = v23
                    } else {
                        v24 := GenS1{F0: -139, F1: -540, F2: "hello"}
                        v25 := GenS2{F0: 0, F1: v5, F2: (v5 | (v5 + (v5 & -795)))}
                        _ = v24
                        _ = v25
                    }
                default:
                    var v26 int32 = 1
                    _ = v26
                }
                switch fi18 {
                case 42:
                    v5 = 2147483647
                default:
                    delete(v3, uint32(v1))
                    if true {
                        switch "" {
                        case "foo bar":
                            var v27 uint8 = v1
                            _ = v27
                        default:
                            var v28 int32 = v5
                            _ = v28
                        }
                        switch (fi21 & fi18) {
                        case 0:
                            if true {
                                v29 := v3[4294967295]
                                _ = v29
                            }
                        case 819:
                            var v30 string = ((("" + "foo bar") + "\n\t") + "")
                            _ = v30
                        case -1:
                            v31 := GenS3{F0: (fi18 & 1)}
                            _ = v31
                        }
                        if false {
                            v32 := []int{fi18, fi21, fi18}
                            var v33 float32 = 1e10
                            if false {
                                switch fi21 {
                                case -9223372036854775808:
                                    v3[966] = genFn0()
                                default:
                                    if (v4 || (("\n\t" == "a") || ((fi21 / (((fi18 % ((fi21) | 1))) | 1)) == -237))) {
                                        v34 := GenS4{F0: (fi21 & (fi18 | fi18)), F1: v1, F2: "\n\t"}
                                        if true {
                                            var v35 int = 42
                                            var v36 string = "a"
                                            var v37 uint16 = 450
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                        _ = v34
                                    } else {
                                        switch -9223372036854775808 {
                                        case 32768:
                                            v38 := GenS5{F0: true}
                                            _ = v38
                                        }
                                        var v39 int = (fi18 + fi18)
                                        var v40 int16 = 212
                                        var v41 float64 = (((0.1 + v22) * v22) * v22)
                                        if false {
                                            switch genFn0() {
                                            case 89:
                                                var v42 string = "foo bar"
                                                _ = v42
                                            case 4294967295:
                                                v43 := GenS3{F0: fi21}
                                                _ = v43
                                            }
                                        } else {
                                            if (69 > 563) {
                                                var v44 uint64 = v2
                                                v1 += v1
                                                var v45 uint16 = 32767
                                                _ = v44
                                                _ = v45
                                            }
                                            v46 := GenS6{F0: 100}
                                            switch "hello" {
                                            case "foo bar":
                                                var v47 uint = 993
                                                _ = v47
                                            case "":
                                                var v48 int32 = v5
                                                _ = v48
                                            case "\n\t":
                                                v22 -= (3.14159 / float64(0.1))
                                            }
                                            _ = v46
                                        }
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                }
                            }
                            _ = v32
                            _ = v33
                        } else {
                            if v4 {
                                if ((v5 << ((uint(v1)) & 63)) == v5) {
                                    v49 := []int{(^(fi18))}
                                    fmt.Println(v5)
                                    v50 := GenS4{F0: 701, F1: v1, F2: ""}
                                    _ = v49
                                    _ = v50
                                }
                                v51 := []int{0, fi21, fi18, fi21}
                                _ = v51
                            } else {
                                fmt.Println(v5)
                                var v52 uint16 = 255
                                var v53 uint32 = 32
                                _ = v52
                                _ = v53
                            }
                            delete(v3, 2)
                            switch "hello" {
                            case "foo bar":
                                var v54 int32 = v5
                                v5 -= (-(v5))
                                var v55 uint32 = 65535
                                if (float32(v22) >= 1.5) {
                                    if true {
                                        var v56 int8 = -1
                                        v55 = v55
                                        if ((v56 << ((4294967295) & 63)) == v56) {
                                            v22 = v22
                                            switch (fi18 | (fi18 >> ((42) & 63))) {
                                            case 128:
                                                v57 := GenS7{F0: v2, F1: 68, F2: 1}
                                                _ = v57
                                            }
                                            v1 = (v1 * v1)
                                            v58 := GenS8{F0: (v5 % (((-(v54))) | 1))}
                                            v59 := GenS3{F0: 449}
                                            _ = v58
                                            _ = v59
                                        }
                                        _ = v56
                                    }
                                    v54 = (v54 / ((749) | 1))
                                } else {
                                    v55 += v55
                                    switch (fi18 - (fi18 - (fi18 << ((uint(v55)) & 63)))) {
                                    case 0:
                                        v4 = (42 == 572)
                                        v55 += v55
                                        var v60 int8 = -1
                                        _ = len(v3)
                                        _ = v60
                                    case -2147483648:
                                        v22 -= v22
                                        v54 -= (-(v5))
                                    }
                                    v2 += (v2 >> ((690) & 63))
                                }
                                fmt.Println(fi18, v4)
                                if false {
                                    var v61 uint32 = 616
                                    v62 := []int{(fi18 | fi18), -88, fi21, 32768}
                                    _ = v61
                                    _ = v62
                                } else {
                                    v63 := GenS3{F0: int(v2)}
                                    fmt.Println(fi21, fi21)
                                    _ = v63
                                }
                                v2 = v2
                                fmt.Println(v4, v54)
                                _ = v54
                                _ = v55
                            }
                        }
                    }
                }
                switch "a" {
                case "":
                    fmt.Println(fi18)
                case "foo bar":
                    var v64 float32 = 1.5
                    _ = v64
                case "\n\t":
                    v65 := []int{(fi21 << ((0) & 63))}
                    var v66 int = (fi21 / (((-(fi18))) | 1))
                    switch "" {
                    case "foo bar":
                        v67 := GenS1{F0: 9223372036854775807, F1: 42, F2: ""}
                        _ = v67
                    }
                    v66 = (-(fi21))
                    delete(v3, 42)
                    _ = v65
                    _ = v66
                default:
                    var v68 int32 = (v5 % (((-(v5))) | 1))
                    _ = v68
                }
                _ = fi21
                _ = v22
            }
        default:
            for fi69 := 0; fi69 < 20; fi69++ {
                v70 := []int{fi18, (fi69 ^ (fi69 / (((fi18 - (fi69 >> ((255) & 63)))) | 1))), genFn0(), (fi69 & fi69)}
                var v71 int64 = -9223372036854775808
                fmt.Println(v1, fi69)
                _ = fi69
                _ = v70
                _ = v71
            }
        }
        v5 = 53
        var v72 bool = (!v4)
        _ = fi18
        _ = v72
    }
    if (384 != int32(v1)) {
        for fi73 := 0; fi73 < 15; fi73++ {
            v2 = (v2 + 18446744073709551615)
            _ = fi73
        }
        v74 := []int{351, -451}
        v74 = append(v74, 42)
        _ = v74
    }
    var v75 int = 1
    for fi76 := 0; fi76 < 11; fi76++ {
        for fi77 := 0; fi77 < 3; fi77++ {
            if v4 {
                var v78 uint = 255
                _ = v78
            } else {
                if (int16(fi76) >= 13) {
                    var v79 uint = 1
                    fmt.Println(v79, v5, v79)
                    _ = v79
                } else {
                    v1 = 2
                }
            }
            v2 = 42
            _ = fi77
        }
        _ = fi76
    }
    switch v75 {
    case -392:
        if (v75 > (v75 / (((v75 + v75)) | 1))) {
            fmt.Println(v75, v1)
        } else {
            delete(v3, 853)
            var v80 uint64 = v2
            if v4 {
                v3[42] = 42
                if (65535 >= 42) {
                    var v81 int = v75
                    _ = v81
                } else {
                    v2 = v80
                }
            }
            _ = v80
        }
    }
    _ = len(v3)
    v82 := GenS8{F0: 2147483647}
    var v83 int8 = int8(v2)
    v4 = v4
    if v4 {
        fmt.Println(v4, v83, v4)
    } else {
        v84 := v3[2147483648]
        var v85 uint = 625
        delete(v3, 65535)
        _ = v84
        _ = v85
    }
    if (0.1 <= 3.14159) {
        if false {
            var v86 int32 = (v5 / ((42) | 1))
            _ = v86
        }
    } else {
        var v87 int8 = v83
        if v4 {
            v88 := []int{v75, genFn0(), (v75 / (((v75 ^ v75)) | 1)), (v75 | (v75 % ((v75) | 1)))}
            for fi89 := 0; fi89 < 5; fi89++ {
                var v90 float64 = 1e10
                switch 32767 {
                case 0:
                    switch (fi89 >> ((2) & 63)) {
                    case 32768:
                        fmt.Println(v5)
                    }
                case 32767:
                    var v91 int8 = (v87 ^ v87)
                    _ = v91
                case -687:
                    v1 += 42
                }
                _ = fi89
                _ = v90
            }
            _ = v88
        } else {
            _ = len(v3)
            fmt.Println(v2)
        }
        _ = v87
    }
    v92 := []int{-2147483648, v75}
    v83 += v83
    v2 = uint64(v1)
    for fi93 := 0; fi93 < 4; fi93++ {
        var v94 uint32 = 2147483647
        _ = fi93
        _ = v94
    }
    var v95 float64 = (1e10 * 1.5)
    var v96 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v75:", v75)
    fmt.Println("v83:", v83)
    fmt.Println("slicelen v92:", len(v92), v92)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v82
    _ = v83
    _ = v92
    _ = v95
    _ = v96
}
