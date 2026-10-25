package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint64; F2 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS9 struct { F0 int16; F1 int64; F2 uint }
type GenS10 struct { F0 int64; F1 bool; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 uint64 = 1
    var v4 bool = false
    var v5 int32 = 431
    var v6 string = ("hello" + ("" + ("" + "\n\t")))
    var v7 uint64 = v3
    if (18446744073709551615 > 255) {
        if false {
            var v8 int32 = v5
            switch -86 {
            case 90:
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := []int{-553, 32767}
                    _ = fi9
                    _ = v10
                }
            }
            _ = v8
        }
        var v11 uint16 = 32767
        _ = v11
    }
    if v4 {
        v12 := v1[127]
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 string = v12
            _ = fi13
            _ = v14
        }
        _ = v12
    } else {
        v15 := GenS1{F0: 814, F1: 4294967295}
        v16 := []int{-2147483648}
        switch (v6 + "") {
        case "\n\t":
            var v17 int8 = 0
            _ = v17
        case "a":
            v18 := GenS2{F0: 42}
            _ = v18
        case "":
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                if (v7 <= (v3 ^ (^(v7)))) {
                    for fi19 := 0; fi19 < 5; fi19++ {
                        var v20 uint = 255
                        _ = fi19
                        _ = v20
                    }
                }
            case "hello":
                _ = len(v16)
            case "\n\t":
                var v21 bool = v4
                _ = v21
            }
        default:
            v22 := GenS3{F0: 182}
            _ = v22
        }
        _ = v15
        _ = v16
    }
    v23 := GenS2{F0: 2}
    fmt.Println(v7, v4, v7)
    var v24 int64 = 898
    if (-164 < 1000) {
        if (-64 >= -32768) {
            var v25 uint8 = 1
            for fi26 := 0; fi26 < 11; fi26++ {
                for fi27 := 0; fi27 < 15; fi27++ {
                    v28 := GenS4{F0: 42, F1: ((v5 << ((1) & 63)) <= v5)}
                    if (float64(3.14159) < 1e10) {
                        if v4 {
                            var v29 float64 = (-2.25 * 0.1)
                            if (!(-19 <= -1)) {
                                v30 := []int{(fi26 - fi26)}
                                fmt.Println(v3)
                                v6 = v6
                                v24 -= v24
                                var v31 float32 = float32(v29)
                                _ = v30
                                _ = v31
                            }
                            _ = v29
                        }
                        fmt.Println(v4, fi26)
                    } else {
                        v7 = 0
                        v32 := GenS1{F0: fi26, F1: v3}
                        v5 = v5
                        v33 := GenS2{F0: 882}
                        var v34 float64 = ((0.0 * -2.25) * (-2.25 / 1.5))
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
            if ((v25 | v25) > v25) {
                v35 := []int{-17}
                v24 += v24
                _ = v35
            } else {
                if (-9223372036854775808 <= -487) {
                    if (264 <= 0) {
                        v36 := GenS5{F0: 32768}
                        v37 := GenS3{F0: (v25 & 7)}
                        _ = v36
                        _ = v37
                    } else {
                        if (472 <= 0) {
                            for fi38 := 0; fi38 < 3; fi38++ {
                                if (true || ((v3 + 888) <= 18446744073709551615)) {
                                    v39 := GenS6{F0: ("foo bar" + ("foo bar" + ((v6 + "") + ("" + "")))), F1: 18446744073709551615, F2: 1}
                                    for fi40 := 0; fi40 < 9; fi40++ {
                                        v41 := GenS7{F0: -1}
                                        v42 := GenS8{F0: 2147483648, F1: (^(v25)), F2: 255}
                                        v5 = v5
                                        _ = fi40
                                        _ = v41
                                        _ = v42
                                    }
                                    if (((float64(float32(3.14159)) - -2.25) - 0.1) >= 0.0) {
                                        v4 = true
                                    }
                                    _ = v39
                                }
                                fmt.Println(v25, v3)
                                _ = fi38
                            }
                            var v43 uint16 = 32767
                            var v44 int32 = -380
                            _ = v43
                            _ = v44
                        }
                    }
                } else {
                    v45 := GenS3{F0: v25}
                    if false {
                        if v4 {
                            var v46 float32 = (1.5 + float32(1e10))
                            fmt.Println(v6, v6, v6)
                            var v47 int64 = int64(v3)
                            _ = v46
                            _ = v47
                        } else {
                            fmt.Println(v4, v7)
                            var v48 int = 1
                            _ = v48
                        }
                        if (!(!(true || ((330 < 815) || (!v4))))) {
                            if v4 {
                                v25 = 42
                                var v49 bool = (4294967295 >= 1)
                                _ = v49
                            } else {
                                fmt.Println(v7, v4, v4)
                                switch (v6 + v6) {
                                case "\n\t":
                                    v5 -= (-(v5))
                                case "a":
                                    if (v4 || v4) {
                                        fmt.Println(v6)
                                        delete(v2, 722)
                                    } else {
                                        delete(v2, 1)
                                    }
                                case "foo bar":
                                    if (255 > 255) {
                                        v50 := GenS5{F0: -856}
                                        for fi51 := 0; fi51 < 1; fi51++ {
                                            for fi52 := 0; fi52 < 4; fi52++ {
                                                if v4 {
                                                    switch (v6 + "\n\t") {
                                                    case "hello":
                                                        v5 -= v5
                                                    case "a":
                                                        fmt.Println(v24)
                                                    case "\n\t":
                                                        var v53 int16 = -32768
                                                        _ = v53
                                                    }
                                                    v6 = ((v6 + v6) + v6)
                                                } else {
                                                    var v54 uint16 = 443
                                                    var v55 uint8 = v25
                                                    var v56 float64 = -2.25
                                                    v57 := GenS2{F0: 2}
                                                    switch (fi51 >> ((255) & 63)) {
                                                    case 1:
                                                        var v58 int8 = 1
                                                        fmt.Println(v25, v24)
                                                        fmt.Println(v7, v4)
                                                        _ = v58
                                                    case 127:
                                                        v59 := GenS8{F0: 620, F1: v55, F2: 732}
                                                        _ = v59
                                                    default:
                                                        v60 := GenS5{F0: (fi51 - (fi52 * 42))}
                                                        _ = v60
                                                    }
                                                    v61 := v1[(fi52 + 1)]
                                                    _ = v54
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                    _ = v61
                                                }
                                                v5 = (^(v5))
                                                _ = fi52
                                            }
                                            _ = fi51
                                        }
                                        if v4 {
                                            if true {
                                                switch 127 {
                                                case 2147483647:
                                                    var v62 int64 = (v24 / ((v24) | 1))
                                                    _ = v62
                                                case -715:
                                                    v24 = (v24 & -9223372036854775808)
                                                }
                                            } else {
                                                v63 := GenS9{F0: -1, F1: (v24 / ((1) | 1)), F2: 4294967295}
                                                var v64 bool = ((v4 && v4) || v4)
                                                fmt.Println(v24)
                                                _ = v63
                                                _ = v64
                                            }
                                            v65 := GenS4{F0: 1000, F1: v4}
                                            _ = v65
                                        }
                                        _ = v50
                                    } else {
                                        for fi66 := 0; fi66 < 8; fi66++ {
                                            v5 -= 42
                                            _ = fi66
                                        }
                                    }
                                }
                            }
                            for fi67 := 0; fi67 < 3; fi67++ {
                                fmt.Println(v6, v7, fi67)
                                v5 -= -1
                                _ = fi67
                            }
                        }
                    } else {
                        v68 := []int{4294967295}
                        v69 := GenS10{F0: v24, F1: true, F2: (^(v25))}
                        v24 = (v24 & v24)
                        _ = v68
                        _ = v69
                    }
                    v1[666] = (("a" + "hello") + ("" + (("\n\t" + v6) + ((v6 + ("\n\t" + "a")) + v6))))
                    _ = v45
                }
                var v70 uint8 = 42
                _ = v70
            }
            _ = v25
        } else {
            v71 := GenS5{F0: 2147483647}
            if v4 {
                var v72 uint8 = 255
                if v4 {
                    if false {
                        v73 := GenS6{F0: ((v6 + (v6 + (v6 + ((((("a" + "a") + (v6 + ("foo bar" + ""))) + (("\n\t" + v6) + v6)) + v6) + "a")))) + ((v6 + v6) + v6)), F1: 357, F2: 42}
                        _ = v73
                    }
                } else {
                    v4 = (0 > 100)
                }
                var v74 uint16 = 201
                _ = v72
                _ = v74
            } else {
                switch v6 {
                case "a":
                    fmt.Println(v3, v3, v6)
                case "\n\t":
                    v5 = v5
                case "hello":
                    v75 := GenS1{F0: 981, F1: v7}
                    _ = v75
                default:
                    for fi76 := 0; fi76 < 15; fi76++ {
                        fmt.Println(v6, fi76, fi76)
                        fmt.Println(v6)
                        var v77 float64 = (float64(0.1) - (3.14159 - 1.5))
                        _ = fi76
                        _ = v77
                    }
                }
                var v78 string = v6
                _ = v78
            }
            _ = v71
        }
    } else {
        v1[-11] = v6
        v79 := GenS5{F0: -299}
        _ = v79
    }
    v1[32768] = ("" + v6)
    fmt.Println(v7, v5, v3)
    v6 = v6
    fmt.Println(v7, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
