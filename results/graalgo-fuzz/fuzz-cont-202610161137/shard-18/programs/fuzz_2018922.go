package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int16; F1 uint32; F2 bool }
type GenS4 struct { F0 int; F1 int32; F2 uint16 }
func genFn0() uint {
    var v4 uint16 = 29
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{(fi5 | fi5), -395}
        v4 = (v4 & 709)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: v4}
            v10 := []int{}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    _ = v4
    return 720
}
func genFn1(p0 int32, p1 uint64, p2 int16) int16 {
    if (true || (false || false)) {
        var v11 float32 = 0.1
        p1 = p1
        _ = v11
    }
    var v12 int64 = 9223372036854775807
    p0 -= 173
    v13 := []int{42, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return (p2 >> ((42) & 63))
}

func main() {
    var v1 uint8 = 94
    v2 := make(map[int]int)
    v3 := []uint32{}
    if true {
        v1 = (v1 % ((91) | 1))
        v14 := v2[-806]
        for fi15 := 0; fi15 < 13; fi15++ {
            v1 = 128
            _ = fi15
        }
        _ = v14
    } else {
        v16 := GenS2{F0: 837, F1: genFn1(-845, 818, -1)}
        _ = v16
    }
    v1 = 127
    var v17 uint32 = 269
    fmt.Println(v1, v17)
    v1 += 0
    _ = len(v2)
    v18 := GenS3{F0: 42, F1: 0, F2: false}
    v1 -= v1
    fmt.Println(v17, v17)
    if false {
        v1 = 159
        _ = len(v2)
        v19 := GenS4{F0: 9223372036854775807, F1: 0, F2: 42}
        _ = v19
    }
    if true {
        if (18446744073709551615 > 657) {
            v20 := GenS1{F0: 356}
            if false {
                for fi21 := 0; fi21 < 4; fi21++ {
                    v17 += 11
                    _ = len(v2)
                    v1 -= (^(v1))
                    _ = fi21
                }
                v17 -= 316
                fmt.Println(v17, v17)
            } else {
                var v22 int64 = int64(v17)
                var v23 int = 611
                _ = v22
                _ = v23
            }
            if (!false) {
                for fi24 := 0; fi24 < 16; fi24++ {
                    fmt.Println(fi24, fi24, v1)
                    if (-1 > 2147483647) {
                        fmt.Println(fi24)
                        v1 = 2
                    }
                    _ = fi24
                }
                if false {
                    v25 := GenS3{F0: -1, F1: (^(v17)), F2: (!true)}
                    for fi26 := 0; fi26 < 1; fi26++ {
                        for fi27 := 0; fi27 < 20; fi27++ {
                            v1 = (^(v1))
                            _ = fi27
                        }
                        if (2 == 417) {
                            for fi28 := 0; fi28 < 9; fi28++ {
                                var v29 uint16 = 2
                                _ = fi28
                                _ = v29
                            }
                            var v30 uint = 65535
                            if len(v3) > 0 {
                                v3[0] = 2
                            }
                            _ = v30
                        }
                        _ = fi26
                    }
                    if (true && ((false && (!false)) || (true && false))) {
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            if false {
                                v1 -= (v1 << ((genFn0()) & 63))
                                if false {
                                    for fi31 := 0; fi31 < 19; fi31++ {
                                        if ((-2.25 + float64(3.14159)) >= 0.0) {
                                            var v32 float64 = 0.1
                                            v17 -= (v17 + (v17 | 2))
                                            _ = v32
                                        }
                                        if true {
                                            v1 = 128
                                            var v33 uint64 = 1
                                            _ = v33
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                var v34 uint32 = (v17 - v17)
                                                _ = v34
                                            case "foo bar":
                                                v35 := v2[fi31]
                                                _ = v35
                                            case "\n\t":
                                                v17 += 42
                                            default:
                                                v36 := v2[(fi31 ^ (-(fi31)))]
                                                _ = v36
                                            }
                                        }
                                        _ = fi31
                                    }
                                }
                            }
                        case "a":
                            if true {
                                if false {
                                    v17 -= 2147483648
                                } else {
                                    v2[933] = -267
                                    fmt.Println(v17, v1)
                                }
                                v17 = v17
                            } else {
                                switch "\n\t" {
                                case "a":
                                    var v37 int16 = -32768
                                    _ = v37
                                case "foo bar":
                                    for fi38 := 0; fi38 < 13; fi38++ {
                                        if false {
                                            switch fi38 {
                                            case 0:
                                                fmt.Println(fi38)
                                            }
                                            if false {
                                                v1 = (v1 >> ((545) & 63))
                                                var v39 uint8 = (v1 / (((v1 ^ v1)) | 1))
                                                fmt.Println(v1, v39)
                                                _ = v39
                                            }
                                            switch 32767 {
                                            case -349:
                                                fmt.Println(fi38)
                                            case 0:
                                                v17 = 762
                                            }
                                        }
                                        v17 += uint32(v1)
                                        _ = fi38
                                    }
                                case "hello":
                                    var v40 int = 9223372036854775807
                                    _ = v40
                                default:
                                    _ = len(v2)
                                }
                                v1 = (v1 * 1)
                                var v41 int8 = 0
                                _ = v41
                            }
                        case "hello":
                            for fi42 := 0; fi42 < 7; fi42++ {
                                for fi43 := 0; fi43 < 20; fi43++ {
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                    _ = fi43
                                }
                                switch fi42 {
                                case 295:
                                    fmt.Println(v17, v1, v1)
                                case 0:
                                    _ = len(v3)
                                case 127:
                                    v17 -= uint32(v1)
                                default:
                                    if true {
                                        v44 := GenS4{F0: fi42, F1: 42, F2: 32767}
                                        _ = v44
                                    } else {
                                        var v45 uint16 = 65535
                                        v1 -= 2
                                        _ = v45
                                    }
                                }
                                _ = fi42
                            }
                        default:
                            v46 := GenS1{F0: 65535}
                            _ = v46
                        }
                    }
                    _ = v25
                } else {
                    v1 = (^(v1))
                    switch 42 {
                    case 0:
                        v17 = v17
                    case -298:
                        for fi47 := 0; fi47 < 4; fi47++ {
                            fmt.Println(v1)
                            if false {
                                var v48 int64 = 0
                                switch fi47 {
                                case -9223372036854775808:
                                    fmt.Println(v48, v1, v48)
                                case 35:
                                    var v49 uint64 = 0
                                    _ = v49
                                default:
                                    fmt.Println(fi47)
                                }
                                _ = v48
                            } else {
                                switch -9223372036854775808 {
                                case 32767:
                                    v50 := GenS4{F0: fi47, F1: 42, F2: 65}
                                    _ = v50
                                case -1:
                                    v1 += (^(v1))
                                case 2147483647:
                                    v1 -= v1
                                default:
                                    v2[fi47] = -2147483648
                                }
                            }
                            v51 := GenS1{F0: 0}
                            _ = fi47
                            _ = v51
                        }
                    default:
                        var v52 bool = true
                        _ = v52
                    }
                    var v53 int32 = -2147483648
                    _ = v53
                }
            }
            _ = v20
        } else {
            for fi54 := 0; fi54 < 6; fi54++ {
                switch fi54 {
                case 1:
                    v17 = v17
                case 815:
                    v17 -= v17
                }
                switch int(v17) {
                case 42:
                    if false {
                        v2[4294967295] = fi54
                    }
                case 1:
                    var v55 bool = true
                    _ = v55
                default:
                    v1 -= (^(v1))
                }
                _ = fi54
            }
            var v56 uint8 = v1
            _ = v56
        }
        var v57 int32 = -689
        _ = v57
    } else {
        _ = len(v3)
    }
    if true {
        v1 = 255
        for fi58 := 0; fi58 < 4; fi58++ {
            if (!(65535 < 361)) {
                v17 -= (^(v17))
                for fi59 := 0; fi59 < 11; fi59++ {
                    fmt.Println(v17, fi58, v17)
                    var v60 int32 = 42
                    fmt.Println(fi58, v60, v60)
                    _ = fi59
                    _ = v60
                }
            } else {
                v17 = 919
                v1 += 244
            }
            var v61 int32 = 1
            v1 -= 42
            _ = fi58
            _ = v61
        }
        var v62 uint16 = 32767
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
