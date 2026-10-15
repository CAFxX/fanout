package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint32; F2 uint32 }

func main() {
    var v1 int = -462
    v2 := make(map[int]bool)
    var v3 float64 = 1e10
    v4 := []int{-2147483648, (v1 >> ((4294967295) & 63))}
    var v5 int8 = 100
    var v6 uint32 = 206
    _ = len(v2)
    var v7 int16 = 42
    if false {
        var v8 int16 = -818
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 9; fi10++ {
                if true {
                    fmt.Println(fi10, fi9, v7)
                }
                if true {
                    v1 -= (fi10 | (fi10 ^ fi10))
                    v6 = (v6 | (v6 ^ v6))
                    var v11 uint16 = 0
                    var v12 int = (v1 ^ (v1 + fi9))
                    v6 = v6
                    if (((1e10 / 3.14159) <= 1e10) || (v3 != (v3 + ((1.5 / v3) * (float64(1e10) / v3))))) {
                        if (!(303 >= -160)) {
                            v1 -= 845
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        var v13 uint64 = 18446744073709551615
                        var v14 uint16 = v11
                        _ = v13
                        _ = v14
                    }
                    var v15 uint16 = v11
                    _ = v11
                    _ = v12
                    _ = v15
                }
                if (true || true) {
                    if true {
                        var v16 bool = false
                        fmt.Println(v8)
                        _ = v16
                    } else {
                        _ = len(v2)
                        var v17 int8 = (-(v5))
                        if len(v4) > 0 {
                            v4[0] = v1
                        }
                        _ = v17
                    }
                    var v18 int64 = 1
                    var v19 uint32 = (v6 ^ (v6 * 4294967295))
                    _ = v18
                    _ = v19
                } else {
                    var v20 string = ((("foo bar" + "a") + "foo bar") + "\n\t")
                    var v21 uint32 = (v6 * (^(v6)))
                    _ = v20
                    _ = v21
                }
                fmt.Println(v6, v1, fi9)
                _ = fi10
            }
            _ = fi9
        }
        v22 := GenS1{F0: 1, F1: 625}
        _ = v22
        _ = v8
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 19; fi24++ {
            _ = len(v2)
            delete(v2, (fi24 + -2147483648))
            _ = fi24
        }
        v7 = (v7 - 217)
        var v25 uint32 = (v6 & v6)
        _ = fi23
        _ = v25
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := GenS1{F0: 255, F1: 561}
        _ = fi26
        _ = v27
    }
    if len(v4) > 0 {
        v4[0] = (v1 << ((255) & 63))
    }
    if (("foo bar" != "hello") || (!false)) {
        v28 := GenS1{F0: 127, F1: 1}
        v29 := GenS1{F0: 1, F1: 42}
        fmt.Println(v3, v7)
        _ = v28
        _ = v29
    }
    v30 := GenS2{F0: 94, F1: ""}
    for fi31 := 0; fi31 < 6; fi31++ {
        v5 = -128
        for fi32 := 0; fi32 < 12; fi32++ {
            if false {
                fmt.Println(v6, v3, v1)
                v6 = 2147483648
                fmt.Println(v7, fi31, v7)
            }
            _ = fi32
        }
        if (!(127 < v5)) {
            _ = len(v2)
        }
        _ = fi31
    }
    var v33 int8 = (-(v5))
    switch v1 {
    case 792:
        var v34 uint32 = 2147483647
        _ = v34
    default:
        var v35 bool = (true || false)
        _ = v35
    }
    if (v6 == v6) {
        fmt.Println(v1, v7, v1)
    }
    switch -823 {
    case 128:
        v3 += 3.14159
    case 32768:
        var v36 float32 = (float32(0.1) + 1e10)
        _ = v36
    case 160:
        v33 = (-(v33))
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        for fi38 := 0; fi38 < 16; fi38++ {
            if true {
                var v39 uint64 = 1
                _ = v39
            } else {
                var v40 bool = (0 == -9223372036854775808)
                _ = v40
            }
            switch ("" + "a") {
            case "a":
                var v41 uint16 = 117
                _ = v41
            case "":
                if (!(v3 == 0.0)) {
                    if ((-1 >= 0) || (1.5 == (1e10 + (float32((v3 / v3)) * (0.1 + -2.25))))) {
                        v1 -= (fi37 << ((386) & 63))
                        fmt.Println(v1)
                    } else {
                        var v42 uint8 = 153
                        var v43 int32 = -1
                        switch -11 {
                        case 1:
                            var v44 bool = true
                            _ = v44
                        case 120:
                            v4 = append(v4, (fi37 ^ fi38))
                        }
                        _ = v42
                        _ = v43
                    }
                } else {
                    var v45 int32 = 1
                    v46 := GenS3{F0: (!(75 >= 4294967295))}
                    v3 -= v3
                    _ = v45
                    _ = v46
                }
            default:
                if true {
                    _ = len(v2)
                    v47 := GenS4{F0: 1, F1: 255, F2: 0}
                    var v48 float32 = 0.1
                    switch "" {
                    case "hello":
                        if true {
                            if true {
                                var v49 int64 = 0
                                v7 = v7
                                v5 = (v5 | -128)
                                var v50 uint16 = 42
                                v51 := GenS2{F0: 366, F1: ""}
                                var v52 float32 = 0.1
                                _ = v49
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                            var v53 uint8 = 1
                            if ((-2147483648 >= 674) && (float32(float64((float32(1e10) - -2.25))) >= (0.1 - (v48 / (((float32(0.1) + v48) / ((v48 / float32(0.1)) + -2.25)) / (((float32(v3) + v48) * 1e10) + v48)))))) {
                                var v54 bool = (-1 < -357)
                                v55 := GenS2{F0: 65535, F1: "hello"}
                                switch "foo bar" {
                                case "hello":
                                    v3 = ((v3 * -2.25) * v3)
                                case "foo bar":
                                    fmt.Println(v54, v5, fi38)
                                case "\n\t":
                                    var v56 float32 = 1.5
                                    _ = v56
                                }
                                var v57 uint16 = 176
                                var v58 bool = (228 <= 42)
                                _ = v54
                                _ = v55
                                _ = v57
                                _ = v58
                            }
                            v33 = v5
                            _ = v53
                        } else {
                            v59 := GenS2{F0: 2, F1: ("hello" + "hello")}
                            _ = v59
                        }
                        if (!(!(((0.0 - (0.1 - 1e10)) <= v48) && ((("\n\t" + "foo bar") > "\n\t") && true)))) {
                            var v60 uint16 = 0
                            if (int8(fi37) != v5) {
                                var v61 int64 = -166
                                v61 -= v61
                                _ = v61
                            } else {
                                if ((!(!(!(!(1 != 255))))) && (true && false)) {
                                    switch "hello" {
                                    case "a":
                                        v48 = 0.0
                                    case "foo bar":
                                        v6 += (v6 | 804)
                                    case "":
                                        v33 -= 0
                                    }
                                } else {
                                    var v62 float32 = v48
                                    _ = len(v4)
                                    fmt.Println(v1)
                                    _ = v62
                                }
                            }
                            var v63 int = -1
                            _ = v60
                            _ = v63
                        } else {
                            v33 = -1
                            _ = len(v4)
                            var v64 int = (v1 - v1)
                            _ = v64
                        }
                        var v65 int = fi37
                        _ = v65
                    case "":
                        v66 := GenS5{F0: "foo bar"}
                        _ = v66
                    case "a":
                        var v67 uint32 = v6
                        _ = v67
                    }
                    var v68 int16 = (v7 & 1)
                    _ = v47
                    _ = v48
                    _ = v68
                } else {
                    if true {
                        v5 = 100
                    }
                    v69 := GenS6{F0: true, F1: v6, F2: v6}
                    _ = v69
                }
            }
            var v70 uint32 = 563
            _ = fi38
            _ = v70
        }
        switch "foo bar" {
        case "":
            fmt.Println(v33, v5)
        default:
            fmt.Println(v7, v3, fi37)
        }
        _ = fi37
    }
    var v71 uint16 = 747
    v72 := GenS5{F0: "\n\t"}
    v71 = (^(v71))
    var v73 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v71
    _ = v72
    _ = v73
}
