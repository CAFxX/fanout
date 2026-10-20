package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint32 }
func genFn0() uint64 {
    v4 := GenS1{F0: 255, F1: 719, F2: 1000}
    var v5 uint = 0
    v6 := []int{2147483647, 2147483647, -178, 42}
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := []int{}
    var v2 int = -646
    var v3 int = (^(v2))
    v2 -= v3
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(fi7, v2)
        _ = fi7
    }
    v2 += 0
    if true {
        for fi8 := 0; fi8 < 12; fi8++ {
            switch (v2 + -806) {
            case 127:
                v2 -= 4294967295
            case 751:
                fmt.Println(fi8)
            default:
                v2 = -2147483648
            }
            if (-188 == 0) {
                v3 = v3
            }
            _ = fi8
        }
        if (1 <= 42) {
            switch v3 {
            case -780:
                var v9 uint64 = 782
                _ = v9
            }
            v10 := GenS2{F0: 2147483647, F1: 0, F2: 1000}
            var v11 int8 = -1
            _ = v10
            _ = v11
        } else {
            fmt.Println(v2, v2)
        }
        v3 += 42
    }
    var v12 uint = 18446744073709551615
    var v13 int64 = -1
    var v14 uint16 = 32767
    v13 = (v13 + -1)
    if true {
        switch (v2 - v2) {
        case 32768:
            var v15 uint = 18446744073709551615
            _ = v15
        default:
            v16 := GenS3{F0: (v13 & (v13 * v13)), F1: v14, F2: 2}
            _ = v16
        }
        if (65535 < 62) {
            var v17 string = "\n\t"
            _ = v17
        }
        for fi18 := 0; fi18 < 14; fi18++ {
            var v19 float32 = -2.25
            _ = fi18
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            _ = len(v1)
            v3 -= v3
            _ = fi20
        }
        v21 := GenS3{F0: (v13 >> ((v12) & 63)), F1: v14, F2: 2147483648}
        _ = v21
    }
    switch "\n\t" {
    case "":
        if len(v1) > 0 {
            v1[0] = (v3 ^ v2)
        }
    case "\n\t":
        if (705 > 195) {
            v22 := GenS1{F0: (v14 - 1), F1: (v13 >> ((v12) & 63)), F2: 817}
            var v23 int = (v2 / (((v2 - (v2 + (^(v2))))) | 1))
            _ = v22
            _ = v23
        } else {
            for fi24 := 0; fi24 < 10; fi24++ {
                if (!true) {
                    switch "foo bar" {
                    case "hello":
                        _ = len(v1)
                    case "a":
                        fmt.Println(fi24, v2)
                    default:
                        for fi25 := 0; fi25 < 20; fi25++ {
                            var v26 uint = v12
                            v2 -= (v3 / (((^(fi24))) | 1))
                            fmt.Println(fi24, v3)
                            _ = fi25
                            _ = v26
                        }
                    }
                    v14 += 32767
                    fmt.Println(v3, v14)
                } else {
                    fmt.Println(v2, v12)
                    var v27 uint = (v12 | v12)
                    if len(v1) > 0 {
                        v1[0] = fi24
                    }
                    _ = v27
                }
                v28 := GenS1{F0: v14, F1: v13, F2: 192}
                _ = fi24
                _ = v28
            }
            v12 = (v12 - 2)
        }
    case "foo bar":
        v29 := GenS3{F0: (v13 * (v13 >> (((v12 * 255)) & 63))), F1: v14, F2: 2}
        _ = v29
    }
    v30 := GenS1{F0: (v14 << (((v12 * (v12 ^ v12))) & 63)), F1: v13, F2: -1}
    var v31 bool = (2 >= uint64(v12))
    v31 = (!(42 != -84))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
