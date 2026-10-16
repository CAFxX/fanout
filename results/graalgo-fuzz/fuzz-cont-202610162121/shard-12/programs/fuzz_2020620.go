package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int8 }
type GenS2 struct { F0 int64; F1 bool; F2 int8 }
type GenS3 struct { F0 string; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 string; F1 uint32; F2 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8 }
func genFn0(p0 uint64) uint8 {
    if (false && ("a" > "foo bar")) {
        var v7 uint32 = 577
        _ = v7
    }
    switch (("foo bar" + "foo bar") + "") {
    case "a":
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 string = "a"
            var v10 uint8 = 210
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    if true {
        if true {
            var v11 uint16 = 255
            _ = v11
        } else {
            p0 = 690
        }
        v12 := GenS1{F0: 0, F1: 42, F2: -1}
        _ = v12
    }
    v13 := []int{0}
    _ = p0
    _ = v13
    return 121
}
func genFn1() uint16 {
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int8 = 0
        v16 := GenS1{F0: int(v15), F1: 42, F2: v15}
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 uint16 = 42
            var v19 uint16 = 706
            if (true && true) {
                v20 := []int{fi17, (-(fi14)), (fi14 * (fi14 * 842)), (fi14 & (fi17 ^ fi17))}
                _ = len(v20)
                v21 := GenS2{F0: 9223372036854775807, F1: (v15 == v15), F2: (v15 >> ((4294967295) & 63))}
                var v22 uint32 = 687
                v18 = 931
                _ = v20
                _ = v21
                _ = v22
            } else {
                v23 := GenS1{F0: (fi17 % ((fi17) | 1)), F1: 4294967295, F2: v15}
                v24 := []int{}
                _ = len(v24)
                _ = v23
                _ = v24
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    switch "hello" {
    case "\n\t":
        v25 := []int{}
        _ = v25
    case "foo bar":
        switch -1 {
        case 4294967295:
            v26 := GenS2{F0: 1, F1: (128 > 1), F2: 0}
            _ = v26
        case 598:
            v27 := GenS2{F0: 9223372036854775807, F1: false, F2: -88}
            _ = v27
        }
    }
    switch "a" {
    case "a":
        var v28 int8 = 0
        _ = v28
    case "\n\t":
        var v29 uint16 = 0
        _ = v29
    case "":
        v30 := GenS2{F0: 628, F1: true, F2: 127}
        _ = v30
    default:
        v31 := []int{-364}
        _ = v31
    }
    return 42
}

func main() {
    var v1 uint64 = 2
    var v2 int32 = -1
    v3 := []uint32{}
    v4 := make(map[string]int)
    v5 := make(map[string]string)
    var v6 int = 557
    v32 := GenS3{F0: "\n\t", F1: 894, F2: 142}
    v5["hello"] = "hello"
    if false {
        v1 -= 939
    } else {
        var v33 int8 = -128
        v34 := GenS4{F0: (v1 >> ((1) & 63)), F1: int32(v6)}
        _ = v33
        _ = v34
    }
    v6 -= v6
    for fi35 := 0; fi35 < 18; fi35++ {
        v1 = (v1 + 525)
        _ = fi35
    }
    fmt.Println(v6, v1, v6)
    v1 -= v1
    v2 = v2
    v2 = (v2 / ((v2) | 1))
    switch (v6 + (v6 + v6)) {
    case 128:
        v36 := GenS1{F0: (v6 | (v6 * (v6 * (v6 * v6)))), F1: 642, F2: 1}
        _ = v36
    case 42:
        if false {
            var v37 int64 = -123
            fmt.Println(v1, v37)
            _ = v37
        } else {
            var v38 uint8 = 255
            _ = v38
        }
    case -9223372036854775808:
        if ((^(v1)) <= (v1 % (((^(v1))) | 1))) {
            for fi39 := 0; fi39 < 20; fi39++ {
                fmt.Println(v1, fi39, fi39)
                v2 -= (v2 & v2)
                _ = fi39
            }
            switch (v6 & v6) {
            case -755:
                for fi40 := 0; fi40 < 17; fi40++ {
                    v41 := GenS3{F0: "", F1: 0, F2: 276}
                    var v42 uint = 778
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            case 128:
                fmt.Println(v1, v2, v1)
            case 1:
                v1 += (v1 >> ((65535) & 63))
            default:
                for fi43 := 0; fi43 < 20; fi43++ {
                    switch v6 {
                    case 0:
                        var v44 uint16 = 42
                        _ = v44
                    }
                    _ = fi43
                }
            }
            v2 = v2
        } else {
            var v45 int64 = 42
            for fi46 := 0; fi46 < 14; fi46++ {
                v47 := GenS5{F0: "foo bar", F1: 2147483647, F2: 0}
                switch (^(v6)) {
                case 128:
                    switch -182 {
                    case -552:
                        if ("" > "") {
                            v3 = append(v3, 2)
                            var v48 int64 = (v45 - 0)
                            v49 := GenS1{F0: (v6 << ((1) & 63)), F1: 2147483648, F2: 1}
                            _ = v48
                            _ = v49
                        } else {
                            var v50 int32 = -592
                            _ = v50
                        }
                    }
                case 2147483647:
                    for fi51 := 0; fi51 < 4; fi51++ {
                        v1 = v1
                        v52 := GenS2{F0: (v45 / ((331) | 1)), F1: true, F2: -128}
                        v53 := GenS1{F0: (v6 & v6), F1: 777, F2: 50}
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                case 127:
                    if (121 < 2) {
                        for fi54 := 0; fi54 < 7; fi54++ {
                            v6 = int(v2)
                            var v55 bool = true
                            _ = fi54
                            _ = v55
                        }
                        var v56 int8 = -128
                        var v57 int8 = v56
                        _ = v56
                        _ = v57
                    } else {
                        v4[("hello" + "foo bar")] = (v6 * v6)
                    }
                default:
                    var v58 uint64 = (^(v1))
                    _ = v58
                }
                var v59 bool = false
                _ = fi46
                _ = v47
                _ = v59
            }
            switch (v6 >> ((978) & 63)) {
            case 32768:
                for fi60 := 0; fi60 < 4; fi60++ {
                    fmt.Println(v2)
                    v61 := GenS6{F0: fi60}
                    _ = fi60
                    _ = v61
                }
            default:
                for fi62 := 0; fi62 < 8; fi62++ {
                    if false {
                        for fi63 := 0; fi63 < 4; fi63++ {
                            var v64 int64 = (v45 + (v45 & (v45 % (((^(v45))) | 1))))
                            if false {
                                v1 = v1
                                v65 := GenS2{F0: v45, F1: true, F2: -128}
                                _ = v65
                            } else {
                                delete(v5, "a")
                            }
                            var v66 uint = 4294967295
                            _ = fi63
                            _ = v64
                            _ = v66
                        }
                    } else {
                        fmt.Println(v45, v45)
                        if ((true || true) && (true && (v45 > v45))) {
                            v67 := GenS7{F0: "\n\t"}
                            v2 = v2
                            var v68 int = v6
                            _ = v67
                            _ = v68
                        } else {
                            if true {
                                for fi69 := 0; fi69 < 5; fi69++ {
                                    v2 = (v2 | v2)
                                    _ = fi69
                                }
                            } else {
                                for fi70 := 0; fi70 < 13; fi70++ {
                                    fmt.Println(v6)
                                    var v71 uint32 = 431
                                    _ = fi70
                                    _ = v71
                                }
                            }
                        }
                    }
                    _ = fi62
                }
            }
            _ = v45
        }
    }
    for fi72 := 0; fi72 < 5; fi72++ {
        if false {
            var v73 uint16 = 42
            v1 = v1
            _ = v73
        }
        _ = fi72
    }
    var v74 uint32 = 935
    v1 += uint64(v74)
    if (int16(v6) >= 32767) {
        v75 := GenS8{F0: 135}
        var v76 int64 = 0
        if (!(false || ((v74 + 4294967295) >= uint32(v1)))) {
            v76 += 0
        }
        _ = v75
        _ = v76
    } else {
        delete(v4, ("hello" + ("foo bar" + "foo bar")))
    }
    fmt.Println(v6, v74, v74)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v74
}
