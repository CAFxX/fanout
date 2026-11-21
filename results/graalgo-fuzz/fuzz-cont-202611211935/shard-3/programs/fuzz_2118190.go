package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int64; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 uint8 }
type GenS9 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS10 struct { F0 uint8 }

func main() {
    var v1 uint16 = 32767
    var v2 int32 = 136
    v3 := make(map[uint32]string)
    var v4 int32 = (v2 - v2)
    v5 := []int{0, 4294967295, 127, 4294967295}
    var v6 int = 128
    for fi7 := 0; fi7 < 14; fi7++ {
        switch ((("foo bar" + "") + "hello") + "hello") {
        case "a":
            v2 -= (v4 % ((v2) | 1))
        }
        switch "a" {
        case "hello":
            v6 = (fi7 & v6)
        case "\n\t":
            switch v6 {
            case 128:
                fmt.Println(v4, v1, v1)
            case -2147483648:
                v1 = v1
            case 127:
                for fi8 := 0; fi8 < 2; fi8++ {
                    v2 += 0
                    v9 := v3[2]
                    if (2 >= 0) {
                        if false {
                            if true {
                                var v10 uint32 = 2
                                _ = v10
                            } else {
                                var v11 int64 = 321
                                fmt.Println(v6)
                                if (-49 != 100) {
                                    if (false && false) {
                                        fmt.Println(v4, v9, v11)
                                        var v12 int = fi8
                                        if (32767 != 211) {
                                            switch v6 {
                                            case 2147483647:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            case 9223372036854775807:
                                                v13 := GenS1{F0: v4, F1: (v12 | -2147483648)}
                                                _ = v13
                                            case -238:
                                                v14 := GenS1{F0: (v4 - -2147483648), F1: 4294967295}
                                                _ = v14
                                            default:
                                                fmt.Println(v11, v2)
                                            }
                                        }
                                        _ = v12
                                    } else {
                                        v15 := GenS1{F0: 172, F1: (v6 & fi7)}
                                        if (!false) {
                                            fmt.Println(v6, v4, fi8)
                                            var v16 uint64 = 0
                                            fmt.Println(v16, v2, v2)
                                            _ = v16
                                        }
                                        fmt.Println(v9, v11)
                                        var v17 int32 = v2
                                        _ = v15
                                        _ = v17
                                    }
                                    var v18 int = fi8
                                    v18 = (v18 ^ 862)
                                    _ = v18
                                } else {
                                    if (127 > -1) {
                                        v19 := GenS2{F0: v11, F1: (^(v11)), F2: 328}
                                        var v20 uint16 = v1
                                        _ = v19
                                        _ = v20
                                    }
                                }
                                _ = v11
                            }
                        } else {
                            v2 += 0
                            var v21 float64 = 1.5
                            _ = v21
                        }
                        v22 := GenS3{F0: 502, F1: v1, F2: 1}
                        switch (fi8 ^ fi8) {
                        case -1:
                            if (39 > 32767) {
                                if (false && (!(v1 > (v1 * (^(v1)))))) {
                                    fmt.Println(v2, v1)
                                    v23 := GenS4{F0: 1000}
                                    _ = v23
                                }
                            }
                        case 268:
                            if false {
                                var v24 uint16 = 42
                                _ = v24
                            } else {
                                if (false || (!(2147483648 >= 0))) {
                                    var v25 uint = 18446744073709551615
                                    _ = v25
                                }
                                v1 = (^(v1))
                            }
                        case -247:
                            switch (fi7 << ((834) & 63)) {
                            case 2147483647:
                                v26 := GenS1{F0: 0, F1: fi8}
                                _ = v26
                            default:
                                v27 := GenS3{F0: 42, F1: v1, F2: 65535}
                                _ = v27
                            }
                        default:
                            if false {
                                v28 := GenS5{F0: v9}
                                _ = v28
                            } else {
                                if true {
                                    if true {
                                        switch 127 {
                                        case 42:
                                            _ = len(v3)
                                        case 4294967295:
                                            if false {
                                                if true {
                                                    v6 = 32767
                                                    if len(v5) > 0 {
                                                        v5[0] = -2147483648
                                                    }
                                                    v29 := GenS6{F0: 127, F1: (v1 >> ((349) & 63)), F2: 1}
                                                    _ = v29
                                                } else {
                                                    _ = len(v5)
                                                }
                                                v30 := GenS2{F0: 9223372036854775807, F1: 42, F2: 15}
                                                v1 = (v1 << ((2) & 63))
                                                var v31 uint = 422
                                                v5 = append(v5, (v6 - -9223372036854775808))
                                                _ = v30
                                                _ = v31
                                            } else {
                                                v9 = (v9 + ("hello" + v9))
                                                if (false && ((float32(((0.0 - (1.5 * (1.5 * float64((((-2.25 + 0.1) / (0.1 - (((1e10 / (float32(float64(0.1)) / (0.1 - 1e10))) * 0.0) * 3.14159))) - ((0.1 - (1e10 / 1.5)) + float32(float64(-2.25)))))))) + 1e10)) / 0.1) >= ((0.0 + (float32(0.1) / 0.1)) / float32(((((1e10 * -2.25) / 1e10) - 0.1) / ((float64(0.1) - float64(3.14159)) / 1.0)))))) {
                                                    var v32 int = fi7
                                                    delete(v3, 2)
                                                    if true {
                                                        if (v2 >= v2) {
                                                            _ = len(v3)
                                                            switch fi8 {
                                                            case 9223372036854775807:
                                                                v5 = append(v5, -847)
                                                            }
                                                        }
                                                    }
                                                    _ = v32
                                                }
                                                var v33 float32 = float32(float64((0.1 - 3.14159)))
                                                _ = v33
                                            }
                                        case 127:
                                            v2 = -1
                                        }
                                    }
                                    if (4294967295 != 18446744073709551615) {
                                        if false {
                                            fmt.Println(v2)
                                            fmt.Println(v6, fi8)
                                            var v34 float32 = -2.25
                                            var v35 int16 = 42
                                            _ = v34
                                            _ = v35
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v36 uint8 = 128
                                        _ = v36
                                    } else {
                                        v6 = fi7
                                    }
                                    var v37 int16 = 774
                                    _ = v37
                                } else {
                                    delete(v3, 1)
                                    v2 = (^(v2))
                                }
                                v38 := GenS3{F0: 2, F1: (v1 - 0), F2: uint(v1)}
                                v3[2147483648] = "hello"
                                switch v9 {
                                case "foo bar":
                                    v6 = fi8
                                case "":
                                    v39 := v3[545]
                                    _ = v39
                                case "a":
                                    v9 = ("" + v9)
                                default:
                                    fmt.Println(fi7)
                                }
                                v4 += (v2 % (((v2 - (v2 & v2))) | 1))
                                switch fi8 {
                                case -9223372036854775808:
                                    v40 := GenS6{F0: 102, F1: 0, F2: 2}
                                    _ = v40
                                case -1:
                                    if false {
                                        var v41 uint64 = 199
                                        if true {
                                            v42 := GenS4{F0: 464}
                                            _ = v42
                                        }
                                        if (!true) {
                                            v2 += (v2 / ((v2) | 1))
                                        } else {
                                            v43 := v3[2147483648]
                                            if false {
                                                v44 := GenS5{F0: v9}
                                                _ = v44
                                            } else {
                                                switch "foo bar" {
                                                case "a":
                                                    var v45 float32 = 0.0
                                                    _ = v45
                                                case "":
                                                    v41 -= (v41 >> ((582) & 63))
                                                default:
                                                    if false {
                                                        v43 = ((v43 + v9) + "foo bar")
                                                        if false {
                                                            var v46 uint64 = (v41 + v41)
                                                            v47 := GenS7{F0: 9223372036854775807}
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            if true {
                                                                var v48 uint8 = 1
                                                                var v49 uint32 = uint32(v48)
                                                                var v50 int = 128
                                                                if true {
                                                                    v51 := GenS3{F0: (v41 << ((65535) & 63)), F1: v1, F2: 42}
                                                                    v52 := GenS8{F0: 2, F1: (v48 >> ((42) & 63))}
                                                                    var v53 bool = (fi7 >= (-(fi8)))
                                                                    _ = v51
                                                                    _ = v52
                                                                    _ = v53
                                                                }
                                                                _ = v48
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                fmt.Println(v43, v9, fi8)
                                                                delete(v3, 4294967295)
                                                            }
                                                            var v54 int8 = -128
                                                            _ = v54
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v43
                                        }
                                        _ = v41
                                    } else {
                                        v55 := GenS9{F0: 0, F1: (^(v1)), F2: 42}
                                        _ = v55
                                    }
                                    switch fi8 {
                                    case -9223372036854775808:
                                        _ = len(v3)
                                    case 32767:
                                        var v56 int32 = (v2 - (^(v2)))
                                        _ = v56
                                    default:
                                        var v57 bool = true
                                        _ = v57
                                    }
                                }
                                v58 := GenS5{F0: ""}
                                _ = len(v3)
                                _ = v38
                                _ = v58
                            }
                        }
                        _ = v22
                    }
                    _ = fi8
                    _ = v9
                }
            }
        default:
            var v59 float64 = (-2.25 + float64(-2.25))
            _ = v59
        }
        _ = fi7
    }
    var v60 int16 = -1
    for fi61 := 0; fi61 < 15; fi61++ {
        fmt.Println(v4, v60, v2)
        switch int(v2) {
        case 1:
            v62 := v3[734]
            _ = v62
        case 495:
            fmt.Println(v60, fi61, fi61)
        }
        switch ("" + ("\n\t" + "foo bar")) {
        case "foo bar":
            var v63 int64 = 967
            _ = v63
        case "a":
            v64 := GenS4{F0: (^(v60))}
            _ = v64
        default:
            for fi65 := 0; fi65 < 16; fi65++ {
                fmt.Println(fi65, v4, v6)
                _ = fi65
            }
        }
        _ = fi61
    }
    v66 := GenS10{F0: 42}
    for fi67 := 0; fi67 < 1; fi67++ {
        var v68 int32 = (v4 % ((v4) | 1))
        v2 = v4
        _ = fi67
        _ = v68
    }
    v2 = (v4 * -269)
    v1 = v1
    for fi69 := 0; fi69 < 9; fi69++ {
        fmt.Println(v6, v60)
        v60 = -1
        if (((v1 / ((42) | 1)) >= (^(v1))) && (!true)) {
            v5 = append(v5, 42)
        } else {
            var v70 string = ""
            if false {
                switch -174 {
                case 32767:
                    var v71 string = ("" + "hello")
                    _ = v71
                case -9:
                    var v72 uint = 211
                    _ = v72
                default:
                    var v73 float64 = (((-2.25 - 3.14159) * 3.14159) / 1.0)
                    _ = v73
                }
            }
            v1 = (v1 & v1)
            _ = v70
        }
        _ = fi69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v66
}
