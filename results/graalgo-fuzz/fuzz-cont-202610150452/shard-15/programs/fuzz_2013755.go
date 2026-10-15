package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS7 struct { F0 bool; F1 uint16; F2 uint32 }
func genFn0(p0 uint8, p1 int) int16 {
    fmt.Println(p1)
    if ((true && (42 > 811)) && (!true)) {
        fmt.Println(p1)
        v6 := GenS1{F0: -208}
        p1 += p1
        _ = v6
    }
    if false {
        var v7 uint64 = 0
        _ = v7
    } else {
        v8 := GenS1{F0: -108}
        var v9 int = (^(p1))
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 string) uint64 {
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := []int{-2147483648, (-(fi10)), fi10}
        _ = fi10
        _ = v11
    }
    switch 0 {
    case 128:
        v12 := []int{208, -2147483648}
        _ = v12
    case 4294967295:
        v13 := GenS2{F0: 255, F1: 868}
        _ = v13
    default:
        var v14 bool = true
        _ = v14
    }
    var v15 uint64 = 4294967295
    if false {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 float32 = 1e10
            v18 := []int{949}
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 uint32 = 4294967295
        _ = v19
    }
    v20 := GenS3{F0: 897, F1: v15, F2: 390}
    _ = p0
    _ = v15
    _ = v20
    return 645
}

