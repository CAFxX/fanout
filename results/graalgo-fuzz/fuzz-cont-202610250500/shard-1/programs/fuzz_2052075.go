package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint, p1 int16) uint {
    var v7 int = 2147483647
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{(fi8 | fi8), fi8, 443, (fi8 - (v7 ^ (fi8 % (((v7 & (v7 / (((^(fi8))) | 1)))) | 1))))}
        _ = len(v9)
        v7 = (v7 % (((fi8 * v7)) | 1))
        _ = fi8
        _ = v9
    }
    v10 := []int{0, v7, v7, v7}
    v11 := GenS1{F0: 42, F1: "hello", F2: 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    return p0
}
func genFn1(p0 uint, p1 int64, p2 uint8) int {
    var v12 float64 = (1e10 * float64(1.5))
    for fi13 := 0; fi13 < 1; fi13++ {
        if true {
            p0 = (^(p0))
        }
        var v14 float32 = 0.1
        if false {
            v15 := GenS2{F0: (p0 / (((p0 * (p0 - p0))) | 1)), F1: false}
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(p0, fi16)
                v17 := GenS3{F0: 65535, F1: (^(p0))}
                v18 := GenS3{F0: 0, F1: genFn0(p0, 576)}
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 int = (fi13 | -655)
            _ = v15
            _ = v19
        } else {
            v20 := []int{(-(fi13)), 32768, fi13}
            var v21 uint64 = 743
            fmt.Println(p0, v12, v14)
            _ = v20
            _ = v21
        }
        _ = fi13
        _ = v14
    }
    p2 = p2
    p2 = p2
    v22 := GenS2{F0: (p0 * (p0 * genFn0(p0, 42))), F1: ((!(893 == 2)) && false)}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v22
    return int(p0)
}

func main() {
    var v1 uint = 2
    var v2 int16 = 0
    var v3 uint8 = 10
    var v4 int = 4294967295
    var v5 uint8 = (v3 << (((v1 << ((v1) & 63))) & 63))
    var v6 uint = v1
    fmt.Println(v5, v2, v3)
    v23 := []int{(v4 - genFn1(v1, -321, (v3 / ((v3) | 1))))}
    for fi24 := 0; fi24 < 4; fi24++ {
        if false {
            v25 := GenS4{F0: 844}
            v1 -= (v6 << (((v6 * v6)) & 63))
            _ = v25
        }
        _ = fi24
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v4 -= v4
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 uint16 = 231
            fmt.Println(v5)
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    v5 = v3
    if false {
        var v29 uint16 = 32767
        var v30 float32 = (-2.25 / (0.1 + 0.0))
        v31 := GenS1{F0: (v29 % ((v29) | 1)), F1: ("" + "foo bar"), F2: -1}
        _ = v29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        var v33 uint64 = 948
        for fi34 := 0; fi34 < 10; fi34++ {
            switch "foo bar" {
            case "foo bar":
                v35 := GenS5{F0: (-(v2))}
                var v36 string = ((("\n\t" + "foo bar") + "hello") + ((((("foo bar" + (((("\n\t" + "a") + ("" + "a")) + "\n\t") + ("" + ("foo bar" + "\n\t")))) + ("\n\t" + "hello")) + "foo bar") + "") + "foo bar"))
                _ = v35
                _ = v36
            case "hello":
                switch "a" {
                case "foo bar":
                    v33 += 2
                case "hello":
                    if (665 >= 65535) {
                        var v37 uint32 = 714
                        switch "" {
                        case "hello":
                            v6 -= (v6 ^ 2)
                        case "\n\t":
                            if true {
                                var v38 int = (fi34 + v4)
                                v39 := GenS4{F0: (v33 - v33)}
                                _ = v38
                                _ = v39
                            }
                        }
                        _ = v37
                    }
                }
            case "":
                var v40 int = (fi32 / (((v4 - fi32)) | 1))
                _ = v40
            }
            v33 = (v33 | (v33 % ((42) | 1)))
            var v41 uint16 = 985
            _ = fi34
            _ = v41
        }
        for fi42 := 0; fi42 < 16; fi42++ {
            v4 = (fi32 & -28)
            _ = fi42
        }
        _ = fi32
        _ = v33
    }
    v43 := GenS1{F0: uint16(v1), F1: ("hello" + ""), F2: -759}
    v44 := GenS6{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
