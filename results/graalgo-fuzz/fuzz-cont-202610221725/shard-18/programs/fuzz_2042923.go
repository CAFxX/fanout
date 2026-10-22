package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint; F1 string }
func genFn0(p0 uint16) int32 {
    var v5 uint16 = p0
    if (!(2147483647 >= 99)) {
        if (127 > 1) {
            var v6 int16 = 42
            _ = v6
        } else {
            var v7 int16 = -575
            v8 := GenS1{F0: (p0 * p0)}
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := GenS2{F0: 18446744073709551615}
                _ = fi9
                _ = v10
            }
            _ = v7
            _ = v8
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            p0 += 301
            switch fi11 {
            case 1:
                p0 = (p0 % ((255) | 1))
            default:
                p0 = 65535
            }
            _ = fi11
        }
    } else {
        if true {
            var v12 int = -962
            v5 = v5
            _ = v12
        }
        switch 127 {
        case -403:
            var v13 uint32 = 747
            _ = v13
        case 127:
            fmt.Println(v5, p0)
        case 32767:
            if (119 >= -1) {
                var v14 int32 = 42
                var v15 string = ("a" + ("" + ("hello" + "a")))
                _ = v14
                _ = v15
            } else {
                v16 := GenS2{F0: 650}
                _ = v16
            }
        }
        var v17 int32 = 2147483647
        _ = v17
    }
    switch "a" {
    case "foo bar":
        if true {
            p0 = 476
            v18 := GenS1{F0: 1}
            _ = v18
        }
    case "hello":
        var v19 uint8 = 1
        _ = v19
    default:
        var v20 uint64 = 18446744073709551615
        _ = v20
    }
    _ = p0
    _ = v5
    return 2147483647
}
func genFn1() int16 {
    var v21 uint = 198
    var v22 uint8 = 143
    _ = v21
    _ = v22
    return -1
}

