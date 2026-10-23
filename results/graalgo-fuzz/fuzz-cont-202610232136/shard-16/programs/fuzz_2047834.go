package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 bool; F1 int8; F2 int16 }
type GenS8 struct { F0 uint16; F1 uint }
type GenS9 struct { F0 int64; F1 string }
type GenS10 struct { F0 int64; F1 bool }
type GenS11 struct { F0 int16; F1 int }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int; F1 uint64; F2 uint16 }
type GenS14 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS15 struct { F0 int32; F1 int8 }
func genFn0() bool {
    v4 := []int{-9223372036854775808, 128, 4294967295, -57}
    var v5 uint8 = 127
    for fi6 := 0; fi6 < 19; fi6++ {
        v5 += v5
        _ = fi6
    }
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 int32) uint32 {
    p0 = (p0 ^ p0)
    v7 := GenS1{F0: "a"}
    if true {
        if false {
            p0 = p0
        } else {
            fmt.Println(p0)
        }
        if ("foo bar" != "foo bar") {
            switch "\n\t" {
            case "a":
                for fi8 := 0; fi8 < 16; fi8++ {
                    p0 += 2147483647
                    p0 -= -2147483648
                    fmt.Println(p0, fi8)
                    _ = fi8
                }
            case "":
                var v9 float32 = -2.25
                _ = v9
            }
            v10 := GenS1{F0: ("hello" + "a")}
            if false {
                if (p0 != p0) {
                    p0 = 0
                    p0 = (p0 & (p0 & p0))
                } else {
                    for fi11 := 0; fi11 < 9; fi11++ {
                        switch ("\n\t" + ("" + ("" + "hello"))) {
                        case "a":
                            p0 -= (p0 | 2147483647)
                        case "foo bar":
                            fmt.Println(p0, p0)
                        case "hello":
                            var v12 bool = false
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 15; fi13++ {
                            fmt.Println(p0, p0, fi13)
                            if (true || true) {
                                v14 := GenS2{F0: 825}
                                switch "hello" {
                                case "hello":
                                    if (-32768 > -252) {
                                        var v15 int = 32767
                                        _ = v15
                                    }
                                case "\n\t":
                                    v16 := []int{(fi11 * 492), fi11, (-(fi11))}
                                    _ = v16
                                case "foo bar":
                                    fmt.Println(fi13, fi13)
                                    var v17 float64 = 1e10
                                    v18 := GenS3{F0: 66, F1: (((false || (0 > -394)) || false) && true)}
                                    fmt.Println(p0, p0)
                                    _ = v17
                                    _ = v18
                                }
                                _ = v14
                            }
                            p0 = p0
                            _ = fi13
                        }
                        _ = fi11
                    }
                }
                v19 := GenS4{F0: 0, F1: -9223372036854775808}
                if (0.0 == ((0.0 - (float64(float32(-2.25)) / 1.0)) - -2.25)) {
                    if (0 > 32767) {
                        v20 := []int{4294967295, 1, -9223372036854775808}
                        if (2147483647 == 636) {
                            fmt.Println(p0)
                            if (!(57 == 18446744073709551615)) {
                                p0 -= p0
                                v21 := GenS4{F0: 4294967295, F1: 128}
                                switch "a" {
                                case "":
                                    if (((3.14159 / (1.5 + 1e10)) - (-2.25 / float64(float32(-2.25)))) == (0.1 + 0.1)) {
                                        p0 = 862
                                        var v22 float64 = 0.1
                                        v23 := GenS5{F0: 42, F1: 127}
                                        _ = v22
                                        _ = v23
                                    }
                                default:
                                    if len(v20) > 0 {
                                        v20[0] = 128
                                    }
                                }
                                _ = v21
                            }
                        }
                        _ = v20
                    } else {
                        var v24 float64 = 0.0
                        if (255 > 2) {
                            v25 := GenS5{F0: 1, F1: 127}
                            _ = v25
                        }
                        var v26 int64 = 42
                        _ = v24
                        _ = v26
                    }
                } else {
                    p0 = -647
                    for fi27 := 0; fi27 < 1; fi27++ {
                        for fi28 := 0; fi28 < 1; fi28++ {
                            v29 := GenS1{F0: "a"}
                            switch "a" {
                            case "a":
                                fmt.Println(fi27, fi27, fi28)
                            case "\n\t":
                                p0 = p0
                            case "hello":
                                fmt.Println(fi27, p0)
                            }
                            v30 := GenS5{F0: 2, F1: 91}
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        _ = fi27
                    }
                    v31 := GenS6{F0: 42, F1: 128}
                    _ = v31
                }
                _ = v19
            }
            _ = v10
        } else {
            var v32 int16 = 432
            v33 := GenS6{F0: int32(v32), F1: 348}
            _ = v32
            _ = v33
        }
    } else {
        v34 := GenS6{F0: -1, F1: 32767}
        if true {
            var v35 float64 = 1e10
            _ = v35
        } else {
            v36 := GenS4{F0: 848, F1: 42}
            var v37 float64 = 1.5
            _ = v36
            _ = v37
        }
        _ = v34
    }
    fmt.Println(p0)
    p0 = (p0 ^ (p0 * p0))
    _ = p0
    _ = v7
    return 2147483648
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -1
    var v3 uint16 = 540
    v38 := v1[0]
    v39 := GenS7{F0: false, F1: 111, F2: v2}
    v38 = ((v38 + ("" + v38)) + (("foo bar" + "") + v38))
    switch "" {
    case "":
        fmt.Println(v3, v3, v3)
    case "foo bar":
        switch 32767 {
        case 0:
            for fi40 := 0; fi40 < 20; fi40++ {
                for fi41 := 0; fi41 < 1; fi41++ {
                    v2 -= 1000
                    v42 := GenS3{F0: 48, F1: true}
                    _ = fi41
                    _ = v42
                }
                v38 = (v38 + (v38 + v38))
                var v43 uint8 = 0
                _ = fi40
                _ = v43
            }
        case -860:
            v44 := GenS8{F0: v3, F1: 42}
            _ = v44
        }
    }
    var v45 int = 1
    switch (v38 + v38) {
    case "hello":
        v46 := GenS9{F0: 9223372036854775807, F1: ("hello" + (v38 + "\n\t"))}
        _ = v46
    }
    v3 = 65535
    var v47 uint32 = 357
    switch (v45 << ((1) & 63)) {
    case 4294967295:
        v2 = (v2 % ((v2) | 1))
    case 32767:
        var v48 uint16 = (v3 & 65535)
        _ = v48
    default:
        for fi49 := 0; fi49 < 14; fi49++ {
            for fi50 := 0; fi50 < 17; fi50++ {
                v3 += (v3 >> ((4294967295) & 63))
                _ = fi50
            }
            fmt.Println(v3)
            _ = fi49
        }
    }
    v51 := []int{v45}
    if (32767 < (v3 | (v3 / (((v3 >> ((uint(v3)) & 63))) | 1)))) {
        v45 += -822
        var v52 string = v38
        _ = v52
    } else {
        for fi53 := 0; fi53 < 10; fi53++ {
            v51 = append(v51, v45)
            var v54 int64 = -1
            v2 = v2
            _ = fi53
            _ = v54
        }
        switch "foo bar" {
        case "hello":
            if len(v51) >= 2 {
                v51 = v51[1:len(v51)]
            }
        case "\n\t":
            if (((!(true && true)) && true) && false) {
                var v55 float64 = -2.25
                for fi56 := 0; fi56 < 17; fi56++ {
                    v57 := GenS1{F0: v38}
                    var v58 int32 = 42
                    for fi59 := 0; fi59 < 18; fi59++ {
                        if len(v51) >= 2 {
                            v51 = v51[1:len(v51)]
                        }
                        var v60 int32 = (v58 + v58)
                        _ = fi59
                        _ = v60
                    }
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                delete(v1, v45)
                _ = v55
            }
        }
    }
    v61 := GenS8{F0: uint16(v47), F1: 606}
    for fi62 := 0; fi62 < 2; fi62++ {
        _ = len(v1)
        var v63 int16 = v2
        _ = fi62
        _ = v63
    }
    v45 = 391
    var v64 uint16 = (v3 >> ((255) & 63))
    for fi65 := 0; fi65 < 12; fi65++ {
        switch "\n\t" {
        case "":
            if (!false) {
                v1[(fi65 >> ((960) & 63))] = "\n\t"
                var v66 float32 = ((1.5 + 1.5) + ((0.0 - 0.0) / 1.0))
                if true {
                    v45 += fi65
                    v67 := GenS10{F0: 202, F1: false}
                    _ = v67
                }
                _ = v66
            } else {
                for fi68 := 0; fi68 < 6; fi68++ {
                    switch "\n\t" {
                    case "a":
                        v69 := GenS4{F0: (v47 & (v47 / ((618) | 1))), F1: 32767}
                        _ = v69
                    case "hello":
                        var v70 int32 = 792
                        _ = v70
                    case "\n\t":
                        var v71 uint32 = v47
                        _ = v71
                    default:
                        switch (v38 + (v38 + "foo bar")) {
                        case "":
                            v1[(^(v45))] = v38
                        default:
                            _ = len(v51)
                        }
                    }
                    if true {
                        v72 := GenS11{F0: v2, F1: fi68}
                        _ = v72
                    } else {
                        v1[-28] = ("" + "a")
                        v73 := GenS12{F0: v64}
                        _ = v73
                    }
                    v38 = (v38 + "")
                    _ = fi68
                }
                for fi74 := 0; fi74 < 10; fi74++ {
                    v75 := GenS13{F0: fi65, F1: 4294967295, F2: (v3 / ((v64) | 1))}
                    if true {
                        if true {
                            v45 = fi65
                            v76 := GenS14{F0: 65535, F1: 76, F2: 0}
                            var v77 float32 = (((1e10 + (((((1e10 / 1.0) - (-2.25 + 1e10)) - 0.1) + (0.1 - 1.5)) - 3.14159)) - 0.0) * 1e10)
                            _ = v76
                            _ = v77
                        }
                        var v78 bool = ((("\n\t" + (v38 + "a")) + v38) > v38)
                        _ = v78
                    } else {
                        var v79 bool = (2 < 112)
                        v80 := GenS15{F0: -738, F1: 0}
                        _ = len(v51)
                        _ = v79
                        _ = v80
                    }
                    var v81 int16 = (v2 / (((v2 & (v2 | -1))) | 1))
                    _ = fi74
                    _ = v75
                    _ = v81
                }
            }
        }
        fmt.Println(v45)
        _ = fi65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v45
    _ = v47
    _ = v51
    _ = v61
    _ = v64
}
