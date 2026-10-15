package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 uint64; F2 string }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 string; F1 int; F2 int }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint32 = 42
    var v3 uint32 = (v2 % ((uint32(v1)) | 1))
    var v4 uint32 = (v2 - v2)
    v5 := GenS1{F0: 667, F1: ((true || (((("" + "\n\t") + "") != "hello") || ((true && false) || true))) || false)}
    var v6 uint32 = (v3 + v4)
    v7 := GenS1{F0: 9223372036854775807, F1: (1 > 19)}
    var v8 int32 = 0
    if ("\n\t" != ("" + "foo bar")) {
        var v9 float32 = float32(1.5)
        if false {
            var v10 float64 = 3.14159
            v4 = (v2 & (v6 >> (((^(v1))) & 63)))
            if true {
                var v11 bool = true
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 uint32 = v4
                    v4 += uint32(v1)
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 8; fi14++ {
                    v3 -= (v4 - 85)
                    _ = fi14
                }
                _ = v11
            }
            _ = v10
        }
        if (v1 != 863) {
            if ((4294967295 >= 42) && false) {
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := []int{(fi15 | fi15), (fi15 + 2147483647), fi15}
                    _ = fi15
                    _ = v16
                }
                v17 := GenS2{F0: 32767}
                for fi18 := 0; fi18 < 4; fi18++ {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        switch (("a" + "\n\t") + "\n\t") {
                        case "foo bar":
                            if ((v3 ^ (v6 | v3)) >= (v3 - (v2 % (((^(v2))) | 1)))) {
                                var v20 uint16 = 42
                                switch "" {
                                case "foo bar":
                                    var v21 bool = (false || true)
                                    _ = v21
                                }
                                if ("foo bar" < ("\n\t" + "a")) {
                                    var v22 int64 = 1
                                    _ = v22
                                } else {
                                    switch (^(fi18)) {
                                    case 32767:
                                        v4 = 1
                                    case 123:
                                        var v23 bool = false
                                        _ = v23
                                    case 0:
                                        switch (("\n\t" + "\n\t") + ("" + "foo bar")) {
                                        case "hello":
                                            var v24 bool = (!true)
                                            _ = v24
                                        case "":
                                            fmt.Println(v3, v8, v20)
                                        default:
                                            var v25 int = -518
                                            _ = v25
                                        }
                                    default:
                                        switch "hello" {
                                        case "foo bar":
                                            v26 := GenS2{F0: 1000}
                                            _ = v26
                                        case "a":
                                            v27 := GenS3{F0: 143, F1: fi18}
                                            _ = v27
                                        default:
                                            fmt.Println(v20, v20)
                                        }
                                    }
                                    var v28 bool = (!((3.14159 / (((1e10 * -2.25) - (0.1 * -2.25)) * float64((v9 - (v9 - (v9 / (v9 - 3.14159))))))) == -2.25))
                                    if ((v20 * v20) != v20) {
                                        v29 := []int{}
                                        _ = v29
                                    }
                                    _ = v28
                                }
                                _ = v20
                            } else {
                                v9 -= 0.1
                                var v30 float64 = 3.14159
                                _ = v30
                            }
                        case "hello":
                            v6 = (v2 | (v4 | 2147483647))
                        case "":
                            v3 += v2
                        default:
                            switch (fi19 % ((fi18) | 1)) {
                            case 127:
                                if (-9223372036854775808 != 0) {
                                    var v31 int32 = v8
                                    v32 := GenS1{F0: -9223372036854775808, F1: (2 < 534)}
                                    var v33 uint8 = 2
                                    var v34 uint64 = 2
                                    var v35 int8 = 42
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                } else {
                                    v36 := []int{}
                                    _ = v36
                                }
                            case 1:
                                fmt.Println(v4, v6)
                            case 32767:
                                var v37 uint8 = 252
                                _ = v37
                            }
                        }
                        _ = fi19
                    }
                    _ = fi18
                }
                _ = v17
            }
            v38 := GenS2{F0: -342}
            var v39 uint = v1
            _ = v38
            _ = v39
        } else {
            var v40 int16 = -32768
            if true {
                if (389 == -9223372036854775808) {
                    for fi41 := 0; fi41 < 1; fi41++ {
                        var v42 bool = false
                        v2 = (^(v4))
                        _ = fi41
                        _ = v42
                    }
                    v43 := GenS4{F0: (v4 << ((v1) & 63)), F1: 255}
                    var v44 uint16 = 34
                    _ = v43
                    _ = v44
                }
            } else {
                var v45 uint32 = 838
                _ = v45
            }
            if false {
                var v46 uint8 = 0
                if false {
                    v4 -= (v3 & (v2 % (((v6 - 609)) | 1)))
                    fmt.Println(v4)
                } else {
                    v2 += v6
                    v47 := GenS5{F0: 1}
                    _ = v47
                }
                if (v8 <= v8) {
                    for fi48 := 0; fi48 < 17; fi48++ {
                        v49 := []int{}
                        _ = fi48
                        _ = v49
                    }
                } else {
                    if true {
                        v50 := GenS5{F0: 0}
                        _ = v50
                    } else {
                        v46 -= (v46 | (v46 / ((v46) | 1)))
                        var v51 uint = (v1 >> ((18446744073709551615) & 63))
                        _ = v51
                    }
                    v52 := GenS2{F0: v40}
                    _ = v52
                }
                _ = v46
            } else {
                for fi53 := 0; fi53 < 18; fi53++ {
                    var v54 int = 32768
                    _ = fi53
                    _ = v54
                }
                var v55 int32 = 0
                for fi56 := 0; fi56 < 12; fi56++ {
                    v8 += 816
                    _ = fi56
                }
                _ = v55
            }
            _ = v40
        }
        _ = v9
    } else {
        if false {
            var v57 float32 = (3.14159 - 3.14159)
            switch -9223372036854775808 {
            case 42:
                if true {
                    v4 = (v6 ^ (v3 % ((v4) | 1)))
                    v4 += v4
                } else {
                    var v58 uint32 = (v3 - (^(v6)))
                    _ = v58
                }
            case 930:
                var v59 string = ("" + "")
                _ = v59
            }
            for fi60 := 0; fi60 < 10; fi60++ {
                for fi61 := 0; fi61 < 4; fi61++ {
                    var v62 bool = false
                    v63 := []int{(fi60 - (fi61 * (fi61 << (((v1 ^ 65535)) & 63)))), fi61, -30}
                    if true {
                        fmt.Println(v4, fi60)
                        if v62 {
                            fmt.Println(v62)
                            switch "" {
                            case "":
                                fmt.Println(v8)
                            case "a":
                                if false {
                                    if true {
                                        v64 := GenS1{F0: 0, F1: (134 > 1)}
                                        switch ("hello" + ("\n\t" + "a")) {
                                        case "foo bar":
                                            v3 = (v2 ^ 65535)
                                        default:
                                            v65 := GenS6{F0: 128, F1: 4294967295, F2: ""}
                                            _ = v65
                                        }
                                        v1 = v1
                                        if (!(!((1 < 127) && (337 != 212)))) {
                                            var v66 float64 = (3.14159 / 1.0)
                                            _ = v66
                                        }
                                        if (!v62) {
                                            var v67 bool = (!(float64(v57) == (3.14159 / 0.1)))
                                            fmt.Println(v62, fi60)
                                            switch (fi60 % ((fi61) | 1)) {
                                            case 9223372036854775807:
                                                var v68 int32 = (v8 ^ (v8 % ((v8) | 1)))
                                                _ = v68
                                            default:
                                                _ = len(v63)
                                            }
                                            _ = v67
                                        }
                                        _ = v64
                                    }
                                    var v69 uint = v1
                                    _ = v69
                                }
                            case "\n\t":
                                var v70 uint16 = 65535
                                _ = v70
                            default:
                                v71 := GenS7{F0: 271, F1: (fi61 % ((fi60) | 1))}
                                _ = v71
                            }
                            var v72 int64 = -9223372036854775808
                            _ = v72
                        } else {
                            switch ("a" + "") {
                            case "a":
                                v62 = false
                            case "":
                                v3 = v6
                            }
                            if false {
                                v73 := GenS8{F0: "hello", F1: 42, F2: fi61}
                                _ = v73
                            } else {
                                var v74 string = "\n\t"
                                var v75 int8 = 0
                                v8 -= int32(v75)
                                var v76 string = v74
                                v77 := GenS7{F0: 400, F1: (fi61 >> ((v1) & 63))}
                                v62 = v62
                                _ = v74
                                _ = v75
                                _ = v76
                                _ = v77
                            }
                            var v78 int64 = int64(v2)
                            var v79 int64 = (v78 << (((v1 << ((65535) & 63))) & 63))
                            v6 = (v4 << (((v1 >> (((v1 ^ (v1 >> ((0) & 63)))) & 63))) & 63))
                            _ = v78
                            _ = v79
                        }
                    } else {
                        var v80 int16 = -342
                        _ = len(v63)
                        _ = v80
                    }
                    v2 = v4
                    _ = fi61
                    _ = v62
                    _ = v63
                }
                var v81 int8 = -128
                _ = fi60
                _ = v81
            }
            _ = v57
        }
        v82 := []int{32768, 161, 32767, -2147483648}
        _ = v82
    }
    v83 := []int{849, 42}
    for fi84 := 0; fi84 < 9; fi84++ {
        var v85 uint64 = 2
        v1 = 733
        _ = fi84
        _ = v85
    }
    switch "foo bar" {
    case "\n\t":
        var v86 string = ("\n\t" + ("hello" + "\n\t"))
        _ = v86
    case "hello":
        for fi87 := 0; fi87 < 3; fi87++ {
            _ = len(v83)
            v1 = v1
            _ = fi87
        }
    case "foo bar":
        _ = len(v83)
    }
    if (false || true) {
        fmt.Println(v1, v1)
        v4 -= 4294967295
        if false {
            var v88 uint = (^(v1))
            v8 = -2147483648
            _ = v88
        } else {
            if false {
                for fi89 := 0; fi89 < 12; fi89++ {
                    fmt.Println(v1, v6, v2)
                    _ = fi89
                }
                fmt.Println(v4)
            }
        }
    } else {
        v90 := GenS7{F0: 65535, F1: 32767}
        v3 += v2
        _ = v90
    }
    var v91 int16 = 1000
    for fi92 := 0; fi92 < 5; fi92++ {
        if true {
            if len(v83) > 0 {
                v83[0] = (fi92 + (fi92 + -71))
            }
        }
        _ = fi92
    }
    var v93 uint64 = 42
    fmt.Println(v8, v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v83:", len(v83), v83)
    fmt.Println("v91:", v91)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v83
    _ = v91
    _ = v93
}
