package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 string }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 int16; F2 int }
type GenS10 struct { F0 int }
func genFn0(p0 int64, p1 uint64, p2 uint32) uint16 {
    v7 := []int{374, 32767, 32767}
    if len(v7) > 0 {
        v7[0] = -214
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 959
    var v3 int16 = 42
    var v4 uint32 = 2147483647
    var v5 int8 = -1
    var v6 uint32 = v4
    v1[2147483648] = (false || true)
    v4 = 2147483648
    switch 1 {
    case 1:
        var v8 uint = 2
        _ = v8
    default:
        v6 = (v4 + (v4 << ((2) & 63)))
    }
    v9 := v1[(v6 ^ v4)]
    fmt.Println(v2, v4)
    v10 := GenS1{F0: "", F1: 2147483647}
    v6 = (v4 >> ((65535) & 63))
    var v11 float64 = ((float64(0.1) + 3.14159) * -2.25)
    switch -198 {
    case 32767:
        var v12 uint = 65535
        _ = v12
    case 128:
        v9 = (!v9)
    case 1:
        if v9 {
            if (((v4 & v6) < 40) && v9) {
                var v13 float32 = 1.5
                v2 -= v2
                _ = v13
            } else {
                v14 := GenS1{F0: "foo bar", F1: 563}
                if true {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v16 := GenS1{F0: ("hello" + ("a" + "hello")), F1: 42}
                        v17 := GenS2{F0: fi15, F1: 2, F2: ("" + (("a" + ("a" + "a")) + "a"))}
                        switch 1 {
                        case -1:
                            for fi18 := 0; fi18 < 4; fi18++ {
                                fmt.Println(v3, fi18, v9)
                                var v19 int8 = v5
                                var v20 int16 = 0
                                var v21 int = -2147483648
                                _ = fi18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                        case 128:
                            var v22 uint64 = v2
                            _ = v22
                        case -478:
                            v23 := GenS1{F0: "foo bar", F1: -1}
                            _ = v23
                        default:
                            _ = len(v1)
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
                _ = v14
            }
            if v9 {
                v11 = ((v11 * (3.14159 - v11)) + (v11 * 0.0))
                v24 := []int{int(v2), -9223372036854775808}
                _ = v24
            } else {
                for fi25 := 0; fi25 < 7; fi25++ {
                    if v9 {
                        var v26 uint8 = 127
                        if false {
                            var v27 float64 = 1e10
                            _ = v27
                        } else {
                            var v28 uint16 = uint16(v26)
                            for fi29 := 0; fi29 < 15; fi29++ {
                                _ = len(v1)
                                v4 += 938
                                if (!("" != "a")) {
                                    switch (fi25 % ((9223372036854775807) | 1)) {
                                    case -2147483648:
                                        var v30 int64 = -172
                                        _ = v30
                                    default:
                                        v31 := GenS2{F0: -473, F1: (^(v26)), F2: "\n\t"}
                                        _ = v31
                                    }
                                    var v32 uint16 = 42
                                    switch 9223372036854775807 {
                                    case 32767:
                                        v6 = (v4 + 267)
                                    }
                                    _ = v32
                                } else {
                                    v33 := GenS2{F0: (fi29 | (fi29 + (fi29 >> ((434) & 63)))), F1: v26, F2: "a"}
                                    v34 := []int{(-(fi29)), 128, int(v3)}
                                    fmt.Println(v2)
                                    _ = v33
                                    _ = v34
                                }
                                _ = fi29
                            }
                            _ = v28
                        }
                        _ = v26
                    }
                    _ = fi25
                }
                if (!true) {
                    if false {
                        for fi35 := 0; fi35 < 18; fi35++ {
                            v5 += v5
                            _ = fi35
                        }
                        switch "foo bar" {
                        case "a":
                            for fi36 := 0; fi36 < 12; fi36++ {
                                var v37 uint = 42
                                if ((v2 / (((v2 >> ((1) & 63))) | 1)) <= v2) {
                                    if v9 {
                                        delete(v1, (v6 * 4294967295))
                                        delete(v1, (v4 | v6))
                                    } else {
                                        fmt.Println(v9, fi36, fi36)
                                        if false {
                                            v38 := []int{fi36, 0, (fi36 % ((int(v5)) | 1))}
                                            v38 = append(v38, (fi36 % (((fi36 >> ((18446744073709551615) & 63))) | 1)))
                                            _ = v38
                                        }
                                    }
                                }
                                for fi39 := 0; fi39 < 18; fi39++ {
                                    v5 = 67
                                    _ = fi39
                                }
                                _ = fi36
                                _ = v37
                            }
                        default:
                            fmt.Println(v6, v9)
                        }
                    } else {
                        var v40 uint32 = (v6 >> ((uint(v4)) & 63))
                        v6 += (^(v6))
                        _ = v40
                    }
                    fmt.Println(v9, v5)
                } else {
                    v41 := GenS3{F0: v2, F1: -9223372036854775808}
                    _ = v41
                }
                var v42 uint32 = v4
                _ = v42
            }
            if false {
                for fi43 := 0; fi43 < 6; fi43++ {
                    v44 := GenS4{F0: 42, F1: v9}
                    _ = len(v1)
                    _ = fi43
                    _ = v44
                }
                var v45 float32 = float32(((float64(3.14159) / (v11 * (v11 + float64((0.1 - 1e10))))) * v11))
                if false {
                    if v9 {
                        var v46 int = int(v6)
                        v3 += 32767
                        _ = v46
                    }
                    v45 -= v45
                    v47 := GenS2{F0: 9223372036854775807, F1: 127, F2: ("hello" + "\n\t")}
                    _ = v47
                } else {
                    switch 660 {
                    case 0:
                        v1[474] = (v45 == v45)
                    case 2147483647:
                        v48 := GenS1{F0: "hello", F1: int32(v2)}
                        _ = v48
                    }
                    switch 42 {
                    case 986:
                        switch -106 {
                        case -520:
                            if v9 {
                                _ = len(v1)
                                var v49 int32 = -2147483648
                                v50 := GenS5{F0: (("hello" + "hello") + (("\n\t" + "") + ("foo bar" + ("a" + "")))), F1: 147}
                                _ = v49
                                _ = v50
                            } else {
                                var v51 uint32 = v6
                                var v52 int = 0
                                if ((v4 | (v6 | v4)) == (v51 >> ((0) & 63))) {
                                    v6 = (v51 / ((2147483647) | 1))
                                    v53 := GenS6{F0: 42, F1: 42}
                                    fmt.Println(v2, v51, v45)
                                    _ = v53
                                } else {
                                    var v54 string = "hello"
                                    var v55 int16 = -581
                                    _ = v54
                                    _ = v55
                                }
                                _ = v51
                                _ = v52
                            }
                        case -9223372036854775808:
                            if ((v3 / ((0) | 1)) != (v3 ^ (v3 & v3))) {
                                v6 += 2
                                var v56 float64 = v11
                                var v57 int8 = v5
                                _ = v56
                                _ = v57
                            }
                        case 967:
                            switch ((("a" + "foo bar") + (("\n\t" + "a") + ((("hello" + "") + "a") + "foo bar"))) + "hello") {
                            case "foo bar":
                                switch int(v4) {
                                case 127:
                                    var v58 uint16 = genFn0(1, (v2 - v2), v4)
                                    _ = v58
                                case 0:
                                    v2 -= 42
                                default:
                                    var v59 uint64 = v2
                                    _ = v59
                                }
                            default:
                                v4 += (v6 - 2147483647)
                            }
                        }
                    case -288:
                        var v60 int16 = v3
                        _ = v60
                    case 4294967295:
                        if v9 {
                            var v61 bool = v9
                            for fi62 := 0; fi62 < 2; fi62++ {
                                fmt.Println(v9)
                                _ = fi62
                            }
                            _ = v61
                        }
                    default:
                        var v63 int16 = int16(v2)
                        _ = v63
                    }
                }
                _ = v45
            }
        }
    }
    if v9 {
        fmt.Println(v6, v6)
    }
    fmt.Println(v4, v9)
    var v64 uint64 = (v2 - (v2 & uint64(v6)))
    var v65 uint = 255
    v64 -= (v64 % ((4294967295) | 1))
    switch -56 {
    case -397:
        _ = len(v1)
    case 128:
        v66 := GenS3{F0: v64, F1: -220}
        _ = v66
    case -312:
        if v9 {
            var v67 float64 = (v11 / 0.1)
            for fi68 := 0; fi68 < 14; fi68++ {
                for fi69 := 0; fi69 < 7; fi69++ {
                    v70 := GenS7{F0: 2147483647, F1: (v2 | v64)}
                    _ = fi69
                    _ = v70
                }
                for fi71 := 0; fi71 < 14; fi71++ {
                    if v9 {
                        switch "" {
                        case "\n\t":
                            var v72 bool = (((!true) && v9) && v9)
                            _ = v72
                        }
                        v73 := GenS7{F0: 2147483647, F1: (v64 & 18446744073709551615)}
                        var v74 int32 = 42
                        _ = v73
                        _ = v74
                    }
                    if v9 {
                        v75 := GenS8{F0: "\n\t"}
                        v76 := GenS9{F0: int8(v2), F1: v3, F2: (fi68 / (((fi71 ^ fi71)) | 1))}
                        var v77 int = 127
                        _ = v75
                        _ = v76
                        _ = v77
                    } else {
                        v9 = ((v67 / v11) > v11)
                        v78 := []int{fi71}
                        v79 := GenS1{F0: "a", F1: 1}
                        _ = v78
                        _ = v79
                    }
                    fmt.Println(fi68)
                    _ = fi71
                }
                _ = fi68
            }
            fmt.Println(v67)
            _ = v67
        } else {
            v2 = 841
            for fi80 := 0; fi80 < 13; fi80++ {
                if (148 >= (v6 & v6)) {
                    v11 = (v11 - -2.25)
                    fmt.Println(v5, v4)
                    for fi81 := 0; fi81 < 8; fi81++ {
                        if ((v3 % (((v3 | v3)) | 1)) <= 1000) {
                            v82 := GenS8{F0: ("a" + "a")}
                            _ = v82
                        }
                        v83 := GenS2{F0: (fi80 - (-(fi81))), F1: 2, F2: "\n\t"}
                        var v84 uint64 = 740
                        _ = fi81
                        _ = v83
                        _ = v84
                    }
                } else {
                    if (v64 == 2) {
                        v85 := []int{}
                        _ = v85
                    } else {
                        switch int(v5) {
                        case 1:
                            v6 -= v6
                        case 492:
                            v86 := []int{fi80, fi80, fi80, (fi80 - fi80)}
                            _ = v86
                        default:
                            for fi87 := 0; fi87 < 19; fi87++ {
                                var v88 uint64 = (^(v2))
                                if ((65535 == 32767) || (!(v9 && (v11 <= v11)))) {
                                    fmt.Println(v65, v4)
                                    v88 -= (v88 << (((v65 / ((v65) | 1))) & 63))
                                    var v89 uint = (v65 - 18446744073709551615)
                                    _ = v89
                                }
                                v2 += v64
                                v9 = v9
                                fmt.Println(v5, v5, v2)
                                v1[923] = v9
                                _ = fi87
                                _ = v88
                            }
                        }
                        if v9 {
                            for fi90 := 0; fi90 < 14; fi90++ {
                                var v91 string = "hello"
                                var v92 float32 = (-2.25 * 1e10)
                                _ = fi90
                                _ = v91
                                _ = v92
                            }
                            v4 = 1
                            v93 := GenS10{F0: fi80}
                            _ = v93
                        }
                        var v94 uint8 = uint8(v65)
                        _ = v94
                    }
                }
                var v95 int32 = -2147483648
                _ = fi80
                _ = v95
            }
            v96 := GenS10{F0: 127}
            _ = v96
        }
    }
    v6 = (v4 + (v4 << ((v65) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v9
}
