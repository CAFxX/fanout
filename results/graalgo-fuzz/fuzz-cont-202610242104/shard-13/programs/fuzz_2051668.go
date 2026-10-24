package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 uint32; F1 int8; F2 string }
func genFn0() uint16 {
    v6 := []int{42}
    if true {
        for fi7 := 0; fi7 < 18; fi7++ {
            switch fi7 {
            case -2147483648:
                if true {
                    for fi8 := 0; fi8 < 14; fi8++ {
                        var v9 uint16 = 1
                        _ = fi8
                        _ = v9
                    }
                    var v10 uint64 = 0
                    _ = v10
                }
            case -1:
                fmt.Println(fi7)
            default:
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 uint32 = 1
                    _ = fi11
                    _ = v12
                }
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi7
        }
        v13 := GenS1{F0: -32768, F1: 2147483647}
        var v14 int64 = 940
        _ = v13
        _ = v14
    } else {
        var v15 string = ("hello" + (("hello" + "") + "a"))
        fmt.Println(v15)
        var v16 int32 = -1
        _ = v15
        _ = v16
    }
    var v17 uint8 = 0
    fmt.Println(v17)
    _ = v17
    _ = v6
    return 1
}
func genFn1(p0 int16) int32 {
    fmt.Println(p0)
    var v18 int32 = 0
    v19 := []int{-9223372036854775808, 0, 280, -101}
    _ = p0
    _ = v18
    _ = v19
    return v18
}

func main() {
    var v1 uint32 = 478
    v2 := []int{-9223372036854775808, -808}
    var v3 bool = true
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 65535
    switch (("foo bar" + "") + "foo bar") {
    case "":
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS2{F0: 1}
            switch "" {
            case "a":
                v22 := GenS3{F0: 0, F1: uint16(v1), F2: 127}
                _ = v22
            case "foo bar":
                v23 := GenS4{F0: 1, F1: 32767, F2: (v5 >> ((321) & 63))}
                _ = v23
            case "\n\t":
                for fi24 := 0; fi24 < 5; fi24++ {
                    var v25 int32 = -1
                    switch "a" {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "a":
                        v25 = genFn1(1)
                    }
                    v25 = (-(v25))
                    _ = fi24
                    _ = v25
                }
            default:
                switch "foo bar" {
                case "a":
                    var v26 float64 = 0.0
                    _ = v26
                case "foo bar":
                    for fi27 := 0; fi27 < 3; fi27++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        fmt.Println(fi20, fi20, v3)
                        switch fi27 {
                        case 1:
                            v3 = true
                        default:
                            v28 := GenS4{F0: 340, F1: 65535, F2: v5}
                            _ = v28
                        }
                        _ = fi27
                    }
                default:
                    fmt.Println(v5, fi20)
                }
            }
            switch "hello" {
            case "":
                v29 := GenS5{F0: 233}
                _ = v29
            case "hello":
                var v30 int8 = 54
                _ = v30
            default:
                for fi31 := 0; fi31 < 16; fi31++ {
                    var v32 bool = (519 == 0)
                    var v33 float64 = 1.5
                    v3 = (4294967295 == 65535)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
            _ = fi20
            _ = v21
        }
    case "foo bar":
        var v34 uint8 = 255
        _ = v34
    default:
        switch "hello" {
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = -492
            }
        case "":
            var v35 uint = 2
            _ = v35
        case "hello":
            v36 := GenS4{F0: 851, F1: 347, F2: v1}
            _ = v36
        default:
            v37 := GenS6{F0: 367, F1: false}
            _ = v37
        }
    }
    fmt.Println(v3, v1, v3)
    v38 := GenS7{F0: (v5 + (v5 % ((v1) | 1))), F1: 1, F2: ("hello" + "foo bar")}
    v3 = ((true && ((1 <= 1) && ((v1 * 384) < v5))) && v3)
    for fi39 := 0; fi39 < 6; fi39++ {
        if (genFn0() < 787) {
            switch fi39 {
            case -9223372036854775808:
                var v40 uint32 = v1
                _ = v40
            }
            _ = len(v4)
            if len(v2) > 0 {
                v2[0] = (^(fi39))
            }
        }
        fmt.Println(fi39)
        var v41 int = 127
        _ = fi39
        _ = v41
    }
    var v42 uint = uint(v5)
    v1 = (v1 << ((v42) & 63))
    v42 += 65535
    v43 := GenS4{F0: 970, F1: 265, F2: 527}
    v3 = (-473 >= 32767)
    v3 = (73 >= 2)
    var v44 int32 = 1
    fmt.Println(v44, v1)
    for fi45 := 0; fi45 < 1; fi45++ {
        if (true || v3) {
            v5 = v5
            v1 -= (v1 >> ((288) & 63))
        } else {
            if (100 == 0) {
                switch "foo bar" {
                case "":
                    v42 = (v42 << ((255) & 63))
                case "a":
                    var v46 int16 = -1
                    _ = v46
                case "foo bar":
                    var v47 uint32 = (v5 - v5)
                    _ = v47
                }
                var v48 int = fi45
                _ = v48
            } else {
                switch ("" + "foo bar") {
                case "a":
                    v49 := v4[(v5 << (((v42 ^ 181)) & 63))]
                    _ = v49
                case "\n\t":
                    v50 := GenS6{F0: v44, F1: (!(((((!(-1 >= 42)) || true) && v3) && v3) && v3))}
                    _ = v50
                case "foo bar":
                    for fi51 := 0; fi51 < 4; fi51++ {
                        v52 := GenS2{F0: -128}
                        fmt.Println(fi45, v5, v3)
                        _ = fi51
                        _ = v52
                    }
                default:
                    v44 = int32(v5)
                }
                fmt.Println(v44)
                var v53 float64 = (1e10 - 1e10)
                _ = v53
            }
            fmt.Println(v1, v1)
        }
        v42 += 65535
        switch fi45 {
        case 32768:
            v54 := GenS6{F0: v44, F1: true}
            _ = v54
        case -593:
            if (-32768 == 1) {
                var v55 int = (fi45 >> ((v42) & 63))
                if ((!v3) || ((int(v5) <= fi45) && false)) {
                    v44 += v44
                    for fi56 := 0; fi56 < 2; fi56++ {
                        var v57 uint = (^(v42))
                        var v58 int64 = 914
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                    fmt.Println(v1)
                } else {
                    for fi59 := 0; fi59 < 13; fi59++ {
                        var v60 string = "\n\t"
                        _ = fi59
                        _ = v60
                    }
                    fmt.Println(v44, v3)
                    var v61 int16 = -203
                    _ = v61
                }
                _ = v55
            }
        case 2147483647:
            switch "a" {
            case "":
                v62 := GenS2{F0: 34}
                _ = v62
            default:
                v5 += (v1 & (^(v5)))
            }
        default:
            for fi63 := 0; fi63 < 14; fi63++ {
                v44 = (-(v44))
                var v64 int32 = (v44 ^ 0)
                var v65 uint16 = 1
                _ = fi63
                _ = v64
                _ = v65
            }
        }
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
