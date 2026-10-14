package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 string; F2 int32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint8, p1 uint16, p2 int64) uint64 {
    v4 := GenS1{F0: -125}
    p1 += 42
    v5 := []int{128, 61, 32768, 1}
    switch 128 {
    case -68:
        for fi6 := 0; fi6 < 13; fi6++ {
            switch (-(fi6)) {
            case 2147483647:
                v7 := GenS2{F0: 0, F1: "\n\t", F2: 0}
                _ = v7
            case 127:
                var v8 uint64 = 4294967295
                _ = v8
            case 32767:
                var v9 uint16 = p1
                _ = v9
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS1{F0: (fi6 >> ((926) & 63))}
                switch (fi6 & (fi6 >> ((uint(p0)) & 63))) {
                case 0:
                    v12 := GenS1{F0: fi6}
                    _ = v12
                default:
                    if ((p2 | (^(p2))) < p2) {
                        var v13 int16 = 32767
                        var v14 float64 = 1e10
                        v15 := GenS1{F0: (-(fi10))}
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        v16 := GenS3{F0: 42, F1: 2}
                        _ = v16
                    }
                }
                _ = fi10
                _ = v11
            }
            p1 += p1
            _ = fi6
        }
    }
    v17 := GenS4{F0: 1, F1: -944}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 uint64, p1 string) int16 {
    v18 := []int{-9223372036854775808, 42}
    v19 := GenS5{F0: p0, F1: -73}
    v20 := GenS4{F0: 37, F1: 42}
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v20
    return -1
}

func main() {
    var v1 int = 4294967295
    v2 := []int{}
    v3 := make(map[string]string)
    v21 := GenS3{F0: 107, F1: 18446744073709551615}
    fmt.Println(v1)
    delete(v3, "\n\t")
    if true {
        v1 -= v1
        v1 = 0
    } else {
        v22 := GenS3{F0: 720, F1: 864}
        _ = v22
    }
    v23 := GenS5{F0: genFn0(151, 0, 735), F1: 387}
    v1 -= (v1 & (v1 / (((v1 | (v1 - (v1 - 0)))) | 1)))
    for fi24 := 0; fi24 < 19; fi24++ {
        v1 = (^(fi24))
        var v25 int16 = 928
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 int32 = -2147483648
            var v28 float32 = 1e10
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi24
        _ = v25
    }
    v1 -= v1
    v29 := GenS4{F0: v1, F1: (v1 | 32767)}
    delete(v3, (("a" + ("a" + "foo bar")) + "\n\t"))
    for fi30 := 0; fi30 < 9; fi30++ {
        var v31 uint = 65535
        switch (fi30 >> ((v31) & 63)) {
        case 128:
            v1 = (v1 * fi30)
        }
        _ = fi30
        _ = v31
    }
    var v32 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v29
    _ = v3
    _ = v32
}
