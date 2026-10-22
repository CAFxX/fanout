package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]uint32)
    var v3 bool = false
    v4 := []int{32768, -1}
    _ = len(v4)
    var v5 float32 = ((((0.0 - 1e10) * -2.25) + 0.0) / 1.5)
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int32 = 42
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if (!(2 < 719)) {
        fmt.Println(v1, v5, v3)
    }
    v9 := GenS1{F0: (("a" + (("a" + (("" + "a") + "")) + "\n\t")) + "hello"), F1: v3}
    switch "\n\t" {
    case "":
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 int32 = 2147483647
            var v12 int64 = 501
            var v13 int = fi10
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "hello":
        if false {
            var v14 bool = v3
            v15 := GenS1{F0: "a", F1: ((^(v1)) < (v1 + v1))}
            _ = v14
            _ = v15
        } else {
            v3 = (-2.25 == (3.14159 + 1.5))
            var v16 uint64 = 1
            for fi17 := 0; fi17 < 15; fi17++ {
                if true {
                    switch fi17 {
                    case -515:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 2147483647:
                        for fi18 := 0; fi18 < 9; fi18++ {
                            var v19 uint = 255
                            fmt.Println(v16, fi18, v1)
                            _ = fi18
                            _ = v19
                        }
                    }
                    if ((!(1e10 < -2.25)) || ((65535 != 2) || v3)) {
                        for fi20 := 0; fi20 < 16; fi20++ {
                            switch "" {
                            case "\n\t":
                                var v21 bool = (!(!(("" < "a") && v3)))
                                _ = v21
                            case "a":
                                v1 += v1
                            case "foo bar":
                                switch "\n\t" {
                                case "hello":
                                    v22 := GenS1{F0: ("" + ("" + "")), F1: (((!v3) || (((v1 & v1) > v1) && (((0.0 - (-2.25 * (0.0 - 0.1))) >= -2.25) && ((1e10 + float32(1.5)) != v5)))) && v3)}
                                    _ = v22
                                }
                            }
                            switch fi20 {
                            case -623:
                                if false {
                                    v2[2147483648] = 267
                                    fmt.Println(fi20)
                                    if len(v4) > 0 {
                                        v4[0] = fi20
                                    }
                                    v23 := GenS2{F0: 4294967295}
                                    v24 := GenS2{F0: 65535}
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v25 float32 = v5
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v25
                                }
                            }
                            if v3 {
                                v1 += (v1 & (v1 - v1))
                                v26 := GenS1{F0: "\n\t", F1: false}
                                _ = v26
                            } else {
                                switch "a" {
                                case "hello":
                                    v1 -= v1
                                case "a":
                                    _ = len(v2)
                                case "\n\t":
                                    v4 = append(v4, fi17)
                                }
                                v27 := GenS3{F0: 0, F1: (v1 ^ v1), F2: 9223372036854775807}
                                _ = len(v4)
                                _ = v27
                            }
                            _ = fi20
                        }
                        for fi28 := 0; fi28 < 6; fi28++ {
                            if false {
                                v16 = v16
                            } else {
                                v16 += v16
                                if (-369 < 42) {
                                    if true {
                                        v5 += (v5 * 0.0)
                                        var v29 float32 = (v5 - -2.25)
                                        var v30 float64 = (1.5 * ((0.0 - 1e10) - ((1.5 * 1e10) / 1.0)))
                                        _ = v29
                                        _ = v30
                                    } else {
                                        fmt.Println(v16, v3)
                                        v31 := GenS4{F0: 65535}
                                        v32 := GenS2{F0: 13}
                                        var v33 int = fi28
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    if (-375 > -154) {
                                        var v34 int = fi28
                                        switch (("hello" + ("foo bar" + (("foo bar" + ("foo bar" + "foo bar")) + ""))) + (("a" + ("hello" + "")) + "\n\t")) {
                                        case "a":
                                            fmt.Println(v3)
                                        case "foo bar":
                                            v16 += v16
                                        default:
                                            fmt.Println(v1)
                                        }
                                        _ = v34
                                    }
                                } else {
                                    fmt.Println(v3)
                                }
                            }
                            var v35 uint16 = 255
                            _ = fi28
                            _ = v35
                        }
                    } else {
                        for fi36 := 0; fi36 < 15; fi36++ {
                            v16 = (v16 | 1)
                            v2[770] = 871
                            v5 = (1.5 / v5)
                            _ = fi36
                        }
                        var v37 float32 = ((1.5 * 0.0) / (v5 + (1.5 * 1.5)))
                        _ = v37
                    }
                }
                _ = fi17
            }
            _ = v16
        }
    default:
        v3 = (v3 || (1 > 65535))
    }
    v38 := GenS1{F0: "\n\t", F1: (!v3)}
    _ = len(v2)
    for fi39 := 0; fi39 < 2; fi39++ {
        var v40 uint32 = 906
        _ = fi39
        _ = v40
    }
    v3 = (0 == 18446744073709551615)
    var v41 uint8 = 146
    v1 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v9
}
