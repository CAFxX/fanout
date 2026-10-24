package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint }
type GenS2 struct { F0 int16; F1 uint8; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS5 struct { F0 int }
func genFn0() uint16 {
    var v6 float64 = 0.1
    v7 := []int{-860, 9223372036854775807, -9223372036854775808}
    if false {
        fmt.Println(v6)
        v8 := GenS1{F0: (226 <= 2), F1: 450, F2: 18446744073709551615}
        v9 := GenS2{F0: 1, F1: 42, F2: 42}
        _ = v8
        _ = v9
    } else {
        var v10 uint64 = 1
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 1; fi12++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v13 int16 = 292
                _ = fi12
                _ = v13
            }
            v10 = v10
            _ = fi11
        }
        _ = v10
    }
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint) string {
    if (!true) {
        if false {
            if false {
                p0 += 2
            }
        } else {
            var v14 uint32 = 2147483647
            _ = v14
        }
        var v15 float64 = 0.0
        p0 = p0
        _ = v15
    }
    fmt.Println(p0)
    _ = p0
    return "a"
}

func main() {
    var v1 int16 = 0
    var v2 int8 = 76
    var v3 bool = (("a" <= "hello") || true)
    var v4 int = 4294967295
    v5 := []int{v4, (v4 | (v4 ^ (v4 >> ((193) & 63)))), 32767}
    var v16 float32 = -2.25
    switch v4 {
    case -1:
        v17 := GenS3{F0: (v3 && (!v3))}
        _ = v17
    default:
        _ = len(v5)
    }
    v4 = (v4 - v4)
    var v18 int64 = 421
    switch "a" {
    case "hello":
        var v19 float64 = ((1e10 / (3.14159 * 1.5)) / 1.0)
        _ = v19
    default:
        for fi20 := 0; fi20 < 13; fi20++ {
            fmt.Println(v3)
            _ = fi20
        }
    }
    switch (v4 - v4) {
    case 127:
        if v3 {
            v21 := GenS2{F0: (^(v1)), F1: 2, F2: 0}
            _ = v21
        }
    case 189:
        v22 := GenS2{F0: (v1 >> ((246) & 63)), F1: 1, F2: 65535}
        _ = v22
    }
    var v23 string = "hello"
    if ((^(v1)) >= (v1 | 610)) {
        if (((42 == 154) && v3) || true) {
            if v3 {
                v4 = (v4 % (((v4 | (-(v4)))) | 1))
                var v24 string = v23
                _ = v24
            }
        }
    }
    if v3 {
        fmt.Println(v2)
        v1 += (v1 ^ -32768)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        for fi25 := 0; fi25 < 18; fi25++ {
            v16 = v16
            _ = fi25
        }
        var v26 int = v4
        for fi27 := 0; fi27 < 7; fi27++ {
            v18 = -598
            _ = fi27
        }
        _ = v26
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int8 = -90
        _ = fi28
        _ = v29
    }
    v30 := GenS4{F0: 18446744073709551615, F1: 255, F2: (^(v2))}
    _ = len(v5)
    v31 := GenS5{F0: (v4 * 1)}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        fmt.Println(v16, v1, v23)
        if (v2 == 100) {
            var v33 bool = true
            if (127 < 127) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                v2 -= int8(v4)
                if ((v1 <= v1) || v3) {
                    if len(v5) > 0 {
                        v5[0] = v4
                    }
                } else {
                    v5 = append(v5, v4)
                }
                var v34 bool = false
                _ = v34
            }
            _ = v33
        } else {
            var v35 int = (fi32 % (((^(v4))) | 1))
            _ = v35
        }
        fmt.Println(v16, fi32, v23)
        _ = fi32
    }
    switch v23 {
    case "\n\t":
        v1 -= -1
    case "foo bar":
        v36 := GenS3{F0: false}
        _ = v36
    case "a":
        var v37 float32 = 0.1
        _ = v37
    }
    v18 -= (-(v18))
    v38 := GenS3{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
    _ = v5
}
