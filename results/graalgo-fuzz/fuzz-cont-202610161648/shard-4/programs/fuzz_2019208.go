package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
type GenS2 struct { F0 int8; F1 uint8; F2 string }
type GenS3 struct { F0 int16; F1 uint }

func main() {
    var v1 int16 = 1
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    var v4 int = -1
    var v5 uint = 1
    var v6 int8 = -128
    fmt.Println(v1, v4, v5)
    v5 = 1
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            switch "" {
            case "\n\t":
                var v9 uint16 = 0
                _ = v9
            }
            v10 := GenS1{F0: 2147483647, F1: (v5 >> ((4294967295) & 63)), F2: v1}
            v11 := []int{551, fi7, v4}
            v12 := v2[v4]
            _ = fi8
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    v4 -= (^(v4))
    for fi13 := 0; fi13 < 9; fi13++ {
        switch "a" {
        case "":
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint16 = 42
                var v16 int64 = 0
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        _ = fi13
    }
    v4 -= (v4 * 32768)
    if false {
        v17 := GenS2{F0: (v6 << ((4294967295) & 63)), F1: 42, F2: "a"}
        if (!(!(true && (((^(v5)) <= (v5 % ((v5) | 1))) || (v6 < (v6 + v6)))))) {
            switch "\n\t" {
            case "foo bar":
                v18 := GenS3{F0: (v1 * (v1 & 828)), F1: v5}
                _ = v18
            case "a":
                fmt.Println(v4, v6)
            }
            switch (v4 + (v4 / ((v4) | 1))) {
            case 127:
                for fi19 := 0; fi19 < 17; fi19++ {
                    fmt.Println(fi19, v4)
                    v1 += (^(v1))
                    v20 := GenS2{F0: v6, F1: 255, F2: "foo bar"}
                    _ = fi19
                    _ = v20
                }
            default:
                var v21 int = -605
                _ = v21
            }
        } else {
            v1 = -32768
        }
        var v22 int64 = 866
        _ = v17
        _ = v22
    } else {
        _ = len(v2)
        fmt.Println(v4, v6)
        v5 = (v5 - v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
