package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint; F2 bool }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint32) int32 {
    v7 := GenS1{F0: 0}
    fmt.Println(p0)
    if false {
        if (1.5 < float32(1e10)) {
            var v8 uint8 = 128
            if true {
                for fi9 := 0; fi9 < 8; fi9++ {
                    fmt.Println(p0, p0)
                    _ = fi9
                }
                var v10 int16 = -32768
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 uint32 = 597
                    v13 := GenS2{F0: 2147483647, F1: 0, F2: (4294967295 == 2)}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = v10
            } else {
                var v14 int8 = 42
                v15 := []int{4294967295, 841, -692, 1}
                p0 -= (p0 + 625)
                _ = v14
                _ = v15
            }
            switch "foo bar" {
            case "foo bar":
                var v16 uint16 = 231
                _ = v16
            case "":
                fmt.Println(p0, v8)
            case "a":
                var v17 uint64 = 837
                _ = v17
            default:
                for fi18 := 0; fi18 < 15; fi18++ {
                    v8 = (v8 + (v8 - 1))
                    v8 = v8
                    var v19 string = ("foo bar" + "a")
                    _ = fi18
                    _ = v19
                }
            }
            _ = v8
        } else {
            if (false && true) {
                var v20 int32 = 113
                if (!true) {
                    v21 := []int{32768}
                    if false {
                        p0 += (p0 | 474)
                        if false {
                            if true {
                                p0 -= 2147483647
                            }
                            var v22 uint8 = 42
                            _ = v22
                        }
                    } else {
                        v23 := GenS1{F0: v20}
                        fmt.Println(v20, v20)
                        fmt.Println(p0, v20)
                        _ = v23
                    }
                    var v24 uint16 = 750
                    _ = v21
                    _ = v24
                }
                if ((false || (false && false)) || (40 >= -265)) {
                    var v25 uint64 = 2
                    p0 += (p0 + (p0 * 2147483647))
                    _ = v25
                }
                _ = v20
            } else {
                fmt.Println(p0)
                v26 := []int{-56, 4294967295, 217}
                var v27 uint16 = 42
                _ = v26
                _ = v27
            }
        }
        fmt.Println(p0)
        p0 = p0
    }
    var v28 string = ((("hello" + ("" + "hello")) + "") + "\n\t")
    if false {
        p0 += (p0 | p0)
        fmt.Println(v28, p0)
        switch v28 {
        case "foo bar":
            var v29 float32 = 3.14159
            _ = v29
        case "a":
            var v30 uint32 = p0
            _ = v30
        default:
            var v31 uint = 2
            _ = v31
        }
    } else {
        if (!false) {
            for fi32 := 0; fi32 < 12; fi32++ {
                var v33 uint8 = 255
                var v34 int64 = 42
                switch 1 {
                case 127:
                    v35 := []int{(fi32 | (-(fi32)))}
                    _ = v35
                default:
                    var v36 uint8 = v33
                    _ = v36
                }
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        var v37 int16 = 1000
        _ = v37
    }
    _ = p0
    _ = v28
    _ = v7
    return 0
}
func genFn1(p0 bool, p1 uint8, p2 uint32) int32 {
    if p0 {
        if p0 {
            v38 := GenS1{F0: 2147483647}
            _ = v38
        }
        v39 := GenS1{F0: 1}
        var v40 uint16 = 2
        _ = v39
        _ = v40
    }
    if (p0 || p0) {
        for fi41 := 0; fi41 < 14; fi41++ {
            v42 := []int{}
            var v43 uint16 = uint16(p1)
            var v44 uint = 18446744073709551615
            _ = fi41
            _ = v42
            _ = v43
            _ = v44
        }
        for fi45 := 0; fi45 < 2; fi45++ {
            p2 = p2
            _ = fi45
        }
        p0 = p0
    } else {
        var v46 string = ("hello" + "a")
        _ = v46
    }
    v47 := []int{-204}
    if p0 {
        var v48 int8 = 45
        var v49 string = ""
        v50 := GenS1{F0: 329}
        _ = v48
        _ = v49
        _ = v50
    }
    if p0 {
        v47 = append(v47, -338)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v47
    return -344
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 1
    var v3 string = (("\n\t" + "foo bar") + ("foo bar" + ("a" + "\n\t")))
    v4 := make(map[uint32]int)
    v5 := make(map[string]bool)
    v6 := make(map[int]uint32)
    v3 = "foo bar"
    for fi51 := 0; fi51 < 18; fi51++ {
        v2 -= (v2 + (-(v2)))
        _ = fi51
    }
    var v52 uint8 = 255
    var v53 string = (v3 + (v3 + v3))
    if false {
        switch v53 {
        case "hello":
            switch 0 {
            case 32768:
                switch v3 {
                case "hello":
                    v54 := GenS3{F0: -130, F1: genFn1((!(((1.5 + ((-2.25 - ((0.0 / 1.5) - 3.14159)) - 1.5)) - 1.5) != 1.5)), 255, 2147483647)}
                    _ = v54
                case "foo bar":
                    for fi55 := 0; fi55 < 1; fi55++ {
                        v56 := GenS4{F0: false, F1: 42}
                        _ = fi55
                        _ = v56
                    }
                case "":
                    var v57 uint8 = (^(v52))
                    _ = v57
                }
            case 4294967295:
                var v58 uint32 = 65535
                _ = v58
            case 2147483647:
                v59 := GenS1{F0: genFn1((!(false || true)), (^(v52)), 42)}
                _ = v59
            default:
                switch "\n\t" {
                case "\n\t":
                    var v60 int = 1
                    _ = v60
                case "a":
                    for fi61 := 0; fi61 < 17; fi61++ {
                        fmt.Println(fi61, v52, v2)
                        v62 := GenS5{F0: 42}
                        _ = fi61
                        _ = v62
                    }
                default:
                    switch 127 {
                    case -2147483648:
                        v2 = -128
                    case 42:
                        for fi63 := 0; fi63 < 13; fi63++ {
                            var v64 uint8 = v52
                            var v65 int = 0
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                    case -906:
                        var v66 float64 = 1e10
                        _ = v66
                    default:
                        var v67 string = ("a" + v53)
                        _ = v67
                    }
                }
            }
        }
    }
    v68 := GenS6{F0: -32768, F1: v52}
    v69 := GenS6{F0: 1, F1: (^(v52))}
    var v70 bool = (float64(((float32(float64((-2.25 + (((-2.25 * ((0.1 * 1e10) - (-2.25 - 0.0))) / (3.14159 * 1e10)) - 1.5)))) - 1.5) / 1.5)) > 3.14159)
    var v71 bool = (1 != 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v68
    _ = v69
    _ = v70
    _ = v71
}
