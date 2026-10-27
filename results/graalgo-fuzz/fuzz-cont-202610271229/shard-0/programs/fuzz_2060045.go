package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 bool; F2 string }
type GenS5 struct { F0 uint; F1 uint; F2 int64 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint16; F1 int32; F2 int64 }
func genFn0() bool {
    var v4 bool = (9223372036854775807 > 42)
    switch (((("hello" + ("" + ((("" + "foo bar") + ("\n\t" + "hello")) + ("hello" + ("foo bar" + ((("foo bar" + (("" + ("hello" + "a")) + "a")) + "hello") + ("hello" + "hello"))))))) + (("foo bar" + "foo bar") + "")) + ("a" + ("a" + "foo bar"))) + ("a" + "")) {
    case "a":
        for fi5 := 0; fi5 < 14; fi5++ {
            if v4 {
                var v6 bool = (32768 != 545)
                switch fi5 {
                case -994:
                    var v7 uint32 = 2147483647
                    _ = v7
                default:
                    switch (fi5 / ((fi5) | 1)) {
                    case 32768:
                        v4 = false
                    default:
                        v8 := []int{fi5}
                        _ = v8
                    }
                }
                _ = v6
            }
            v9 := GenS1{F0: 2147483648, F1: 2, F2: 32767}
            _ = fi5
            _ = v9
        }
    default:
        v10 := GenS2{F0: 1}
        _ = v10
    }
    v11 := []int{-424}
    _ = v11
    _ = v4
    return v4
}

func main() {
    var v1 uint = 65535
    v2 := []int{}
    v3 := make(map[int]int)
    v12 := GenS3{F0: "hello"}
    if true {
        var v13 uint = (v1 ^ v1)
        v14 := GenS4{F0: true, F1: (("\n\t" + (("foo bar" + "a") + ("\n\t" + (("foo bar" + (("hello" + "\n\t") + ((("\n\t" + ("hello" + ("" + "foo bar"))) + "a") + "a"))) + "\n\t")))) > "\n\t"), F2: (("foo bar" + "a") + ("\n\t" + ""))}
        for fi15 := 0; fi15 < 14; fi15++ {
            for fi16 := 0; fi16 < 10; fi16++ {
                fmt.Println(fi15, fi15)
                v1 = v13
                if false {
                    if (((1.5 * 1e10) < -2.25) && false) {
                        switch ("hello" + "a") {
                        case "\n\t":
                            if (42 <= -706) {
                                switch -621 {
                                case 42:
                                    var v17 float64 = float64(-2.25)
                                    _ = v17
                                case -788:
                                    if (!(false || (true && false))) {
                                        var v18 uint8 = 183
                                        _ = v18
                                    }
                                default:
                                    v19 := GenS5{F0: v13, F1: (v1 % ((v13) | 1)), F2: 0}
                                    _ = v19
                                }
                            } else {
                                fmt.Println(fi16, fi16)
                            }
                        case "a":
                            fmt.Println(v13, fi15)
                        default:
                            var v20 string = "foo bar"
                            _ = v20
                        }
                    } else {
                        if (("\n\t" + ((("hello" + "a") + "hello") + "foo bar")) <= "hello") {
                            v21 := GenS6{F0: 1000, F1: 128}
                            v1 -= v1
                            _ = v21
                        }
                        switch fi15 {
                        case 430:
                            var v22 uint16 = 221
                            v3[(fi15 << (((v1 - (v1 % ((v13) | 1)))) & 63))] = (fi15 / (((fi15 - 1)) | 1))
                            _ = v22
                        default:
                            if ((fi16 ^ fi15) != fi15) {
                                if false {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if true {
                                        var v23 bool = true
                                        _ = v23
                                    }
                                    var v24 float64 = -2.25
                                    _ = v24
                                } else {
                                    var v25 float64 = 3.14159
                                    v26 := GenS7{F0: 2, F1: 32767}
                                    _ = v25
                                    _ = v26
                                }
                                v27 := GenS8{F0: (^(v1))}
                                v13 = (v13 % (((^(v1))) | 1))
                                _ = v27
                            } else {
                                v28 := v3[(^(fi16))]
                                var v29 float64 = 1e10
                                _ = v28
                                _ = v29
                            }
                            switch fi16 {
                            case 9223372036854775807:
                                if (!((65535 < 0) || (1 == 42))) {
                                    if (!false) {
                                        fmt.Println(fi16)
                                        switch "hello" {
                                        case "":
                                            v30 := GenS9{F0: 1}
                                            _ = v30
                                        case "foo bar":
                                            v1 += v1
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    switch fi15 {
                                    case 0:
                                        v1 += v1
                                    case 2147483647:
                                        v13 += v13
                                    case 1:
                                        v1 += v1
                                        v2 = append(v2, -331)
                                    default:
                                        var v31 float64 = -2.25
                                        _ = v31
                                    }
                                    v32 := GenS10{F0: 32767, F1: -722, F2: 128}
                                    _ = v32
                                }
                            default:
                                if false {
                                    fmt.Println(fi15)
                                    v1 = (v13 | 557)
                                }
                            }
                        }
                    }
                    var v33 uint64 = 4294967295
                    _ = v33
                } else {
                    v13 += 18446744073709551615
                }
                _ = fi16
            }
            _ = fi15
        }
        _ = v13
        _ = v14
    } else {
        for fi34 := 0; fi34 < 2; fi34++ {
            fmt.Println(v1, fi34)
            v1 = (v1 << ((v1) & 63))
            _ = fi34
        }
        for fi35 := 0; fi35 < 14; fi35++ {
            v3[(fi35 | (fi35 + (^(fi35))))] = (fi35 >> ((4294967295) & 63))
            fmt.Println(v1)
            _ = fi35
        }
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        for fi37 := 0; fi37 < 13; fi37++ {
            var v38 int64 = 1
            var v39 float64 = 3.14159
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = fi36
    }
    v1 += 42
    fmt.Println(v1)
    fmt.Println(v1)
    v1 -= v1
    for fi40 := 0; fi40 < 18; fi40++ {
        v1 = v1
        fmt.Println(fi40, fi40)
        _ = fi40
    }
    fmt.Println(v1)
    for fi41 := 0; fi41 < 8; fi41++ {
        if false {
            v42 := GenS1{F0: 4294967295, F1: 65535, F2: 32767}
            v1 = (v1 - v1)
            v1 += 22
            _ = v42
        } else {
            var v43 int8 = 1
            _ = v43
        }
        fmt.Println(v1, v1)
        for fi44 := 0; fi44 < 20; fi44++ {
            var v45 uint32 = 2147483648
            _ = fi44
            _ = v45
        }
        _ = fi41
    }
    v46 := v3[0]
    var v47 int8 = 15
    v46 += (-(v46))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v46
    _ = v47
}
