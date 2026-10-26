package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint16; F2 int32 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 int16; F1 int32; F2 uint8 }
type GenS10 struct { F0 int16; F1 int64 }
type GenS11 struct { F0 uint32 }
func genFn0(p0 bool) int8 {
    v4 := []int{-2147483648, 42, 4294967295}
    if p0 {
        var v5 uint = 18446744073709551615
        _ = v5
    } else {
        fmt.Println(p0)
        for fi6 := 0; fi6 < 5; fi6++ {
            p0 = true
            for fi7 := 0; fi7 < 17; fi7++ {
                p0 = (!(false || ((fi7 != 32767) && (987 > 42))))
                switch "\n\t" {
                case "\n\t":
                    p0 = p0
                case "hello":
                    v4 = append(v4, -47)
                default:
                    switch ("" + "") {
                    case "":
                        p0 = p0
                    case "\n\t":
                        if true {
                            var v8 uint16 = 2
                            if (!(fi6 < 137)) {
                                p0 = p0
                                p0 = (806 < 4294967295)
                            } else {
                                var v9 uint32 = 184
                                _ = v9
                            }
                            switch fi6 {
                            case -2147483648:
                                if p0 {
                                    p0 = ("hello" > "hello")
                                    var v10 int64 = 948
                                    _ = v10
                                } else {
                                    fmt.Println(fi6)
                                    var v11 int = (fi6 / ((-469) | 1))
                                    var v12 int64 = 9223372036854775807
                                    _ = len(v4)
                                    v13 := GenS1{F0: -128, F1: 743}
                                    if (0 <= 9223372036854775807) {
                                        if (1 < 0) {
                                            if p0 {
                                                v14 := GenS2{F0: 1000, F1: 2147483648}
                                                _ = v14
                                            } else {
                                                v15 := GenS3{F0: 42, F1: (v8 - v8), F2: 2147483647}
                                                var v16 int8 = 11
                                                var v17 uint8 = uint8(v8)
                                                _ = v15
                                                _ = v16
                                                _ = v17
                                            }
                                            v11 += fi7
                                            if len(v4) > 0 {
                                                v4[0] = (fi7 << ((0) & 63))
                                            }
                                            var v18 int8 = -1
                                            fmt.Println(v18, v11, fi7)
                                            _ = v18
                                        } else {
                                            if p0 {
                                                v11 = (fi6 & fi6)
                                            }
                                            switch "" {
                                            case "\n\t":
                                                var v19 string = "\n\t"
                                                _ = v19
                                            case "foo bar":
                                                v12 = v12
                                            case "hello":
                                                fmt.Println(fi7, v8, fi6)
                                            default:
                                                var v20 int = 32767
                                                _ = v20
                                            }
                                        }
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if ((2 < 1) || p0) {
                                            fmt.Println(v8)
                                        } else {
                                            v21 := GenS4{F0: 0, F1: false}
                                            if (1 >= 2) {
                                                var v22 int16 = 32767
                                                switch ((("\n\t" + "hello") + "a") + "hello") {
                                                case "\n\t":
                                                    switch (^(fi7)) {
                                                    case 9:
                                                        _ = len(v4)
                                                    case -9223372036854775808:
                                                        v4 = append(v4, 2147483647)
                                                    case 745:
                                                        v23 := GenS2{F0: -1, F1: 321}
                                                        _ = v23
                                                    }
                                                case "a":
                                                    if (p0 || (p0 || (!true))) {
                                                        fmt.Println(v12, v12)
                                                    }
                                                case "":
                                                    _ = len(v4)
                                                }
                                                p0 = (true || (4294967295 != 2))
                                                _ = v22
                                            }
                                            var v24 int32 = 0
                                            _ = v21
                                            _ = v24
                                        }
                                        var v25 uint16 = v8
                                        _ = v25
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v26 bool = p0
                                    _ = v11
                                    _ = v12
                                    _ = v13
                                    _ = v26
                                }
                            case 9223372036854775807:
                                p0 = (0 != (^(fi6)))
                            case 128:
                                p0 = p0
                            }
                            _ = v8
                        }
                    case "hello":
                        v4 = append(v4, -949)
                    default:
                        switch fi6 {
                        case 420:
                            var v27 int32 = -734
                            _ = v27
                        default:
                            p0 = true
                        }
                    }
                }
                _ = fi7
            }
            var v28 int16 = 966
            _ = fi6
            _ = v28
        }
        for fi29 := 0; fi29 < 3; fi29++ {
            for fi30 := 0; fi30 < 3; fi30++ {
                var v31 int64 = 42
                p0 = p0
                if len(v4) > 0 {
                    v4[0] = (^(fi30))
                }
                _ = fi30
                _ = v31
            }
            v32 := GenS4{F0: -1, F1: false}
            _ = fi29
            _ = v32
        }
    }
    v4 = append(v4, 270)
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 int64 = 42
    var v2 uint = 65535
    v3 := make(map[uint32]uint32)
    var v33 int8 = 42
    v34 := GenS2{F0: 32767, F1: 4294967295}
    fmt.Println(v1, v33, v1)
    var v35 uint16 = 2
    v33 = v33
    var v36 uint16 = v35
    var v37 float64 = (3.14159 * (0.1 + 1e10))
    var v38 float32 = 1e10
    if false {
        v3[4294967295] = 42
        _ = len(v3)
    } else {
        var v39 uint = (v2 | v2)
        for fi40 := 0; fi40 < 13; fi40++ {
            switch fi40 {
            case 9223372036854775807:
                if true {
                    v41 := GenS4{F0: 32767, F1: false}
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 17; fi42++ {
                        fmt.Println(fi40, v37)
                        _ = fi42
                    }
                    var v43 int32 = 2147483647
                    _ = v43
                }
            case 128:
                v44 := []int{int(v39), int(v33), (-(fi40))}
                _ = v44
            case 32767:
                var v45 string = ""
                _ = v45
            }
            if false {
                switch "a" {
                case "":
                    var v46 float64 = (float64(v38) * v37)
                    _ = v46
                case "foo bar":
                    switch -2147483648 {
                    case 42:
                        v47 := GenS5{F0: uint64(v39)}
                        _ = v47
                    case 127:
                        switch fi40 {
                        case 822:
                            v48 := GenS6{F0: (v2 % (((^(v39))) | 1))}
                            _ = v48
                        case 4294967295:
                            v35 = v36
                        default:
                            if (-2147483648 != 0) {
                                v36 += v35
                                switch (("a" + "a") + ("hello" + "")) {
                                case "hello":
                                    for fi49 := 0; fi49 < 9; fi49++ {
                                        v37 = 1e10
                                        _ = fi49
                                    }
                                }
                                if (231 >= 1) {
                                    v50 := GenS7{F0: (fi40 << ((v2) & 63)), F1: int(v33)}
                                    v2 = uint(v36)
                                    _ = v50
                                } else {
                                    v51 := []int{fi40, (fi40 + (^(fi40))), 4294967295}
                                    var v52 float32 = float32((v37 / 3.14159))
                                    fmt.Println(v1, v52, v35)
                                    _ = v51
                                    _ = v52
                                }
                            } else {
                                if true {
                                    v53 := GenS4{F0: 32767, F1: (330 != 4294967295)}
                                    v54 := GenS7{F0: fi40, F1: fi40}
                                    var v55 int32 = 863
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                } else {
                                    v56 := []int{}
                                    _ = v56
                                }
                                for fi57 := 0; fi57 < 3; fi57++ {
                                    v37 = v37
                                    var v58 bool = false
                                    v58 = v58
                                    _ = fi57
                                    _ = v58
                                }
                                v59 := GenS4{F0: -1, F1: (!false)}
                                _ = v59
                            }
                        }
                    case 702:
                        v60 := []int{(fi40 << (((v39 / ((v2) | 1))) & 63)), int(v1)}
                        _ = v60
                    default:
                        v36 += v36
                    }
                case "hello":
                    var v61 bool = true
                    _ = v61
                default:
                    var v62 int64 = 42
                    _ = v62
                }
            } else {
                v33 += (^(v33))
                if ("hello" <= ("a" + "\n\t")) {
                    _ = len(v3)
                    var v63 int = 88
                    _ = v63
                } else {
                    for fi64 := 0; fi64 < 7; fi64++ {
                        v65 := GenS8{F0: "foo bar", F1: false}
                        _ = fi64
                        _ = v65
                    }
                }
                if false {
                    var v66 bool = (255 == 42)
                    _ = v66
                } else {
                    v37 = 3.14159
                    for fi67 := 0; fi67 < 2; fi67++ {
                        var v68 float32 = v38
                        v69 := []int{-9223372036854775808, fi67}
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                    if (fi40 > fi40) {
                        if true {
                            v39 += (v39 + (v39 ^ v2))
                        } else {
                            fmt.Println(v36)
                            v70 := GenS5{F0: 398}
                            v39 += v39
                            _ = v70
                        }
                        fmt.Println(v38)
                    } else {
                        fmt.Println(v36, v1)
                        v2 = v2
                    }
                }
            }
            _ = fi40
        }
        _ = v39
    }
    if true {
        var v71 bool = (v35 <= (v35 - (v35 | 42)))
        _ = v71
    } else {
        var v72 int32 = -2147483648
        var v73 uint8 = 49
        switch "\n\t" {
        case "foo bar":
            if false {
                var v74 int8 = int8(v36)
                if ((v35 >> (((v2 & (v2 * v2))) & 63)) > v36) {
                    var v75 int16 = 32767
                    _ = v75
                } else {
                    var v76 uint = (v2 / ((uint(v36)) | 1))
                    var v77 uint = (v76 + 1)
                    var v78 uint16 = (v35 + v36)
                    _ = v76
                    _ = v77
                    _ = v78
                }
                var v79 bool = false
                _ = v74
                _ = v79
            } else {
                if true {
                    v72 += (v72 / ((-2147483648) | 1))
                    var v80 uint8 = uint8(v2)
                    v1 = (v1 << ((v2) & 63))
                    _ = v80
                } else {
                    var v81 string = ("" + "hello")
                    _ = v81
                }
                for fi82 := 0; fi82 < 13; fi82++ {
                    for fi83 := 0; fi83 < 17; fi83++ {
                        v84 := v3[2147483648]
                        v85 := GenS2{F0: 1, F1: (v84 & (^(v84)))}
                        var v86 uint8 = v73
                        _ = fi83
                        _ = v84
                        _ = v85
                        _ = v86
                    }
                    v87 := v3[65535]
                    _ = fi82
                    _ = v87
                }
                delete(v3, 2147483647)
            }
        case "":
            v35 += 0
        }
        _ = v72
        _ = v73
    }
    var v88 uint8 = uint8(v35)
    if true {
        v88 = (^(v88))
        v89 := GenS9{F0: -32768, F1: 397, F2: 1}
        if (v38 <= (v38 + 0.0)) {
            v90 := []int{866, 128}
            if (42 >= 4294967295) {
                v36 = v35
                var v91 uint = 334
                _ = v91
            } else {
                switch (("hello" + "") + "") {
                case "":
                    v92 := GenS10{F0: -32768, F1: (-(v1))}
                    _ = v92
                default:
                    v93 := GenS4{F0: 0, F1: ((v2 * v2) >= (v2 << ((v2) & 63)))}
                    _ = v93
                }
            }
            _ = v90
        }
        _ = v89
    }
    v36 = v35
    switch "\n\t" {
    case "\n\t":
        switch 32768 {
        case -1:
            v94 := GenS11{F0: 2147483648}
            _ = v94
        }
    }
    v95 := []int{-1, 90, int(v33), 9223372036854775807}
    var v96 float64 = ((1e10 - 3.14159) * v37)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v88:", v88)
    fmt.Println("slicelen v95:", len(v95), v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v88
    _ = v95
    _ = v96
}
