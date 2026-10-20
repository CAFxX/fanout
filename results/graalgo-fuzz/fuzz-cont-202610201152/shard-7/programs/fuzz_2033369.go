package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int16; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 bool; F1 int; F2 int16 }
type GenS8 struct { F0 int8; F1 int8 }
func genFn0(p0 uint32) uint8 {
    p0 -= (p0 & 587)
    p0 = 477
    var v6 int = -1
    _ = p0
    _ = v6
    return 151
}
func genFn1() int32 {
    if false {
        var v7 uint8 = 1
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS1{F0: 9223372036854775807, F1: "", F2: 177}
            v10 := GenS2{F0: 370}
            v7 += (^(v7))
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint32 = 567
            var v13 int = fi11
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v7
    }
    switch "a" {
    case "\n\t":
        var v14 string = ("a" + "\n\t")
        _ = v14
    case "hello":
        var v15 int = 1
        _ = v15
    default:
        var v16 int16 = 606
        _ = v16
    }
    if (!false) {
        v17 := []int{32767}
        if true {
            var v18 int32 = 229
            if true {
                for fi19 := 0; fi19 < 15; fi19++ {
                    v18 = v18
                    _ = fi19
                }
            }
            _ = v18
        }
        var v20 float64 = -2.25
        _ = v17
        _ = v20
    }
    var v21 int8 = 1
    _ = v21
    return 42
}

