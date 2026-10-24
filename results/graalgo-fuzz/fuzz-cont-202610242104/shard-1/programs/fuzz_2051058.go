package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 bool }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int8; F2 int64 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint16) int64 {
    if ((-9223372036854775808 >= -1) || false) {
        for fi4 := 0; fi4 < 6; fi4++ {
            switch fi4 {
            case -32:
                var v5 uint64 = 4294967295
                _ = v5
            case 9223372036854775807:
                var v6 uint64 = 246
                _ = v6
            case 2147483647:
                v7 := GenS1{F0: 1, F1: 42}
                _ = v7
            }
            _ = fi4
        }
        v8 := GenS1{F0: 992, F1: 0}
        _ = v8
    }
    var v9 int8 = 81
    for fi10 := 0; fi10 < 7; fi10++ {
        p0 = (p0 & 0)
        v9 = (v9 - (v9 / ((v9) | 1)))
        if true {
            v9 += v9
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 uint = 385
                v13 := GenS1{F0: 0, F1: 179}
                v14 := GenS1{F0: 18446744073709551615, F1: 2}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        } else {
            switch (fi10 << ((1) & 63)) {
            case 2147483647:
                v9 -= 0
            }
        }
        _ = fi10
    }
    switch "foo bar" {
    case "hello":
        if true {
            var v15 uint64 = 1
            var v16 uint8 = 67
            v17 := []int{127}
            _ = v15
            _ = v16
            _ = v17
        }
    case "foo bar":
        var v18 float64 = ((1e10 - 1.5) * (float64(-2.25) * 3.14159))
        _ = v18
    case "a":
        v9 = v9
    }
    _ = p0
    _ = v9
    return 236
}

