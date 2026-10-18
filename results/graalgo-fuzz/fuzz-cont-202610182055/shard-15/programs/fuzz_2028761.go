package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 int8; F1 int64; F2 int }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 bool }
func genFn0(p0 int, p1 uint64, p2 int8) int64 {
    var v5 uint64 = (p1 + p1)
    var v6 int16 = 872
    v6 = (v6 % (((v6 << ((119) & 63))) | 1))
    p0 -= (p0 << ((528) & 63))
    if (4294967295 < 1) {
        switch 0 {
        case 4294967295:
            for fi7 := 0; fi7 < 3; fi7++ {
                for fi8 := 0; fi8 < 8; fi8++ {
                    p2 -= -30
                    if (1 != 128) {
                        p0 -= fi7
                        var v9 uint32 = 42
                        _ = v9
                    } else {
                        v10 := []int{(fi8 & (p0 % ((fi8) | 1))), (fi8 | fi7)}
                        _ = v10
                    }
                    _ = fi8
                }
                if true {
                    v11 := []int{p0, (^(fi7)), (^(fi7)), -1}
                    p0 = (fi7 & -9223372036854775808)
                    _ = v11
                }
                if (233 < 42) {
                    switch int(p2) {
                    case 42:
                        var v12 float64 = 3.14159
                        _ = v12
                    case -865:
                        var v13 float64 = 1.5
                        _ = v13
                    case -1:
                        var v14 bool = false
                        _ = v14
                    }
                } else {
                    v15 := GenS1{F0: 840, F1: 1}
                    fmt.Println(fi7, v6, p0)
                    var v16 uint8 = 42
                    _ = v15
                    _ = v16
                }
                _ = fi7
            }
        case -1:
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 int8 = p2
                v19 := GenS2{F0: 92, F1: 0, F2: 1}
                _ = fi17
                _ = v18
                _ = v19
            }
        case 32768:
            if true {
                switch (p0 / (((p0 * int(p2))) | 1)) {
                case 32767:
                    p1 = p1
                case 4294967295:
                    switch (p0 - p0) {
                    case 432:
                        v5 = v5
                    case -9223372036854775808:
                        v6 = (v6 >> ((42) & 63))
                    default:
                        v20 := GenS1{F0: 1, F1: 1}
                        _ = v20
                    }
                }
            }
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := []int{}
            if (42 < 42) {
                v23 := GenS3{F0: 255, F1: (^(v5)), F2: (-(v6))}
                fmt.Println(p1, fi21, v5)
                _ = v23
            } else {
                p2 += p2
                var v24 uint16 = uint16(v5)
                _ = v24
            }
            p2 -= (p2 >> ((2) & 63))
            _ = fi21
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}
func genFn1() uint32 {
    v25 := GenS3{F0: 65535, F1: 0, F2: 492}
    var v26 int16 = -1
    v27 := []int{}
    var v28 string = ((("\n\t" + "hello") + "\n\t") + "\n\t")
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 uint = 255
    var v4 int32 = -68
    for fi29 := 0; fi29 < 2; fi29++ {
        for fi30 := 0; fi30 < 5; fi30++ {
            v3 = (v3 & v3)
            _ = fi30
        }
        for fi31 := 0; fi31 < 12; fi31++ {
            v3 = v3
            _ = fi31
        }
        _ = fi29
    }
    v32 := GenS4{F0: 748}
    v33 := GenS5{F0: -1, F1: 128}
    if (!(true && true)) {
        if false {
            v4 += v4
            for fi34 := 0; fi34 < 17; fi34++ {
                _ = len(v2)
                if true {
                    v35 := GenS3{F0: 255, F1: 0, F2: -1}
                    var v36 uint64 = 1
                    _ = v35
                    _ = v36
                } else {
                    var v37 float64 = (-2.25 / 1.0)
                    v38 := GenS3{F0: 110, F1: 18446744073709551615, F2: 622}
                    _ = v37
                    _ = v38
                }
                if ((fi34 % (((-(fi34))) | 1)) >= 127) {
                    if (18446744073709551615 >= 4294967295) {
                        v39 := GenS2{F0: 106, F1: 102, F2: 2}
                        _ = v39
                    }
                    v3 += (^(v3))
                    switch ("foo bar" + "hello") {
                    case "":
                        v4 -= (v4 + (^(v4)))
                    case "foo bar":
                        v4 += -424
                    case "hello":
                        var v40 uint64 = 2
                        _ = v40
                    default:
                        v3 = (v3 ^ (^(v3)))
                    }
                } else {
                    v41 := GenS6{F0: -32768, F1: false}
                    for fi42 := 0; fi42 < 16; fi42++ {
                        v43 := GenS5{F0: (fi42 / ((32767) | 1)), F1: fi34}
                        v3 += (v3 & v3)
                        _ = fi42
                        _ = v43
                    }
                    var v44 uint8 = 128
                    _ = v41
                    _ = v44
                }
                _ = fi34
            }
        }
    } else {
        for fi45 := 0; fi45 < 8; fi45++ {
            if false {
                var v46 bool = false
                if (2147483647 <= 4294967295) {
                    v3 += 1
                }
                _ = v46
            } else {
                if (!(v3 != 255)) {
                    for fi47 := 0; fi47 < 4; fi47++ {
                        if true {
                            v48 := []int{}
                            v49 := GenS1{F0: 42, F1: 42}
                            v3 = 4294967295
                            _ = v48
                            _ = v49
                        }
                        v50 := GenS3{F0: 26, F1: 2, F2: -979}
                        if false {
                            v51 := []int{-1, 782, fi45, fi45}
                            var v52 uint16 = 32767
                            if false {
                                v4 += v4
                            } else {
                                if (!(0.1 >= (1.5 / float32(1e10)))) {
                                    if false {
                                        var v53 string = "foo bar"
                                        _ = v53
                                    } else {
                                        var v54 int32 = (v4 * (v4 / ((797) | 1)))
                                        var v55 uint = (v3 + 4294967295)
                                        v56 := GenS7{F0: 127, F1: -9223372036854775808, F2: fi45}
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                }
                                if false {
                                    if true {
                                        v1[775] = (fi45 % (((-(fi45))) | 1))
                                    }
                                }
                                v3 -= 18446744073709551615
                            }
                            var v57 uint8 = 128
                            _ = v51
                            _ = v52
                            _ = v57
                        } else {
                            switch (fi47 + -9223372036854775808) {
                            case 32768:
                                var v58 int64 = genFn0(fi47, 18446744073709551615, 0)
                                _ = v58
                            default:
                                v59 := GenS8{F0: v3, F1: fi47}
                                _ = v59
                            }
                            fmt.Println(v3, fi47, fi45)
                            fmt.Println(v3, v4, fi45)
                        }
                        _ = fi47
                        _ = v50
                    }
                }
                v60 := []int{(fi45 << ((65535) & 63)), fi45}
                fmt.Println(fi45, v3, v4)
                _ = v60
            }
            v61 := []int{}
            _ = fi45
            _ = v61
        }
        var v62 int8 = -1
        _ = v62
    }
    for fi63 := 0; fi63 < 2; fi63++ {
        v64 := []int{fi63, (fi63 - fi63), 32767}
        var v65 float64 = 1e10
        _ = fi63
        _ = v64
        _ = v65
    }
    if false {
        _ = len(v1)
    }
    switch 4294967295 {
    case -1:
        var v66 uint16 = 968
        _ = v66
    default:
        switch -326 {
        case -9223372036854775808:
            v1[genFn1()] = 42
        case 0:
            var v67 uint32 = 70
            _ = v67
        default:
            for fi68 := 0; fi68 < 8; fi68++ {
                fmt.Println(v3, v4)
                fmt.Println(v4, fi68)
                v69 := GenS9{F0: 154}
                _ = fi68
                _ = v69
            }
        }
    }
    v70 := GenS4{F0: 2147483647}
    var v71 bool = (463 > 624)
    for fi72 := 0; fi72 < 8; fi72++ {
        v73 := v1[65535]
        v74 := GenS9{F0: 87}
        _ = fi72
        _ = v73
        _ = v74
    }
    v75 := []int{127}
    if v71 {
        for fi76 := 0; fi76 < 5; fi76++ {
            fmt.Println(v71, v3)
            v77 := GenS8{F0: v3, F1: (fi76 >> (((v3 >> (((^(v3))) & 63))) & 63))}
            _ = fi76
            _ = v77
        }
        v78 := GenS10{F0: (!false)}
        _ = v78
    }
    var v79 float64 = 0.0
    var v80 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v75:", len(v75), v75)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v70
    _ = v71
    _ = v75
    _ = v79
    _ = v80
}
