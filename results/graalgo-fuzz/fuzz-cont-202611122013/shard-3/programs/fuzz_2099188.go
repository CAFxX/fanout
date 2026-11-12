package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int; F2 bool }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 906
    var v3 uint = 2
    v4 := make(map[uint32]string)
    var v5 string = "foo bar"
    fmt.Println(v2)
    v6 := v4[(^(v2))]
    var v7 float32 = 3.14159
    v3 = v3
    v8 := []int{-1}
    switch "hello" {
    case "":
        v3 += 902
    case "hello":
        if false {
            var v9 int16 = 32767
            switch -9223372036854775808 {
            case -1:
                _ = len(v8)
            default:
                v7 -= v7
            }
            v10 := GenS1{F0: 0, F1: ((42 > 61) || true), F2: 0}
            _ = v10
            _ = v9
        } else {
            fmt.Println(v6)
            var v11 uint = (v3 * (v3 & (v3 >> ((v3) & 63))))
            _ = v11
        }
    case "foo bar":
        fmt.Println(v3, v2)
    default:
        switch v5 {
        case "\n\t":
            v8 = append(v8, 42)
        }
    }
    var v12 string = v5
    fmt.Println(v3)
    v13 := GenS2{F0: -128}
    v14 := GenS2{F0: 100}
    if (float64(v7) > float64(v7)) {
        for fi15 := 0; fi15 < 15; fi15++ {
            if (((((float64(v7) - (1.5 - ((float64(((1e10 * ((v7 - v7) - (v7 - float32(0.0)))) / (0.1 - v7))) + ((3.14159 - 0.0) * (0.0 + (((1e10 / -2.25) * ((-2.25 * 1.5) - -2.25)) * (0.0 / (-2.25 / -2.25)))))) * 0.1))) * ((-2.25 - 1e10) / float64(((v7 + (((float32((((3.14159 * (1e10 + 1e10)) + 0.0) * float64(v7))) + float32(0.1)) * 3.14159) * 0.1)) * v7)))) < ((1.5 - 1e10) * 0.0)) || true) || true) {
                switch (((v5 + "\n\t") + "foo bar") + v6) {
                case "":
                    v16 := GenS3{F0: (fi15 & fi15), F1: fi15, F2: (v6 != "\n\t")}
                    _ = v16
                case "foo bar":
                    fmt.Println(v2, v12, v5)
                case "\n\t":
                    var v17 bool = false
                    _ = v17
                default:
                    switch (fi15 * (fi15 * 2147483647)) {
                    case 2147483647:
                        v18 := GenS3{F0: -994, F1: -1, F2: (4294967295 <= 42)}
                        _ = v18
                    }
                }
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 int8 = -126
                    v7 = -2.25
                    _ = fi19
                    _ = v20
                }
            } else {
                v3 = v3
                if (1 <= -9223372036854775808) {
                    v6 = v12
                    v2 = v2
                    var v21 uint8 = 42
                    _ = v21
                }
            }
            _ = fi15
        }
        v5 = ("a" + "\n\t")
    }
    var v22 int32 = -2147483648
    v23 := GenS1{F0: 596, F1: ((true || false) || true), F2: 32767}
    var v24 int = -311
    var v25 uint32 = (v2 - v2)
    if false {
        _ = len(v8)
        var v26 float64 = float64((-2.25 / (v7 / (v7 * (v7 - 1.5)))))
        _ = v26
    }
    v4[(v25 ^ v25)] = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
