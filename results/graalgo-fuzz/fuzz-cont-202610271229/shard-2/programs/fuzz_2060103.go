package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 int }
func genFn0(p0 int) uint {
    p0 = (p0 | (p0 & -9223372036854775808))
    v7 := GenS1{F0: -1, F1: 2147483648}
    _ = p0
    _ = v7
    return 4294967295
}
func genFn1(p0 uint64) string {
    p0 += p0
    for fi8 := 0; fi8 < 7; fi8++ {
        if (-1 == 377) {
            p0 = (^(p0))
        } else {
            switch "" {
            case "":
                v9 := []int{}
                _ = v9
            }
            switch fi8 {
            case 9223372036854775807:
                p0 -= p0
            case 2147483647:
                var v10 uint64 = p0
                _ = v10
            case 374:
                if false {
                    v11 := GenS1{F0: fi8, F1: 42}
                    var v12 uint8 = 43
                    _ = v11
                    _ = v12
                } else {
                    v13 := []int{127, (fi8 ^ (-(fi8)))}
                    _ = v13
                }
            default:
                p0 += p0
            }
            fmt.Println(fi8)
        }
        _ = fi8
    }
    _ = p0
    return ""
}

func main() {
    var v1 uint8 = 1
    var v2 uint16 = 42
    var v3 uint16 = 0
    var v4 int16 = -1
    var v5 int32 = 0
    var v6 uint = 276
    v14 := GenS2{F0: 1, F1: (v3 & (v3 * 1)), F2: -816}
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(v3, v4)
        switch "" {
        case "foo bar":
            v1 += 219
        case "hello":
            var v16 int32 = -2147483648
            _ = v16
        case "a":
            v2 -= (v3 & (v2 >> ((0) & 63)))
        default:
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 uint = genFn0(-1)
                var v19 int8 = 127
                var v20 float64 = float64(1e10)
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            v22 := []int{}
            _ = len(v22)
            fmt.Println(fi21, v4, v3)
            _ = fi21
            _ = v22
        }
        _ = fi15
    }
    fmt.Println(v1, v2)
    v6 = (v6 & v6)
    fmt.Println(v1, v3, v5)
    if (!(-865 <= 76)) {
        v1 += (v1 & (^(v1)))
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        switch fi23 {
        case -942:
            switch 4294967295 {
            case -9223372036854775808:
                var v24 uint8 = (v1 + (v1 + v1))
                _ = v24
            default:
                fmt.Println(v4, v4)
            }
        case 9223372036854775807:
            if ((^(v6)) == 218) {
                v5 = v5
            } else {
                if true {
                    v1 = v1
                }
                v25 := []int{}
                _ = v25
            }
        case 1:
            if (true || false) {
                var v26 string = "a"
                v6 = 930
                v3 = 583
                _ = v26
            }
        }
        _ = fi23
    }
    switch -857 {
    case -166:
        v27 := GenS2{F0: -45, F1: (^(v3)), F2: 2147483647}
        _ = v27
    }
    v3 = 65535
    if false {
        var v28 bool = true
        _ = v28
    } else {
        var v29 bool = false
        _ = v29
    }
    switch "\n\t" {
    case "foo bar":
        var v30 uint16 = (v3 - v3)
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
