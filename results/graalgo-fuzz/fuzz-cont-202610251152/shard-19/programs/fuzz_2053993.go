package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 int; F1 int; F2 int8 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16; F1 int32; F2 int8 }
type GenS9 struct { F0 bool; F1 int64 }
func genFn0(p0 int8) uint16 {
    v4 := []int{2147483647, 704, -2147483648}
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 float64 = (float64(-2.25) - -2.25)
        if false {
            var v7 int32 = 0
            v8 := GenS1{F0: 42, F1: "hello", F2: 878}
            _ = v7
            _ = v8
        }
        if (p0 == (p0 | (p0 | p0))) {
            for fi9 := 0; fi9 < 13; fi9++ {
                if len(v4) > 0 {
                    v4[0] = fi9
                }
                _ = fi9
            }
        } else {
            var v10 uint = 1
            _ = v10
        }
        _ = fi5
        _ = v6
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        p0 = (p0 >> ((741) & 63))
        v4 = append(v4, (fi11 | 1))
        _ = fi11
    }
    _ = p0
    _ = v4
    return 1
}

func main() {
    var v1 bool = (true && (!false))
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    for fi12 := 0; fi12 < 20; fi12++ {
        if (!(794 > -994)) {
            if v1 {
                v1 = ((fi12 << ((376) & 63)) != fi12)
            }
            var v13 bool = false
            _ = v13
        } else {
            fmt.Println(v1, fi12)
            fmt.Println(fi12, v1)
            var v14 uint8 = 2
            _ = v14
        }
        v15 := GenS1{F0: genFn0(42), F1: ("hello" + "a"), F2: 18446744073709551615}
        v16 := []int{fi12}
        _ = fi12
        _ = v15
        _ = v16
    }
    v1 = false
    v17 := GenS2{F0: -2147483648, F1: 32768, F2: 118}
    if v1 {
        var v18 uint8 = 134
        var v19 int16 = -529
        var v20 int = 2147483647
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 float32 = (-2.25 / 3.14159)
    v22 := []int{-2147483648}
    if v1 {
        v23 := GenS3{F0: true, F1: 65535, F2: 128}
        switch "foo bar" {
        case "":
            v24 := GenS3{F0: v1, F1: 673, F2: 1}
            _ = v24
        case "hello":
            var v25 int = 652
            _ = v25
        case "\n\t":
            if v1 {
                if v1 {
                    v21 -= v21
                    var v26 uint32 = 363
                    _ = v26
                } else {
                    switch 32768 {
                    case 168:
                        v27 := GenS3{F0: (-784 != 850), F1: 2147483647, F2: 128}
                        _ = v27
                    default:
                        v28 := GenS3{F0: v1, F1: 42, F2: 127}
                        _ = v28
                    }
                }
                v29 := GenS1{F0: 2, F1: ("foo bar" + "\n\t"), F2: 473}
                switch 4294967295 {
                case 4294967295:
                    if (-1 < 100) {
                        var v30 uint16 = genFn0(0)
                        _ = v30
                    } else {
                        for fi31 := 0; fi31 < 19; fi31++ {
                            for fi32 := 0; fi32 < 2; fi32++ {
                                v1 = (42 == -1)
                                v22 = append(v22, (-(fi31)))
                                _ = fi32
                            }
                            _ = fi31
                        }
                        for fi33 := 0; fi33 < 10; fi33++ {
                            var v34 uint64 = 2
                            if len(v22) > 0 {
                                v22[0] = -373
                            }
                            _ = fi33
                            _ = v34
                        }
                    }
                default:
                    switch -151 {
                    case -2147483648:
                        for fi35 := 0; fi35 < 1; fi35++ {
                            var v36 int16 = -892
                            fmt.Println(v21, v36, v21)
                            _ = fi35
                            _ = v36
                        }
                    case 32768:
                        v1 = (!v1)
                    case -351:
                        var v37 int64 = -9223372036854775808
                        _ = v37
                    default:
                        switch "hello" {
                        case "\n\t":
                            if ((1 == 42) || true) {
                                v38 := GenS4{F0: "foo bar", F1: ""}
                                _ = v38
                            }
                        case "foo bar":
                            if true {
                                if true {
                                    v1 = (!(v1 || (100 == 62)))
                                    var v39 uint32 = 329
                                    _ = v39
                                }
                                for fi40 := 0; fi40 < 18; fi40++ {
                                    for fi41 := 0; fi41 < 9; fi41++ {
                                        v1 = v1
                                        _ = len(v22)
                                        if false {
                                            v2["hello"] = v1
                                        } else {
                                            v42 := GenS5{F0: 32767}
                                            var v43 int64 = 397
                                            _ = v42
                                            _ = v43
                                        }
                                        _ = fi41
                                    }
                                    var v44 int64 = 910
                                    if (2147483648 > 42) {
                                        var v45 float32 = (float32(1.5) - (0.0 + float32(0.1)))
                                        v46 := GenS6{F0: v44}
                                        _ = len(v3)
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = fi40
                                    _ = v44
                                }
                            } else {
                                for fi47 := 0; fi47 < 17; fi47++ {
                                    v21 = (1.5 / v21)
                                    _ = fi47
                                }
                            }
                        case "":
                            for fi48 := 0; fi48 < 1; fi48++ {
                                v49 := GenS3{F0: v1, F1: 2147483647, F2: 2}
                                v1 = v1
                                _ = fi48
                                _ = v49
                            }
                        default:
                            if (42 > 111) {
                                var v50 uint = 255
                                _ = v50
                            } else {
                                if v1 {
                                    var v51 uint16 = genFn0(22)
                                    var v52 uint64 = 590
                                    _ = v51
                                    _ = v52
                                } else {
                                    v22 = append(v22, -9223372036854775808)
                                    v53 := GenS7{F0: 2147483647}
                                    _ = v53
                                }
                                v1 = v1
                                var v54 string = ("hello" + "hello")
                                _ = v54
                            }
                        }
                    }
                }
                _ = v29
            }
        }
        _ = v23
    }
    v3[630] = 0
    fmt.Println(v1)
    v55 := GenS8{F0: 526, F1: -2147483648, F2: -60}
    if (109 < 127) {
        v1 = (true || v1)
        var v56 float64 = 0.1
        _ = v56
    } else {
        v21 += v21
        switch (("\n\t" + "foo bar") + "") {
        case "\n\t":
            if (!false) {
                switch 32767 {
                case 606:
                    var v57 bool = v1
                    _ = v57
                case 32767:
                    for fi58 := 0; fi58 < 15; fi58++ {
                        for fi59 := 0; fi59 < 13; fi59++ {
                            var v60 int8 = 35
                            v60 = v60
                            fmt.Println(fi59, v1)
                            _ = fi59
                            _ = v60
                        }
                        _ = fi58
                    }
                }
                fmt.Println(v1, v1)
            } else {
                fmt.Println(v21, v21)
            }
        case "hello":
            if v1 {
                v21 -= 3.14159
                _ = len(v22)
            }
        default:
            for fi61 := 0; fi61 < 18; fi61++ {
                v1 = (!v1)
                v1 = v1
                _ = fi61
            }
        }
    }
    v62 := GenS2{F0: 9223372036854775807, F1: -456, F2: -76}
    v1 = (!(!(!v1)))
    var v63 uint16 = 32767
    var v64 int32 = 42
    switch ("" + ("a" + "")) {
    case "hello":
        var v65 int64 = -449
        _ = v65
    }
    _ = len(v3)
    if true {
        v66 := GenS9{F0: (0 <= 1), F1: 1}
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v55
    _ = v62
    _ = v63
    _ = v64
}
