package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int16; F1 int64; F2 uint8 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 int8; F2 uint }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint64) bool {
    p1 -= 842
    for fi6 := 0; fi6 < 4; fi6++ {
        p1 += (p1 / ((18446744073709551615) | 1))
        _ = fi6
    }
    var v7 int64 = 9223372036854775807
    v8 := GenS1{F0: 42, F1: -811, F2: -2147483648}
    var v9 float32 = float32(1.5)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 int) string {
    v10 := []int{p0}
    if false {
        for fi11 := 0; fi11 < 10; fi11++ {
            p0 += (p0 * 42)
            _ = len(v10)
            _ = fi11
        }
    } else {
        p0 -= 9223372036854775807
    }
    if len(v10) > 0 {
        v10[0] = p0
    }
    var v12 bool = false
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 float32 = 0.0
        if ((!v12) || (!v12)) {
            switch p0 {
            case 2147483647:
                p0 = (^(fi13))
            case 514:
                if (!((333 > 0) && (("hello" + ("foo bar" + "")) == "\n\t"))) {
                    v14 = v14
                }
            default:
                v15 := GenS1{F0: 6, F1: -1, F2: 0}
                _ = v15
            }
        }
        var v16 string = (((("\n\t" + "\n\t") + "a") + "foo bar") + ("\n\t" + "\n\t"))
        _ = fi13
        _ = v14
        _ = v16
    }
    _ = p0
    _ = v10
    _ = v12
    return "foo bar"
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 32767
    var v3 float64 = (1e10 + (1.5 * float64(0.1)))
    v4 := []int{4294967295, 9223372036854775807, -144}
    var v5 uint = 0
    var v17 int = 9223372036854775807
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            switch (fi18 / ((fi18) | 1)) {
            case 128:
                v3 -= ((v3 / v3) + 1.5)
            case 2147483647:
                v20 := GenS1{F0: 150, F1: 1, F2: -945}
                _ = v20
            default:
                var v21 int32 = 1
                var v22 string = "hello"
                fmt.Println(fi18, fi18)
                _ = v21
                _ = v22
            }
            _ = fi19
        }
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 int = (fi18 + -866)
            switch (fi18 & fi18) {
            case 0:
                fmt.Println(fi23, v3)
            case 9223372036854775807:
                switch "\n\t" {
                case "":
                    v25 := GenS1{F0: 45, F1: -178, F2: 1}
                    _ = v25
                case "\n\t":
                    v26 := GenS1{F0: 128, F1: 889, F2: -545}
                    _ = v26
                case "hello":
                    var v27 uint16 = (^(v2))
                    _ = v27
                }
            default:
                v28 := v1[fi23]
                _ = v28
            }
            if (42 >= 127) {
                switch fi18 {
                case -2147483648:
                    v29 := GenS1{F0: 255, F1: 42, F2: 2147483647}
                    _ = v29
                case -9223372036854775808:
                    delete(v1, v17)
                }
                var v30 uint = 4294967295
                _ = v30
            }
            _ = fi23
            _ = v24
        }
        _ = fi18
    }
    if true {
        _ = len(v4)
    }
    if ((v17 - v17) == (v17 - 671)) {
        v31 := GenS2{F0: (v2 << (((v5 - (v5 / ((v5) | 1)))) & 63)), F1: 242}
        delete(v1, (v17 >> (((v5 & v5)) & 63)))
        var v32 uint8 = 127
        _ = v31
        _ = v32
    }
    v3 += v3
    for fi33 := 0; fi33 < 3; fi33++ {
        _ = len(v4)
        _ = fi33
    }
    if false {
        var v34 int32 = -510
        var v35 int64 = 42
        var v36 int = v17
        _ = v34
        _ = v35
        _ = v36
    } else {
        v37 := GenS3{F0: -205, F1: -9223372036854775808, F2: 0}
        _ = v37
    }
    var v38 float64 = float64((0.1 * 0.0))
    _ = len(v4)
    v5 = (v5 + v5)
    v38 = (1e10 * v38)
    for fi39 := 0; fi39 < 13; fi39++ {
        for fi40 := 0; fi40 < 13; fi40++ {
            v38 += -2.25
            var v41 int8 = 42
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    v3 -= (0.1 - v38)
    switch (("\n\t" + "a") + "") {
    case "foo bar":
        switch (v17 % (((v17 * 32767)) | 1)) {
        case 0:
            fmt.Println(v17, v5)
        case -405:
            switch (v17 / ((v17) | 1)) {
            case 2147483647:
                if ("a" == ("\n\t" + "")) {
                    v42 := GenS1{F0: 140, F1: -2147483648, F2: 42}
                    _ = v42
                }
            case 147:
                if (1e10 <= 0.0) {
                    v43 := GenS4{F0: (^(v17)), F1: -1}
                    for fi44 := 0; fi44 < 7; fi44++ {
                        v5 = v5
                        _ = fi44
                    }
                    _ = v43
                } else {
                    var v45 int64 = 0
                    var v46 int = (v17 ^ v17)
                    _ = v45
                    _ = v46
                }
            case 32767:
                switch "foo bar" {
                case "a":
                    var v47 int = (v17 * v17)
                    _ = v47
                case "\n\t":
                    v48 := v1[-235]
                    _ = v48
                case "":
                    for fi49 := 0; fi49 < 2; fi49++ {
                        v50 := GenS1{F0: 177, F1: -1, F2: 0}
                        var v51 int = 32767
                        v2 += (v2 - 997)
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                }
            default:
                var v52 int = v17
                _ = v52
            }
        case -1:
            fmt.Println(v2, v38, v17)
        }
    case "a":
        var v53 uint8 = 81
        _ = v53
    }
    switch v17 {
    case -1:
        v54 := GenS4{F0: v17, F1: 560}
        _ = v54
    case -459:
        v17 = (v17 * v17)
    default:
        if len(v4) > 0 {
            v4[0] = (v17 * (v17 - 32767))
        }
    }
    for fi55 := 0; fi55 < 7; fi55++ {
        v3 += v38
        for fi56 := 0; fi56 < 20; fi56++ {
            v57 := GenS5{F0: 42}
            _ = fi56
            _ = v57
        }
        _ = fi55
    }
    if ("foo bar" >= "") {
        if true {
            v5 = 395
            var v58 int8 = 127
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v58
        } else {
            for fi59 := 0; fi59 < 6; fi59++ {
                var v60 uint64 = 18446744073709551615
                _ = fi59
                _ = v60
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        v2 = 65535
        v61 := GenS6{F0: 4294967295, F1: 44, F2: v5}
        _ = v61
    } else {
        var v62 float32 = -2.25
        v63 := GenS4{F0: v17, F1: -1}
        switch v17 {
        case 32768:
            for fi64 := 0; fi64 < 2; fi64++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v5 = 42
                _ = fi64
            }
        }
        _ = v62
        _ = v63
    }
    for fi65 := 0; fi65 < 7; fi65++ {
        v66 := GenS4{F0: v17, F1: -9223372036854775808}
        switch "" {
        case "\n\t":
            v67 := GenS7{F0: ((float64(1e10) * v38) >= v38), F1: true}
            _ = v67
        case "hello":
            var v68 float64 = (v3 / (v38 - v38))
            _ = v68
        case "foo bar":
            switch "foo bar" {
            case "a":
                switch 4294967295 {
                case 127:
                    var v69 uint32 = 42
                    _ = v69
                case -972:
                    var v70 uint16 = (v2 | v2)
                    _ = v70
                }
            case "foo bar":
                for fi71 := 0; fi71 < 2; fi71++ {
                    var v72 bool = true
                    _ = fi71
                    _ = v72
                }
            case "\n\t":
                v73 := GenS3{F0: 32767, F1: 9223372036854775807, F2: 2}
                _ = v73
            default:
                switch (-(v17)) {
                case 9223372036854775807:
                    fmt.Println(fi65, v17)
                case 893:
                    v74 := GenS8{F0: 0}
                    _ = v74
                case 127:
                    v75 := GenS5{F0: 289}
                    _ = v75
                }
            }
        }
        _ = fi65
        _ = v66
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
