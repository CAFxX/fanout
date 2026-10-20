package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 float64 = 1e10
    var v2 int8 = 54
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    v5 := []string{((("foo bar" + "hello") + "hello") + "a"), "hello", ""}
    var v6 int = 2147483647
    var v7 uint = 255
    v6 += (v6 << (((v7 << ((v7) & 63))) & 63))
    v6 = 83
    fmt.Println(v2)
    v8 := GenS1{F0: 42, F1: -1, F2: 42}
    v4[65535] = true
    v1 = 3.14159
    switch "" {
    case "":
        if (((42 == 1) || ((1000 < -1) || true)) || (!false)) {
            var v9 int = 366
            _ = v9
        } else {
            v10 := GenS2{F0: 1}
            if true {
                if (255 < 0) {
                    v11 := GenS1{F0: 5, F1: 42, F2: 1000}
                    _ = v11
                } else {
                    if (999 == 1) {
                        v12 := GenS2{F0: 2}
                        _ = v12
                    } else {
                        fmt.Println(v7, v1)
                    }
                }
                if (v2 < v2) {
                    v1 += v1
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                v7 -= (v7 | v7)
            }
            _ = v10
        }
    case "\n\t":
        if (1 < 18446744073709551615) {
            v13 := GenS1{F0: 42, F1: 0, F2: 32767}
            var v14 uint32 = 65535
            _ = v13
            _ = v14
        }
    case "a":
        if (true && (1 > 2147483647)) {
            if true {
                var v15 uint64 = 1
                _ = v15
            } else {
                fmt.Println(v2, v6, v1)
            }
        } else {
            fmt.Println(v6, v1)
        }
    default:
        if false {
            if (false || false) {
                v16 := v4[2147483647]
                fmt.Println(v16, v16, v2)
                if true {
                    var v17 uint16 = 2
                    _ = v17
                } else {
                    v18 := GenS3{F0: uint64(v7)}
                    v2 = v2
                    _ = v18
                }
                _ = v16
            }
            v5 = append(v5, "")
        }
    }
    var v19 int8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
