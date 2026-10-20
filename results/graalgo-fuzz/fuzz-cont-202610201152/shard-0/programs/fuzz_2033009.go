package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint16; F2 int }
type GenS3 struct { F0 uint8; F1 int; F2 uint32 }
type GenS4 struct { F0 uint16; F1 string; F2 bool }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint64; F1 uint16 }
type GenS11 struct { F0 int32; F1 int16; F2 int32 }
func genFn0(p0 int64) int64 {
    if (42 < 127) {
        v7 := GenS1{F0: false}
        _ = v7
    } else {
        if (0 != 2147483647) {
            p0 -= p0
        } else {
            fmt.Println(p0)
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            if true {
                fmt.Println(p0, fi8)
            }
            _ = fi8
        }
    }
    var v9 uint16 = 32767
    _ = p0
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]uint32)
    var v3 string = "a"
    var v4 float32 = ((float32(3.14159) - (3.14159 * (3.14159 + float32((-2.25 + 0.1))))) / 1e10)
    var v5 int8 = 25
    v6 := make(map[string]string)
    var v10 uint16 = v1
    if (2147483647 == -737) {
        v11 := []int{128}
        _ = v11
    } else {
        switch v3 {
        case "\n\t":
            v10 = v1
        }
    }
    var v12 float64 = (3.14159 + (float64((float32(0.0) + float32(-2.25))) / 0.1))
    v13 := []int{32768, -805, -9223372036854775808, 1}
    fmt.Println(v4)
    if false {
        if (755 >= -1) {
            v14 := GenS2{F0: 1, F1: v1, F2: -507}
            _ = v14
        } else {
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 uint64 = 4294967295
                fmt.Println(v4)
                _ = fi15
                _ = v16
            }
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            var v17 bool = (0 == 50)
            _ = v17
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: false}
            switch v3 {
            case "":
                fmt.Println(v3, v12, v5)
            case "foo bar":
                _ = len(v6)
            }
            _ = fi18
            _ = v19
        }
    }
    switch 2147483647 {
    case 127:
        var v20 int16 = -718
        _ = v20
    case 4294967295:
        fmt.Println(v4, v5)
    case -9223372036854775808:
        v3 = (v3 + "")
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v4 = v4
        _ = fi21
    }
    switch v3 {
    case "foo bar":
        var v22 uint8 = 2
        _ = v22
    }
    if false {
        v5 = 4
        var v23 int = 2147483647
        _ = v23
    } else {
        switch 2147483647 {
        case 128:
            var v24 int32 = 1
            _ = v24
        case 9223372036854775807:
            if (1 != int(v5)) {
                v4 = (1.5 / v4)
                for fi25 := 0; fi25 < 9; fi25++ {
                    v3 = ((("hello" + v3) + v3) + v3)
                    _ = fi25
                }
                if (v10 >= v10) {
                    if (-1 >= -293) {
                        var v26 int = 26
                        _ = v26
                    }
                } else {
                    switch (((v3 + v3) + v3) + v3) {
                    case "hello":
                        for fi27 := 0; fi27 < 18; fi27++ {
                            var v28 float32 = (0.1 - v4)
                            v4 += float32(v12)
                            if false {
                                fmt.Println(v1, v10)
                                v12 -= v12
                            }
                            _ = fi27
                            _ = v28
                        }
                    case "\n\t":
                        var v29 uint32 = 657
                        _ = v29
                    default:
                        v30 := GenS3{F0: 0, F1: 512, F2: 42}
                        _ = v30
                    }
                    if (0 != 2) {
                        var v31 uint16 = 65535
                        _ = v31
                    } else {
                        var v32 int32 = -342
                        v33 := v2[v3]
                        _ = v32
                        _ = v33
                    }
                }
            }
        default:
            var v34 uint8 = 2
            _ = v34
        }
        var v35 float64 = (0.0 + float64(v4))
        var v36 float32 = v4
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 uint32 = 0
        for fi39 := 0; fi39 < 17; fi39++ {
            if (v3 < v3) {
                if (127 < 1) {
                    if true {
                        var v40 int64 = -293
                        v13 = append(v13, (fi39 ^ -2147483648))
                        _ = v40
                    } else {
                        if false {
                            var v41 int64 = 42
                            switch (("hello" + v3) + "a") {
                            case "\n\t":
                                if (((-(v41)) > -699) && false) {
                                    v42 := GenS4{F0: (v1 | (^(v10))), F1: v3, F2: false}
                                    var v43 int64 = (v41 | 9223372036854775807)
                                    delete(v6, v3)
                                    v44 := GenS5{F0: 4294967295, F1: 506}
                                    switch v3 {
                                    case "\n\t":
                                        switch fi37 {
                                        case 4294967295:
                                            if false {
                                                v13 = append(v13, (fi39 >> ((469) & 63)))
                                                v12 -= v12
                                            }
                                        case 2147483647:
                                            v5 = -26
                                        case 127:
                                            if len(v13) > 0 {
                                                v13[0] = 983
                                            }
                                        default:
                                            v45 := GenS3{F0: 1, F1: -2147483648, F2: (v38 ^ 42)}
                                            _ = v45
                                        }
                                    case "foo bar":
                                        v46 := GenS6{F0: 1}
                                        _ = v46
                                    case "a":
                                        var v47 uint = 65535
                                        _ = v47
                                    default:
                                        if ((v5 % ((v5) | 1)) != v5) {
                                            if (255 > 208) {
                                                v10 += 1
                                                v1 = v10
                                            }
                                            var v48 float64 = 1.5
                                            switch v3 {
                                            case "a":
                                                if (true || false) {
                                                    v38 -= (v38 / (((^(v38))) | 1))
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                }
                                            }
                                            _ = v48
                                        }
                                    }
                                    _ = len(v2)
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                            }
                            fmt.Println(v10)
                            _ = v41
                        }
                    }
                    switch (^(fi37)) {
                    case 4294967295:
                        v12 -= (v12 * (v12 * 1e10))
                    default:
                        fmt.Println(v10, fi39)
                        switch v3 {
                        case "foo bar":
                            if (!(fi39 <= (^(fi37)))) {
                                switch (^(fi37)) {
                                case 2147483647:
                                    v49 := v6[""]
                                    _ = v49
                                case 4294967295:
                                    v4 += v4
                                default:
                                    switch (v3 + v3) {
                                    case "hello":
                                        v50 := GenS7{F0: true, F1: (^(v1))}
                                        _ = v50
                                    case "":
                                        v13 = append(v13, (fi37 | -1))
                                    }
                                }
                            } else {
                                var v51 uint = 42
                                _ = v51
                            }
                        case "hello":
                            switch (fi39 / ((fi39) | 1)) {
                            case 674:
                                v52 := GenS1{F0: false}
                                _ = v52
                            case -9223372036854775808:
                                var v53 int = -777
                                _ = v53
                            case 839:
                                v13 = append(v13, (-(fi37)))
                            default:
                                var v54 int32 = -1
                                _ = v54
                            }
                        case "":
                            fmt.Println(v4)
                        default:
                            if (v5 != -1) {
                                v38 += 2
                                v12 = ((0.1 + (v12 - ((-2.25 / 1e10) + (-2.25 / (float64(v4) + v12))))) * 0.0)
                                switch (fi39 >> ((255) & 63)) {
                                case -170:
                                    switch v3 {
                                    case "hello":
                                        var v55 uint16 = 2
                                        _ = v55
                                    case "\n\t":
                                        v13 = append(v13, fi37)
                                    default:
                                        var v56 int16 = 974
                                        var v57 int32 = 566
                                        v58 := GenS6{F0: (v56 * v56)}
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    }
                                default:
                                    var v59 bool = false
                                    _ = v59
                                }
                                fmt.Println(v3)
                            } else {
                                switch int(v10) {
                                case -1:
                                    var v60 int16 = 1
                                    _ = v60
                                default:
                                    v38 += v38
                                }
                                v4 -= v4
                            }
                        }
                        fmt.Println(v10, v38, v5)
                    }
                }
                if (18446744073709551615 != 648) {
                    delete(v6, (v3 + ((v3 + v3) + "")))
                    var v61 int16 = -952
                    _ = v61
                } else {
                    switch v3 {
                    case "foo bar":
                        _ = len(v6)
                    case "":
                        v62 := GenS8{F0: (^(v38)), F1: 0}
                        _ = v62
                    }
                }
                _ = len(v6)
            } else {
                v63 := GenS7{F0: true, F1: v1}
                var v64 string = v3
                v64 = (v3 + "")
                _ = v63
                _ = v64
            }
            v1 += (v10 & 32767)
            if true {
                delete(v6, "\n\t")
                switch v3 {
                case "hello":
                    v5 -= v5
                }
                var v65 uint8 = 42
                _ = v65
            } else {
                if false {
                    if (0 != 4294967295) {
                        v66 := GenS9{F0: 42}
                        var v67 uint8 = 115
                        _ = v66
                        _ = v67
                    } else {
                        v5 = (v5 % (((-(v5))) | 1))
                    }
                    v68 := GenS10{F0: 42, F1: v1}
                    var v69 bool = false
                    _ = v68
                    _ = v69
                }
                var v70 uint16 = 72
                v71 := GenS5{F0: 65535, F1: 471}
                _ = v70
                _ = v71
            }
            _ = fi39
        }
        switch "a" {
        case "hello":
            if (4294967295 < 18446744073709551615) {
                for fi72 := 0; fi72 < 2; fi72++ {
                    v73 := GenS8{F0: v38, F1: -1}
                    var v74 uint16 = (^(v1))
                    _ = fi72
                    _ = v73
                    _ = v74
                }
            } else {
                v75 := GenS11{F0: int32(v10), F1: 152, F2: 42}
                fmt.Println(v5)
                _ = v75
            }
        }
        _ = fi37
        _ = v38
    }
    var v76 float64 = 0.0
    v77 := GenS2{F0: 42, F1: v10, F2: -870}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
}
