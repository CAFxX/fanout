package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 bool; F1 string; F2 int32 }
func genFn0(p0 int) uint64 {
    p0 += -317
    v4 := GenS1{F0: 659}
    if true {
        v5 := GenS2{F0: 2, F1: 792}
        switch "hello" {
        case "a":
            p0 = 127
        case "":
            fmt.Println(p0)
        }
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 int8 = -125
            var v8 uint = 701
            if (2 >= 32767) {
                switch "\n\t" {
                case "hello":
                    for fi9 := 0; fi9 < 1; fi9++ {
                        v10 := []int{2147483647, fi9, fi6}
                        v7 += v7
                        v11 := GenS2{F0: 42, F1: 1}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                default:
                    var v12 uint8 = 128
                    _ = v12
                }
                p0 = p0
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    } else {
        v13 := GenS2{F0: 2, F1: -527}
        p0 -= p0
        _ = v13
    }
    _ = p0
    _ = v4
    return 2
}
func genFn1(p0 int32) int {
    if true {
        v14 := []int{0}
        p0 = 42
        _ = v14
    }
    var v15 int = 686
    fmt.Println(p0, p0)
    for fi16 := 0; fi16 < 17; fi16++ {
        p0 = (p0 + (p0 % ((1) | 1)))
        _ = fi16
    }
    var v17 int64 = 11
    _ = p0
    _ = v15
    _ = v17
    return v15
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 42
    var v3 uint = 354
    switch "hello" {
    case "":
        switch -477 {
        case 127:
            var v18 uint = uint(v2)
            _ = v18
        default:
            if (!true) {
                var v19 int32 = -1
                v19 += (v19 / ((42) | 1))
                _ = v19
            } else {
                v2 = v2
                v20 := GenS3{F0: (v3 + v3)}
                v21 := []int{-186, 128, 42, -985}
                _ = v20
                _ = v21
            }
        }
    case "a":
        var v22 bool = (1e10 <= ((0.0 - (1.5 / (-2.25 - 3.14159))) - 1.5))
        _ = v22
    default:
        var v23 int16 = 32767
        _ = v23
    }
    v24 := GenS2{F0: genFn0(-457), F1: 42}
    v3 += 0
    for fi25 := 0; fi25 < 2; fi25++ {
        switch -666 {
        case 128:
            switch genFn1(2147483647) {
            case -9223372036854775808:
                v26 := []int{genFn1(-2147483648), (^(fi25)), (^(fi25)), -82}
                _ = v26
            case 128:
                var v27 uint64 = genFn0(genFn1(-356))
                _ = v27
            case 32767:
                fmt.Println(v1, v1, fi25)
            default:
                if true {
                    v3 += v3
                    switch genFn1(-137) {
                    case -961:
                        v3 = (v3 - (v3 | v3))
                    case -757:
                        if (true && true) {
                            if (288 != 0) {
                                v1 = genFn0(fi25)
                                var v28 uint32 = 2
                                fmt.Println(v3, v2)
                                _ = v28
                            }
                            v2 += (v2 << (((v3 << (((v3 / (((v3 ^ (v3 + (v3 - (v3 / (((v3 * (v3 >> (((v3 >> ((v3) & 63))) & 63)))) | 1)))))) | 1))) & 63))) & 63))
                        }
                    case 42:
                        var v29 uint64 = (v1 & (^(v1)))
                        _ = v29
                    }
                } else {
                    v2 = v2
                    switch (-(fi25)) {
                    case 834:
                        var v30 int8 = -128
                        _ = v30
                    default:
                        var v31 uint8 = 62
                        _ = v31
                    }
                    var v32 int8 = -128
                    _ = v32
                }
            }
        }
        v33 := []int{genFn1(2147483647), 127, (-(fi25))}
        switch "a" {
        case "\n\t":
            var v34 bool = false
            _ = v34
        case "a":
            for fi35 := 0; fi35 < 1; fi35++ {
                v1 = 430
                v36 := GenS4{F0: 0, F1: (fi35 - 42)}
                v37 := GenS5{F0: (1 == 0), F1: ("hello" + ("" + "\n\t")), F2: -376}
                _ = fi35
                _ = v36
                _ = v37
            }
        case "hello":
            var v38 string = "a"
            _ = v38
        }
        _ = fi25
        _ = v33
    }
    v39 := []int{genFn1(2147483647)}
    for fi40 := 0; fi40 < 6; fi40++ {
        switch (fi40 | (fi40 << ((2) & 63))) {
        case -1:
            var v41 uint64 = v1
            _ = v41
        case -132:
            v42 := GenS2{F0: uint64(v2), F1: -1}
            _ = v42
        case 517:
            fmt.Println(v2)
        default:
            v2 = v2
        }
        _ = fi40
    }
    var v43 uint32 = 65535
    if false {
        var v44 bool = true
        _ = v44
    } else {
        v43 = 42
        v2 -= 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v43
}
