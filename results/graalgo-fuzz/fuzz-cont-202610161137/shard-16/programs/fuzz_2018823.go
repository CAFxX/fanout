package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int32; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int16; F2 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS10 struct { F0 uint; F1 uint64 }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 int8; F1 uint64 }
func genFn0(p0 int, p1 uint8) int8 {
    switch "foo bar" {
    case "foo bar":
        v5 := []int{p0, p0}
        _ = v5
    case "a":
        switch ("\n\t" + "") {
        case "":
            var v6 string = (("\n\t" + "") + (("foo bar" + "foo bar") + "foo bar"))
            _ = v6
        case "a":
            for fi7 := 0; fi7 < 20; fi7++ {
                if true {
                    if true {
                        var v8 int64 = -1
                        if ((!false) && true) {
                            v9 := GenS1{F0: v8, F1: 1}
                            p1 -= (p1 >> ((1) & 63))
                            _ = v9
                        }
                        for fi10 := 0; fi10 < 12; fi10++ {
                            var v11 int16 = int16(fi7)
                            v12 := []int{fi7}
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        _ = v8
                    }
                }
                p0 -= fi7
                _ = fi7
            }
        default:
            if (0.0 <= -2.25) {
                p0 += 127
                var v13 uint8 = 128
                for fi14 := 0; fi14 < 15; fi14++ {
                    p1 = v13
                    v15 := []int{}
                    _ = fi14
                    _ = v15
                }
                _ = v13
            } else {
                v16 := []int{}
                _ = v16
            }
        }
    default:
        v17 := GenS1{F0: -507, F1: 596}
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := GenS2{F0: 65535, F1: -1, F2: 42}
        switch "\n\t" {
        case "\n\t":
            v20 := []int{p0, p0}
            _ = v20
        case "":
            p0 -= (-(fi18))
        case "foo bar":
            if true {
                for fi21 := 0; fi21 < 11; fi21++ {
                    v22 := []int{}
                    _ = fi21
                    _ = v22
                }
                if (2 > 0) {
                    p0 = 1
                    if ((p1 - (p1 >> ((149) & 63))) >= p1) {
                        p1 -= (^(p1))
                        var v23 int16 = 293
                        v24 := GenS3{F0: 2}
                        _ = v23
                        _ = v24
                    } else {
                        p0 = -902
                        v25 := GenS2{F0: 255, F1: 931, F2: 829}
                        _ = v25
                    }
                    p0 -= (p0 / ((p0) | 1))
                }
                p1 += 127
            } else {
                v26 := []int{(-(fi18)), fi18}
                var v27 bool = false
                switch "a" {
                case "hello":
                    switch "\n\t" {
                    case "foo bar":
                        p0 -= p0
                    case "":
                        v26 = append(v26, 32768)
                    default:
                        p0 += (^(p0))
                    }
                case "a":
                    v27 = (uint16(p1) > 243)
                case "foo bar":
                    v28 := GenS4{F0: 0}
                    _ = v28
                default:
                    for fi29 := 0; fi29 < 15; fi29++ {
                        switch "\n\t" {
                        case "":
                            var v30 uint8 = p1
                            _ = v30
                        case "foo bar":
                            v27 = ((1 != 4294967295) || true)
                            var v31 int8 = 42
                            if len(v26) > 0 {
                                v26[0] = fi29
                            }
                            if v27 {
                                var v32 int16 = -941
                                fmt.Println(p0)
                                _ = v32
                            } else {
                                if v27 {
                                    if (1 == 65535) {
                                        var v33 uint16 = 42
                                        _ = v33
                                    }
                                } else {
                                    v34 := GenS5{F0: 0, F1: int16(fi29), F2: int64(v31)}
                                    v31 += (v31 + (^(v31)))
                                    if v27 {
                                        if true {
                                            var v35 int16 = 483
                                            var v36 uint16 = 439
                                            _ = v35
                                            _ = v36
                                        }
                                        if true {
                                            var v37 int32 = int32(v31)
                                            fmt.Println(fi29, p1, v37)
                                            var v38 uint8 = p1
                                            _ = v37
                                            _ = v38
                                        }
                                    } else {
                                        var v39 uint8 = 2
                                        fmt.Println(v31, p0, v31)
                                        v31 = (v31 >> ((2) & 63))
                                        if (!(v39 == v39)) {
                                            _ = len(v26)
                                        }
                                        _ = v39
                                    }
                                    _ = v34
                                }
                                var v40 int32 = -2147483648
                                _ = v40
                            }
                            _ = v31
                        }
                        var v41 int8 = 100
                        _ = fi29
                        _ = v41
                    }
                }
                _ = v26
                _ = v27
            }
        }
        v42 := GenS5{F0: 2147483647, F1: 42, F2: -1}
        _ = fi18
        _ = v19
        _ = v42
    }
    v43 := GenS6{F0: 42}
    v44 := GenS2{F0: 32767, F1: -1, F2: -1}
    v45 := []int{(^(p0))}
    _ = p0
    _ = p1
    _ = v43
    _ = v44
    _ = v45
    return 0
}

func main() {
    v1 := []int{130}
    var v2 int8 = 100
    v3 := []string{("\n\t" + ("\n\t" + "foo bar")), "foo bar", ("hello" + "foo bar")}
    var v4 uint64 = 672
    var v46 int = -12
    if len(v1) > 0 {
        v1[0] = v46
    }
    if true {
        for fi47 := 0; fi47 < 15; fi47++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v48 int64 = 874
            v49 := GenS7{F0: -2147483648}
            _ = fi47
            _ = v48
            _ = v49
        }
    } else {
        v2 = v2
        switch ("\n\t" + "\n\t") {
        case "a":
            var v50 uint = 837
            _ = v50
        case "hello":
            var v51 uint32 = 42
            _ = v51
        default:
            switch "hello" {
            case "foo bar":
                v52 := GenS6{F0: 42}
                _ = v52
            case "\n\t":
                v4 = (v4 ^ 18446744073709551615)
            case "hello":
                if true {
                    v46 -= v46
                    var v53 uint64 = 1
                    var v54 float32 = -2.25
                    _ = v53
                    _ = v54
                }
            }
        }
    }
    var v55 int16 = -1
    var v56 bool = true
    v3 = append(v3, "")
    var v57 int = (v46 + (v46 << ((uint(v4)) & 63)))
    var v58 int8 = (v2 * v2)
    var v59 int16 = (v55 << ((58) & 63))
    v60 := GenS8{F0: (v57 & v57)}
    if len(v1) > 0 {
        v1[0] = 270
    }
    for fi61 := 0; fi61 < 6; fi61++ {
        var v62 uint64 = v4
        var v63 uint = 1
        switch fi61 {
        case 366:
            v58 = (v2 / (((v58 - v58)) | 1))
        case 1:
            v55 += -32768
        case 4294967295:
            v64 := GenS9{F0: 2, F1: 1, F2: 1}
            _ = v64
        default:
            for fi65 := 0; fi65 < 20; fi65++ {
                v66 := GenS10{F0: (v63 << (((^(v63))) & 63)), F1: 98}
                v67 := GenS11{F0: (2147483648 <= 0), F1: 65535}
                v59 += (v59 % (((v55 | v55)) | 1))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi65
                _ = v66
                _ = v67
            }
        }
        _ = fi61
        _ = v62
        _ = v63
    }
    var v68 uint = 18446744073709551615
    v69 := GenS12{F0: 42, F1: 18446744073709551615}
    v58 = genFn0((^(v57)), 2)
    var v70 bool = (251 < 20)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v68
    _ = v69
    _ = v70
}
