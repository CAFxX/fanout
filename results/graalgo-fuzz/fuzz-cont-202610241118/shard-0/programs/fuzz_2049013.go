package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 int32 }
func genFn0() string {
    if (false || (42 <= 1)) {
        if true {
            var v4 uint32 = 65535
            _ = v4
        } else {
            switch 716 {
            case 218:
                if (-128 != -26) {
                    var v5 string = "a"
                    var v6 float32 = float32(3.14159)
                    _ = v5
                    _ = v6
                }
            case 4294967295:
                switch "hello" {
                case "a":
                    if (65535 < 0) {
                        var v7 uint64 = 1
                        _ = v7
                    }
                case "hello":
                    var v8 float32 = 0.1
                    _ = v8
                default:
                    var v9 float32 = 3.14159
                    _ = v9
                }
            }
        }
        v10 := []int{-2147483648, 290, -51}
        var v11 bool = (172 >= -1)
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 2, F1: 1}
    switch -396 {
    case -9223372036854775808:
        v13 := []int{9223372036854775807}
        _ = v13
    case 722:
        var v14 string = "foo bar"
        _ = v14
    default:
        var v15 int16 = -84
        _ = v15
    }
    _ = v12
    return "\n\t"
}

func main() {
    v1 := []int{}
    var v2 uint = 1
    v3 := make(map[int]int)
    switch "" {
    case "hello":
        var v16 uint16 = 0
        _ = v16
    case "":
        v3[-595] = -1
    case "a":
        v2 = (v2 + (v2 << ((18446744073709551615) & 63)))
    }
    v17 := GenS1{F0: 128, F1: 1}
    v2 = 2
    v18 := GenS2{F0: 42}
    if false {
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 float32 = -2.25
            v21 := v3[(fi19 >> (((v2 + v2)) & 63))]
            v20 -= v20
            _ = fi19
            _ = v20
            _ = v21
        }
        if (-1 == 103) {
            for fi22 := 0; fi22 < 12; fi22++ {
                v1 = append(v1, 32767)
                _ = fi22
            }
            switch -2147483648 {
            case -1:
                for fi23 := 0; fi23 < 19; fi23++ {
                    var v24 int64 = -1
                    _ = fi23
                    _ = v24
                }
            }
            v2 = 926
        } else {
            v2 = v2
            fmt.Println(v2)
        }
        if false {
            if (773 != uint32(v2)) {
                v25 := GenS3{F0: 1, F1: 974}
                v26 := GenS1{F0: 42, F1: 0}
                var v27 uint = (v2 & 65535)
                _ = v25
                _ = v26
                _ = v27
            } else {
                switch 2147483647 {
                case 32768:
                    var v28 uint = (v2 / ((v2) | 1))
                    _ = v28
                case -688:
                    v2 += (v2 * v2)
                }
                var v29 string = "\n\t"
                _ = v29
            }
        } else {
            var v30 int32 = 0
            if (false && (v2 != v2)) {
                for fi31 := 0; fi31 < 7; fi31++ {
                    fmt.Println(v30)
                    _ = fi31
                }
                if (true || true) {
                    fmt.Println(v30, v2)
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch 128 {
                    case 360:
                        v2 += v2
                    case 32768:
                        v2 -= 522
                    }
                    v1 = append(v1, 128)
                }
                var v32 float64 = (1.5 + 0.1)
                _ = v32
            } else {
                v33 := GenS1{F0: 128, F1: 9223372036854775807}
                _ = v33
            }
            _ = v30
        }
    }
    v2 = (v2 ^ 4294967295)
    for fi34 := 0; fi34 < 8; fi34++ {
        v2 -= (v2 >> ((v2) & 63))
        v35 := GenS4{F0: v2, F1: 42, F2: 2}
        _ = fi34
        _ = v35
    }
    if ((184 != 127) && (-90 == 1)) {
        v36 := GenS1{F0: 2, F1: -855}
        var v37 uint16 = 65535
        _ = v36
        _ = v37
    } else {
        v2 += (v2 + (v2 * v2))
    }
    switch "hello" {
    case "\n\t":
        for fi38 := 0; fi38 < 12; fi38++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v39 uint64 = 0
            _ = fi38
            _ = v39
        }
    case "foo bar":
        v2 = (v2 + v2)
    case "":
        if (!true) {
            v40 := GenS5{F0: 4294967295, F1: 736}
            for fi41 := 0; fi41 < 8; fi41++ {
                v2 = 173
                _ = fi41
            }
            if false {
                var v42 uint = (v2 >> (((v2 >> ((v2) & 63))) & 63))
                v42 = v2
                _ = v42
            } else {
                for fi43 := 0; fi43 < 15; fi43++ {
                    switch "foo bar" {
                    case "\n\t":
                        v2 = 255
                    default:
                        fmt.Println(v2, v2)
                    }
                    delete(v3, (fi43 ^ (fi43 + (^(fi43)))))
                    _ = fi43
                }
                if true {
                    var v44 uint = (v2 + 2)
                    if (v44 != (v2 << (((^(v2))) & 63))) {
                        v45 := GenS4{F0: 18446744073709551615, F1: 42, F2: 0}
                        v2 = v2
                        for fi46 := 0; fi46 < 8; fi46++ {
                            v2 += v44
                            _ = fi46
                        }
                        _ = v45
                    } else {
                        v47 := GenS6{F0: 508, F1: -1}
                        v2 -= (v44 ^ (v44 / (((^(v44))) | 1)))
                        _ = v47
                    }
                    v2 += v44
                    _ = v44
                } else {
                    var v48 int64 = 1
                    for fi49 := 0; fi49 < 8; fi49++ {
                        if true {
                            v48 = (v48 >> ((v2) & 63))
                            for fi50 := 0; fi50 < 2; fi50++ {
                                switch "\n\t" {
                                case "\n\t":
                                    var v51 uint32 = 1
                                    _ = v51
                                }
                                v2 = v2
                                if false {
                                    v48 -= v48
                                    switch (fi50 - 4294967295) {
                                    case 2147483647:
                                        switch "hello" {
                                        case "a":
                                            if (1 != 0) {
                                                var v52 uint64 = 18446744073709551615
                                                v53 := v3[fi49]
                                                switch 1 {
                                                case -45:
                                                    fmt.Println(fi49)
                                                }
                                                _ = v52
                                                _ = v53
                                            }
                                        default:
                                            var v54 int64 = int64(v2)
                                            _ = v54
                                        }
                                    case 0:
                                        var v55 string = "\n\t"
                                        _ = v55
                                    }
                                    v2 = v2
                                }
                                if true {
                                    var v56 int8 = -122
                                    var v57 bool = (("foo bar" == ("" + "hello")) || (42 != 255))
                                    _ = v56
                                    _ = v57
                                }
                                _ = fi50
                            }
                            _ = len(v1)
                        } else {
                            if false {
                                _ = len(v3)
                                var v58 int64 = -417
                                v59 := GenS4{F0: v2, F1: 42, F2: 0}
                                _ = v58
                                _ = v59
                            } else {
                                for fi60 := 0; fi60 < 6; fi60++ {
                                    v61 := v3[(fi60 + fi60)]
                                    var v62 int16 = -1
                                    v63 := GenS7{F0: (v48 % (((v48 * -901)) | 1))}
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                }
                                switch "foo bar" {
                                case "a":
                                    v2 += (v2 + v2)
                                case "foo bar":
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch "a" {
                                        case "foo bar":
                                            v2 -= (v2 / (((^(v2))) | 1))
                                        case "":
                                            v48 -= v48
                                        case "a":
                                            v64 := GenS8{F0: v48}
                                            _ = v64
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = (fi49 >> (((v2 ^ 2)) & 63))
                                            }
                                        }
                                    default:
                                        if ("hello" == "") {
                                            v65 := GenS9{F0: 340, F1: 127}
                                            var v66 uint8 = 128
                                            for fi67 := 0; fi67 < 15; fi67++ {
                                                var v68 int8 = -25
                                                fmt.Println(fi49, v68)
                                                _ = fi67
                                                _ = v68
                                            }
                                            _ = v65
                                            _ = v66
                                        } else {
                                            if (true && false) {
                                                var v69 int64 = v48
                                                v2 += 683
                                                _ = v69
                                            } else {
                                                v70 := GenS10{F0: 0}
                                                var v71 uint64 = 42
                                                _ = v70
                                                _ = v71
                                            }
                                        }
                                    }
                                }
                                v2 += (v2 + v2)
                            }
                            _ = len(v1)
                            fmt.Println(v2, v48)
                        }
                        _ = fi49
                    }
                    for fi72 := 0; fi72 < 3; fi72++ {
                        switch fi72 {
                        case 904:
                            v2 = 36
                        }
                        fmt.Println(fi72)
                        _ = fi72
                    }
                    _ = v48
                }
            }
            _ = v40
        } else {
            var v73 uint32 = 141
            for fi74 := 0; fi74 < 15; fi74++ {
                if true {
                    fmt.Println(v2, fi74, v2)
                    var v75 int16 = 1000
                    var v76 string = "foo bar"
                    _ = v75
                    _ = v76
                }
                _ = fi74
            }
            _ = v73
        }
    default:
        if true {
            v2 = v2
        }
    }
    v77 := GenS10{F0: 300}
    if len(v1) > 0 {
        v1[0] = 32767
    }
    if false {
        v78 := GenS9{F0: 1, F1: 46}
        v1 = append(v1, 675)
        v79 := v3[-432]
        _ = v78
        _ = v79
    } else {
        var v80 uint8 = uint8(v2)
        v81 := GenS9{F0: 703, F1: -128}
        var v82 int16 = int16(v80)
        _ = v80
        _ = v81
        _ = v82
    }
    switch -1 {
    case -332:
        if false {
            v2 += 4294967295
        } else {
            var v83 bool = (1 < -9223372036854775808)
            if (!v83) {
                var v84 float64 = 1.5
                if v83 {
                    var v85 int = 198
                    _ = v85
                } else {
                    if v83 {
                        v86 := v3[-2147483648]
                        _ = v86
                    }
                    var v87 int = 361
                    switch v87 {
                    case 128:
                        v83 = (v84 < 0.1)
                    default:
                        var v88 string = "hello"
                        _ = v88
                    }
                    _ = v87
                }
                var v89 int64 = 0
                _ = v84
                _ = v89
            }
            v2 = 142
            _ = v83
        }
    case -728:
        v90 := GenS2{F0: 956}
        _ = v90
    }
    for fi91 := 0; fi91 < 14; fi91++ {
        v2 += v2
        v2 += 65535
        _ = fi91
    }
    v2 = (v2 + (v2 | (v2 >> ((v2) & 63))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v77
}
