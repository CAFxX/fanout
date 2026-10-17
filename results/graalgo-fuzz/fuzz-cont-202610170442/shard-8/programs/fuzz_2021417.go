package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint32; F1 int8; F2 uint32 }
func genFn0(p0 uint) int32 {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := []int{fi4}
        var v6 uint = 2
        if (-27 == -1) {
            v6 += v6
            fmt.Println(fi4, fi4)
        } else {
            var v7 int8 = 100
            p0 = v6
            _ = v7
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    p0 += (^(p0))
    if false {
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(fi8, p0)
            _ = fi8
        }
        p0 = (p0 | (p0 / ((p0) | 1)))
        v9 := []int{1, 0, -2147483648, -1}
        _ = v9
    }
    _ = p0
    return -1
}

func main() {
    var v1 int8 = -1
    var v2 int16 = 32767
    var v3 bool = (false && (0 >= 255))
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := []int{fi10}
        var v12 float32 = 0.0
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int = -399
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := []int{fi14}
        switch "a" {
        case "":
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 uint32 = 2147483647
                var v18 uint8 = 2
                _ = fi16
                _ = v17
                _ = v18
            }
        case "\n\t":
            v15 = append(v15, (fi14 << ((4294967295) & 63)))
        case "foo bar":
            var v19 uint8 = 40
            _ = v19
        }
        var v20 uint16 = 255
        _ = fi14
        _ = v15
        _ = v20
    }
    v13 = (v13 * v13)
    v21 := []int{v13, v13}
    var v22 uint8 = 127
    for fi23 := 0; fi23 < 4; fi23++ {
        v3 = v3
        for fi24 := 0; fi24 < 8; fi24++ {
            if true {
                var v25 int8 = v1
                if (!v3) {
                    fmt.Println(v1, v13, fi24)
                    var v26 int32 = 2147483647
                    _ = v26
                }
                _ = v25
            }
            switch "hello" {
            case "foo bar":
                if ("\n\t" >= "hello") {
                    var v27 bool = false
                    switch "" {
                    case "a":
                        v21 = append(v21, v13)
                    case "hello":
                        fmt.Println(fi24, v13, fi23)
                    case "":
                        v27 = v3
                    default:
                        if v27 {
                            v2 = (^(v2))
                        }
                    }
                    _ = v27
                } else {
                    var v28 uint32 = 547
                    var v29 int32 = 2147483647
                    _ = v28
                    _ = v29
                }
                fmt.Println(fi24, v13)
            case "\n\t":
                var v30 float32 = (0.0 - (((1.5 + 0.0) - float32(float64((float32(0.1) / (1.5 * 0.1))))) / 1.0))
                v2 = (v2 << ((18446744073709551615) & 63))
                v21 = append(v21, fi23)
                if (false && v3) {
                    if (((("foo bar" + "hello") + "foo bar") + "foo bar") > "foo bar") {
                        if v3 {
                            var v31 uint32 = 4294967295
                            var v32 uint32 = (v31 & 150)
                            var v33 int32 = 0
                            if len(v21) > 0 {
                                v21[0] = -279
                            }
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            _ = v31
                            _ = v32
                            _ = v33
                        } else {
                            if (false && (v30 == v30)) {
                                var v34 int8 = v1
                                if (!(!v3)) {
                                    var v35 float32 = 0.1
                                    _ = v35
                                } else {
                                    v36 := GenS1{F0: -935, F1: 128}
                                    fmt.Println(fi23)
                                    v37 := GenS1{F0: 526, F1: fi23}
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    fmt.Println(fi24)
                                    _ = v36
                                    _ = v37
                                }
                                v2 = int16(v13)
                                _ = v34
                            } else {
                                v38 := GenS1{F0: -2147483648, F1: (-(v13))}
                                _ = len(v21)
                                switch "" {
                                case "hello":
                                    v22 -= v22
                                case "\n\t":
                                    var v39 float32 = (v30 + (v30 / v30))
                                    _ = v39
                                case "foo bar":
                                    if v3 {
                                        if v3 {
                                            switch (fi23 & (fi23 / (((v13 % ((-2147483648) | 1))) | 1))) {
                                            case 9223372036854775807:
                                                switch "" {
                                                case "hello":
                                                    if (!(((9223372036854775807 <= 42) || v3) || true)) {
                                                        var v40 int64 = 1
                                                        v22 -= v22
                                                        v41 := GenS2{F0: v22, F1: (v1 <= int8(v2))}
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                default:
                                                    var v42 string = "\n\t"
                                                    v43 := GenS1{F0: 42, F1: 404}
                                                    _ = v42
                                                    _ = v43
                                                }
                                            default:
                                                var v44 uint32 = 2147483647
                                                _ = v44
                                            }
                                            v45 := GenS2{F0: v22, F1: ("" != "")}
                                            _ = v45
                                        }
                                    } else {
                                        v1 -= 0
                                    }
                                default:
                                    if v3 {
                                        switch ("a" + "hello") {
                                        case "foo bar":
                                            var v46 uint8 = (v22 + v22)
                                            _ = v46
                                        }
                                        if (!v3) {
                                            var v47 uint8 = v22
                                            var v48 int = (fi24 & (^(fi23)))
                                            if len(v21) > 0 {
                                                v21[0] = fi23
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                        v22 -= 2
                                    }
                                }
                                _ = v38
                            }
                            v13 -= fi23
                        }
                        v13 = (fi23 & (v13 + v13))
                    } else {
                        v22 = (^(v22))
                        if v3 {
                            if (!v3) {
                                v21 = append(v21, 288)
                                v2 += (v2 - (v2 << ((uint(v22)) & 63)))
                                if (1 == 4294967295) {
                                    if len(v21) > 0 {
                                        v21[0] = 42
                                    }
                                }
                            } else {
                                var v49 uint8 = v22
                                v2 += v2
                                _ = v49
                            }
                            v13 += fi23
                        } else {
                            switch "a" {
                            case "a":
                                switch "" {
                                case "a":
                                    v2 += -1
                                default:
                                    v50 := GenS2{F0: v22, F1: v3}
                                    _ = v50
                                }
                            }
                            var v51 int = (fi23 - 32768)
                            v1 = v1
                            _ = v51
                        }
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                    }
                    switch "a" {
                    case "foo bar":
                        v13 = fi23
                    case "a":
                        var v52 int32 = -821
                        _ = v52
                    case "\n\t":
                        switch "foo bar" {
                        case "a":
                            v30 = (1e10 * 1e10)
                        case "foo bar":
                            if (v3 && (v3 || false)) {
                                if ((!(-21 == (v1 << ((42) & 63)))) || (((v22 == (v22 + v22)) || (!(391 > 42))) && true)) {
                                    var v53 int64 = 42
                                    _ = v53
                                } else {
                                    v3 = false
                                    var v54 uint32 = 85
                                    _ = v54
                                }
                            } else {
                                var v55 string = ""
                                if (!(2 != 18446744073709551615)) {
                                    switch fi24 {
                                    case 127:
                                        var v56 string = (v55 + "hello")
                                        _ = v56
                                    case 2147483647:
                                        v57 := GenS3{F0: 310, F1: -98, F2: 0}
                                        _ = v57
                                    }
                                    var v58 float32 = (v30 - ((0.0 / 1.5) / v30))
                                    v21 = append(v21, fi24)
                                    _ = v58
                                }
                                v2 += v2
                                _ = v55
                            }
                        case "":
                            var v59 uint32 = 167
                            _ = v59
                        default:
                            var v60 uint64 = 441
                            _ = v60
                        }
                    }
                    v22 = (v22 / (((v22 % ((v22) | 1))) | 1))
                    v30 = v30
                } else {
                    v2 = v2
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                }
                _ = v30
            default:
                var v61 int = (v13 * (fi23 - (fi23 ^ (-(fi23)))))
                fmt.Println(v3, fi23, v22)
                _ = v61
            }
            var v62 uint8 = (^(v22))
            var v63 uint16 = 2
            _ = fi24
            _ = v62
            _ = v63
        }
        var v64 uint8 = (^(v22))
        _ = fi23
        _ = v64
    }
    var v65 uint16 = 42
    var v66 uint = 4294967295
    if false {
        for fi67 := 0; fi67 < 17; fi67++ {
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            for fi68 := 0; fi68 < 3; fi68++ {
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
                var v69 uint32 = 2
                var v70 uint8 = v22
                _ = fi68
                _ = v69
                _ = v70
            }
            fmt.Println(v2, v66, v22)
            _ = fi67
        }
        if v3 {
            for fi71 := 0; fi71 < 4; fi71++ {
                for fi72 := 0; fi72 < 12; fi72++ {
                    fmt.Println(v65, v65, fi72)
                    var v73 bool = v3
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    switch v13 {
                    case -977:
                        var v74 int = fi71
                        _ = v74
                    case -9223372036854775808:
                        v21 = append(v21, (v13 >> ((uint(v65)) & 63)))
                    default:
                        _ = len(v21)
                    }
                    _ = fi72
                    _ = v73
                }
                _ = fi71
            }
            fmt.Println(v13)
        }
    } else {
        _ = len(v21)
        if v3 {
            var v75 int64 = -153
            var v76 float64 = (0.1 / 1.0)
            _ = v75
            _ = v76
        } else {
            _ = len(v21)
        }
        var v77 int32 = 639
        _ = v77
    }
    var v78 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v65
    _ = v66
    _ = v78
}
