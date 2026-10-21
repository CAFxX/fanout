package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 bool) string {
    fmt.Println(p0)
    v4 := GenS1{F0: 1}
    p0 = (-943 < 32767)
    _ = p0
    _ = v4
    return (((("foo bar" + (("a" + "hello") + "a")) + "a") + "\n\t") + "hello")
}
func genFn1() int16 {
    v5 := []int{}
    v5 = append(v5, -190)
    _ = v5
    return 464
}

func main() {
    v1 := []int{1, 444}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 2
    if (233 <= 255) {
        for fi6 := 0; fi6 < 10; fi6++ {
            if ((fi6 >> ((813) & 63)) > 127) {
                v3 -= (v3 * (v3 + (^(v3))))
            }
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := GenS1{F0: 188}
                v3 = 4294967295
                v3 = v3
                _ = fi7
                _ = v8
            }
            var v9 float64 = ((1e10 + 1e10) + 0.0)
            _ = fi6
            _ = v9
        }
        if true {
            _ = len(v2)
            v3 -= (^(v3))
            delete(v2, 65535)
        }
        fmt.Println(v3)
    }
    if len(v1) > 0 {
        v1[0] = 128
    }
    var v10 int8 = 100
    switch ((("hello" + "") + "foo bar") + "a") {
    case "\n\t":
        v3 += v3
    case "a":
        var v11 int = 32768
        _ = v11
    case "foo bar":
        var v12 uint8 = 114
        _ = v12
    default:
        v10 = -126
    }
    v10 = -128
    var v13 int = 4294967295
    if (1000 >= -1) {
        for fi14 := 0; fi14 < 13; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                fmt.Println(fi14, v13)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch ("\n\t" + "hello") {
                case "a":
                    v16 := GenS2{F0: 2, F1: 2, F2: 255}
                    _ = v16
                case "\n\t":
                    if (((463 > 789) && false) && ((42 < 2147483647) || true)) {
                        if (false || true) {
                            v1 = append(v1, 9223372036854775807)
                            if false {
                                v10 = 1
                                v13 -= 32768
                                fmt.Println(fi14)
                                var v17 uint64 = (v3 >> ((246) & 63))
                                v3 = 169
                                _ = v17
                            }
                        }
                        fmt.Println(v3, v3)
                        v18 := GenS3{F0: true}
                        _ = v18
                    }
                }
                if len(v1) > 0 {
                    v1[0] = (fi15 / (((-(v13))) | 1))
                }
                var v19 int64 = -1
                _ = fi15
                _ = v19
            }
            var v20 int32 = 0
            _ = fi14
            _ = v20
        }
    } else {
        v10 = (v10 & (v10 * -1))
        if len(v1) > 0 {
            v1[0] = v13
        }
        fmt.Println(v13, v3, v13)
    }
    v21 := GenS1{F0: 182}
    v22 := GenS2{F0: 42, F1: 2, F2: 617}
    var v23 uint8 = 127
    for fi24 := 0; fi24 < 13; fi24++ {
        if true {
            v25 := GenS2{F0: 65535, F1: uint32(v23), F2: 42}
            _ = v25
        }
        v26 := GenS1{F0: 0}
        var v27 uint32 = 860
        _ = fi24
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 bool = ((760 < 0) && true)
        _ = fi28
        _ = v29
    }
    var v30 int16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
}
