package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint; F2 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 int }
func genFn0(p0 uint, p1 uint16, p2 uint8) int8 {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: (uint8(p1) != 62)}
        p1 += (p1 % ((p1) | 1))
        p1 -= p1
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "hello":
        for fi9 := 0; fi9 < 12; fi9++ {
            switch (fi9 - 127) {
            case 158:
                for fi10 := 0; fi10 < 4; fi10++ {
                    p1 -= (p1 | p1)
                    switch fi10 {
                    case -907:
                        p1 += p1
                    default:
                        var v11 int8 = -1
                        _ = v11
                    }
                    v12 := GenS1{F0: true}
                    _ = fi10
                    _ = v12
                }
            case 4294967295:
                var v13 uint64 = 4294967295
                _ = v13
            default:
                switch fi9 {
                case 2147483647:
                    v14 := GenS2{F0: -1}
                    _ = v14
                case 9223372036854775807:
                    p2 = (p2 * p2)
                case -158:
                    p2 -= (p2 | 156)
                default:
                    switch fi9 {
                    case 32767:
                        p2 -= p2
                    case 2147483647:
                        p0 -= (p0 | p0)
                    case -9223372036854775808:
                        if false {
                            for fi15 := 0; fi15 < 5; fi15++ {
                                v16 := []int{}
                                switch ("foo bar" + (("foo bar" + ("" + "a")) + (("foo bar" + "\n\t") + "a"))) {
                                case "":
                                    var v17 int64 = -9223372036854775808
                                    _ = v17
                                }
                                _ = fi15
                                _ = v16
                            }
                            var v18 int8 = 18
                            var v19 uint64 = 42
                            _ = v18
                            _ = v19
                        } else {
                            p2 = 2
                            p1 += p1
                        }
                    default:
                        var v20 float32 = (1.5 / float32((-2.25 - 1.5)))
                        _ = v20
                    }
                }
            }
            if true {
                fmt.Println(fi9)
            }
            _ = fi9
        }
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := GenS3{F0: -128, F1: p2}
        var v23 int16 = 836
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1() uint16 {
    if ((29 >= 1) && (262 <= 512)) {
        var v24 bool = true
        _ = v24
    }
    var v25 int8 = -128
    _ = v25
    return 379
}