func main() {
    var v1 uint64 = 2
    var v2 uint = 1
    var v3 string = (("\n\t" + "\n\t") + "")
    for fi19 := 0; fi19 < 10; fi19++ {
        if false {
            v1 -= 42
            switch (fi19 & fi19) {
            case 639:
                v20 := []int{fi19, (fi19 * (fi19 / ((-452) | 1))), -1, (-(fi19))}
                _ = v20
            case 127:
                if true {
                    for fi21 := 0; fi21 < 7; fi21++ {
                        v22 := GenS2{F0: 1, F1: 42, F2: true}
                        _ = fi21
                        _ = v22
                    }
                } else {
                    switch 32768 {
                    case -755:
                        var v23 float64 = float64(1e10)
                        _ = v23
                    case 127:
                        v24 := []int{fi19, (fi19 + (fi19 % ((32767) | 1))), (fi19 ^ 2147483647), 32767}
                        _ = v24
                    case 42:
                        var v25 float64 = float64(0.0)
                        _ = v25
                    default:
                        v3 = v3
                    }
                    v26 := []int{fi19}
                    var v27 uint16 = 902
                    _ = v26
                    _ = v27
                }
            case 739:
                if false {
                    v28 := []int{}
                    _ = v28
                } else {
                    v29 := []int{(fi19 / ((fi19) | 1)), 128}
                    for fi30 := 0; fi30 < 12; fi30++ {
                        if ((v3 + ("a" + v3)) != ("hello" + ("a" + v3))) {
                            v3 = v3
                            v29 = append(v29, (fi30 - 128))
                            v3 = ""
                        }
                        fmt.Println(v3, v1)
                        if false {
                            var v31 int16 = 1000
                            _ = v31
                        } else {
                            var v32 uint8 = 103
                            _ = len(v29)
                            var v33 int64 = -9223372036854775808
                            if (!true) {
                                fmt.Println(fi30, v33)
                            }
                            _ = v32
                            _ = v33
                        }
                        _ = fi30
                    }
                    _ = v29
                }
            }
            if false {
                switch (v3 + "\n\t") {
                case "":
                    v3 = ""
                }
            }
        } else {
            if false {
                for fi34 := 0; fi34 < 19; fi34++ {
                    var v35 uint8 = 250
                    v1 -= (v1 - (v1 >> ((v2) & 63)))
                    _ = fi34
                    _ = v35
                }
                switch fi19 {
                case 2147483647:
                    var v36 uint = (v2 % ((v2) | 1))
                    _ = v36
                case -1:
                    fmt.Println(fi19)
                case 9223372036854775807:
                    v1 = v1
                default:
                    for fi37 := 0; fi37 < 5; fi37++ {
                        v1 = v1
                        _ = fi37
                    }
                }
                if (2 < 339) {
                    fmt.Println(v1, v2, v2)
                    for fi38 := 0; fi38 < 2; fi38++ {
                        var v39 uint64 = v1
                        _ = fi38
                        _ = v39
                    }
                }
            } else {
                var v40 uint64 = 250
                if (genFn0(42) < genFn0(255)) {
                    var v41 int32 = 42
                    v42 := GenS1{F0: v40, F1: 2}
                    var v43 int8 = 100
                    _ = v41
                    _ = v42
                    _ = v43
                } else {
                    var v44 uint16 = 0
                    v45 := GenS3{F0: 821, F1: ("a" + (v3 + ("\n\t" + v3))), F2: v44}
                    _ = v44
                    _ = v45
                }
                fmt.Println(v40, v40)
                _ = v40
            }
            switch fi19 {
            case 32767:
                v1 -= v1
            case -2147483648:
                if false {
                    if true {
                        fmt.Println(fi19, v1)
                        var v46 uint32 = 661
                        _ = v46
                    } else {
                        if true {
                            switch ("hello" + "foo bar") {
                            case "\n\t":
                                if (!true) {
                                    v47 := GenS2{F0: -1, F1: 32767, F2: (false || ((0.0 + 3.14159) >= 1e10))}
                                    _ = v47
                                } else {
                                    v48 := GenS4{F0: v2, F1: 608}
                                    v3 = "foo bar"
                                    for fi49 := 0; fi49 < 16; fi49++ {
                                        v50 := []int{fi19, -111}
                                        _ = fi49
                                        _ = v50
                                    }
                                    _ = v48
                                }
                            case "foo bar":
                                v1 = v1
                            case "":
                                for fi51 := 0; fi51 < 17; fi51++ {
                                    if (v2 >= v2) {
                                        if (-127 > 52) {
                                            var v52 float32 = -2.25
                                            var v53 int64 = 42
                                            switch v3 {
                                            case "hello":
                                                var v54 uint = (v2 >> ((427) & 63))
                                                _ = v54
                                            case "\n\t":
                                                v55 := GenS4{F0: 255, F1: 2}
                                                _ = v55
                                            }
                                            v56 := GenS3{F0: genFn0(727), F1: "foo bar", F2: 2}
                                            _ = v52
                                            _ = v53
                                            _ = v56
                                        } else {
                                            if false {
                                                fmt.Println(v3, v1)
                                            }
                                            v57 := GenS5{F0: 42, F1: 42, F2: -9223372036854775808}
                                            v58 := GenS6{F0: v3, F1: v3}
                                            _ = v57
                                            _ = v58
                                        }
                                        switch fi19 {
                                        case 1:
                                            v59 := GenS3{F0: 1, F1: "foo bar", F2: 1}
                                            _ = v59
                                        case 4294967295:
                                            if (128 >= fi51) {
                                                v60 := GenS7{F0: 42}
                                                _ = v60
                                            }
                                        case -553:
                                            if (1e10 >= -2.25) {
                                                var v61 int64 = 1
                                                v62 := GenS5{F0: 2, F1: int8(v61), F2: (v61 | v61)}
                                                _ = v61
                                                _ = v62
                                            } else {
                                                v2 += v2
                                                var v63 float32 = 1.5
                                                _ = v63
                                            }
                                        }
                                        v64 := GenS4{F0: (^(v2)), F1: 754}
                                        _ = v64
                                    }
                                    _ = fi51
                                }
                            }
                            switch v3 {
                            case "\n\t":
                                var v65 uint = 4294967295
                                _ = v65
                            }
                        }
                    }
                }
            }
            for fi66 := 0; fi66 < 2; fi66++ {
                fmt.Println(v1, v3, v1)
                _ = fi66
            }
        }
        _ = fi19
    }
    v67 := GenS6{F0: "foo bar", F1: v3}
    var v68 int32 = -2147483648
    for fi69 := 0; fi69 < 1; fi69++ {
        for fi70 := 0; fi70 < 17; fi70++ {
            switch "a" {
            case "\n\t":
                var v71 uint64 = v1
                _ = v71
            case "a":
                if (true && false) {
                    fmt.Println(v1, v3, v1)
                    if false {
                        v72 := GenS8{F0: true}
                        _ = v72
                    }
                    v73 := []int{}
                    _ = v73
                }
            default:
                var v74 int64 = 0
                _ = v74
            }
            fmt.Println(v2)
            v75 := GenS8{F0: true}
            v1 += (^(v1))
            _ = fi70
            _ = v75
        }
        _ = fi69
    }
    fmt.Println(v1, v2)
    v76 := GenS2{F0: 0, F1: -756, F2: ((-2.25 > -2.25) && true)}
    if false {
        if true {
            var v77 uint8 = 127
            for fi78 := 0; fi78 < 6; fi78++ {
                switch v3 {
                case "\n\t":
                    switch fi78 {
                    case 127:
                        v79 := GenS9{F0: 4294967295, F1: v1}
                        _ = v79
                    case -970:
                        if false {
                            for fi80 := 0; fi80 < 2; fi80++ {
                                var v81 int8 = 23
                                var v82 uint8 = (v77 & (v77 >> (((v2 / ((v2) | 1))) & 63)))
                                v82 += v82
                                _ = fi80
                                _ = v81
                                _ = v82
                            }
                            v3 = (v3 + ("a" + "a"))
                            if (32767 == 2) {
                                for fi83 := 0; fi83 < 4; fi83++ {
                                    v84 := GenS8{F0: (!(v3 <= ((v3 + v3) + v3)))}
                                    v3 = v3
                                    _ = fi83
                                    _ = v84
                                }
                                v85 := []int{}
                                _ = v85
                            }
                        }
                    case 32768:
                        v86 := []int{fi78, (fi78 % ((fi78) | 1))}
                        _ = v86
                    }
                default:
                    v2 += v2
                }
                var v87 uint32 = 383
                v88 := GenS5{F0: 795, F1: 10, F2: 9223372036854775807}
                _ = fi78
                _ = v87
                _ = v88
            }
            _ = v77
        }
        v89 := []int{-9223372036854775808, 32768}
        _ = v89
    }
    fmt.Println(v3)
    v3 = v3
    if false {
        for fi90 := 0; fi90 < 3; fi90++ {
            var v91 float32 = 1.5
            v92 := GenS1{F0: (v1 + v1), F1: (^(v1))}
            _ = fi90
            _ = v91
            _ = v92
        }
        v3 = "foo bar"
    }
    var v93 uint32 = 65535
    var v94 uint32 = 42
    v1 = v1
    var v95 int32 = (^(v68))
    v96 := []int{-721}
    fmt.Println(v1)
    var v97 int64 = 1
    for fi98 := 0; fi98 < 2; fi98++ {
        for fi99 := 0; fi99 < 1; fi99++ {
            switch fi98 {
            case 1:
                var v100 float32 = -2.25
                _ = v100
            case 32767:
                switch v3 {
                case "\n\t":
                    var v101 int64 = -1
                    _ = v101
                }
            case 32768:
                var v102 bool = true
                _ = v102
            }
            _ = fi99
        }
        var v103 float64 = 3.14159
        v2 += v2
        _ = fi98
        _ = v103
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v68:", v68)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    fmt.Println("slicelen v96:", len(v96), v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v67
    _ = v68
    _ = v76
    _ = v93
    _ = v94
    _ = v95
    _ = v96
    _ = v97
}
