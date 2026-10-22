package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 uint64; F1 int; F2 int16 }
func genFn0(p0 uint64, p1 bool, p2 uint64) uint64 {
    p1 = p1
    for fi6 := 0; fi6 < 8; fi6++ {
        if false {
            v7 := GenS1{F0: 1}
            _ = v7
        } else {
            p0 = p0
            for fi8 := 0; fi8 < 1; fi8++ {
                if (((!p1) && (!(32767 <= 32767))) || p1) {
                    switch 2147483647 {
                    case 128:
                        p2 = (p2 >> ((0) & 63))
                    case 959:
                        p2 = (p0 - (p2 ^ p2))
                        p0 -= 19
                    default:
                        switch fi6 {
                        case -587:
                            var v9 int = fi8
                            _ = v9
                        case 128:
                            switch "" {
                            case "hello":
                                v10 := GenS1{F0: 437}
                                _ = v10
                            default:
                                p1 = ((("hello" + "a") < ("hello" + "foo bar")) || (428 <= 635))
                            }
                        case -765:
                            var v11 uint32 = 4294967295
                            _ = v11
                        default:
                            fmt.Println(fi8)
                        }
                        v12 := GenS1{F0: 0}
                        v13 := []int{}
                        _ = v12
                        _ = v13
                    }
                    v14 := []int{(fi8 + -1), 128}
                    _ = v14
                } else {
                    var v15 uint32 = 42
                    p0 = 71
                    _ = v15
                }
                var v16 uint64 = p2
                _ = fi8
                _ = v16
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return 384
}
func genFn1(p0 string) int32 {
    fmt.Println(p0)
    v17 := GenS2{F0: 360, F1: 1, F2: 255}
    var v18 float64 = (3.14159 / (-2.25 / 1e10))
    var v19 uint64 = 18446744073709551615
    var v20 uint8 = 1
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    _ = v20
    return 2147483647
}

func main() {
    var v1 uint32 = 2
    var v2 int16 = 1
    v3 := []uint32{v1, (v1 + v1), (v1 ^ v1), (v1 % (((v1 & (v1 ^ v1))) | 1))}
    v4 := []int{4294967295, 1, 0, 539}
    var v5 uint = 1
    var v21 int64 = 24
    v22 := GenS2{F0: (v21 * v21), F1: 0, F2: uint(v1)}
    fmt.Println(v5, v21)
    v5 += (^(v5))
    switch "\n\t" {
    case "a":
        for fi23 := 0; fi23 < 2; fi23++ {
            fmt.Println(v5, fi23, v5)
            _ = fi23
        }
    }
    var v24 int64 = (v21 - (v21 | int64(v5)))
    for fi25 := 0; fi25 < 1; fi25++ {
        fmt.Println(v21, v5)
        var v26 uint16 = 65535
        fmt.Println(v1, v5)
        _ = fi25
        _ = v26
    }
    v2 -= v2
    v5 += (v5 / (((v5 >> ((v5) & 63))) | 1))
    var v27 int32 = -1
    var v28 uint = v5
    _ = len(v4)
    switch "" {
    case "a":
        v5 = v5
    case "foo bar":
        for fi29 := 0; fi29 < 1; fi29++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (!(2 <= 2)) {
                var v30 uint8 = 2
                _ = v30
            }
            _ = fi29
        }
    default:
        for fi31 := 0; fi31 < 13; fi31++ {
            v27 += (v27 + v27)
            var v32 uint8 = 0
            var v33 int8 = 1
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    if true {
        var v34 int16 = (v2 - (v2 / ((v2) | 1)))
        v24 += (v21 >> ((65535) & 63))
        if (((("a" + "") + ("" + "")) < ("\n\t" + "")) || (v21 == v21)) {
            v4 = append(v4, 42)
            v24 += v24
            if false {
                v35 := GenS3{F0: 1, F1: 32767, F2: (v34 & -1)}
                v21 += (v24 >> (((v5 / ((1) | 1))) & 63))
                v5 = v28
                _ = v35
            }
        }
        _ = v34
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        for fi37 := 0; fi37 < 4; fi37++ {
            v4 = append(v4, fi37)
            _ = fi37
        }
        var v38 int16 = v2
        for fi39 := 0; fi39 < 18; fi39++ {
            v21 = (v24 * v24)
            _ = fi39
        }
        _ = fi36
        _ = v38
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        var v41 int32 = genFn1("foo bar")
        _ = fi40
        _ = v41
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
