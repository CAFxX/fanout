package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 string; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 uint8; F1 int32; F2 int }
type GenS9 struct { F0 bool; F1 string; F2 int64 }
func genFn0(p0 int32, p1 int16) int8 {
    p0 = (p0 - (p0 * p0))
    v6 := []int{128, -715, -785, 9223372036854775807}
    var v7 int = -9223372036854775808
    var v8 int64 = 42
    switch "" {
    case "hello":
        var v9 bool = false
        _ = v9
    case "foo bar":
        if false {
            v10 := GenS1{F0: "", F1: 31, F2: v8}
            v11 := GenS2{F0: 699, F1: p0}
            _ = v10
            _ = v11
        }
    case "":
        v12 := GenS1{F0: (("\n\t" + "a") + ("\n\t" + "")), F1: 0, F2: v8}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []int{-2147483648, -9223372036854775808, -2147483648, 32768}
    var v2 bool = false
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    fmt.Println(v2)
    if (255 != 212) {
        v2 = (!(255 == 255))
        v4[("foo bar" + "a")] = false
    } else {
        var v13 int16 = 86
        _ = v13
    }
    v2 = v2
    if ((42 == -762) && true) {
        v14 := GenS2{F0: 615, F1: 1}
        _ = v14
    }
    v2 = false
    var v15 uint64 = 0
    v15 = v15
    var v16 uint32 = 568
    for fi17 := 0; fi17 < 16; fi17++ {
        switch "hello" {
        case "\n\t":
            v18 := GenS3{F0: (v16 ^ 384), F1: (v16 - (v16 << ((255) & 63)))}
            _ = v18
        case "foo bar":
            v16 = 0
        case "":
            fmt.Println(v16, v15, v15)
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 uint32 = (v16 >> ((1) & 63))
        var v21 uint32 = (v16 % ((v20) | 1))
        _ = fi19
        _ = v20
        _ = v21
    }
    if (-133 != -882) {
        if (4294967295 <= 256) {
            switch 1 {
            case 9223372036854775807:
                if true {
                    fmt.Println(v15, v15, v15)
                    if ("a" < "foo bar") {
                        if (v2 && true) {
                            var v22 string = "foo bar"
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 11; fi23++ {
                                v24 := GenS1{F0: ("hello" + "hello"), F1: 92, F2: 9223372036854775807}
                                _ = fi23
                                _ = v24
                            }
                            _ = len(v3)
                        }
                        var v25 uint8 = 178
                        if v2 {
                            v3[165] = (v16 * v16)
                        }
                        _ = v25
                    } else {
                        v26 := GenS3{F0: (^(v16)), F1: v16}
                        _ = v26
                    }
                } else {
                    if (32767 != 65535) {
                        for fi27 := 0; fi27 < 10; fi27++ {
                            v28 := GenS4{F0: 0, F1: v16}
                            var v29 bool = ("hello" <= "a")
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        if (65535 != 92) {
                            fmt.Println(v2)
                        } else {
                            var v30 int32 = 335
                            if len(v1) > 0 {
                                v1[0] = -1
                            }
                            v16 = 2147483647
                            _ = v30
                        }
                    } else {
                        for fi31 := 0; fi31 < 15; fi31++ {
                            var v32 float64 = float64((1e10 * float32(1e10)))
                            var v33 uint8 = uint8(v16)
                            v32 = v32
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS5{F0: "foo bar", F1: 563, F2: 778}
                        _ = v34
                    }
                    v35 := GenS5{F0: "a", F1: -32768, F2: 434}
                    _ = v35
                }
            case 0:
                for fi36 := 0; fi36 < 8; fi36++ {
                    v15 += (v15 ^ uint64(v16))
                    for fi37 := 0; fi37 < 9; fi37++ {
                        fmt.Println(fi36)
                        _ = fi37
                    }
                    v5[(fi36 % ((-1) | 1))] = (("hello" + "") + "foo bar")
                    _ = fi36
                }
            case 2147483647:
                delete(v4, ("a" + "hello"))
            default:
                switch "\n\t" {
                case "foo bar":
                    var v38 int = 540
                    _ = v38
                case "":
                    switch 0 {
                    case 1:
                        _ = len(v1)
                    case 140:
                        v39 := GenS6{F0: 233, F1: v15, F2: 1}
                        _ = v39
                    case 4294967295:
                        switch 127 {
                        case -292:
                            switch "" {
                            case "\n\t":
                                v15 = (v15 % ((18446744073709551615) | 1))
                            case "a":
                                v16 += 1
                            default:
                                v40 := GenS5{F0: (("a" + "hello") + "hello"), F1: 1, F2: 1}
                                _ = v40
                            }
                        case -2147483648:
                            v41 := GenS7{F0: 9223372036854775807, F1: genFn0(0, 32767)}
                            _ = v41
                        }
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case "\n\t":
                    var v42 int64 = -9223372036854775808
                    _ = v42
                default:
                    fmt.Println(v2)
                }
            }
            v43 := GenS7{F0: 9223372036854775807, F1: -128}
            _ = v43
        } else {
            if (!((!v2) || ("\n\t" == "a"))) {
                var v44 uint = uint(v16)
                _ = v44
            } else {
                v15 = (v15 & 0)
            }
        }
        v16 = 0
        for fi45 := 0; fi45 < 16; fi45++ {
            v15 = (v15 + 42)
            v46 := v4[""]
            v15 = v15
            _ = fi45
            _ = v46
        }
    } else {
        for fi47 := 0; fi47 < 4; fi47++ {
            fmt.Println(fi47)
            _ = fi47
        }
        if (v2 && ((-212 < 127) || v2)) {
            if true {
                if (((^(v15)) >= v15) || (((v15 / ((18446744073709551615) | 1)) > v15) || (42 != v15))) {
                    v15 = v15
                    switch -213 {
                    case 32768:
                        for fi48 := 0; fi48 < 3; fi48++ {
                            var v49 uint8 = 128
                            _ = fi48
                            _ = v49
                        }
                    case 1:
                        if (genFn0(42, 32767) < 100) {
                            for fi50 := 0; fi50 < 19; fi50++ {
                                switch 2147483647 {
                                case 825:
                                    if ((1 >= 65535) || (!false)) {
                                        v15 += (v15 % ((42) | 1))
                                        var v51 bool = true
                                        _ = v51
                                    } else {
                                        for fi52 := 0; fi52 < 16; fi52++ {
                                            v53 := GenS8{F0: 232, F1: 0, F2: 743}
                                            _ = len(v1)
                                            var v54 bool = (1.5 < (-2.25 + 3.14159))
                                            _ = fi52
                                            _ = v53
                                            _ = v54
                                        }
                                        for fi55 := 0; fi55 < 8; fi55++ {
                                            var v56 uint = 255
                                            v57 := GenS4{F0: 821, F1: v16}
                                            _ = fi55
                                            _ = v56
                                            _ = v57
                                        }
                                    }
                                default:
                                    if false {
                                        v58 := GenS5{F0: "a", F1: 0, F2: 2}
                                        v15 = 2
                                        var v59 uint8 = 70
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                _ = fi50
                            }
                            if (128 != 42) {
                                var v60 uint8 = 42
                                var v61 int32 = 42
                                v62 := GenS9{F0: ((1 < v16) && v2), F1: ("hello" + (("" + "\n\t") + "a")), F2: -1}
                                _ = v60
                                _ = v61
                                _ = v62
                            } else {
                                if v2 {
                                    v16 = (^(v16))
                                    v63 := GenS8{F0: 255, F1: 42, F2: -9223372036854775808}
                                    _ = v63
                                }
                                for fi64 := 0; fi64 < 6; fi64++ {
                                    _ = len(v5)
                                    _ = fi64
                                }
                            }
                        } else {
                            for fi65 := 0; fi65 < 8; fi65++ {
                                var v66 string = (("" + (("\n\t" + ("a" + "foo bar")) + "foo bar")) + "")
                                v67 := GenS5{F0: v66, F1: -32768, F2: 909}
                                _ = fi65
                                _ = v66
                                _ = v67
                            }
                        }
                    case -9223372036854775808:
                        v2 = (!(32767 >= 32767))
                    default:
                        var v68 uint32 = (v16 & (^(v16)))
                        _ = v68
                    }
                }
                var v69 int8 = 127
                _ = v69
            }
        }
    }
    if ((false && v2) || false) {
        v16 += (v16 - 2)
    } else {
        switch "hello" {
        case "foo bar":
            v2 = (42 != uint16(v16))
        default:
            if (v16 != (v16 >> ((65535) & 63))) {
                v2 = (!false)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                delete(v4, "")
            } else {
                v70 := GenS6{F0: 2, F1: uint64(v16), F2: (v15 + v15)}
                _ = v70
            }
        }
        v16 = 936
        for fi71 := 0; fi71 < 8; fi71++ {
            if v2 {
                for fi72 := 0; fi72 < 7; fi72++ {
                    v16 = (^(v16))
                    _ = fi72
                }
            }
            _ = fi71
        }
    }
    v15 = v15
    var v73 uint64 = 0
    v15 = (v15 / ((v15) | 1))
    var v74 int64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v73
    _ = v74
}
