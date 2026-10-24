package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 uint; F2 uint16 }
func genFn0(p0 int32, p1 uint64, p2 int8) uint {
    switch 0 {
    case 127:
        var v6 float32 = 0.1
        _ = v6
    case -2147483648:
        v7 := GenS1{F0: 2, F1: (p2 ^ 0)}
        _ = v7
    default:
        if false {
            p0 -= (p0 | p0)
        }
    }
    v8 := GenS1{F0: 128, F1: (p2 % (((^(p2))) | 1))}
    p2 = -122
    p1 = p1
    v9 := GenS1{F0: 128, F1: -22}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1() uint32 {
    var v10 int8 = -46
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int8 = v10
        _ = fi11
        _ = v12
    }
    v13 := []int{128}
    if (true || (1 > 238)) {
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 float32 = 1e10
            v16 := GenS2{F0: 65535, F1: ("" + ("hello" + ("\n\t" + ("foo bar" + "foo bar"))))}
            var v17 float32 = 0.1
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    }
    if ((false && (0.1 > (3.14159 - 0.1))) && (1000 > 32767)) {
        v10 -= (-(v10))
        var v18 bool = false
        if (!v18) {
            v10 -= -43
        } else {
            if v18 {
                if ((255 > 106) || (!(18446744073709551615 != 0))) {
                    var v19 int8 = (v10 & v10)
                    for fi20 := 0; fi20 < 4; fi20++ {
                        switch (fi20 ^ (^(fi20))) {
                        case -9223372036854775808:
                            v18 = (!(4294967295 < 18446744073709551615))
                        default:
                            var v21 bool = (0 < 1)
                            _ = v21
                        }
                        _ = fi20
                    }
                    _ = v19
                }
                v10 -= v10
                v13 = append(v13, 32768)
            } else {
                if (42 != 2147483647) {
                    for fi22 := 0; fi22 < 8; fi22++ {
                        switch (-(fi22)) {
                        case 42:
                            var v23 bool = ((656 != 431) || (50 < 255))
                            _ = v23
                        case -480:
                            if false {
                                if v18 {
                                    fmt.Println(v10, fi22, fi22)
                                    v13 = append(v13, fi22)
                                }
                            } else {
                                v24 := GenS2{F0: 0, F1: "\n\t"}
                                _ = v24
                            }
                        case -1:
                            v25 := GenS3{F0: -2147483648, F1: (("\n\t" + (("foo bar" + ((("" + ("a" + "hello")) + "a") + ((("a" + "") + "foo bar") + (("" + "foo bar") + "\n\t")))) + "hello")) == ((("foo bar" + "") + (("\n\t" + "\n\t") + ("hello" + ""))) + (("a" + "a") + "")))}
                            _ = v25
                        }
                        _ = fi22
                    }
                    switch ((("\n\t" + (((("foo bar" + "") + "foo bar") + ("" + "hello")) + "foo bar")) + "hello") + "\n\t") {
                    case "foo bar":
                        if (1 >= 1) {
                            var v26 uint16 = 2
                            _ = v26
                        }
                    case "hello":
                        var v27 string = "\n\t"
                        _ = v27
                    case "\n\t":
                        if v18 {
                            for fi28 := 0; fi28 < 8; fi28++ {
                                v10 -= -1
                                _ = fi28
                            }
                            if true {
                                for fi29 := 0; fi29 < 19; fi29++ {
                                    v18 = ("" == "")
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    fmt.Println(v18, v10, v10)
                                    _ = fi29
                                }
                                v30 := GenS2{F0: 414, F1: ("" + "\n\t")}
                                v31 := GenS1{F0: 2, F1: (v10 | v10)}
                                _ = v30
                                _ = v31
                            }
                            v32 := GenS4{F0: 42}
                            _ = v32
                        }
                    default:
                        var v33 string = "\n\t"
                        _ = v33
                    }
                    v18 = v18
                } else {
                    for fi34 := 0; fi34 < 20; fi34++ {
                        v10 += v10
                        _ = fi34
                    }
                    fmt.Println(v18, v18)
                }
            }
            var v35 float32 = (1e10 - -2.25)
            _ = v35
        }
        _ = v18
    } else {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        v10 = 1
        var v36 string = "\n\t"
        _ = v36
    }
    _ = v10
    _ = v13
    return 2147483648
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    v3 := []string{("\n\t" + ""), "foo bar", "hello"}
    var v4 uint32 = 0
    v5 := make(map[int]uint32)
    v37 := GenS5{F0: -492, F1: 0, F2: 32767}
    var v38 uint32 = 2
    for fi39 := 0; fi39 < 5; fi39++ {
        for fi40 := 0; fi40 < 6; fi40++ {
            var v41 int16 = -1
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    v4 = (v38 ^ v4)
    if (42 <= 42) {
        fmt.Println(v38, v4)
    }
    v4 = genFn1()
    v4 = (v38 / ((v4) | 1))
    for fi42 := 0; fi42 < 8; fi42++ {
        v1[fi42] = (^(fi42))
        if (!(!true)) {
            var v43 uint8 = uint8(v4)
            _ = v43
        }
        switch fi42 {
        case 4294967295:
            var v44 bool = false
            _ = v44
        }
        _ = fi42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
