package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint; F2 bool }
type GenS3 struct { F0 int }
func genFn0(p0 uint64, p1 uint32, p2 uint32) uint32 {
    fmt.Println(p1, p0, p0)
    var v5 int64 = -996
    switch "hello" {
    case "foo bar":
        v6 := []int{int(p0)}
        _ = v6
    case "":
        var v7 uint16 = 255
        _ = v7
    default:
        if true {
            for fi8 := 0; fi8 < 13; fi8++ {
                switch ("a" + "hello") {
                case "":
                    v9 := GenS1{F0: 0}
                    _ = v9
                case "\n\t":
                    fmt.Println(p0)
                }
                v10 := GenS1{F0: 948}
                _ = fi8
                _ = v10
            }
            v5 -= (-(v5))
        } else {
            if (65535 >= 42) {
                v11 := GenS1{F0: 1000}
                for fi12 := 0; fi12 < 8; fi12++ {
                    p0 += (p0 >> ((93) & 63))
                    p1 = uint32(p0)
                    if ((103 < 0) || (false && false)) {
                        if (156 == int(p1)) {
                            var v13 bool = false
                            if (15 != 44) {
                                var v14 int16 = -207
                                for fi15 := 0; fi15 < 13; fi15++ {
                                    var v16 int32 = -2147483648
                                    _ = fi15
                                    _ = v16
                                }
                                v17 := []int{-359, -184, fi12}
                                _ = v14
                                _ = v17
                            } else {
                                switch (("foo bar" + "hello") + "") {
                                case "a":
                                    if (v13 && v13) {
                                        fmt.Println(p2, p0, p2)
                                        v5 -= v5
                                    } else {
                                        var v18 string = ("a" + "a")
                                        fmt.Println(v13)
                                        _ = v18
                                    }
                                }
                                v19 := GenS1{F0: -542}
                                for fi20 := 0; fi20 < 2; fi20++ {
                                    v21 := GenS2{F0: 2, F1: 1, F2: (1 != 0)}
                                    _ = fi20
                                    _ = v21
                                }
                                _ = v19
                            }
                            v22 := []int{}
                            _ = v13
                            _ = v22
                        } else {
                            fmt.Println(p2, v5)
                            var v23 int32 = -2147483648
                            _ = v23
                        }
                    }
                    _ = fi12
                }
                var v24 int32 = 2147483647
                _ = v11
                _ = v24
            }
            v5 = 0
            p0 -= p0
        }
    }
    v5 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    var v1 bool = (!(1 != 42))
    v2 := []uint32{2, 0, 484, 647}
    v3 := make(map[int]bool)
    var v4 float64 = -2.25
    v25 := v3[32768]
    v26 := GenS3{F0: 810}
    if len(v2) > 0 {
        v2[0] = genFn0(772, 42, 2147483647)
    }
    var v27 uint8 = 2
    switch 4294967295 {
    case 220:
        v1 = (0 <= 32767)
    }
    switch -55 {
    case -1:
        v27 = v27
    case 32767:
        var v28 uint32 = 2147483647
        _ = v28
    case 467:
        v25 = v1
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 string = "hello"
        if (69 >= -128) {
            var v31 int32 = 0
            _ = v31
        } else {
            switch (v30 + "") {
            case "\n\t":
                for fi32 := 0; fi32 < 5; fi32++ {
                    v25 = (42 < -82)
                    if len(v2) > 0 {
                        v2[0] = genFn0(2, 2147483648, 2147483647)
                    }
                    var v33 float64 = (v4 / 1.0)
                    _ = fi32
                    _ = v33
                }
            default:
                var v34 bool = (!(!(fi29 >= fi29)))
                _ = v34
            }
        }
        _ = fi29
        _ = v30
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 uint32 = 42
        _ = fi35
        _ = v36
    }
    v37 := v3[127]
    v38 := GenS3{F0: -2147483648}
    var v39 int = -2147483648
    var v40 string = "\n\t"
    v4 += (v4 + v4)
    var v41 float64 = v4
    var v42 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
