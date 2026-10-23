package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int }
type GenS2 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 uint8; F1 int64; F2 int32 }
func genFn0(p0 int8, p1 uint8, p2 int8) int8 {
    v4 := []int{}
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 bool = true
        var v7 bool = (false && (int64(p1) >= -1))
        v6 = (true && (v7 || false))
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        p2 = -27
        var v9 float32 = 1e10
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p0
}
func genFn1() uint16 {
    var v10 int64 = 42
    v10 -= v10
    fmt.Println(v10)
    _ = v10
    return 401
}

func main() {
    var v1 int64 = 42
    var v2 uint = 18446744073709551615
    var v3 bool = (3.14159 == (((3.14159 - (3.14159 + ((-2.25 * 1e10) + 0.1))) * (float32((1e10 / -2.25)) - 0.0)) - ((1e10 * (1e10 / 1e10)) - 1e10)))
    var v11 uint64 = 2
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 int = (fi13 >> ((v2) & 63))
            switch (v12 + "a") {
            case "\n\t":
                v16 := GenS1{F0: 1, F1: uint16(v11), F2: -2147483648}
                _ = v16
            default:
                switch "a" {
                case "":
                    fmt.Println(v1, v1)
                case "a":
                    var v17 uint = (v2 | 4294967295)
                    _ = v17
                case "foo bar":
                    v15 = 729
                default:
                    v18 := []int{v15, fi13, 9223372036854775807, (fi14 / (((^(fi13))) | 1))}
                    _ = v18
                }
            }
            var v19 int64 = v1
            _ = fi14
            _ = v15
            _ = v19
        }
        _ = fi13
    }
    var v20 uint8 = 2
    v20 += v20
    for fi21 := 0; fi21 < 2; fi21++ {
        switch -299 {
        case 127:
            v20 = v20
        case -1:
            fmt.Println(v1, v20)
        case 13:
            v22 := []int{}
            _ = v22
        }
        fmt.Println(v2, v12)
        _ = fi21
    }
    switch "" {
    case "":
        v23 := []int{-489}
        _ = v23
    default:
        if false {
            fmt.Println(v3, v12, v1)
            fmt.Println(v11, v3)
            if v3 {
                var v24 int64 = int64(v2)
                if (1 != 32768) {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        var v26 string = "a"
                        if true {
                            v3 = (!(true || (!(((v2 * (^(v2))) != (v2 | (v2 | (v2 | 18446744073709551615)))) && (-2147483648 < 42)))))
                            var v27 float64 = (float64(((((-2.25 + 1e10) - 1e10) + ((0.0 * 0.1) - 1.5)) + -2.25)) - 1.5)
                            _ = v27
                        }
                        var v28 int8 = 91
                        _ = fi25
                        _ = v26
                        _ = v28
                    }
                    v29 := GenS1{F0: v20, F1: 1, F2: -792}
                    _ = v29
                }
                v30 := GenS2{F0: v11, F1: -128, F2: (v20 | (v20 / (((v20 >> ((v2) & 63))) | 1)))}
                _ = v24
                _ = v30
            } else {
                var v31 bool = (uint64(v2) < 171)
                if v31 {
                    fmt.Println(v31)
                    var v32 uint16 = 2
                    _ = v32
                }
                if (42 >= uint8(v11)) {
                    v1 = v1
                    for fi33 := 0; fi33 < 12; fi33++ {
                        switch (fi33 << ((1) & 63)) {
                        case -1:
                            v34 := []int{(^(fi33)), fi33, fi33, -2147483648}
                            _ = v34
                        case -2147483648:
                            for fi35 := 0; fi35 < 19; fi35++ {
                                var v36 int32 = -887
                                var v37 string = "a"
                                fmt.Println(v31, v20, v1)
                                v1 -= (v1 + (^(v1)))
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        case 32767:
                            var v38 uint8 = v20
                            _ = v38
                        default:
                            switch v12 {
                            case "hello":
                                v2 = 547
                            case "a":
                                switch v12 {
                                case "a":
                                    v20 -= (v20 & (v20 + (^(v20))))
                                case "\n\t":
                                    v31 = (!(!(!true)))
                                default:
                                    if v3 {
                                        for fi39 := 0; fi39 < 3; fi39++ {
                                            var v40 uint8 = v20
                                            switch ("a" + "a") {
                                            case "a":
                                                var v41 bool = (!v3)
                                                _ = v41
                                            default:
                                                v42 := []int{fi39}
                                                _ = v42
                                            }
                                            if v3 {
                                                v43 := []int{}
                                                var v44 int = fi39
                                                _ = len(v43)
                                                _ = v43
                                                _ = v44
                                            } else {
                                                fmt.Println(v11, fi39, v40)
                                                if v31 {
                                                    fmt.Println(fi33)
                                                    fmt.Println(v12)
                                                    switch 207 {
                                                    case -708:
                                                        var v45 int8 = 1
                                                        _ = v45
                                                    }
                                                    var v46 int32 = -1
                                                    v46 -= v46
                                                    _ = v46
                                                } else {
                                                    v47 := GenS2{F0: uint64(v20), F1: 127, F2: (v20 + (v20 - v20))}
                                                    fmt.Println(fi33, v31, v12)
                                                    fmt.Println(fi33)
                                                    _ = v47
                                                }
                                                if ((-(fi39)) >= fi33) {
                                                    switch fi33 {
                                                    case 32767:
                                                        fmt.Println(v12, fi33)
                                                    case 771:
                                                        var v48 uint8 = (v40 >> ((v2) & 63))
                                                        _ = v48
                                                    case -2147483648:
                                                        v31 = (v31 && (!(127 > int8(v11))))
                                                    default:
                                                        if true {
                                                            v49 := GenS1{F0: (^(v40)), F1: uint16(v11), F2: fi39}
                                                            _ = v49
                                                        } else {
                                                            v50 := GenS3{F0: v12}
                                                            var v51 int = fi33
                                                            switch (v12 + "a") {
                                                            case "\n\t":
                                                                if (!v3) {
                                                                    switch (v12 + (v12 + "hello")) {
                                                                    case "hello":
                                                                        switch "" {
                                                                        case "":
                                                                            v52 := []int{}
                                                                            _ = v52
                                                                        case "\n\t":
                                                                            if (1 < 0) {
                                                                                switch (-(fi33)) {
                                                                                case 42:
                                                                                    var v53 bool = v31
                                                                                    _ = v53
                                                                                case -146:
                                                                                    v54 := GenS3{F0: v12}
                                                                                    _ = v54
                                                                                default:
                                                                                    var v55 int8 = 127
                                                                                    _ = v55
                                                                                }
                                                                                v56 := GenS3{F0: "foo bar"}
                                                                                var v57 string = "\n\t"
                                                                                _ = v56
                                                                                _ = v57
                                                                            } else {
                                                                                switch (v12 + (v12 + ("hello" + v12))) {
                                                                                case "hello":
                                                                                    v58 := []int{fi39, 32768, (^(fi33))}
                                                                                    _ = v58
                                                                                case "":
                                                                                    var v59 uint32 = 0
                                                                                    _ = v59
                                                                                default:
                                                                                    var v60 string = ((v12 + "\n\t") + v12)
                                                                                    _ = v60
                                                                                }
                                                                                fmt.Println(v31, v3)
                                                                                v20 = (^(v20))
                                                                                var v61 int8 = -128
                                                                                if ("foo bar" <= ("" + v12)) {
                                                                                    if v31 {
                                                                                        var v62 int64 = v1
                                                                                        v61 -= v61
                                                                                        v20 += (v20 / (((v40 | 2)) | 1))
                                                                                        _ = v62
                                                                                    } else {
                                                                                        v51 += fi33
                                                                                        if (287 <= 65535) {
                                                                                            v63 := GenS1{F0: (v40 + (v40 >> ((v2) & 63))), F1: 2, F2: (fi33 ^ v51)}
                                                                                            fmt.Println(v51)
                                                                                            _ = v63
                                                                                        } else {
                                                                                            if v31 {
                                                                                                v64 := []int{42, (-(v51)), 32768}
                                                                                                _ = v64
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if true {
                                                                                        fmt.Println(v1, v1)
                                                                                        v65 := []int{fi33, fi39, 127}
                                                                                        v66 := GenS4{F0: -32768, F1: 0}
                                                                                        v40 = v20
                                                                                        var v67 int = (-(fi39))
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    }
                                                                                } else {
                                                                                    var v68 uint16 = 1
                                                                                    _ = v68
                                                                                }
                                                                                _ = v61
                                                                            }
                                                                        case "a":
                                                                            fmt.Println(v3, fi39)
                                                                        }
                                                                    }
                                                                    if (-1 <= 1000) {
                                                                        if (v3 && ((-32768 != -735) && v31)) {
                                                                            v69 := []int{(fi39 + (fi33 - 0)), (fi33 - (fi39 + int(v1)))}
                                                                            if (v40 >= (^(v40))) {
                                                                                var v70 int8 = -63
                                                                                _ = v70
                                                                            }
                                                                            _ = v69
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v1, v1)
                                                                        var v71 int16 = -509
                                                                        v72 := GenS4{F0: (v71 & -1), F1: 355}
                                                                        var v73 int16 = int16(fi39)
                                                                        switch fi39 {
                                                                        case -2147483648:
                                                                            v74 := GenS3{F0: v12}
                                                                            _ = v74
                                                                        case 564:
                                                                            var v75 int = (fi33 | fi39)
                                                                            _ = v75
                                                                        default:
                                                                            v40 -= (v40 - (v20 ^ v20))
                                                                        }
                                                                        _ = v71
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    fmt.Println(v40, fi39, v2)
                                                                } else {
                                                                    fmt.Println(v3, v51, v2)
                                                                    if false {
                                                                        var v76 uint8 = v40
                                                                        switch v12 {
                                                                        case "\n\t":
                                                                            v77 := GenS2{F0: 19, F1: -21, F2: v40}
                                                                            _ = v77
                                                                        case "foo bar":
                                                                            var v78 uint32 = 306
                                                                            _ = v78
                                                                        default:
                                                                            if v3 {
                                                                                v3 = true
                                                                                v20 = 2
                                                                                var v79 float64 = -2.25
                                                                                _ = v79
                                                                            } else {
                                                                                v31 = v3
                                                                                var v80 int64 = v1
                                                                                _ = v80
                                                                            }
                                                                        }
                                                                        v81 := GenS3{F0: "a"}
                                                                        _ = v76
                                                                        _ = v81
                                                                    }
                                                                }
                                                            case "hello":
                                                                v2 = (v2 - 2)
                                                            }
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    }
                                                } else {
                                                    v40 -= v40
                                                    switch "" {
                                                    case "":
                                                        v82 := GenS1{F0: 2, F1: 490, F2: (fi33 ^ 4294967295)}
                                                        _ = v82
                                                    default:
                                                        if true {
                                                            v20 -= (v20 % (((v20 << ((v2) & 63))) | 1))
                                                        }
                                                    }
                                                }
                                            }
                                            _ = fi39
                                            _ = v40
                                        }
                                    }
                                }
                            }
                        }
                        v83 := GenS5{F0: (v1 | (v1 * -9223372036854775808)), F1: true}
                        if (v2 == (v2 + (v2 - v2))) {
                            var v84 float64 = 3.14159
                            _ = v84
                        }
                        _ = fi33
                        _ = v83
                    }
                    for fi85 := 0; fi85 < 14; fi85++ {
                        fmt.Println(v3, v3, v20)
                        var v86 int = fi85
                        switch (-(v86)) {
                        case -9223372036854775808:
                            v87 := GenS4{F0: 83, F1: -2147483648}
                            _ = v87
                        default:
                            if ((v12 > v12) && v31) {
                                for fi88 := 0; fi88 < 19; fi88++ {
                                    switch "" {
                                    case "foo bar":
                                        v31 = false
                                    case "a":
                                        var v89 int = 1
                                        _ = v89
                                    }
                                    _ = fi88
                                }
                                v1 = (v1 & (v1 + (v1 & (v1 + 42))))
                                var v90 int32 = -1
                                _ = v90
                            }
                        }
                        _ = fi85
                        _ = v86
                    }
                }
                _ = v31
            }
        } else {
            fmt.Println(v11, v3)
            v3 = (127 > 42)
        }
    }
    v91 := GenS3{F0: ("hello" + ("\n\t" + ((((v12 + v12) + (("\n\t" + (("a" + (((v12 + "foo bar") + v12) + ((v12 + (v12 + "foo bar")) + v12))) + v12)) + (v12 + (v12 + "")))) + v12) + "")))}
    var v92 uint32 = 4294967295
    var v93 bool = v3
    if (1 > uint32(v20)) {
        fmt.Println(v2, v92)
    } else {
        v94 := GenS2{F0: (^(v11)), F1: int8(v11), F2: 42}
        _ = v94
    }
    v95 := []int{42, 127}
    var v96 int32 = -2147483648
    var v97 int16 = 1000
    v98 := GenS6{F0: (v20 >> ((v2) & 63)), F1: int64(v2), F2: v96}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("slicelen v95:", len(v95), v95)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v91
    _ = v92
    _ = v93
    _ = v95
    _ = v96
    _ = v97
    _ = v98
}
