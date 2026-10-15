package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 int32; F1 int8 }
func genFn0(p0 uint8) uint32 {
    var v7 uint = 892
    v8 := GenS1{F0: 32767, F1: 65535}
    p0 = (p0 | p0)
    fmt.Println(p0, p0)
    if (v7 < v7) {
        var v9 bool = (!false)
        var v10 int32 = 809
        for fi11 := 0; fi11 < 3; fi11++ {
            p0 = (p0 | (^(p0)))
            v10 = 0
            v12 := GenS1{F0: 32767, F1: 2}
            _ = fi11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := []int{0}
    var v2 int8 = 100
    v3 := make(map[uint32]uint32)
    v4 := []uint32{116, 4294967295, 219, 42}
    var v5 uint8 = 127
    var v6 uint64 = 1
    if (("\n\t" + "") <= ("\n\t" + "a")) {
        var v13 uint32 = 65535
        v5 = uint8(v13)
        for fi14 := 0; fi14 < 1; fi14++ {
            v1 = append(v1, (fi14 ^ (fi14 & fi14)))
            _ = fi14
        }
        _ = v13
    }
    switch (("\n\t" + "a") + "foo bar") {
    case "hello":
        v5 += v5
    case "":
        fmt.Println(v6, v2, v6)
    default:
        if (0 < 42) {
            v15 := GenS2{F0: "foo bar", F1: "hello"}
            var v16 uint64 = 42
            fmt.Println(v6, v6, v5)
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        if true {
            var v18 bool = false
            v19 := GenS1{F0: 42, F1: 4294967295}
            v2 = (^(v2))
            _ = v18
            _ = v19
        } else {
            if false {
                switch (fi17 >> ((42) & 63)) {
                case 478:
                    v2 = v2
                }
                for fi20 := 0; fi20 < 20; fi20++ {
                    fmt.Println(v2)
                    _ = fi20
                }
            }
            var v21 uint16 = 303
            v4 = append(v4, 2)
            _ = v21
        }
        fmt.Println(v6)
        _ = fi17
    }
    v6 -= (v6 * (v6 << ((255) & 63)))
    v5 += v5
    _ = len(v3)
    switch "a" {
    case "":
        if len(v4) > 0 {
            v4[0] = genFn0((v5 / (((v5 / ((128) | 1))) | 1)))
        }
    case "foo bar":
        v5 = (v5 & v5)
    }
    v22 := GenS2{F0: "hello", F1: "foo bar"}
    for fi23 := 0; fi23 < 18; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            switch ("" + "a") {
            case "\n\t":
                v5 = 42
            case "":
                switch "\n\t" {
                case "":
                    delete(v3, 42)
                case "\n\t":
                    var v25 uint64 = 304
                    if ((-(v2)) == v2) {
                        if false {
                            if (199 <= (^(v5))) {
                                _ = len(v1)
                                v2 = int8(v5)
                                fmt.Println(v25, fi24, v25)
                                v6 = (v6 ^ 42)
                                var v26 int8 = (^(v2))
                                switch fi24 {
                                case 2147483647:
                                    var v27 int16 = -813
                                    _ = v27
                                case -256:
                                    switch "" {
                                    case "\n\t":
                                        v28 := GenS3{F0: 497, F1: 0}
                                        _ = v28
                                    case "":
                                        var v29 uint32 = 2147483647
                                        _ = v29
                                    default:
                                        v26 = 127
                                    }
                                default:
                                    v25 -= v6
                                }
                                if false {
                                    v5 += 128
                                    fmt.Println(fi23, v26)
                                    fmt.Println(fi23)
                                    v5 = 42
                                } else {
                                    fmt.Println(v5, v25)
                                    v2 = 0
                                    switch fi24 {
                                    case -984:
                                        fmt.Println(v5, v6, v2)
                                    case 4294967295:
                                        var v30 float64 = 1.5
                                        _ = v30
                                    default:
                                        v26 = -1
                                    }
                                }
                                _ = v26
                            }
                            var v31 uint16 = 65535
                            _ = v31
                        } else {
                            var v32 float64 = float64(float32(3.14159))
                            v6 = (v25 | v25)
                            v5 -= 70
                            _ = v32
                        }
                        fmt.Println(v2, v6, v25)
                        switch (((("\n\t" + "hello") + ("\n\t" + "hello")) + ("a" + ("a" + ""))) + "foo bar") {
                        case "\n\t":
                            fmt.Println(v5, v25, fi23)
                        case "hello":
                            v5 += 80
                        default:
                            v33 := GenS4{F0: -302, F1: (v6 << ((493) & 63))}
                            _ = v33
                        }
                    } else {
                        fmt.Println(v2, v5)
                        var v34 int8 = v2
                        delete(v3, 928)
                        _ = v34
                    }
                    var v35 int = fi24
                    _ = v25
                    _ = v35
                case "foo bar":
                    v36 := GenS3{F0: 1, F1: -889}
                    _ = v36
                }
            case "a":
                switch (-(fi24)) {
                case 473:
                    switch fi23 {
                    case 42:
                        var v37 uint64 = (v6 - (v6 & (v6 ^ v6)))
                        _ = v37
                    case -548:
                        v38 := GenS5{F0: 0, F1: v6, F2: v5}
                        _ = v38
                    case -286:
                        fmt.Println(v5, v2)
                    }
                case 32768:
                    fmt.Println(v5, fi23)
                default:
                    if false {
                        if true {
                            v6 = v6
                        } else {
                            v2 = int8(fi23)
                            var v39 int64 = -872
                            v6 = v6
                            delete(v3, genFn0(v5))
                            _ = v39
                        }
                    }
                }
            default:
                var v40 int = (fi23 - (fi24 * fi24))
                fmt.Println(v6, fi23, v6)
                v41 := v3[590]
                _ = v40
                _ = v41
            }
            v3[uint32(v5)] = 1
            _ = fi24
        }
        for fi42 := 0; fi42 < 19; fi42++ {
            switch "\n\t" {
            case "hello":
                var v43 uint32 = 4294967295
                _ = v43
            case "":
                v5 += 127
            default:
                v44 := GenS6{F0: ("" < "hello"), F1: "foo bar"}
                _ = v44
            }
            _ = fi42
        }
        if ((116 == 0) && true) {
            v45 := v3[uint32(v6)]
            v46 := GenS7{F0: -2147483648, F1: (v2 * v2)}
            _ = v45
            _ = v46
        } else {
            v4 = append(v4, 603)
            for fi47 := 0; fi47 < 18; fi47++ {
                v5 += (v5 >> ((42) & 63))
                fmt.Println(fi23, v5)
                var v48 uint16 = 32767
                _ = fi47
                _ = v48
            }
            if len(v4) > 0 {
                v4[0] = genFn0((v5 << ((2) & 63)))
            }
        }
        _ = fi23
    }
    for fi49 := 0; fi49 < 3; fi49++ {
        var v50 uint64 = v6
        v51 := GenS7{F0: 1, F1: 42}
        _ = fi49
        _ = v50
        _ = v51
    }
    v4 = append(v4, 42)
    v5 += (^(v5))
    for fi52 := 0; fi52 < 11; fi52++ {
        v6 += (v6 + 2)
        var v53 int = (^(fi52))
        if (!(((uint16(v6) <= uint16(v5)) || false) || (729 > 481))) {
            v53 -= fi52
            if len(v1) > 0 {
                v1[0] = fi52
            }
            v53 -= -2147483648
        }
        _ = fi52
        _ = v53
    }
    fmt.Println(v2, v5, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
