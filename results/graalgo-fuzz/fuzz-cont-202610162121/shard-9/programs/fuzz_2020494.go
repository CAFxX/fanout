package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint16; F1 uint16; F2 int }
type GenS8 struct { F0 int }
func genFn0(p0 uint16, p1 uint64) int {
    v7 := GenS1{F0: 127, F1: 42}
    if false {
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS2{F0: fi9}
                v11 := GenS3{F0: 65535, F1: -407, F2: 2}
                switch 32768 {
                case 127:
                    p0 -= 255
                case 32767:
                    v12 := []int{-169, 128}
                    _ = v12
                case 2147483647:
                    var v13 int32 = 1
                    _ = v13
                default:
                    if ((p0 % ((255) | 1)) >= (p0 & 65535)) {
                        v14 := []int{(fi8 - (fi8 | fi8)), 650, fi9, fi9}
                        _ = v14
                    }
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            p1 = p1
            if true {
                var v16 int = (fi15 % ((fi15) | 1))
                for fi17 := 0; fi17 < 8; fi17++ {
                    if false {
                        fmt.Println(fi15)
                        var v18 uint64 = (p1 % (((p1 >> ((80) & 63))) | 1))
                        v19 := GenS4{F0: p0, F1: int32(fi17)}
                        _ = v18
                        _ = v19
                    } else {
                        switch "a" {
                        case "foo bar":
                            p0 -= (^(p0))
                        case "":
                            var v20 uint = 2
                            _ = v20
                        case "\n\t":
                            var v21 float64 = 1e10
                            _ = v21
                        }
                        v22 := GenS5{F0: 0}
                        fmt.Println(p0, p1)
                        _ = v22
                    }
                    switch "\n\t" {
                    case "\n\t":
                        var v23 int16 = -259
                        if (1e10 != float32(1.5)) {
                            v24 := GenS4{F0: (p0 - (p0 | p0)), F1: 1}
                            if true {
                                v16 = v16
                                if true {
                                    v16 -= (fi15 ^ -931)
                                    var v25 int8 = 1
                                    _ = v25
                                }
                                if true {
                                    switch (fi17 << ((18446744073709551615) & 63)) {
                                    case -509:
                                        var v26 int64 = -1
                                        _ = v26
                                    case 32768:
                                        v27 := []int{}
                                        _ = v27
                                    }
                                }
                            }
                            _ = v24
                        }
                        _ = v23
                    }
                    _ = fi17
                }
                _ = v16
            } else {
                switch "hello" {
                case "hello":
                    v28 := []int{(fi15 * fi15), fi15}
                    _ = v28
                case "":
                    var v29 uint = 18446744073709551615
                    _ = v29
                default:
                    for fi30 := 0; fi30 < 17; fi30++ {
                        var v31 float64 = (0.1 - 0.1)
                        v31 = -2.25
                        v32 := GenS5{F0: 127}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
                p0 = 32767
            }
            v33 := GenS6{F0: 0, F1: p1, F2: (p1 ^ (p1 - p1))}
            _ = fi15
            _ = v33
        }
        v34 := GenS7{F0: p0, F1: 367, F2: -713}
        _ = v34
    } else {
        fmt.Println(p0, p1)
        if false {
            if (true || false) {
                for fi35 := 0; fi35 < 8; fi35++ {
                    for fi36 := 0; fi36 < 5; fi36++ {
                        p0 = p0
                        _ = fi36
                    }
                    var v37 uint8 = 127
                    _ = fi35
                    _ = v37
                }
                switch ("hello" + "foo bar") {
                case "hello":
                    v38 := []int{-9223372036854775808}
                    _ = v38
                }
            } else {
                for fi39 := 0; fi39 < 16; fi39++ {
                    var v40 int32 = 0
                    _ = fi39
                    _ = v40
                }
            }
        }
        var v41 float32 = 0.0
        _ = v41
    }
    var v42 float64 = ((1.5 + ((0.0 + -2.25) / 0.1)) + 0.1)
    _ = p0
    _ = p1
    _ = v42
    _ = v7
    return 128
}

func main() {
    v1 := []int{32768, 529, 9223372036854775807, -569}
    v2 := make(map[int]uint32)
    var v3 uint8 = 255
    var v4 int32 = 0
    v5 := make(map[uint32]bool)
    v6 := make(map[string]int)
    v43 := v6["hello"]
    v4 = 354
    fmt.Println(v4)
    var v44 int8 = 0
    v3 = 127
    v45 := GenS8{F0: v43}
    v3 += (v3 & (v3 * (v3 - v3)))
    v46 := GenS7{F0: 982, F1: 974, F2: v43}
    for fi47 := 0; fi47 < 12; fi47++ {
        v48 := GenS5{F0: (v3 ^ 0)}
        _ = fi47
        _ = v48
    }
    _ = len(v6)
    v44 += v44
    for fi49 := 0; fi49 < 19; fi49++ {
        var v50 int64 = -9223372036854775808
        var v51 int8 = v44
        _ = fi49
        _ = v50
        _ = v51
    }
    v3 = v3
    v44 -= -128
    for fi52 := 0; fi52 < 14; fi52++ {
        switch (("hello" + (("foo bar" + "foo bar") + "hello")) + (("hello" + "hello") + "foo bar")) {
        case "a":
            v44 -= (^(v44))
        default:
            var v53 float64 = (0.1 / 1.5)
            _ = v53
        }
        _ = fi52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
