package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int; F1 uint32; F2 int16 }
type GenS6 struct { F0 bool; F1 bool; F2 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 int64; F1 int; F2 int }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint8 }
func genFn0(p0 uint, p1 int8, p2 int) uint16 {
    v5 := GenS1{F0: 0}
    v6 := GenS2{F0: (p0 | 486)}
    p1 -= p1
    var v7 uint = 2
    switch ("a" + "") {
    case "foo bar":
        switch "hello" {
        case "foo bar":
            var v8 uint64 = 4294967295
            _ = v8
        case "":
            if true {
                v9 := []int{0, 828, 127}
                _ = v9
            }
        default:
            v10 := []int{(p2 << (((v7 % (((v7 ^ 600)) | 1))) & 63)), 1, 127, (p2 & 2147483647)}
            _ = v10
        }
    case "":
        var v11 int32 = 2147483647
        _ = v11
    case "hello":
        v12 := GenS1{F0: 1}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 680
}
func genFn1(p0 bool, p1 uint64) int32 {
    v13 := GenS3{F0: 0, F1: 1}
    p1 = 2
    var v14 int32 = 699
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := []int{(fi15 | (fi15 & 32768)), 0, fi15, 127}
        for fi17 := 0; fi17 < 9; fi17++ {
            fmt.Println(p0, p1)
            v18 := GenS3{F0: 898, F1: 255}
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    var v19 uint64 = p1
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v19
    return (v14 & (-(v14)))
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]int)
    v3 := []int{-1, 127}
    v4 := make(map[string]bool)
    for fi20 := 0; fi20 < 4; fi20++ {
        fmt.Println(fi20, fi20)
        var v21 uint16 = 967
        _ = fi20
        _ = v21
    }
    switch ("\n\t" + ("\n\t" + "\n\t")) {
    case "hello":
        delete(v2, 524)
    case "foo bar":
        v1 += 0
    default:
        fmt.Println(v1)
    }
    if (!true) {
        var v22 bool = (false && (!false))
        if (v1 > (v1 % ((v1) | 1))) {
            switch 4294967295 {
            case 42:
                switch 32768 {
                case -9223372036854775808:
                    if (true || (1e10 == ((1e10 - 3.14159) + 0.1))) {
                        var v23 float64 = 1e10
                        v23 = (0.0 / ((0.1 * v23) + v23))
                        switch "a" {
                        case "a":
                            v22 = true
                        default:
                            if (v22 || v22) {
                                for fi24 := 0; fi24 < 14; fi24++ {
                                    for fi25 := 0; fi25 < 6; fi25++ {
                                        if true {
                                            v1 = v1
                                        }
                                        v23 += float64(1e10)
                                        _ = fi25
                                    }
                                    _ = fi24
                                }
                                v4[("" + "\n\t")] = (!(v22 && (float32((float64(3.14159) * v23)) > 3.14159)))
                            } else {
                                for fi26 := 0; fi26 < 3; fi26++ {
                                    var v27 int16 = 1
                                    v2[2147483648] = fi26
                                    _ = fi26
                                    _ = v27
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v22 = v22
                            }
                        }
                        _ = v23
                    } else {
                        v28 := GenS1{F0: -32768}
                        switch 4294967295 {
                        case 128:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 4294967295:
                            switch 0 {
                            case -1:
                                _ = len(v3)
                            default:
                                v22 = v22
                            }
                        case 748:
                            _ = len(v3)
                        }
                        for fi29 := 0; fi29 < 5; fi29++ {
                            fmt.Println(v1, fi29)
                            _ = len(v3)
                            fmt.Println(v1, v1)
                            _ = fi29
                        }
                        _ = v28
                    }
                }
            }
        } else {
            v3 = append(v3, 42)
            var v30 bool = (0 < 32768)
            _ = v30
        }
        v1 += (v1 | 2)
        _ = v22
    }
    var v31 float64 = float64((1e10 / 1.0))
    v1 -= 42
    v31 = v31
    var v32 bool = (32767 == 2)
    v31 = (3.14159 + float64(((((0.0 + 0.1) / 3.14159) / 3.14159) * 1.5)))
    if v32 {
        switch -2147483648 {
        case -737:
            v33 := GenS1{F0: 42}
            _ = v33
        case 0:
            var v34 int8 = 1
            _ = v34
        case -907:
            switch (("a" + "\n\t") + "\n\t") {
            case "":
                var v35 uint16 = 1
                _ = v35
            default:
                var v36 int = 128
                _ = v36
            }
        default:
            v37 := GenS4{F0: (v1 - v1), F1: 127, F2: (v1 & v1)}
            _ = v37
        }
        if (1.5 > ((3.14159 * (float32((float64(3.14159) / v31)) / (1e10 * ((float32(0.0) / ((3.14159 / 0.1) * (-2.25 / 1.0))) - -2.25)))) - (-2.25 + (1.5 * ((1.5 / (3.14159 * float32((v31 / v31)))) - float32(v31)))))) {
            v31 += float64(((1e10 + (0.1 + 0.1)) / float32((float64((0.1 + 3.14159)) + ((v31 / v31) + v31)))))
        } else {
            v31 += 3.14159
        }
        v31 = (float64(((-2.25 / (3.14159 * 1e10)) * 0.1)) * v31)
    } else {
        var v38 float32 = ((0.0 + (((3.14159 + 0.1) * 0.1) * 0.0)) * float32(((v31 + (v31 / ((float64(-2.25) + v31) * (0.1 * float64(0.1))))) * v31)))
        if v32 {
            v39 := GenS5{F0: 9223372036854775807, F1: 1, F2: 1000}
            fmt.Println(v38, v32)
            var v40 float64 = v31
            _ = v39
            _ = v40
        } else {
            var v41 int16 = 1000
            v42 := GenS6{F0: false, F1: true, F2: -105}
            _ = v41
            _ = v42
        }
        for fi43 := 0; fi43 < 8; fi43++ {
            fmt.Println(v32, fi43, fi43)
            var v44 uint32 = 2147483647
            _ = fi43
            _ = v44
        }
        _ = v38
    }
    v31 += (v31 - v31)
    v45 := GenS7{F0: 255}
    v46 := GenS8{F0: true, F1: -92}
    v47 := GenS7{F0: 42}
    for fi48 := 0; fi48 < 6; fi48++ {
        for fi49 := 0; fi49 < 11; fi49++ {
            v32 = (!v32)
            if (!v32) {
                switch (("\n\t" + (("hello" + "a") + "hello")) + "foo bar") {
                case "a":
                    v32 = true
                default:
                    if (!v32) {
                        delete(v2, 42)
                        if (v32 || ((v1 >> ((65535) & 63)) >= v1)) {
                            v50 := GenS6{F0: false, F1: (42 <= -1), F2: -128}
                            _ = v50
                        }
                    } else {
                        if (2 <= 416) {
                            v51 := GenS8{F0: v32, F1: 1}
                            if (int32(v1) > -536) {
                                var v52 string = "hello"
                                _ = v52
                            }
                            _ = v51
                        } else {
                            var v53 float32 = (((1.5 / 1.0) * 1e10) - ((-2.25 - 1e10) * -2.25))
                            var v54 int = 0
                            switch (v54 - fi49) {
                            case 32768:
                                v1 = v1
                            case -154:
                                if (0 >= 825) {
                                    v55 := v2[4294967295]
                                    _ = v55
                                } else {
                                    switch -402 {
                                    case 873:
                                        v2[2147483647] = (fi48 << ((507) & 63))
                                        v31 += (v31 * 1e10)
                                        v54 = (fi48 >> ((443) & 63))
                                    case 651:
                                        if (0 >= -87) {
                                            if (v32 || true) {
                                                v53 += ((1.5 + v53) - 0.0)
                                                v32 = (!(v32 && v32))
                                                v56 := GenS1{F0: -1}
                                                v3 = append(v3, 128)
                                                _ = v56
                                            }
                                            var v57 uint8 = (^(v1))
                                            _ = v57
                                        }
                                    default:
                                        var v58 int8 = -1
                                        var v59 uint = 970
                                        _ = v58
                                        _ = v59
                                    }
                                }
                            case -264:
                                v60 := GenS9{F0: 460, F1: v54, F2: (-(v54))}
                                _ = v60
                            default:
                                switch fi48 {
                                case 83:
                                    if (759 <= 858) {
                                        var v61 int = fi49
                                        switch ("a" + ((("foo bar" + "hello") + ("hello" + ("" + "foo bar"))) + "\n\t")) {
                                        case "\n\t":
                                            fmt.Println(v61, v32)
                                        default:
                                            var v62 int8 = -1
                                            _ = v62
                                        }
                                        _ = v61
                                    }
                                }
                            }
                            _ = v53
                            _ = v54
                        }
                        fmt.Println(v32)
                        switch ("a" + "") {
                        case "foo bar":
                            fmt.Println(fi48, v32, fi49)
                        case "\n\t":
                            var v63 int32 = -2147483648
                            _ = v63
                        case "":
                            if (v31 > ((((v31 * 1.5) / v31) - 3.14159) * v31)) {
                                _ = len(v3)
                                fmt.Println(fi49)
                                if false {
                                    var v64 int8 = int8(fi49)
                                    _ = v64
                                } else {
                                    v31 = 0.1
                                    v31 += v31
                                }
                                v32 = (false || v32)
                                v65 := GenS10{F0: 63, F1: 0}
                                _ = v65
                            }
                        default:
                            if ((-781 == 0) || false) {
                                var v66 uint = 18446744073709551615
                                var v67 string = "foo bar"
                                _ = v66
                                _ = v67
                            }
                            var v68 uint = 42
                            _ = v68
                        }
                        if true {
                            var v69 uint16 = 390
                            var v70 bool = (112 != 4294967295)
                            v71 := GenS11{F0: 42}
                            _ = v69
                            _ = v70
                            _ = v71
                        }
                    }
                }
            }
            _ = fi49
        }
        v72 := GenS12{F0: (^(v1))}
        _ = fi48
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v45
    _ = v46
    _ = v47
}
