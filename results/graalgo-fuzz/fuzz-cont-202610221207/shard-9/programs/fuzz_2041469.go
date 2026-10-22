package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int64; F2 int }
type GenS4 struct { F0 bool; F1 uint32; F2 int8 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int; F1 int64; F2 string }
type GenS7 struct { F0 int; F1 int32; F2 int64 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int; F1 string; F2 int8 }
func genFn0(p0 int16, p1 int, p2 int64) int64 {
    for fi4 := 0; fi4 < 4; fi4++ {
        p2 += 871
        var v5 string = "\n\t"
        v6 := GenS1{F0: fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    if false {
        fmt.Println(p0, p1, p1)
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            fmt.Println(p0, p1)
            _ = fi7
        }
        v8 := GenS2{F0: 48, F1: 2}
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint = 928
        if false {
            if true {
                v11 := GenS1{F0: (^(p1))}
                p2 += int64(p0)
                var v12 uint = (v10 & 18446744073709551615)
                _ = v11
                _ = v12
            }
        }
        _ = fi9
        _ = v10
    }
    v13 := GenS2{F0: 628, F1: 65535}
    for fi14 := 0; fi14 < 1; fi14++ {
        switch "a" {
        case "\n\t":
            v15 := []int{}
            _ = v15
        case "hello":
            for fi16 := 0; fi16 < 4; fi16++ {
                if false {
                    v17 := []int{fi14, 254, p1}
                    var v18 float32 = ((float32((1e10 * 0.0)) - 0.0) - 3.14159)
                    var v19 float32 = v18
                    _ = v17
                    _ = v18
                    _ = v19
                }
                p1 = (fi14 & p1)
                _ = fi16
            }
        case "":
            switch "hello" {
            case "":
                for fi20 := 0; fi20 < 3; fi20++ {
                    p1 = (fi20 / ((fi14) | 1))
                    _ = fi20
                }
            case "hello":
                switch -2147483648 {
                case 32767:
                    if true {
                        var v21 uint32 = 4294967295
                        p0 += 32767
                        p0 -= (p0 + p0)
                        _ = v21
                    }
                case -9223372036854775808:
                    fmt.Println(p2, p2)
                }
            }
        default:
            v22 := []int{(fi14 + (p1 * 2147483647)), -833, (p1 | (p1 >> ((4294967295) & 63)))}
            _ = v22
        }
        p0 -= 1000
        var v23 uint8 = 86
        _ = fi14
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p2
}
func genFn1(p0 int32, p1 int8, p2 int32) uint64 {
    switch "\n\t" {
    case "foo bar":
        fmt.Println(p2, p2)
    case "hello":
        v24 := []int{}
        _ = v24
    case "\n\t":
        switch 166 {
        case 268:
            switch "foo bar" {
            case "":
                var v25 uint = 255
                _ = v25
            case "\n\t":
                v26 := GenS3{F0: 0, F1: genFn0(1, 32767, -1), F2: 127}
                _ = v26
            }
        }
    }
    v27 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 4294967295
    var v3 uint8 = 224
    var v28 float32 = -2.25
    v29 := GenS4{F0: ((v3 == 0) || (4294967295 == 4294967295)), F1: 4294967295, F2: -1}
    v3 += v3
    fmt.Println(v3, v28, v28)
    var v30 int = -855
    var v31 int = v30
    fmt.Println(v30, v3, v31)
    if false {
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 uint = 1
            var v34 int = v31
            var v35 float64 = 1e10
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        v31 -= 1
        var v36 float32 = v28
        _ = v36
    } else {
        switch v30 {
        case 2147483647:
            var v37 int16 = 42
            _ = v37
        case 4294967295:
            if ((v31 / ((v30) | 1)) < -53) {
                var v38 bool = ((1.5 - 1.5) < 3.14159)
                v39 := []int{}
                _ = v38
                _ = v39
            }
        default:
            if true {
                v31 = v30
                for fi40 := 0; fi40 < 4; fi40++ {
                    v31 -= -9223372036854775808
                    v41 := v1[376]
                    _ = fi40
                    _ = v41
                }
            }
        }
        v42 := GenS2{F0: uint16(v3), F1: 177}
        v30 += 32767
        _ = v42
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        if true {
            if true {
                switch "a" {
                case "a":
                    for fi44 := 0; fi44 < 18; fi44++ {
                        var v45 uint16 = 255
                        v31 = -26
                        v46 := []int{}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                default:
                    var v47 string = "a"
                    _ = v47
                }
                v48 := GenS5{F0: ("a" + ("a" + "")), F1: (false || (v3 != 255))}
                for fi49 := 0; fi49 < 16; fi49++ {
                    v2 = (^(v2))
                    v2 += (v2 - 853)
                    _ = fi49
                }
                _ = v48
            }
            if (-1 != 0) {
                v50 := GenS5{F0: "foo bar", F1: (genFn1(383, -1, -2147483648) >= 0)}
                delete(v1, 173)
                _ = v50
            } else {
                var v51 uint64 = v2
                if (-10 > 310) {
                    if (!true) {
                        v28 -= (3.14159 * float32(-2.25))
                        v30 += v31
                        if (true || (false && (!(32767 > 2)))) {
                            var v52 uint16 = 255
                            v53 := v1[955]
                            v54 := GenS6{F0: 721, F1: -9223372036854775808, F2: v53}
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                    }
                    for fi55 := 0; fi55 < 18; fi55++ {
                        fmt.Println(v30, v28, fi43)
                        fmt.Println(v3)
                        _ = fi55
                    }
                } else {
                    if (true || false) {
                        v56 := GenS2{F0: 32767, F1: 0}
                        _ = v56
                    }
                    switch v30 {
                    case 550:
                        var v57 int64 = -1
                        _ = v57
                    case -2147483648:
                        var v58 uint16 = 870
                        _ = v58
                    case 0:
                        var v59 string = "hello"
                        _ = v59
                    }
                }
                for fi60 := 0; fi60 < 17; fi60++ {
                    v61 := v1[2147483647]
                    var v62 int16 = 1
                    v2 += (^(v51))
                    var v63 uint8 = (^(v3))
                    _ = fi60
                    _ = v61
                    _ = v62
                    _ = v63
                }
                _ = v51
            }
        }
        for fi64 := 0; fi64 < 1; fi64++ {
            v3 = (v3 % ((v3) | 1))
            v2 -= (v2 & (v2 ^ 0))
            if (4294967295 > 1) {
                v65 := GenS7{F0: fi64, F1: 738, F2: 1}
                var v66 int16 = -614
                _ = v65
                _ = v66
            }
            fmt.Println(fi64, v31, v30)
            _ = fi64
        }
        v67 := GenS8{F0: -1, F1: v30}
        _ = fi43
        _ = v67
    }
    var v68 uint32 = 42
    for fi69 := 0; fi69 < 16; fi69++ {
        for fi70 := 0; fi70 < 3; fi70++ {
            v68 += (v68 / ((v68) | 1))
            fmt.Println(v31, fi70, v30)
            v30 = int(v68)
            var v71 uint32 = v68
            _ = fi70
            _ = v71
        }
        _ = fi69
    }
    if (!(0.1 == 0.1)) {
        for fi72 := 0; fi72 < 10; fi72++ {
            v31 = (^(v30))
            var v73 int32 = -722
            _ = fi72
            _ = v73
        }
        v31 = (v31 / ((1) | 1))
    }
    v74 := v1[uint32(v2)]
    if false {
        v75 := GenS9{F0: -2147483648, F1: v74, F2: -1}
        _ = v75
    } else {
        switch v74 {
        case "foo bar":
            v74 = "hello"
        case "a":
            v74 = v74
        case "\n\t":
            var v76 string = v74
            _ = v76
        default:
            if false {
                v30 = -1
                switch "foo bar" {
                case "":
                    for fi77 := 0; fi77 < 4; fi77++ {
                        v3 = (v3 / (((^(v3))) | 1))
                        fmt.Println(v2, v2)
                        _ = fi77
                    }
                case "a":
                    var v78 int64 = genFn0(0, v31, 1)
                    _ = v78
                case "\n\t":
                    for fi79 := 0; fi79 < 20; fi79++ {
                        var v80 uint32 = v68
                        if true {
                            var v81 string = (((v74 + "foo bar") + v74) + v74)
                            _ = v81
                        } else {
                            var v82 uint8 = (v3 >> ((1) & 63))
                            v80 = (v68 + (v80 / (((v68 - (v80 * 65535))) | 1)))
                            _ = v82
                        }
                        _ = fi79
                        _ = v80
                    }
                }
                var v83 int32 = 310
                _ = v83
            }
        }
        _ = len(v1)
    }
    for fi84 := 0; fi84 < 17; fi84++ {
        var v85 uint32 = (^(v68))
        var v86 uint = 42
        v87 := []int{127, -849, v30, (v31 / ((v30) | 1))}
        _ = fi84
        _ = v85
        _ = v86
        _ = v87
    }
    for fi88 := 0; fi88 < 12; fi88++ {
        switch "" {
        case "foo bar":
            fmt.Println(fi88)
        case "\n\t":
            if false {
                var v89 uint8 = (v3 - (v3 - 128))
                v90 := []int{(fi88 << ((150) & 63)), v31}
                _ = v89
                _ = v90
            }
        default:
            var v91 uint = 4294967295
            _ = v91
        }
        _ = fi88
    }
    if true {
        var v92 string = "hello"
        _ = v92
    } else {
        if (false || ((((!((!((v68 >> ((2) & 63)) != v68)) || ((v2 != 851) || (1 >= 111)))) && true) && (v3 == (v3 / ((127) | 1)))) || (!((v3 + 127) < v3)))) {
            if (-32768 <= 0) {
                v68 -= v68
                var v93 int = (v31 + 128)
                _ = v93
            }
            fmt.Println(v31, v2, v2)
            if (!((((3.14159 - (0.0 * (1.5 - (float64(0.0) + 1e10)))) / (3.14159 * float64(v28))) + ((0.0 - 1e10) - 3.14159)) != 3.14159)) {
                v94 := GenS8{F0: 42, F1: 32767}
                fmt.Println(v74, v31)
                _ = v94
            }
        } else {
            switch "foo bar" {
            case "a":
                v30 = v30
            }
        }
        var v95 int8 = 38
        _ = v95
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v68:", v68)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v68
    _ = v74
}
