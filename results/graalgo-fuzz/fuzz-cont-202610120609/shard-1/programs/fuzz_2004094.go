package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int }
func genFn0(p0 string) int {
    var v6 uint16 = 2
    if ((^(v6)) == 266) {
        if (65535 >= 2147483648) {
            var v7 string = p0
            v8 := []int{-694}
            var v9 int64 = -9223372036854775808
            _ = v7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 409, F1: v6}
        _ = v10
    } else {
        if true {
            v11 := GenS2{F0: 2, F1: 1}
            _ = v11
        } else {
            switch (p0 + p0) {
            case "a":
                switch 128 {
                case 665:
                    var v12 bool = (true && ("a" >= (p0 + p0)))
                    _ = v12
                case 0:
                    var v13 float64 = 1.5
                    _ = v13
                case 144:
                    var v14 int16 = 1
                    _ = v14
                default:
                    switch "a" {
                    case "foo bar":
                        v15 := []int{1}
                        _ = v15
                    case "\n\t":
                        switch 280 {
                        case -2147483648:
                            fmt.Println(v6, v6)
                        default:
                            for fi16 := 0; fi16 < 12; fi16++ {
                                fmt.Println(fi16, fi16, v6)
                                _ = fi16
                            }
                        }
                    }
                }
            case "":
                v17 := []int{}
                _ = v17
            case "hello":
                p0 = (("" + "") + "\n\t")
            }
            if true {
                if (p0 == (p0 + ("" + p0))) {
                    v18 := GenS1{F0: 42, F1: v6}
                    _ = v18
                } else {
                    fmt.Println(v6)
                    v19 := []int{-9223372036854775808, -517, 128, 4294967295}
                    for fi20 := 0; fi20 < 4; fi20++ {
                        var v21 string = p0
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
                v22 := GenS3{F0: 42, F1: 908, F2: -1}
                _ = v22
            } else {
                v6 = 923
                for fi23 := 0; fi23 < 10; fi23++ {
                    var v24 bool = true
                    var v25 uint16 = v6
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                fmt.Println(v6, p0)
            }
        }
        v6 = (^(v6))
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 string = "hello"
        var v28 int32 = -1
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        v30 := []int{fi29}
        _ = fi29
        _ = v30
    }
    v31 := GenS3{F0: 1, F1: 1, F2: -202}
    _ = p0
    _ = v31
    _ = v6
    return -397
}
func genFn1(p0 uint) bool {
    switch -239 {
    case 127:
        fmt.Println(p0)
    case 846:
        if (42 >= 127) {
            fmt.Println(p0)
            p0 -= p0
        }
    default:
        var v32 uint32 = 42
        _ = v32
    }
    var v33 uint = (^(p0))
    _ = p0
    _ = v33
    return false
}

func main() {
    v1 := []uint32{42, 1}
    v2 := []int{-2147483648, 9223372036854775807}
    v3 := make(map[uint32]bool)
    var v4 uint32 = 530
    v5 := make(map[uint32]uint32)
    var v34 uint = 405
    for fi35 := 0; fi35 < 10; fi35++ {
        if (42 != 426) {
            _ = len(v2)
        } else {
            v4 = v4
        }
        _ = fi35
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        switch "\n\t" {
        case "hello":
            delete(v5, (v4 | v4))
        case "foo bar":
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 uint = v34
                var v39 uint32 = v4
                _ = fi37
                _ = v38
                _ = v39
            }
        default:
            for fi40 := 0; fi40 < 11; fi40++ {
                var v41 int32 = -2147483648
                var v42 float64 = 3.14159
                v42 = v42
                _ = fi40
                _ = v41
                _ = v42
            }
        }
        _ = fi36
    }
    fmt.Println(v34, v4)
    var v43 uint64 = 18446744073709551615
    fmt.Println(v4, v43)
    v43 -= v43
    v34 = 18446744073709551615
    v43 = v43
    var v44 float32 = (((1e10 - (1e10 + 0.0)) * 0.1) * float32(3.14159))
    v45 := GenS1{F0: (v4 - v4), F1: 32767}
    for fi46 := 0; fi46 < 2; fi46++ {
        var v47 int16 = -416
        v43 = v43
        v47 += (v47 / (((v47 ^ (v47 | (^(v47))))) | 1))
        _ = fi46
        _ = v47
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
