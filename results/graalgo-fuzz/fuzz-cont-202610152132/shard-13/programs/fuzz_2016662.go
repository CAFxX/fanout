package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int16; F2 int32 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 1
    v3 := []uint32{2, 65535, 464}
    v3 = append(v3, 2)
    var v4 int8 = 100
    var v5 string = (("" + "a") + "\n\t")
    v1[883] = 42
    switch 4294967295 {
    case 127:
        var v6 float64 = float64(1e10)
        _ = v6
    case 172:
        var v7 uint16 = uint16(v2)
        _ = v7
    default:
        for fi8 := 0; fi8 < 12; fi8++ {
            if false {
                if true {
                    v9 := v1[2147483647]
                    if true {
                        switch fi8 {
                        case 128:
                            if (!((v5 + ((v5 + ((v5 + (("\n\t" + ((v5 + "\n\t") + "foo bar")) + ("" + v5))) + v5)) + ("\n\t" + ""))) > v5)) {
                                fmt.Println(v4, v5)
                                v2 = (v2 - v2)
                            } else {
                                v10 := GenS1{F0: 32768}
                                if ((-103 >= 42) && true) {
                                    fmt.Println(v5, fi8)
                                    v11 := GenS1{F0: -894}
                                    _ = v11
                                } else {
                                    fmt.Println(v9)
                                    if true {
                                        for fi12 := 0; fi12 < 17; fi12++ {
                                            var v13 float32 = 3.14159
                                            var v14 float32 = (-2.25 + 1e10)
                                            var v15 int16 = 42
                                            if (0.1 < (v14 * float32(3.14159))) {
                                                fmt.Println(v2, v15, fi12)
                                                var v16 bool = false
                                                _ = v16
                                            } else {
                                                if ((v4 & (-(v4))) > (v4 << ((255) & 63))) {
                                                    v13 -= float32(3.14159)
                                                } else {
                                                    v1[v9] = 2147483648
                                                    var v17 uint32 = 776
                                                    _ = v17
                                                }
                                                var v18 int = (fi12 << ((255) & 63))
                                                v18 += (v18 - fi12)
                                                _ = v18
                                            }
                                            _ = fi12
                                            _ = v13
                                            _ = v14
                                            _ = v15
                                        }
                                        fmt.Println(v2, v9)
                                        var v19 float64 = (0.0 - 0.1)
                                        _ = v19
                                    } else {
                                        v2 += (v2 & (v2 << ((65535) & 63)))
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    var v20 uint64 = v2
                                    _ = v20
                                }
                                var v21 int32 = 42
                                _ = v10
                                _ = v21
                            }
                        default:
                            v2 += (v2 | (v2 | 4294967295))
                        }
                        if (323 <= 255) {
                            v4 = v4
                            _ = len(v1)
                        } else {
                            var v22 uint32 = (^(v9))
                            for fi23 := 0; fi23 < 13; fi23++ {
                                var v24 int32 = 614
                                switch "a" {
                                case "\n\t":
                                    v9 = v9
                                default:
                                    if false {
                                        v25 := GenS1{F0: -9223372036854775808}
                                        _ = v25
                                    }
                                }
                                if true {
                                    if (1.5 > 3.14159) {
                                        fmt.Println(fi8, v4, fi23)
                                        _ = len(v1)
                                    }
                                    switch int(v9) {
                                    case -190:
                                        var v26 bool = false
                                        _ = v26
                                    case 1:
                                        v27 := GenS2{F0: (-(v4)), F1: 819, F2: (v24 | v24)}
                                        _ = v27
                                    default:
                                        switch -672 {
                                        case -9223372036854775808:
                                            delete(v1, (v22 | (v22 - uint32(v2))))
                                        case -2147483648:
                                            var v28 uint32 = v22
                                            _ = v28
                                        default:
                                            fmt.Println(fi23, v2, v9)
                                            var v29 float32 = 1e10
                                            _ = v29
                                        }
                                    }
                                    v22 = (v22 & v22)
                                }
                                v5 = (v5 + "\n\t")
                                v22 = 1
                                v4 = (v4 - v4)
                                v30 := GenS2{F0: (-(v4)), F1: 1, F2: (v24 - (^(v24)))}
                                _ = fi23
                                _ = v24
                                _ = v30
                            }
                            var v31 int8 = (v4 * (v4 << ((uint(v2)) & 63)))
                            _ = v22
                            _ = v31
                        }
                        if true {
                            var v32 int = 2147483647
                            _ = v32
                        }
                    } else {
                        v33 := v1[4294967295]
                        for fi34 := 0; fi34 < 3; fi34++ {
                            v1[(v9 * 754)] = v9
                            var v35 float32 = 1e10
                            _ = fi34
                            _ = v35
                        }
                        if (1e10 != (0.1 * 0.0)) {
                            for fi36 := 0; fi36 < 3; fi36++ {
                                if ((^(fi36)) != 128) {
                                    if false {
                                        fmt.Println(v9)
                                    }
                                    fmt.Println(v4)
                                    var v37 bool = true
                                    _ = v37
                                } else {
                                    v9 -= 0
                                    fmt.Println(fi36)
                                    v4 -= (v4 % (((v4 - v4)) | 1))
                                }
                                v38 := v1[v33]
                                v9 -= 1
                                v39 := GenS2{F0: -38, F1: -1, F2: 2147483647}
                                v1[(v38 & uint32(v2))] = 65535
                                _ = fi36
                                _ = v38
                                _ = v39
                            }
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v9, v4, v4)
                            if (-2147483648 < (-(fi8))) {
                                if (true && (false && (v4 > (^(v4))))) {
                                    if ((!false) || (0 == 221)) {
                                        if (-840 < 42) {
                                            var v40 string = "foo bar"
                                            _ = v40
                                        } else {
                                            fmt.Println(fi8, fi8, v5)
                                            v2 = 18446744073709551615
                                            if ((v33 / ((v9) | 1)) <= v33) {
                                                fmt.Println(v9, v5)
                                            }
                                        }
                                        for fi41 := 0; fi41 < 18; fi41++ {
                                            v5 = (((((v5 + "a") + ("hello" + (v5 + v5))) + "\n\t") + "hello") + v5)
                                            if (0 > uint8(v2)) {
                                                _ = len(v1)
                                                var v42 uint = 62
                                                if ((!(0.0 != float64(0.0))) && ((^(v2)) != (v2 / ((v2) | 1)))) {
                                                    var v43 int8 = (v4 | v4)
                                                    if true {
                                                        switch 980 {
                                                        case -443:
                                                            switch "hello" {
                                                            case "hello":
                                                                v44 := GenS3{F0: (v42 | (v42 % (((v42 >> ((100) & 63))) | 1))), F1: 2}
                                                                _ = v44
                                                            case "\n\t":
                                                                v42 = (^(v42))
                                                            }
                                                        default:
                                                            if true {
                                                                var v45 float64 = ((1e10 + float64(float32(0.0))) * 0.0)
                                                                _ = v45
                                                            }
                                                        }
                                                        if false {
                                                            v4 += (v4 / ((-1) | 1))
                                                            v2 = v2
                                                            var v46 uint = (v42 | 48)
                                                            _ = v46
                                                        } else {
                                                            v42 += (v42 >> ((v42) & 63))
                                                            var v47 uint8 = 2
                                                            _ = v47
                                                        }
                                                    } else {
                                                        v42 = uint(v33)
                                                        v48 := GenS2{F0: (v43 ^ v43), F1: 1000, F2: 2147483647}
                                                        _ = v48
                                                    }
                                                    _ = v43
                                                } else {
                                                    var v49 int16 = -32768
                                                    _ = v49
                                                }
                                                _ = v42
                                            } else {
                                                v5 = v5
                                                if len(v3) > 0 {
                                                    v3[0] = (v33 >> ((65535) & 63))
                                                }
                                            }
                                            var v50 uint8 = 127
                                            _ = fi41
                                            _ = v50
                                        }
                                    } else {
                                        var v51 uint64 = (v2 - v2)
                                        var v52 int = fi8
                                        v53 := GenS1{F0: (^(v52))}
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    }
                                    v54 := v1[v33]
                                    v2 -= (v2 + v2)
                                    _ = v54
                                }
                                for fi55 := 0; fi55 < 19; fi55++ {
                                    fmt.Println(v33, fi8, v9)
                                    _ = fi55
                                }
                                var v56 int8 = (v4 << ((65535) & 63))
                                _ = v56
                            }
                        }
                        _ = v33
                    }
                    _ = v9
                } else {
                    var v57 string = ("\n\t" + v5)
                    var v58 float64 = 3.14159
                    _ = v57
                    _ = v58
                }
            }
            switch ((("\n\t" + "foo bar") + v5) + "foo bar") {
            case "\n\t":
                var v59 float32 = 1.5
                _ = v59
            case "a":
                v2 = (v2 % ((v2) | 1))
            case "":
                for fi60 := 0; fi60 < 8; fi60++ {
                    switch (fi8 % (((-(fi8))) | 1)) {
                    case 1:
                        if (((v2 >> ((208) & 63)) == v2) || (3.14159 <= (((float32((((1.5 - 1.5) / (3.14159 - (0.1 * 1e10))) * float64(0.0))) * 1e10) - (3.14159 * (1e10 + 3.14159))) - float32((0.1 + 3.14159))))) {
                            fmt.Println(v2, v5)
                            if true {
                                var v61 uint8 = 127
                                switch 4294967295 {
                                case -991:
                                    if (((1 <= 65535) || (!((0.1 / (3.14159 / 1.0)) >= 1.5))) && ((v2 & uint64(v61)) >= v2)) {
                                        if (128 > (v61 - v61)) {
                                            var v62 bool = (1 == v61)
                                            v63 := GenS2{F0: (v4 / ((6) | 1)), F1: 32767, F2: 249}
                                            _ = v62
                                            _ = v63
                                        } else {
                                            var v64 float64 = 0.0
                                            var v65 bool = (0 <= 41)
                                            _ = v64
                                            _ = v65
                                        }
                                        if (true && true) {
                                            fmt.Println(fi60, v4)
                                            if (295 != 0) {
                                                var v66 float32 = 0.1
                                                var v67 bool = true
                                                _ = v66
                                                _ = v67
                                            }
                                        } else {
                                            v5 = "foo bar"
                                        }
                                        if false {
                                            var v68 uint32 = 1
                                            _ = v68
                                        }
                                    } else {
                                        var v69 int16 = -899
                                        v61 += v61
                                        var v70 int16 = (v69 / ((-1) | 1))
                                        _ = v69
                                        _ = v70
                                    }
                                default:
                                    if true {
                                        if (((v2 != 684) && (-825 < -1)) || ("foo bar" > "a")) {
                                            fmt.Println(v61, fi8)
                                            var v71 bool = (42 > 1)
                                            _ = v71
                                        } else {
                                            var v72 int16 = 754
                                            v5 = v5
                                            fmt.Println(v2, v5, v72)
                                            var v73 int = fi8
                                            v74 := GenS1{F0: 32767}
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                        }
                                        v61 = (v61 + v61)
                                        var v75 int = (fi60 / ((fi60) | 1))
                                        _ = v75
                                    } else {
                                        var v76 int8 = (v4 & -8)
                                        _ = v76
                                    }
                                }
                                var v77 uint16 = 32767
                                _ = v61
                                _ = v77
                            }
                            switch fi60 {
                            case 4294967295:
                                if len(v3) > 0 {
                                    v3[0] = 2147483647
                                }
                            }
                        }
                    }
                    fmt.Println(v5, fi60)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi60
                }
            default:
                var v78 uint8 = 1
                _ = v78
            }
            switch v5 {
            case "\n\t":
                var v79 string = "hello"
                _ = v79
            case "hello":
                for fi80 := 0; fi80 < 16; fi80++ {
                    if (true && ((((0.0 + 3.14159) / (float64(1e10) + (-2.25 + 1e10))) <= (3.14159 - 0.0)) && (true || false))) {
                        v1[4294967295] = 42
                        switch v5 {
                        case "\n\t":
                            var v81 string = (v5 + v5)
                            _ = v81
                        case "":
                            v4 -= 127
                        case "a":
                            v4 += -1
                        default:
                            var v82 uint8 = 89
                            _ = v82
                        }
                    } else {
                        fmt.Println(v2, v5, fi80)
                    }
                    _ = fi80
                }
            }
            _ = fi8
        }
    }
    fmt.Println(v5, v2, v2)
    var v83 string = "a"
    var v84 uint = 0
    v5 = ("hello" + "foo bar")
    switch ((("a" + v83) + (("a" + (v5 + ("a" + v5))) + v5)) + v83) {
    case "foo bar":
        v83 = (v5 + "hello")
    case "a":
        v85 := GenS4{F0: 544, F1: v2}
        _ = v85
    }
    fmt.Println(v5, v83)
    v86 := GenS4{F0: 623, F1: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v84
    _ = v86
}