func main() {
    v1 := []int{0, -1}
    v2 := []uint32{}
    v3 := []string{}
    var v4 uint32 = 2147483647
    var v5 int64 = -1
    v6 := []int{32768, 0, 32768, int(v5)}
    if true {
        fmt.Println(v4, v5)
        if false {
            for fi26 := 0; fi26 < 10; fi26++ {
                switch "" {
                case "a":
                    v5 = (^(v5))
                default:
                    if (-128 >= 0) {
                        v4 -= v4
                        var v27 string = "\n\t"
                        if (("" + v27) > "hello") {
                            for fi28 := 0; fi28 < 3; fi28++ {
                                if false {
                                    var v29 int8 = 72
                                    _ = v29
                                } else {
                                    v30 := GenS4{F0: -770, F1: 2}
                                    _ = v30
                                }
                                _ = fi28
                            }
                            v31 := GenS4{F0: 1000, F1: 953}
                            _ = v31
                        } else {
                            if (138 < 128) {
                                if true {
                                    fmt.Println(fi26, fi26, fi26)
                                    v6 = append(v6, (^(fi26)))
                                } else {
                                    if false {
                                        if true {
                                            if ((v5 >> ((0) & 63)) >= -9223372036854775808) {
                                                if false {
                                                    v5 = v5
                                                    if true {
                                                        v27 = v27
                                                    }
                                                    for fi32 := 0; fi32 < 13; fi32++ {
                                                        var v33 uint8 = uint8(v4)
                                                        var v34 bool = false
                                                        if (!v34) {
                                                            fmt.Println(v27, v27)
                                                        } else {
                                                            fmt.Println(v33, v33, fi32)
                                                            switch v27 {
                                                            case "":
                                                                var v35 int = -947
                                                                _ = v35
                                                            default:
                                                                var v36 uint8 = (v33 % ((v33) | 1))
                                                                _ = v36
                                                            }
                                                        }
                                                        fmt.Println(v5, fi26)
                                                        if (v34 && v34) {
                                                            switch v27 {
                                                            case "":
                                                                if (v34 && (!(!((v34 && v34) && v34)))) {
                                                                    switch ("" + v27) {
                                                                    case "a":
                                                                        if v34 {
                                                                            v27 = "\n\t"
                                                                            v27 = v27
                                                                        } else {
                                                                            _ = len(v2)
                                                                            var v37 uint8 = 255
                                                                            var v38 int64 = v5
                                                                            _ = len(v2)
                                                                            _ = v37
                                                                            _ = v38
                                                                        }
                                                                    default:
                                                                        v4 -= v4
                                                                    }
                                                                    var v39 uint8 = v33
                                                                    _ = v39
                                                                } else {
                                                                    v4 = 947
                                                                    v5 = -9223372036854775808
                                                                    _ = len(v6)
                                                                }
                                                            case "hello":
                                                                var v40 uint32 = v4
                                                                _ = v40
                                                            default:
                                                                fmt.Println(v33, fi32, v33)
                                                            }
                                                            v5 -= int64(v4)
                                                        }
                                                        _ = fi32
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                }
                                                v41 := GenS4{F0: 1000, F1: 4294967295}
                                                _ = len(v6)
                                                _ = v41
                                            }
                                        } else {
                                            v42 := GenS5{F0: 20, F1: 1, F2: -1}
                                            _ = v42
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 65535
                                        }
                                    }
                                    var v43 int16 = 854
                                    _ = v43
                                }
                            }
                        }
                        _ = v27
                    }
                }
                v4 -= 4294967295
                _ = fi26
            }
            v5 = (v5 - (v5 >> ((18446744073709551615) & 63)))
        }
    } else {
        v44 := GenS4{F0: 32767, F1: 519}
        _ = v44
    }
    var v45 uint64 = 1
    fmt.Println(v4, v4)
    v46 := GenS1{F0: true}
    if (!(v45 >= 18446744073709551615)) {
        var v47 float32 = 3.14159
        _ = v47
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v48 := GenS1{F0: false}
        _ = v48
    }
    switch 9223372036854775807 {
    case 0:
        v49 := GenS1{F0: (!true)}
        _ = v49
    case 128:
        v50 := GenS6{F0: (v5 / (((v5 | (v5 * v5))) | 1))}
        _ = v50
    case 524:
        var v51 bool = (("\n\t" + "a") >= "\n\t")
        _ = v51
    default:
        switch 1 {
        case 556:
            switch 128 {
            case 32767:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case 1:
                var v52 int32 = 42
                _ = v52
            default:
                for fi53 := 0; fi53 < 1; fi53++ {
                    if true {
                        switch "a" {
                        case "foo bar":
                            for fi54 := 0; fi54 < 5; fi54++ {
                                v4 -= (v4 * 40)
                                if true {
                                    _ = len(v1)
                                    v55 := GenS7{F0: 0}
                                    _ = v55
                                }
                                var v56 uint32 = v4
                                _ = fi54
                                _ = v56
                            }
                        case "":
                            var v57 int16 = 693
                            _ = v57
                        }
                    } else {
                        for fi58 := 0; fi58 < 18; fi58++ {
                            var v59 uint = 635
                            _ = fi58
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 4; fi60++ {
                            switch (fi53 + 167) {
                            case 100:
                                v45 = 4294967295
                                var v61 int16 = 1
                                switch 731 {
                                case 9223372036854775807:
                                    fmt.Println(v4, v4, v4)
                                    v4 -= 2147483648
                                case 4294967295:
                                    switch -183 {
                                    case 9223372036854775807:
                                        var v62 int64 = (v5 & -9223372036854775808)
                                        _ = v62
                                    }
                                case -757:
                                    v4 = (v4 << ((889) & 63))
                                default:
                                    var v63 int64 = 42
                                    _ = v63
                                }
                                _ = v61
                            case 32768:
                                var v64 int = fi60
                                _ = v64
                            case 269:
                                v45 = v45
                            }
                            v4 = 142
                            _ = fi60
                        }
                        switch (-(fi53)) {
                        case 0:
                            v65 := GenS8{F0: (v45 + (v45 | v45)), F1: (fi53 ^ 127)}
                            _ = v65
                        case 9223372036854775807:
                            var v66 uint32 = v4
                            _ = v66
                        default:
                            v4 += (v4 - 65535)
                        }
                    }
                    _ = fi53
                }
            }
        case 554:
            v5 += (v5 | (v5 ^ (v5 << ((1) & 63))))
        }
    }
    v5 += 0
    var v67 uint32 = v4
    switch "a" {
    case "foo bar":
        v4 = v67
    case "":
        fmt.Println(v67)
    case "\n\t":
        v67 = v67
    default:
        switch "foo bar" {
        case "a":
            v67 += (v67 % (((v4 | (v67 + 362))) | 1))
        }
    }
    v45 -= (v45 << ((4294967295) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v67
}
