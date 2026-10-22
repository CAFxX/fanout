package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 int16; F1 string; F2 int64 }

func main() {
    v1 := []uint32{1}
    var v2 uint8 = 1
    v3 := make(map[int]int)
    var v4 uint64 = 2
    v5 := []int{-9223372036854775808, -169, -2147483648, 42}
    var v6 uint32 = 4294967295
    switch "foo bar" {
    case "hello":
        v2 += v2
    case "a":
        v6 = 2
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        if len(v5) > 0 {
            v5[0] = -672
        }
        fmt.Println(v4, v6, v2)
        var v8 uint16 = 1
        _ = fi7
        _ = v8
    }
    fmt.Println(v6, v4)
    switch -121 {
    case -1:
        v9 := GenS1{F0: 833}
        _ = v9
    case 32767:
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 string = "foo bar"
            for fi12 := 0; fi12 < 12; fi12++ {
                v3[42] = fi12
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
    case 290:
        v6 = v6
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        if true {
            v2 += 1
            v14 := GenS2{F0: (v2 & 42)}
            _ = v14
        } else {
            v6 = v6
            if (1 < 262) {
                for fi15 := 0; fi15 < 19; fi15++ {
                    fmt.Println(v4, v4)
                    fmt.Println(fi15)
                    _ = fi15
                }
                v2 = (v2 - v2)
                v6 -= (^(v6))
            } else {
                v4 = (v4 % ((v4) | 1))
                v4 = (v4 + v4)
            }
        }
        _ = fi13
    }
    var v16 uint8 = uint8(v6)
    var v17 int8 = 87
    v17 = 127
    if ((^(v6)) != v6) {
        switch "foo bar" {
        case "hello":
            var v18 uint64 = v4
            _ = v18
        default:
            v2 = 0
        }
        if true {
            v19 := GenS2{F0: v2}
            fmt.Println(v2, v4)
            if false {
                var v20 bool = false
                _ = v20
            } else {
                for fi21 := 0; fi21 < 17; fi21++ {
                    for fi22 := 0; fi22 < 19; fi22++ {
                        if (1000 == 1) {
                            if true {
                                var v23 uint8 = (v16 % ((255) | 1))
                                v23 = (v23 | v23)
                                _ = v23
                            } else {
                                var v24 float32 = (1.5 * 0.1)
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    if (2147483647 <= 0) {
                                        v24 = ((((v24 / v24) + (float32((float64(((float32(0.1) / (-2.25 / v24)) / v24)) / 1.5)) - ((1.5 + v24) * v24))) - 0.1) * v24)
                                    } else {
                                        if (false || true) {
                                            v6 -= v6
                                            _ = len(v5)
                                            v5 = append(v5, fi21)
                                            if len(v5) > 0 {
                                                v5[0] = (fi22 & fi22)
                                            }
                                        } else {
                                            var v25 int16 = int16(v4)
                                            if false {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                default:
                                                    switch "\n\t" {
                                                    case "":
                                                        fmt.Println(v4)
                                                        var v26 uint = 18446744073709551615
                                                        v25 = -243
                                                        _ = v26
                                                    case "\n\t":
                                                        switch (fi22 * (fi21 * (fi21 >> ((1) & 63)))) {
                                                        case 4294967295:
                                                            v25 += v25
                                                        case -2147483648:
                                                            if true {
                                                                v16 += (v2 / (((v16 / (((v2 ^ v16)) | 1))) | 1))
                                                                v27 := GenS3{F0: 0, F1: 2}
                                                                var v28 float64 = 3.14159
                                                                fmt.Println(v4)
                                                                if true {
                                                                    _ = len(v1)
                                                                    fmt.Println(fi22, fi21)
                                                                } else {
                                                                    _ = len(v1)
                                                                    v29 := GenS4{F0: v4, F1: (^(v16)), F2: v17}
                                                                    _ = v29
                                                                }
                                                                _ = v27
                                                                _ = v28
                                                            } else {
                                                                v30 := GenS1{F0: 2}
                                                                var v31 int16 = (v25 >> ((2) & 63))
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                        default:
                                                            v24 = v24
                                                        }
                                                    case "a":
                                                        var v32 float64 = (1.5 / 1.0)
                                                        _ = v32
                                                    }
                                                }
                                            } else {
                                                var v33 string = "foo bar"
                                                v25 = 0
                                                if true {
                                                    fmt.Println(fi22, v2)
                                                    switch 32768 {
                                                    case -1:
                                                        var v34 int = (fi22 | (^(fi22)))
                                                        _ = v34
                                                    }
                                                } else {
                                                    fmt.Println(v4, v17, fi22)
                                                    if false {
                                                        if (false && true) {
                                                            switch -188 {
                                                            case -718:
                                                                v25 += (v25 * (v25 - v25))
                                                            default:
                                                                v3[(fi21 << ((255) & 63))] = 150
                                                            }
                                                        }
                                                        v2 -= v2
                                                        var v35 int8 = (-(v17))
                                                        _ = v35
                                                    }
                                                    if ((v4 >= (v4 % ((0) | 1))) || (18446744073709551615 <= 4294967295)) {
                                                        v36 := GenS2{F0: (v2 * (^(v16)))}
                                                        _ = v36
                                                    } else {
                                                        v4 += v4
                                                        v37 := GenS2{F0: v2}
                                                        switch (-(fi21)) {
                                                        case 42:
                                                            if true {
                                                                v38 := GenS2{F0: v2}
                                                                fmt.Println(v2)
                                                                _ = v38
                                                            } else {
                                                                if false {
                                                                    if (v4 < (v4 / ((v4) | 1))) {
                                                                        switch fi22 {
                                                                        case -2147483648:
                                                                            v16 = (v2 + 176)
                                                                        case 728:
                                                                            switch (((("hello" + v33) + v33) + ((v33 + ((v33 + v33) + ((("" + "foo bar") + v33) + (v33 + v33)))) + ("\n\t" + "\n\t"))) + (v33 + (v33 + (v33 + "a")))) {
                                                                            case "a":
                                                                                _ = len(v3)
                                                                            case "":
                                                                                if true {
                                                                                    v4 -= 18446744073709551615
                                                                                    if true {
                                                                                        var v39 int64 = 9223372036854775807
                                                                                        _ = v39
                                                                                    }
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                } else {
                                                                                    v40 := GenS4{F0: (v4 >> ((1) & 63)), F1: 97, F2: (^(v17))}
                                                                                    _ = v40
                                                                                }
                                                                            }
                                                                        case 4294967295:
                                                                            fmt.Println(fi21)
                                                                        }
                                                                        var v41 int32 = 0
                                                                        _ = v41
                                                                    } else {
                                                                        v24 -= -2.25
                                                                        v42 := GenS5{F0: true}
                                                                        _ = v42
                                                                    }
                                                                    delete(v3, 9223372036854775807)
                                                                    fmt.Println(v33, v25, v2)
                                                                } else {
                                                                    var v43 int8 = int8(v4)
                                                                    if len(v5) > 0 {
                                                                        v5[0] = fi21
                                                                    }
                                                                    _ = v43
                                                                }
                                                                switch (fi22 ^ fi22) {
                                                                case -291:
                                                                    var v44 int8 = 42
                                                                    _ = v44
                                                                case -406:
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (v6 / (((v6 / (((v6 % ((4294967295) | 1))) | 1))) | 1))
                                                                    }
                                                                case 2147483647:
                                                                    var v45 bool = true
                                                                    _ = v45
                                                                default:
                                                                    v16 += v2
                                                                }
                                                                if len(v5) > 0 {
                                                                    v5[0] = (fi21 % (((fi22 & fi21)) | 1))
                                                                }
                                                            }
                                                        case -2147483648:
                                                            v16 -= (v16 | v2)
                                                        default:
                                                            v17 += (v17 >> ((675) & 63))
                                                        }
                                                        _ = v37
                                                    }
                                                    v16 = 2
                                                }
                                                _ = v33
                                            }
                                            v5 = append(v5, (fi21 / ((fi21) | 1)))
                                            v6 = (v6 - v6)
                                            _ = v25
                                        }
                                    }
                                }
                                v17 += 127
                                switch "\n\t" {
                                case "foo bar":
                                    var v46 float64 = (-2.25 - 0.1)
                                    _ = v46
                                case "":
                                    if false {
                                        var v47 uint32 = v6
                                        _ = v47
                                    } else {
                                        var v48 float64 = 0.1
                                        switch "hello" {
                                        case "foo bar":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        _ = v48
                                    }
                                case "hello":
                                    var v49 int64 = -865
                                    _ = v49
                                }
                                v16 += 1
                                _ = v24
                            }
                            switch "a" {
                            case "":
                                var v50 int8 = (v17 ^ (v17 / ((-5) | 1)))
                                _ = v50
                            }
                            if true {
                                var v51 int16 = 1000
                                _ = v51
                            }
                        } else {
                            var v52 float64 = 3.14159
                            v6 = (v6 + 2)
                            v2 -= 127
                            v53 := GenS3{F0: -1, F1: 127}
                            switch fi22 {
                            case 2147483647:
                                var v54 uint32 = (^(v6))
                                _ = v54
                            case 4294967295:
                                if len(v5) > 0 {
                                    v5[0] = fi21
                                }
                            case -2147483648:
                                var v55 int8 = v17
                                _ = v55
                            }
                            _ = v52
                            _ = v53
                        }
                        v56 := GenS1{F0: 65535}
                        v57 := GenS5{F0: true}
                        _ = fi22
                        _ = v56
                        _ = v57
                    }
                    v16 = 0
                    v6 = v6
                    _ = fi21
                }
                var v58 uint64 = (v4 >> ((369) & 63))
                v59 := GenS6{F0: 288, F1: -1}
                _ = v58
                _ = v59
            }
            _ = v19
        }
        var v60 int64 = 1
        _ = v60
    }
    v61 := GenS3{F0: int8(v6), F1: uint8(v4)}
    var v62 string = ""
    for fi63 := 0; fi63 < 1; fi63++ {
        v64 := GenS7{F0: -628, F1: (v62 + v62), F2: 358}
        _ = fi63
        _ = v64
    }
    var v65 uint16 = 32767
    var v66 float32 = (1.5 * 1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v65
    _ = v66
}
