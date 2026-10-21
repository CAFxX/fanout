package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int64; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint }

func main() {
    var v1 int32 = -801
    v2 := make(map[string]string)
    var v3 uint8 = 255
    v4 := make(map[int]string)
    var v5 uint = 65535
    var v6 uint16 = 1
    var v7 uint32 = 4294967295
    switch "" {
    case "":
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(fi8, v5)
            v6 = (^(v6))
            v1 -= v1
            _ = fi8
        }
    case "hello":
        v9 := GenS1{F0: 966}
        _ = v9
    case "\n\t":
        fmt.Println(v7)
    default:
        switch 256 {
        case -9223372036854775808:
            switch -238 {
            case 9223372036854775807:
                if true {
                    switch -9223372036854775808 {
                    case 4294967295:
                        var v10 int = 719
                        _ = v10
                    case -1:
                        var v11 uint16 = 91
                        _ = v11
                    case 127:
                        switch 1 {
                        case -1:
                            fmt.Println(v6, v3)
                        case -444:
                            var v12 uint8 = 60
                            _ = v12
                        case -366:
                            v13 := GenS1{F0: 32768}
                            _ = v13
                        default:
                            v3 = v3
                        }
                    }
                }
            case 128:
                var v14 uint8 = 123
                _ = v14
            }
        default:
            if ((-93 != 0) || (-1 >= 9223372036854775807)) {
                v15 := []int{128}
                _ = v15
            } else {
                fmt.Println(v7, v7)
                fmt.Println(v7, v3, v7)
                for fi16 := 0; fi16 < 20; fi16++ {
                    v17 := []int{(fi16 & (fi16 / (((-(fi16))) | 1))), fi16, fi16, int(v1)}
                    v6 -= v6
                    _ = fi16
                    _ = v17
                }
            }
        }
    }
    var v18 int = -653
    v18 = 128
    v19 := GenS2{F0: v7, F1: -1, F2: 991}
    v20 := GenS1{F0: v18}
    var v21 int16 = 42
    var v22 uint16 = v6
    v18 -= int(v6)
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := v4[fi24]
            delete(v4, (v18 ^ fi23))
            v4[475] = v25
            var v26 float32 = (1.5 + 0.1)
            v22 -= uint16(v7)
            if (false || true) {
                v27 := []int{(-(fi24)), 1}
                v28 := v2["foo bar"]
                _ = v27
                _ = v28
            }
            delete(v2, v25)
            v29 := []int{(fi23 / (((v18 * (fi24 | (fi24 - (^(v18)))))) | 1)), (-(fi23))}
            _ = fi24
            _ = v25
            _ = v26
            _ = v29
        }
        switch (fi23 % ((fi23) | 1)) {
        case 110:
            v6 += 2
        case 127:
            switch fi23 {
            case 32767:
                switch "\n\t" {
                case "hello":
                    v30 := GenS2{F0: 2147483647, F1: -9223372036854775808, F2: 443}
                    _ = v30
                case "":
                    v31 := GenS3{F0: false}
                    _ = v31
                default:
                    for fi32 := 0; fi32 < 16; fi32++ {
                        if false {
                            fmt.Println(v22)
                            var v33 float32 = 0.0
                            if true {
                                fmt.Println(v33)
                            }
                            _ = v33
                        } else {
                            v34 := GenS1{F0: (v18 * v18)}
                            v5 += (v5 & v5)
                            _ = v34
                        }
                        switch (("a" + ("" + "foo bar")) + ("\n\t" + "a")) {
                        case "":
                            var v35 int8 = -1
                            _ = v35
                        case "hello":
                            v7 = (v7 / ((974) | 1))
                        case "\n\t":
                            v1 += v1
                        }
                        _ = fi32
                    }
                }
            case 0:
                v36 := GenS4{F0: true, F1: 18446744073709551615}
                _ = v36
            case 32768:
                fmt.Println(v18, v3, v18)
            default:
                var v37 string = "a"
                _ = v37
            }
        case 2147483647:
            v38 := GenS5{F0: 255}
            _ = v38
        default:
            var v39 uint64 = 1
            _ = v39
        }
        for fi40 := 0; fi40 < 14; fi40++ {
            var v41 int64 = 189
            if (v7 < (v7 & v7)) {
                v2["a"] = ("" + "a")
                v42 := []int{-1, fi40, int(v21), v18}
                v5 = v5
                _ = v42
            }
            _ = fi40
            _ = v41
        }
        _ = fi23
    }
    var v43 uint64 = 4294967295
    var v44 int16 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
