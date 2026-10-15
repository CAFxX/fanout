package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 string; F2 int16 }
type GenS5 struct { F0 uint32; F1 int; F2 int64 }
type GenS6 struct { F0 string; F1 int32 }
func genFn0() int32 {
    if (false && false) {
        for fi4 := 0; fi4 < 8; fi4++ {
            if (false && (!false)) {
                var v5 int8 = 127
                _ = v5
            } else {
                v6 := []int{}
                _ = v6
            }
            v7 := []int{(^(fi4)), (fi4 / (((^(fi4))) | 1))}
            v8 := GenS1{F0: (fi4 << ((698) & 63))}
            _ = fi4
            _ = v7
            _ = v8
        }
        if true {
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 int16 = 1000
                v11 := GenS1{F0: 4294967295}
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            if (1.5 >= 1e10) {
                var v12 float32 = float32(1e10)
                _ = v12
            } else {
                v13 := []int{127, 9223372036854775807, -744, -1}
                switch "a" {
                case "":
                    var v14 uint16 = 42
                    _ = v14
                case "\n\t":
                    v15 := GenS1{F0: 42}
                    _ = v15
                }
                _ = v13
            }
            var v16 float64 = 0.0
            if (180 <= 128) {
                var v17 uint32 = 1
                var v18 bool = true
                _ = v17
                _ = v18
            } else {
                v16 += v16
                var v19 uint64 = 4294967295
                _ = v19
            }
            _ = v16
        }
    }
    if (-32 >= 1) {
        var v20 float32 = 1.5
        var v21 float32 = (v20 - 0.0)
        fmt.Println(v21, v20)
        _ = v20
        _ = v21
    }
    var v22 int8 = 29
    v23 := GenS1{F0: -221}
    _ = v22
    _ = v23
    return 2147483647
}
func genFn1(p0 string, p1 uint64, p2 int) int8 {
    switch p0 {
    case "\n\t":
        switch 4294967295 {
        case 42:
            fmt.Println(p1, p1)
        case -929:
            fmt.Println(p1, p2)
        default:
            v24 := []int{42, 4294967295, (p2 >> ((255) & 63))}
            _ = v24
        }
    case "foo bar":
        v25 := GenS2{F0: 0, F1: p1, F2: 201}
        _ = v25
    case "":
        var v26 uint32 = 2
        _ = v26
    default:
        for fi27 := 0; fi27 < 15; fi27++ {
            var v28 int = (p2 + p2)
            var v29 int16 = 32767
            p1 = 18446744073709551615
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    v30 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -220
    var v3 uint16 = 65535
    v3 = 1
    var v31 int8 = 82
    if false {
        if (-32768 >= -32768) {
            v2 += v2
            var v32 float64 = 0.1
            _ = v32
        }
        v31 += v31
    } else {
        var v33 float32 = float32(1.5)
        fmt.Println(v3, v31, v31)
        _ = v33
    }
    v34 := GenS2{F0: uint64(v3), F1: 4294967295, F2: 71}
    if true {
        if (true || ("hello" != "foo bar")) {
            if (!true) {
                var v35 string = ("foo bar" + "a")
                _ = v35
            } else {
                var v36 float64 = 3.14159
                _ = v36
            }
        }
        var v37 float64 = ((1.5 * (1e10 - (float64(((0.1 + (float32((1e10 + (0.0 - (((((((1.5 - float64(-2.25)) * 3.14159) / (-2.25 + 0.0)) / 1.0) * float64(((((-2.25 * (1.5 - 1.5)) / 3.14159) / (((-2.25 / 1e10) + (0.0 / -2.25)) + 0.1)) * float32(1e10)))) / float64(3.14159)) - ((0.0 - -2.25) + 1.5))))) + (float32(1e10) + (0.0 / 1.0)))) - 1e10)) * 3.14159))) / (((0.0 * (1e10 / (0.1 + 1e10))) - ((1.5 - 0.1) - float64((float32((0.1 - 0.0)) + 1e10)))) + 3.14159))
        switch ("hello" + (("" + ("\n\t" + "hello")) + "foo bar")) {
        case "\n\t":
            switch ("foo bar" + ("a" + "foo bar")) {
            case "":
                var v38 int32 = genFn0()
                _ = v38
            case "hello":
                delete(v1, 679)
            case "a":
                switch (v2 << ((325) & 63)) {
                case 42:
                    v39 := GenS3{F0: -2147483648}
                    _ = v39
                case 32768:
                    var v40 uint8 = 127
                    _ = v40
                }
            default:
                var v41 bool = (0 > 42)
                _ = v41
            }
        case "a":
            v2 = 4294967295
        }
        _ = v37
    } else {
        if (false || (float64(-2.25) != (0.1 / float64(3.14159)))) {
            for fi42 := 0; fi42 < 8; fi42++ {
                for fi43 := 0; fi43 < 11; fi43++ {
                    var v44 uint8 = 2
                    _ = fi43
                    _ = v44
                }
                v2 -= v2
                _ = len(v1)
                _ = fi42
            }
        }
        var v45 int8 = v31
        v31 += v45
        _ = v45
    }
    switch (^(v2)) {
    case -469:
        for fi46 := 0; fi46 < 14; fi46++ {
            v47 := GenS1{F0: (v2 | fi46)}
            switch (fi46 * (v2 - int(v3))) {
            case 2147483647:
                v48 := []int{v2}
                _ = v48
            default:
                for fi49 := 0; fi49 < 15; fi49++ {
                    if false {
                        var v50 uint16 = v3
                        switch 42 {
                        case -1:
                            fmt.Println(v2, fi49, v3)
                            v51 := GenS1{F0: v2}
                            _ = v51
                        default:
                            v50 -= v50
                        }
                        var v52 bool = false
                        _ = v50
                        _ = v52
                    }
                    v53 := GenS4{F0: 2147483648, F1: ("hello" + (((("hello" + "\n\t") + "\n\t") + "\n\t") + "foo bar")), F2: -212}
                    v54 := GenS3{F0: -2147483648}
                    _ = fi49
                    _ = v53
                    _ = v54
                }
            }
            var v55 float64 = -2.25
            _ = fi46
            _ = v47
            _ = v55
        }
    case 127:
        var v56 string = "foo bar"
        _ = v56
    case 419:
        if (false || true) {
            for fi57 := 0; fi57 < 14; fi57++ {
                v58 := GenS5{F0: 812, F1: fi57, F2: -1}
                v59 := GenS4{F0: 2, F1: ("a" + ""), F2: int16(v31)}
                v31 -= v31
                _ = fi57
                _ = v58
                _ = v59
            }
        } else {
            fmt.Println(v3, v3, v31)
        }
    default:
        v60 := GenS3{F0: 940}
        _ = v60
    }
    v61 := []int{}
    var v62 float64 = -2.25
    v1[1] = (9223372036854775807 == -942)
    if len(v61) > 0 {
        v61[0] = v2
    }
    for fi63 := 0; fi63 < 6; fi63++ {
        fmt.Println(v31, v2)
        var v64 float64 = (v62 + v62)
        var v65 int = 909
        _ = fi63
        _ = v64
        _ = v65
    }
    v66 := GenS6{F0: ("" + "a"), F1: 42}
    if true {
        for fi67 := 0; fi67 < 19; fi67++ {
            var v68 int32 = 0
            _ = fi67
            _ = v68
        }
        if len(v61) >= 2 {
            v61 = v61[1:len(v61)]
        }
    } else {
        for fi69 := 0; fi69 < 1; fi69++ {
            if (42 > 42) {
                v62 -= ((((v62 + (1.5 * v62)) - v62) / ((v62 / float64((0.1 - -2.25))) * 0.0)) / v62)
                v70 := GenS1{F0: (v2 * (-(fi69)))}
                _ = v70
            }
            for fi71 := 0; fi71 < 17; fi71++ {
                v61 = append(v61, fi69)
                _ = fi71
            }
            _ = fi69
        }
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v61
    _ = v62
    _ = v66
}
