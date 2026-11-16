package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint8 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = false
    var v3 int = 0
    var v4 uint64 = v1
    v5 := []uint32{4, 1, 2, 2147483648}
    v6 := []uint32{}
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v2 = (1.5 != 3.14159)
    v1 += (v4 + v4)
    var v10 int = (v3 - (^(v3)))
    if v2 {
        var v11 int = 2147483647
        _ = v11
    } else {
        for fi12 := 0; fi12 < 14; fi12++ {
            if (!v2) {
                var v13 int = -2147483648
                v3 -= 32768
                v14 := GenS1{F0: v3, F1: 787, F2: 0}
                _ = v13
                _ = v14
            }
            switch "" {
            case "hello":
                for fi15 := 0; fi15 < 12; fi15++ {
                    v1 += (v4 | (v1 >> ((65535) & 63)))
                    _ = fi15
                }
            case "foo bar":
                fmt.Println(v10, v1, fi12)
            case "":
                var v16 int8 = 1
                _ = v16
            }
            var v17 uint32 = 795
            _ = fi12
            _ = v17
        }
        v18 := GenS1{F0: -9223372036854775808, F1: 42, F2: 244}
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 uint16 = 42
        var v21 uint16 = (v20 * (v20 * (v20 << ((0) & 63))))
        _ = len(v6)
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: (-(v10)), F1: 42, F2: 42}
    v23 := GenS1{F0: 0, F1: 32767, F2: 68}
    switch "\n\t" {
    case "a":
        var v24 float64 = float64(-2.25)
        _ = v24
    default:
        v3 = v10
    }
    switch 778 {
    case 42:
        v5 = append(v5, 2147483647)
    case -918:
        var v25 uint8 = 42
        _ = v25
    }
    if (!false) {
        var v26 int32 = 0
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v26
    }
    var v27 int32 = 1
    for fi28 := 0; fi28 < 19; fi28++ {
        fmt.Println(v2, v2)
        if v2 {
            switch (v3 / ((v10) | 1)) {
            case -1:
                switch v3 {
                case 42:
                    var v29 uint64 = v4
                    _ = v29
                case -2147483648:
                    v10 = (fi28 * 202)
                default:
                    var v30 uint8 = 46
                    _ = v30
                }
            case -9223372036854775808:
                for fi31 := 0; fi31 < 1; fi31++ {
                    v10 -= (v10 ^ (fi28 / ((-890) | 1)))
                    v4 = (^(v1))
                    v32 := GenS1{F0: (fi31 & (v10 + (v3 & (-(fi28))))), F1: 0, F2: 107}
                    _ = fi31
                    _ = v32
                }
            case -2147483648:
                var v33 string = "hello"
                _ = v33
            default:
                v27 = (v27 & v27)
            }
            v27 = (v27 | 42)
        } else {
            if false {
                fmt.Println(v1, fi28, v2)
            }
        }
        fmt.Println(v4, v3, v2)
        _ = fi28
    }
    var v34 int64 = -121
    var v35 int32 = v27
    if v2 {
        if v2 {
            for fi36 := 0; fi36 < 13; fi36++ {
                _ = len(v6)
                fmt.Println(fi36, fi36)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi36
            }
        }
        var v37 float64 = 0.0
        for fi38 := 0; fi38 < 11; fi38++ {
            var v39 bool = true
            _ = fi38
            _ = v39
        }
        _ = v37
    }
    if v2 {
        v34 -= (v34 * v34)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
