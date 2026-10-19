package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 string }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 int64; F2 int }
type GenS7 struct { F0 int16; F1 int32; F2 int }
func genFn0(p0 uint64) uint64 {
    if false {
        v7 := GenS1{F0: 1, F1: 65535, F2: "\n\t"}
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS2{F0: 0, F1: (p0 + (^(p0)))}
            switch (fi8 + fi8) {
            case 2147483647:
                for fi10 := 0; fi10 < 5; fi10++ {
                    switch fi10 {
                    case 127:
                        v11 := []int{}
                        _ = v11
                    default:
                        v12 := GenS1{F0: 128, F1: 434, F2: "a"}
                        _ = v12
                    }
                    v13 := []int{}
                    _ = fi10
                    _ = v13
                }
            case 127:
                var v14 int8 = -1
                _ = v14
            case 128:
                for fi15 := 0; fi15 < 2; fi15++ {
                    var v16 string = "\n\t"
                    p0 -= (p0 / ((0) | 1))
                    v17 := []int{}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = fi8
            _ = v9
        }
        for fi18 := 0; fi18 < 9; fi18++ {
            if (true || true) {
                v19 := []int{}
                p0 -= 2
                _ = v19
            } else {
                if false {
                    fmt.Println(p0, p0)
                    switch "foo bar" {
                    case "foo bar":
                        v20 := []int{0, 128, -73}
                        _ = v20
                    case "hello":
                        fmt.Println(fi18, p0)
                    }
                    v21 := []int{}
                    _ = v21
                } else {
                    switch ("\n\t" + (("a" + ("a" + ("a" + "\n\t"))) + "hello")) {
                    case "foo bar":
                        var v22 int8 = 0
                        _ = v22
                    case "hello":
                        p0 = (p0 % ((p0) | 1))
                    default:
                        fmt.Println(fi18)
                    }
                }
            }
            p0 = p0
            v23 := []int{fi18, (fi18 + (fi18 % ((fi18) | 1))), (fi18 | 127)}
            _ = fi18
            _ = v23
        }
        _ = v7
    }
    fmt.Println(p0)
    if true {
        if false {
            fmt.Println(p0)
        } else {
            v24 := GenS2{F0: 42, F1: (p0 | 423)}
            var v25 uint16 = 255
            _ = v24
            _ = v25
        }
        switch -558 {
        case 32768:
            v26 := []int{270, 42, 612, 128}
            _ = v26
        case -2147483648:
            v27 := GenS1{F0: 240, F1: 42, F2: ("\n\t" + ("hello" + ""))}
            _ = v27
        case 42:
            for fi28 := 0; fi28 < 7; fi28++ {
                v29 := GenS3{F0: 426}
                v30 := GenS2{F0: 100, F1: p0}
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            if true {
                p0 -= (^(p0))
            }
        }
        p0 = p0
    }
    switch 128 {
    case 128:
        fmt.Println(p0)
    case 1:
        if (float64((-2.25 / 1.0)) >= ((float64(3.14159) / ((float64(3.14159) - 0.0) + float64(3.14159))) - (0.0 / 0.1))) {
            var v31 int64 = -9223372036854775808
            for fi32 := 0; fi32 < 14; fi32++ {
                for fi33 := 0; fi33 < 19; fi33++ {
                    v31 += int64(fi32)
                    v34 := GenS1{F0: 89, F1: 2, F2: "\n\t"}
                    _ = fi33
                    _ = v34
                }
                for fi35 := 0; fi35 < 2; fi35++ {
                    fmt.Println(p0)
                    if (!(-32768 == -32768)) {
                        v36 := []int{fi35}
                        _ = v36
                    }
                    v37 := []int{-740, fi32, fi35}
                    _ = fi35
                    _ = v37
                }
                _ = fi32
            }
            _ = v31
        }
    case 32768:
        if false {
            var v38 uint32 = 65535
            if (18446744073709551615 <= p0) {
                for fi39 := 0; fi39 < 12; fi39++ {
                    v40 := []int{fi39, fi39, fi39}
                    _ = fi39
                    _ = v40
                }
            } else {
                switch 1 {
                case -1:
                    var v41 uint64 = p0
                    _ = v41
                case 128:
                    if ((0.0 - 0.1) <= float32((0.0 + (-2.25 - (1.5 + 1.5))))) {
                        v38 = (v38 << ((65535) & 63))
                        v42 := []int{42, 128, 32767}
                        for fi43 := 0; fi43 < 4; fi43++ {
                            v42 = append(v42, fi43)
                            if (true && true) {
                                v42 = append(v42, fi43)
                            } else {
                                v38 += v38
                            }
                            _ = fi43
                        }
                        _ = v42
                    }
                default:
                    v44 := []int{-9223372036854775808}
                    _ = v44
                }
                switch 4294967295 {
                case 42:
                    var v45 int32 = 42
                    _ = v45
                case -463:
                    var v46 bool = (true || ((255 == 42) || (0 != (v38 & (v38 | (v38 - v38))))))
                    _ = v46
                }
                var v47 int16 = 0
                _ = v47
            }
            _ = v38
        }
    }
    _ = p0
    return (p0 + 674)
}

func main() {
    var v1 int16 = 1000
    var v2 int8 = 1
    var v3 uint16 = 65535
    v4 := make(map[string]int)
    var v5 float32 = ((0.0 + ((3.14159 - -2.25) - 0.1)) - (0.1 + -2.25))
    var v6 uint32 = 0
    if ((980 < 583) && (-127 >= 400)) {
        v4["a"] = 128
        var v48 uint16 = v3
        _ = v48
    } else {
        fmt.Println(v1, v5)
        var v49 uint8 = 1
        _ = v49
    }
    fmt.Println(v5)
    v5 += float32(1.5)
    switch ("a" + ((("a" + "hello") + (((("hello" + (("hello" + "hello") + "\n\t")) + "\n\t") + ("" + "foo bar")) + "foo bar")) + ("foo bar" + ("hello" + "a")))) {
    case "\n\t":
        v50 := GenS1{F0: 1, F1: v3, F2: "\n\t"}
        _ = v50
    case "hello":
        delete(v4, (("foo bar" + "foo bar") + "a"))
    default:
        if false {
            for fi51 := 0; fi51 < 2; fi51++ {
                if (fi51 >= fi51) {
                    v52 := v4[("foo bar" + "foo bar")]
                    _ = v52
                }
                v1 = v1
                _ = fi51
            }
            for fi53 := 0; fi53 < 2; fi53++ {
                var v54 uint32 = (v6 >> ((255) & 63))
                v6 = 42
                _ = fi53
                _ = v54
            }
            switch 42 {
            case 32767:
                v1 -= (^(v1))
            case -2147483648:
                _ = len(v4)
            case 9223372036854775807:
                v55 := []int{651, 414, int(v3)}
                _ = v55
            default:
                var v56 bool = false
                _ = v56
            }
        } else {
            for fi57 := 0; fi57 < 11; fi57++ {
                if (42 <= -464) {
                    if (((1.5 - ((((1e10 * 0.1) / 0.1) / (1e10 / float64(((1e10 - (float32(1e10) / v5)) + v5)))) + 3.14159)) + 1e10) <= -2.25) {
                        switch "" {
                        case "\n\t":
                            v58 := GenS4{F0: 255}
                            _ = v58
                        default:
                            v59 := GenS5{F0: 2}
                            _ = v59
                        }
                        var v60 uint32 = v6
                        if (true || true) {
                            v61 := GenS4{F0: 1}
                            if (!(!true)) {
                                v4[""] = 127
                                v6 -= 65535
                                if false {
                                    v62 := []int{fi57, fi57, (fi57 / (((fi57 & 127)) | 1)), 61}
                                    _ = v62
                                }
                            } else {
                                v63 := GenS6{F0: 105, F1: 630, F2: 798}
                                v60 -= v6
                                var v64 int = fi57
                                _ = v63
                                _ = v64
                            }
                            switch (fi57 % (((fi57 >> ((1) & 63))) | 1)) {
                            case -992:
                                fmt.Println(fi57)
                            case -9223372036854775808:
                                fmt.Println(fi57, v6)
                            default:
                                v65 := GenS7{F0: (v1 | (v1 ^ -32768)), F1: 1, F2: (fi57 | (fi57 | (fi57 << ((433) & 63))))}
                                _ = v65
                            }
                            _ = v61
                        }
                        _ = v60
                    } else {
                        v6 = (v6 | v6)
                    }
                    v1 -= (v1 * (v1 >> ((0) & 63)))
                }
                var v66 int8 = (v2 % (((v2 / ((93) | 1))) | 1))
                for fi67 := 0; fi67 < 8; fi67++ {
                    v4[(("a" + "foo bar") + ("a" + ""))] = (fi67 * 32767)
                    v4[((("hello" + "") + "foo bar") + ("" + "\n\t"))] = (fi57 & fi67)
                    if true {
                        if (2 != 42) {
                            v66 -= 74
                        }
                        switch "hello" {
                        case "":
                            var v68 float32 = ((v5 / float32(3.14159)) / v5)
                            _ = v68
                        case "\n\t":
                            if (!false) {
                                v2 = v2
                            } else {
                                var v69 int64 = -9223372036854775808
                                var v70 uint64 = 1
                                v71 := []int{fi57, (fi67 * (fi57 / ((fi57) | 1)))}
                                _ = v69
                                _ = v70
                                _ = v71
                            }
                        case "foo bar":
                            var v72 uint32 = v6
                            _ = v72
                        }
                    }
                    _ = fi67
                }
                _ = fi57
                _ = v66
            }
            if (!false) {
                fmt.Println(v3, v1)
                for fi73 := 0; fi73 < 5; fi73++ {
                    for fi74 := 0; fi74 < 17; fi74++ {
                        var v75 int16 = -735
                        _ = fi74
                        _ = v75
                    }
                    _ = fi73
                }
            }
        }
    }
    v76 := v4[""]
    var v77 float64 = (float64(v5) * 0.1)
    var v78 int32 = -2147483648
    v5 = (3.14159 - v5)
    var v79 int32 = 0
    switch "" {
    case "\n\t":
        var v80 string = ("a" + "\n\t")
        _ = v80
    case "":
        var v81 float32 = (v5 * ((v5 + 1e10) / v5))
        _ = v81
    case "hello":
        for fi82 := 0; fi82 < 12; fi82++ {
            if (v6 > (v6 + 2)) {
                fmt.Println(v3, v6, v2)
            } else {
                v6 = 191
                v83 := v4["a"]
                _ = v83
            }
            _ = fi82
        }
    }
    v84 := v4["foo bar"]
    fmt.Println(v76, v1, v84)
    if false {
        v76 += (v84 % ((32768) | 1))
        v1 = v1
        v84 -= v84
    }
    switch "foo bar" {
    case "\n\t":
        v84 -= (-(v84))
    case "foo bar":
        fmt.Println(v77)
    default:
        v6 = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v84
}
