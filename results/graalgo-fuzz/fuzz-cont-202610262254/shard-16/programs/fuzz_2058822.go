package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 uint }
func genFn0(p0 int, p1 uint8) uint64 {
    for fi4 := 0; fi4 < 4; fi4++ {
        for fi5 := 0; fi5 < 6; fi5++ {
            v6 := GenS1{F0: 373, F1: 0}
            _ = fi5
            _ = v6
        }
        var v7 string = "foo bar"
        v7 = v7
        _ = fi4
        _ = v7
    }
    p1 += (p1 & 42)
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint64, p1 int8, p2 uint64) int64 {
    v8 := GenS1{F0: 1, F1: 32767}
    var v9 uint = 42
    var v10 int = 127
    var v11 string = (("foo bar" + "foo bar") + "foo bar")
    v9 -= (^(v9))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 261
}

func main() {
    var v1 int32 = 427
    var v2 uint8 = 2
    var v3 uint64 = 1
    var v12 int = 0
    v13 := []int{v12}
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: 0, F1: 42}
        if (((!(!(!(42 == v3)))) && (4294967295 > 4294967295)) && true) {
            for fi16 := 0; fi16 < 9; fi16++ {
                v17 := GenS1{F0: 42, F1: -1}
                v3 = genFn0(42, (v2 - (^(v2))))
                if false {
                    v18 := GenS1{F0: 2147483648, F1: 64}
                    _ = v18
                } else {
                    if false {
                        v1 -= v1
                        var v19 uint8 = v2
                        v1 += (v1 >> ((959) & 63))
                        _ = v19
                    } else {
                        var v20 int = 9223372036854775807
                        _ = v20
                    }
                    if true {
                        if len(v13) > 0 {
                            v13[0] = (v12 & fi14)
                        }
                    } else {
                        fmt.Println(v2, v12)
                        var v21 float64 = ((-2.25 * -2.25) + (1.5 + 1e10))
                        var v22 int64 = -9223372036854775808
                        _ = v21
                        _ = v22
                    }
                    if (0 < 2) {
                        switch (("\n\t" + ((("" + "") + "hello") + "")) + ("a" + "a")) {
                        case "foo bar":
                            if false {
                                switch (fi14 ^ -988) {
                                case -495:
                                    v23 := GenS2{F0: 91}
                                    _ = v23
                                default:
                                    if (!false) {
                                        var v24 uint16 = 255
                                        _ = v24
                                    }
                                }
                                if (false || false) {
                                    switch ("foo bar" + "foo bar") {
                                    case "a":
                                        fmt.Println(fi16, v1, fi16)
                                    case "hello":
                                        if false {
                                            v12 = v12
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            fmt.Println(fi14, v12, v3)
                                        }
                                    default:
                                        v25 := GenS1{F0: 1, F1: 0}
                                        _ = v25
                                    }
                                    var v26 uint16 = 255
                                    _ = v26
                                } else {
                                    v1 += -10
                                    switch "foo bar" {
                                    case "":
                                        v2 = v2
                                    }
                                }
                            } else {
                                fmt.Println(v3)
                            }
                        }
                        v1 = -1
                    } else {
                        fmt.Println(v3, v2)
                    }
                    if ((65535 >= 447) && false) {
                        v2 += 189
                    } else {
                        var v27 string = "hello"
                        v28 := GenS1{F0: 2147483647, F1: 1}
                        var v29 bool = false
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                }
                _ = fi16
                _ = v17
            }
            fmt.Println(v2, fi14)
            v12 -= -2147483648
        }
        _ = fi14
        _ = v15
    }
    if (!false) {
        var v30 int = (v12 & v12)
        fmt.Println(v1, v3, v12)
        for fi31 := 0; fi31 < 19; fi31++ {
            v32 := GenS1{F0: 2147483647, F1: 42}
            v2 += 138
            _ = fi31
            _ = v32
        }
        _ = v30
    } else {
        v33 := GenS1{F0: 2, F1: 0}
        _ = v33
    }
    if (!false) {
        for fi34 := 0; fi34 < 10; fi34++ {
            v35 := GenS1{F0: 4294967295, F1: -635}
            _ = len(v13)
            var v36 int64 = -994
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    fmt.Println(v3)
    var v37 bool = true
    v12 = 127
    v38 := GenS2{F0: v3}
    for fi39 := 0; fi39 < 14; fi39++ {
        for fi40 := 0; fi40 < 16; fi40++ {
            if len(v13) > 0 {
                v13[0] = fi39
            }
            v41 := GenS2{F0: v3}
            if len(v13) > 0 {
                v13[0] = fi40
            }
            _ = fi40
            _ = v41
        }
        v2 = (v2 ^ uint8(v3))
        _ = fi39
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        if v37 {
            if v37 {
                _ = len(v13)
                v43 := GenS3{F0: 1, F1: 2, F2: 0}
                _ = v43
            }
        } else {
            var v44 uint16 = 76
            _ = v44
        }
        if (!(!v37)) {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        }
        if (!v37) {
            fmt.Println(v37)
            if (-52 <= fi42) {
                var v45 float64 = ((1.5 / 0.1) + (((float64(1e10) / 0.1) / 0.1) / 1.0))
                for fi46 := 0; fi46 < 15; fi46++ {
                    if false {
                        switch -9223372036854775808 {
                        case -9223372036854775808:
                            _ = len(v13)
                        }
                        fmt.Println(fi46)
                        switch 4294967295 {
                        case 42:
                            var v47 float64 = v45
                            _ = v47
                        default:
                            if v37 {
                                v1 = v1
                                v48 := GenS3{F0: 65535, F1: 11, F2: 0}
                                _ = v48
                            }
                        }
                    }
                    _ = fi46
                }
                _ = v45
            } else {
                var v49 int = -617
                fmt.Println(v1, v49, v3)
                _ = v49
            }
            if len(v13) > 0 {
                v13[0] = fi42
            }
        }
        _ = fi42
    }
    var v50 int = (v12 / ((v12) | 1))
    if (4294967295 <= 760) {
        for fi51 := 0; fi51 < 14; fi51++ {
            if v37 {
                for fi52 := 0; fi52 < 20; fi52++ {
                    fmt.Println(v50)
                    v1 -= (v1 + 761)
                    v13 = append(v13, 138)
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    var v53 int8 = -128
                    _ = fi52
                    _ = v53
                }
                fmt.Println(v1, v1, v37)
                var v54 uint32 = 1
                _ = v54
            } else {
                var v55 int64 = -1
                _ = len(v13)
                var v56 float32 = 1e10
                _ = v55
                _ = v56
            }
            var v57 int64 = -213
            v13 = append(v13, (v50 - -2147483648))
            _ = fi51
            _ = v57
        }
        if v37 {
            if v37 {
                if v37 {
                    for fi58 := 0; fi58 < 14; fi58++ {
                        v12 = (v12 | (^(v50)))
                        v59 := GenS4{F0: 65535, F1: v50}
                        v60 := GenS4{F0: 18446744073709551615, F1: (v12 % ((v12) | 1))}
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                } else {
                    v61 := GenS4{F0: 2, F1: (v12 << ((655) & 63))}
                    v13 = append(v13, v50)
                    switch (v12 << ((4294967295) & 63)) {
                    case 128:
                        if (2 <= 255) {
                            var v62 int32 = -529
                            for fi63 := 0; fi63 < 16; fi63++ {
                                v1 += v62
                                _ = fi63
                            }
                            _ = v62
                        } else {
                            v37 = (-1 != 0)
                            for fi64 := 0; fi64 < 9; fi64++ {
                                var v65 int16 = 42
                                var v66 int64 = -9223372036854775808
                                var v67 bool = (1 < 127)
                                _ = fi64
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                            for fi68 := 0; fi68 < 17; fi68++ {
                                v50 = fi68
                                v50 = (v12 + (^(v50)))
                                var v69 int16 = -185
                                _ = fi68
                                _ = v69
                            }
                        }
                    case -2147483648:
                        v70 := GenS5{F0: 255}
                        _ = v70
                    case 9223372036854775807:
                        switch (v12 - v50) {
                        case 4294967295:
                            var v71 uint32 = 65535
                            _ = v71
                        }
                    }
                    _ = v61
                }
            } else {
                if (100 > 127) {
                    fmt.Println(v2, v37, v1)
                    var v72 string = ""
                    _ = len(v13)
                    _ = v72
                } else {
                    v12 = v50
                }
            }
            var v73 bool = v37
            _ = v73
        }
    } else {
        v2 -= v2
    }
    if len(v13) > 0 {
        v13[0] = (v12 << ((18446744073709551615) & 63))
    }
    if ((!v37) && (42 != v2)) {
        v50 = v12
    } else {
        for fi74 := 0; fi74 < 9; fi74++ {
            switch v50 {
            case -799:
                switch "foo bar" {
                case "hello":
                    if len(v13) > 0 {
                        v13[0] = (-(fi74))
                    }
                }
            case 9223372036854775807:
                v3 = v3
            }
            _ = fi74
        }
    }
    var v75 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v50:", v50)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v50
    _ = v75
}
