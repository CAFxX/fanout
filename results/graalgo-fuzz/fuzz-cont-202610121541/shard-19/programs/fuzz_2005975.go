package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int64; F2 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 int16; F1 string; F2 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int32; F1 int32 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint }
func genFn0(p0 uint64, p1 int64) int64 {
    p1 += p1
    p1 -= p1
    v5 := GenS1{F0: (!false), F1: 144, F2: (!true)}
    _ = p0
    _ = p1
    _ = v5
    return p1
}
func genFn1(p0 int8) uint16 {
    v6 := GenS1{F0: true, F1: 1, F2: false}
    if (1000 <= 0) {
        if (!(!false)) {
            var v7 int16 = 91
            if (((!(145 > 0)) && (!(false && true))) && ((p0 % ((p0) | 1)) > (p0 - (p0 | -1)))) {
                v8 := []int{}
                v9 := GenS2{F0: 15, F1: 18446744073709551615}
                _ = v8
                _ = v9
            } else {
                for fi10 := 0; fi10 < 18; fi10++ {
                    switch 358 {
                    case -545:
                        var v11 int16 = v7
                        _ = v11
                    case 598:
                        v12 := GenS3{F0: (-(p0))}
                        _ = v12
                    case 9223372036854775807:
                        if false {
                            var v13 uint16 = 42
                            v14 := []int{(-(fi10)), fi10}
                            _ = v13
                            _ = v14
                        }
                    default:
                        p0 = (p0 / ((p0) | 1))
                    }
                    if (!false) {
                        v15 := GenS4{F0: p0, F1: -9223372036854775808, F2: (fi10 >> ((255) & 63))}
                        switch -364 {
                        case 42:
                            for fi16 := 0; fi16 < 5; fi16++ {
                                var v17 int8 = p0
                                fmt.Println(fi16, v17)
                                var v18 float32 = (3.14159 * -2.25)
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        }
                        _ = v15
                    } else {
                        v19 := GenS5{F0: 1}
                        _ = v19
                    }
                    p0 = p0
                    _ = fi10
                }
            }
            _ = v7
        } else {
            if (1 != 9223372036854775807) {
                v20 := []int{}
                _ = len(v20)
                _ = v20
            }
            if (-1 < -1) {
                if (false && true) {
                    switch -855 {
                    case -2147483648:
                        p0 = p0
                    case 42:
                        var v21 uint64 = 1
                        _ = v21
                    }
                } else {
                    switch "hello" {
                    case "a":
                        p0 = -30
                    case "":
                        p0 += (p0 & p0)
                    default:
                        var v22 uint16 = 0
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v24 := GenS6{F0: 42, F1: "hello"}
                        _ = fi23
                        _ = v24
                    }
                    if ((-(p0)) > (p0 ^ -128)) {
                        var v25 int64 = 9223372036854775807
                        fmt.Println(p0, v25)
                        v26 := []int{32767, 9223372036854775807, 4294967295}
                        _ = v25
                        _ = v26
                    }
                }
                var v27 uint16 = 65535
                if (!false) {
                    v28 := GenS7{F0: 590, F1: (("a" + "hello") + "a"), F2: 446}
                    switch ("a" + "hello") {
                    case "hello":
                        for fi29 := 0; fi29 < 3; fi29++ {
                            var v30 uint64 = 301
                            for fi31 := 0; fi31 < 18; fi31++ {
                                fmt.Println(fi31)
                                var v32 float32 = (float32(1e10) + (0.0 / 1.5))
                                v30 = (v30 | v30)
                                v33 := GenS4{F0: (p0 + (p0 + int8(v30))), F1: 0, F2: 9223372036854775807}
                                switch ((("hello" + "a") + ("foo bar" + ("foo bar" + ((("hello" + ("hello" + "foo bar")) + "") + "")))) + "a") {
                                case "\n\t":
                                    var v34 int32 = 370
                                    _ = v34
                                case "":
                                    p0 = 100
                                    v30 = v30
                                    switch "" {
                                    case "foo bar":
                                        switch "hello" {
                                        case "\n\t":
                                            var v35 int8 = -128
                                            v36 := GenS5{F0: (^(v27))}
                                            _ = v35
                                            _ = v36
                                        }
                                    case "hello":
                                        var v37 string = ""
                                        _ = v37
                                    default:
                                        fmt.Println(fi29, v30)
                                    }
                                    v38 := []int{fi29, 441}
                                    _ = v38
                                case "foo bar":
                                    v39 := []int{fi29, fi31, (fi31 + 42)}
                                    _ = v39
                                default:
                                    if true {
                                        fmt.Println(p0)
                                        switch fi31 {
                                        case 127:
                                            v30 = (v30 % ((v30) | 1))
                                        case 32768:
                                            v40 := []int{978}
                                            _ = v40
                                        case -112:
                                            v41 := []int{}
                                            _ = v41
                                        default:
                                            v42 := GenS5{F0: 2}
                                            _ = v42
                                        }
                                    }
                                }
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            _ = fi29
                            _ = v30
                        }
                    default:
                        switch ("" + "") {
                        case "a":
                            fmt.Println(p0, p0)
                        case "hello":
                            if true {
                                p0 -= 127
                                var v43 int = 394
                                fmt.Println(v43, v27)
                                _ = v43
                            } else {
                                v44 := []int{-9223372036854775808, 0, 32768}
                                for fi45 := 0; fi45 < 6; fi45++ {
                                    var v46 int64 = 9223372036854775807
                                    for fi47 := 0; fi47 < 6; fi47++ {
                                        switch ("\n\t" + "hello") {
                                        case "hello":
                                            if len(v44) > 0 {
                                                v44[0] = fi47
                                            }
                                        case "\n\t":
                                            var v48 bool = true
                                            _ = v48
                                        }
                                        _ = fi47
                                    }
                                    _ = fi45
                                    _ = v46
                                }
                                _ = v44
                            }
                        case "foo bar":
                            for fi49 := 0; fi49 < 19; fi49++ {
                                var v50 bool = ((p0 | (p0 ^ (p0 >> ((1) & 63)))) == (^(p0)))
                                switch (("a" + ("\n\t" + "")) + "\n\t") {
                                case "\n\t":
                                    v51 := GenS2{F0: 280, F1: 42}
                                    _ = v51
                                case "hello":
                                    if (v50 || ((fi49 * 32767) >= (-(fi49)))) {
                                        v52 := []int{fi49, fi49, 32768}
                                        _ = v52
                                    }
                                }
                                _ = fi49
                                _ = v50
                            }
                        default:
                            for fi53 := 0; fi53 < 19; fi53++ {
                                v54 := GenS8{F0: 128}
                                fmt.Println(p0, fi53, p0)
                                _ = fi53
                                _ = v54
                            }
                        }
                    }
                    _ = v28
                }
                _ = v27
            } else {
                v55 := []int{-897, -2147483648, 127}
                p0 += p0
                for fi56 := 0; fi56 < 14; fi56++ {
                    if len(v55) > 0 {
                        v55[0] = (fi56 - 441)
                    }
                    _ = fi56
                }
                _ = v55
            }
        }
    }
    _ = p0
    _ = v6
    return 32767
}

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]bool)
    var v3 uint8 = 202
    var v4 int64 = 388
    if true {
        var v57 int = 32768
        v3 = (v3 & v3)
        _ = v57
    } else {
        if (true && true) {
            v58 := []int{42, -471, 883, -439}
            for fi59 := 0; fi59 < 3; fi59++ {
                _ = len(v2)
                v1 -= v1
                v3 = v3
                _ = fi59
            }
            fmt.Println(v1, v4)
            _ = v58
        } else {
            if true {
                var v60 int = 32767
                switch "\n\t" {
                case "hello":
                    v1 = (v1 ^ int8(v60))
                default:
                    var v61 float32 = (1e10 / (1.5 * 1e10))
                    _ = v61
                }
                for fi62 := 0; fi62 < 2; fi62++ {
                    var v63 bool = (false && false)
                    _ = fi62
                    _ = v63
                }
                _ = v60
            }
            if true {
                if true {
                    for fi64 := 0; fi64 < 16; fi64++ {
                        var v65 bool = (0.1 <= float32(float64(0.1)))
                        var v66 int32 = 0
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                } else {
                    v1 = (v1 >> ((82) & 63))
                }
            } else {
                if (!true) {
                    fmt.Println(v1)
                    var v67 int64 = -991
                    v68 := GenS5{F0: 255}
                    _ = v67
                    _ = v68
                }
            }
        }
        v69 := GenS1{F0: (!false), F1: v3, F2: false}
        var v70 uint16 = 635
        _ = v69
        _ = v70
    }
    var v71 uint64 = 670
    v72 := []int{128, 127, 544}
    v73 := GenS9{F0: 434}
    var v74 uint32 = 42
    if false {
        fmt.Println(v4, v71, v4)
        v75 := GenS8{F0: v3}
        _ = v75
    }
    if false {
        v76 := GenS10{F0: 1, F1: 0}
        v4 -= (-(v4))
        _ = v76
    } else {
        v71 -= (v71 ^ 1)
        if true {
            v77 := GenS1{F0: false, F1: 166, F2: false}
            var v78 int64 = (v4 - -9223372036854775808)
            _ = v77
            _ = v78
        }
        for fi79 := 0; fi79 < 19; fi79++ {
            if (((v71 % ((2) | 1)) < 137) || false) {
                v1 += -1
                delete(v2, v74)
                v80 := GenS2{F0: 2147483647, F1: 42}
                _ = v80
            }
            var v81 string = ""
            v81 = v81
            _ = fi79
            _ = v81
        }
    }
    v82 := GenS11{F0: -32768}
    v83 := GenS12{F0: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v82
    _ = v83
}
