package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint64, p1 int8) bool {
    var v4 int = -9223372036854775808
    var v5 int64 = 0
    if (255 <= 1) {
        p0 -= 4294967295
        var v6 uint32 = 579
        _ = v6
    }
    var v7 uint = 0
    var v8 int64 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 uint16, p1 uint16) string {
    p0 = (p0 - p1)
    switch 127 {
    case 0:
        if ((^(p0)) != (p1 - p0)) {
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := []int{574, (fi9 % ((fi9) | 1))}
                for fi11 := 0; fi11 < 10; fi11++ {
                    fmt.Println(p0, fi11)
                    _ = fi11
                }
                for fi12 := 0; fi12 < 9; fi12++ {
                    v13 := GenS1{F0: 38, F1: false}
                    if (true && (false && true)) {
                        v14 := GenS2{F0: 954, F1: 4294967295, F2: 624}
                        if len(v10) > 0 {
                            v10[0] = (fi12 + 0)
                        }
                        _ = v14
                    } else {
                        v15 := GenS2{F0: 369, F1: 816, F2: 1}
                        p1 += 42
                        p0 = 710
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
            p0 += (p1 - p1)
        } else {
            fmt.Println(p1, p0)
            v16 := GenS1{F0: 127, F1: false}
            _ = v16
        }
    }
    switch 197 {
    case 128:
        v17 := []int{127, 143}
        _ = v17
    case 395:
        p1 -= (p0 / ((p0) | 1))
    }
    _ = p0
    _ = p1
    return "a"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{2147483647, 2147483648, 2147483647, 104}
    v3 := []int{-1, 128, 4294967295, 32767}
    v18 := GenS3{F0: 108, F1: 120}
    switch 928 {
    case 128:
        switch "a" {
        case "":
            var v19 int = 2147483647
            _ = v19
        case "hello":
            if (-83 <= 476) {
                if false {
                    v20 := GenS3{F0: 42, F1: 42}
                    var v21 float64 = (1e10 + 0.1)
                    _ = v20
                    _ = v21
                } else {
                    var v22 uint16 = 2
                    v22 += (^(v22))
                    v23 := GenS4{F0: 1, F1: 0}
                    _ = v22
                    _ = v23
                }
                var v24 string = ("a" + "hello")
                switch (v24 + v24) {
                case "foo bar":
                    switch -2147483648 {
                    case 128:
                        v25 := GenS2{F0: 42, F1: 4294967295, F2: 42}
                        _ = v25
                    case 9223372036854775807:
                        v24 = v24
                    }
                case "\n\t":
                    for fi26 := 0; fi26 < 2; fi26++ {
                        var v27 int8 = 0
                        switch 32767 {
                        case 599:
                            var v28 int16 = 1000
                            _ = v28
                        default:
                            switch (fi26 + (fi26 | fi26)) {
                            case 213:
                                var v29 float64 = 0.1
                                _ = v29
                            case 128:
                                v30 := GenS2{F0: 108, F1: 42, F2: 2147483647}
                                _ = v30
                            case -828:
                                v24 = v24
                            default:
                                var v31 uint = 18446744073709551615
                                _ = v31
                            }
                        }
                        _ = fi26
                        _ = v27
                    }
                case "":
                    fmt.Println(v24)
                default:
                    v24 = v24
                }
                _ = v24
            }
        default:
            if true {
                switch 32768 {
                case -3:
                    for fi32 := 0; fi32 < 17; fi32++ {
                        var v33 uint32 = 2
                        v33 -= (v33 - 267)
                        _ = fi32
                        _ = v33
                    }
                case 4294967295:
                    var v34 uint64 = 18446744073709551615
                    _ = v34
                default:
                    var v35 int16 = -260
                    _ = v35
                }
                if true {
                    var v36 float64 = ((float64(0.0) - (3.14159 + 0.1)) * 0.1)
                    _ = v36
                } else {
                    for fi37 := 0; fi37 < 3; fi37++ {
                        var v38 float64 = (1.5 - 1e10)
                        var v39 uint64 = 57
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    var v40 int64 = 575
                    _ = v40
                }
            }
        }
    case 9223372036854775807:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        switch ("a" + ("" + "foo bar")) {
        case "hello":
            var v41 int16 = 32767
            _ = v41
        case "\n\t":
            var v42 uint32 = 219
            _ = v42
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            v43 := GenS2{F0: 42, F1: 0, F2: 1}
            _ = v43
        }
    }
    var v44 int32 = -341
    v45 := GenS4{F0: (v44 >> ((18446744073709551615) & 63)), F1: 4294967295}
    switch (("" + "\n\t") + "foo bar") {
    case "":
        for fi46 := 0; fi46 < 11; fi46++ {
            v47 := v1[42]
            v44 += v44
            _ = fi46
            _ = v47
        }
    default:
        if (false || (177 > v44)) {
            for fi48 := 0; fi48 < 18; fi48++ {
                v44 = -2147483648
                _ = fi48
            }
            var v49 string = "hello"
            var v50 int = int(v44)
            _ = v49
            _ = v50
        } else {
            for fi51 := 0; fi51 < 19; fi51++ {
                var v52 float64 = 0.1
                v1[685] = 770
                fmt.Println(v52, v44)
                _ = fi51
                _ = v52
            }
            v53 := GenS2{F0: 42, F1: 2147483648, F2: 2147483648}
            _ = len(v1)
            _ = v53
        }
    }
    if (42 != 42) {
        for fi54 := 0; fi54 < 6; fi54++ {
            var v55 int64 = 1
            delete(v1, 757)
            _ = fi54
            _ = v55
        }
        var v56 int64 = 42
        _ = v56
    }
    var v57 uint = 255
    if (32767 == 980) {
        var v58 int = 42
        fmt.Println(v57, v58)
        if false {
            v58 += 9223372036854775807
        } else {
            if false {
                switch "a" {
                case "\n\t":
                    switch "\n\t" {
                    case "foo bar":
                        v1[2] = 400
                    }
                }
                if (2 == 481) {
                    for fi59 := 0; fi59 < 9; fi59++ {
                        var v60 uint = (v57 & v57)
                        _ = fi59
                        _ = v60
                    }
                }
                var v61 int32 = v44
                _ = v61
            }
        }
        _ = v58
    } else {
        var v62 string = (("" + "\n\t") + "a")
        _ = v62
    }
    _ = len(v1)
    var v63 uint32 = 735
    var v64 uint64 = 0
    var v65 bool = (v64 > v64)
    var v66 uint64 = v64
    var v67 uint = 42
    v68 := GenS5{F0: 105}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    fmt.Println("v57:", v57)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v57
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
