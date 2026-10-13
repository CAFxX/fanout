package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int32; F2 string }
func genFn0(p0 int, p1 uint) uint16 {
    switch (p0 & (p0 & -426)) {
    case 470:
        var v6 float32 = 3.14159
        _ = v6
    case -1:
        var v7 uint8 = 255
        _ = v7
    case 127:
        p1 = (^(p1))
    }
    var v8 float32 = 1.5
    if false {
        var v9 uint16 = 16
        var v10 int64 = 1
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 float32 = v8
            _ = fi11
            _ = v12
        }
        _ = v10
        _ = v9
    } else {
        v13 := GenS1{F0: -9223372036854775808}
        if true {
            v14 := []int{9223372036854775807, (-(p0)), (p0 * p0)}
            var v15 uint = p1
            _ = v14
            _ = v15
        }
        v16 := []int{-814, (-(p0))}
        _ = v13
        _ = v16
    }
    fmt.Println(p0, v8, p1)
    for fi17 := 0; fi17 < 19; fi17++ {
        v8 = v8
        switch ("a" + "hello") {
        case "hello":
            var v18 int8 = 100
            _ = v18
        case "\n\t":
            v19 := []int{(-(p0)), (fi17 + (-(p0)))}
            _ = v19
        case "foo bar":
            v20 := []int{}
            _ = v20
        }
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = v8
    return 65535
}
func genFn1(p0 int16, p1 bool) int32 {
    switch -173 {
    case 9223372036854775807:
        v21 := GenS2{F0: 0}
        _ = v21
    }
    v22 := GenS2{F0: 0}
    _ = p0
    _ = p1
    _ = v22
    return 0
}

