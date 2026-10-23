package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 bool; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 int; F1 uint8 }
func genFn0(p0 uint8, p1 int16) int16 {
    p1 = p1
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: int8(p0), F1: 0, F2: 1}
    var v8 int16 = p1
    v9 := GenS1{F0: 42, F1: p0, F2: 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 uint64 = 577
    var v2 int32 = -1
    v3 := []uint32{}
    v4 := make(map[uint32]string)
    v5 := make(map[int]uint32)
    switch "foo bar" {
    case "a":
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    case "foo bar":
        for fi10 := 0; fi10 < 13; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                v2 = (v2 >> ((65535) & 63))
                v2 += (v2 - (-(v2)))
                var v12 uint8 = 109
                _ = fi11
                _ = v12
            }
            delete(v5, fi10)
            _ = fi10
        }
    case "":
        v13 := GenS1{F0: -128, F1: 244, F2: v2}
        _ = v13
    default:
        v1 = (v1 - 42)
    }
    if true {
        switch int(v1) {
        case 9223372036854775807:
            v14 := v4[698]
            _ = v14
        default:
            if false {
                v1 -= 4294967295
                fmt.Println(v2, v1)
            }
        }
    }
    switch "" {
    case "hello":
        fmt.Println(v2)
    case "a":
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v15 float64 = 0.1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        if (255 >= 255) {
            if true {
                v17 := GenS2{F0: true, F1: 0, F2: v1}
                var v18 uint16 = 2
                _ = v17
                _ = v18
            }
            if (((fi16 | -2147483648) == (^(fi16))) || true) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if false {
                    switch fi16 {
                    case -1:
                        v19 := GenS3{F0: 42}
                        _ = v19
                    case 500:
                        _ = len(v3)
                    case 128:
                        v20 := GenS1{F0: 100, F1: 0, F2: (v2 / (((v2 / (((v2 - (v2 / (((v2 | v2)) | 1)))) | 1))) | 1))}
                        _ = v20
                    }
                    switch "a" {
                    case "foo bar":
                        for fi21 := 0; fi21 < 11; fi21++ {
                            fmt.Println(v1, fi21, v2)
                            _ = fi21
                        }
                    }
                } else {
                    for fi22 := 0; fi22 < 19; fi22++ {
                        var v23 int64 = -1
                        var v24 int8 = 0
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 12; fi25++ {
                        if true {
                            if true {
                                v1 = v1
                                var v26 string = ("" + "a")
                                _ = v26
                            }
                            switch "\n\t" {
                            case "":
                                if (false && false) {
                                    fmt.Println(v1)
                                    if (((((v1 | 262) >= v1) || true) || false) && (true || (2147483648 >= 989))) {
                                        if true {
                                            delete(v5, 127)
                                            _ = len(v3)
                                            var v27 float32 = 0.1
                                            _ = v27
                                        } else {
                                            v28 := GenS1{F0: -30, F1: 127, F2: (v2 | 42)}
                                            switch fi25 {
                                            case 128:
                                                v1 = 460
                                            case 9223372036854775807:
                                                var v29 uint16 = 65535
                                                _ = v29
                                            }
                                            _ = v28
                                        }
                                        v30 := GenS3{F0: 237}
                                        fmt.Println(v1, fi16)
                                        _ = v30
                                    }
                                    var v31 uint64 = (v1 | 0)
                                    var v32 bool = (!true)
                                    _ = v31
                                    _ = v32
                                }
                            case "a":
                                v33 := GenS3{F0: 2}
                                _ = v33
                            case "hello":
                                var v34 bool = (true || true)
                                _ = v34
                            default:
                                var v35 int16 = 776
                                _ = v35
                            }
                        }
                        _ = fi25
                    }
                    switch fi16 {
                    case -916:
                        var v36 int32 = 0
                        _ = v36
                    }
                }
                fmt.Println(v2, v2)
            } else {
                for fi37 := 0; fi37 < 19; fi37++ {
                    v38 := GenS4{F0: -1}
                    v1 = 1
                    _ = fi37
                    _ = v38
                }
            }
        } else {
            var v39 int32 = v2
            _ = v39
        }
        if (v2 >= v2) {
            var v40 int = (fi16 | (fi16 | (^(fi16))))
            _ = v40
        }
        v41 := GenS2{F0: false, F1: 2, F2: v1}
        _ = fi16
        _ = v41
    }
    if (3.14159 < 1.5) {
        if true {
            switch "\n\t" {
            case "a":
                switch 629 {
                case -9223372036854775808:
                    fmt.Println(v2)
                case 128:
                    v1 -= (v1 & (^(v1)))
                case 127:
                    v1 = v1
                default:
                    if (0 > 222) {
                        if len(v3) > 0 {
                            v3[0] = 940
                        }
                        var v42 uint = 65535
                        v1 = v1
                        _ = v42
                    } else {
                        delete(v4, 2147483648)
                        if (13 >= 42) {
                            var v43 int = 2147483647
                            var v44 float64 = -2.25
                            _ = v43
                            _ = v44
                        } else {
                            if false {
                                switch ((("a" + "foo bar") + (("\n\t" + ("a" + "hello")) + "\n\t")) + "hello") {
                                case "foo bar":
                                    var v45 uint = 255
                                    _ = v45
                                case "a":
                                    switch -1 {
                                    case 128:
                                        _ = len(v3)
                                    case 127:
                                        if (9223372036854775807 >= -1) {
                                            for fi46 := 0; fi46 < 1; fi46++ {
                                                _ = len(v3)
                                                _ = fi46
                                            }
                                            fmt.Println(v1, v1)
                                            switch "hello" {
                                            case "":
                                                v2 -= v2
                                            default:
                                                for fi47 := 0; fi47 < 6; fi47++ {
                                                    for fi48 := 0; fi48 < 5; fi48++ {
                                                        v49 := GenS5{F0: 255, F1: 2147483647, F2: -1}
                                                        var v50 string = (("foo bar" + "hello") + (((("hello" + "hello") + ("a" + (("hello" + "hello") + "a"))) + ((((("foo bar" + ("\n\t" + "hello")) + "a") + "") + ("" + "foo bar")) + "hello")) + "a"))
                                                        var v51 float64 = (1e10 + 0.1)
                                                        _ = fi48
                                                        _ = v49
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    _ = fi47
                                                }
                                            }
                                        } else {
                                            var v52 int64 = 0
                                            _ = v52
                                        }
                                    }
                                default:
                                    switch "\n\t" {
                                    case "foo bar":
                                        switch 2147483647 {
                                        case -9223372036854775808:
                                            v3 = append(v3, 4294967295)
                                        }
                                    case "\n\t":
                                        delete(v4, 2147483647)
                                    case "":
                                        var v53 uint32 = 65535
                                        _ = v53
                                    }
                                }
                                var v54 int32 = 1
                                _ = v54
                            } else {
                                var v55 uint8 = 0
                                if false {
                                    _ = len(v3)
                                } else {
                                    _ = len(v3)
                                    v2 -= (v2 ^ 2147483647)
                                }
                                v3 = append(v3, 0)
                                _ = v55
                            }
                            fmt.Println(v1, v2)
                        }
                    }
                }
            case "hello":
                v56 := GenS4{F0: 9223372036854775807}
                _ = v56
            case "foo bar":
                v57 := GenS4{F0: -680}
                _ = v57
            default:
                for fi58 := 0; fi58 < 2; fi58++ {
                    for fi59 := 0; fi59 < 3; fi59++ {
                        if (!(!true)) {
                            var v60 uint16 = 689
                            _ = v60
                        } else {
                            v3 = append(v3, 1)
                        }
                        if false {
                            var v61 uint = 42
                            v61 = (v61 + (v61 % ((736) | 1)))
                            _ = v61
                        } else {
                            fmt.Println(fi59, fi58, v1)
                            v2 += v2
                            v62 := GenS6{F0: 42, F1: 0}
                            _ = v62
                        }
                        var v63 int8 = 42
                        v64 := GenS7{F0: v2, F1: 34}
                        _ = fi59
                        _ = v63
                        _ = v64
                    }
                    for fi65 := 0; fi65 < 11; fi65++ {
                        fmt.Println(fi58, fi58)
                        _ = len(v3)
                        var v66 float64 = float64(0.0)
                        _ = fi65
                        _ = v66
                    }
                    _ = fi58
                }
            }
            v67 := GenS7{F0: -1, F1: 255}
            _ = v67
        }
    }
    var v68 uint = 2
    v69 := GenS8{F0: 127, F1: 0}
    switch -906 {
    case -297:
        fmt.Println(v1, v1, v1)
    }
    var v70 uint = (^(v68))
    for fi71 := 0; fi71 < 12; fi71++ {
        var v72 float32 = 1.5
        if true {
            v72 = ((v72 + 0.0) + 1e10)
            v1 += (v1 ^ 1)
            var v73 string = ""
            _ = v73
        } else {
            v70 = v70
            v70 += (v70 * v70)
            for fi74 := 0; fi74 < 18; fi74++ {
                switch fi71 {
                case 4294967295:
                    var v75 string = ""
                    _ = v75
                }
                delete(v4, 18)
                _ = fi74
            }
        }
        _ = fi71
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v68:", v68)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
    _ = v70
}
