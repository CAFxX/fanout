package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int; F1 int32; F2 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int32; F1 int32 }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 int32 }
func genFn0(p0 int, p1 int32, p2 int16) string {
    v7 := GenS1{F0: 705, F1: (p0 >> ((309) & 63))}
    if false {
        var v8 uint = 700
        _ = v8
    } else {
        var v9 int16 = p2
        var v10 bool = false
        _ = v10
        _ = v9
    }
    p1 = (p1 / ((p1) | 1))
    p2 = (p2 << ((2) & 63))
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return ("a" + "\n\t")
}

func main() {
    var v1 bool = true
    v2 := []uint32{658, 4294967295, 570, 242}
    v3 := []int{42, -836, 4294967295, 127}
    v4 := []int{-155, 0}
    var v5 bool = false
    var v6 float64 = 0.1
    var v12 uint16 = 255
    switch 32767 {
    case 1:
        v1 = true
    case -1:
        v1 = true
    default:
        if (144 == 42) {
            var v13 uint64 = 645
            if (-1 < -39) {
                v12 = (v12 % ((255) | 1))
            }
            _ = v13
        } else {
            for fi14 := 0; fi14 < 17; fi14++ {
                for fi15 := 0; fi15 < 6; fi15++ {
                    v16 := GenS1{F0: 871, F1: 941}
                    v12 -= v12
                    v1 = v5
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 19; fi17++ {
                    var v18 int16 = 42
                    _ = fi17
                    _ = v18
                }
                switch "foo bar" {
                case "a":
                    if v5 {
                        if (((((v12 / (((v12 / ((32767) | 1))) | 1)) <= (^(v12))) && true) || v1) || true) {
                            switch "a" {
                            case "\n\t":
                                switch "" {
                                case "":
                                    for fi19 := 0; fi19 < 11; fi19++ {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = len(v3)
                                        if (v6 >= (v6 - -2.25)) {
                                            v20 := GenS2{F0: 984}
                                            fmt.Println(fi19, v6)
                                            v1 = (!(false || (!v1)))
                                            _ = v20
                                        }
                                        _ = fi19
                                    }
                                case "hello":
                                    for fi21 := 0; fi21 < 18; fi21++ {
                                        v12 += (v12 | 32767)
                                        fmt.Println(v6, fi21)
                                        v22 := GenS2{F0: 2}
                                        _ = fi21
                                        _ = v22
                                    }
                                case "a":
                                    v23 := GenS2{F0: 2}
                                    _ = v23
                                }
                            case "hello":
                                if (-102 >= int8(fi14)) {
                                    v6 = (((3.14159 - (float64(3.14159) / (float64(-2.25) + v6))) * v6) * float64((float32(1e10) + (3.14159 * -2.25))))
                                    var v24 int = fi14
                                    v25 := GenS2{F0: 1}
                                    _ = v24
                                    _ = v25
                                } else {
                                    v26 := GenS3{F0: 9223372036854775807, F1: 213}
                                    _ = v26
                                }
                            case "":
                                v5 = (-2147483648 < 2147483647)
                            }
                            var v27 int = (fi14 | -191)
                            _ = v27
                        }
                    }
                default:
                    v1 = (2 < 2)
                }
                _ = fi14
            }
            fmt.Println(v1, v5)
        }
    }
    v28 := GenS4{F0: -884, F1: -1, F2: -79}
    var v29 string = "a"
    v5 = false
    for fi30 := 0; fi30 < 17; fi30++ {
        fmt.Println(v29)
        _ = fi30
    }
    v1 = (!(!false))
    switch ("" + v29) {
    case "\n\t":
        _ = len(v2)
    case "hello":
        if true {
            v29 = v29
        } else {
            if false {
                var v31 int16 = -1
                if (!v1) {
                    for fi32 := 0; fi32 < 18; fi32++ {
                        for fi33 := 0; fi33 < 13; fi33++ {
                            v34 := GenS4{F0: fi32, F1: 0, F2: v31}
                            _ = fi33
                            _ = v34
                        }
                        _ = fi32
                    }
                    if v5 {
                        v4 = append(v4, 32768)
                    } else {
                        if (0 < 0) {
                            if true {
                                v5 = (1.5 < (float32(v6) - 1.5))
                                v35 := GenS2{F0: 533}
                                _ = v35
                            }
                        } else {
                            v36 := GenS2{F0: 0}
                            for fi37 := 0; fi37 < 15; fi37++ {
                                for fi38 := 0; fi38 < 2; fi38++ {
                                    if ((v29 + ((v29 + "hello") + ("foo bar" + ("\n\t" + v29)))) == "hello") {
                                        fmt.Println(v31)
                                        v29 = "foo bar"
                                        v3 = append(v3, fi37)
                                    } else {
                                        v12 = (v12 & v12)
                                        v39 := GenS1{F0: 255, F1: fi37}
                                        var v40 int8 = -24
                                        _ = v39
                                        _ = v40
                                    }
                                    var v41 string = (("\n\t" + "") + v29)
                                    v12 = (v12 & 0)
                                    _ = fi38
                                    _ = v41
                                }
                                for fi42 := 0; fi42 < 4; fi42++ {
                                    var v43 int32 = 42
                                    if len(v3) > 0 {
                                        v3[0] = fi37
                                    }
                                    v29 = (("hello" + v29) + v29)
                                    _ = fi42
                                    _ = v43
                                }
                                v44 := GenS4{F0: (fi37 << ((0) & 63)), F1: 1, F2: 1000}
                                _ = fi37
                                _ = v44
                            }
                            var v45 uint = 675
                            _ = v36
                            _ = v45
                        }
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v29, v29)
                        case "hello":
                            v6 += float64(1e10)
                        }
                        v1 = (uint(v12) == 42)
                    }
                    if false {
                        if true {
                            v29 = v29
                            if v1 {
                                for fi46 := 0; fi46 < 7; fi46++ {
                                    fmt.Println(v5)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if true {
                                        v1 = ((v29 <= v29) && true)
                                        for fi47 := 0; fi47 < 10; fi47++ {
                                            fmt.Println(v31)
                                            var v48 bool = v1
                                            v49 := GenS5{F0: (fi46 / (((fi46 | fi47)) | 1))}
                                            _ = fi47
                                            _ = v48
                                            _ = v49
                                        }
                                    }
                                    _ = fi46
                                }
                                v50 := GenS6{F0: 1, F1: 257, F2: v12}
                                _ = v50
                            }
                            v51 := GenS3{F0: -732, F1: 65535}
                            _ = v51
                        } else {
                            v29 = (v29 + ((v29 + v29) + v29))
                        }
                        for fi52 := 0; fi52 < 10; fi52++ {
                            v29 = (("a" + (v29 + v29)) + v29)
                            v53 := GenS5{F0: (fi52 - 9223372036854775807)}
                            switch fi52 {
                            case 0:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 1:
                                if (v12 >= v12) {
                                    switch ("\n\t" + v29) {
                                    case "foo bar":
                                        switch fi52 {
                                        case -1:
                                            if v5 {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                if v1 {
                                                    var v54 int = fi52
                                                    switch v54 {
                                                    case -2147483648:
                                                        v54 = (v54 >> ((4294967295) & 63))
                                                    }
                                                    _ = len(v3)
                                                    _ = v54
                                                }
                                                for fi55 := 0; fi55 < 1; fi55++ {
                                                    var v56 uint16 = v12
                                                    _ = fi55
                                                    _ = v56
                                                }
                                            } else {
                                                v1 = false
                                                var v57 int = (fi52 * fi52)
                                                var v58 uint16 = 42
                                                _ = v57
                                                _ = v58
                                            }
                                        case -294:
                                            var v59 int64 = -11
                                            _ = v59
                                        }
                                    case "":
                                        var v60 int8 = -128
                                        _ = v60
                                    case "a":
                                        for fi61 := 0; fi61 < 20; fi61++ {
                                            if ((!false) && false) {
                                                var v62 int8 = -128
                                                _ = v62
                                            } else {
                                                if v1 {
                                                    var v63 int32 = 42
                                                    _ = v63
                                                } else {
                                                    v6 -= 0.1
                                                    v12 -= 2
                                                    v64 := GenS7{F0: -9223372036854775808, F1: 385}
                                                    _ = v64
                                                }
                                                var v65 uint = 42
                                                v3 = append(v3, 32767)
                                                _ = v65
                                            }
                                            v1 = (-9223372036854775808 != 739)
                                            _ = fi61
                                        }
                                    default:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    if v5 {
                                        v2 = append(v2, 2147483647)
                                        v29 = "\n\t"
                                    }
                                    if (!v1) {
                                        switch v29 {
                                        case "a":
                                            if ((false || (v12 != 1)) || false) {
                                                v12 = v12
                                                v29 = "\n\t"
                                            }
                                        }
                                        var v66 int = 32768
                                        fmt.Println(v31, v12, v31)
                                        _ = v66
                                    } else {
                                        v67 := GenS4{F0: (fi52 / (((fi52 + fi52)) | 1)), F1: int32(v31), F2: 1}
                                        var v68 int8 = 0
                                        var v69 uint16 = v12
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    }
                                } else {
                                    v70 := GenS7{F0: 32768, F1: -1}
                                    var v71 float64 = (v6 * (v6 * (0.0 + (((v6 + v6) - (v6 / v6)) * v6))))
                                    if v1 {
                                        v6 = v6
                                        var v72 uint8 = 39
                                        _ = v72
                                    }
                                    _ = v70
                                    _ = v71
                                }
                            }
                            _ = fi52
                            _ = v53
                        }
                    } else {
                        for fi73 := 0; fi73 < 14; fi73++ {
                            var v74 int8 = 42
                            if v1 {
                                var v75 uint = 1
                                var v76 uint = v75
                                var v77 uint16 = 149
                                _ = v75
                                _ = v76
                                _ = v77
                            } else {
                                for fi78 := 0; fi78 < 13; fi78++ {
                                    switch ("" + v29) {
                                    case "\n\t":
                                        _ = len(v2)
                                    case "":
                                        var v79 int16 = v31
                                        _ = v79
                                    case "hello":
                                        fmt.Println(v74, fi73, v31)
                                    }
                                    if (!v5) {
                                        if len(v3) > 0 {
                                            v3[0] = fi73
                                        }
                                        v80 := GenS8{F0: 584, F1: 42}
                                        _ = v80
                                    }
                                    _ = fi78
                                }
                                switch v29 {
                                case "":
                                    if len(v4) > 0 {
                                        v4[0] = 9223372036854775807
                                    }
                                case "foo bar":
                                    var v81 int16 = (^(v31))
                                    _ = v81
                                }
                            }
                            _ = fi73
                            _ = v74
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                } else {
                    var v82 int64 = -9223372036854775808
                    if (-295 == -1) {
                        v29 = ("" + v29)
                        switch v29 {
                        case "\n\t":
                            v31 = v31
                        case "foo bar":
                            for fi83 := 0; fi83 < 10; fi83++ {
                                var v84 int32 = -1
                                for fi85 := 0; fi85 < 15; fi85++ {
                                    v86 := GenS5{F0: -575}
                                    var v87 int8 = 100
                                    _ = fi85
                                    _ = v86
                                    _ = v87
                                }
                                _ = fi83
                                _ = v84
                            }
                        case "a":
                            for fi88 := 0; fi88 < 20; fi88++ {
                                v89 := GenS7{F0: fi88, F1: int64(v12)}
                                var v90 int64 = v82
                                _ = fi88
                                _ = v89
                                _ = v90
                            }
                        default:
                            v29 = ((v29 + ((((((v29 + ("a" + "")) + v29) + ("\n\t" + v29)) + "a") + v29) + (v29 + "hello"))) + "")
                        }
                        fmt.Println(v1)
                    }
                    _ = v82
                }
                var v91 uint = 0
                _ = v31
                _ = v91
            }
            v92 := GenS9{F0: 9223372036854775807, F1: 42}
            var v93 uint16 = (v12 + v12)
            _ = v92
            _ = v93
        }
    default:
        v94 := GenS10{F0: 332}
        _ = v94
    }
    var v95 string = "hello"
    v29 = v29
    v5 = false
    var v96 int32 = -2147483648
    v96 = 2147483647
    _ = len(v4)
    v4 = append(v4, -105)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v95
    _ = v96
}
