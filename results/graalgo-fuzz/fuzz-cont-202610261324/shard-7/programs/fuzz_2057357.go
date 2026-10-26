package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int8 }
type GenS2 struct { F0 int; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 string; F2 int32 }
type GenS6 struct { F0 int8; F1 uint; F2 uint32 }
type GenS7 struct { F0 bool; F1 uint64; F2 int16 }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 int8; F1 int; F2 bool }
func genFn0(p0 uint16) int16 {
    p0 -= (^(p0))
    p0 = 65535
    switch ("" + "foo bar") {
    case "":
        var v7 uint8 = 128
        _ = v7
    }
    if (true || false) {
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(fi8)
            p0 = 32767
            _ = fi8
        }
        if (!((p0 >> ((716) & 63)) == (p0 - (p0 | p0)))) {
            var v9 uint = 2
            if (!((((0.1 - 0.0) - ((1.5 - float64(3.14159)) * float64(0.0))) != 0.0) || (true && false))) {
                if ((!true) || (uint64(v9) >= 453)) {
                    switch int(v9) {
                    case -757:
                        var v10 int = 9223372036854775807
                        _ = v10
                    }
                    p0 = (p0 * uint16(v9))
                }
                var v11 int = 127
                _ = v11
            } else {
                if (!true) {
                    var v12 uint32 = 2147483648
                    v13 := GenS1{F0: uint32(p0), F1: (!(!true)), F2: 100}
                    _ = v12
                    _ = v13
                }
                fmt.Println(p0, p0)
            }
            p0 = p0
            _ = v9
        } else {
            fmt.Println(p0)
        }
    } else {
        var v14 uint64 = 18446744073709551615
        v15 := GenS1{F0: 223, F1: (1 <= 127), F2: 127}
        _ = v14
        _ = v15
    }
    _ = p0
    return 1000
}
func genFn1() uint32 {
    if ((true || (!true)) && true) {
        v16 := []int{}
        switch "foo bar" {
        case "\n\t":
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        default:
            switch "foo bar" {
            case "a":
                if true {
                    var v17 bool = (!(!(!true)))
                    _ = v17
                } else {
                    v18 := GenS2{F0: 9223372036854775807, F1: 0, F2: -28}
                    var v19 int8 = -128
                    v20 := GenS3{F0: 4294967295}
                    _ = v18
                    _ = v19
                    _ = v20
                }
            case "\n\t":
                var v21 int8 = 97
                _ = v21
            default:
                var v22 float32 = 0.1
                _ = v22
            }
        }
        if false {
            var v23 int16 = -1
            _ = v23
        } else {
            if false {
                _ = len(v16)
                var v24 int = 32767
                v24 -= v24
                _ = v24
            } else {
                _ = len(v16)
            }
            _ = len(v16)
        }
        _ = v16
    } else {
        for fi25 := 0; fi25 < 17; fi25++ {
            v26 := []int{}
            _ = fi25
            _ = v26
        }
    }
    if false {
        var v27 uint = 2
        v27 += v27
        switch -9223372036854775808 {
        case -344:
            fmt.Println(v27)
        case 421:
            v27 -= (^(v27))
        case 1:
            if true {
                if (0 >= 740) {
                    v28 := GenS3{F0: 0}
                    if (-1 != 172) {
                        v27 = v27
                        for fi29 := 0; fi29 < 2; fi29++ {
                            v27 += (^(v27))
                            _ = fi29
                        }
                    }
                    _ = v28
                }
                v27 = 2
                v30 := GenS3{F0: 872}
                _ = v30
            }
        }
        _ = v27
    } else {
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 float64 = ((0.1 + (3.14159 / 3.14159)) + 0.0)
            _ = fi31
            _ = v32
        }
    }
    var v33 int64 = 0
    _ = v33
    return 65535
}

