package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 string; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 bool; F1 int8; F2 uint64 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int8; F1 uint8; F2 string }
func genFn0(p0 int16) uint16 {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int = 1
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int8 = 116
            fmt.Println(v10, p0, v10)
            p0 = (^(p0))
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := []int{(fi11 << ((2) & 63)), fi11}
            _ = fi11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    v13 := GenS1{F0: 1}
    _ = p0
    _ = v13
    return 97
}
func genFn1(p0 uint64, p1 uint8) int {
    p0 -= p0
    if (p0 == 611) {
        switch "hello" {
        case "hello":
            p1 = (p1 << ((255) & 63))
        case "\n\t":
            v14 := GenS1{F0: 127}
            _ = v14
        default:
            fmt.Println(p1, p1)
        }
    }
    var v15 bool = ((p1 % ((0) | 1)) != 154)
    _ = p0
    _ = p1
    _ = v15
    return 127
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = -210
    v3 := make(map[uint32]int)
    v4 := []string{"foo bar", ""}
    var v5 string = "hello"
    var v6 uint64 = (v1 + v1)
    v1 = v1
    for fi16 := 0; fi16 < 4; fi16++ {
        v1 += (v1 & (v6 - (v6 >> ((255) & 63))))
        _ = len(v3)
        if false {
            switch fi16 {
            case 772:
                var v17 int32 = 2147483647
                _ = v17
            case -2147483648:
                if true {
                    fmt.Println(v6, v5)
                }
            }
        }
        _ = fi16
    }
    if (127 == 1) {
        v18 := GenS1{F0: 124}
        v19 := GenS1{F0: uint8(v1)}
        _ = v18
        _ = v19
    } else {
        if (false && false) {
            v6 -= 4294967295
            var v20 int16 = 1000
            fmt.Println(v2, v6, v6)
            _ = v20
        } else {
            fmt.Println(v2, v1)
        }
        v21 := GenS2{F0: -1, F1: (v5 + (((v5 + v5) + ((v5 + (((v5 + v5) + v5) + (v5 + ("hello" + "hello")))) + "a")) + (((v5 + v5) + "") + v5))), F2: 183}
        v22 := GenS3{F0: -1}
        _ = v21
        _ = v22
    }
    switch genFn1(v6, 0) {
    case 42:
        delete(v3, 1)
    case 128:
        v5 = v5
    case -110:
        switch -9223372036854775808 {
        case 1:
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 float32 = (-2.25 * 1.5)
                _ = fi23
                _ = v24
            }
        case -1:
            if (!true) {
                switch "" {
                case "":
                    v5 = v5
                case "\n\t":
                    v1 += (v6 * v6)
                default:
                    var v25 float32 = ((3.14159 / 1.0) * -2.25)
                    _ = v25
                }
                for fi26 := 0; fi26 < 9; fi26++ {
                    for fi27 := 0; fi27 < 7; fi27++ {
                        v4 = append(v4, "hello")
                        _ = fi27
                    }
                    _ = fi26
                }
            } else {
                v1 -= (v1 * v1)
            }
        default:
            v5 = ("foo bar" + ("a" + ("hello" + "a")))
        }
    default:
        if (219 < 251) {
            v28 := GenS3{F0: 42}
            v1 = (v6 % ((v1) | 1))
            _ = v28
        } else {
            v29 := GenS1{F0: 2}
            _ = v29
        }
    }
    fmt.Println(v6)
    if len(v4) > 0 {
        v4[0] = ("\n\t" + v5)
    }
    if true {
        var v30 int16 = 1
        v1 += 676
        _ = v30
    } else {
        delete(v3, 793)
    }
    if (false && true) {
        switch v5 {
        case "a":
            v31 := GenS4{F0: 1}
            _ = v31
        case "hello":
            if false {
                v32 := GenS4{F0: 65535}
                _ = v32
            }
        default:
            v5 = (v5 + v5)
        }
        v5 = v5
    } else {
        v33 := GenS4{F0: 65535}
        v1 = v6
        var v34 int8 = 1
        _ = v33
        _ = v34
    }
    v35 := GenS5{F0: v2, F1: (^(v6))}
    switch v5 {
    case "hello":
        fmt.Println(v1, v2, v1)
    case "foo bar":
        var v36 int = 9223372036854775807
        _ = v36
    case "a":
        var v37 int = genFn1(v1, 9)
        _ = v37
    }
    v38 := GenS6{F0: true, F1: 0, F2: 0}
    v6 = v6
    fmt.Println(v5, v2, v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        v40 := GenS7{F0: v5, F1: -604}
        v41 := GenS2{F0: -642, F1: ("\n\t" + v5), F2: 4294967295}
        for fi42 := 0; fi42 < 17; fi42++ {
            v6 -= v1
            _ = fi42
        }
        _ = fi39
        _ = v40
        _ = v41
    }
    _ = len(v4)
    v43 := GenS8{F0: -128, F1: 122, F2: v5}
    v1 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
