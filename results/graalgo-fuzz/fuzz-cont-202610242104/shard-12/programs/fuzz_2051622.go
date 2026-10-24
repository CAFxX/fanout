package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int8; F1 bool; F2 int32 }
func genFn0(p0 bool, p1 int) int64 {
    var v5 string = "hello"
    var v6 uint32 = 2147483647
    var v7 float64 = float64(1e10)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -911
}
func genFn1(p0 uint64) int16 {
    for fi8 := 0; fi8 < 1; fi8++ {
        switch "\n\t" {
        case "a":
            var v9 uint32 = 4294967295
            _ = v9
        case "foo bar":
            p0 = (p0 << ((18446744073709551615) & 63))
        }
        _ = fi8
    }
    if (false || ((0.1 - 1e10) != 0.1)) {
        v10 := []int{9223372036854775807, -1, 938, 42}
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int16 = -1
            _ = fi11
            _ = v12
        }
        p0 = 719
        _ = v10
    }
    _ = p0
    return 314
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 479
    v3 := make(map[int]string)
    var v4 uint16 = (^(v2))
    var v13 int8 = -114
    if true {
        var v14 int32 = -1
        v2 -= v2
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS1{F0: 42, F1: 30}
            if (-32768 > -32768) {
                v17 := []int{2147483647, fi15, fi15}
                v17 = append(v17, (-(fi15)))
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        fmt.Println(v13, v4, v4)
        var v18 int16 = -1
        _ = v18
    }
    v19 := GenS2{F0: -981, F1: 4294967295}
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(v13, v13, v2)
        v2 -= (v2 - (^(v4)))
        _ = fi20
    }
    var v21 int64 = int64(v13)
    v4 += 1
    v22 := []int{}
    fmt.Println(v1, v1)
    switch "hello" {
    case "hello":
        v23 := GenS2{F0: int32(v4), F1: 539}
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        for fi25 := 0; fi25 < 13; fi25++ {
            v26 := v3[(fi25 >> ((0) & 63))]
            _ = fi25
            _ = v26
        }
        switch fi24 {
        case 353:
            v27 := GenS1{F0: 540, F1: v21}
            _ = v27
        case 4294967295:
            switch (fi24 - fi24) {
            case 32767:
                _ = len(v22)
            case -9223372036854775808:
                if (false && ("\n\t" < (("foo bar" + "foo bar") + "a"))) {
                    if false {
                        if true {
                            v28 := GenS3{F0: int32(v1), F1: (v13 ^ v13)}
                            for fi29 := 0; fi29 < 20; fi29++ {
                                v30 := GenS4{F0: (v13 >> ((0) & 63)), F1: (false || true), F2: -2147483648}
                                _ = fi29
                                _ = v30
                            }
                            _ = v28
                        } else {
                            v31 := GenS2{F0: 42, F1: 2}
                            var v32 int64 = 1
                            var v33 bool = (2 < 788)
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        v22 = append(v22, fi24)
                    }
                    var v34 int64 = (v21 | (-(v21)))
                    _ = v34
                } else {
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                }
            default:
                switch (fi24 - fi24) {
                case 32768:
                    var v35 string = ("foo bar" + "\n\t")
                    _ = v35
                default:
                    var v36 uint32 = 2147483648
                    _ = v36
                }
            }
        default:
            v13 -= (v13 ^ (v13 & (v13 + (v13 >> ((270) & 63)))))
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
