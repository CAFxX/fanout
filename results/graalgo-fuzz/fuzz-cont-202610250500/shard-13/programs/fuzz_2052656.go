package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int8; F1 bool; F2 bool }
type GenS5 struct { F0 bool; F1 string; F2 uint16 }
type GenS6 struct { F0 uint16; F1 int64; F2 int8 }
type GenS7 struct { F0 int16 }

func main() {
    var v1 float64 = (1e10 + 1.5)
    v2 := []uint32{920}
    v3 := make(map[string]uint32)
    v4 := []string{"hello", "\n\t"}
    v1 += v1
    switch 32767 {
    case 0:
        fmt.Println(v1)
    case 1:
        fmt.Println(v1)
    case 32767:
        _ = len(v2)
    }
    var v5 int16 = 32767
    v1 = 0.1
    v1 = (v1 * v1)
    switch 448 {
    case 128:
        var v6 float64 = (-2.25 / ((((0.1 - v1) - v1) - (v1 + float64(((1.5 * 0.0) / 1.0)))) - (v1 * 3.14159)))
        _ = v6
    case 1:
        var v7 bool = (true && false)
        _ = v7
    default:
        var v8 string = "hello"
        _ = v8
    }
    var v9 float64 = v1
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v10 := GenS1{F0: false}
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := v3[""]
        fmt.Println(v9, v12, v12)
        var v13 int16 = 1000
        _ = fi11
        _ = v12
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 19; fi14++ {
            if (0 <= 393) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v5 -= (v5 | 1)
                if (0 <= 2) {
                    v15 := v3[""]
                    _ = v15
                }
            } else {
                if (18446744073709551615 <= 956) {
                    v16 := GenS2{F0: 87, F1: 18446744073709551615, F2: 32767}
                    v5 -= v5
                    v1 -= v9
                    _ = v16
                } else {
                    v5 -= (-(v5))
                    var v17 float32 = float32((v9 / float64(((1e10 - 0.0) / float32((v1 + float64((0.0 - 1e10))))))))
                    v3[""] = 65535
                    _ = v17
                }
            }
            v18 := GenS1{F0: (true || ((v5 << ((4294967295) & 63)) == v5))}
            _ = fi14
            _ = v18
        }
        var v19 float64 = v9
        v20 := GenS1{F0: true}
        _ = v19
        _ = v20
    }
    switch "\n\t" {
    case "":
        if true {
            if ((v5 ^ (v5 + (^(v5)))) == v5) {
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint16 = 347
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if false {
                            var v24 uint32 = 4294967295
                            _ = v24
                        }
                        var v25 float32 = 0.1
                        _ = fi23
                        _ = v25
                    }
                    _ = fi21
                    _ = v22
                }
                for fi26 := 0; fi26 < 1; fi26++ {
                    v27 := GenS1{F0: false}
                    _ = fi26
                    _ = v27
                }
                v28 := GenS1{F0: ((v5 - (v5 - (^(v5)))) > (v5 | (v5 ^ (v5 << ((2) & 63)))))}
                _ = v28
            }
        } else {
            if (!false) {
                var v29 float32 = (3.14159 - float32(v1))
                v30 := GenS1{F0: (!true)}
                _ = len(v2)
                _ = v29
                _ = v30
            } else {
                for fi31 := 0; fi31 < 6; fi31++ {
                    if false {
                        var v32 int32 = 2147483647
                        for fi33 := 0; fi33 < 10; fi33++ {
                            v34 := GenS3{F0: 191, F1: 255}
                            _ = fi33
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 11; fi35++ {
                            delete(v3, ("a" + ""))
                            var v36 uint32 = 0
                            _ = fi35
                            _ = v36
                        }
                        _ = v32
                    }
                    for fi37 := 0; fi37 < 9; fi37++ {
                        v38 := GenS2{F0: 1, F1: 4294967295, F2: 0}
                        switch (fi31 << ((61) & 63)) {
                        case -2147483648:
                            fmt.Println(fi31, v5, fi37)
                        case -941:
                            v1 -= ((v1 * v9) - v1)
                        default:
                            if (0 != 436) {
                                v39 := GenS4{F0: -128, F1: ((127 != 114) && (100 == 127)), F2: false}
                                v5 = (v5 / (((v5 / (((v5 * 1)) | 1))) | 1))
                                _ = v39
                            }
                        }
                        v1 = (0.0 + ((-2.25 - float64(0.0)) + v9))
                        _ = fi37
                        _ = v38
                    }
                    _ = fi31
                }
                switch "hello" {
                case "hello":
                    fmt.Println(v5, v9)
                case "":
                    v1 += ((1e10 / (v9 * v1)) + float64((float32(v1) - 3.14159)))
                default:
                    v40 := GenS1{F0: false}
                    _ = v40
                }
                if false {
                    fmt.Println(v9, v5)
                    v3["hello"] = 4294967295
                    if (!false) {
                        var v41 bool = true
                        _ = len(v3)
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 16; fi42++ {
                            for fi43 := 0; fi43 < 4; fi43++ {
                                if true {
                                    switch "hello" {
                                    case "\n\t":
                                        var v44 string = "foo bar"
                                        _ = v44
                                    case "hello":
                                        var v45 uint = 216
                                        _ = v45
                                    case "":
                                        v46 := GenS5{F0: ((v5 + v5) == v5), F1: "foo bar", F2: 32767}
                                        _ = v46
                                    default:
                                        var v47 int32 = 42
                                        _ = v47
                                    }
                                    if false {
                                        v48 := GenS1{F0: false}
                                        var v49 int64 = 42
                                        _ = v48
                                        _ = v49
                                    }
                                } else {
                                    v50 := GenS1{F0: false}
                                    _ = v50
                                }
                                v51 := GenS2{F0: 18446744073709551615, F1: 0, F2: 32767}
                                var v52 float64 = v1
                                fmt.Println(v9)
                                v53 := GenS6{F0: 513, F1: 650, F2: 1}
                                _ = fi43
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                            for fi54 := 0; fi54 < 18; fi54++ {
                                var v55 int16 = v5
                                fmt.Println(fi42, v9)
                                v55 -= v5
                                _ = fi54
                                _ = v55
                            }
                            if (1 != 1) {
                                var v56 int8 = 90
                                _ = v56
                            }
                            _ = fi42
                        }
                        v2 = append(v2, 4294967295)
                    }
                }
            }
            if ((("" + ("" + "foo bar")) + "\n\t") != "hello") {
                v57 := GenS2{F0: 18446744073709551615, F1: 4294967295, F2: 931}
                _ = v57
            }
        }
    }
    v58 := GenS2{F0: 4294967295, F1: 4294967295, F2: 65535}
    if ((v5 * v5) == v5) {
        switch int(v5) {
        case -1:
            v5 -= (^(v5))
        case 1:
            if (!true) {
                fmt.Println(v9, v9, v5)
                var v59 uint16 = 469
                v60 := GenS7{F0: -83}
                _ = v59
                _ = v60
            }
        default:
            v3["\n\t"] = 906
        }
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v9
}
