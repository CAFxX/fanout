package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS3 struct { F0 int64; F1 string; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 int; F1 uint64; F2 uint }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64; F1 int; F2 int16 }
type GenS12 struct { F0 uint64; F1 bool; F2 int16 }
func genFn0(p0 uint32, p1 int) uint8 {
    v5 := []int{p1}
    v6 := GenS1{F0: p1, F1: (87 == 1)}
    switch (("\n\t" + "hello") + "\n\t") {
    case "foo bar":
        var v7 bool = (!false)
        _ = v7
    case "":
        p0 = (p0 << ((2) & 63))
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int16 = -169
        v5 = append(v5, (fi8 / (((p1 >> ((0) & 63))) | 1)))
        fmt.Println(fi8, p0)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch p1 {
        case 1:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case -1:
            fmt.Println(p0, p0, p0)
        case -830:
            if (!(-9223372036854775808 == 0)) {
                switch 128 {
                case -1:
                    switch "a" {
                    case "\n\t":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "a":
                        var v11 uint64 = 2
                        _ = v11
                    default:
                        v12 := GenS2{F0: 253, F1: 0, F2: 42}
                        _ = v12
                    }
                case 32768:
                    if (1 < 18446744073709551615) {
                        var v13 uint = 42
                        for fi14 := 0; fi14 < 4; fi14++ {
                            if true {
                                var v15 float64 = ((1.5 / 1.0) / 1e10)
                                _ = v15
                            } else {
                                if (-581 < 482) {
                                    var v16 float64 = 1e10
                                    fmt.Println(p0, fi14)
                                    if len(v5) > 0 {
                                        v5[0] = (fi10 % ((fi14) | 1))
                                    }
                                    if (false || true) {
                                        v17 := GenS1{F0: 962, F1: (42 >= 0)}
                                        var v18 bool = false
                                        p1 = 2147483647
                                        _ = v17
                                        _ = v18
                                    } else {
                                        v13 = 65535
                                        var v19 int8 = -74
                                        _ = v19
                                    }
                                    _ = v16
                                } else {
                                    var v20 uint16 = 1
                                    v21 := GenS3{F0: 9223372036854775807, F1: "a", F2: 2}
                                    _ = v20
                                    _ = v21
                                }
                                var v22 int = (p1 & (p1 / ((fi10) | 1)))
                                _ = v22
                            }
                            if len(v5) > 0 {
                                v5[0] = fi14
                            }
                            fmt.Println(p1, p0, fi14)
                            _ = fi14
                        }
                        _ = v13
                    }
                case 506:
                    if (0 == 0) {
                        _ = len(v5)
                    }
                }
                _ = len(v5)
                for fi23 := 0; fi23 < 7; fi23++ {
                    p0 = p0
                    _ = fi23
                }
            } else {
                p1 = p1
                if false {
                    switch ("\n\t" + ("hello" + "")) {
                    case "\n\t":
                        for fi24 := 0; fi24 < 20; fi24++ {
                            _ = len(v5)
                            v25 := GenS1{F0: (fi24 & (p1 << ((42) & 63))), F1: true}
                            v26 := GenS1{F0: p1, F1: false}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    case "foo bar":
                        p0 = (p0 & (p0 >> ((1) & 63)))
                    case "hello":
                        p0 += (p0 >> ((0) & 63))
                    default:
                        if ((^(p0)) < (^(p0))) {
                            var v27 uint8 = 1
                            _ = v27
                        } else {
                            for fi28 := 0; fi28 < 15; fi28++ {
                                v29 := GenS4{F0: 0}
                                _ = fi28
                                _ = v29
                            }
                        }
                    }
                } else {
                    var v30 int8 = -1
                    switch fi10 {
                    case -448:
                        for fi31 := 0; fi31 < 1; fi31++ {
                            var v32 uint32 = (p0 - p0)
                            v33 := GenS3{F0: -903, F1: (("a" + "\n\t") + "hello"), F2: 246}
                            var v34 int8 = v30
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    }
                    _ = v30
                }
            }
        }
        p0 = (p0 ^ 1)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 158
}
func genFn1() uint64 {
    v35 := GenS5{F0: 1000}
    var v36 bool = (9223372036854775807 != 42)
    var v37 uint8 = 44
    _ = v35
    _ = v36
    _ = v37
    return 511
}

func main() {
    var v1 int64 = 42
    var v2 float64 = 0.1
    var v3 uint32 = 0
    var v4 float64 = 3.14159
    v38 := []int{}
    fmt.Println(v2, v2)
    for fi39 := 0; fi39 < 4; fi39++ {
        var v40 uint16 = 835
        for fi41 := 0; fi41 < 10; fi41++ {
            var v42 int = (^(fi39))
            var v43 uint64 = 2
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi39
        _ = v40
    }
    if false {
        v1 = v1
        for fi44 := 0; fi44 < 14; fi44++ {
            var v45 uint = 234
            _ = fi44
            _ = v45
        }
        if (1 <= 1) {
            v1 = v1
            if false {
                var v46 float32 = ((3.14159 * (0.0 - 3.14159)) * (((((0.1 - 0.0) / 1.5) - 1e10) + ((-2.25 / (0.1 * (((1e10 * 0.1) - (3.14159 - 1.5)) + 0.1))) - 1.5)) - (0.0 / (((float32(v4) + 1e10) / 0.1) - -2.25))))
                for fi47 := 0; fi47 < 1; fi47++ {
                    fmt.Println(v3, v4)
                    switch "hello" {
                    case "\n\t":
                        switch fi47 {
                        case -972:
                            for fi48 := 0; fi48 < 8; fi48++ {
                                var v49 uint8 = 240
                                _ = fi48
                                _ = v49
                            }
                        case 150:
                            _ = len(v38)
                        default:
                            if false {
                                switch 80 {
                                case 376:
                                    v50 := GenS6{F0: true, F1: "a"}
                                    _ = v50
                                case -192:
                                    switch "foo bar" {
                                    case "hello":
                                        v4 = 1.5
                                    case "a":
                                        if len(v38) > 0 {
                                            v38[0] = (fi47 * 4294967295)
                                        }
                                    case "\n\t":
                                        switch fi47 {
                                        case 385:
                                            var v51 uint32 = 548
                                            _ = v51
                                        case 128:
                                            v2 = (v2 - float64(((v46 - ((1.5 - (v46 + 3.14159)) * v46)) - v46)))
                                        case 2147483647:
                                            _ = len(v38)
                                        }
                                    default:
                                        if false {
                                            _ = len(v38)
                                        }
                                    }
                                case 32768:
                                    v38 = append(v38, 139)
                                }
                            } else {
                                v1 = (v1 / (((v1 * (v1 << ((0) & 63)))) | 1))
                                if true {
                                    v3 = (v3 << ((27) & 63))
                                } else {
                                    if len(v38) >= 2 {
                                        v38 = v38[1:len(v38)]
                                    }
                                }
                            }
                        }
                    case "a":
                        if false {
                            for fi52 := 0; fi52 < 15; fi52++ {
                                v2 += (v2 / ((v4 + (v4 + v2)) + v2))
                                var v53 int8 = 127
                                v1 -= v1
                                _ = fi52
                                _ = v53
                            }
                            for fi54 := 0; fi54 < 20; fi54++ {
                                fmt.Println(fi47, v1, fi47)
                                _ = fi54
                            }
                            for fi55 := 0; fi55 < 17; fi55++ {
                                var v56 bool = (-2147483648 <= 0)
                                if v56 {
                                    v57 := GenS7{F0: 0, F1: 42}
                                    switch 842 {
                                    case -587:
                                        if v56 {
                                            fmt.Println(v2, v56)
                                            v58 := GenS8{F0: (fi47 * 4294967295), F1: genFn1(), F2: 467}
                                            v59 := GenS6{F0: (!false), F1: "a"}
                                            _ = v58
                                            _ = v59
                                        } else {
                                            _ = len(v38)
                                            _ = len(v38)
                                        }
                                    default:
                                        var v60 uint8 = 42
                                        var v61 uint32 = 65535
                                        _ = v60
                                        _ = v61
                                    }
                                    if (100 == 37) {
                                        var v62 int16 = 32767
                                        _ = v62
                                    } else {
                                        var v63 string = ("foo bar" + "\n\t")
                                        switch fi47 {
                                        case -2147483648:
                                            var v64 bool = (false && (-141 != 0))
                                            _ = v64
                                        case 497:
                                            v4 -= (v4 * ((v4 + v4) * v4))
                                        }
                                        var v65 float32 = 0.1
                                        v4 = -2.25
                                        _ = v63
                                        _ = v65
                                    }
                                    _ = v57
                                } else {
                                    v66 := GenS9{F0: genFn0((v3 >> ((49) & 63)), fi55)}
                                    var v67 string = "foo bar"
                                    _ = v66
                                    _ = v67
                                }
                                _ = len(v38)
                                var v68 uint64 = 208
                                v69 := GenS10{F0: 127}
                                _ = fi55
                                _ = v56
                                _ = v68
                                _ = v69
                            }
                        }
                    default:
                        if len(v38) > 0 {
                            v38[0] = fi47
                        }
                    }
                    switch fi47 {
                    case 461:
                        var v70 uint32 = (v3 - v3)
                        _ = v70
                    default:
                        for fi71 := 0; fi71 < 16; fi71++ {
                            fmt.Println(fi71, v46)
                            var v72 float32 = (float32((((v2 * (1e10 + 0.0)) * v4) / (float64(v46) / 1.5))) + (float32(-2.25) + (v46 - float32(v4))))
                            _ = fi71
                            _ = v72
                        }
                    }
                    _ = fi47
                }
                fmt.Println(v1, v1)
                _ = v46
            } else {
                v2 += v4
                v3 -= (v3 | v3)
            }
            var v73 uint8 = 1
            _ = v73
        }
    }
    v2 -= v4
    if (!false) {
        v3 = (v3 & 0)
        switch 421 {
        case 1:
            var v74 int = 9223372036854775807
            _ = v74
        case -1:
            switch ("hello" + "\n\t") {
            case "a":
                var v75 uint32 = v3
                _ = v75
            case "":
                v76 := GenS11{F0: 4294967295, F1: 538, F2: 42}
                _ = v76
            case "hello":
                v3 += (^(v3))
            default:
                _ = len(v38)
            }
        }
        v1 += (v1 * -159)
    }
    var v77 string = ("" + ((("\n\t" + ("a" + "a")) + "\n\t") + (("foo bar" + "") + ("foo bar" + ("\n\t" + "foo bar")))))
    v2 = v2
    v78 := GenS12{F0: 18446744073709551615, F1: true, F2: 1}
    fmt.Println(v2, v2)
    var v79 uint64 = 4294967295
    var v80 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v77:", v77)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
