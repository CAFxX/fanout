package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16; F1 int; F2 int8 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 string; F1 string; F2 uint8 }

func main() {
    v1 := []int{4294967295, 0, 9223372036854775807, -9223372036854775808}
    v2 := []int{-1, 32767, 2147483647}
    var v3 int32 = 0
    var v4 int64 = 9223372036854775807
    fmt.Println(v4)
    if (false || false) {
        var v5 int64 = v4
        _ = v5
    }
    v3 -= (v3 - v3)
    v3 += (v3 & v3)
    fmt.Println(v4, v4)
    var v6 float64 = (((-2.25 / (0.0 - 1e10)) - 1e10) + 0.1)
    fmt.Println(v4)
    v6 -= 1e10
    v1 = append(v1, 835)
    var v7 int8 = 0
    if true {
        v4 = (v4 * (v4 ^ 9223372036854775807))
        v8 := GenS1{F0: 32768, F1: 128}
        _ = v8
    } else {
        v4 = 42
        v6 = 0.0
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v2 = append(v2, fi9)
        switch (fi9 << ((2) & 63)) {
        case 0:
            var v10 int64 = v4
            _ = v10
        case 599:
            for fi11 := 0; fi11 < 10; fi11++ {
                if (922 > 0) {
                    if ((false || (true || true)) && false) {
                        var v12 int32 = v3
                        _ = v12
                    } else {
                        if true {
                            fmt.Println(fi11)
                        }
                        var v13 float64 = v6
                        _ = v13
                    }
                    v14 := GenS2{F0: 32767, F1: -1, F2: (^(v7))}
                    _ = v14
                }
                v3 = -1
                v15 := GenS3{F0: 42, F1: 42}
                _ = fi11
                _ = v15
            }
        }
        if (!true) {
            var v16 int32 = v3
            v7 += v7
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v2)
            } else {
                v17 := GenS4{F0: (v3 << ((42) & 63)), F1: 1}
                for fi18 := 0; fi18 < 13; fi18++ {
                    fmt.Println(v4)
                    var v19 uint16 = 2
                    v20 := GenS5{F0: 2, F1: (v7 % ((1) | 1))}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v21 := GenS6{F0: "hello", F1: "hello", F2: 128}
                _ = v17
                _ = v21
            }
            _ = v16
        } else {
            switch "\n\t" {
            case "foo bar":
                v2 = append(v2, -667)
            case "hello":
                _ = len(v1)
            case "a":
                var v22 uint8 = 197
                _ = v22
            }
        }
        _ = fi9
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
