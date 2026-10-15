package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint }
type GenS4 struct { F0 int; F1 int64; F2 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS7 struct { F0 string; F1 int }
func genFn0(p0 uint16) int8 {
    p0 = p0
    var v7 int = -633
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 uint64, p1 int32, p2 int64) int32 {
    v8 := GenS1{F0: p1, F1: 2}
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS2{F0: (p2 + (p2 >> ((18446744073709551615) & 63))), F1: (-(p1))}
        _ = fi9
        _ = v10
    }
    v11 := []int{-1, 9223372036854775807, 4294967295, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return p1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 int = -2147483648
    v4 := make(map[uint32]string)
    var v5 int16 = 801
    v6 := make(map[string]int)
    v3 = (-(v3))
    var v12 int16 = 1
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 bool = true
        fmt.Println(fi13)
        v14 = v14
        _ = fi13
        _ = v14
    }
    if true {
        var v15 uint = 255
        _ = v15
    }
    v5 += (-(v5))
    v16 := GenS3{F0: 0, F1: 42, F2: 65535}
    v17 := []int{(v3 | v3), v3}
    if true {
        switch (("" + "") + "\n\t") {
        case "a":
            var v18 uint32 = 4294967295
            _ = v18
        case "foo bar":
            fmt.Println(v3, v12, v5)
        default:
            v12 = v5
        }
    } else {
        v19 := v6["foo bar"]
        v20 := GenS4{F0: (v19 >> ((257) & 63)), F1: 9223372036854775807, F2: 212}
        _ = v19
        _ = v20
    }
    fmt.Println(v3, v5)
    v21 := GenS5{F0: (v12 >> ((538) & 63))}
    var v22 string = "\n\t"
    if false {
        var v23 string = v22
        _ = v23
    } else {
        switch "a" {
        case "\n\t":
            var v24 int8 = 127
            _ = v24
        case "a":
            for fi25 := 0; fi25 < 9; fi25++ {
                v3 = 4294967295
                for fi26 := 0; fi26 < 15; fi26++ {
                    v22 = ("foo bar" + v22)
                    _ = fi26
                }
                v22 = ""
                _ = fi25
            }
        case "foo bar":
            for fi27 := 0; fi27 < 16; fi27++ {
                if false {
                    v5 = (v12 - 32767)
                    if (4294967295 > 2) {
                        switch v3 {
                        case 32767:
                            switch fi27 {
                            case 461:
                                v28 := GenS4{F0: -2147483648, F1: 1, F2: 42}
                                _ = v28
                            case 0:
                                v5 = v12
                            case -9223372036854775808:
                                if true {
                                    v5 += v5
                                    fmt.Println(v3, v5, v22)
                                }
                            }
                        case 234:
                            var v29 int = fi27
                            _ = v29
                        default:
                            if false {
                                v22 = v22
                                for fi30 := 0; fi30 < 3; fi30++ {
                                    if (!true) {
                                        var v31 uint16 = 65535
                                        v22 = "a"
                                        _ = v31
                                    }
                                    v3 = (^(v3))
                                    var v32 float64 = (-2.25 / (1.5 / (1e10 - 1.5)))
                                    _ = fi30
                                    _ = v32
                                }
                            }
                        }
                        for fi33 := 0; fi33 < 6; fi33++ {
                            if true {
                                v12 = v5
                                var v34 uint64 = 4294967295
                                v6[v22] = fi33
                                _ = v34
                            }
                            _ = fi33
                        }
                    } else {
                        var v35 int8 = 42
                        _ = v35
                    }
                } else {
                    fmt.Println(fi27, v3, v3)
                    v22 = v22
                }
                var v36 bool = (!false)
                if v36 {
                    v37 := GenS1{F0: 337, F1: 0}
                    for fi38 := 0; fi38 < 6; fi38++ {
                        v39 := GenS6{F0: 42, F1: 18446744073709551615, F2: 2147483648}
                        delete(v6, v22)
                        v4[0] = "hello"
                        _ = fi38
                        _ = v39
                    }
                    var v40 float32 = -2.25
                    _ = v37
                    _ = v40
                }
                _ = fi27
                _ = v36
            }
        }
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        switch ((((((v22 + v22) + v22) + (v22 + v22)) + ((v22 + (v22 + v22)) + v22)) + ("a" + v22)) + ("" + ("a" + ""))) {
        case "a":
            for fi42 := 0; fi42 < 11; fi42++ {
                switch (v22 + (((v22 + v22) + ((v22 + v22) + v22)) + ("a" + (((v22 + v22) + (v22 + v22)) + "hello")))) {
                case "":
                    fmt.Println(v22)
                case "foo bar":
                    var v43 float64 = ((float64(0.1) - -2.25) + 0.1)
                    _ = v43
                }
                if ((-1 >= (v5 >> ((0) & 63))) || true) {
                    switch v3 {
                    case -9223372036854775808:
                        fmt.Println(fi42, v12)
                    case 849:
                        var v44 int16 = v5
                        _ = v44
                    case -508:
                        v45 := GenS1{F0: genFn1(1, -1, 856), F1: 0}
                        var v46 float64 = -2.25
                        _ = v45
                        _ = v46
                    }
                    v3 = (fi42 ^ (fi42 / ((fi42) | 1)))
                    v47 := GenS7{F0: v22, F1: (fi42 & fi42)}
                    _ = v47
                } else {
                    if (true && (true && true)) {
                        if len(v17) > 0 {
                            v17[0] = (-(fi41))
                        }
                        if (int8(fi41) >= 1) {
                            var v48 string = v22
                            if (!true) {
                                _ = len(v2)
                                v12 = (v12 << ((166) & 63))
                                switch ("\n\t" + v48) {
                                case "":
                                    var v49 float64 = (-2.25 + (1e10 - ((float64((1.5 + (((-2.25 / -2.25) / 3.14159) + -2.25))) * 1.5) * -2.25)))
                                    v49 += v49
                                    _ = v49
                                default:
                                    var v50 int = -28
                                    _ = v50
                                }
                            } else {
                                _ = len(v17)
                                if (-799 == 1) {
                                    v48 = v22
                                } else {
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                }
                                v17 = append(v17, (v3 | fi42))
                                fmt.Println(v48, v48)
                                if true {
                                    v3 += -9223372036854775808
                                }
                            }
                            _ = v48
                        } else {
                            v5 += (v5 % ((v5) | 1))
                        }
                    } else {
                        v22 = ((v22 + v22) + v22)
                        switch "hello" {
                        case "a":
                            _ = len(v6)
                        case "\n\t":
                            v12 -= v5
                        }
                    }
                }
                _ = fi42
            }
        case "\n\t":
            var v51 int16 = -372
            _ = v51
        }
        switch ((v22 + v22) + (v22 + v22)) {
        case "hello":
            v22 = "hello"
        default:
            var v52 int16 = v12
            _ = v52
        }
        fmt.Println(v12, v12)
        _ = fi41
    }
    fmt.Println(v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
