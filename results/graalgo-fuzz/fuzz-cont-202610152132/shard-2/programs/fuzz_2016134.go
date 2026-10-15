package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 int64; F1 int; F2 uint32 }
type GenS9 struct { F0 int; F1 int32 }
func genFn0(p0 uint) uint {
    if (1 < 9223372036854775807) {
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := GenS1{F0: "a"}
            _ = fi6
            _ = v7
        }
        p0 += (p0 % ((p0) | 1))
    } else {
        p0 = (p0 << (((p0 | (p0 >> (((^(p0))) & 63)))) & 63))
        switch ("a" + ("\n\t" + "foo bar")) {
        case "hello":
            v8 := []int{145, 0, 32767, -9223372036854775808}
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := []int{fi9}
                var v11 uint8 = 42
                _ = fi9
                _ = v10
                _ = v11
            }
        default:
            v12 := []int{912, -2147483648, 9223372036854775807}
            _ = v12
        }
        if (!true) {
            v13 := GenS1{F0: "hello"}
            switch ("\n\t" + ("foo bar" + "a")) {
            case "hello":
                fmt.Println(p0)
            case "foo bar":
                if (true && false) {
                    p0 += 255
                    if (32767 != -1) {
                        p0 = (p0 + (^(p0)))
                        var v14 uint16 = 255
                        _ = v14
                    }
                }
            case "":
                var v15 uint16 = 0
                _ = v15
            default:
                p0 += p0
            }
            var v16 float32 = -2.25
            _ = v13
            _ = v16
        }
    }
    fmt.Println(p0)
    v17 := []int{128, 4294967295}
    if (42 <= 681) {
        v18 := GenS2{F0: 42}
        _ = v18
    } else {
        if (1e10 >= float32(float64(3.14159))) {
            p0 -= (p0 + (p0 >> (((p0 * p0)) & 63)))
            v19 := GenS2{F0: 450}
            _ = v19
        } else {
            var v20 bool = false
            _ = v20
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = v17
    return 2
}
func genFn1() string {
    var v21 uint64 = 376
    var v22 uint16 = 2
    for fi23 := 0; fi23 < 1; fi23++ {
        if false {
            for fi24 := 0; fi24 < 20; fi24++ {
                switch "foo bar" {
                case "\n\t":
                    switch 0 {
                    case 9223372036854775807:
                        if (579 >= genFn0(843)) {
                            v21 = (v21 % (((v21 * v21)) | 1))
                        } else {
                            fmt.Println(fi24, fi23, v22)
                        }
                    }
                case "a":
                    v22 = v22
                case "":
                    var v25 float32 = 0.0
                    _ = v25
                default:
                    v26 := []int{(^(fi24))}
                    _ = v26
                }
                switch "foo bar" {
                case "hello":
                    if ((^(fi23)) != 1) {
                        v27 := GenS3{F0: "\n\t"}
                        _ = v27
                    } else {
                        v22 = v22
                        v28 := GenS4{F0: -1}
                        _ = v28
                    }
                    v29 := GenS4{F0: -128}
                    _ = v29
                case "\n\t":
                    v30 := []int{}
                    _ = v30
                case "a":
                    switch "" {
                    case "":
                        fmt.Println(fi23, v22)
                    default:
                        v31 := GenS4{F0: -93}
                        _ = v31
                    }
                default:
                    var v32 float32 = (1e10 * float32((0.1 + 0.1)))
                    _ = v32
                }
                v21 = (v21 & v21)
                _ = fi24
            }
            if false {
                for fi33 := 0; fi33 < 13; fi33++ {
                    var v34 int = (fi23 | fi33)
                    v21 = (v21 - 1)
                    if (!(-745 == 2147483647)) {
                        switch "hello" {
                        case "\n\t":
                            var v35 int8 = -116
                            _ = v35
                        case "":
                            fmt.Println(v34, fi23, v21)
                        case "a":
                            fmt.Println(fi23, fi23)
                        }
                        switch "a" {
                        case "a":
                            var v36 bool = true
                            _ = v36
                        }
                    }
                    _ = fi33
                    _ = v34
                }
                v37 := GenS3{F0: (("\n\t" + (("hello" + ("foo bar" + ("foo bar" + "hello"))) + "hello")) + "a")}
                v38 := []int{fi23, 0}
                _ = v37
                _ = v38
            } else {
                if ((true || false) && true) {
                    if true {
                        v39 := GenS4{F0: -123}
                        _ = v39
                    } else {
                        switch -9223372036854775808 {
                        case -9223372036854775808:
                            var v40 bool = true
                            _ = v40
                        }
                        v41 := []int{fi23}
                        v42 := GenS5{F0: "a", F1: true}
                        _ = v41
                        _ = v42
                    }
                    v43 := GenS6{F0: -128}
                    _ = v43
                } else {
                    for fi44 := 0; fi44 < 11; fi44++ {
                        v45 := GenS6{F0: int8(v22)}
                        if true {
                            var v46 int = -9223372036854775808
                            _ = v46
                        } else {
                            var v47 float32 = (-2.25 * -2.25)
                            v47 += (3.14159 / v47)
                            fmt.Println(fi44, fi44, v47)
                            v47 = (v47 - float32(3.14159))
                            _ = v47
                        }
                        _ = fi44
                        _ = v45
                    }
                }
            }
            switch "hello" {
            case "hello":
                var v48 string = ""
                _ = v48
            case "\n\t":
                v49 := []int{-684, 684, fi23, (fi23 / ((fi23) | 1))}
                _ = v49
            case "a":
                fmt.Println(fi23)
            default:
                v22 = v22
            }
        } else {
            var v50 int = fi23
            switch (fi23 - (v50 / (((fi23 * (v50 + 32768))) | 1))) {
            case 32768:
                v51 := GenS7{F0: ("\n\t" + "a"), F1: 42}
                _ = v51
            case 642:
                var v52 uint32 = 617
                _ = v52
            case 4294967295:
                if false {
                    var v53 uint64 = v21
                    _ = v53
                } else {
                    var v54 uint32 = 624
                    v55 := GenS7{F0: "foo bar", F1: -259}
                    v56 := []int{(^(v50)), fi23, v50}
                    _ = v54
                    _ = v55
                    _ = v56
                }
            default:
                if false {
                    var v57 float32 = -2.25
                    v57 += 3.14159
                    v21 = v21
                    _ = v57
                } else {
                    var v58 uint8 = 1
                    _ = v58
                }
            }
            _ = v50
        }
        v59 := []int{}
        if len(v59) >= 2 {
            v59 = v59[1:len(v59)]
        }
        _ = fi23
        _ = v59
    }
    v60 := []int{128, 127, -2147483648}
    if (!(true || (2 < 2))) {
        var v61 uint64 = 250
        _ = v61
    }
    _ = v21
    _ = v22
    _ = v60
    return "\n\t"
}

func main() {
    var v1 uint32 = 0
    var v2 int16 = 0
    var v3 string = ("" + ("foo bar" + "\n\t"))
    v4 := make(map[uint32]string)
    v5 := make(map[string]int)
    v2 -= (v2 ^ v2)
    var v62 int16 = v2
    var v63 float64 = float64((float32((1e10 / float64(1e10))) + -2.25))
    for fi64 := 0; fi64 < 14; fi64++ {
        for fi65 := 0; fi65 < 11; fi65++ {
            switch (v3 + v3) {
            case "":
                v66 := GenS8{F0: -40, F1: fi64, F2: 0}
                _ = v66
            case "\n\t":
                v1 -= (v1 * v1)
            }
            _ = fi65
        }
        if (v2 < (v2 >> ((42) & 63))) {
            switch v3 {
            case "hello":
                var v67 int8 = -1
                _ = v67
            default:
                var v68 uint32 = 911
                _ = v68
            }
        } else {
            v1 = v1
        }
        for fi69 := 0; fi69 < 20; fi69++ {
            v62 += 42
            fmt.Println(v63, v2, fi69)
            _ = fi69
        }
        _ = fi64
    }
    for fi70 := 0; fi70 < 6; fi70++ {
        fmt.Println(fi70)
        v71 := GenS9{F0: -556, F1: 1}
        _ = fi70
        _ = v71
    }
    for fi72 := 0; fi72 < 5; fi72++ {
        v3 = ("\n\t" + v3)
        var v73 int16 = -567
        _ = fi72
        _ = v73
    }
    fmt.Println(v1, v62)
    var v74 uint16 = 2
    var v75 int = -1
    switch (("a" + v3) + "") {
    case "foo bar":
        var v76 uint8 = 42
        _ = v76
    case "":
        var v77 uint32 = (v1 / (((v1 * 2147483648)) | 1))
        _ = v77
    }
    var v78 int16 = (^(v62))
    v79 := GenS4{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
    _ = v74
    _ = v75
    _ = v78
    _ = v79
}
