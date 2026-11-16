package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint32; F1 int32; F2 string }
func genFn0(p0 int) string {
    p0 = -1
    v7 := []int{}
    if (!(127 != 0)) {
        _ = len(v7)
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint8 = 1
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            if false {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 float64 = 1e10
                    p0 = fi11
                    v13 := GenS1{F0: (fi11 - (fi11 ^ fi11)), F1: ("" + ("" + "a")), F2: 728}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            switch "hello" {
            case "\n\t":
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            _ = fi10
        }
    }
    if false {
        for fi14 := 0; fi14 < 3; fi14++ {
            _ = len(v7)
            _ = fi14
        }
    }
    _ = p0
    _ = v7
    return (("a" + "hello") + ("a" + "a"))
}
func genFn1(p0 int) int {
    var v15 int8 = 27
    v16 := GenS1{F0: p0, F1: "", F2: 367}
    _ = p0
    _ = v15
    _ = v16
    return 9223372036854775807
}

func main() {
    var v1 int32 = 1
    var v2 uint = 575
    var v3 float32 = (1.5 - 1e10)
    v4 := []uint32{2147483647, 65535, 618}
    v5 := []uint32{2, 380, 0}
    var v6 uint = (^(v2))
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    switch "hello" {
    case "a":
        var v17 uint64 = 756
        _ = v17
    case "foo bar":
        var v18 uint64 = 1
        _ = v18
    default:
        switch genFn1(1) {
        case 503:
            var v19 bool = true
            _ = v19
        }
    }
    switch 9223372036854775807 {
    case 42:
        v4 = append(v4, 4294967295)
    case -1:
        var v20 string = ("a" + (("\n\t" + "a") + ("hello" + ("hello" + "a"))))
        _ = v20
    }
    v21 := GenS1{F0: 606, F1: "", F2: 2147483647}
    v6 -= (v6 & (v2 % (((v6 << (((v2 - v2)) & 63))) | 1)))
    v22 := GenS1{F0: 2147483647, F1: "", F2: 0}
    var v23 uint = v6
    var v24 float64 = (3.14159 / 1.0)
    if (!false) {
        fmt.Println(v2, v1)
        switch "foo bar" {
        case "\n\t":
            switch 128 {
            case -9223372036854775808:
                for fi25 := 0; fi25 < 13; fi25++ {
                    v24 = (float64((v3 / (float32(float64((v3 - ((float32(v24) * (float32(0.0) * v3)) * v3)))) + v3))) - 3.14159)
                    _ = fi25
                }
            case -115:
                if true {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        if (128 <= 0) {
                            for fi27 := 0; fi27 < 15; fi27++ {
                                fmt.Println(v23, v2)
                                v23 += v2
                                if (!false) {
                                    v23 -= (v2 << (((v2 >> (((v6 & (v6 << ((1) & 63)))) & 63))) & 63))
                                    v4 = append(v4, 2147483647)
                                    if ((true || (!(42 <= 0))) || true) {
                                        v28 := GenS2{F0: ""}
                                        _ = v28
                                    }
                                } else {
                                    fmt.Println(fi27)
                                    if false {
                                        var v29 float64 = v24
                                        var v30 string = "foo bar"
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v24 -= (v24 * 0.1)
                                        var v31 float64 = 3.14159
                                        if len(v5) > 0 {
                                            v5[0] = 65535
                                        }
                                        _ = v31
                                    }
                                    var v32 int = (fi26 / ((genFn1(32768)) | 1))
                                    _ = v32
                                }
                                _ = fi27
                            }
                            v33 := GenS3{F0: (v1 % ((863) | 1))}
                            v34 := GenS4{F0: (fi26 << ((v6) & 63)), F1: "foo bar"}
                            _ = v33
                            _ = v34
                        } else {
                            var v35 int8 = int8(fi26)
                            if len(v4) > 0 {
                                v4[0] = 65535
                            }
                            for fi36 := 0; fi36 < 7; fi36++ {
                                if ((fi26 >> (((^(v6))) & 63)) != -173) {
                                    v37 := GenS5{F0: -9223372036854775808, F1: 909, F2: 440}
                                    v1 = -1
                                    v3 -= (1e10 - v3)
                                    v23 -= 2
                                    if len(v5) > 0 {
                                        v5[0] = 1
                                    }
                                    _ = v37
                                }
                                var v38 uint32 = 0
                                fmt.Println(fi36, v3)
                                var v39 uint = v6
                                _ = fi36
                                _ = v38
                                _ = v39
                            }
                            _ = v35
                        }
                        v40 := GenS3{F0: v1}
                        _ = fi26
                        _ = v40
                    }
                    if len(v4) > 0 {
                        v4[0] = 635
                    }
                    v41 := GenS2{F0: (("" + "foo bar") + ("hello" + ("hello" + "hello")))}
                    _ = v41
                }
            default:
                v3 += (float32(float64(v3)) - 0.0)
            }
        case "":
            for fi42 := 0; fi42 < 6; fi42++ {
                if (42 >= v1) {
                    switch ("foo bar" + ("foo bar" + "foo bar")) {
                    case "foo bar":
                        v43 := GenS6{F0: 121, F1: v2, F2: 255}
                        _ = v43
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v44 uint = v23
                    _ = v44
                }
                _ = fi42
            }
        }
        var v45 int8 = -1
        _ = v45
    }
    v3 = (v3 / v3)
    _ = len(v5)
    switch "hello" {
    case "\n\t":
        var v46 uint64 = 1
        _ = v46
    }
    v24 = (float64((0.1 - v3)) / (0.0 / v24))
    for fi47 := 0; fi47 < 2; fi47++ {
        v48 := GenS3{F0: -2147483648}
        for fi49 := 0; fi49 < 16; fi49++ {
            fmt.Println(v23, v23)
            var v50 bool = (true && false)
            _ = fi49
            _ = v50
        }
        var v51 uint16 = 868
        _ = fi47
        _ = v48
        _ = v51
    }
    v52 := GenS7{F0: 0, F1: 42, F2: "hello"}
    var v53 int64 = 42
    v54 := GenS2{F0: ("a" + ("\n\t" + ""))}
    switch -1 {
    case -1:
        var v55 float32 = (v3 + -2.25)
        _ = v55
    default:
        var v56 uint = v6
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
