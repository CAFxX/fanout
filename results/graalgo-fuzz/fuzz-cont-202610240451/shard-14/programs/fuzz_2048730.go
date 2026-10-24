package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }
type GenS2 struct { F0 string }
func genFn0(p0 int, p1 string, p2 uint64) uint8 {
    p2 -= p2
    var v6 int32 = 2147483647
    switch "a" {
    case "hello":
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 127
}
func genFn1(p0 uint16, p1 uint8) uint8 {
    if (9223372036854775807 >= 2147483647) {
        v7 := GenS1{F0: 926, F1: p1, F2: 1}
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(fi8, p1, p0)
        _ = fi8
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 0
    var v3 int8 = 127
    var v4 int16 = -31
    var v5 uint8 = 127
    var v9 uint = 331
    if (!(42 != 0)) {
        v3 = v3
        var v10 int8 = v3
        switch 0 {
        case 32767:
            v2 = v2
        default:
            var v11 int = -200
            _ = v11
        }
        _ = v10
    } else {
        var v12 uint32 = 65535
        var v13 float32 = -2.25
        var v14 int64 = -9223372036854775808
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 float32 = 0.1
    v16 := v1[-394]
    var v17 int16 = (v4 * (v4 - (v4 / ((42) | 1))))
    if (716 != 2) {
        for fi18 := 0; fi18 < 6; fi18++ {
            v15 = (v15 + v15)
            switch "hello" {
            case "foo bar":
                v19 := v1[fi18]
                _ = v19
            case "a":
                v9 += (^(v9))
            case "":
                switch (^(fi18)) {
                case 1:
                    delete(v1, (fi18 >> ((4294967295) & 63)))
                }
            default:
                for fi20 := 0; fi20 < 7; fi20++ {
                    fmt.Println(v4)
                    if (uint(v5) < 0) {
                        switch 326 {
                        case 128:
                            if (65535 != v9) {
                                v21 := GenS1{F0: -1, F1: (v5 >> ((v9) & 63)), F2: 42}
                                _ = v21
                            }
                            fmt.Println(v9, v9)
                            var v22 int8 = v3
                            _ = v22
                        case 32767:
                            v2 -= 571
                            var v23 bool = (32767 <= v17)
                            v2 -= v2
                            _ = v23
                        case -488:
                            v15 = (v15 + (0.0 + v15))
                        }
                        var v24 float64 = -2.25
                        v25 := GenS1{F0: 0, F1: (v5 - 42), F2: -683}
                        _ = v24
                        _ = v25
                    } else {
                        switch ("" + "foo bar") {
                        case "\n\t":
                            v26 := []int{(fi18 * (fi20 ^ (-(fi18)))), (-(fi20)), -803}
                            _ = v26
                        case "hello":
                            fmt.Println(v9, fi18)
                        default:
                            v17 += v17
                        }
                    }
                    var v27 float32 = 1e10
                    if v16 {
                        v4 = (^(v4))
                    } else {
                        v28 := GenS1{F0: -151, F1: v5, F2: -519}
                        fmt.Println(v16, fi18)
                        _ = v28
                    }
                    v29 := v1[42]
                    _ = fi20
                    _ = v27
                    _ = v29
                }
            }
            v30 := GenS2{F0: "foo bar"}
            _ = fi18
            _ = v30
        }
        if false {
            v31 := []int{}
            for fi32 := 0; fi32 < 18; fi32++ {
                var v33 int16 = 42
                var v34 bool = false
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = v31
        } else {
            var v35 int16 = (v4 / ((42) | 1))
            _ = v35
        }
    } else {
        v9 -= (v9 << ((v9) & 63))
    }
    if (v3 >= (-(v3))) {
        var v36 int8 = v3
        delete(v1, 773)
        _ = v36
    }
    switch -1 {
    case 4294967295:
        var v37 int = 128
        _ = v37
    case -725:
        fmt.Println(v2)
    case 864:
        var v38 uint = (v9 ^ v9)
        _ = v38
    default:
        v9 = v9
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        v40 := []int{}
        _ = fi39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
