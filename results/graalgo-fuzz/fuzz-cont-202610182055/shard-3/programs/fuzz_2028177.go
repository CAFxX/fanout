package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int }

func main() {
    var v1 uint64 = 926
    var v2 float64 = -2.25
    var v3 string = "foo bar"
    var v4 int16 = -252
    fmt.Println(v1, v4, v2)
    v4 = (v4 & (v4 << ((4294967295) & 63)))
    v4 = 1000
    var v5 int64 = int64(v4)
    var v6 int32 = 27
    var v7 uint = 65535
    fmt.Println(v1)
    if ((0.1 + 0.1) == (float32((v2 + 1e10)) / 1e10)) {
        for fi8 := 0; fi8 < 13; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                v2 -= v2
                if (1e10 <= ((1e10 * 1.5) / v2)) {
                    v10 := []int{(fi8 >> ((v7) & 63))}
                    _ = v10
                }
                v7 = 449
                _ = fi9
            }
            _ = fi8
        }
        if true {
            fmt.Println(v6, v3, v5)
            if false {
                var v11 int8 = 127
                _ = v11
            } else {
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 bool = ((v7 + (v7 * (v7 >> (((v7 | (v7 & (^(v7))))) & 63)))) == (v7 << (((v7 - 864)) & 63)))
                    v6 += (-(v6))
                    _ = fi12
                    _ = v13
                }
            }
            var v14 int = 127
            _ = v14
        }
    } else {
        v15 := []int{9223372036854775807, 264, -2147483648}
        if (!true) {
            v16 := GenS1{F0: v7, F1: (v5 / ((689) | 1)), F2: -697}
            var v17 float64 = (float64(-2.25) - (v2 * (0.0 * ((v2 * -2.25) / v2))))
            _ = v16
            _ = v17
        } else {
            v6 -= v6
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 string = v3
                _ = fi18
                _ = v19
            }
            if true {
                v20 := GenS1{F0: v7, F1: v5, F2: 1}
                _ = v20
            } else {
                fmt.Println(v3, v5, v7)
                v1 = v1
            }
        }
        _ = v15
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 float64 = float64((1e10 - (1e10 * 1e10)))
        v23 := []int{(fi21 * (^(fi21))), (fi21 >> ((v7) & 63)), -2147483648}
        var v24 uint32 = 2
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    v7 = 743
    if true {
        var v25 float32 = 0.1
        _ = v25
    }
    v26 := GenS1{F0: 65535, F1: (v5 + v5), F2: -88}
    v7 += v7
    for fi27 := 0; fi27 < 7; fi27++ {
        switch v3 {
        case "a":
            var v28 int = (fi27 ^ fi27)
            _ = v28
        case "":
            v1 = (v1 & (v1 >> ((945) & 63)))
        case "\n\t":
            var v29 bool = (v7 > v7)
            _ = v29
        default:
            fmt.Println(v6, fi27)
        }
        v4 = v4
        if true {
            v30 := []int{32767, 32767, (fi27 << ((v7) & 63)), 127}
            _ = v30
        }
        _ = fi27
    }
    if ((v7 | v7) < (v7 ^ 4294967295)) {
        v31 := []int{731}
        v1 -= 203
        if (32767 != 0) {
            v7 = v7
        }
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
