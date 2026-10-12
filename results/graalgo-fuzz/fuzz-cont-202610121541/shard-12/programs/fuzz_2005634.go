package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int32; F2 int32 }

func main() {
    var v1 uint16 = 323
    var v2 int8 = -1
    var v3 int8 = v2
    v4 := make(map[string]bool)
    var v5 float64 = 1.5
    var v6 int8 = v2
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(fi8)
        switch (fi8 | (fi8 / (((fi8 % (((fi8 | (fi8 << ((2) & 63)))) | 1))) | 1))) {
        case 127:
            if ((v5 / v5) < v5) {
                v1 += v1
                v9 := []int{fi8, 128, fi8}
                _ = len(v9)
                _ = v9
            }
        }
        _ = fi8
    }
    v10 := GenS1{F0: (v3 / ((-2) | 1)), F1: (v7 + (v7 + v7))}
    if (4294967295 != 18446744073709551615) {
        var v11 uint16 = v1
        _ = v11
    } else {
        v6 = v3
    }
    switch 568 {
    case -1:
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS2{F0: 42, F1: 42}
            v14 := GenS3{F0: 978}
            v7 = ""
            _ = fi12
            _ = v13
            _ = v14
        }
    case -328:
        switch -686 {
        case -2147483648:
            v2 = (^(v6))
        }
    case 32767:
        fmt.Println(v2, v5, v6)
    }
    v6 -= (v2 * 0)
    if true {
        v7 = ("" + "\n\t")
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS4{F0: v3, F1: 0, F2: 0}
        var v17 uint16 = (v1 / ((v1) | 1))
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
