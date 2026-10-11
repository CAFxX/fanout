package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int8 = -52
    var v2 uint = 2
    var v3 uint = 18446744073709551615
    var v4 int32 = 0
    v2 = (v2 / (((v3 - v2)) | 1))
    v4 = v4
    fmt.Println(v4)
    for fi5 := 0; fi5 < 2; fi5++ {
        v4 = (v4 | v4)
        if true {
            fmt.Println(v2, fi5)
            v6 := []int{(fi5 / (((fi5 >> ((v3) & 63))) | 1))}
            fmt.Println(fi5, v2)
            _ = v6
        } else {
            v3 += (v3 * (^(v3)))
            var v7 int64 = 19
            switch (fi5 ^ fi5) {
            case 9223372036854775807:
                v3 = v3
            default:
                var v8 bool = (!false)
                _ = v8
            }
            _ = v7
        }
        _ = fi5
    }
    var v9 float64 = 1.5
    v2 -= (v2 / (((^(v3))) | 1))
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: 42, F1: 0}
        if (!false) {
            var v12 int = fi10
            var v13 int16 = -492
            switch "" {
            case "\n\t":
                var v14 float64 = ((v9 + v9) / v9)
                _ = v14
            default:
                if true {
                    v15 := []int{(fi10 / (((v12 ^ (fi10 & (v12 & v12)))) | 1))}
                    _ = v15
                }
            }
            _ = v12
            _ = v13
        } else {
            v16 := GenS2{F0: 255}
            var v17 uint8 = 42
            _ = v16
            _ = v17
        }
        _ = fi10
        _ = v11
    }
    v18 := []int{9223372036854775807, 42, 1}
    var v19 string = "a"
    for fi20 := 0; fi20 < 7; fi20++ {
        switch v19 {
        case "hello":
            v2 += 2
        case "":
            var v21 int = 32767
            _ = v21
        default:
            v4 = (v4 >> ((42) & 63))
        }
        _ = fi20
    }
    v22 := GenS2{F0: 255}
    v4 = (v4 / ((2147483647) | 1))
    fmt.Println(v4, v3)
    v18 = append(v18, 42)
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 int = (fi23 - fi23)
        switch "a" {
        case "a":
            v24 -= (v24 * v24)
        case "foo bar":
            v25 := GenS3{F0: 0}
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
