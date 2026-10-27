package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 int8; F2 int32 }
func genFn0(p0 bool, p1 int64) uint8 {
    p1 -= p1
    if true {
        if false {
            v7 := GenS1{F0: 42, F1: 244}
            _ = v7
        } else {
            if p0 {
                p1 = p1
            } else {
                var v8 uint8 = 127
                _ = v8
            }
            for fi9 := 0; fi9 < 16; fi9++ {
                switch "\n\t" {
                case "a":
                    var v10 int32 = 0
                    _ = v10
                case "\n\t":
                    fmt.Println(p1)
                case "":
                    fmt.Println(p1, p0, fi9)
                }
                _ = fi9
            }
            fmt.Println(p1, p0)
        }
        if true {
            if (!p0) {
                fmt.Println(p0, p1)
            }
            var v11 float32 = 1.5
            var v12 uint = 1
            _ = v11
            _ = v12
        } else {
            p1 = (p1 % ((p1) | 1))
            fmt.Println(p1, p1)
        }
        p0 = (!false)
    } else {
        fmt.Println(p1, p1)
        p0 = p0
    }
    v13 := []int{2147483647}
    fmt.Println(p1, p0)
    if p0 {
        switch "\n\t" {
        case "foo bar":
            var v14 int32 = -1
            _ = v14
        default:
            var v15 int16 = 1
            _ = v15
        }
        switch ("a" + ("\n\t" + "a")) {
        case "foo bar":
            var v16 int = 4294967295
            _ = v16
        }
        var v17 uint = 18446744073709551615
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v13
    return 49
}

func main() {
    v1 := []uint32{42}
    var v2 int16 = -736
    var v3 uint32 = 114
    v4 := make(map[int]uint32)
    v5 := []uint32{(v3 - 1), (^(v3))}
    var v6 bool = (0 > 255)
    v18 := GenS2{F0: true, F1: 1}
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := GenS2{F0: ((!(true && false)) || v6), F1: 127}
            v22 := GenS3{F0: uint(v3), F1: 354}
            v2 -= (-(v2))
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    var v23 float32 = 1e10
    if (1 >= -1) {
        for fi24 := 0; fi24 < 15; fi24++ {
            switch (("\n\t" + (("" + "") + ("\n\t" + "\n\t"))) + (("hello" + ("" + "hello")) + "a")) {
            case "":
                var v25 uint8 = 127
                _ = v25
            case "a":
                v3 = (v3 % ((0) | 1))
            case "foo bar":
                v26 := GenS2{F0: (-42 >= -95), F1: genFn0(((!v6) || v6), 222)}
                _ = v26
            }
            _ = fi24
        }
    } else {
        for fi27 := 0; fi27 < 4; fi27++ {
            v28 := GenS4{F0: 9223372036854775807}
            if len(v5) > 0 {
                v5[0] = 76
            }
            _ = fi27
            _ = v28
        }
        fmt.Println(v2, v6)
        v2 -= (v2 | v2)
    }
    v4[152] = (v3 ^ v3)
    var v29 int32 = 0
    v5 = append(v5, (v3 % ((42) | 1)))
    for fi30 := 0; fi30 < 13; fi30++ {
        switch (("\n\t" + ("foo bar" + "hello")) + (("foo bar" + "a") + ("foo bar" + "foo bar"))) {
        case "":
            v2 -= v2
        case "hello":
            for fi31 := 0; fi31 < 14; fi31++ {
                var v32 bool = (!v6)
                _ = fi31
                _ = v32
            }
        default:
            v2 -= v2
        }
        switch "\n\t" {
        case "foo bar":
            v33 := GenS1{F0: 2, F1: 42}
            _ = v33
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v3 = (v3 & (v3 * (^(v3))))
        _ = fi30
    }
    var v34 int16 = (v2 + 0)
    v2 += (v2 + v2)
    for fi35 := 0; fi35 < 10; fi35++ {
        var v36 uint16 = 1
        fmt.Println(v3, v36)
        _ = fi35
        _ = v36
    }
    v37 := GenS5{F0: 32767, F1: -30, F2: v29}
    var v38 uint8 = uint8(v3)
    for fi39 := 0; fi39 < 11; fi39++ {
        if true {
            switch "foo bar" {
            case "":
                for fi40 := 0; fi40 < 8; fi40++ {
                    v41 := GenS3{F0: 65535, F1: 4294967295}
                    if true {
                        v42 := GenS1{F0: 42, F1: v3}
                        var v43 uint8 = (v38 / (((v38 << ((18446744073709551615) & 63))) | 1))
                        _ = v42
                        _ = v43
                    }
                    _ = fi40
                    _ = v41
                }
            case "\n\t":
                for fi44 := 0; fi44 < 20; fi44++ {
                    v23 = v23
                    delete(v4, int(v29))
                    _ = fi44
                }
            default:
                var v45 int = (fi39 * (fi39 - (fi39 & (fi39 + fi39))))
                _ = v45
            }
        } else {
            var v46 uint = 0
            _ = v46
        }
        v4[fi39] = v3
        v2 = 0
        _ = fi39
    }
    v29 += (^(v29))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
