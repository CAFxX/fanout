package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint16; F1 int64; F2 int }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS8 struct { F0 int; F1 int32; F2 int32 }
type GenS9 struct { F0 uint16 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            fmt.Println(fi4)
            fmt.Println(fi4, fi4)
            _ = fi5
        }
        v6 := GenS1{F0: "a", F1: 0}
        var v7 uint8 = 203
        _ = fi4
        _ = v6
        _ = v7
    }
    if (float64(((((0.1 + (((-2.25 - 0.0) / 1.5) / 1e10)) / ((float32(3.14159) + 3.14159) + ((1.5 - 0.0) * -2.25))) + (float32(3.14159) + 1e10)) + (((1.5 / 1.0) + 1e10) / 1e10))) < (1e10 * 0.1)) {
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := []int{(fi8 - 358)}
            _ = len(v9)
            _ = fi8
            _ = v9
        }
    } else {
        if true {
            v10 := []int{}
            _ = v10
        } else {
            if true {
                v11 := GenS1{F0: ("hello" + "hello"), F1: 0}
                _ = v11
            }
            var v12 uint16 = 163
            _ = v12
        }
        v13 := GenS1{F0: "\n\t", F1: 0}
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint16 = 0
        _ = fi14
        _ = v15
    }
    switch "a" {
    case "hello":
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 uint64 = 460
            v17 -= v17
            for fi18 := 0; fi18 < 3; fi18++ {
                if false {
                    var v19 int = -2147483648
                    _ = v19
                } else {
                    v20 := GenS2{F0: 1, F1: v17, F2: (v17 << ((400) & 63))}
                    v17 = (v17 / ((571) | 1))
                    _ = v20
                }
                v21 := GenS1{F0: "foo bar", F1: 29}
                v22 := []int{-532, fi18, (fi18 << ((339) & 63)), 9223372036854775807}
                _ = fi18
                _ = v21
                _ = v22
            }
            _ = fi16
            _ = v17
        }
    case "a":
        var v23 uint = 255
        _ = v23
    case "foo bar":
        v24 := GenS1{F0: "", F1: 127}
        _ = v24
    }
    var v25 string = ("foo bar" + "foo bar")
    _ = v25
    return 2
}
func genFn1(p0 uint32, p1 int16) int16 {
    p0 -= 2
    for fi26 := 0; fi26 < 13; fi26++ {
        for fi27 := 0; fi27 < 20; fi27++ {
            p1 += p1
            _ = fi27
        }
        if (!(!(-1 >= 186))) {
            if (0 > 65535) {
                p1 = (p1 | p1)
                if true {
                    for fi28 := 0; fi28 < 1; fi28++ {
                        switch -1 {
                        case 9223372036854775807:
                            var v29 uint8 = 127
                            if ((1e10 * 1.5) < 3.14159) {
                                v29 -= v29
                            }
                            _ = v29
                        case 42:
                            var v30 uint64 = 18446744073709551615
                            _ = v30
                        default:
                            v31 := GenS2{F0: (^(p1)), F1: 848, F2: 1}
                            _ = v31
                        }
                        p0 -= (p0 ^ 1)
                        _ = fi28
                    }
                    fmt.Println(p0)
                    v32 := GenS2{F0: p1, F1: 4294967295, F2: 463}
                    _ = v32
                } else {
                    v33 := []int{(-(fi26))}
                    v34 := GenS1{F0: "hello", F1: 41}
                    _ = v33
                    _ = v34
                }
                for fi35 := 0; fi35 < 1; fi35++ {
                    v36 := GenS3{F0: "foo bar", F1: p1}
                    if (true && false) {
                        var v37 int16 = (p1 - (p1 >> ((883) & 63)))
                        _ = v37
                    }
                    _ = fi35
                    _ = v36
                }
            } else {
                var v38 uint = 611
                v39 := []int{(fi26 / ((127) | 1)), -283, -2147483648, fi26}
                _ = v38
                _ = v39
            }
            for fi40 := 0; fi40 < 16; fi40++ {
                var v41 int8 = -128
                v42 := GenS4{F0: 2, F1: int64(v41), F2: fi40}
                _ = fi40
                _ = v41
                _ = v42
            }
        } else {
            if true {
                var v43 int16 = p1
                _ = v43
            }
        }
        for fi44 := 0; fi44 < 17; fi44++ {
            if (true && ((-1 < 66) || false)) {
                switch ("a" + "\n\t") {
                case "a":
                    switch "hello" {
                    case "foo bar":
                        v45 := []int{(-(fi26)), (fi44 / ((fi26) | 1))}
                        _ = v45
                    case "":
                        var v46 uint16 = 0
                        _ = v46
                    case "a":
                        switch "\n\t" {
                        case "hello":
                            v47 := GenS4{F0: 42, F1: 0, F2: fi26}
                            _ = v47
                        case "\n\t":
                            v48 := []int{(fi44 % ((fi26) | 1)), fi26}
                            _ = v48
                        default:
                            p0 += p0
                        }
                        var v49 float32 = -2.25
                        v50 := GenS2{F0: p1, F1: 18446744073709551615, F2: 1}
                        var v51 string = "a"
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                case "\n\t":
                    p0 -= 0
                case "":
                    v52 := []int{fi44, (fi26 % (((fi44 + -9223372036854775808)) | 1)), fi44}
                    _ = v52
                }
                switch ((("a" + "\n\t") + "\n\t") + "") {
                case "a":
                    if (true && ((-2.25 * (1e10 - (0.1 + float32(1.5)))) < float32(-2.25))) {
                        v53 := GenS5{F0: 0, F1: 1}
                        _ = v53
                    } else {
                        var v54 int32 = 801
                        _ = v54
                    }
                case "foo bar":
                    var v55 uint32 = p0
                    _ = v55
                case "hello":
                    fmt.Println(p1, p0, fi44)
                }
                fmt.Println(fi26)
            }
            p0 = p0
            var v56 uint = 18446744073709551615
            v56 -= v56
            if (float64((1e10 + 0.1)) != 1.5) {
                v57 := GenS6{F0: (("a" + "\n\t") + ("" + "a")), F1: 1}
                _ = v57
            } else {
                v58 := []int{}
                _ = v58
            }
            _ = fi44
            _ = v56
        }
        _ = fi26
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 uint = 65535
    var v2 uint16 = 63
    var v3 int8 = 41
    var v59 float64 = 3.14159
    fmt.Println(v2)
    switch ("foo bar" + "") {
    case "hello":
        v59 += v59
    case "\n\t":
        v60 := []int{32768}
        _ = v60
    case "":
        var v61 int = 128
        _ = v61
    default:
        var v62 float64 = (float64((float32((v59 + (v59 * -2.25))) - -2.25)) + v59)
        _ = v62
    }
    var v63 int16 = 0
    var v64 string = "foo bar"
    var v65 float32 = 0.1
    v3 -= 0
    v3 += v3
    v66 := GenS7{F0: v2, F1: 2, F2: 2}
    for fi67 := 0; fi67 < 7; fi67++ {
        if (true && (-617 > 934)) {
            var v68 float32 = v65
            var v69 float32 = v68
            _ = v68
            _ = v69
        }
        switch v64 {
        case "foo bar":
            switch 4294967295 {
            case -809:
                v70 := []int{}
                _ = v70
            case 32768:
                v71 := GenS7{F0: (v2 ^ uint16(v1)), F1: v1, F2: 483}
                _ = v71
            case -846:
                v72 := GenS1{F0: v64, F1: (v3 % ((v3) | 1))}
                _ = v72
            default:
                switch (v64 + "a") {
                case "hello":
                    for fi73 := 0; fi73 < 9; fi73++ {
                        v74 := GenS8{F0: 9223372036854775807, F1: 307, F2: 2147483647}
                        _ = fi73
                        _ = v74
                    }
                case "\n\t":
                    v75 := GenS9{F0: (v2 % (((v2 - 32767)) | 1))}
                    _ = v75
                }
            }
        }
        _ = fi67
    }
    var v76 uint = v1
    var v77 string = v64
    var v78 uint = 18446744073709551615
    if (v2 == 42) {
        v78 += v76
        fmt.Println(v77, v76, v77)
        if (v77 < v64) {
            fmt.Println(v76)
            var v79 int16 = v63
            fmt.Println(v2)
            _ = v79
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v59:", v59)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v59
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v76
    _ = v77
    _ = v78
}
