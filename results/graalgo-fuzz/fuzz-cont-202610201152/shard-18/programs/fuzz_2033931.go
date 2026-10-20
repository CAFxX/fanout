package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int8; F2 int64 }
type GenS3 struct { F0 bool; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint; F1 int16; F2 bool }
type GenS5 struct { F0 int16 }
func genFn0(p0 int8) string {
    p0 = p0
    p0 -= (p0 - -1)
    var v7 float32 = ((1.5 - 1e10) - -2.25)
    _ = p0
    _ = v7
    return ("foo bar" + ((("foo bar" + "hello") + ("" + "\n\t")) + (("" + "foo bar") + (("foo bar" + "\n\t") + "hello"))))
}
func genFn1(p0 uint32, p1 uint32, p2 int) int16 {
    for fi8 := 0; fi8 < 12; fi8++ {
        if false {
            p0 -= p0
            p1 = 42
            p2 -= p2
        } else {
            p0 -= (p0 * 0)
        }
        _ = fi8
    }
    var v9 uint = 65535
    p0 += 548
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 32767
}

func main() {
    var v1 uint = 1
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    var v5 int16 = 473
    v6 := []int{-2147483648, -9223372036854775808, 305, 42}
    v5 += (-(v5))
    if (123 == -1) {
        fmt.Println(v1, v1)
    }
    v1 += (v1 + v1)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int32 = 442
        _ = fi10
        _ = v11
    }
    var v12 uint = v1
    v3[2147483648] = 2147483647
    for fi13 := 0; fi13 < 18; fi13++ {
        v5 = (-(v5))
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 uint32 = 4294967295
            var v16 int16 = v5
            _ = fi14
            _ = v15
            _ = v16
        }
        v12 = v1
        _ = fi13
    }
    v1 = (^(v1))
    v5 = (v5 & genFn1(1, 65535, 363))
    fmt.Println(v1)
    v17 := GenS1{F0: (v1 >> (((v12 * v12)) & 63))}
    v1 += 862
    v5 -= v5
    switch 0 {
    case 32767:
        switch "hello" {
        case "":
            for fi18 := 0; fi18 < 1; fi18++ {
                if false {
                    var v19 int64 = 9223372036854775807
                    var v20 uint8 = 2
                    switch "hello" {
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "a":
                        v3[0] = -2147483648
                    default:
                        switch ("\n\t" + ("\n\t" + ("" + "hello"))) {
                        case "foo bar":
                            var v21 int8 = 42
                            _ = v21
                        case "":
                            switch ("foo bar" + "a") {
                            case "a":
                                var v22 int64 = -1
                                _ = v22
                            case "foo bar":
                                var v23 float32 = 3.14159
                                _ = v23
                            case "hello":
                                var v24 int16 = v5
                                _ = v24
                            }
                        case "\n\t":
                            if (false && false) {
                                if false {
                                    v12 -= 435
                                    fmt.Println(v20)
                                } else {
                                    fmt.Println(fi18)
                                }
                                v19 -= v19
                            } else {
                                switch ((("foo bar" + "foo bar") + ("a" + "hello")) + ("foo bar" + "foo bar")) {
                                case "\n\t":
                                    switch -572 {
                                    case 1:
                                        if (65535 >= 2) {
                                            v6 = append(v6, 4294967295)
                                        } else {
                                            if false {
                                                var v25 int8 = 100
                                                _ = v25
                                            } else {
                                                fmt.Println(v12, v12, v1)
                                            }
                                            var v26 uint16 = 32767
                                            v20 = (v20 << ((v1) & 63))
                                            _ = v26
                                        }
                                    case -472:
                                        var v27 float32 = 1e10
                                        _ = v27
                                    case 32767:
                                        if ((1e10 == -2.25) && true) {
                                            fmt.Println(v20)
                                        }
                                    default:
                                        v12 += v1
                                    }
                                case "hello":
                                    v28 := GenS2{F0: "\n\t", F1: 42, F2: v19}
                                    _ = v28
                                case "foo bar":
                                    for fi29 := 0; fi29 < 4; fi29++ {
                                        fmt.Println(fi29, v20)
                                        var v30 bool = (!false)
                                        switch fi29 {
                                        case 32767:
                                            switch (fi29 / ((fi18) | 1)) {
                                            case -2147483648:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case -1:
                                                v5 = (v5 % ((42) | 1))
                                            case 1:
                                                if true {
                                                    v31 := v3[65535]
                                                    _ = v31
                                                } else {
                                                    v3[245] = fi29
                                                    var v32 int64 = (v19 % ((v19) | 1))
                                                    v33 := GenS2{F0: "hello", F1: 42, F2: v32}
                                                    _ = v32
                                                    _ = v33
                                                }
                                            }
                                        }
                                        _ = fi29
                                        _ = v30
                                    }
                                default:
                                    for fi34 := 0; fi34 < 10; fi34++ {
                                        var v35 int64 = v19
                                        if false {
                                            if false {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                delete(v3, 2147483647)
                                                v36 := GenS2{F0: ("\n\t" + "a"), F1: 72, F2: (v19 & -1)}
                                                _ = v36
                                            } else {
                                                v20 = (v20 ^ v20)
                                                if true {
                                                    fmt.Println(fi18)
                                                }
                                                if false {
                                                    v5 = (-(v5))
                                                }
                                                if (v12 < (^(v1))) {
                                                    switch "\n\t" {
                                                    case "a":
                                                        v37 := GenS3{F0: ((((((((("foo bar" + "hello") + "foo bar") + "") + ("\n\t" + "foo bar")) + "foo bar") + "hello") + ("" + ("a" + "foo bar"))) <= "a") && false), F1: 18446744073709551615, F2: -1}
                                                        _ = v37
                                                    default:
                                                        var v38 uint = v12
                                                        _ = v38
                                                    }
                                                    v39 := GenS4{F0: v12, F1: (v5 >> ((65535) & 63)), F2: (!true)}
                                                    _ = v39
                                                }
                                                var v40 bool = ((v12 >= 2) && true)
                                                _ = v40
                                            }
                                            var v41 uint64 = 0
                                            _ = v41
                                        }
                                        _ = fi34
                                        _ = v35
                                    }
                                }
                                v42 := GenS5{F0: v5}
                                _ = v42
                            }
                        }
                    }
                    _ = v19
                    _ = v20
                }
                var v43 uint = 1
                v44 := GenS1{F0: (v43 + v43)}
                _ = fi18
                _ = v43
                _ = v44
            }
        case "hello":
            var v45 uint16 = 272
            _ = v45
        default:
            fmt.Println(v5)
        }
    }
    v5 -= (-(v5))
    var v46 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
