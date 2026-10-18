package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int8 }
type GenS2 struct { F0 int8; F1 int; F2 bool }
type GenS3 struct { F0 string }
func genFn0(p0 int, p1 uint64, p2 uint32) int16 {
    for fi4 := 0; fi4 < 14; fi4++ {
        p0 -= (p0 << ((uint(p2)) & 63))
        _ = fi4
    }
    if (p2 > 907) {
        switch ("hello" + ("foo bar" + "a")) {
        case "foo bar":
            v5 := []int{(p0 + p0), (p0 - (p0 & (p0 % ((127) | 1))))}
            _ = v5
        default:
            v6 := GenS1{F0: 2, F1: false, F2: -1}
            _ = v6
        }
        v7 := GenS2{F0: 127, F1: (p0 & -1), F2: true}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 string = "foo bar"
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := GenS1{F0: (p1 & p1), F1: ((42 != 1) || false), F2: -128}
                _ = fi10
                _ = v11
            }
            var v12 float32 = (-2.25 + -2.25)
            _ = fi8
            _ = v12
            _ = v9
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            p1 = (p1 << ((65535) & 63))
            var v14 int = (fi13 << ((648) & 63))
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        if (!(0 <= 0)) {
            p2 -= 2147483647
        } else {
            var v16 int = 42
            _ = v16
        }
        _ = fi15
    }
    switch p0 {
    case 9223372036854775807:
        var v17 uint8 = 1
        _ = v17
    case -9223372036854775808:
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 uint64 = p1
            p2 = p2
            _ = fi18
            _ = v19
        }
    default:
        var v20 bool = false
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}
func genFn1(p0 uint8) uint32 {
    p0 = p0
    var v21 float64 = (3.14159 + -2.25)
    v22 := GenS2{F0: 127, F1: -9223372036854775808, F2: true}
    v23 := GenS2{F0: 0, F1: 1, F2: true}
    _ = p0
    _ = v21
    _ = v22
    _ = v23
    return 2147483647
}

func main() {
    var v1 uint64 = 114
    var v2 int64 = 1
    var v3 int8 = 127
    v1 += (v1 % ((v1) | 1))
    v24 := GenS2{F0: v3, F1: 921, F2: true}
    var v25 int32 = -877
    v2 = 9223372036854775807
    switch "" {
    case "foo bar":
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 bool = false
            if v27 {
                fmt.Println(v1, fi26)
                var v28 uint16 = 2
                for fi29 := 0; fi29 < 13; fi29++ {
                    if v27 {
                        if v27 {
                            v1 = 1
                            fmt.Println(fi26)
                        }
                        if (!(!v27)) {
                            v1 = (^(v1))
                        }
                        if (4294967295 == 698) {
                            var v30 int16 = int16(v25)
                            if v27 {
                                v31 := []int{(fi29 ^ (fi29 + (fi26 * (^(fi26))))), (fi29 << ((1) & 63)), fi26, int(v28)}
                                var v32 int64 = v2
                                _ = v31
                                _ = v32
                            }
                            v33 := GenS1{F0: 411, F1: true, F2: v3}
                            _ = v30
                            _ = v33
                        }
                    } else {
                        var v34 int = fi29
                        var v35 float32 = (-2.25 * -2.25)
                        _ = v34
                        _ = v35
                    }
                    _ = fi29
                }
                _ = v28
            } else {
                v36 := []int{(fi26 * fi26), 175, fi26}
                _ = v36
            }
            v1 = (v1 >> ((288) & 63))
            _ = fi26
            _ = v27
        }
    case "\n\t":
        for fi37 := 0; fi37 < 9; fi37++ {
            v1 -= v1
            _ = fi37
        }
    }
    switch 2147483647 {
    case 954:
        if false {
            if false {
                fmt.Println(v1, v2)
            }
            v38 := []int{4294967295, 32768, 9223372036854775807}
            _ = v38
        }
    case 2147483647:
        switch 127 {
        case 644:
            for fi39 := 0; fi39 < 16; fi39++ {
                for fi40 := 0; fi40 < 20; fi40++ {
                    v3 = int8(v25)
                    _ = fi40
                }
                _ = fi39
            }
        case -830:
            v1 = 4294967295
        case 0:
            v41 := GenS3{F0: "a"}
            _ = v41
        }
    }
    var v42 string = ("hello" + (("hello" + "hello") + ("a" + "\n\t")))
    if (((v25 | (v25 ^ (v25 << ((2) & 63)))) > v25) && false) {
        v43 := []int{996, -87, 0}
        v2 = v2
        _ = v43
    }
    var v44 int = 127
    var v45 int64 = v2
    for fi46 := 0; fi46 < 20; fi46++ {
        v45 += -1
        _ = fi46
    }
    var v47 uint16 = 2
    v48 := GenS1{F0: v1, F1: true, F2: v3}
    var v49 uint8 = 210
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v42
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v49
}
