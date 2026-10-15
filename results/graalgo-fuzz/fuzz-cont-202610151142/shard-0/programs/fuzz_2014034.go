package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint64; F2 int }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int16; F1 int64; F2 int16 }
func genFn0() uint32 {
    var v4 int32 = -22
    v4 += -1
    _ = v4
    return 4294967295
}
func genFn1(p0 int16, p1 int) int64 {
    var v5 string = ("\n\t" + "")
    switch v5 {
    case "":
        for fi6 := 0; fi6 < 8; fi6++ {
            p0 += (p0 << ((255) & 63))
            switch p1 {
            case -624:
                if true {
                    var v7 bool = (p1 < (fi6 + (p1 | (p1 - fi6))))
                    v7 = (-555 == fi6)
                    var v8 uint32 = genFn0()
                    _ = v7
                    _ = v8
                }
            case 141:
                for fi9 := 0; fi9 < 2; fi9++ {
                    p1 += (^(fi9))
                    v10 := GenS1{F0: (p1 ^ (p1 - fi6)), F1: 65535, F2: 18446744073709551615}
                    var v11 int16 = (p0 & p0)
                    var v12 int64 = -897
                    switch (v5 + (("\n\t" + (v5 + (((v5 + (("hello" + (v5 + (v5 + v5))) + v5)) + v5) + ""))) + (v5 + v5))) {
                    case "\n\t":
                        switch "" {
                        case "foo bar":
                            v13 := GenS1{F0: (fi6 + p1), F1: 122, F2: 42}
                            _ = v13
                        case "\n\t":
                            var v14 int16 = v11
                            _ = v14
                        }
                    default:
                        switch "hello" {
                        case "hello":
                            if (!(!(!false))) {
                                v15 := GenS2{F0: -2147483648, F1: 141, F2: p1}
                                _ = v15
                            }
                        case "":
                            v16 := GenS2{F0: 166, F1: 2, F2: (-(p1))}
                            _ = v16
                        case "a":
                            var v17 int16 = (-(v11))
                            _ = v17
                        }
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            case 42:
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 int16 = -144
                    v20 := []int{p1, (p1 ^ 1), (fi6 >> ((1) & 63)), fi18}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
            _ = fi6
        }
    case "foo bar":
        var v21 float64 = ((1.5 / 1e10) / -2.25)
        _ = v21
    default:
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 int = p1
            _ = fi22
            _ = v23
        }
    }
    v24 := GenS3{F0: (v5 + (((((("foo bar" + "a") + (((v5 + v5) + "") + v5)) + (v5 + v5)) + v5) + v5) + v5)), F1: (p1 << ((233) & 63))}
    _ = p0
    _ = p1
    _ = v24
    _ = v5
    return -1
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -267
    var v3 int = 0
    switch "hello" {
    case "\n\t":
        v25 := GenS2{F0: -2147483648, F1: 12, F2: (v3 / (((v3 - (v3 - v3))) | 1))}
        _ = v25
    case "foo bar":
        var v26 int64 = -1
        _ = v26
    case "":
        if true {
            if (true || false) {
                v1[v3] = (1 < 127)
                v2 -= (v2 ^ v2)
            }
            var v27 uint8 = 107
            _ = v27
        }
    default:
        var v28 float32 = float32((0.1 - -2.25))
        _ = v28
    }
    if true {
        v3 -= v3
    }
    _ = len(v1)
    if true {
        v3 -= 127
    }
    var v29 string = (("foo bar" + "a") + "")
    var v30 int = (v3 >> ((42) & 63))
    v3 = (v30 * v30)
    v2 += v2
    v31 := GenS4{F0: (v2 << ((0) & 63)), F1: 0, F2: (^(v2))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
