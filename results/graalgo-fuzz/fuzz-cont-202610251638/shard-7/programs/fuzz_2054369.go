package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int32; F2 int16 }
type GenS6 struct { F0 string; F1 uint; F2 int64 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 int; F1 uint16; F2 uint }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 string; F1 uint; F2 int }
type GenS11 struct { F0 int; F1 int32 }
type GenS12 struct { F0 int32; F1 uint }

func main() {
    var v1 string = ""
    var v2 bool = false
    var v3 float64 = 3.14159
    var v4 float32 = -2.25
    var v5 int8 = 42
    v6 := GenS1{F0: 2, F1: 645, F2: -9223372036854775808}
    var v7 float64 = (((v3 / 3.14159) - v3) / v3)
    v8 := GenS1{F0: 255, F1: 65535, F2: 0}
    if false {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int32 = 251
            v11 := GenS1{F0: 42, F1: 65535, F2: fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS2{F0: 2, F1: 2147483648}
        _ = fi12
        _ = v13
    }
    v14 := []int{4294967295, 897, 176}
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint8 = 1
        if (!true) {
            var v17 uint16 = 32767
            for fi18 := 0; fi18 < 4; fi18++ {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = fi18
            }
            _ = v17
        }
        if (false && false) {
            if true {
                if true {
                    fmt.Println(v7, v3, v4)
                } else {
                    v19 := GenS2{F0: 2, F1: 598}
                    _ = v19
                }
                for fi20 := 0; fi20 < 4; fi20++ {
                    v14 = append(v14, fi20)
                    var v21 uint16 = 75
                    v16 += (^(v16))
                    v22 := GenS3{F0: fi20, F1: (v16 * v16), F2: 42}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 uint32 = 42
                _ = v23
            } else {
                switch "" {
                case "\n\t":
                    v7 += float64(0.0)
                default:
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
                if false {
                    v14 = append(v14, fi15)
                    _ = len(v14)
                } else {
                    v7 = ((1.5 * v3) / 0.1)
                    v24 := GenS4{F0: 661}
                    if (42 == 18446744073709551615) {
                        switch fi15 {
                        case 1:
                            if len(v14) > 0 {
                                v14[0] = (fi15 / (((fi15 << ((854) & 63))) | 1))
                            }
                        case -169:
                            var v25 float32 = 0.0
                            _ = v25
                        case 42:
                            for fi26 := 0; fi26 < 20; fi26++ {
                                _ = len(v14)
                                v7 -= 0.1
                                if (fi15 < (fi26 - fi15)) {
                                    v4 = 1.5
                                } else {
                                    var v27 int = fi15
                                    switch v1 {
                                    case "":
                                        v14 = append(v14, (v27 % (((fi15 >> ((0) & 63))) | 1)))
                                    default:
                                        v2 = v2
                                    }
                                    _ = v27
                                }
                                _ = fi26
                            }
                        default:
                            for fi28 := 0; fi28 < 10; fi28++ {
                                var v29 uint64 = 42
                                if false {
                                    v30 := GenS5{F0: (v29 & 0), F1: 2147483647, F2: 13}
                                    switch ("" + v1) {
                                    case "foo bar":
                                        if v2 {
                                            if len(v14) > 0 {
                                                v14[0] = 32768
                                            }
                                            v2 = (!false)
                                            v29 = v29
                                            var v31 uint32 = 203
                                            _ = v31
                                        } else {
                                            if v2 {
                                                v32 := GenS4{F0: uint32(v16)}
                                                fmt.Println(v2, fi28, fi28)
                                                _ = v32
                                            }
                                        }
                                    case "hello":
                                        v33 := GenS6{F0: v1, F1: 637, F2: 542}
                                        _ = v33
                                    default:
                                        switch "foo bar" {
                                        case "hello":
                                            v5 -= (-(v5))
                                        case "a":
                                            var v34 float32 = v4
                                            _ = v34
                                        }
                                    }
                                    _ = v30
                                }
                                v16 = 127
                                _ = fi28
                                _ = v29
                            }
                        }
                        switch v1 {
                        case "foo bar":
                            for fi35 := 0; fi35 < 8; fi35++ {
                                var v36 int = fi35
                                var v37 uint8 = 1
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        case "\n\t":
                            for fi38 := 0; fi38 < 9; fi38++ {
                                switch v1 {
                                case "hello":
                                    var v39 uint = 0
                                    _ = v39
                                }
                                _ = fi38
                            }
                        case "hello":
                            switch v1 {
                            case "foo bar":
                                v7 = 3.14159
                            case "hello":
                                fmt.Println(v7, fi15, v5)
                            case "a":
                                v1 = v1
                            default:
                                v3 += 3.14159
                            }
                        default:
                            if ((!false) || (255 != 75)) {
                                fmt.Println(v3)
                                v40 := GenS4{F0: 1}
                                switch (^(fi15)) {
                                case -682:
                                    v41 := GenS7{F0: -993, F1: -115}
                                    _ = v41
                                case 2147483647:
                                    var v42 float64 = v3
                                    _ = v42
                                case 4294967295:
                                    if len(v14) > 0 {
                                        v14[0] = fi15
                                    }
                                default:
                                    for fi43 := 0; fi43 < 1; fi43++ {
                                        var v44 int16 = 32767
                                        fmt.Println(v44, v3, v3)
                                        _ = fi43
                                        _ = v44
                                    }
                                }
                                _ = v40
                            }
                        }
                        v45 := GenS8{F0: (fi15 | -1), F1: 180, F2: 4294967295}
                        _ = v45
                    }
                    _ = v24
                }
            }
            fmt.Println(v5)
            var v46 int16 = 32767
            _ = v46
        } else {
            if false {
                v2 = v2
            } else {
                if (!true) {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
                switch -1 {
                case 591:
                    switch "\n\t" {
                    case "\n\t":
                        var v47 int64 = 1
                        _ = v47
                    default:
                        v48 := GenS8{F0: fi15, F1: 2, F2: 18446744073709551615}
                        _ = v48
                    }
                }
            }
            fmt.Println(v4)
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v7)
    v49 := GenS4{F0: 776}
    v5 -= v5
    v50 := GenS3{F0: 2147483647, F1: 224, F2: int32(v5)}
    switch (v1 + ("\n\t" + (v1 + v1))) {
    case "a":
        if v2 {
            v2 = ("hello" <= ((v1 + "hello") + v1))
        } else {
            switch "a" {
            case "":
                for fi51 := 0; fi51 < 20; fi51++ {
                    var v52 uint32 = 314
                    var v53 int8 = -40
                    var v54 uint64 = 1
                    _ = fi51
                    _ = v52
                    _ = v53
                    _ = v54
                }
            }
        }
    }
    if true {
        fmt.Println(v2, v2)
        switch v1 {
        case "":
            for fi55 := 0; fi55 < 20; fi55++ {
                if (817 <= 18446744073709551615) {
                    if false {
                        var v56 int = (fi55 + 9223372036854775807)
                        var v57 uint32 = 241
                        _ = v56
                        _ = v57
                    } else {
                        fmt.Println(fi55)
                        fmt.Println(v1)
                    }
                    v4 -= v4
                    if v2 {
                        v14 = append(v14, 32768)
                        switch (fi55 << ((423) & 63)) {
                        case 1:
                            _ = len(v14)
                        default:
                            for fi58 := 0; fi58 < 14; fi58++ {
                                var v59 int32 = 729
                                v4 = v4
                                if v2 {
                                    if v2 {
                                        var v60 uint16 = 1
                                        switch fi55 {
                                        case 4294967295:
                                            v61 := GenS9{F0: 0, F1: 466}
                                            _ = v61
                                        case 128:
                                            if len(v14) > 0 {
                                                v14[0] = -2147483648
                                            }
                                        default:
                                            if true {
                                                v5 -= v5
                                            } else {
                                                if true {
                                                    v1 = "a"
                                                    v4 += v4
                                                    v5 += -128
                                                } else {
                                                    v62 := GenS10{F0: "a", F1: 2, F2: 9223372036854775807}
                                                    _ = v62
                                                }
                                                var v63 uint8 = 2
                                                _ = len(v14)
                                                _ = v63
                                            }
                                        }
                                        _ = v60
                                    }
                                    fmt.Println(fi55)
                                    var v64 bool = (v1 == "foo bar")
                                    _ = v64
                                }
                                _ = fi58
                                _ = v59
                            }
                        }
                    }
                } else {
                    var v65 string = "\n\t"
                    var v66 int8 = (v5 << ((441) & 63))
                    _ = v65
                    _ = v66
                }
                _ = fi55
            }
        case "hello":
            v1 = (v1 + ("a" + (("\n\t" + v1) + ((v1 + "\n\t") + v1))))
        }
    } else {
        if true {
            var v67 int16 = 1
            for fi68 := 0; fi68 < 18; fi68++ {
                for fi69 := 0; fi69 < 14; fi69++ {
                    switch fi69 {
                    case 128:
                        v70 := GenS11{F0: fi68, F1: 851}
                        _ = v70
                    case -644:
                        v71 := GenS12{F0: -1, F1: 18446744073709551615}
                        _ = v71
                    default:
                        v2 = true
                    }
                    _ = fi69
                }
                _ = fi68
            }
            _ = v67
        }
        for fi72 := 0; fi72 < 10; fi72++ {
            switch v1 {
            case "foo bar":
                v73 := GenS10{F0: "hello", F1: 862, F2: fi72}
                _ = v73
            case "":
                fmt.Println(v2)
            case "\n\t":
                if (v4 == float32(0.0)) {
                    switch fi72 {
                    case 478:
                        var v74 int32 = 332
                        _ = v74
                    case -321:
                        var v75 uint16 = 2
                        _ = v75
                    case 2147483647:
                        v7 = (v3 * (v7 + float64((v4 / (v4 + (float32(0.1) / float32(float64(-2.25))))))))
                    }
                }
            }
            v7 = 3.14159
            _ = fi72
        }
    }
    var v76 float32 = (0.0 / (v4 - (0.0 * float32((((v7 - v3) + v3) + v7)))))
    v77 := GenS6{F0: "foo bar", F1: 0, F2: -421}
    var v78 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v7
    _ = v76
    _ = v77
    _ = v78
    _ = v8
}
