package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint; F1 int32 }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 string; F1 uint; F2 int32 }
type GenS12 struct { F0 int; F1 uint8; F2 int64 }
type GenS13 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    var v3 int8 = -1
    var v4 uint = 2
    v5 := []int{-1, 1}
    var v6 float32 = 3.14159
    v6 = v6
    var v7 uint8 = uint8(v4)
    v5 = append(v5, 263)
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = -639
    }
    v5 = append(v5, 128)
    v8 := GenS1{F0: v7, F1: (("hello" <= ("foo bar" + ("\n\t" + "foo bar"))) || true)}
    var v9 uint = 65535
    switch 32768 {
    case 4294967295:
        fmt.Println(v6)
    case -478:
        fmt.Println(v6, v4, v6)
    case 1:
        fmt.Println(v3)
    default:
        var v10 int8 = v3
        _ = v10
    }
    var v11 float32 = -2.25
    switch ("" + "\n\t") {
    case "\n\t":
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v7, v4)
        case "":
            v6 -= ((v11 / (v11 * 3.14159)) - v6)
        case "a":
            v6 -= 3.14159
        default:
            switch "\n\t" {
            case "a":
                var v12 float64 = 0.0
                _ = v12
            case "hello":
                var v13 uint8 = (v7 << ((4294967295) & 63))
                _ = v13
            }
        }
    case "foo bar":
        fmt.Println(v6, v6)
    case "":
        var v14 int8 = 33
        _ = v14
    default:
        fmt.Println(v6, v9)
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint = v9
        if (!(1e10 < (float64(1e10) / 1.5))) {
            switch "a" {
            case "a":
                var v17 bool = false
                _ = v17
            case "":
                switch (-(fi15)) {
                case 9223372036854775807:
                    var v18 int8 = (v3 >> ((v9) & 63))
                    _ = v18
                default:
                    _ = len(v2)
                }
            default:
                if true {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v4 -= 462
                        v16 = v9
                        var v20 bool = (v7 < (^(v7)))
                        _ = fi19
                        _ = v20
                    }
                } else {
                    switch -2147483648 {
                    case 9223372036854775807:
                        switch ("hello" + "hello") {
                        case "foo bar":
                            if (522 != -1) {
                                for fi21 := 0; fi21 < 11; fi21++ {
                                    v22 := GenS2{F0: -2147483648, F1: (v4 / ((v4) | 1))}
                                    _ = fi21
                                    _ = v22
                                }
                                if true {
                                    delete(v1, "foo bar")
                                    var v23 float64 = (0.1 / 1e10)
                                    _ = v23
                                }
                            }
                        case "hello":
                            switch 32767 {
                            case 32768:
                                fmt.Println(v16, v11)
                            case -9223372036854775808:
                                var v24 int8 = v3
                                _ = v24
                            default:
                                if false {
                                    var v25 float32 = v6
                                    _ = v25
                                }
                            }
                        }
                    case -2147483648:
                        var v26 int32 = -654
                        _ = v26
                    case -815:
                        switch "\n\t" {
                        case "a":
                            var v27 int32 = -1
                            _ = v27
                        case "\n\t":
                            for fi28 := 0; fi28 < 3; fi28++ {
                                var v29 int = fi15
                                v30 := GenS1{F0: uint8(v16), F1: false}
                                var v31 string = ("foo bar" + "")
                                _ = fi28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        default:
                            var v32 bool = ((v7 * (v7 << ((v16) & 63))) < v7)
                            _ = v32
                        }
                    }
                    v33 := GenS3{F0: v4, F1: 0}
                    _ = v33
                }
            }
        } else {
            var v34 uint = (v9 + (^(v9)))
            _ = v34
        }
        _ = fi15
        _ = v16
    }
    var v35 string = ""
    for fi36 := 0; fi36 < 17; fi36++ {
        v9 = v4
        if (v35 < v35) {
            v9 += (v4 >> ((v9) & 63))
            switch "foo bar" {
            case "foo bar":
                v3 = -128
            }
        }
        if (float32((float64(v6) / 1.5)) != v6) {
            fmt.Println(v35, v9)
            if (1 >= -528) {
                v37 := GenS4{F0: 65535, F1: 946, F2: 32767}
                fmt.Println(v4, v11, v9)
                _ = v37
            } else {
                switch fi36 {
                case 298:
                    for fi38 := 0; fi38 < 11; fi38++ {
                        v39 := GenS3{F0: v9, F1: 0}
                        if false {
                            if false {
                                v4 = v9
                                var v40 string = v35
                                if (((v9 << ((0) & 63)) <= v4) || false) {
                                    if (503 != 4294967295) {
                                        switch int(v3) {
                                        case 128:
                                            switch v40 {
                                            case "\n\t":
                                                var v41 int16 = 1
                                                _ = v41
                                            case "a":
                                                var v42 int32 = 1
                                                _ = v42
                                            default:
                                                if ((^(v3)) != 0) {
                                                    fmt.Println(fi38, v6)
                                                    v43 := GenS5{F0: 255, F1: -559}
                                                    _ = v43
                                                } else {
                                                    var v44 bool = ((v3 * v3) == 34)
                                                    _ = v44
                                                }
                                            }
                                        case 0:
                                            v9 -= 18446744073709551615
                                        case -292:
                                            switch v40 {
                                            case "foo bar":
                                                v7 -= (v7 * 127)
                                            case "hello":
                                                switch fi38 {
                                                case 42:
                                                    if true {
                                                        if true {
                                                            v35 = "foo bar"
                                                        } else {
                                                            if true {
                                                                v11 -= 3.14159
                                                                v40 = v35
                                                                v11 = ((0.1 / ((float32((0.1 / 0.1)) + float32(3.14159)) * 1.5)) - 1e10)
                                                            } else {
                                                                if (true && ((1 < 2147483647) || false)) {
                                                                    var v45 int32 = 812
                                                                    fmt.Println(v6)
                                                                    _ = v45
                                                                } else {
                                                                    v3 -= v3
                                                                    delete(v2, fi36)
                                                                    fmt.Println(v11, v11)
                                                                    var v46 float32 = v11
                                                                    _ = v46
                                                                }
                                                                switch v40 {
                                                                case "a":
                                                                    v47 := GenS3{F0: 1, F1: 629}
                                                                    _ = v47
                                                                case "foo bar":
                                                                    v4 -= (^(v9))
                                                                case "hello":
                                                                    var v48 uint = 1
                                                                    _ = v48
                                                                default:
                                                                    var v49 string = v35
                                                                    _ = v49
                                                                }
                                                            }
                                                            v4 -= (v9 / ((v4) | 1))
                                                            fmt.Println(v40, v4, v11)
                                                        }
                                                    } else {
                                                        if true {
                                                            _ = len(v1)
                                                            _ = len(v2)
                                                            var v50 int32 = -698
                                                            _ = v50
                                                        } else {
                                                            v3 -= (v3 + int8(v9))
                                                            if false {
                                                                v51 := GenS3{F0: 753, F1: 42}
                                                                var v52 int32 = -2147483648
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                        }
                                                    }
                                                case 0:
                                                    v53 := GenS6{F0: 358, F1: 161}
                                                    _ = v53
                                                default:
                                                    fmt.Println(v3, v11)
                                                }
                                            }
                                        default:
                                            var v54 int32 = -2147483648
                                            _ = v54
                                        }
                                        v6 += float32(3.14159)
                                        v6 += 1e10
                                    } else {
                                        v7 = 194
                                        v40 = (("foo bar" + v35) + (((v40 + ((v40 + ((v35 + ("hello" + "hello")) + "\n\t")) + (((v35 + (v40 + v35)) + "foo bar") + "\n\t"))) + v35) + "hello"))
                                        v3 += -77
                                    }
                                    v55 := GenS6{F0: 0, F1: -165}
                                    fmt.Println(v7)
                                    _ = v55
                                }
                                _ = v40
                            }
                        } else {
                            if false {
                                v11 = ((float32(-2.25) * 1e10) + v6)
                            }
                            switch 844 {
                            case 0:
                                if (-1 <= 0) {
                                    if true {
                                        delete(v1, (v35 + (v35 + (v35 + "hello"))))
                                        v5 = append(v5, (fi38 + (fi36 ^ fi36)))
                                    } else {
                                        if true {
                                            v9 += 2
                                            if (v4 >= 65535) {
                                                var v56 string = v35
                                                _ = v56
                                            }
                                        }
                                        switch fi36 {
                                        case 32767:
                                            switch fi36 {
                                            case 9223372036854775807:
                                                switch (v35 + v35) {
                                                case "a":
                                                    v5 = append(v5, 944)
                                                case "":
                                                    v57 := v2[(fi38 / ((207) | 1))]
                                                    _ = v57
                                                }
                                            case 2147483647:
                                                v2[fi38] = true
                                            case -564:
                                                var v58 uint64 = 42
                                                _ = v58
                                            }
                                            if (2 != 764) {
                                                v59 := GenS3{F0: (v4 << ((uint(v7)) & 63)), F1: 819}
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if true {
                                                    v60 := GenS7{F0: (v4 - 42), F1: -2147483648}
                                                    _ = v60
                                                } else {
                                                    if (0 < -32768) {
                                                        v61 := GenS8{F0: 18446744073709551615, F1: int16(v3)}
                                                        v62 := GenS6{F0: 255, F1: -65}
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                }
                                                _ = v59
                                            }
                                            v63 := v1[v35]
                                            _ = v63
                                        case 929:
                                            if (false && false) {
                                                v6 += v6
                                                v35 = ("" + v35)
                                                var v64 uint = 65535
                                                v65 := GenS9{F0: 18446744073709551615}
                                                _ = v64
                                                _ = v65
                                            }
                                        case 411:
                                            v4 = (^(v9))
                                        default:
                                            v9 -= v9
                                        }
                                    }
                                }
                            case 996:
                                var v66 uint = v9
                                _ = v66
                            case 1:
                                if true {
                                    fmt.Println(fi36, fi36)
                                    v4 += v4
                                }
                            }
                        }
                        var v67 int16 = 32767
                        _ = fi38
                        _ = v39
                        _ = v67
                    }
                case 0:
                    switch fi36 {
                    case 4294967295:
                        v68 := GenS10{F0: 2147483647}
                        _ = v68
                    }
                case 32768:
                    var v69 int16 = -799
                    _ = v69
                default:
                    for fi70 := 0; fi70 < 18; fi70++ {
                        v7 = (v7 % ((210) | 1))
                        var v71 int64 = 9223372036854775807
                        if (("hello" + v35) != v35) {
                            v35 = "\n\t"
                            switch fi36 {
                            case 9223372036854775807:
                                v3 = v3
                            case 128:
                                var v72 uint8 = 1
                                _ = v72
                            case -2147483648:
                                v11 = (v11 / (3.14159 / (v11 - v6)))
                            }
                            fmt.Println(v4, v4, v9)
                        } else {
                            switch v35 {
                            case "a":
                                v71 += (v71 & (v71 << ((v9) & 63)))
                            case "foo bar":
                                v7 = v7
                            case "hello":
                                v73 := GenS11{F0: (v35 + "foo bar"), F1: 0, F2: 1}
                                _ = v73
                            default:
                                var v74 int8 = v3
                                _ = v74
                            }
                            var v75 int8 = v3
                            v35 = "foo bar"
                            v7 = 1
                            var v76 uint = v9
                            _ = v75
                            _ = v76
                        }
                        _ = fi70
                        _ = v71
                    }
                }
            }
            v4 += v9
        } else {
            switch ((v35 + v35) + ("a" + (v35 + v35))) {
            case "":
                var v77 uint16 = 65535
                _ = v77
            case "foo bar":
                var v78 float64 = 0.1
                _ = v78
            }
            var v79 bool = true
            _ = v79
        }
        _ = fi36
    }
    fmt.Println(v11)
    for fi80 := 0; fi80 < 10; fi80++ {
        for fi81 := 0; fi81 < 15; fi81++ {
            v6 -= float32((float64((1e10 + 0.1)) / 1e10))
            v82 := GenS12{F0: fi80, F1: (v7 * (v7 * v7)), F2: 42}
            _ = fi81
            _ = v82
        }
        for fi83 := 0; fi83 < 5; fi83++ {
            if true {
                v9 = 0
            } else {
                v11 = v6
                if false {
                    fmt.Println(v9)
                    if true {
                        v35 = v35
                    } else {
                        if false {
                            if (204 < 127) {
                                var v84 float32 = (v6 + float32(-2.25))
                                _ = v84
                            } else {
                                _ = len(v5)
                                if len(v5) > 0 {
                                    v5[0] = fi83
                                }
                                var v85 float64 = -2.25
                                _ = v85
                            }
                            var v86 int64 = 331
                            _ = v86
                        } else {
                            v9 = (v9 ^ (v9 << (((v4 >> ((v9) & 63))) & 63)))
                            fmt.Println(v6, v4, v35)
                        }
                    }
                } else {
                    var v87 uint64 = 15
                    _ = v87
                }
            }
            v3 = (v3 + v3)
            _ = fi83
        }
        v88 := GenS13{F0: (v9 / ((v9) | 1))}
        _ = fi80
        _ = v88
    }
    var v89 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v89:", v89)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v89
    _ = v9
}