func main() {
    var v1 uint = 255
    var v2 uint32 = 2
    var v3 uint8 = 2
    var v4 uint32 = (v2 << ((v1) & 63))
    var v5 float64 = float64(((-2.25 + 1.5) - (0.0 - -2.25)))
    if (true || true) {
        v23 := GenS2{F0: 2}
        fmt.Println(v2, v4)
        _ = v23
    } else {
        v24 := []int{-1}
        if ((!true) && (!true)) {
            v5 = v5
            if (4294967295 > 788) {
                var v25 uint64 = 2
                v26 := GenS2{F0: (^(v25))}
                _ = v25
                _ = v26
            }
        }
        _ = len(v24)
        _ = v24
    }
    if false {
        v27 := GenS2{F0: 18446744073709551615}
        fmt.Println(v2, v4, v5)
        _ = v27
    } else {
        if (!false) {
            v28 := GenS2{F0: 42}
            _ = v28
        } else {
            v29 := GenS3{F0: -372}
            for fi30 := 0; fi30 < 12; fi30++ {
                v31 := []int{(^(fi30)), (fi30 % ((fi30) | 1)), fi30}
                var v32 int8 = -128
                if true {
                    v33 := GenS3{F0: fi30}
                    switch ("a" + "hello") {
                    case "a":
                        var v34 int16 = int16(v32)
                        _ = v34
                    default:
                        var v35 uint64 = 1
                        _ = v35
                    }
                    _ = v33
                }
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v29
        }
        v4 = v4
        var v36 string = (("hello" + ("" + ("hello" + ("\n\t" + "hello")))) + "foo bar")
        _ = v36
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        v38 := []int{(fi37 * (^(fi37))), fi37, 9223372036854775807}
        _ = fi37
        _ = v38
    }
    v1 = v1
    for fi39 := 0; fi39 < 11; fi39++ {
        var v40 int8 = 100
        _ = fi39
        _ = v40
    }
    if false {
        switch "hello" {
        case "\n\t":
            v41 := GenS3{F0: 768}
            _ = v41
        case "a":
            var v42 uint = 42
            _ = v42
        default:
            var v43 bool = (float64(float32((v5 + 0.1))) >= v5)
            _ = v43
        }
        if false {
            for fi44 := 0; fi44 < 6; fi44++ {
                for fi45 := 0; fi45 < 12; fi45++ {
                    var v46 uint16 = genFn0((fi45 * (fi44 / ((fi45) | 1))), v1)
                    _ = fi45
                    _ = v46
                }
                v47 := []int{fi44, 414, (fi44 << ((936) & 63))}
                v5 = (3.14159 / v5)
                _ = fi44
                _ = v47
            }
            if (v5 != (v5 / (1.5 * (v5 * v5)))) {
                v4 += (v2 / ((v2) | 1))
                var v48 uint16 = 32767
                for fi49 := 0; fi49 < 9; fi49++ {
                    v50 := GenS2{F0: uint64(v1)}
                    _ = fi49
                    _ = v50
                }
                _ = v48
            } else {
                v51 := []int{int(v3), 0, 97}
                v5 = float64(float32(v5))
                if false {
                    for fi52 := 0; fi52 < 19; fi52++ {
                        var v53 int8 = 127
                        var v54 int64 = 9223372036854775807
                        var v55 int = 128
                        _ = fi52
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                }
                _ = v51
            }
            v2 += 678
        }
        var v56 uint16 = genFn0(-2147483648, v1)
        _ = v56
    }
    var v57 int16 = 42
    for fi58 := 0; fi58 < 11; fi58++ {
        for fi59 := 0; fi59 < 12; fi59++ {
            fmt.Println(v4)
            var v60 uint = (v1 / ((42) | 1))
            v61 := GenS1{F0: (^(fi59))}
            var v62 string = "foo bar"
            switch "" {
            case "hello":
                var v63 int32 = 0
                _ = v63
            case "foo bar":
                var v64 float64 = (v5 + (v5 - 3.14159))
                _ = v64
            case "\n\t":
                if true {
                    var v65 uint = v1
                    var v66 int64 = -1
                    if true {
                        v3 = 132
                    }
                    _ = v65
                    _ = v66
                } else {
                    var v67 int32 = 0
                    if (!(-128 < 42)) {
                        v68 := []int{(-(fi59)), (fi58 << ((255) & 63)), 0}
                        if len(v68) >= 2 {
                            v68 = v68[1:len(v68)]
                        }
                        v69 := GenS1{F0: 2147483647}
                        _ = v68
                        _ = v69
                    }
                    if true {
                        var v70 bool = (!(!true))
                        var v71 float32 = (3.14159 + 0.0)
                        var v72 bool = ("a" >= v62)
                        if v70 {
                            v3 -= 98
                        } else {
                            if true {
                                fmt.Println(v72)
                                v67 = v67
                                v73 := []int{9223372036854775807, 9223372036854775807, (fi59 | int(v4))}
                                _ = v73
                            }
                            fmt.Println(fi59, v62)
                            v72 = true
                        }
                        _ = v70
                        _ = v71
                        _ = v72
                    }
                    v74 := []int{0, fi59, -438}
                    var v75 int8 = 0
                    v76 := GenS4{F0: (v67 % ((genFn1(v57, false)) | 1)), F1: 65535, F2: 798}
                    var v77 uint = uint(v3)
                    fmt.Println(v67, fi59, v5)
                    if (((v5 / float64(0.1)) + (0.1 + float64((((float32(v5) + (-2.25 - (float32(v5) - (0.1 - (-2.25 - ((0.0 / (float32(1.5) * 1e10)) / 3.14159)))))) / float32(v5)) - (-2.25 + float32(float64(3.14159))))))) != float64(1.5)) {
                        v78 := GenS3{F0: (^(fi58))}
                        _ = v78
                    } else {
                        var v79 int8 = -1
                        _ = len(v74)
                        _ = v79
                    }
                    switch fi58 {
                    case 42:
                        v80 := GenS3{F0: fi58}
                        _ = v80
                    case 127:
                        if true {
                            v67 = int32(v1)
                            fmt.Println(v75, fi58, v62)
                        } else {
                            var v81 uint64 = 57
                            if len(v74) > 0 {
                                v74[0] = 4294967295
                            }
                            if ((v57 * 1) > -968) {
                                v82 := GenS3{F0: int(v77)}
                                _ = v82
                            } else {
                                if (42 != v67) {
                                    fmt.Println(v75, fi58, v67)
                                    var v83 int8 = (v75 ^ v75)
                                    if (false || ((^(v81)) >= (^(v81)))) {
                                        v3 -= uint8(v4)
                                    } else {
                                        var v84 uint32 = (v2 % ((v4) | 1))
                                        v5 -= (float64(0.0) + -2.25)
                                        _ = v84
                                    }
                                    _ = v83
                                }
                            }
                            _ = v81
                        }
                    case -326:
                        fmt.Println(v5)
                    }
                    _ = v67
                    _ = v74
                    _ = v75
                    _ = v76
                    _ = v77
                }
            }
            _ = fi59
            _ = v60
            _ = v61
            _ = v62
        }
        for fi85 := 0; fi85 < 17; fi85++ {
            var v86 int64 = 9223372036854775807
            v87 := GenS1{F0: fi58}
            switch "foo bar" {
            case "\n\t":
                var v88 float32 = (0.0 / (float32(1.5) + 3.14159))
                _ = v88
            case "foo bar":
                v89 := []int{(fi58 << ((502) & 63)), fi85}
                _ = v89
            case "hello":
                v2 = uint32(v1)
            default:
                var v90 uint16 = 255
                _ = v90
            }
            _ = fi85
            _ = v86
            _ = v87
        }
        _ = fi58
    }
    v91 := GenS3{F0: 128}
    var v92 bool = false
    for fi93 := 0; fi93 < 18; fi93++ {
        v94 := GenS5{F0: v2, F1: 486, F2: "foo bar"}
        fmt.Println(v5)
        _ = fi93
        _ = v94
    }
    v95 := []int{}
    switch "" {
    case "":
        for fi96 := 0; fi96 < 16; fi96++ {
            if len(v95) > 0 {
                v95[0] = -9223372036854775808
            }
            _ = fi96
        }
    case "a":
        var v97 bool = (0 == 100)
        _ = v97
    }
    var v98 uint16 = genFn0(42, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v92:", v92)
    fmt.Println("slicelen v95:", len(v95), v95)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v91
    _ = v92
    _ = v95
    _ = v98
}
