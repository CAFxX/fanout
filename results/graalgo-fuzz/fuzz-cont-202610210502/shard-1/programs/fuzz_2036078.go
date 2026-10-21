package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint64; F2 int }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    v3 := []string{"hello"}
    var v4 string = "a"
    var v5 int16 = -509
    var v6 int64 = 0
    v7 := GenS1{F0: v4}
    v4 = (v4 + (v4 + (v4 + (v4 + "a"))))
    for fi8 := 0; fi8 < 6; fi8++ {
        v5 -= -723
        v5 -= v5
        _ = fi8
    }
    var v9 int = 32767
    v10 := GenS2{F0: (v9 * 9223372036854775807), F1: 746, F2: v9}
    var v11 uint8 = 217
    v12 := GenS1{F0: v4}
    v9 -= v9
    var v13 float64 = 1e10
    v9 += 42
    switch v4 {
    case "":
        if true {
            var v14 float64 = float64((float32((((v13 / v13) * (1e10 - (v13 / 1e10))) + v13)) / 1.0))
            _ = v14
        } else {
            v9 += v9
            v15 := GenS3{F0: v9}
            _ = v15
        }
    case "a":
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 uint8 = v11
            if (2147483647 <= 0) {
                switch (fi16 >> ((65535) & 63)) {
                case 128:
                    v17 -= (v11 >> ((1) & 63))
                case -2147483648:
                    v18 := GenS3{F0: fi16}
                    _ = v18
                case 1:
                    var v19 uint16 = 676
                    _ = v19
                default:
                    var v20 int32 = -994
                    _ = v20
                }
                switch v4 {
                case "\n\t":
                    v21 := GenS1{F0: ("a" + (("" + v4) + v4))}
                    _ = v21
                case "hello":
                    if true {
                        v5 = (v5 + v5)
                        v11 += 188
                        fmt.Println(v9, v11, v17)
                    }
                case "foo bar":
                    var v22 bool = false
                    _ = v22
                }
            }
            v4 = v4
            _ = fi16
            _ = v17
        }
    case "\n\t":
        var v23 int64 = v6
        _ = v23
    }
    if (!(-1 < 720)) {
        v24 := v1[2]
        var v25 int64 = v6
        v24 = "a"
        _ = v24
        _ = v25
    }
    var v26 uint32 = 2147483648
    switch v9 {
    case 127:
        var v27 float32 = 0.0
        _ = v27
    default:
        var v28 uint64 = 0
        _ = v28
    }
    var v29 uint8 = (v11 << ((2) & 63))
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v4, v29)
    case "":
        if (!false) {
            switch v9 {
            case 4294967295:
                var v30 uint32 = (v26 >> ((1) & 63))
                _ = v30
            default:
                var v31 int64 = v6
                _ = v31
            }
        } else {
            v32 := GenS1{F0: (((v4 + (v4 + "foo bar")) + "a") + v4)}
            v6 = (v6 | (v6 % (((v6 / ((323) | 1))) | 1)))
            _ = len(v1)
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
