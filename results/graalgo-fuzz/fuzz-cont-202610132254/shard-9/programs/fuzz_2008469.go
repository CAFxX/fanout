package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 int; F2 int32 }
type GenS8 struct { F0 bool; F1 uint; F2 uint32 }
type GenS9 struct { F0 uint; F1 uint8 }
type GenS10 struct { F0 uint; F1 uint64 }
type GenS11 struct { F0 int32; F1 uint32; F2 int }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int64 }
func genFn0(p0 int, p1 int64, p2 uint) uint {
    v4 := []int{9223372036854775807, (p0 | -1)}
    for fi5 := 0; fi5 < 13; fi5++ {
        p1 = (p1 / (((p1 ^ p1)) | 1))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 255
}
func genFn1(p0 int16, p1 uint32, p2 uint8) int8 {
    fmt.Println(p2, p0, p2)
    if false {
        v6 := GenS1{F0: -1}
        var v7 int8 = 100
        if true {
            var v8 string = ("hello" + "\n\t")
            var v9 uint8 = p2
            _ = v8
            _ = v9
        } else {
            v10 := []int{1}
            _ = v10
        }
        _ = v6
        _ = v7
    } else {
        p0 = 1
        var v11 int16 = (p0 & (p0 << ((527) & 63)))
        v12 := []int{-1}
        _ = v11
        _ = v12
    }
    fmt.Println(p2)
    p0 -= 1000
    _ = p0
    _ = p1
    _ = p2
    return 6
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 4294967295
    var v3 uint8 = 158
    fmt.Println(v2, v2)
    switch -735 {
    case -87:
        fmt.Println(v2, v3)
    case -917:
        var v13 string = "foo bar"
        _ = v13
    default:
        var v14 uint32 = 438
        _ = v14
    }
    if ("" < "hello") {
        var v15 int8 = 1
        _ = v15
    } else {
        v2 = v2
    }
    v2 -= v2
    switch 599 {
    case -841:
        switch -58 {
        case -533:
            v3 -= v3
        }
    case 127:
        v2 = uint(v3)
    case -696:
        v2 += (v2 + v2)
    }
    var v16 int16 = -901
    fmt.Println(v3, v3, v3)
    for fi17 := 0; fi17 < 9; fi17++ {
        if true {
            if false {
                fmt.Println(v16, fi17)
                v18 := GenS1{F0: -1}
                _ = v18
            } else {
                if ((float64(float32(1e10)) * 0.0) == float64(-2.25)) {
                    if true {
                        if true {
                            for fi19 := 0; fi19 < 16; fi19++ {
                                fmt.Println(fi17, v16)
                                fmt.Println(fi19)
                                if false {
                                    v20 := GenS1{F0: 42}
                                    _ = v20
                                }
                                _ = fi19
                            }
                        }
                        fmt.Println(fi17, v2)
                        v21 := GenS2{F0: -436}
                        _ = v21
                    }
                    if ((!true) || true) {
                        v16 = v16
                        v22 := GenS1{F0: 0}
                        for fi23 := 0; fi23 < 17; fi23++ {
                            delete(v1, "a")
                            _ = fi23
                        }
                        _ = v22
                    }
                    v16 = 42
                }
                v16 = (v16 % ((1) | 1))
                fmt.Println(fi17)
            }
        }
        if (!(0.1 > 3.14159)) {
            switch (-(fi17)) {
            case -1:
                if false {
                    if (2147483648 == 82) {
                        for fi24 := 0; fi24 < 11; fi24++ {
                            if true {
                                v2 -= genFn0((fi24 % ((fi24) | 1)), -108, (v2 - (v2 << ((genFn0(fi24, 9223372036854775807, (v2 & (^(v2))))) & 63))))
                                var v25 int = (^(fi24))
                                var v26 bool = (v16 >= v16)
                                _ = v25
                                _ = v26
                            } else {
                                v27 := GenS3{F0: 42, F1: -1, F2: 1}
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v2, v16, v16)
                                case "foo bar":
                                    v28 := GenS4{F0: v3, F1: -9223372036854775808, F2: 523}
                                    var v29 uint = v2
                                    v30 := GenS1{F0: int32(v16)}
                                    var v31 int8 = 42
                                    v32 := []int{fi17, int(v16), fi17, fi24}
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                var v33 int32 = 1
                                _ = v27
                                _ = v33
                            }
                            v2 = 65535
                            _ = fi24
                        }
                        fmt.Println(v3, v3)
                        for fi34 := 0; fi34 < 15; fi34++ {
                            if (int64(v3) > 385) {
                                fmt.Println(v16, v2)
                                v16 -= -32768
                                v35 := GenS5{F0: (v16 - 0)}
                                _ = v35
                            }
                            _ = fi34
                        }
                    } else {
                        switch fi17 {
                        case 128:
                            var v36 uint32 = 1
                            _ = v36
                        case 42:
                            for fi37 := 0; fi37 < 7; fi37++ {
                                switch "\n\t" {
                                case "hello":
                                    fmt.Println(fi17, v16)
                                }
                                fmt.Println(fi17, v3)
                                var v38 uint8 = 99
                                _ = fi37
                                _ = v38
                            }
                        case 888:
                            if (false || true) {
                                v39 := GenS5{F0: v16}
                                fmt.Println(v3, v2, v16)
                                _ = v39
                            }
                        default:
                            for fi40 := 0; fi40 < 13; fi40++ {
                                v41 := GenS6{F0: 977}
                                v2 += (v2 << (((v2 << (((v2 - (v2 - 65535))) & 63))) & 63))
                                _ = fi40
                                _ = v41
                            }
                        }
                        if (genFn1(v16, 2, 255) < 41) {
                            var v42 int = fi17
                            fmt.Println(v2, v3, v42)
                            _ = v42
                        } else {
                            for fi43 := 0; fi43 < 20; fi43++ {
                                var v44 int64 = 1
                                var v45 bool = true
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                        }
                    }
                } else {
                    fmt.Println(v2, v16, v16)
                    if ((v16 * 87) == (v16 >> ((v2) & 63))) {
                        if false {
                            v3 = 42
                            if false {
                                var v46 uint32 = 4294967295
                                var v47 string = ""
                                _ = v46
                                _ = v47
                            } else {
                                if false {
                                    var v48 uint8 = v3
                                    v49 := []int{(fi17 ^ fi17), (fi17 / (((fi17 - fi17)) | 1)), (fi17 / (((fi17 << ((495) & 63))) | 1)), (fi17 - fi17)}
                                    var v50 int16 = 504
                                    _ = v48
                                    _ = v49
                                    _ = v50
                                }
                            }
                        }
                        for fi51 := 0; fi51 < 11; fi51++ {
                            var v52 uint = v2
                            v3 = (v3 % ((v3) | 1))
                            _ = fi51
                            _ = v52
                        }
                    } else {
                        v53 := []int{fi17, fi17, (fi17 / (((-(fi17))) | 1)), (fi17 ^ fi17)}
                        v3 = (v3 & (v3 | v3))
                        _ = v53
                    }
                    v54 := []int{316, -2147483648, (fi17 + fi17)}
                    _ = v54
                }
            }
            v16 = (v16 | v16)
            v55 := GenS7{F0: 742, F1: fi17, F2: 258}
            _ = v55
        } else {
            if (!true) {
                v56 := GenS8{F0: (!(-892 <= v16)), F1: v2, F2: 273}
                _ = v56
            } else {
                if ((v2 + genFn0((-(fi17)), int64(fi17), (v2 & (v2 + v2)))) >= (^(v2))) {
                    for fi57 := 0; fi57 < 17; fi57++ {
                        if (1 > 331) {
                            if (!true) {
                                var v58 string = ""
                                v59 := GenS7{F0: 42, F1: -788, F2: -2147483648}
                                _ = v58
                                _ = v59
                            }
                            var v60 uint64 = 18446744073709551615
                            if false {
                                var v61 uint32 = 890
                                delete(v1, "a")
                                v62 := GenS4{F0: (v3 << ((v2) & 63)), F1: -594, F2: 1}
                                _ = v61
                                _ = v62
                            }
                            v63 := GenS9{F0: v2, F1: v3}
                            _ = v60
                            _ = v63
                        }
                        _ = fi57
                    }
                    var v64 int16 = 0
                    _ = v64
                }
                v3 = (^(v3))
                switch ("\n\t" + "hello") {
                case "hello":
                    if ((957 < 267) || true) {
                        v3 += (v3 & (v3 % ((uint8(v2)) | 1)))
                        for fi65 := 0; fi65 < 17; fi65++ {
                            v66 := []int{fi17, (-(fi17)), (fi17 & 9223372036854775807)}
                            var v67 uint32 = 2
                            _ = fi65
                            _ = v66
                            _ = v67
                        }
                        v16 = v16
                    } else {
                        v1["foo bar"] = "\n\t"
                        fmt.Println(v3)
                        v68 := GenS8{F0: false, F1: (v2 >> (((v2 >> (((v2 * (v2 - genFn0((fi17 ^ fi17), 1, (v2 - 1))))) & 63))) & 63)), F2: 603}
                        _ = v68
                    }
                }
            }
            for fi69 := 0; fi69 < 17; fi69++ {
                switch ("\n\t" + "a") {
                case "a":
                    v70 := GenS10{F0: (v2 & v2), F1: 978}
                    _ = v70
                case "foo bar":
                    v71 := GenS11{F0: 21, F1: 917, F2: -405}
                    _ = v71
                case "hello":
                    var v72 int8 = 0
                    _ = v72
                default:
                    v3 -= uint8(v2)
                }
                _ = fi69
            }
            if false {
                if (false && false) {
                    v73 := GenS10{F0: (v2 / (((v2 | v2)) | 1)), F1: 455}
                    v16 += v16
                    v74 := GenS4{F0: 255, F1: 90, F2: 65535}
                    _ = v73
                    _ = v74
                } else {
                    v2 += (v2 / (((v2 >> ((v2) & 63))) | 1))
                }
            }
        }
        _ = fi17
    }
    v75 := GenS9{F0: (^(v2)), F1: (^(v3))}
    fmt.Println(v16, v2, v3)
    var v76 uint16 = 577
    v77 := GenS12{F0: int16(v2)}
    v78 := GenS11{F0: -2147483648, F1: 1, F2: 0}
    fmt.Println(v2, v16, v16)
    if false {
        if false {
            for fi79 := 0; fi79 < 14; fi79++ {
                v76 += v76
                for fi80 := 0; fi80 < 11; fi80++ {
                    v2 -= (^(v2))
                    _ = fi80
                }
                if (!false) {
                    fmt.Println(v3)
                }
                _ = fi79
            }
        } else {
            v16 = (v16 & (v16 | v16))
        }
        fmt.Println(v76, v2, v2)
    } else {
        switch 2147483647 {
        case -9223372036854775808:
            var v81 int64 = -966
            _ = v81
        case 42:
            var v82 float32 = ((0.0 * ((1.5 - -2.25) * 1.5)) * 3.14159)
            _ = v82
        case 1:
            var v83 int8 = 0
            _ = v83
        default:
            for fi84 := 0; fi84 < 15; fi84++ {
                var v85 int = fi84
                if (!true) {
                    if false {
                        v1["\n\t"] = ("a" + "foo bar")
                        var v86 int = v85
                        _ = v86
                    }
                    var v87 int16 = (-(v16))
                    _ = v87
                } else {
                    v88 := GenS10{F0: (v2 >> (((v2 << ((255) & 63))) & 63)), F1: 0}
                    v89 := []int{(fi84 >> ((42) & 63)), 127, (^(fi84))}
                    v90 := GenS13{F0: -9223372036854775808}
                    _ = v88
                    _ = v89
                    _ = v90
                }
                switch (-(v85)) {
                case -9223372036854775808:
                    fmt.Println(v76, v3, v85)
                case 984:
                    fmt.Println(v3)
                default:
                    var v91 uint32 = 42
                    _ = v91
                }
                _ = fi84
                _ = v85
            }
        }
        v92 := GenS2{F0: 210}
        for fi93 := 0; fi93 < 19; fi93++ {
            for fi94 := 0; fi94 < 19; fi94++ {
                v2 = v2
                var v95 uint8 = 143
                _ = fi94
                _ = v95
            }
            _ = fi93
        }
        _ = v92
    }
    v16 += -1
    switch 4294967295 {
    case 32767:
        v3 -= 0
    default:
        v3 = 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
