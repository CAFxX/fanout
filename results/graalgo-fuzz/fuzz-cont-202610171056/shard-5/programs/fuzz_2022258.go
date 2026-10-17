package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint32; F1 int; F2 int8 }
func genFn0() string {
    v4 := []int{-9223372036854775808, 0, -609}
    var v5 int64 = 883
    var v6 int = -267
    var v7 uint8 = 144
    fmt.Println(v7)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := []uint32{2}
    v2 := []uint32{2147483648, 484}
    v3 := make(map[int]bool)
    var v8 float64 = (-2.25 / float64(0.1))
    v9 := GenS1{F0: ("" + "a"), F1: 1}
    if true {
        if (1 <= 1) {
            for fi10 := 0; fi10 < 2; fi10++ {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 uint8 = 186
                    fmt.Println(fi11, fi11)
                    var v13 uint = 42
                    v12 -= (v12 - v12)
                    v12 = (v12 % ((v12) | 1))
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = len(v1)
                delete(v3, fi10)
                _ = fi10
            }
            var v14 uint = 18
            _ = v14
        } else {
            var v15 float64 = float64((3.14159 - 1.5))
            v15 = v15
            for fi16 := 0; fi16 < 1; fi16++ {
                v17 := GenS1{F0: "hello", F1: -94}
                _ = fi16
                _ = v17
            }
            _ = v15
        }
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        fmt.Println(fi18, v8)
        for fi19 := 0; fi19 < 14; fi19++ {
            v8 += v8
            var v20 uint64 = 18446744073709551615
            switch (fi18 >> ((1) & 63)) {
            case -9223372036854775808:
                fmt.Println(fi19, v20)
            }
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS2{F0: 800, F1: 1, F2: -28}
        if (597 < 2) {
            v8 += v8
            v3[(fi21 & fi21)] = true
            switch "" {
            case "\n\t":
                v8 = v8
            default:
                var v23 uint16 = 408
                _ = v23
            }
        } else {
            if (true && false) {
                v8 = ((1.5 / v8) / (((v8 / (-2.25 / v8)) - v8) - v8))
                fmt.Println(fi21, fi21)
            }
            v24 := GenS2{F0: 917, F1: (fi21 * fi21), F2: -6}
            var v25 uint32 = 2147483647
            _ = v24
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    if (2 >= 32767) {
        for fi26 := 0; fi26 < 20; fi26++ {
            for fi27 := 0; fi27 < 7; fi27++ {
                v2 = append(v2, 999)
                v8 -= 1e10
                _ = fi27
            }
            var v28 int64 = 42
            fmt.Println(fi26, v28)
            _ = fi26
            _ = v28
        }
        v8 -= v8
    }
    var v29 float32 = 0.1
    v29 += float32((v8 - (v8 * 1.5)))
    for fi30 := 0; fi30 < 2; fi30++ {
        var v31 int8 = -82
        _ = fi30
        _ = v31
    }
    var v32 float64 = 1e10
    if false {
        fmt.Println(v32, v29, v8)
    } else {
        delete(v3, -2147483648)
        if (!((!(true || true)) && (!(!(65535 <= 607))))) {
            _ = len(v2)
            _ = len(v3)
        } else {
            v8 = v8
        }
    }
    switch 1 {
    case -540:
        v33 := GenS2{F0: 2147483648, F1: 0, F2: 0}
        _ = v33
    case -47:
        for fi34 := 0; fi34 < 20; fi34++ {
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 int = (fi34 / ((-2147483648) | 1))
                switch "a" {
                case "foo bar":
                    var v37 int = (fi34 ^ 1)
                    _ = v37
                case "a":
                    v36 = 32767
                case "hello":
                    v29 += float32(1.5)
                    fmt.Println(fi35, fi34, v29)
                default:
                    v38 := v3[(-(v36))]
                    _ = v38
                }
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    default:
        if true {
            if (0 > 1) {
                var v39 int = 1
                for fi40 := 0; fi40 < 5; fi40++ {
                    v32 -= float64(v29)
                    _ = fi40
                }
                v3[v39] = false
                _ = v39
            }
            fmt.Println(v32, v32, v32)
        } else {
            for fi41 := 0; fi41 < 18; fi41++ {
                v29 += 3.14159
                _ = fi41
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v8
    _ = v9
}
