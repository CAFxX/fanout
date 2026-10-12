package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int32; F1 int16; F2 bool }
func genFn0() uint32 {
    var v6 int32 = -165
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{(fi7 ^ 40), (fi7 * (fi7 + fi7)), (fi7 >> ((4294967295) & 63))}
        v6 += v6
        _ = fi7
        _ = v8
    }
    _ = v6
    return 4294967295
}
func genFn1(p0 uint) int16 {
    if (true && false) {
        var v9 int16 = 32767
        _ = v9
    }
    if true {
        v10 := []int{-1, 32767, 2147483647, -9223372036854775808}
        _ = v10
    }
    v11 := GenS1{F0: 32767}
    v12 := GenS2{F0: 2147483647, F1: 9223372036854775807}
    fmt.Println(p0)
    _ = p0
    _ = v11
    _ = v12
    return 835
}

func main() {
    var v1 uint32 = 0
    var v2 int = 2147483647
    var v3 int16 = -752
    var v4 string = ""
    v5 := make(map[int]bool)
    if true {
        var v13 int8 = 105
        switch v4 {
        case "":
            for fi14 := 0; fi14 < 8; fi14++ {
                v15 := GenS2{F0: -1, F1: (-(v2))}
                v4 = v4
                v4 = v4
                _ = fi14
                _ = v15
            }
        case "\n\t":
            var v16 int64 = 9223372036854775807
            _ = v16
        case "a":
            var v17 int32 = 448
            _ = v17
        default:
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 uint64 = 18446744073709551615
                var v20 uint32 = (^(v1))
                v20 -= genFn0()
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        if false {
            fmt.Println(v1, v3, v3)
            var v21 uint16 = 255
            _ = v21
        } else {
            v1 = v1
            v2 -= v2
        }
        _ = v13
    } else {
        var v22 uint64 = 543
        v22 -= (v22 + v22)
        v22 = 0
        _ = v22
    }
    var v23 float64 = 0.0
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := []int{v2, 42, -2147483648, fi24}
        for fi26 := 0; fi26 < 3; fi26++ {
            v23 = float64(0.1)
            v1 -= 376
            var v27 int = fi24
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v1, v2)
    var v28 uint = 737
    for fi29 := 0; fi29 < 7; fi29++ {
        v30 := GenS2{F0: -2147483648, F1: (fi29 / ((v2) | 1))}
        v31 := GenS1{F0: 32767}
        _ = fi29
        _ = v30
        _ = v31
    }
    switch v4 {
    case "":
        _ = len(v5)
    case "hello":
        for fi32 := 0; fi32 < 20; fi32++ {
            if false {
                var v33 string = (v4 + v4)
                if (0 <= 678) {
                    fmt.Println(v28, v28)
                }
                _ = v33
            }
            v23 = (float64((0.1 - (float32(v23) * (1e10 + 1e10)))) - ((float64((1.5 * 3.14159)) + float64((1e10 + ((0.1 * ((float32(1.5) + ((float32(1e10) - -2.25) + (float32(v23) - -2.25))) / (-2.25 * (((0.1 * -2.25) * 3.14159) / 1.0)))) / -2.25)))) + (v23 / v23)))
            switch (v4 + v4) {
            case "\n\t":
                var v34 string = (v4 + ((v4 + v4) + ((v4 + "hello") + v4)))
                _ = v34
            }
            _ = fi32
        }
    case "\n\t":
        v3 += (v3 << (((v28 | 65535)) & 63))
    }
    if (1 < 42) {
        var v35 int16 = genFn1((v28 | (v28 / ((v28) | 1))))
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            var v36 float64 = v23
            _ = v36
        case "hello":
            if true {
                v35 += (v3 - (^(v3)))
                switch v2 {
                case -174:
                    switch "hello" {
                    case "foo bar":
                        v2 += 9223372036854775807
                    }
                case 32767:
                    v28 = (v28 + v28)
                case 9223372036854775807:
                    var v37 bool = false
                    _ = v37
                default:
                    v38 := []int{-407, (v2 >> (((v28 * 73)) & 63)), (v2 + (^(v2))), v2}
                    _ = v38
                }
                fmt.Println(v1)
            }
        case "foo bar":
            for fi39 := 0; fi39 < 12; fi39++ {
                v23 += ((v23 / -2.25) * (-2.25 * v23))
                _ = fi39
            }
        }
        _ = v35
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        v5[-9223372036854775808] = true
        v4 = (v4 + v4)
        v41 := GenS3{F0: 2147483647, F1: (v3 * (v3 % (((v3 & (v3 ^ 0))) | 1))), F2: (9223372036854775807 != v2)}
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 6; fi42++ {
        var v43 int = (fi42 << (((v28 % ((42) | 1))) & 63))
        if ((v2 ^ v43) != v43) {
            if true {
                v2 -= 251
                v23 -= (3.14159 + v23)
                v1 = genFn0()
            }
            v4 = "\n\t"
        } else {
            fmt.Println(v28, v1)
            v44 := GenS2{F0: 1, F1: (v43 ^ 4294967295)}
            _ = v44
        }
        _ = fi42
        _ = v43
    }
    v3 += (v3 + v3)
    v45 := GenS1{F0: 379}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