func main() {
    var v1 uint16 = 589
    var v2 uint = 65535
    v3 := make(map[string]uint32)
    var v4 int32 = 1
    v5 := make(map[int]int)
    v1 -= (v1 + v1)
    if true {
        v4 += (v4 & (^(v4)))
    }
    var v21 uint16 = (v1 & (v1 << (((^(v2))) & 63)))
    v22 := []int{25}
    for fi23 := 0; fi23 < 8; fi23++ {
        v1 = 42
        v24 := GenS4{F0: 1}
        _ = fi23
        _ = v24
    }
    var v25 uint8 = 144
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 uint8 = 2
            fmt.Println(v28, fi26)
            var v29 int = fi26
            _ = fi27
            _ = v28
            _ = v29
        }
        switch (-(fi26)) {
        case -1:
            var v30 int8 = -122
            _ = v30
        case 2147483647:
            if (975 != 2) {
                if false {
                    if false {
                        if true {
                            if len(v22) > 0 {
                                v22[0] = fi26
                            }
                        } else {
                            var v31 int32 = (v4 ^ (v4 | int32(v2)))
                            if (v2 <= (v2 - (v2 >> ((v2) & 63)))) {
                                fmt.Println(v25)
                            } else {
                                for fi32 := 0; fi32 < 16; fi32++ {
                                    switch "a" {
                                    case "":
                                        delete(v5, (fi32 / (((fi26 % ((fi26) | 1))) | 1)))
                                    default:
                                        v21 -= v21
                                    }
                                    v4 = -1
                                    _ = fi32
                                }
                                if false {
                                    if true {
                                        if len(v22) >= 2 {
                                            v22 = v22[1:len(v22)]
                                        }
                                        if true {
                                            if ((-44 > -1) || true) {
                                                for fi33 := 0; fi33 < 5; fi33++ {
                                                    var v34 uint16 = (v1 << (((v2 + v2)) & 63))
                                                    switch (((("hello" + "hello") + "") + "") + "foo bar") {
                                                    case "hello":
                                                        v5[(fi26 & (fi33 % (((fi33 << ((v2) & 63))) | 1)))] = (fi26 / ((fi33) | 1))
                                                        v25 = (v25 & v25)
                                                        var v35 float32 = 1e10
                                                        _ = v35
                                                    }
                                                    _ = fi33
                                                    _ = v34
                                                }
                                                if (-32768 != genFn0(v25, fi26)) {
                                                    v4 = v31
                                                    v4 -= v4
                                                    v3["foo bar"] = 1
                                                }
                                            } else {
                                                var v36 uint64 = 18446744073709551615
                                                fmt.Println(v25)
                                                var v37 uint = (v2 >> (((^(v2))) & 63))
                                                _ = v36
                                                _ = v37
                                            }
                                        }
                                        for fi38 := 0; fi38 < 4; fi38++ {
                                            v39 := GenS5{F0: (v25 | v25), F1: 1}
                                            _ = fi38
                                            _ = v39
                                        }
                                    }
                                    var v40 float32 = (-2.25 / 1.0)
                                    _ = v40
                                } else {
                                    v41 := GenS6{F0: 42, F1: 42, F2: (v21 / ((uint16(v25)) | 1))}
                                    if false {
                                        for fi42 := 0; fi42 < 2; fi42++ {
                                            v5[fi42] = 32768
                                            v43 := GenS4{F0: 210}
                                            var v44 int8 = -80
                                            v45 := GenS7{F0: ((v44 + v44) != (v44 - v44)), F1: (^(v21)), F2: 0}
                                            if (v4 <= (v4 | 56)) {
                                                v31 += 225
                                                v46 := GenS4{F0: 700}
                                                _ = v46
                                            }
                                            delete(v5, fi42)
                                            _ = fi42
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                        }
                                        v47 := GenS1{F0: 393}
                                        _ = v47
                                    } else {
                                        v4 = v31
                                    }
                                    _ = v41
                                }
                                fmt.Println(v21, v21)
                            }
                            _ = v31
                        }
                    }
                }
            }
        case 127:
            if true {
                var v48 uint16 = (v1 | (v21 | 32767))
                _ = v48
            }
        }
        _ = fi26
    }
    if (!(!false)) {
        v21 = v21
    } else {
        var v49 int64 = -178
        v49 -= int64(v25)
        switch 9223372036854775807 {
        case 127:
            var v50 float64 = (3.14159 * (-2.25 - 0.0))
            _ = v50
        case 1:
            var v51 int16 = 42
            _ = v51
        case 32768:
            fmt.Println(v49, v21)
        }
        _ = v49
    }
    var v52 int16 = 696
    var v53 int32 = (v4 * 362)
    v52 -= v52
    if true {
        if (0.1 > -2.25) {
            v54 := GenS7{F0: (!false), F1: (^(v1)), F2: 860}
            for fi55 := 0; fi55 < 5; fi55++ {
                v52 = (v52 | genFn0(106, (fi55 | 9223372036854775807)))
                for fi56 := 0; fi56 < 1; fi56++ {
                    _ = len(v3)
                    _ = fi56
                }
                _ = fi55
            }
            _ = v54
        }
    } else {
        var v57 uint8 = (v25 - 0)
        v58 := GenS1{F0: (v53 + v53)}
        _ = v57
        _ = v58
    }
    var v59 int = -1000
    if (true && (false || ("\n\t" != ""))) {
        if true {
            v2 += v2
            if (1 > 110) {
                var v60 int = v59
                if ((true || true) && true) {
                    delete(v5, (v60 % (((v60 % ((v60) | 1))) | 1)))
                    fmt.Println(v1, v59)
                }
                _ = v60
            } else {
                v2 = (v2 * (v2 | (v2 * (v2 & 325))))
                var v61 uint64 = uint64(v2)
                v4 += (v53 | v4)
                _ = v61
            }
            switch v59 {
            case -533:
                for fi62 := 0; fi62 < 4; fi62++ {
                    var v63 int32 = 42
                    v21 = (v21 | uint16(v25))
                    _ = fi62
                    _ = v63
                }
            case -2147483648:
                v53 += 1
            default:
                if false {
                    for fi64 := 0; fi64 < 13; fi64++ {
                        var v65 float64 = -2.25
                        var v66 uint32 = 1
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                } else {
                    var v67 uint32 = 65535
                    v2 -= v2
                    _ = v67
                }
            }
        } else {
            switch "" {
            case "a":
                var v68 int32 = 1
                _ = v68
            }
            for fi69 := 0; fi69 < 3; fi69++ {
                for fi70 := 0; fi70 < 14; fi70++ {
                    if false {
                        v59 += int(v52)
                    } else {
                        v71 := GenS4{F0: 639}
                        v59 = (v59 + -54)
                        var v72 int = (^(fi69))
                        _ = v71
                        _ = v72
                    }
                    _ = fi70
                }
                _ = fi69
            }
            fmt.Println(v53)
        }
        v21 += 42
    } else {
        fmt.Println(v4, v25)
        switch (((("hello" + ("hello" + ("hello" + "hello"))) + "hello") + "\n\t") + "hello") {
        case "hello":
            v73 := GenS2{F0: (v21 >> (((v2 * v2)) & 63)), F1: 2147483647}
            _ = v73
        }
        var v74 float64 = 0.0
        _ = v74
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v59
}
