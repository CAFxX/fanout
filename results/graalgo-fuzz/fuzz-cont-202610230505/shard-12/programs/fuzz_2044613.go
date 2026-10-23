package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 int32; F1 string }
func genFn0(p0 string) uint8 {
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return 1
}

func main() {
    var v1 uint32 = 416
    var v2 int64 = -1
    var v3 uint16 = 32767
    var v4 uint8 = 0
    var v5 uint8 = (v4 / ((v4) | 1))
    var v6 uint32 = (v1 - (v1 % (((v1 + v1)) | 1)))
    if (!true) {
        v1 -= (v1 / ((941) | 1))
        v2 -= (v2 << ((377) & 63))
        switch 9223372036854775807 {
        case -9223372036854775808:
            v7 := []int{}
            _ = v7
        default:
            v2 += -749
        }
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int32 = 887
        var v10 uint8 = (v4 | v5)
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        for fi11 := 0; fi11 < 7; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                switch (fi12 % ((fi11) | 1)) {
                case 2147483647:
                    if false {
                        v5 = 128
                        v13 := GenS1{F0: 0, F1: ("" + "a"), F2: 806}
                        _ = v13
                    } else {
                        var v14 int64 = (v2 & -421)
                        if (true || false) {
                            var v15 int = (fi11 ^ (fi11 * (fi11 / ((fi11) | 1))))
                            var v16 int8 = 100
                            _ = v15
                            _ = v16
                        }
                        _ = v14
                    }
                }
                fmt.Println(v4, v2)
                v17 := GenS1{F0: 0, F1: "", F2: (^(v3))}
                _ = fi12
                _ = v17
            }
            var v18 int = 1
            v19 := GenS1{F0: -1, F1: (("foo bar" + "foo bar") + "foo bar"), F2: (v3 << ((2) & 63))}
            _ = fi11
            _ = v18
            _ = v19
        }
    }
    switch 32768 {
    case -9223372036854775808:
        var v20 int32 = 42
        _ = v20
    case 128:
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := GenS2{F0: v5, F1: 9}
            fmt.Println(v1, fi21)
            for fi23 := 0; fi23 < 2; fi23++ {
                fmt.Println(v4, fi23, v3)
                if (!true) {
                    v24 := GenS3{F0: false}
                    v3 += uint16(v6)
                    _ = v24
                } else {
                    v25 := GenS4{F0: 740, F1: -32768}
                    if false {
                        var v26 int32 = 420
                        v2 = v2
                        v2 = v2
                        _ = v26
                    }
                    v2 -= (v2 ^ -9223372036854775808)
                    var v27 int = 676
                    _ = v25
                    _ = v27
                }
                v28 := []int{}
                _ = fi23
                _ = v28
            }
            _ = fi21
            _ = v22
        }
    default:
        if ((v6 - v1) != v6) {
            v6 = v1
        }
    }
    var v29 int8 = 42
    v2 -= -1
    v3 -= (^(v3))
    fmt.Println(v1, v6)
    switch ("a" + ((("" + ("\n\t" + "")) + "hello") + "foo bar")) {
    case "\n\t":
        var v30 uint64 = 330
        _ = v30
    case "hello":
        v31 := []int{-7, -2147483648}
        _ = v31
    case "foo bar":
        v32 := GenS5{F0: 42, F1: "foo bar"}
        _ = v32
    default:
        v4 -= v5
    }
    var v33 bool = true
    v29 -= int8(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
