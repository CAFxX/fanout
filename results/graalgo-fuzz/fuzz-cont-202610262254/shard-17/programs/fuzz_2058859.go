package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint8; F1 int16; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint32; F2 int }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8; F1 uint8 }
func genFn0() int64 {
    var v5 string = "hello"
    var v6 int = -313
    v7 := GenS1{F0: 2, F1: 18446744073709551615}
    for fi8 := 0; fi8 < 16; fi8++ {
        v6 = (^(v6))
        v9 := []int{(v6 / (((v6 % ((-872) | 1))) | 1))}
        _ = fi8
        _ = v9
    }
    _ = v5
    _ = v6
    _ = v7
    return 932
}
func genFn1(p0 int, p1 int, p2 uint8) int64 {
    p1 = (p1 / (((p1 | (^(p1)))) | 1))
    v10 := GenS1{F0: (p2 % ((p2) | 1)), F1: 65535}
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := []int{(p1 & p0), p1, 359, p0}
        _ = len(v12)
        var v13 int16 = -401
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int16 = -1
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 uint16 = 299
        for fi17 := 0; fi17 < 14; fi17++ {
            if (4294967295 <= 42) {
                switch (fi17 ^ 417) {
                case 500:
                    fmt.Println(v14, p2, fi15)
                }
                switch (("" + "foo bar") + "\n\t") {
                case "":
                    switch ("a" + "hello") {
                    case "a":
                        v18 := []int{127}
                        var v19 int8 = -1
                        v20 := GenS1{F0: (p2 - (p2 & p2)), F1: 18446744073709551615}
                        _ = v18
                        _ = v19
                        _ = v20
                    case "":
                        p1 = fi17
                    case "foo bar":
                        if (!true) {
                            switch ("\n\t" + "a") {
                            case "hello":
                                if (!false) {
                                    v14 += 42
                                }
                            case "":
                                switch 128 {
                                case 4294967295:
                                    var v21 int64 = 716
                                    if ((v21 & genFn0()) >= v21) {
                                        v14 = v14
                                        switch (fi15 * (fi15 + 4294967295)) {
                                        case 128:
                                            v14 = v14
                                        case 0:
                                            v21 = v21
                                        default:
                                            p1 -= 32767
                                            fmt.Println(p2)
                                            var v22 int64 = 349
                                            _ = v22
                                        }
                                        var v23 int16 = 1000
                                        _ = v23
                                    }
                                    _ = v21
                                case -354:
                                    p1 -= fi15
                                case 32768:
                                    v24 := []int{fi17, fi17, fi17, p0}
                                    _ = v24
                                default:
                                    v16 = v16
                                }
                            case "a":
                                v16 = v16
                            }
                            v25 := GenS1{F0: (^(p2)), F1: 77}
                            switch (fi15 >> ((255) & 63)) {
                            case 32768:
                                p2 += p2
                            default:
                                var v26 int64 = -696
                                _ = v26
                            }
                            _ = v25
                        }
                    default:
                        var v27 bool = ((18446744073709551615 >= 1) || false)
                        _ = v27
                    }
                case "\n\t":
                    if (0 != -1) {
                        v28 := GenS1{F0: 95, F1: 42}
                        var v29 string = ("foo bar" + (("\n\t" + "foo bar") + ("a" + "foo bar")))
                        _ = v28
                        _ = v29
                    }
                default:
                    p0 += (fi15 - (^(p0)))
                }
                v16 -= v16
            } else {
                p0 -= p1
                v30 := GenS2{F0: uint8(v16), F1: v14, F2: false}
                _ = v30
            }
            switch "hello" {
            case "\n\t":
                var v31 uint64 = 4294967295
                _ = v31
            default:
                var v32 uint8 = p2
                _ = v32
            }
            v33 := GenS3{F0: (p2 / (((p2 << ((65535) & 63))) | 1))}
            _ = fi17
            _ = v33
        }
        v16 = v16
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v14
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 1
    var v2 uint = 4294967295
    var v3 int = 2147483647
    var v4 int = -2147483648
    var v34 int = v4
    var v35 uint8 = 1
    v36 := GenS3{F0: 156}
    for fi37 := 0; fi37 < 12; fi37++ {
        for fi38 := 0; fi38 < 20; fi38++ {
            switch "" {
            case "hello":
                var v39 uint32 = uint32(v2)
                _ = v39
            case "\n\t":
                v40 := []int{-717}
                _ = v40
            case "foo bar":
                v41 := GenS4{F0: (^(v35))}
                _ = v41
            }
            v1 = (v1 << ((uint(v1)) & 63))
            v1 = 0
            _ = fi38
        }
        v42 := GenS5{F0: "\n\t", F1: 4294967295, F2: (fi37 >> (((^(v2))) & 63))}
        _ = fi37
        _ = v42
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        v4 = (fi43 >> ((2) & 63))
        if true {
            v44 := GenS2{F0: v35, F1: -1, F2: ((1.5 * ((-2.25 + (1.5 - (1.5 + 1e10))) + 0.0)) < (0.0 - (1e10 - -2.25)))}
            switch "hello" {
            case "":
                v1 = v1
            }
            _ = v44
        }
        _ = fi43
    }
    if (v3 < (-(v3))) {
        switch v4 {
        case 618:
            v35 += (^(v35))
        case 486:
            fmt.Println(v34)
        }
        switch v4 {
        case -490:
            if ((v35 == v35) && true) {
                var v45 uint = (v2 - v2)
                v4 -= (^(v3))
                v46 := GenS5{F0: "a", F1: 2147483647, F2: (v3 ^ v34)}
                _ = v45
                _ = v46
            } else {
                var v47 float32 = 0.1
                v48 := []int{(v34 | (v3 << ((v2) & 63)))}
                v49 := GenS2{F0: v35, F1: -1, F2: (!(42 > 65535))}
                _ = v47
                _ = v48
                _ = v49
            }
        }
    } else {
        v1 -= 32767
        if false {
            v34 = v3
            v3 = (v34 & v3)
        }
        v50 := GenS2{F0: (^(v35)), F1: -13, F2: true}
        _ = v50
    }
    if false {
        v51 := []int{-280, v34}
        _ = v51
    } else {
        var v52 string = ""
        _ = v52
    }
    for fi53 := 0; fi53 < 4; fi53++ {
        var v54 string = "a"
        v55 := GenS2{F0: v35, F1: 0, F2: true}
        fmt.Println(v2, v34)
        _ = fi53
        _ = v54
        _ = v55
    }
    if (true && true) {
        switch (("" + (("" + "hello") + "\n\t")) + "a") {
        case "a":
            for fi56 := 0; fi56 < 19; fi56++ {
                if ((true || (!false)) && false) {
                    v57 := []int{v3}
                    _ = v57
                }
                _ = fi56
            }
        case "hello":
            for fi58 := 0; fi58 < 6; fi58++ {
                switch v4 {
                case 128:
                    v59 := GenS6{F0: v2, F1: (^(v1))}
                    _ = v59
                case -933:
                    var v60 uint8 = v35
                    _ = v60
                }
                v35 = 131
                v34 += v4
                _ = fi58
            }
        case "":
            v4 += (v34 - 9223372036854775807)
        }
    } else {
        var v61 bool = true
        for fi62 := 0; fi62 < 3; fi62++ {
            for fi63 := 0; fi63 < 18; fi63++ {
                v64 := []int{(fi62 + (fi62 >> (((v2 * (v2 / ((v2) | 1)))) & 63))), (v3 % (((^(fi63))) | 1)), (-(fi63)), -408}
                fmt.Println(v34, v35)
                _ = fi63
                _ = v64
            }
            v61 = v61
            _ = fi62
        }
        _ = v61
    }
    for fi65 := 0; fi65 < 7; fi65++ {
        v66 := GenS7{F0: 32767}
        if false {
            v67 := []int{}
            v68 := GenS2{F0: (v35 >> (((^(v2))) & 63)), F1: -32768, F2: ((56 >= 1) && true)}
            v69 := GenS5{F0: (("" + "foo bar") + "hello"), F1: 2, F2: (fi65 | (v3 | v3))}
            _ = v67
            _ = v68
            _ = v69
        } else {
            switch "\n\t" {
            case "\n\t":
                for fi70 := 0; fi70 < 18; fi70++ {
                    v71 := []int{}
                    if false {
                        v34 = (fi70 % ((v3) | 1))
                        var v72 int16 = -242
                        _ = v72
                    }
                    fmt.Println(v4, v4)
                    _ = fi70
                    _ = v71
                }
            case "foo bar":
                var v73 float64 = 0.0
                _ = v73
            }
            for fi74 := 0; fi74 < 19; fi74++ {
                v3 -= fi65
                v3 = v3
                _ = fi74
            }
        }
        _ = fi65
        _ = v66
    }
    fmt.Println(v34)
    v75 := GenS7{F0: 32767}
    if false {
        var v76 uint64 = 1
        fmt.Println(v2, v35)
        v77 := []int{(v34 ^ (v34 << ((0) & 63))), v3, (v3 * (v34 + (v4 | 32767))), v3}
        _ = v76
        _ = v77
    }
    v78 := []int{(v4 & v34), v34}
    switch "hello" {
    case "":
        v79 := GenS8{F0: v35, F1: v35}
        _ = v79
    }
    v4 = -884
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v78:", len(v78), v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v75
    _ = v78
}
