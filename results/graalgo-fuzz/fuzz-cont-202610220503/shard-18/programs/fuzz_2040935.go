package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 bool; F1 int; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 int }
func genFn0(p0 int, p1 int32, p2 uint32) uint8 {
    switch 506 {
    case -2147483648:
        fmt.Println(p1)
    case 32768:
        v7 := GenS1{F0: 268, F1: 32}
        _ = v7
    }
    v8 := GenS1{F0: -9223372036854775808, F1: 928}
    p2 += (p2 % ((42) | 1))
    if true {
        v9 := []int{128, p0, 205, (^(p0))}
        switch ("foo bar" + "hello") {
        case "\n\t":
            var v10 int64 = -1
            _ = v10
        }
        p0 = p0
        _ = v9
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 uint16 = 42
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 127
}
func genFn1(p0 uint16) bool {
    v13 := []int{286, -2147483648, 9223372036854775807, 2147483647}
    if true {
        v13 = append(v13, 311)
    } else {
        for fi14 := 0; fi14 < 10; fi14++ {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            p0 = p0
            if ((!(1 <= (^(p0)))) && false) {
                var v15 uint16 = (p0 | 2)
                v15 = 121
                for fi16 := 0; fi16 < 12; fi16++ {
                    if len(v13) > 0 {
                        v13[0] = (fi16 - (fi16 / (((fi14 / ((fi14) | 1))) | 1)))
                    }
                    if true {
                        var v17 float64 = float64((0.1 + 0.1))
                        v18 := GenS1{F0: -453, F1: 2}
                        v15 += (p0 | (v15 / ((65535) | 1)))
                        _ = v17
                        _ = v18
                    }
                    var v19 uint16 = 32767
                    if len(v13) > 0 {
                        v13[0] = -2147483648
                    }
                    _ = fi16
                    _ = v19
                }
                _ = v15
            }
            _ = fi14
        }
        switch (("hello" + ("\n\t" + "hello")) + "a") {
        case "hello":
            var v20 uint64 = 0
            _ = v20
        case "":
            v21 := GenS2{F0: true, F1: 32768, F2: 4294967295}
            _ = v21
        case "\n\t":
            for fi22 := 0; fi22 < 7; fi22++ {
                p0 -= p0
                switch "" {
                case "":
                    p0 -= (p0 % ((p0) | 1))
                case "a":
                    var v23 uint16 = (p0 << ((65535) & 63))
                    _ = v23
                case "\n\t":
                    if false {
                        for fi24 := 0; fi24 < 13; fi24++ {
                            p0 += 32767
                            p0 = (p0 % (((p0 + (p0 / (((p0 ^ 42)) | 1)))) | 1))
                            _ = fi24
                        }
                        fmt.Println(p0, fi22)
                        v25 := GenS1{F0: 134, F1: 42}
                        _ = v25
                    }
                default:
                    var v26 bool = (1 != 510)
                    _ = v26
                }
                _ = fi22
            }
        }
    }
    if false {
        for fi27 := 0; fi27 < 17; fi27++ {
            p0 = p0
            if ((473 < fi27) || false) {
                fmt.Println(p0, fi27)
                if len(v13) > 0 {
                    v13[0] = -9223372036854775808
                }
                v28 := GenS3{F0: 0, F1: 0}
                _ = v28
            }
            _ = fi27
        }
    }
    p0 = 275
    for fi29 := 0; fi29 < 17; fi29++ {
        p0 -= p0
        p0 = p0
        _ = fi29
    }
    _ = p0
    _ = v13
    return true
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"foo bar"}
    var v3 bool = (3.14159 < 0.0)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 2
    v6 := []uint32{699, 406}
    _ = len(v6)
    v3 = (!v3)
    v30 := GenS4{F0: "foo bar"}
    var v31 uint32 = 247
    if len(v6) > 0 {
        v6[0] = v31
    }
    v32 := GenS5{F0: 255}
    v33 := GenS6{F0: v3, F1: -9223372036854775808}
    v34 := GenS5{F0: 662}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
