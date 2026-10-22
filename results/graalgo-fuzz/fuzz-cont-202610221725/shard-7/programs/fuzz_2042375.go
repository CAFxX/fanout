package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint8; F1 string; F2 uint16 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []uint32{42, 2147483647}
    var v2 int64 = 0
    var v3 int8 = -1
    v4 := make(map[int]uint32)
    var v5 int32 = -1
    v6 := []string{("" + "\n\t"), "hello", "hello"}
    fmt.Println(v3, v2, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2, v3)
    _ = len(v6)
    switch -929 {
    case -1:
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint8 = 30
            _ = fi7
            _ = v8
        }
    case -9223372036854775808:
        switch ("\n\t" + "") {
        case "\n\t":
            if ((((715 <= 65535) && (-628 < v5)) || (1.5 < (0.0 * (0.0 / 0.1)))) || (412 != 4294967295)) {
                v9 := GenS1{F0: 3, F1: 42}
                v10 := GenS2{F0: 25, F1: "\n\t", F2: 0}
                _ = v10
                _ = v9
            } else {
                v5 = (v5 / ((v5) | 1))
                v11 := GenS2{F0: 191, F1: "", F2: 2}
                _ = v11
            }
        case "hello":
            v5 = (v5 % ((v5) | 1))
        default:
            if true {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            } else {
                var v12 int = 1
                var v13 uint64 = 0
                _ = v12
                _ = v13
            }
        }
    case -442:
        var v14 uint = 18446744073709551615
        _ = v14
    default:
        v2 -= v2
    }
    v6 = append(v6, "hello")
    v1 = append(v1, 4294967295)
    v3 += v3
    if true {
        fmt.Println(v5, v2)
    } else {
        v15 := GenS1{F0: 2, F1: 628}
        if false {
            for fi16 := 0; fi16 < 3; fi16++ {
                v2 = (v2 / (((v2 / (((v2 << ((12) & 63))) | 1))) | 1))
                fmt.Println(fi16, v3, v2)
                v17 := GenS3{F0: int64(v5)}
                _ = fi16
                _ = v17
            }
            fmt.Println(v5, v3, v2)
        } else {
            v18 := GenS3{F0: (v2 ^ (-(v2)))}
            var v19 uint64 = 4294967295
            _ = v18
            _ = v19
        }
        _ = v15
    }
    switch "foo bar" {
    case "":
        switch 127 {
        case -270:
            var v20 uint = 314
            _ = v20
        case 128:
            if (0 < 42) {
                v2 -= v2
                v5 = 1
            }
        }
    }
    var v21 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
