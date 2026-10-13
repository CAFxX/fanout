package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 int64 }

func main() {
    var v1 float32 = (-2.25 * (0.1 - -2.25))
    var v2 uint16 = 42
    var v3 int16 = -32768
    v4 := make(map[uint32]uint32)
    v1 -= float32((1e10 - 0.1))
    var v5 int8 = -36
    fmt.Println(v2)
    var v6 string = "hello"
    _ = len(v4)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 1, F1: 65535, F2: fi7}
        if true {
            v9 := GenS2{F0: fi7, F1: (v3 * v3)}
            v10 := []int{fi7}
            if true {
                var v11 bool = (true && false)
                delete(v4, 626)
                v12 := GenS1{F0: fi7, F1: 0, F2: fi7}
                _ = v11
                _ = v12
            }
            _ = v10
            _ = v9
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            switch "foo bar" {
            case "a":
                v3 += (v3 ^ 1000)
            default:
                v2 = (v2 / (((v2 / (((v2 >> ((42) & 63))) | 1))) | 1))
            }
            if false {
                v14 := GenS2{F0: (fi7 ^ (fi13 * fi13)), F1: (v3 + (v3 + (v3 % ((v3) | 1))))}
                var v15 bool = (true || (234 >= 221))
                v16 := []int{}
                if (380 == 2) {
                    switch 63 {
                    case 128:
                        fmt.Println(fi13, fi13)
                    case 4294967295:
                        var v17 int16 = 42
                        _ = v17
                    case 2147483647:
                        if (((3.14159 / 1.0) < 3.14159) || false) {
                            fmt.Println(v5, v2, v1)
                            v18 := v4[1]
                            v2 = (v2 * v2)
                            var v19 uint16 = v2
                            if (0 < 1) {
                                v20 := GenS1{F0: (fi13 & (-(fi7))), F1: (v2 - (v19 & 255)), F2: (^(fi13))}
                                var v21 int64 = 0
                                _ = v20
                                _ = v21
                            } else {
                                if false {
                                    if len(v16) > 0 {
                                        v16[0] = (fi7 | fi7)
                                    }
                                    v22 := GenS3{F0: (v5 << ((18446744073709551615) & 63))}
                                    if (!v15) {
                                        v23 := v4[v18]
                                        delete(v4, 404)
                                        if (v5 <= (v5 << ((879) & 63))) {
                                            v24 := GenS1{F0: fi13, F1: 2, F2: fi7}
                                            var v25 float32 = (((v1 * v1) * v1) * v1)
                                            v23 = v18
                                            _ = v24
                                            _ = v25
                                        }
                                        _ = v23
                                    }
                                    var v26 uint64 = 2
                                    _ = v22
                                    _ = v26
                                } else {
                                    v27 := GenS3{F0: (v5 / (((v5 / ((v5) | 1))) | 1))}
                                    v19 = v19
                                    _ = v27
                                }
                                v5 += v5
                            }
                            _ = v18
                            _ = v19
                        } else {
                            v28 := GenS2{F0: 2147483647, F1: 42}
                            _ = v28
                        }
                    }
                    v2 -= (v2 - 32767)
                }
                if false {
                    if true {
                        if (2147483647 != 2147483647) {
                            fmt.Println(fi7, v5)
                            v5 = v5
                        } else {
                            if v15 {
                                v1 = v1
                                if v15 {
                                    var v29 uint16 = (^(v2))
                                    if (!(4294967295 >= 445)) {
                                        if false {
                                            fmt.Println(fi7, v29)
                                            switch (v6 + "foo bar") {
                                            case "":
                                                switch (fi7 - (fi13 % (((fi7 >> ((710) & 63))) | 1))) {
                                                case 282:
                                                    fmt.Println(v3, v2, fi13)
                                                default:
                                                    if (!v15) {
                                                        var v30 int32 = 606
                                                        _ = v30
                                                    } else {
                                                        if (0 > uint32(v29)) {
                                                            v31 := GenS2{F0: (fi7 / ((int(v29)) | 1)), F1: v3}
                                                            var v32 int64 = -9223372036854775808
                                                            _ = v31
                                                            _ = v32
                                                        } else {
                                                            v15 = v15
                                                            v29 += v29
                                                            if v15 {
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    if v15 {
                                                                        if (65535 == (^(v29))) {
                                                                            if len(v16) > 0 {
                                                                                v16[0] = (fi7 / ((fi7) | 1))
                                                                            }
                                                                            if len(v16) > 0 {
                                                                                v16[0] = (fi13 << ((1) & 63))
                                                                            }
                                                                            v33 := GenS1{F0: fi13, F1: v2, F2: 32767}
                                                                            _ = v33
                                                                        } else {
                                                                            fmt.Println(v3, v2)
                                                                        }
                                                                        switch (-(fi7)) {
                                                                        case 32767:
                                                                            fmt.Println(v5)
                                                                        case -157:
                                                                            v34 := GenS4{F0: v5, F1: 2}
                                                                            _ = v34
                                                                        case -694:
                                                                            _ = len(v16)
                                                                        }
                                                                    } else {
                                                                        var v35 uint32 = 1
                                                                        v36 := GenS2{F0: fi7, F1: -1}
                                                                        v15 = true
                                                                        _ = v35
                                                                        _ = v36
                                                                    }
                                                                case "foo bar":
                                                                    switch v6 {
                                                                    case "foo bar":
                                                                        var v37 uint64 = 819
                                                                        _ = v37
                                                                    case "a":
                                                                        if false {
                                                                            v16 = append(v16, 32768)
                                                                            v1 += float32((0.1 - 1.5))
                                                                        }
                                                                    }
                                                                case "":
                                                                    v29 += 639
                                                                }
                                                                switch fi13 {
                                                                case 359:
                                                                    fmt.Println(v1)
                                                                case 2147483647:
                                                                    v5 = 127
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            case "a":
                                                var v38 int = (fi13 & fi7)
                                                _ = v38
                                            case "hello":
                                                v5 -= v5
                                            }
                                        }
                                    }
                                    _ = v29
                                } else {
                                    delete(v4, 0)
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    var v39 int16 = (v3 >> ((2) & 63))
                                    _ = v39
                                }
                            }
                        }
                        if (!(!(266 != 2))) {
                            var v40 uint64 = 18446744073709551615
                            if ((955 < 742) && v15) {
                                if (false && (42 != 42)) {
                                    var v41 float64 = 3.14159
                                    _ = v41
                                } else {
                                    if (true && false) {
                                        if v15 {
                                            switch fi7 {
                                            case 2147483647:
                                                var v42 uint8 = 45
                                                _ = v42
                                            case -1:
                                                switch (fi7 / (((fi13 ^ (fi13 ^ fi13))) | 1)) {
                                                case -9223372036854775808:
                                                    fmt.Println(v1, v5, v15)
                                                case 802:
                                                    switch (^(fi7)) {
                                                    case 32768:
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            fmt.Println(fi7)
                                                        case "foo bar":
                                                            v43 := GenS2{F0: (fi7 | fi7), F1: (v3 * 1000)}
                                                            _ = v43
                                                        case "a":
                                                            switch v6 {
                                                            case "hello":
                                                                v44 := GenS5{F0: ""}
                                                                _ = v44
                                                            case "\n\t":
                                                                v5 = (-(v5))
                                                            }
                                                        }
                                                        v1 += v1
                                                    case -2147483648:
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                    case 9223372036854775807:
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                        var v45 int64 = -577
                                                        _ = v45
                                                    default:
                                                        v1 = float32(0.0)
                                                    }
                                                case 0:
                                                    if true {
                                                        if (-2.25 == float64(v1)) {
                                                            v1 += v1
                                                            _ = len(v16)
                                                        } else {
                                                            switch v6 {
                                                            case "a":
                                                                if true {
                                                                    v46 := GenS4{F0: -1, F1: 2}
                                                                    _ = v46
                                                                } else {
                                                                    v15 = (!(128 == 26))
                                                                }
                                                            case "foo bar":
                                                                v15 = (1 > 0)
                                                            case "":
                                                                v47 := GenS1{F0: fi7, F1: 558, F2: (fi7 << ((42) & 63))}
                                                                _ = v47
                                                            }
                                                        }
                                                    }
                                                }
                                            case 127:
                                                var v48 string = ("\n\t" + v6)
                                                _ = v48
                                            default:
                                                v4[2147483647] = 2
                                            }
                                            if v15 {
                                                if true {
                                                    var v49 int32 = 2147483647
                                                    if true {
                                                        v50 := GenS2{F0: fi13, F1: 1}
                                                        v51 := GenS6{F0: 0, F1: -1}
                                                        switch (fi13 / ((-9223372036854775808) | 1)) {
                                                        case 4294967295:
                                                            delete(v4, 635)
                                                            v5 -= (-(v5))
                                                            if (v3 > (v3 << ((1) & 63))) {
                                                                if (-2.25 != float64(v1)) {
                                                                    v6 = "a"
                                                                    switch fi13 {
                                                                    case -210:
                                                                        v52 := v4[130]
                                                                        _ = v52
                                                                    case -155:
                                                                        v1 += 0.0
                                                                    }
                                                                } else {
                                                                    v16 = append(v16, fi7)
                                                                    if ("foo bar" == ((v6 + ("hello" + ((v6 + "hello") + v6))) + "hello")) {
                                                                        if (v15 && (-9223372036854775808 >= 0)) {
                                                                            fmt.Println(fi13)
                                                                            fmt.Println(v5, v40, v5)
                                                                        } else {
                                                                            if (v15 && v15) {
                                                                                var v53 int8 = 110
                                                                                v54 := GenS2{F0: (-(fi13)), F1: (v3 - (v3 | (v3 * (v3 | (v3 & v3)))))}
                                                                                _ = v53
                                                                                _ = v54
                                                                            } else {
                                                                                var v55 bool = (-559 > (fi7 << ((346) & 63)))
                                                                                _ = v55
                                                                            }
                                                                            if (((1.5 + -2.25) / -2.25) < -2.25) {
                                                                                v6 = v6
                                                                                v4[2147483647] = 785
                                                                                v1 = v1
                                                                                v56 := GenS6{F0: 1, F1: -9223372036854775808}
                                                                                v15 = (0.0 >= 3.14159)
                                                                                v40 = v40
                                                                                _ = v56
                                                                            }
                                                                        }
                                                                        var v57 int64 = 42
                                                                        _ = v57
                                                                    } else {
                                                                        var v58 int32 = -457
                                                                        v2 = (v2 ^ (v2 << ((42) & 63)))
                                                                        _ = v58
                                                                    }
                                                                    v59 := v4[0]
                                                                    _ = v59
                                                                }
                                                                v4[0] = 65535
                                                            }
                                                        case 2147483647:
                                                            var v60 float64 = (1e10 * -2.25)
                                                            _ = v60
                                                        }
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        fmt.Println(fi7, v2, v1)
                                                        v15 = v15
                                                        if (!(false && v15)) {
                                                            v40 -= (^(v40))
                                                        } else {
                                                            v61 := v4[2147483648]
                                                            var v62 uint32 = v61
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                    }
                                                    _ = v49
                                                } else {
                                                    v63 := GenS4{F0: (v5 << ((uint(v40)) & 63)), F1: uint8(v2)}
                                                    var v64 string = ((v6 + "") + (v6 + ("\n\t" + v6)))
                                                    v16 = append(v16, -185)
                                                    if v15 {
                                                        var v65 bool = (v40 >= (v40 | v40))
                                                        _ = v65
                                                    } else {
                                                        if true {
                                                            v5 -= v5
                                                        }
                                                        fmt.Println(v6, v1)
                                                    }
                                                    _ = v63
                                                    _ = v64
                                                }
                                            } else {
                                                var v66 uint32 = 2147483648
                                                _ = v66
                                            }
                                            switch fi13 {
                                            case -1:
                                                if (v15 && true) {
                                                    fmt.Println(v6, v2)
                                                }
                                            case 32767:
                                                if len(v16) > 0 {
                                                    v16[0] = 2147483647
                                                }
                                            }
                                        }
                                        v1 += v1
                                        v2 = v2
                                    } else {
                                        fmt.Println(v1, v2, v1)
                                        var v67 int32 = 1
                                        _ = v67
                                    }
                                }
                                switch fi7 {
                                case 734:
                                    v40 -= (^(v40))
                                case 32767:
                                    fmt.Println(v3)
                                default:
                                    switch -235 {
                                    case 9223372036854775807:
                                        var v68 float32 = (1.5 / 0.1)
                                        _ = v68
                                    case -924:
                                        v69 := GenS4{F0: v5, F1: 1}
                                        _ = v69
                                    case 0:
                                        fmt.Println(v6)
                                    default:
                                        var v70 uint8 = 255
                                        _ = v70
                                    }
                                }
                            }
                            fmt.Println(v6, fi13, v15)
                            _ = v40
                        } else {
                            var v71 uint16 = (^(v2))
                            v3 += (v3 >> ((4294967295) & 63))
                            _ = v71
                        }
                    }
                    v72 := GenS6{F0: 4294967295, F1: -1}
                    var v73 int32 = 42
                    _ = v72
                    _ = v73
                }
                if v15 {
                    var v74 uint64 = 2
                    _ = v74
                }
                _ = v14
                _ = v15
                _ = v16
            }
            var v75 uint16 = 760
            _ = fi13
            _ = v75
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v3, v6)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
