package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }

func main() {
    v1 := []int{-2147483648}
    var v2 int32 = 531
    v3 := []uint32{0, 618}
    var v4 uint32 = 2147483647
    var v5 float32 = -2.25
    v6 := []int{32767, -992}
    v5 = v5
    switch 1 {
    case 4294967295:
        var v7 float64 = (1e10 / float64(v5))
        _ = v7
    default:
        var v8 float32 = float32((1.5 * 1.5))
        _ = v8
    }
    if (32767 != 1000) {
        if false {
            for fi9 := 0; fi9 < 1; fi9++ {
                switch "hello" {
                case "a":
                    for fi10 := 0; fi10 < 14; fi10++ {
                        var v11 string = "foo bar"
                        fmt.Println(fi9)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi10
                        _ = v11
                    }
                }
                v12 := GenS1{F0: (fi9 - fi9), F1: 1}
                _ = fi9
                _ = v12
            }
            v4 = 2147483648
            v2 -= (v2 >> ((42) & 63))
        }
    } else {
        var v13 uint32 = v4
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = v4
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v14 := GenS1{F0: -203, F1: 42}
    for fi15 := 0; fi15 < 16; fi15++ {
        if (v5 == (float32(1.5) / -2.25)) {
            fmt.Println(v2, v4, fi15)
            var v16 int64 = -9223372036854775808
            _ = v16
        } else {
            v17 := GenS1{F0: -304, F1: 42}
            var v18 bool = false
            _ = v17
            _ = v18
        }
        var v19 uint32 = (v4 >> ((4294967295) & 63))
        _ = fi15
        _ = v19
    }
    switch 0 {
    case -184:
        if (!((false && false) && false)) {
            var v20 string = "foo bar"
            v2 -= v2
            var v21 uint16 = 144
            _ = v20
            _ = v21
        } else {
            if (!(1 == (v4 / ((2147483648) | 1)))) {
                for fi22 := 0; fi22 < 4; fi22++ {
                    for fi23 := 0; fi23 < 2; fi23++ {
                        var v24 int = -1
                        _ = fi23
                        _ = v24
                    }
                    _ = fi22
                }
            } else {
                fmt.Println(v4, v2, v5)
                v2 -= v2
                var v25 int32 = v2
                _ = v25
            }
            fmt.Println(v4, v2, v4)
        }
    case 229:
        v26 := GenS1{F0: -2147483648, F1: 42}
        _ = v26
    case 9223372036854775807:
        v3 = append(v3, v4)
    }
    var v27 int64 = 0
    v5 -= v5
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 float64 = 0.0
        fmt.Println(fi28, v29)
        _ = fi28
        _ = v29
    }
    if (!((-674 <= 0) && (2 > (v4 ^ v4)))) {
        if true {
            for fi30 := 0; fi30 < 8; fi30++ {
                v27 -= v27
                switch (fi30 * (fi30 >> ((497) & 63))) {
                case -1:
                    v31 := GenS1{F0: fi30, F1: 42}
                    _ = v31
                default:
                    v5 += v5
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    v4 += (v4 % (((v4 ^ (v4 >> ((1) & 63)))) | 1))
                    _ = fi32
                }
                _ = fi30
            }
            fmt.Println(v4, v27, v4)
        } else {
            v33 := GenS1{F0: 340, F1: 42}
            v6 = append(v6, -572)
            v34 := GenS1{F0: 127, F1: 96}
            _ = v33
            _ = v34
        }
    } else {
        var v35 int = 32767
        var v36 string = "hello"
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        if len(v6) > 0 {
            v6[0] = (fi37 | int(v2))
        }
        fmt.Println(v5)
        _ = fi37
    }
    fmt.Println(v4)
    switch "\n\t" {
    case "":
        v27 = (v27 / ((1) | 1))
    case "hello":
        switch -9223372036854775808 {
        case 42:
            var v38 int16 = 312
            _ = v38
        case 0:
            var v39 uint8 = 42
            _ = v39
        case 128:
            fmt.Println(v27, v2, v4)
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    if (false || true) {
        if ((v4 > (v4 >> ((1) & 63))) && false) {
            var v40 int = 802
            v2 = 42
            v40 -= (^(v40))
            _ = v40
        } else {
            if (1000 < 264) {
                if true {
                    var v41 float32 = v5
                    _ = v41
                }
                _ = len(v1)
            }
            v4 -= v4
        }
        if false {
            v27 += 798
            fmt.Println(v27)
        }
    }
    var v42 uint32 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