func main() {
    var v1 int32 = 2147483647
    v2 := []int{867}
    var v3 uint16 = 32767
    var v4 string = "a"
    v5 := make(map[string]uint32)
    var v22 int = -281
    var v23 string = (v4 + "hello")
    _ = len(v2)
    var v24 int32 = genFn1()
    v3 = 2
    if false {
        for fi25 := 0; fi25 < 15; fi25++ {
            if true {
                var v26 uint16 = 32767
                _ = v26
            } else {
                v3 -= (v3 % ((v3) | 1))
                for fi27 := 0; fi27 < 3; fi27++ {
                    var v28 uint16 = (v3 / ((831) | 1))
                    _ = fi27
                    _ = v28
                }
                var v29 uint16 = 712
                _ = v29
            }
            if (1 > 42) {
                if (false && false) {
                    if false {
                        v22 += v22
                    } else {
                        fmt.Println(v24, v22)
                        v3 -= (v3 | 65535)
                    }
                    v30 := GenS3{F0: 2, F1: v3, F2: 86}
                    fmt.Println(v24)
                    _ = v30
                } else {
                    if false {
                        var v31 uint64 = 0
                        var v32 uint16 = v3
                        _ = v31
                        _ = v32
                    } else {
                        if (!(332 > 0)) {
                            var v33 uint16 = (v3 & v3)
                            if false {
                                fmt.Println(v23, v3, v23)
                                for fi34 := 0; fi34 < 9; fi34++ {
                                    switch v4 {
                                    case "\n\t":
                                        v4 = ((("foo bar" + v23) + "\n\t") + v4)
                                    default:
                                        v3 -= (v33 + 2)
                                    }
                                    v35 := GenS1{F0: 0, F1: v23, F2: 866}
                                    _ = fi34
                                    _ = v35
                                }
                            } else {
                                fmt.Println(fi25)
                                var v36 uint32 = 2
                                _ = v36
                            }
                            _ = v33
                        } else {
                            var v37 uint = 255
                            _ = v37
                        }
                        var v38 uint32 = 4294967295
                        _ = v38
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi39 := 0; fi39 < 10; fi39++ {
                        switch v4 {
                        case "a":
                            v3 += (v3 * v3)
                        case "\n\t":
                            var v40 uint32 = 4294967295
                            if len(v2) > 0 {
                                v2[0] = 32767
                            }
                            v23 = (v23 + (v4 + (v23 + v4)))
                            _ = v40
                        case "foo bar":
                            v41 := GenS2{F0: v24}
                            _ = v41
                        }
                        var v42 uint32 = 0
                        switch (v4 + v23) {
                        case "foo bar":
                            v4 = (v23 + ((v4 + ("" + v4)) + "a"))
                        case "\n\t":
                            v43 := v5[((v4 + v4) + (v4 + v4))]
                            _ = v43
                        case "hello":
                            v22 -= (fi25 << ((4294967295) & 63))
                        }
                        _ = fi39
                        _ = v42
                    }
                }
                v44 := GenS3{F0: 1, F1: (v3 << ((2) & 63)), F2: 1}
                _ = v44
            } else {
                v45 := v5[v4]
                var v46 uint8 = 128
                _ = v45
                _ = v46
            }
            fmt.Println(v22)
            _ = fi25
        }
        v47 := GenS3{F0: 465, F1: (v3 | (^(v3))), F2: 2}
        _ = v47
    }
    v3 = (v3 >> ((65535) & 63))
    v24 -= (v24 % (((v1 >> ((1) & 63))) | 1))
    for fi48 := 0; fi48 < 13; fi48++ {
        v49 := GenS3{F0: 529, F1: 231, F2: 255}
        v23 = v4
        v22 = v22
        _ = fi48
        _ = v49
    }
    var v50 uint16 = v3
    if (-545 != 32767) {
        v23 = "\n\t"
        for fi51 := 0; fi51 < 12; fi51++ {
            for fi52 := 0; fi52 < 8; fi52++ {
                v53 := GenS3{F0: 401, F1: (v3 ^ v50), F2: 128}
                v50 -= v3
                fmt.Println(fi51, v1)
                v54 := GenS2{F0: 525}
                v4 = "\n\t"
                var v55 int16 = -189
                if true {
                    var v56 uint8 = 127
                    _ = len(v5)
                    _ = v56
                }
                _ = fi52
                _ = v53
                _ = v54
                _ = v55
            }
            if true {
                var v57 string = ((((v23 + "hello") + v23) + v4) + "a")
                var v58 int64 = 847
                var v59 bool = false
                _ = v57
                _ = v58
                _ = v59
            } else {
                for fi60 := 0; fi60 < 5; fi60++ {
                    v61 := GenS3{F0: 1, F1: v3, F2: 1}
                    v62 := GenS4{F0: -962, F1: -1, F2: 665}
                    var v63 int64 = 42
                    _ = fi60
                    _ = v61
                    _ = v62
                    _ = v63
                }
                switch v22 {
                case -1:
                    v50 = (v50 >> ((1) & 63))
                default:
                    v24 -= (v1 << ((65535) & 63))
                }
            }
            var v64 bool = true
            _ = fi51
            _ = v64
        }
    } else {
        if ((false && false) || (!(!(v22 >= 1)))) {
            for fi65 := 0; fi65 < 20; fi65++ {
                if ((!(42 == 43)) || true) {
                    _ = len(v5)
                    v50 -= (v3 - 551)
                }
                _ = fi65
            }
            var v66 uint64 = 65
            _ = v66
        }
    }
    v1 += (v1 % ((v24) | 1))
    if (!(808 == 820)) {
        v22 = (v22 * (v22 + (v22 >> ((651) & 63))))
        switch v23 {
        case "foo bar":
            var v67 uint16 = 32767
            _ = v67
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = (v22 >> ((1) & 63))
            }
        default:
            var v68 uint32 = 4294967295
            _ = v68
        }
        if (!(!true)) {
            switch (v4 + (v23 + v4)) {
            case "foo bar":
                switch v23 {
                case "":
                    if (true && false) {
                        v69 := GenS3{F0: 4294967295, F1: (v50 + (^(v50))), F2: 42}
                        if (((float32(3.14159) * 1e10) / 1e10) < (-2.25 + ((((3.14159 + 1.5) / 1.0) - (3.14159 * 3.14159)) - (1e10 * 1e10)))) {
                            if (!true) {
                                v70 := GenS5{F0: 1, F1: 839}
                                v23 = v4
                                if len(v2) > 0 {
                                    v2[0] = (v22 + (v22 * v22))
                                }
                                _ = v70
                            }
                        } else {
                            var v71 int = (v22 >> ((42) & 63))
                            _ = v71
                        }
                        _ = v69
                    }
                case "a":
                    for fi72 := 0; fi72 < 8; fi72++ {
                        if ("foo bar" > "a") {
                            _ = len(v2)
                            if false {
                                v73 := GenS4{F0: -32768, F1: -483, F2: 18446744073709551615}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v73
                            }
                        } else {
                            v74 := GenS6{F0: true, F1: 153}
                            fmt.Println(v3, v22)
                            if (v24 > (^(v24))) {
                                v22 = 0
                                if (0 > -1) {
                                    v23 = ("a" + (v23 + v23))
                                    v75 := v5[v4]
                                    _ = v75
                                } else {
                                    v76 := GenS3{F0: 18446744073709551615, F1: (^(v50)), F2: 1}
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v76
                                }
                                v77 := GenS7{F0: ((1.5 / 1.0) > (3.14159 + float64(3.14159))), F1: fi72, F2: 448}
                                _ = v77
                            } else {
                                v23 = v23
                            }
                            _ = v74
                        }
                        switch (v22 * v22) {
                        case -2147483648:
                            fmt.Println(v3, v4)
                        }
                        _ = fi72
                    }
                case "foo bar":
                    v24 = (v24 << ((65535) & 63))
                default:
                    switch v22 {
                    case 42:
                        v23 = v23
                    case -9223372036854775808:
                        v78 := GenS7{F0: (v50 != 1), F1: v22, F2: -51}
                        _ = v78
                    }
                }
            }
        } else {
            v79 := GenS8{F0: -128, F1: 100}
            var v80 int16 = int16(v50)
            _ = v79
            _ = v80
        }
    }
    var v81 float64 = 0.1
    var v82 int64 = -341
    v3 = (^(v3))
    var v83 int64 = (v82 >> ((42) & 63))
    v84 := GenS6{F0: (float64((3.14159 / 1e10)) != 3.14159), F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v81
    _ = v82
    _ = v83
    _ = v84
}