func main() {
    var v1 int64 = 42
    v2 := []string{"\n\t", ("a" + ("\n\t" + (("foo bar" + "\n\t") + "")))}
    v3 := make(map[int]uint32)
    var v4 float32 = 0.0
    var v5 int64 = v1
    var v6 int32 = 1
    v4 = (((1e10 * (((v4 / (((v4 + 0.1) / v4) * v4)) - ((1e10 - v4) * 1.5)) * 1.5)) * float32(0.1)) + 1e10)
    switch -193 {
    case 0:
        switch "\n\t" {
        case "foo bar":
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 uint = 2
                _ = fi34
                _ = v35
            }
        case "hello":
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                v36 := GenS2{F0: 773, F1: 462, F2: -126}
                _ = v36
            case "a":
                _ = len(v2)
            case "hello":
                v37 := GenS4{F0: 81}
                _ = v37
            }
        case "\n\t":
            v5 -= -1
        default:
            v38 := GenS5{F0: 0, F1: ("hello" + "foo bar"), F2: v6}
            _ = v38
        }
    default:
        v39 := GenS1{F0: 402, F1: (("a" > "a") || true), F2: 1}
        _ = v39
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        v41 := v3[32768]
        v2 = append(v2, "\n\t")
        fmt.Println(v5)
        _ = fi40
        _ = v41
    }
    if true {
        switch "a" {
        case "":
            v42 := GenS6{F0: 26, F1: 255, F2: 42}
            _ = v42
        default:
            switch "a" {
            case "":
                if true {
                    v4 = v4
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v4, v1, v4)
                }
            case "a":
                switch 318 {
                case -9223372036854775808:
                    for fi43 := 0; fi43 < 18; fi43++ {
                        var v44 uint = 283
                        _ = fi43
                        _ = v44
                    }
                case 976:
                    fmt.Println(v4, v4, v4)
                }
            case "hello":
                var v45 uint8 = 255
                _ = v45
            }
        }
        for fi46 := 0; fi46 < 4; fi46++ {
            var v47 bool = false
            if true {
                var v48 int32 = (-(v6))
                var v49 int8 = 1
                _ = v48
                _ = v49
            }
            for fi50 := 0; fi50 < 15; fi50++ {
                var v51 uint16 = 32767
                v4 -= 0.0
                if len(v2) > 0 {
                    v2[0] = "a"
                }
                _ = fi50
                _ = v51
            }
            _ = fi46
            _ = v47
        }
        for fi52 := 0; fi52 < 12; fi52++ {
            for fi53 := 0; fi53 < 17; fi53++ {
                var v54 uint32 = 2
                var v55 int = (fi53 >> ((255) & 63))
                if true {
                    v1 = (v5 | v5)
                    delete(v3, v55)
                } else {
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                }
                _ = fi53
                _ = v54
                _ = v55
            }
            _ = fi52
        }
    } else {
        var v56 int64 = -1
        v57 := v3[-9223372036854775808]
        _ = v56
        _ = v57
    }
    var v58 uint64 = 0
    v2 = append(v2, "a")
    var v59 int32 = v6
    v60 := GenS5{F0: 1, F1: ("foo bar" + ""), F2: (v6 & (v59 ^ (v6 * -704)))}
    v6 = v6
    var v61 uint = 42
    v6 = 42
    if (("a" == "a") && (v61 != (v61 << ((131) & 63)))) {
        if false {
            v1 -= (v1 - (v5 + (v5 - v5)))
            v62 := GenS7{F0: ((false || (v61 == 1)) || true), F1: v58, F2: 42}
            if (4294967295 <= -681) {
                var v63 bool = (false || (float32(1e10) < 0.1))
                _ = v63
            } else {
                v64 := v3[121]
                var v65 float64 = -2.25
                _ = v64
                _ = v65
            }
            _ = v62
        } else {
            var v66 float32 = -2.25
            var v67 int64 = (^(v1))
            v68 := GenS8{F0: 9223372036854775807, F1: (v58 << ((621) & 63))}
            _ = v66
            _ = v67
            _ = v68
        }
    }
    fmt.Println(v61, v5, v6)
    v4 -= v4
    var v69 uint = (v61 & v61)
    if true {
        if false {
            switch ("a" + ("foo bar" + "\n\t")) {
            case "foo bar":
                var v70 float64 = 1e10
                _ = v70
            }
            v6 = (v6 >> (((v69 % ((747) | 1))) & 63))
            v69 = (^(v69))
        } else {
            if (32767 != 1000) {
                _ = len(v2)
            }
            for fi71 := 0; fi71 < 15; fi71++ {
                v61 += 42
                for fi72 := 0; fi72 < 6; fi72++ {
                    var v73 int = (^(fi71))
                    var v74 float32 = 3.14159
                    v58 = 42
                    v61 = 4
                    _ = fi72
                    _ = v73
                    _ = v74
                }
                _ = fi71
            }
        }
        if true {
            _ = len(v3)
        }
    } else {
        for fi75 := 0; fi75 < 1; fi75++ {
            var v76 uint = v61
            _ = fi75
            _ = v76
        }
        switch (("\n\t" + "foo bar") + "foo bar") {
        case "\n\t":
            v69 += (v61 | (v69 * (v61 * (^(v61)))))
        case "foo bar":
            for fi77 := 0; fi77 < 5; fi77++ {
                for fi78 := 0; fi78 < 17; fi78++ {
                    v79 := GenS1{F0: 65, F1: true, F2: 88}
                    v5 -= v1
                    _ = fi78
                    _ = v79
                }
                fmt.Println(v6, v5)
                _ = fi77
            }
        default:
            fmt.Println(v5, v58, v1)
        }
    }
    if (true || ((false || (!true)) && true)) {
        var v80 string = ("a" + "a")
        v4 = 1e10
        if ((genFn1() == 37) && true) {
            fmt.Println(v4)
            v6 += v59
        }
        _ = v80
    } else {
        v69 = v69
        if false {
            v1 = v5
            v81 := GenS7{F0: true, F1: (v58 + (v58 >> ((v61) & 63))), F2: genFn0(225)}
            _ = v81
        } else {
            v82 := GenS9{F0: 0, F1: 42, F2: true}
            _ = v82
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v69
}
