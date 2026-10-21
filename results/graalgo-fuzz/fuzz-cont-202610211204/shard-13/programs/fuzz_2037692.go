package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint8; F1 string }
func genFn0(p0 bool, p1 int, p2 int32) int32 {
    v5 := GenS1{F0: -1}
    switch "\n\t" {
    case "hello":
        switch "" {
        case "hello":
            v6 := []int{1, p1, p1}
            _ = v6
        case "":
            for fi7 := 0; fi7 < 12; fi7++ {
                p0 = true
                if ((0.0 + 0.0) < (1.5 + 3.14159)) {
                    var v8 uint32 = 2
                    v9 := []int{(fi7 + 32767), p1, p1}
                    if p0 {
                        _ = len(v9)
                        var v10 uint32 = v8
                        _ = v10
                    }
                    _ = v8
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v12 := GenS2{F0: 42, F1: 2, F2: 202}
                        v13 := []int{(fi11 >> ((220) & 63))}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v14 := []int{(fi7 << ((65535) & 63)), 0, 127}
                    _ = v14
                }
                _ = fi7
            }
        default:
            v15 := GenS1{F0: 127}
            _ = v15
        }
    case "":
        switch (p1 << ((18446744073709551615) & 63)) {
        case 1:
            var v16 int64 = -9223372036854775808
            _ = v16
        case 42:
            v17 := []int{(p1 >> ((18446744073709551615) & 63)), (p1 << ((784) & 63)), (p1 * (p1 + 128)), (p1 & p1)}
            _ = v17
        case 884:
            p2 = p2
        default:
            var v18 uint16 = 42
            _ = v18
        }
    default:
        fmt.Println(p0, p1)
    }
    v19 := GenS2{F0: 1000, F1: 4294967295, F2: 4294967295}
    fmt.Println(p1, p0)
    p2 += (p2 + (-(p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v5
    return (-(p2))
}
func genFn1(p0 int32, p1 bool, p2 int64) uint32 {
    for fi20 := 0; fi20 < 4; fi20++ {
        switch "\n\t" {
        case "foo bar":
            p2 += (p2 * (p2 - 1))
        default:
            var v21 int16 = 604
            _ = v21
        }
        v22 := []int{fi20, (fi20 << ((4294967295) & 63)), (fi20 >> ((65535) & 63))}
        v23 := GenS3{F0: -2147483648, F1: 42}
        _ = fi20
        _ = v22
        _ = v23
    }
    p2 = (-(p2))
    var v24 uint64 = 42
    p2 = p2
    var v25 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    return 2147483648
}

func main() {
    v1 := []string{"hello", "a", "\n\t", "foo bar"}
    var v2 int = 128
    var v3 float64 = float64(-2.25)
    var v4 uint = 0
    var v26 uint16 = 255
    var v27 bool = ((v2 * v2) > v2)
    v26 += v26
    var v28 int8 = 0
    v4 = (v4 >> (((v4 | v4)) & 63))
    var v29 bool = (v4 >= (v4 + (v4 << (((v4 + v4)) & 63))))
    v30 := GenS3{F0: 722, F1: 65535}
    var v31 float64 = v3
    fmt.Println(v26, v3, v2)
    switch 876 {
    case 982:
        v26 += 778
    case 42:
        var v32 int8 = v28
        _ = v32
    }
    var v33 uint = (v4 >> ((2) & 63))
    for fi34 := 0; fi34 < 6; fi34++ {
        fmt.Println(v31, v28, v29)
        for fi35 := 0; fi35 < 20; fi35++ {
            v36 := GenS4{F0: 159, F1: ""}
            _ = fi35
            _ = v36
        }
        var v37 uint8 = 174
        _ = fi34
        _ = v37
    }
    if (0.0 == 1e10) {
        fmt.Println(v29)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi38 := 0; fi38 < 15; fi38++ {
            for fi39 := 0; fi39 < 8; fi39++ {
                fmt.Println(v3)
                var v40 uint16 = v26
                _ = len(v1)
                _ = fi39
                _ = v40
            }
            v26 = (^(v26))
            switch v2 {
            case 4294967295:
                v33 = (v33 % ((18446744073709551615) | 1))
            case 100:
                switch "\n\t" {
                case "foo bar":
                    v4 = v33
                default:
                    v33 += 65535
                }
            case 2147483647:
                for fi41 := 0; fi41 < 1; fi41++ {
                    if true {
                        var v42 uint64 = 312
                        var v43 uint32 = genFn1(0, ((false || v27) || true), -244)
                        _ = v42
                        _ = v43
                    } else {
                        fmt.Println(fi41, v33, v3)
                        fmt.Println(v31)
                    }
                    fmt.Println(v27)
                    _ = fi41
                }
            default:
                v3 = float64(1e10)
            }
            _ = fi38
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
}
