package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint64; F1 string; F2 uint }
type GenS5 struct { F0 int32; F1 uint64; F2 int16 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 int8 }
func genFn0(p0 int16, p1 string, p2 uint64) int {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int32 = 384
        _ = fi7
        _ = v8
    }
    if (1 != 65535) {
        v9 := []int{435, 32768}
        _ = v9
    } else {
        p1 = "a"
        v10 := []int{128, -522, -1, 204}
        for fi11 := 0; fi11 < 1; fi11++ {
            switch p1 {
            case "foo bar":
                v12 := GenS1{F0: 1}
                _ = v12
            default:
                v13 := GenS2{F0: 2147483647, F1: 127}
                _ = v13
            }
            v14 := GenS3{F0: 255, F1: 523, F2: -1}
            _ = fi11
            _ = v14
        }
        _ = v10
    }
    v15 := GenS3{F0: 2, F1: 2, F2: -128}
    if false {
        if false {
            for fi16 := 0; fi16 < 16; fi16++ {
                v17 := []int{(fi16 & fi16), (fi16 << ((65535) & 63)), 2147483647, int(p2)}
                switch ("foo bar" + p1) {
                case "foo bar":
                    v18 := GenS4{F0: (p2 * p2), F1: p1, F2: 0}
                    _ = v18
                case "a":
                    if (4294967295 < 2) {
                        switch fi16 {
                        case -9223372036854775808:
                            for fi19 := 0; fi19 < 10; fi19++ {
                                p1 = p1
                                fmt.Println(fi19, fi19, p2)
                                _ = fi19
                            }
                        }
                    }
                case "\n\t":
                    var v20 float64 = 0.1
                    _ = v20
                }
                _ = fi16
                _ = v17
            }
            var v21 float64 = (-2.25 + -2.25)
            _ = v21
        } else {
            v22 := GenS1{F0: (-(p0))}
            var v23 uint64 = (p2 ^ 385)
            _ = v22
            _ = v23
        }
        var v24 int64 = 42
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return -518
}
func genFn1() int64 {
    var v25 int64 = -994
    v26 := []int{genFn0(1000, "foo bar", 4294967295), genFn0(0, "hello", 18446744073709551615), -354, 1}
    if (1.5 == -2.25) {
        var v27 int64 = 418
        if false {
            var v28 float32 = (0.0 * -2.25)
            switch "hello" {
            case "":
                fmt.Println(v25, v25)
            case "\n\t":
                v28 -= 0.0
            }
            v29 := GenS5{F0: 1, F1: 42, F2: 32767}
            _ = v28
            _ = v29
        } else {
            if true {
                fmt.Println(v27)
            } else {
                switch -2147483648 {
                case 275:
                    v27 += 9223372036854775807
                }
                v30 := GenS5{F0: 42, F1: 0, F2: 42}
                _ = v30
            }
            for fi31 := 0; fi31 < 19; fi31++ {
                var v32 int64 = v25
                _ = fi31
                _ = v32
            }
            var v33 int16 = -1
            _ = v33
        }
        var v34 int8 = int8(v25)
        _ = v27
        _ = v34
    } else {
        switch "a" {
        case "hello":
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 uint32 = 83
                _ = fi35
                _ = v36
            }
        case "":
            var v37 uint16 = 255
            _ = v37
        default:
            var v38 float64 = 0.1
            _ = v38
        }
        var v39 uint16 = 213
        v39 += v39
        _ = v39
    }
    _ = v25
    _ = v26
    return v25
}

