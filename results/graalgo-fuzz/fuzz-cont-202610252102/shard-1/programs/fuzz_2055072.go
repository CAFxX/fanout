package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 int8, p1 int64) int16 {
    if true {
        v7 := []int{}
        p0 += -1
        _ = v7
    } else {
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS1{F0: 100, F1: -32768, F2: 1}
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint64 = 2
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1() uint8 {
    var v12 bool = (!(128 != 128))
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := GenS1{F0: -66, F1: 42, F2: 184}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v16 := GenS1{F0: 1, F1: 42, F2: -875}
    switch ("foo bar" + (("a" + "hello") + (("" + "\n\t") + "hello"))) {
    case "\n\t":
        v17 := GenS1{F0: 42, F1: 1000, F2: -1}
        _ = v17
    case "foo bar":
        switch "hello" {
        case "a":
            v18 := []int{4294967295, 1}
            _ = v18
        case "\n\t":
            v19 := GenS2{F0: 0}
            _ = v19
        case "hello":
            fmt.Println(v12)
        default:
            v12 = (65535 > 1)
        }
    case "a":
        v12 = ("" < "a")
    }
    _ = v12
    _ = v16
    return 255
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 int64 = 1
    v4 := []uint32{2147483648, 1, 278, 2}
    v5 := []uint32{0, 65535, 0}
    v6 := []uint32{0}
    var v20 int8 = -1
    if true {
        v1[42] = 1
    } else {
        if true {
            for fi21 := 0; fi21 < 1; fi21++ {
                fmt.Println(v20, v3, v3)
                _ = fi21
            }
        } else {
            if true {
                if false {
                    var v22 uint = 18446744073709551615
                    var v23 int64 = (v3 - (^(v3)))
                    _ = v22
                    _ = v23
                } else {
                    if false {
                        var v24 string = (("" + "\n\t") + "\n\t")
                        _ = v24
                    }
                    v20 -= 109
                }
                v25 := GenS1{F0: v20, F1: -32768, F2: v3}
                _ = v25
            } else {
                if false {
                    _ = len(v2)
                    if false {
                        v20 = v20
                    }
                    for fi26 := 0; fi26 < 5; fi26++ {
                        for fi27 := 0; fi27 < 8; fi27++ {
                            var v28 int = (fi26 >> ((553) & 63))
                            _ = fi27
                            _ = v28
                        }
                        _ = fi26
                    }
                }
                if (!false) {
                    var v29 uint = 18446744073709551615
                    _ = v29
                } else {
                    switch "foo bar" {
                    case "hello":
                        if (!true) {
                            if (false && (-999 > 32768)) {
                                v3 = (-(v3))
                                var v30 int = 42
                                var v31 int8 = (v20 / ((127) | 1))
                                _ = v30
                                _ = v31
                            } else {
                                v20 = (-(v20))
                                var v32 int32 = -929
                                _ = v32
                            }
                            var v33 uint = 255
                            _ = v33
                        }
                    case "a":
                        var v34 int32 = 2147483647
                        _ = v34
                    case "foo bar":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    var v35 float64 = 3.14159
                    _ = v35
                }
            }
            for fi36 := 0; fi36 < 14; fi36++ {
                for fi37 := 0; fi37 < 18; fi37++ {
                    if (1 == 0) {
                        fmt.Println(v3, fi36, v3)
                        switch fi37 {
                        case 1:
                            v3 = (^(v3))
                        case 42:
                            fmt.Println(v20)
                        default:
                            v3 += v3
                        }
                    }
                    v4 = append(v4, 65535)
                    _ = fi37
                }
                v3 -= 1
                _ = fi36
            }
            var v38 int8 = v20
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        switch "" {
        case "\n\t":
            if (true || true) {
                fmt.Println(fi39, v3)
            } else {
                for fi40 := 0; fi40 < 19; fi40++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi40
                }
                v3 = v3
                var v41 float64 = 1.5
                _ = v41
            }
        case "a":
            var v42 int64 = (-(v3))
            _ = v42
        default:
            var v43 int64 = v3
            _ = v43
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi39
    }
    var v44 int32 = 2147483647
    switch "hello" {
    case "foo bar":
        v1[-9223372036854775808] = 421
    }
    for fi45 := 0; fi45 < 1; fi45++ {
        delete(v1, (fi45 | 4294967295))
        if true {
            v3 -= v3
            for fi46 := 0; fi46 < 15; fi46++ {
                fmt.Println(fi46)
                _ = fi46
            }
        }
        _ = fi45
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    v3 -= (v3 | 1)
    var v47 int64 = (v3 / (((v3 / (((v3 % ((v3) | 1))) | 1))) | 1))
    v3 += int64(v20)
    v3 = 42
    v48 := GenS2{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
