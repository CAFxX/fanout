package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int32; F1 uint32; F2 uint }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int64; F1 int8; F2 int64 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 int; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 int8; F1 uint32 }
func genFn0(p0 bool, p1 uint) uint64 {
    switch "a" {
    case "foo bar":
        p0 = (100 != 1)
    }
    if p0 {
        if (1.5 == 3.14159) {
            if (!false) {
                v6 := []int{}
                _ = v6
            } else {
                for fi7 := 0; fi7 < 16; fi7++ {
                    if p0 {
                        var v8 int32 = 1
                        var v9 string = ("hello" + "foo bar")
                        _ = v8
                        _ = v9
                    }
                    p0 = false
                    for fi10 := 0; fi10 < 8; fi10++ {
                        p1 -= 65535
                        v11 := []int{(fi10 / (((fi7 % ((fi7) | 1))) | 1)), (fi7 >> ((255) & 63)), (-(fi10)), (fi10 + fi7)}
                        v12 := GenS1{F0: true, F1: ("" + "hello")}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi7
                }
                p0 = false
            }
        } else {
            p1 = 65535
            if true {
                var v13 uint16 = 0
                _ = v13
            } else {
                if p0 {
                    v14 := GenS1{F0: (float32(0.1) != 1e10), F1: ""}
                    _ = v14
                }
                if true {
                    var v15 uint = (p1 >> (((p1 / ((1) | 1))) & 63))
                    v16 := GenS1{F0: (!true), F1: ""}
                    _ = v15
                    _ = v16
                } else {
                    switch 126 {
                    case 32768:
                        var v17 uint32 = 65535
                        _ = v17
                    }
                    v18 := []int{0, 4294967295}
                    _ = v18
                }
                p1 = (^(p1))
            }
        }
        v19 := GenS1{F0: (238 < 246), F1: ""}
        _ = v19
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint64) int32 {
    var v20 bool = (255 <= 150)
    var v21 string = "hello"
    v20 = ((42 <= 0) || true)
    _ = p0
    _ = v20
    _ = v21
    return 1
}

func main() {
    var v1 int64 = -877
    v2 := make(map[string]uint32)
    var v3 int = -1
    var v4 int = v3
    var v5 int32 = 42
    fmt.Println(v3, v3, v1)
    fmt.Println(v4, v1)
    if (false || true) {
        var v22 float32 = (0.1 + (0.0 * 0.1))
        var v23 int = 128
        _ = v22
        _ = v23
    } else {
        if ((1.5 * 0.1) > 0.0) {
            v24 := []int{}
            for fi25 := 0; fi25 < 13; fi25++ {
                v1 += v1
                for fi26 := 0; fi26 < 19; fi26++ {
                    v24 = append(v24, (v4 / (((v3 ^ v4)) | 1)))
                    var v27 int32 = v5
                    v28 := GenS2{F0: -1, F1: 791, F2: 42}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v29 := GenS1{F0: true, F1: "a"}
                _ = fi25
                _ = v29
            }
            _ = v24
        } else {
            if false {
                v30 := []int{(v4 & -441), v4, (v3 & 32768), (v4 + v3)}
                v31 := GenS3{F0: 212, F1: 56}
                _ = v30
                _ = v31
            } else {
                var v32 int = -9223372036854775808
                _ = v32
            }
            v5 -= (^(v5))
        }
        if (true || (false && false)) {
            fmt.Println(v3)
            switch "hello" {
            case "":
                fmt.Println(v5, v5)
            case "a":
                for fi33 := 0; fi33 < 13; fi33++ {
                    for fi34 := 0; fi34 < 16; fi34++ {
                        v35 := []int{v4}
                        fmt.Println(fi33)
                        var v36 string = ((("" + "\n\t") + "hello") + ("\n\t" + "foo bar"))
                        var v37 int32 = (v5 ^ v5)
                        v3 += int(v37)
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    switch (fi33 | fi33) {
                    case 485:
                        var v38 int32 = genFn1(0)
                        _ = v38
                    default:
                        v39 := GenS2{F0: 42, F1: 0, F2: 53}
                        _ = v39
                    }
                    v40 := GenS1{F0: (!(0 == 255)), F1: ("a" + ("a" + "hello"))}
                    _ = fi33
                    _ = v40
                }
            case "foo bar":
                fmt.Println(v4, v1)
            default:
                var v41 float32 = 1.5
                _ = v41
            }
        }
    }
    v42 := GenS3{F0: 127, F1: 30}
    switch 9223372036854775807 {
    case 32767:
        var v43 uint64 = 4294967295
        _ = v43
    case -790:
        v5 = v5
    }
    v5 = (-(v5))
    if (((124 == 128) && true) || (42 != 0)) {
        v3 += (v3 ^ 1)
    }
    v44 := GenS4{F0: (v1 % (((v1 / ((v1) | 1))) | 1)), F1: int8(v3), F2: (v1 | (v1 + 379))}
    var v45 uint = 20
    v46 := GenS3{F0: 127, F1: 0}
    v4 = v3
    if false {
        if (4294967295 != v4) {
            v4 = (^(v3))
            for fi47 := 0; fi47 < 10; fi47++ {
                v48 := GenS5{F0: 100, F1: 255}
                v49 := GenS1{F0: false, F1: "foo bar"}
                v5 -= 881
                _ = fi47
                _ = v48
                _ = v49
            }
            v50 := []int{v4, (v4 ^ (v4 << ((2) & 63)))}
            _ = v50
        } else {
            for fi51 := 0; fi51 < 3; fi51++ {
                v52 := GenS1{F0: (false || false), F1: "a"}
                _ = fi51
                _ = v52
            }
            v3 -= 32768
        }
    }
    if (4294967295 < v3) {
        if (!(4294967295 <= 1)) {
            v53 := GenS4{F0: v1, F1: 100, F2: (v1 >> ((v45) & 63))}
            var v54 int32 = v5
            _ = v53
            _ = v54
        }
        v55 := v2[(((("hello" + ((("" + "a") + "") + "hello")) + "foo bar") + "foo bar") + (("" + "\n\t") + "foo bar"))]
        _ = v55
    }
    var v56 bool = false
    if (42 == (v3 | v4)) {
        switch 128 {
        case 32767:
            v57 := GenS6{F0: 11, F1: v5}
            _ = v57
        case 127:
            for fi58 := 0; fi58 < 19; fi58++ {
                v56 = ((v4 / ((v3) | 1)) != v4)
                var v59 uint8 = 0
                fmt.Println(v59, v59)
                _ = fi58
                _ = v59
            }
        default:
            if (false && true) {
                v56 = (!(v56 || false))
                switch "" {
                case "a":
                    v1 = v1
                default:
                    delete(v2, "\n\t")
                }
                for fi60 := 0; fi60 < 16; fi60++ {
                    if (528 != 1000) {
                        switch "a" {
                        case "hello":
                            for fi61 := 0; fi61 < 6; fi61++ {
                                var v62 float64 = (3.14159 / 1.5)
                                fmt.Println(v5, fi60)
                                v63 := GenS7{F0: v3, F1: 255, F2: 18446744073709551615}
                                _ = fi61
                                _ = v62
                                _ = v63
                            }
                        case "a":
                            for fi64 := 0; fi64 < 12; fi64++ {
                                switch 9223372036854775807 {
                                case -1:
                                    switch "hello" {
                                    case "hello":
                                        v65 := v2["foo bar"]
                                        _ = v65
                                    }
                                case 0:
                                    var v66 string = "a"
                                    _ = v66
                                case 32768:
                                    switch "foo bar" {
                                    case "hello":
                                        v67 := GenS8{F0: 255, F1: 728}
                                        _ = v67
                                    }
                                }
                                if true {
                                    fmt.Println(v45, v3)
                                } else {
                                    v5 += v5
                                }
                                _ = fi64
                            }
                        case "":
                            switch v4 {
                            case -86:
                                var v68 uint32 = 107
                                _ = v68
                            case 32768:
                                v69 := GenS9{F0: 56, F1: 351}
                                _ = v69
                            default:
                                if (1e10 <= ((1.5 * (float64(0.0) / (0.0 - 1e10))) + 1e10)) {
                                    switch ("foo bar" + (("foo bar" + "a") + "a")) {
                                    case "a":
                                        var v70 float32 = (0.0 - 0.1)
                                        _ = v70
                                    case "foo bar":
                                        if true {
                                            for fi71 := 0; fi71 < 8; fi71++ {
                                                v72 := []int{(^(v4)), -253, (fi60 % (((^(v4))) | 1)), 127}
                                                var v73 uint64 = genFn0((0 != (v5 / (((-(v5))) | 1))), (^(v45)))
                                                _ = fi71
                                                _ = v72
                                                _ = v73
                                            }
                                            fmt.Println(v4)
                                            v45 = (^(v45))
                                        }
                                    default:
                                        if false {
                                            fmt.Println(fi60, v56)
                                            fmt.Println(v56, v5, v1)
                                        }
                                    }
                                    var v74 float64 = 1.5
                                    _ = v74
                                }
                            }
                        default:
                            v75 := GenS3{F0: 128, F1: 42}
                            _ = v75
                        }
                    }
                    _ = fi60
                }
            } else {
                v2["a"] = 65535
                v5 = (v5 >> ((v45) & 63))
                v56 = ((v45 % ((v45) | 1)) == (v45 ^ v45))
            }
        }
        if (0 >= 32767) {
            switch v4 {
            case -973:
                v76 := []int{}
                _ = v76
            case 9223372036854775807:
                switch (^(v3)) {
                case -83:
                    v77 := []int{(v4 ^ (v3 / (((v3 | (v3 / ((v3) | 1)))) | 1))), (-(v4))}
                    _ = v77
                case -2147483648:
                    var v78 int = 390
                    _ = v78
                }
            default:
                v4 = 42
            }
        }
    }
    var v79 float32 = 0.1
    var v80 int = -258
    v45 = v45
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v56
    _ = v79
    _ = v80
}
