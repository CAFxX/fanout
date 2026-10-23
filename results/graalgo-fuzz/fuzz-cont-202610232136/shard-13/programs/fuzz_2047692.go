package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 int; F2 int16 }
type GenS8 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS9 struct { F0 int16; F1 uint8 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]int)
    var v3 uint16 = 1
    var v4 uint64 = 87
    v5 := make(map[int]int)
    var v6 uint16 = (v3 / ((2) | 1))
    if (!(939 >= 406)) {
        if true {
            v7 := []int{32767, 42}
            v6 = v3
            var v8 float32 = -2.25
            _ = v7
            _ = v8
        }
    }
    v4 += v4
    v9 := GenS1{F0: 246, F1: 0}
    v6 -= v3
    fmt.Println(v4, v4)
    for fi10 := 0; fi10 < 1; fi10++ {
        switch fi10 {
        case -299:
            v4 += v4
        case 977:
            v6 = v3
        }
        v11 := GenS1{F0: (fi10 + int(v1)), F1: v6}
        v3 = 533
        _ = fi10
        _ = v11
    }
    v12 := []int{128, 4294967295}
    v6 -= 1
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS2{F0: (float64(-2.25) <= float64(0.1)), F1: v6}
        fmt.Println(v3)
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS2{F0: false, F1: 2}
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v4 = 18446744073709551615
        v12 = append(v12, int(v3))
        _ = fi17
    }
    if false {
        switch 739 {
        case -135:
            var v18 int64 = -1
            _ = v18
        case -304:
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 int16 = -874
                delete(v2, "a")
                v3 = 42
                _ = fi19
                _ = v20
            }
        case 36:
            if (!(false && false)) {
                if (1.5 > 3.14159) {
                    if (((1.5 / (0.0 - ((0.1 * (((0.0 - 0.1) + 1.5) + float64(1.5))) - (0.1 + 1.5)))) > ((0.1 / 0.1) + 1.5)) && (4294967295 <= 1)) {
                        v21 := GenS2{F0: false, F1: v3}
                        if (4294967295 >= 65535) {
                            for fi22 := 0; fi22 < 10; fi22++ {
                                v6 = v3
                                for fi23 := 0; fi23 < 20; fi23++ {
                                    if (32767 >= -864) {
                                        var v24 int8 = 1
                                        v1 = (v1 % (((v1 / ((v1) | 1))) | 1))
                                        _ = v24
                                    } else {
                                        fmt.Println(fi23, fi22)
                                    }
                                    if true {
                                        var v25 int16 = 428
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        _ = v25
                                    } else {
                                        var v26 int32 = 0
                                        var v27 int32 = v26
                                        v5[fi22] = fi23
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = fi23
                                }
                                for fi28 := 0; fi28 < 16; fi28++ {
                                    v3 -= v3
                                    if (0.1 == 1e10) {
                                        v6 = v3
                                    }
                                    v1 -= (v1 & v1)
                                    var v29 uint32 = 2147483648
                                    _ = fi28
                                    _ = v29
                                }
                                _ = fi22
                            }
                        } else {
                            delete(v2, "\n\t")
                        }
                        _ = v21
                    }
                }
                v4 = 578
            }
        }
        v30 := GenS2{F0: true, F1: 220}
        _ = v30
    }
    switch ("\n\t" + "") {
    case "a":
        var v31 uint = 1
        _ = v31
    case "hello":
        _ = len(v2)
    case "foo bar":
        var v32 uint16 = (v6 ^ 0)
        _ = v32
    default:
        v3 = 2
    }
    v1 += v1
    v33 := GenS3{F0: "\n\t"}
    if false {
        var v34 float32 = 1.5
        if (!false) {
            for fi35 := 0; fi35 < 2; fi35++ {
                if false {
                    var v36 int = fi35
                    _ = v36
                } else {
                    if false {
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = len(v12)
                        v37 := GenS2{F0: false, F1: (v6 * 42)}
                        _ = v37
                    } else {
                        _ = len(v12)
                        v1 += v1
                    }
                }
                for fi38 := 0; fi38 < 15; fi38++ {
                    v1 += 42
                    v39 := GenS4{F0: -1, F1: -471}
                    if false {
                        var v40 float32 = float32((float64(v34) * 0.1))
                        switch "hello" {
                        case "foo bar":
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(fi35, fi38, fi38)
                            case "":
                                _ = len(v12)
                                _ = len(v12)
                                v41 := GenS4{F0: 42, F1: 220}
                                _ = v41
                            }
                        }
                        v4 -= (v4 << ((255) & 63))
                        _ = v40
                    } else {
                        var v42 uint64 = v4
                        v5[fi35] = (fi35 ^ 4294967295)
                        _ = v42
                    }
                    _ = fi38
                    _ = v39
                }
                var v43 float64 = -2.25
                _ = fi35
                _ = v43
            }
        } else {
            if false {
                var v44 int16 = 32767
                _ = v44
            }
        }
        var v45 uint8 = 42
        _ = v34
        _ = v45
    }
    if false {
        v3 = v6
    } else {
        fmt.Println(v4)
        var v46 uint = 1
        _ = v46
    }
    if ((v1 < 1) && false) {
        v47 := GenS5{F0: (v1 & (v1 ^ (v1 - (v1 & 42))))}
        v2["a"] = 1
        _ = v47
    } else {
        for fi48 := 0; fi48 < 1; fi48++ {
            v49 := GenS1{F0: -1, F1: (v6 | (v3 << ((206) & 63)))}
            if len(v12) > 0 {
                v12[0] = (fi48 + 0)
            }
            var v50 uint32 = 2147483647
            _ = fi48
            _ = v49
            _ = v50
        }
        switch ("a" + "foo bar") {
        case "a":
            for fi51 := 0; fi51 < 12; fi51++ {
                if (!false) {
                    v4 -= (v4 + (v4 | (v4 * 4294967295)))
                }
                _ = fi51
            }
        case "hello":
            if len(v12) > 0 {
                v12[0] = 591
            }
        }
        v1 += (v1 ^ (v1 & -9223372036854775808))
    }
    for fi52 := 0; fi52 < 4; fi52++ {
        if (!(4294967295 < 4294967295)) {
            if false {
                for fi53 := 0; fi53 < 18; fi53++ {
                    if len(v12) > 0 {
                        v12[0] = fi53
                    }
                    if (-32768 == -64) {
                        var v54 string = ""
                        v55 := GenS1{F0: (fi52 % ((fi53) | 1)), F1: v6}
                        _ = v54
                        _ = v55
                    } else {
                        v5[9223372036854775807] = fi52
                        if false {
                            var v56 uint32 = 65535
                            var v57 int16 = 32767
                            _ = v56
                            _ = v57
                        } else {
                            v3 = (v6 << ((617) & 63))
                            var v58 uint8 = 88
                            _ = v58
                        }
                        v12 = append(v12, fi52)
                    }
                    var v59 uint8 = 1
                    _ = fi53
                    _ = v59
                }
                var v60 uint = 409
                _ = v60
            }
        }
        for fi61 := 0; fi61 < 12; fi61++ {
            var v62 uint8 = 2
            v6 = (v3 / ((v6) | 1))
            var v63 uint16 = 65535
            _ = fi61
            _ = v62
            _ = v63
        }
        if true {
            for fi64 := 0; fi64 < 14; fi64++ {
                if true {
                    var v65 int16 = 1
                    if false {
                        v65 = v65
                        v1 -= v1
                        v65 -= v65
                        if (-15 < 127) {
                            switch fi64 {
                            case -9223372036854775808:
                                if ((v4 << ((588) & 63)) >= 18446744073709551615) {
                                    v3 -= v3
                                    var v66 uint = 221
                                    _ = v66
                                }
                            case 32767:
                                var v67 uint64 = v4
                                fmt.Println(fi64, v3, v67)
                                _ = v67
                            }
                            var v68 uint16 = v3
                            v68 = (v6 * 0)
                            if true {
                                v6 -= (^(v6))
                                var v69 uint = 42
                                var v70 uint = 255
                                var v71 bool = false
                                var v72 uint16 = 2
                                switch "hello" {
                                case "":
                                    v73 := GenS6{F0: -1}
                                    _ = v73
                                case "\n\t":
                                    if false {
                                        if ((-1 >= 0) || v71) {
                                            var v74 uint8 = 2
                                            var v75 uint8 = (^(v74))
                                            var v76 int = fi64
                                            v76 += 344
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                        } else {
                                            if (1 < 200) {
                                                var v77 float64 = 0.1
                                                v3 -= uint16(v69)
                                                v4 -= (v4 - (v4 | (^(v4))))
                                                _ = v77
                                            }
                                            if (((v4 | (v4 + (v4 >> (((v69 * v69)) & 63)))) < v4) || v71) {
                                                if v71 {
                                                    fmt.Println(v72, v72, v4)
                                                } else {
                                                    v78 := GenS7{F0: false, F1: 128, F2: (v65 * (v65 * v65))}
                                                    _ = v78
                                                }
                                            } else {
                                                v79 := GenS4{F0: 0, F1: -212}
                                                v6 -= (v68 % ((544) | 1))
                                                _ = v79
                                            }
                                            v71 = ((v70 == (v70 >> (((v70 << (((^(v70))) & 63))) & 63))) || v71)
                                        }
                                        if ((-(v65)) > 1) {
                                            if v71 {
                                                v4 = v4
                                                fmt.Println(fi64, v68, v65)
                                                if (0.0 == 0.1) {
                                                    if v71 {
                                                        v80 := v5[128]
                                                        _ = v80
                                                    }
                                                    v65 = v65
                                                    v12 = append(v12, fi64)
                                                } else {
                                                    switch (fi52 % ((fi52) | 1)) {
                                                    case 2147483647:
                                                        if (1000 < v65) {
                                                            v81 := GenS8{F0: 42, F1: 77, F2: 127}
                                                            v65 -= (v65 << ((639) & 63))
                                                            _ = len(v12)
                                                            _ = v81
                                                        } else {
                                                            v6 = v6
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v71 = (!((0.0 / 1.5) == (1e10 / 0.1)))
                                                            case "hello":
                                                                var v82 uint = 123
                                                                _ = v82
                                                            case "foo bar":
                                                                v69 -= v70
                                                            }
                                                            v4 += uint64(v70)
                                                            fmt.Println(v71)
                                                            var v83 uint32 = uint32(v72)
                                                            _ = v83
                                                        }
                                                    default:
                                                        fmt.Println(v4, v4, v72)
                                                    }
                                                }
                                            }
                                        }
                                        if true {
                                            v3 -= (v68 * (v68 - (^(v6))))
                                            v71 = (v4 != 1)
                                        } else {
                                            if len(v12) > 0 {
                                                v12[0] = fi64
                                            }
                                            v84 := GenS1{F0: 2147483647, F1: v3}
                                            var v85 uint64 = (v4 >> (((v70 | v70)) & 63))
                                            _ = v84
                                            _ = v85
                                        }
                                    }
                                case "hello":
                                    fmt.Println(v1)
                                default:
                                    v86 := GenS9{F0: 1, F1: 2}
                                    _ = v86
                                }
                                var v87 string = ("a" + ("a" + ("foo bar" + "\n\t")))
                                var v88 uint8 = 42
                                _ = v69
                                _ = v70
                                _ = v71
                                _ = v72
                                _ = v87
                                _ = v88
                            }
                            _ = v68
                        }
                    }
                    fmt.Println(v3, v3)
                    _ = v65
                }
                _ = fi64
            }
            v1 = v1
        } else {
            v5[(fi52 ^ 1)] = (^(fi52))
            v6 -= (v3 % (((v6 / ((v6) | 1))) | 1))
            fmt.Println(fi52, v1, v1)
        }
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
