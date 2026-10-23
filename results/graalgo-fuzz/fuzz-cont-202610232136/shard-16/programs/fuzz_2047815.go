package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint8; F1 int16; F2 bool }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int32 }
func genFn0(p0 int64) int16 {
    if true {
        v7 := GenS1{F0: 42, F1: 42}
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        switch fi8 {
        case 2147483647:
            var v9 int8 = 100
            _ = v9
        case 1:
            v10 := GenS2{F0: -128, F1: -9223372036854775808}
            _ = v10
        case 32767:
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 uint64 = 460
                _ = fi11
                _ = v12
            }
        }
        var v13 int8 = 1
        var v14 uint32 = 4294967295
        _ = fi8
        _ = v13
        _ = v14
    }
    p0 -= p0
    for fi15 := 0; fi15 < 3; fi15++ {
        if true {
            var v16 uint32 = 42
            v17 := GenS1{F0: v16, F1: 1}
            _ = v16
            _ = v17
        } else {
            p0 = -9223372036854775808
            var v18 string = "hello"
            p0 = p0
            _ = v18
        }
        v19 := GenS3{F0: 65535, F1: 0}
        _ = fi15
        _ = v19
    }
    v20 := []int{4294967295}
    _ = p0
    _ = v20
    return -32768
}
func genFn1(p0 string, p1 uint16, p2 uint16) uint8 {
    p1 = (^(p1))
    for fi21 := 0; fi21 < 17; fi21++ {
        switch (fi21 & (^(fi21))) {
        case -2147483648:
            if false {
                p0 = "foo bar"
                for fi22 := 0; fi22 < 15; fi22++ {
                    v23 := GenS4{F0: int64(p1), F1: 100}
                    var v24 int32 = 2147483647
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            } else {
                var v25 bool = true
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 float64 = (1.5 + 0.0)
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
        default:
            var v28 int = fi21
            _ = v28
        }
        _ = fi21
    }
    p0 = "hello"
    if true {
        v29 := []int{24, -1, 127}
        switch p0 {
        case "a":
            p0 = p0
        case "foo bar":
            if true {
                var v30 uint = 65535
                _ = v30
            }
        case "\n\t":
            var v31 int16 = int16(p1)
            _ = v31
        default:
            var v32 float32 = 0.0
            _ = v32
        }
        _ = v29
    } else {
        if ((((-2.25 + float64(((1e10 - 0.0) - ((1.5 - 3.14159) - -2.25)))) + (float64((1.5 - 0.1)) * 0.0)) * (1.5 * 1.5)) <= (1e10 / 1e10)) {
            if (42 <= 2147483647) {
                v33 := []int{-9223372036854775808, 9223372036854775807, -2147483648}
                _ = v33
            } else {
                p2 -= p2
            }
        } else {
            var v34 string = ((("hello" + "foo bar") + "hello") + "hello")
            if true {
                v35 := GenS2{F0: -128, F1: -9223372036854775808}
                _ = v35
            } else {
                p2 = p2
                if true {
                    p2 -= p2
                    v36 := GenS5{F0: 128, F1: 32767, F2: (!false)}
                    _ = v36
                }
                var v37 uint8 = 157
                _ = v37
            }
            p1 += (^(p1))
            _ = v34
        }
        switch 142 {
        case -235:
            p0 = p0
        }
        for fi38 := 0; fi38 < 16; fi38++ {
            p2 = p1
            for fi39 := 0; fi39 < 4; fi39++ {
                if (-32768 == genFn0(0)) {
                    switch fi39 {
                    case 4294967295:
                        if (0 == 1) {
                            switch (("" + p0) + p0) {
                            case "hello":
                                switch fi39 {
                                case 0:
                                    p2 = (p1 | 869)
                                case 32767:
                                    p2 -= p1
                                    var v40 float32 = float32((1e10 - 1e10))
                                    v40 -= v40
                                    _ = v40
                                default:
                                    fmt.Println(fi39, fi39, p0)
                                }
                            case "foo bar":
                                if (244 == (p2 - p1)) {
                                    fmt.Println(fi38, fi38, p0)
                                } else {
                                    p0 = p0
                                    var v41 string = p0
                                    _ = v41
                                }
                            default:
                                var v42 string = "foo bar"
                                _ = v42
                            }
                            fmt.Println(fi39, p0, p2)
                            if (!(!(!((!false) || (true && false))))) {
                                var v43 int64 = 677
                                p0 = p0
                                _ = v43
                            } else {
                                fmt.Println(p2, fi38, p0)
                            }
                        } else {
                            p2 = p1
                            if (0.0 >= -2.25) {
                                fmt.Println(fi39)
                            } else {
                                var v44 int64 = 76
                                v45 := GenS6{F0: true, F1: 42}
                                _ = v44
                                _ = v45
                            }
                            p0 = ("\n\t" + "")
                        }
                    default:
                        var v46 uint32 = 800
                        var v47 uint32 = (v46 * v46)
                        _ = v46
                        _ = v47
                    }
                    var v48 int16 = 1
                    fmt.Println(fi38, v48, p1)
                    _ = v48
                }
                v49 := []int{-9223372036854775808, fi38}
                v50 := GenS7{F0: 1000}
                _ = fi39
                _ = v49
                _ = v50
            }
            fmt.Println(p0, p2)
            _ = fi38
        }
    }
    var v51 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v51
    return 255
}

func main() {
    var v1 int = 4294967295
    var v2 string = "\n\t"
    v3 := make(map[int]uint32)
    var v4 string = ("hello" + v2)
    var v5 bool = (!(true && false))
    var v6 uint64 = 42
    var v52 int8 = 42
    if v5 {
        if (v5 || v5) {
            v3[v1] = 4294967295
            v52 = v52
            var v53 bool = v5
            _ = v53
        } else {
            var v54 int8 = -8
            v5 = v5
            for fi55 := 0; fi55 < 19; fi55++ {
                v56 := GenS2{F0: v52, F1: -70}
                v4 = "\n\t"
                _ = fi55
                _ = v56
            }
            _ = v54
        }
        if (((3.14159 * (0.1 + (1.5 / (1.5 / -2.25)))) / ((1.5 - 0.0) * 0.1)) >= (3.14159 * 1.5)) {
            for fi57 := 0; fi57 < 18; fi57++ {
                var v58 uint8 = 224
                if (true || v5) {
                    v3[-1] = 4294967295
                    v58 -= v58
                    var v59 uint8 = 1
                    _ = v59
                } else {
                    v52 += (v52 >> ((18446744073709551615) & 63))
                }
                var v60 uint16 = 255
                _ = fi57
                _ = v58
                _ = v60
            }
            v4 = v2
        }
        v2 = "a"
    }
    if v5 {
        fmt.Println(v5, v4)
        var v61 int32 = 1
        for fi62 := 0; fi62 < 10; fi62++ {
            v6 = 1
            _ = fi62
        }
        _ = v61
    }
    if (v52 != v52) {
        var v63 int16 = 814
        for fi64 := 0; fi64 < 15; fi64++ {
            for fi65 := 0; fi65 < 12; fi65++ {
                var v66 int16 = (v63 % ((v63) | 1))
                v67 := GenS8{F0: (v63 | v63)}
                v5 = (331 < 42)
                _ = fi65
                _ = v66
                _ = v67
            }
            _ = fi64
        }
        v68 := GenS4{F0: 892, F1: v52}
        _ = v63
        _ = v68
    }
    var v69 string = (("hello" + (v2 + "\n\t")) + "")
    v70 := []int{128, v1, v1}
    var v71 uint8 = 127
    v72 := GenS9{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v70:", len(v70), v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
