package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 uint }
func genFn0(p0 int16, p1 uint8, p2 uint64) int {
    var v4 bool = false
    switch "hello" {
    case "\n\t":
        var v5 string = "foo bar"
        _ = v5
    case "":
        var v6 uint64 = (p2 * (p2 + 0))
        _ = v6
    default:
        v4 = false
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32768
}
func genFn1() int16 {
    v7 := GenS1{F0: 42, F1: -1, F2: 1}
    v8 := GenS1{F0: 91, F1: 1000, F2: 2}
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 int64 = -177
    var v2 int32 = 1
    var v3 float32 = 3.14159
    v9 := GenS1{F0: 0, F1: genFn1(), F2: 623}
    v10 := GenS2{F0: true, F1: (v1 & v1), F2: 65535}
    if false {
        var v11 bool = false
        _ = v11
    }
    v12 := []int{}
    if true {
        if false {
            v3 -= 1.5
        }
    }
    if ((!true) && (42 > genFn0(-805, 234, 1))) {
        v13 := GenS3{F0: v1, F1: genFn1()}
        _ = v13
    } else {
        v14 := GenS4{F0: 0}
        v2 = 0
        _ = v14
    }
    var v15 uint16 = 833
    v16 := GenS1{F0: -128, F1: 0, F2: 0}
    var v17 int64 = (v1 ^ 1)
    _ = len(v12)
    v18 := GenS5{F0: uint(v15)}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v1 -= (^(v17))
        for fi20 := 0; fi20 < 8; fi20++ {
            switch fi19 {
            case 0:
                fmt.Println(fi20, fi20)
            case 105:
                v21 := GenS2{F0: ((127 != -128) || ((v15 <= (v15 & v15)) || false)), F1: v1, F2: 0}
                _ = v21
            default:
                v15 -= v15
            }
            v22 := GenS5{F0: 255}
            v15 += v15
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            v15 += 42
            fmt.Println(v17, fi19)
            var v23 uint16 = (v15 ^ v15)
            v24 := GenS1{F0: -1, F1: genFn1(), F2: 42}
            _ = fi20
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi19
    }
    if false {
        switch ("" + ("\n\t" + ("foo bar" + ""))) {
        case "foo bar":
            for fi25 := 0; fi25 < 18; fi25++ {
                if true {
                    v1 -= v1
                } else {
                    switch 42 {
                    case -672:
                        _ = len(v12)
                    case -2147483648:
                        if false {
                            var v26 int64 = v17
                            v27 := GenS4{F0: 272}
                            v3 = v3
                            _ = v26
                            _ = v27
                        } else {
                            switch "a" {
                            case "hello":
                                v12 = append(v12, genFn0(42, 176, 18446744073709551615))
                            case "a":
                                fmt.Println(v1, v15)
                            }
                            if (1 >= uint64(v15)) {
                                fmt.Println(v17, v1, v2)
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            } else {
                                fmt.Println(fi25, v1)
                                fmt.Println(v17)
                            }
                        }
                    case 268:
                        v28 := GenS6{F0: 42, F1: v2}
                        _ = v28
                    }
                    var v29 bool = (!(71 >= 255))
                    if false {
                        if false {
                            v30 := GenS7{F0: true, F1: (v2 % (((v2 ^ 0)) | 1))}
                            switch "hello" {
                            case "hello":
                                switch genFn0(781, 128, 0) {
                                case 2147483647:
                                    for fi31 := 0; fi31 < 9; fi31++ {
                                        var v32 uint32 = 2
                                        _ = fi31
                                        _ = v32
                                    }
                                case -537:
                                    for fi33 := 0; fi33 < 1; fi33++ {
                                        var v34 int8 = 1
                                        var v35 int8 = 1
                                        fmt.Println(v35)
                                        if true {
                                            fmt.Println(v2)
                                        } else {
                                            var v36 bool = v29
                                            _ = v36
                                        }
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                case 127:
                                    v17 += v1
                                }
                            case "":
                                v12 = append(v12, fi25)
                            }
                            _ = v30
                        } else {
                            var v37 string = (("\n\t" + "hello") + "")
                            for fi38 := 0; fi38 < 13; fi38++ {
                                v15 += (v15 + (v15 & v15))
                                v39 := GenS2{F0: (!v29), F1: (v1 << ((2) & 63)), F2: 1}
                                v3 -= v3
                                _ = fi38
                                _ = v39
                            }
                            _ = v37
                        }
                    }
                    _ = v29
                }
                var v40 uint16 = 595
                _ = fi25
                _ = v40
            }
        case "\n\t":
            if true {
                v41 := GenS8{F0: 138}
                switch -664 {
                case 2147483647:
                    v3 -= (v3 - (v3 - 1e10))
                case 430:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                case -9223372036854775808:
                    var v42 uint32 = 716
                    _ = v42
                default:
                    v43 := GenS5{F0: 387}
                    _ = v43
                }
                _ = v41
            }
        case "hello":
            var v44 float32 = float32(1e10)
            _ = v44
        default:
            v17 = v1
        }
        var v45 uint16 = (v15 / ((v15) | 1))
        fmt.Println(v17, v1)
        _ = v45
    } else {
        for fi46 := 0; fi46 < 10; fi46++ {
            v15 = 255
            if ((fi46 % (((^(fi46))) | 1)) >= (^(fi46))) {
                switch fi46 {
                case -9223372036854775808:
                    var v47 uint64 = 42
                    _ = v47
                case -2147483648:
                    var v48 int8 = -1
                    _ = v48
                case 32768:
                    fmt.Println(v3, v15)
                default:
                    v49 := GenS6{F0: 42, F1: v2}
                    _ = v49
                }
                fmt.Println(v3, v1)
            } else {
                switch (("\n\t" + "hello") + "\n\t") {
                case "":
                    var v50 bool = ((0.0 >= (1e10 / 1.5)) && (((int64(v2) <= 778) && false) || true))
                    _ = v50
                case "hello":
                    v2 -= v2
                default:
                    var v51 int32 = v2
                    _ = v51
                }
                v2 -= (v2 / ((353) | 1))
                for fi52 := 0; fi52 < 4; fi52++ {
                    v53 := GenS4{F0: 725}
                    _ = fi52
                    _ = v53
                }
            }
            var v54 uint8 = 162
            _ = fi46
            _ = v54
        }
        var v55 string = ("a" + "")
        if (v15 > v15) {
            var v56 int = -9223372036854775808
            _ = v56
        }
        _ = v55
    }
    for fi57 := 0; fi57 < 2; fi57++ {
        if false {
            for fi58 := 0; fi58 < 9; fi58++ {
                v17 += (v1 << ((18446744073709551615) & 63))
                v15 = v15
                _ = fi58
            }
            v17 = -484
        }
        _ = fi57
    }
    if (((float32(-2.25) + ((1.5 - (v3 + (1.5 / (float32(3.14159) * float32(((float64(v3) * 0.0) + 0.1)))))) - 0.1)) / v3) < v3) {
        v2 = 42
        v12 = append(v12, -1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
