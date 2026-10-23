package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 bool = (42 >= 1)
    var v2 bool = v1
    var v3 string = "hello"
    var v4 uint16 = 0
    v5 := GenS1{F0: 1, F1: 997}
    if (18446744073709551615 < 255) {
        v3 = ("a" + v3)
        var v6 string = "foo bar"
        var v7 uint8 = 0
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: false, F1: uint(v4)}
    v9 := GenS2{F0: true, F1: 42}
    switch (v3 + v3) {
    case "\n\t":
        v2 = (!v2)
    case "foo bar":
        v10 := GenS3{F0: v4, F1: -30}
        _ = v10
    case "hello":
        v11 := GenS3{F0: 42, F1: -32768}
        _ = v11
    default:
        fmt.Println(v3, v1)
    }
    v12 := []int{-723}
    if ((18446744073709551615 > 18446744073709551615) && false) {
        var v13 int64 = -1
        _ = v13
    } else {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    v14 := GenS2{F0: (633 < 709), F1: 18446744073709551615}
    v12 = append(v12, -476)
    for fi15 := 0; fi15 < 11; fi15++ {
        if v2 {
            if ((!(!(-1 > 1))) || v1) {
                v16 := GenS1{F0: -1, F1: (fi15 + fi15)}
                v4 = v4
                _ = v16
            } else {
                if ((!((v3 + "") == (v3 + "foo bar"))) || v2) {
                    var v17 string = ((v3 + (v3 + v3)) + "")
                    _ = v17
                }
                for fi18 := 0; fi18 < 3; fi18++ {
                    switch 9223372036854775807 {
                    case 4294967295:
                        if len(v12) > 0 {
                            v12[0] = (fi15 >> ((255) & 63))
                        }
                    case 127:
                        var v19 uint32 = 4294967295
                        v1 = (!false)
                        var v20 int16 = 0
                        _ = v19
                        _ = v20
                    case 9223372036854775807:
                        var v21 float64 = ((1.5 - (0.0 + 0.1)) + 1.5)
                        switch "\n\t" {
                        case "hello":
                            if (fi15 < 2147483647) {
                                if (2147483647 == -2147483648) {
                                    if (181 < fi15) {
                                        if ((!(v3 <= "foo bar")) && (1 != 0)) {
                                            fmt.Println(fi18, v3, fi15)
                                            v22 := GenS1{F0: -696, F1: (fi15 | (fi15 + fi15))}
                                            v2 = v2
                                            v23 := GenS1{F0: 1, F1: fi18}
                                            v21 -= ((v21 * (v21 + v21)) - 3.14159)
                                            _ = v22
                                            _ = v23
                                        }
                                        v12 = append(v12, (fi15 / (((fi18 % ((fi18) | 1))) | 1)))
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        if (716 >= 4294967295) {
                                            if len(v12) > 0 {
                                                v12[0] = fi18
                                            }
                                            v2 = ((1 < 1) && (42 != 1))
                                            v24 := GenS1{F0: -59, F1: 720}
                                            _ = v24
                                        } else {
                                            _ = len(v12)
                                            v25 := GenS2{F0: (473 == 4294967295), F1: 964}
                                            v1 = (42 > uint(v4))
                                            _ = v25
                                        }
                                        var v26 uint64 = 2
                                        _ = v26
                                    }
                                    if len(v12) > 0 {
                                        v12[0] = (^(fi15))
                                    }
                                } else {
                                    switch (^(fi18)) {
                                    case -326:
                                        switch v3 {
                                        case "foo bar":
                                            if false {
                                                switch 128 {
                                                case 0:
                                                    v1 = (v1 || v1)
                                                }
                                                var v27 int16 = -32768
                                                v1 = true
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                                _ = v27
                                            } else {
                                                switch "a" {
                                                case "":
                                                    var v28 int64 = 460
                                                    _ = v28
                                                case "a":
                                                    var v29 bool = false
                                                    _ = v29
                                                case "hello":
                                                    var v30 bool = (858 == 1000)
                                                    _ = v30
                                                }
                                                if len(v12) > 0 {
                                                    v12[0] = -2147483648
                                                }
                                            }
                                        case "":
                                            fmt.Println(v21)
                                        case "a":
                                            v31 := GenS2{F0: v1, F1: 18446744073709551615}
                                            _ = v31
                                        default:
                                            var v32 float32 = float32((1e10 * v21))
                                            _ = v32
                                        }
                                    }
                                }
                                var v33 int32 = 42
                                v34 := GenS2{F0: (v1 || (2147483647 > 65535)), F1: 476}
                                v2 = v2
                                var v35 int = fi18
                                v21 = float64((0.1 - 0.0))
                                v36 := GenS1{F0: v33, F1: 9223372036854775807}
                                _ = v33
                                _ = v34
                                _ = v35
                                _ = v36
                            } else {
                                v21 += 1e10
                            }
                        case "a":
                            switch v3 {
                            case "":
                                var v37 string = ("a" + "a")
                                _ = v37
                            case "foo bar":
                                v38 := GenS1{F0: 0, F1: -309}
                                _ = v38
                            default:
                                switch ((v3 + (("\n\t" + "") + "foo bar")) + ((v3 + "") + v3)) {
                                case "\n\t":
                                    var v39 int32 = 423
                                    _ = v39
                                case "hello":
                                    v40 := GenS2{F0: true, F1: 2}
                                    if v2 {
                                        v2 = ((-59 <= -121) && (85 != int64(v4)))
                                        v1 = true
                                    } else {
                                        v4 -= (v4 | (^(v4)))
                                    }
                                    _ = v40
                                case "a":
                                    v21 += ((v21 + v21) / 1.5)
                                    v4 += v4
                                    v12 = append(v12, fi15)
                                    switch v3 {
                                    case "\n\t":
                                        v4 -= v4
                                    default:
                                        if v2 {
                                            if v2 {
                                                fmt.Println(fi18, fi18, fi18)
                                                if v2 {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v21 += 0.1
                                                    }
                                                    if v2 {
                                                        var v41 int16 = 32767
                                                        if v2 {
                                                            var v42 string = (("foo bar" + v3) + ("" + v3))
                                                            _ = v42
                                                        }
                                                        fmt.Println(v2, v21)
                                                        _ = v41
                                                    }
                                                    var v43 uint32 = uint32(v4)
                                                    _ = v43
                                                }
                                                switch (v3 + "\n\t") {
                                                case "":
                                                    fmt.Println(v2, v3)
                                                case "a":
                                                    var v44 int32 = -2147483648
                                                    _ = v44
                                                case "foo bar":
                                                    if v2 {
                                                        if false {
                                                            if v2 {
                                                                var v45 uint = 2
                                                                v12 = append(v12, -86)
                                                                if false {
                                                                    if true {
                                                                        if false {
                                                                            var v46 bool = false
                                                                            fmt.Println(v3, v21, v2)
                                                                            var v47 int16 = 1
                                                                            _ = v46
                                                                            _ = v47
                                                                        } else {
                                                                            v48 := GenS4{F0: 387, F1: 639}
                                                                            var v49 uint32 = 2147483647
                                                                            var v50 float64 = 1.5
                                                                            _ = len(v12)
                                                                            _ = v48
                                                                            _ = v49
                                                                            _ = v50
                                                                        }
                                                                        switch (fi15 % ((fi18) | 1)) {
                                                                        case -9223372036854775808:
                                                                            var v51 uint16 = v4
                                                                            _ = v51
                                                                        case 726:
                                                                            if v1 {
                                                                                var v52 int16 = -260
                                                                                v45 += 4294967295
                                                                                _ = v52
                                                                            }
                                                                            if ((!v2) && (!v1)) {
                                                                                var v53 uint64 = 2
                                                                                switch 32768 {
                                                                                case 4294967295:
                                                                                    v53 += 705
                                                                                case 128:
                                                                                    var v54 float64 = (v21 / v21)
                                                                                    _ = v54
                                                                                case 9223372036854775807:
                                                                                    v21 = (v21 - 0.0)
                                                                                    if len(v12) > 0 {
                                                                                        v12[0] = fi15
                                                                                    }
                                                                                default:
                                                                                    if len(v12) > 0 {
                                                                                        v12[0] = 9223372036854775807
                                                                                    }
                                                                                }
                                                                                v21 += v21
                                                                                _ = v53
                                                                            } else {
                                                                                v45 += (^(v45))
                                                                            }
                                                                            if v1 {
                                                                                var v55 int64 = -1
                                                                                fmt.Println(v4, v55, v1)
                                                                                var v56 bool = v2
                                                                                _ = v55
                                                                                _ = v56
                                                                            } else {
                                                                                var v57 int16 = -32768
                                                                                _ = v57
                                                                            }
                                                                        case 32768:
                                                                            switch fi18 {
                                                                            case 657:
                                                                                var v58 bool = v2
                                                                                _ = v58
                                                                            case 0:
                                                                                fmt.Println(v21)
                                                                            default:
                                                                                switch v3 {
                                                                                case "a":
                                                                                    if len(v12) > 0 {
                                                                                        v12[0] = fi18
                                                                                    }
                                                                                case "":
                                                                                    v12 = append(v12, 443)
                                                                                case "foo bar":
                                                                                    v59 := GenS1{F0: 834, F1: (fi15 + (fi18 | (fi15 >> ((v45) & 63))))}
                                                                                    _ = v59
                                                                                default:
                                                                                    fmt.Println(v2, v4, v45)
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    if (!true) {
                                                                        v21 = 1e10
                                                                        v60 := GenS1{F0: 194, F1: fi18}
                                                                        fmt.Println(fi18)
                                                                        _ = v60
                                                                    } else {
                                                                        v45 = 1
                                                                        v2 = v2
                                                                        v4 = (v4 & v4)
                                                                        switch 78 {
                                                                        case 1:
                                                                            v61 := GenS5{F0: v4, F1: true, F2: 567}
                                                                            _ = v61
                                                                        case 149:
                                                                            var v62 uint32 = 2147483648
                                                                            _ = v62
                                                                        case 4294967295:
                                                                            if len(v12) > 0 {
                                                                                v12[0] = (fi18 >> ((v45) & 63))
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v63 uint64 = 1
                                                                    var v64 int8 = -128
                                                                    _ = v63
                                                                    _ = v64
                                                                }
                                                                _ = v45
                                                            }
                                                            var v65 uint64 = 0
                                                            v21 += ((v21 + (float64(0.0) * float64(0.1))) / v21)
                                                            _ = v65
                                                        } else {
                                                            v12 = append(v12, fi15)
                                                            if (1 > 128) {
                                                                if len(v12) >= 2 {
                                                                    v12 = v12[1:len(v12)]
                                                                }
                                                            } else {
                                                                var v66 uint64 = 0
                                                                if len(v12) > 0 {
                                                                    v12[0] = (-(fi15))
                                                                }
                                                                _ = v66
                                                            }
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                        }
                                                        if len(v12) >= 2 {
                                                            v12 = v12[1:len(v12)]
                                                        }
                                                        v3 = v3
                                                    }
                                                }
                                            } else {
                                                if v2 {
                                                    fmt.Println(fi15, fi15)
                                                }
                                            }
                                            var v67 int16 = -1
                                            _ = v67
                                        }
                                    }
                                default:
                                    v2 = (1 <= 4294967295)
                                    v2 = (42 != -1)
                                }
                            }
                        case "\n\t":
                            var v68 uint = 255
                            _ = v68
                        default:
                            v4 = v4
                        }
                        switch (^(fi18)) {
                        case 9223372036854775807:
                            _ = len(v12)
                        case 32767:
                            var v69 uint32 = 407
                            var v70 uint = 255
                            _ = v69
                            _ = v70
                        case 32768:
                            fmt.Println(v3)
                            var v71 uint = 18446744073709551615
                            if (fi18 <= 1) {
                                v3 = v3
                                v12 = append(v12, 83)
                                var v72 int8 = 100
                                _ = v72
                            }
                            _ = v71
                        }
                        _ = v21
                    default:
                        v1 = (!(0 < uint(v4)))
                    }
                    _ = len(v12)
                    v4 = 570
                    var v73 float64 = 3.14159
                    _ = fi18
                    _ = v73
                }
                v74 := GenS1{F0: 1, F1: 32768}
                _ = v74
            }
            for fi75 := 0; fi75 < 1; fi75++ {
                var v76 int32 = 2147483647
                v77 := GenS4{F0: 0, F1: 9223372036854775807}
                v1 = ((127 == 42) && false)
                var v78 int16 = 42
                if (int64(fi75) > int64(v78)) {
                    v4 = 472
                    v79 := GenS2{F0: v2, F1: 145}
                    if false {
                        switch fi75 {
                        case -779:
                            var v80 uint = 255
                            _ = v80
                        }
                    } else {
                        var v81 bool = (true && (v1 || ((1.5 > (0.1 - ((float32(0.1) * -2.25) + ((((3.14159 + 3.14159) + 0.0) - 1e10) + -2.25)))) && (!(((v2 && (0.1 > 0.1)) && v1) || v1)))))
                        fmt.Println(v3, fi15)
                        _ = v81
                    }
                    _ = v79
                } else {
                    switch ("a" + v3) {
                    case "\n\t":
                        v4 = 255
                    case "a":
                        var v82 uint = 1
                        _ = v82
                    case "foo bar":
                        v4 -= (v4 + (v4 / ((v4) | 1)))
                    }
                }
                _ = fi75
                _ = v76
                _ = v77
                _ = v78
            }
        } else {
            for fi83 := 0; fi83 < 20; fi83++ {
                _ = len(v12)
                if v1 {
                    if len(v12) > 0 {
                        v12[0] = -2147483648
                    }
                    if v2 {
                        v3 = "a"
                        v2 = v2
                    }
                    var v84 int32 = -741
                    _ = v84
                } else {
                    _ = len(v12)
                    if (!v1) {
                        if v1 {
                            var v85 uint = 65535
                            _ = v85
                        }
                    }
                    var v86 int32 = 811
                    _ = v86
                }
                _ = fi83
            }
            v87 := GenS2{F0: false, F1: 0}
            for fi88 := 0; fi88 < 2; fi88++ {
                v2 = false
                v89 := GenS6{F0: 0}
                _ = fi88
                _ = v89
            }
            _ = v87
        }
        v90 := GenS4{F0: 4294967295, F1: 42}
        fmt.Println(fi15)
        _ = fi15
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
