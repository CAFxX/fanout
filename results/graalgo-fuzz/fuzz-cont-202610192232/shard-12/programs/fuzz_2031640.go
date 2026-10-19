package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint16 }
func genFn0(p0 bool, p1 uint16, p2 uint8) int32 {
    switch 324 {
    case -736:
        p0 = (575 < 0)
    case 127:
        v6 := []int{}
        _ = v6
    case 42:
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 bool = false
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 1, F1: 1, F2: 65535}
    v10 := []int{-9223372036854775808, 128, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -54
}
func genFn1(p0 uint32, p1 int32) uint {
    v11 := GenS2{F0: p1}
    switch "a" {
    case "\n\t":
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 string = "\n\t"
            if false {
                fmt.Println(p1, fi12, p0)
            } else {
                v14 := []int{fi12, (^(fi12))}
                fmt.Println(v13, p1, p1)
                _ = v14
            }
            _ = fi12
            _ = v13
        }
    case "hello":
        v15 := []int{128, 577}
        _ = v15
    default:
        var v16 bool = (false || false)
        _ = v16
    }
    p0 = (^(p0))
    if true {
        var v17 uint = 243
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    var v3 int = 153
    v4 := make(map[int]string)
    var v5 uint = 2
    var v18 uint16 = 255
    v3 += (v3 + (v3 << ((v5) & 63)))
    if true {
        fmt.Println(v18, v18)
        var v19 uint32 = 4294967295
        _ = v19
    }
    var v20 uint32 = 2147483647
    v21 := GenS2{F0: genFn0(false, (v18 >> (((v5 | (v5 << ((v5) & 63)))) & 63)), 1)}
    if (!((v20 + 0) != (v20 / ((v20) | 1)))) {
        v20 -= 2
        for fi22 := 0; fi22 < 1; fi22++ {
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 int32 = 2147483647
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        var v25 uint = (^(v5))
        _ = v25
    } else {
        switch "" {
        case "hello":
            var v26 uint64 = 0
            _ = v26
        default:
            v27 := GenS3{F0: v3, F1: false}
            _ = v27
        }
    }
    var v28 float64 = 3.14159
    if false {
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 uint64 = 18446744073709551615
            _ = fi29
            _ = v30
        }
        v3 -= (v3 % ((1) | 1))
    } else {
        v31 := GenS4{F0: 127, F1: 42}
        if true {
            v18 = (v18 ^ (^(v18)))
        }
        _ = v31
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        v18 = v18
    case "":
        if true {
            v28 += (v28 - (((0.0 / v28) - v28) * float64(-2.25)))
            var v32 int64 = 0
            fmt.Println(v20, v5)
            _ = v32
        }
    }
    v28 -= ((v28 - (float64((1e10 + 0.1)) - 1e10)) / 1.5)
    for fi33 := 0; fi33 < 3; fi33++ {
        fmt.Println(fi33)
        fmt.Println(v5)
        v34 := GenS4{F0: 42, F1: 167}
        _ = fi33
        _ = v34
    }
    v35 := GenS5{F0: (v5 & v5), F1: v18}
    v20 = (^(v20))
    fmt.Println(v3, v20)
    for fi36 := 0; fi36 < 15; fi36++ {
        _ = len(v1)
        fmt.Println(v18, v20, v3)
        _ = fi36
    }
    switch ("a" + "\n\t") {
    case "hello":
        if true {
            var v37 uint = v5
            _ = v37
        } else {
            var v38 int = (v3 ^ (^(v3)))
            for fi39 := 0; fi39 < 13; fi39++ {
                var v40 bool = false
                var v41 int16 = 0
                var v42 uint8 = 2
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
            if false {
                if (42 == -2147483648) {
                    fmt.Println(v3, v5)
                    v43 := GenS2{F0: 46}
                    _ = v43
                } else {
                    v18 -= uint16(v5)
                }
                if true {
                    v3 += (v38 + (v38 / (((^(v3))) | 1)))
                    for fi44 := 0; fi44 < 9; fi44++ {
                        var v45 uint32 = 4294967295
                        _ = fi44
                        _ = v45
                    }
                }
            }
            _ = v38
        }
    default:
        fmt.Println(v3)
    }
    var v46 int8 = 64
    var v47 uint8 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