func main() {
    var v1 int32 = 2147483647
    var v2 int8 = 100
    v3 := []uint32{65535, 2147483647, 702, 255}
    var v4 bool = (!true)
    v5 := make(map[int]uint32)
    v6 := []uint32{1}
    var v40 uint32 = 869
    v2 = v2
    if v4 {
        var v41 float64 = (0.0 / -2.25)
        v4 = v4
        _ = v41
    }
    var v42 int64 = -9223372036854775808
    for fi43 := 0; fi43 < 15; fi43++ {
        for fi44 := 0; fi44 < 9; fi44++ {
            switch fi43 {
            case 9223372036854775807:
                v2 -= (v2 << ((0) & 63))
            default:
                v42 -= 0
            }
            _ = fi44
        }
        var v45 bool = ((42 == 4294967295) && false)
        _ = fi43
        _ = v45
    }
    if v4 {
        v46 := GenS6{F0: v2, F1: 42}
        _ = v46
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        v42 += genFn1()
        v40 += 42
        _ = fi47
    }
    if ((v42 + 0) != genFn1()) {
        v1 = (v1 - (^(v1)))
        var v48 bool = true
        v1 = v1
        _ = v48
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v42, v4)
    case "\n\t":
        v40 = (^(v40))
    }
    var v49 uint64 = 575
    v5[-37] = v40
    v50 := GenS4{F0: (v49 - v49), F1: ((("a" + "") + "\n\t") + ("\n\t" + (("\n\t" + ((("" + "") + "\n\t") + "")) + ("a" + ("hello" + "\n\t"))))), F2: 232}
    _ = len(v6)
    if (v2 > v2) {
        if true {
            for fi51 := 0; fi51 < 4; fi51++ {
                switch "hello" {
                case "\n\t":
                    v1 += -1
                case "hello":
                    for fi52 := 0; fi52 < 12; fi52++ {
                        v49 = v49
                        var v53 uint32 = 857
                        _ = fi52
                        _ = v53
                    }
                case "foo bar":
                    for fi54 := 0; fi54 < 3; fi54++ {
                        if true {
                            v4 = (true && v4)
                        } else {
                            v49 = 527
                        }
                        v55 := GenS7{F0: 0, F1: 65535}
                        v56 := GenS3{F0: 65535, F1: 2, F2: (-(v2))}
                        v57 := GenS2{F0: -1, F1: (^(v2))}
                        _ = fi54
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                default:
                    var v58 int16 = 32767
                    _ = v58
                }
                var v59 int = fi51
                _ = fi51
                _ = v59
            }
            var v60 uint16 = 820
            for fi61 := 0; fi61 < 4; fi61++ {
                for fi62 := 0; fi62 < 3; fi62++ {
                    if (0 >= 128) {
                        var v63 uint64 = v49
                        _ = v63
                    }
                    _ = fi62
                }
                v42 -= v42
                if ((0.1 + 1e10) == (0.0 * float32(float64(float32(3.14159))))) {
                    var v64 bool = (v4 && v4)
                    _ = v64
                } else {
                    _ = len(v5)
                    fmt.Println(v60)
                    if ((34 <= 1000) && v4) {
                        if len(v3) > 0 {
                            v3[0] = 232
                        }
                        v49 -= v49
                        for fi65 := 0; fi65 < 3; fi65++ {
                            fmt.Println(v2)
                            if len(v6) > 0 {
                                v6[0] = 187
                            }
                            v4 = (!((1.5 / float64(float32(1.5))) > 1.5))
                            switch ("foo bar" + "foo bar") {
                            case "foo bar":
                                v42 -= (-(v42))
                            case "":
                                v2 += (v2 | int8(v40))
                            case "\n\t":
                                switch genFn0(-75, ((("foo bar" + "") + "foo bar") + ""), 42) {
                                case 2147483647:
                                    v66 := GenS8{F0: -1, F1: true}
                                    _ = v66
                                default:
                                    v3 = append(v3, v40)
                                }
                                var v67 uint8 = 127
                                if true {
                                    v68 := GenS5{F0: 0, F1: (^(v49)), F2: -32768}
                                    _ = v68
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        v49 = (v49 % ((v49) | 1))
                                    case "a":
                                        v40 = v40
                                    }
                                    var v69 int32 = 1
                                    _ = v69
                                }
                                _ = v67
                            }
                            _ = fi65
                        }
                    } else {
                        var v70 float32 = -2.25
                        v71 := GenS9{F0: (v2 ^ 100)}
                        fmt.Println(v42)
                        _ = v70
                        _ = v71
                    }
                }
                _ = fi61
            }
            _ = v60
        } else {
            v72 := GenS7{F0: v40, F1: 255}
            var v73 int8 = v2
            _ = v72
            _ = v73
        }
        var v74 int32 = (^(v1))
        fmt.Println(v49)
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
