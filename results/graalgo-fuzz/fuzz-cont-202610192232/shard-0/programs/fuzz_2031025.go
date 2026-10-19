package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int; F2 string }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int64) int {
    v6 := []int{-512}
    if true {
        if ((!true) || ((("foo bar" + "") + "\n\t") < "a")) {
            var v7 uint8 = 255
            for fi8 := 0; fi8 < 20; fi8++ {
                fmt.Println(p0, p0)
                p0 = (p0 * 724)
                _ = fi8
            }
            _ = v7
        }
    } else {
        for fi9 := 0; fi9 < 13; fi9++ {
            switch fi9 {
            case 0:
                var v10 float64 = 0.0
                _ = v10
            case -1:
                var v11 uint64 = 399
                _ = v11
            case 32767:
                for fi12 := 0; fi12 < 18; fi12++ {
                    p0 = p0
                    _ = fi12
                }
            default:
                for fi13 := 0; fi13 < 10; fi13++ {
                    if false {
                        p0 = -879
                    }
                    switch "a" {
                    case "hello":
                        p0 = (p0 * 0)
                    default:
                        var v14 int16 = 1
                        _ = v14
                    }
                    _ = fi13
                }
            }
            _ = fi9
        }
    }
    p0 += p0
    var v15 string = "\n\t"
    _ = p0
    _ = v15
    _ = v6
    return -785
}
func genFn1(p0 uint16, p1 int, p2 uint32) int16 {
    p0 = p0
    v16 := []int{p1, p1, (p1 + (p1 >> ((65535) & 63))), p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 568
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 string = "hello"
    v5 := []string{(v4 + v4), (v4 + ("\n\t" + v4)), v4}
    v17 := GenS1{F0: v4, F1: 128, F2: 2147483647}
    for fi18 := 0; fi18 < 12; fi18++ {
        delete(v1, (fi18 % ((genFn0(-9223372036854775808)) | 1)))
        _ = fi18
    }
    var v19 int8 = 127
    if (v19 == v19) {
        var v20 uint = 2
        var v21 int64 = 1
        _ = v20
        _ = v21
    } else {
        if (true || false) {
            if len(v5) > 0 {
                v5[0] = "a"
            }
            v19 = v19
            var v22 uint16 = 55
            _ = v22
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 int64 = 263
            switch v4 {
            case "hello":
                v25 := v3[-1]
                _ = v25
            default:
                v24 = v24
            }
            if false {
                v24 = v24
            }
            _ = fi23
            _ = v24
        }
        fmt.Println(v19)
    }
    if false {
        var v26 uint8 = 54
        _ = v26
    } else {
        v19 += (v19 << ((65535) & 63))
    }
    var v27 int = 32767
    var v28 int32 = -2147483648
    var v29 uint32 = 145
    v29 += v29
    if len(v5) > 0 {
        v5[0] = v4
    }
    switch "a" {
    case "a":
        v30 := GenS1{F0: v4, F1: 127, F2: (v27 + (v27 ^ (-(v27))))}
        _ = v30
    case "hello":
        v4 = ""
    default:
        var v31 uint32 = v29
        _ = v31
    }
    var v32 int16 = 1000
    if (!(true && true)) {
        if (true || (!true)) {
            if ((v19 / ((int8(v32)) | 1)) != 127) {
                if len(v5) > 0 {
                    v5[0] = v4
                }
                if len(v5) > 0 {
                    v5[0] = v4
                }
            } else {
                for fi33 := 0; fi33 < 2; fi33++ {
                    var v34 int32 = int32(fi33)
                    v4 = "hello"
                    if true {
                        if len(v5) > 0 {
                            v5[0] = v4
                        }
                    } else {
                        if (true || false) {
                            var v35 uint64 = 42
                            _ = v35
                        }
                        if true {
                            if true {
                                var v36 int8 = v19
                                for fi37 := 0; fi37 < 4; fi37++ {
                                    if (v32 != v32) {
                                        v3[(-(fi37))] = (v29 >> ((65535) & 63))
                                    } else {
                                        var v38 int16 = (^(v32))
                                        v39 := GenS1{F0: (v4 + v4), F1: 255, F2: (v27 + fi33)}
                                        _ = v38
                                        _ = v39
                                    }
                                    if false {
                                        var v40 uint8 = 0
                                        switch v4 {
                                        case "":
                                            v4 = v4
                                        case "hello":
                                            v41 := GenS1{F0: (((v4 + v4) + v4) + "foo bar"), F1: v40, F2: genFn0(-824)}
                                            _ = v41
                                        case "a":
                                            v42 := GenS2{F0: (v32 / ((genFn1(32767, (v27 << ((0) & 63)), v29)) | 1))}
                                            _ = v42
                                        default:
                                            _ = len(v2)
                                        }
                                        var v43 float64 = 1e10
                                        _ = v40
                                        _ = v43
                                    }
                                    _ = fi37
                                }
                                var v44 int8 = v19
                                _ = v36
                                _ = v44
                            } else {
                                fmt.Println(v4, v28, v34)
                            }
                            for fi45 := 0; fi45 < 13; fi45++ {
                                fmt.Println(v34, v29, v28)
                                v46 := GenS2{F0: v32}
                                _ = fi45
                                _ = v46
                            }
                        }
                    }
                    _ = fi33
                    _ = v34
                }
                for fi47 := 0; fi47 < 3; fi47++ {
                    v48 := GenS3{F0: uint16(v29), F1: 4294967295, F2: v4}
                    for fi49 := 0; fi49 < 15; fi49++ {
                        v3[(v27 & fi49)] = (v29 - v29)
                        v4 = (v4 + "")
                        if (!(972 < 18446744073709551615)) {
                            switch fi47 {
                            case 9223372036854775807:
                                fmt.Println(v32, v4)
                            case -674:
                                var v50 uint = 302
                                _ = v50
                            default:
                                fmt.Println(v4, v32)
                            }
                            v29 = v29
                            v28 += v28
                        } else {
                            v51 := GenS2{F0: 0}
                            _ = len(v5)
                            _ = v51
                        }
                        _ = fi49
                    }
                    v52 := GenS4{F0: v28, F1: v28}
                    _ = fi47
                    _ = v48
                    _ = v52
                }
                for fi53 := 0; fi53 < 16; fi53++ {
                    switch v4 {
                    case "foo bar":
                        v54 := GenS1{F0: (((v4 + "foo bar") + v4) + v4), F1: 103, F2: (-(fi53))}
                        _ = v54
                    case "a":
                        for fi55 := 0; fi55 < 8; fi55++ {
                            var v56 int64 = -9223372036854775808
                            _ = fi55
                            _ = v56
                        }
                    case "\n\t":
                        switch "foo bar" {
                        case "":
                            var v57 int32 = -1
                            _ = v57
                        case "\n\t":
                            for fi58 := 0; fi58 < 10; fi58++ {
                                fmt.Println(v4, fi53, fi58)
                                var v59 string = v4
                                _ = fi58
                                _ = v59
                            }
                        case "a":
                            v60 := GenS5{F0: 9223372036854775807, F1: v19}
                            _ = v60
                        }
                    }
                    var v61 string = ("\n\t" + ("hello" + v4))
                    _ = fi53
                    _ = v61
                }
            }
        }
        v62 := GenS6{F0: 18446744073709551615}
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
