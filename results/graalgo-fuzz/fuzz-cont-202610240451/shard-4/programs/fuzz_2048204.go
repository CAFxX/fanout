package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS2 struct { F0 bool; F1 int8; F2 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0(p0 int8, p1 int32) int32 {
    if false {
        p0 -= 42
        if false {
            p1 = p1
        }
    }
    p1 -= p1
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return p1
}
func genFn1() int {
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 uint16 = 90
        v6 := []int{(-(fi4)), 0, fi4, (fi4 >> ((18446744073709551615) & 63))}
        v5 -= (v5 / ((32767) | 1))
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint16 = 119
    v8 := []int{-6}
    switch 2147483647 {
    case -755:
        v9 := GenS1{F0: 805, F1: 18446744073709551615, F2: 0}
        _ = v9
    case 2147483647:
        var v10 uint32 = 65535
        _ = v10
    case 127:
        var v11 uint32 = 65535
        _ = v11
    default:
        v12 := GenS1{F0: 1, F1: 393, F2: 42}
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 uint64 = 298
        fmt.Println(v7, fi13, fi13)
        _ = fi13
        _ = v14
    }
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 32767
    v2 := []int{4294967295, 0}
    var v3 int = 32767
    var v15 int16 = 1000
    v16 := GenS2{F0: false, F1: 127, F2: (v1 - v1)}
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            v15 += (v15 & -1)
            var v19 uint = 18446744073709551615
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    switch "" {
    case "":
        if false {
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS2{F0: false, F1: -128, F2: 255}
                v22 := GenS3{F0: fi20, F1: 1}
                _ = fi20
                _ = v21
                _ = v22
            }
            v23 := GenS4{F0: (^(v15)), F1: 128}
            _ = v23
        } else {
            v24 := GenS5{F0: 1, F1: (1e10 != 0.0)}
            switch "" {
            case "\n\t":
                var v25 uint8 = 2
                _ = v25
            case "":
                fmt.Println(v1)
            default:
                v3 -= v3
            }
            _ = v24
        }
    case "hello":
        switch "hello" {
        case "\n\t":
            for fi26 := 0; fi26 < 11; fi26++ {
                v3 += 58
                v27 := GenS1{F0: 371, F1: 807, F2: 42}
                _ = fi26
                _ = v27
            }
        }
    }
    v15 -= (^(v15))
    if false {
        v28 := GenS6{F0: (v15 % (((v15 % ((1) | 1))) | 1)), F1: (float64(3.14159) >= (0.0 + 1.5))}
        v15 = (-(v15))
        _ = v28
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v29 := GenS5{F0: 396, F1: (0.1 <= 1e10)}
    fmt.Println(v15)
    v1 -= (v1 ^ v1)
    v30 := GenS7{F0: 2, F1: 20, F2: 255}
    for fi31 := 0; fi31 < 11; fi31++ {
        v1 = (v1 >> ((4294967295) & 63))
        v32 := GenS1{F0: 2147483648, F1: 0, F2: 0}
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        switch (((("a" + "") + "hello") + "foo bar") + "hello") {
        case "\n\t":
            var v34 int = (fi33 | (v3 ^ 1))
            _ = v34
        case "foo bar":
            v1 = (^(v1))
        }
        v35 := GenS4{F0: v15, F1: 128}
        var v36 bool = (184 != 65535)
        _ = fi33
        _ = v35
        _ = v36
    }
    if len(v2) > 0 {
        v2[0] = genFn1()
    }
    var v37 int8 = 127
    fmt.Println(v3, v15)
    if (!(0.1 >= (1.5 * (1.5 / 1.0)))) {
        v15 = v15
    } else {
        if (false || false) {
            fmt.Println(v1, v1)
            fmt.Println(v15, v37)
        }
    }
    var v38 int8 = (v37 << ((846) & 63))
    switch -774 {
    case 1:
        switch v3 {
        case 4294967295:
            if false {
                v15 += int16(v3)
                fmt.Println(v1)
            } else {
                _ = len(v2)
                fmt.Println(v1)
                var v39 uint32 = 0
                _ = v39
            }
        case 128:
            v40 := GenS7{F0: 2, F1: -863, F2: 128}
            _ = v40
        case 2147483647:
            for fi41 := 0; fi41 < 2; fi41++ {
                if (-99 >= -666) {
                    switch genFn1() {
                    case 9223372036854775807:
                        v42 := GenS4{F0: 1000, F1: 7}
                        _ = v42
                    default:
                        var v43 int = 1
                        _ = v43
                    }
                }
                if len(v2) > 0 {
                    v2[0] = (^(v3))
                }
                var v44 int = (fi41 >> ((4294967295) & 63))
                _ = fi41
                _ = v44
            }
        }
    case 42:
        for fi45 := 0; fi45 < 17; fi45++ {
            fmt.Println(v38, v3)
            if (0 != 2) {
                v46 := GenS7{F0: 18446744073709551615, F1: v15, F2: 42}
                _ = v46
            }
            _ = fi45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
}
