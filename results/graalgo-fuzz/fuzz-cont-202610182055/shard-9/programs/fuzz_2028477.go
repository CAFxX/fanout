package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int16; F1 string; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16; F1 uint8; F2 int32 }
type GenS8 struct { F0 string; F1 uint8; F2 uint64 }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 int16; F1 int32; F2 uint64 }
func genFn0(p0 int8) int32 {
    switch "foo bar" {
    case "foo bar":
        p0 = (p0 | (p0 ^ p0))
    case "":
        p0 = (p0 + p0)
    default:
        for fi4 := 0; fi4 < 19; fi4++ {
            for fi5 := 0; fi5 < 17; fi5++ {
                switch fi5 {
                case 955:
                    if false {
                        v6 := GenS1{F0: (fi4 * (fi4 % ((fi4) | 1))), F1: 215}
                        p0 -= p0
                        p0 = (p0 % ((p0) | 1))
                        _ = v6
                    } else {
                        v7 := GenS2{F0: 42, F1: "a", F2: 255}
                        switch fi5 {
                        case -355:
                            v8 := GenS3{F0: 510, F1: int16(fi5), F2: 65535}
                            _ = v8
                        default:
                            var v9 int = (^(fi4))
                            _ = v9
                        }
                        v10 := []int{(fi5 - -9223372036854775808), (fi4 & 128)}
                        _ = v10
                        _ = v7
                    }
                case 127:
                    fmt.Println(fi4, fi4)
                case -378:
                    fmt.Println(p0, fi4)
                }
                _ = fi5
            }
            var v11 int32 = 15
            _ = fi4
            _ = v11
        }
    }
    fmt.Println(p0)
    fmt.Println(p0)
    var v12 uint8 = 127
    _ = p0
    _ = v12
    return 736
}

func main() {
    var v1 uint32 = 679
    v2 := []string{("\n\t" + ((("" + "hello") + "foo bar") + ("hello" + "\n\t"))), "", "", "foo bar"}
    v3 := []uint32{2}
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(fi13, fi13)
        _ = fi13
    }
    v14 := GenS4{F0: true, F1: uint64(v1)}
    v1 = 14
    var v15 int32 = 1
    if (100 <= -72) {
        v1 = (v1 + (v1 % (((v1 / ((v1) | 1))) | 1)))
        for fi16 := 0; fi16 < 15; fi16++ {
            if (!(!false)) {
                if true {
                    switch 127 {
                    case -1:
                        fmt.Println(v1, v1)
                    case 0:
                        v15 = 1
                    case 32768:
                        if (127 <= 89) {
                            var v17 uint32 = v1
                            fmt.Println(v1)
                            v15 += genFn0(-110)
                            _ = v17
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            for fi18 := 0; fi18 < 13; fi18++ {
                                v19 := GenS3{F0: 166, F1: 580, F2: 58}
                                v20 := GenS5{F0: -32768, F1: 18446744073709551615}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                        }
                    }
                }
            }
            _ = fi16
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := GenS6{F0: fi21}
            for fi23 := 0; fi23 < 10; fi23++ {
                v1 = v1
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    } else {
        switch 127 {
        case 547:
            for fi24 := 0; fi24 < 6; fi24++ {
                if len(v2) > 0 {
                    v2[0] = (("" + "\n\t") + "hello")
                }
                v15 += (v15 * (v15 + genFn0(100)))
                v25 := GenS7{F0: 1, F1: 1, F2: v15}
                _ = fi24
                _ = v25
            }
        case -1:
            var v26 uint = 2
            _ = v26
        default:
            var v27 string = ("hello" + "foo bar")
            _ = v27
        }
        if true {
            if true {
                if true {
                    v1 -= v1
                } else {
                    fmt.Println(v15)
                    for fi28 := 0; fi28 < 20; fi28++ {
                        var v29 float32 = 0.1
                        switch "a" {
                        case "\n\t":
                            for fi30 := 0; fi30 < 2; fi30++ {
                                v31 := GenS8{F0: ("a" + (("a" + "") + "a")), F1: 116, F2: 75}
                                var v32 int = -9223372036854775808
                                if false {
                                    fmt.Println(fi28)
                                    v29 = 1e10
                                } else {
                                    if (false || false) {
                                        fmt.Println(v32, v15, v15)
                                        v15 += -2147483648
                                    } else {
                                        v29 = (v29 / v29)
                                        var v33 int8 = 1
                                        v15 -= v15
                                        fmt.Println(v15, fi30, fi30)
                                        var v34 uint32 = (v1 * (v1 & v1))
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        case "a":
                            v35 := GenS9{F0: (fi28 ^ int(v1)), F1: fi28}
                            _ = v35
                        default:
                            var v36 uint16 = 66
                            _ = v36
                        }
                        _ = fi28
                        _ = v29
                    }
                }
                v1 = 2
            } else {
                var v37 int = 127
                _ = v37
            }
            v1 = (v1 - (^(v1)))
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        for fi39 := 0; fi39 < 18; fi39++ {
            if false {
                v15 = v15
                v15 += v15
            }
            _ = fi39
        }
        fmt.Println(v1, fi38, fi38)
        _ = fi38
    }
    v40 := GenS10{F0: 1, F1: v15, F2: 962}
    for fi41 := 0; fi41 < 2; fi41++ {
        switch "foo bar" {
        case "foo bar":
            if true {
                switch "hello" {
                case "a":
                    v3 = append(v3, (^(v1)))
                default:
                    for fi42 := 0; fi42 < 8; fi42++ {
                        fmt.Println(fi41, fi42)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi42
                    }
                }
            }
        case "\n\t":
            for fi43 := 0; fi43 < 16; fi43++ {
                v15 += (v15 % (((v15 & (v15 / ((v15) | 1)))) | 1))
                _ = fi43
            }
        default:
            fmt.Println(v1)
        }
        _ = fi41
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v40
}