func main() {
    var v1 float32 = 0.1
    var v2 int32 = -2147483648
    v3 := make(map[uint32]int)
    v4 := make(map[int]string)
    v23 := GenS1{F0: 65535}
    fmt.Println(v2)
    for fi24 := 0; fi24 < 4; fi24++ {
        v1 += v1
        var v25 int8 = -1
        for fi26 := 0; fi26 < 2; fi26++ {
            delete(v4, (fi26 / ((1) | 1)))
            delete(v4, (-(fi26)))
            fmt.Println(v25, fi26)
            if false {
                switch "hello" {
                case "hello":
                    v27 := GenS2{F0: 4294967295}
                    _ = v27
                case "\n\t":
                    var v28 float64 = 0.0
                    _ = v28
                case "a":
                    var v29 int = 9223372036854775807
                    _ = v29
                default:
                    var v30 bool = false
                    _ = v30
                }
                fmt.Println(v1)
                v25 = (v25 % (((^(v25))) | 1))
            } else {
                v3[1] = fi24
                delete(v3, 24)
            }
            v31 := GenS3{F0: 529, F1: (fi24 + (fi26 >> ((2) & 63)))}
            _ = fi26
            _ = v31
        }
        _ = fi24
        _ = v25
    }
    switch 2147483647 {
    case 2147483647:
        if true {
            if true {
                v32 := v4[32768]
                _ = v32
            }
            var v33 uint8 = 163
            if false {
                switch 0 {
                case 134:
                    if (671 == 0) {
                        v34 := GenS2{F0: 65535}
                        for fi35 := 0; fi35 < 1; fi35++ {
                            if true {
                                var v36 uint = 2
                                for fi37 := 0; fi37 < 19; fi37++ {
                                    if (994 <= 4294967295) {
                                        var v38 uint8 = v33
                                        switch (fi37 * fi37) {
                                        case 127:
                                            v39 := GenS4{F0: 255, F1: 741}
                                            _ = v39
                                        case 1:
                                            var v40 int = (fi37 / ((-9223372036854775808) | 1))
                                            _ = v40
                                        case 128:
                                            var v41 string = ((("\n\t" + "foo bar") + "hello") + "")
                                            _ = v41
                                        }
                                        v2 -= (v2 << ((v36) & 63))
                                        _ = v38
                                    } else {
                                        var v42 int64 = 174
                                        v43 := GenS4{F0: 42, F1: 1}
                                        v44 := []int{42, fi37, 128}
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    switch 127 {
                                    case 32768:
                                        var v45 int8 = -128
                                        _ = v45
                                    case 665:
                                        var v46 int = 2147483647
                                        _ = v46
                                    case -188:
                                        delete(v4, (fi35 | fi35))
                                    default:
                                        var v47 int = (fi37 % ((fi37) | 1))
                                        _ = v47
                                    }
                                    if true {
                                        v48 := GenS3{F0: 2, F1: (fi37 >> ((v36) & 63))}
                                        var v49 uint32 = 464
                                        _ = v48
                                        _ = v49
                                    } else {
                                        var v50 float64 = (float64(v1) * (3.14159 - 3.14159))
                                        _ = v50
                                    }
                                    _ = fi37
                                }
                                v51 := GenS1{F0: 0}
                                _ = v36
                                _ = v51
                            } else {
                                v2 -= v2
                                v52 := []int{(fi35 * fi35)}
                                _ = v52
                            }
                            if true {
                                for fi53 := 0; fi53 < 14; fi53++ {
                                    v54 := GenS5{F0: false, F1: 752}
                                    v1 = 1e10
                                    v55 := []int{}
                                    _ = fi53
                                    _ = v54
                                    _ = v55
                                }
                            }
                            _ = fi35
                        }
                        _ = v34
                    }
                }
                v33 = v33
            }
            _ = v33
        } else {
            v56 := []int{-1, -270, 42, 128}
            for fi57 := 0; fi57 < 18; fi57++ {
                fmt.Println(v1)
                v1 -= (((1e10 / 0.1) + v1) / (float32((1e10 + (((float64(3.14159) - float64(0.0)) - float64(float32(1.5))) * 1e10))) + v1))
                _ = fi57
            }
            _ = v56
        }
    case 4294967295:
        v58 := []int{-673, 32768, -1, -2147483648}
        _ = v58
    default:
        for fi59 := 0; fi59 < 16; fi59++ {
            switch ("" + "hello") {
            case "a":
                var v60 uint32 = 4294967295
                _ = v60
            case "\n\t":
                switch "hello" {
                case "hello":
                    switch "foo bar" {
                    case "foo bar":
                        v61 := GenS6{F0: v2, F1: (^(fi59))}
                        _ = v61
                    }
                case "a":
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v1, v1)
                    case "hello":
                        var v62 int8 = 42
                        _ = v62
                    }
                }
            case "foo bar":
                v3[42] = (fi59 - (^(fi59)))
            }
            var v63 int = (fi59 - fi59)
            v1 = 1e10
            _ = fi59
            _ = v63
        }
    }
    v4[668] = ""
    if true {
        for fi64 := 0; fi64 < 13; fi64++ {
            var v65 float32 = 1e10
            var v66 int64 = 1
            _ = fi64
            _ = v65
            _ = v66
        }
        v67 := GenS7{F0: 1}
        _ = len(v4)
        _ = v67
    } else {
        v68 := []int{}
        v2 -= (v2 / ((15) | 1))
        var v69 string = ""
        _ = v68
        _ = v69
    }
    delete(v3, 2147483647)
    v2 += (v2 % ((42) | 1))
    v70 := v3[4294967295]
    switch v70 {
    case 1:
        if ((32767 < 255) && true) {
            if false {
                var v71 float32 = float32(1.5)
                var v72 int16 = 447
                _ = v71
                _ = v72
            }
            v73 := []int{int(v2), (v70 - 9223372036854775807), -1}
            _ = v73
        } else {
            for fi74 := 0; fi74 < 17; fi74++ {
                v75 := []int{v70}
                v76 := GenS8{F0: 42, F1: "foo bar"}
                _ = fi74
                _ = v75
                _ = v76
            }
        }
    case -9223372036854775808:
        delete(v3, 391)
    }
    if (254 < 2147483647) {
        for fi77 := 0; fi77 < 14; fi77++ {
            switch 1 {
            case 42:
                v78 := GenS7{F0: 1}
                _ = v78
            }
            v2 = (^(v2))
            _ = fi77
        }
        var v79 int64 = 0
        for fi80 := 0; fi80 < 1; fi80++ {
            if (1 == 255) {
                if (18446744073709551615 > 1) {
                    var v81 uint32 = 800
                    v1 = 0.0
                    _ = v81
                } else {
                    v3[1] = v70
                }
            } else {
                if (127 < 36) {
                    v82 := GenS2{F0: 56}
                    _ = v82
                }
            }
            _ = fi80
        }
        _ = v79
    }
    var v83 uint32 = 2147483648
    var v84 uint64 = 18446744073709551615
    delete(v4, (v70 | 128))
    switch ((("" + "a") + "\n\t") + "") {
    case "hello":
        v85 := GenS2{F0: 4294967295}
        _ = v85
    case "":
        for fi86 := 0; fi86 < 1; fi86++ {
            v87 := GenS6{F0: int32(fi86), F1: fi86}
            _ = fi86
            _ = v87
        }
    case "\n\t":
        v1 -= v1
    }
    fmt.Println(v70)
    switch (v70 + (-(v70))) {
    case 0:
        for fi88 := 0; fi88 < 5; fi88++ {
            v3[1] = (fi88 - -9223372036854775808)
            _ = fi88
        }
    case 128:
        for fi89 := 0; fi89 < 4; fi89++ {
            var v90 float32 = (v1 / ((((v1 / float32(((1e10 / float64(((float32((((((float64((v1 * v1)) + 3.14159) + float64((float32(1.5) * (-2.25 - v1)))) + 1.5) * (((0.1 * (0.1 - 3.14159)) * -2.25) * (0.0 - (-2.25 * 3.14159)))) + (1e10 / (3.14159 + float64(float32((((3.14159 * 0.1) + 1.5) * 3.14159))))))) - 3.14159) * 3.14159))) + -2.25))) * 1.5) - v1) - v1))
            _ = fi89
            _ = v90
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v70:", v70)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v70
    _ = v83
    _ = v84
}
