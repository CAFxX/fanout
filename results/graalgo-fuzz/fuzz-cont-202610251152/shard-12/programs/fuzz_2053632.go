package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
func genFn0(p0 int, p1 uint64) int16 {
    var v7 int = p0
    p0 += (v7 / (((p0 << ((65535) & 63))) | 1))
    _ = p0
    _ = p1
    _ = v7
    return -169
}
func genFn1() string {
    var v8 float32 = (1.5 + 1e10)
    v9 := []int{-561, 259}
    _ = v8
    _ = v9
    return ("a" + ("" + "\n\t"))
}

func main() {
    var v1 uint = 65535
    var v2 uint = 4294967295
    var v3 int16 = 1000
    var v4 int32 = 113
    v5 := make(map[int]uint32)
    v6 := make(map[string]bool)
    if true {
        var v10 int8 = 100
        _ = v10
    }
    v1 = v1
    if (-32768 > v3) {
        var v11 uint64 = 495
        var v12 uint16 = 0
        fmt.Println(v3, v4, v3)
        _ = v11
        _ = v12
    } else {
        switch -2147483648 {
        case 128:
            fmt.Println(v4, v3)
        case -2147483648:
            var v13 float64 = ((0.1 * 1.5) / 1.5)
            _ = v13
        case -9223372036854775808:
            if (!false) {
                v14 := GenS1{F0: "foo bar", F1: 32768}
                if false {
                    switch ("foo bar" + ("hello" + "foo bar")) {
                    case "\n\t":
                        fmt.Println(v1)
                    }
                } else {
                    fmt.Println(v2, v1)
                    for fi15 := 0; fi15 < 19; fi15++ {
                        for fi16 := 0; fi16 < 6; fi16++ {
                            var v17 uint64 = 4294967295
                            var v18 int32 = (v4 / (((v4 | (v4 << (((^(v2))) & 63)))) | 1))
                            fmt.Println(v4, v2)
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        if true {
                            var v19 uint16 = 492
                            for fi20 := 0; fi20 < 17; fi20++ {
                                var v21 int = (fi20 | 32768)
                                _ = fi20
                                _ = v21
                            }
                            _ = v19
                        }
                        _ = fi15
                    }
                    fmt.Println(v2)
                }
                _ = v14
            }
        }
        var v22 uint16 = 633
        if ((!(!(((2 > 198) && true) || true))) && (0 == 1)) {
            v1 = v2
        } else {
            v23 := GenS2{F0: 42, F1: 1, F2: (v2 << ((2) & 63))}
            v2 = (v1 << (((v1 * v1)) & 63))
            v24 := GenS2{F0: -755, F1: 2147483648, F2: v2}
            _ = v23
            _ = v24
        }
        _ = v22
    }
    var v25 int32 = v4
    if ((v3 % ((v3) | 1)) == v3) {
        for fi26 := 0; fi26 < 7; fi26++ {
            switch fi26 {
            case 32767:
                var v27 int16 = v3
                _ = v27
            case 636:
                fmt.Println(fi26)
            }
            _ = fi26
        }
        var v28 bool = true
        _ = v28
    }
    var v29 int = 127
    v25 = (-(v4))
    v1 -= 0
    if (2 < 0) {
        var v30 uint = 639
        v30 -= v1
        _ = len(v6)
        _ = v30
    }
    fmt.Println(v3, v1)
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
