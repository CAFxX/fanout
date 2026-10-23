package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 bool; F1 bool; F2 int16 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint64) int8 {
    p0 -= (p0 + 851)
    for fi4 := 0; fi4 < 2; fi4++ {
        p0 += (p0 + (p0 & 1))
        _ = fi4
    }
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int8 = -39
        var v7 int16 = -1
        v8 := GenS1{F0: "hello", F1: 2, F2: -9223372036854775808}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    p0 -= p0
    _ = p0
    return -128
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 255
    var v3 float32 = ((0.1 - -2.25) * 1.5)
    v9 := GenS2{F0: "", F1: (!(975 > 255))}
    v2 -= v2
    var v10 int64 = 0
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS2{F0: ("foo bar" + ("" + "hello")), F1: (2147483647 >= 0)}
        var v13 int8 = -1
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        switch "foo bar" {
        case "\n\t":
            if true {
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 uint = 0
                    _ = fi15
                    _ = v16
                }
            } else {
                var v17 float64 = (1e10 * 0.1)
                _ = v17
            }
        }
        if (v10 >= (v10 << ((675) & 63))) {
            switch (fi14 << ((255) & 63)) {
            case 128:
                switch 819 {
                case -9223372036854775808:
                    if true {
                        var v18 int = 2147483647
                        _ = v18
                    } else {
                        v19 := GenS1{F0: ("" + ("" + ("hello" + "hello"))), F1: 128, F2: 9223372036854775807}
                        var v20 int16 = 63
                        var v21 int32 = -942
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
            case 42:
                var v22 uint8 = 1
                _ = v22
            }
        }
        _ = fi14
    }
    v3 -= v3
    fmt.Println(v2)
    v23 := GenS3{F0: -1}
    var v24 uint32 = 737
    switch 32767 {
    case -2147483648:
        v25 := GenS4{F0: 18446744073709551615, F1: v24}
        _ = v25
    case 2147483647:
        if false {
            if ("hello" >= ("foo bar" + "hello")) {
                v26 := GenS3{F0: 42}
                v27 := GenS5{F0: true, F1: (true || true), F2: 32767}
                _ = v26
                _ = v27
            } else {
                _ = len(v1)
            }
        } else {
            switch -538 {
            case -2147483648:
                if false {
                    var v28 float32 = (1e10 - 1.5)
                    v3 = 3.14159
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        v30 := []int{}
                        var v31 int = fi29
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS6{F0: (v10 + (-(v10))), F1: 1}
                    var v33 float64 = 1e10
                    _ = v32
                    _ = v33
                }
            case 0:
                switch 113 {
                case -211:
                    v34 := GenS4{F0: 4294967295, F1: 829}
                    _ = v34
                default:
                    var v35 int16 = 632
                    _ = v35
                }
            case 4294967295:
                v3 = v3
            default:
                v2 = v2
            }
            var v36 uint64 = 447
            _ = v36
        }
    case 42:
        v3 -= ((v3 + ((v3 / v3) * v3)) + float32(0.1))
    default:
        if true {
            var v37 int16 = 242
            _ = v37
        }
    }
    v38 := GenS7{F0: v24}
    v39 := GenS7{F0: (v24 % ((v24) | 1))}
    for fi40 := 0; fi40 < 8; fi40++ {
        v41 := GenS1{F0: "a", F1: 169, F2: v10}
        v1[346] = v24
        _ = fi40
        _ = v41
    }
    var v42 int32 = -794
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v38
    _ = v39
    _ = v42
    _ = v9
}
