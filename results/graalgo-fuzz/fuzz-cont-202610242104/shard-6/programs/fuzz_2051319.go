package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS3 struct { F0 int64; F1 string; F2 int }

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{2147483648, 363, 2, 2147483647}
    v3 := make(map[int]string)
    var v4 bool = false
    switch -107 {
    case 9223372036854775807:
        v5 := GenS1{F0: 2147483647}
        _ = v5
    }
    v6 := GenS1{F0: 2147483647}
    v7 := GenS1{F0: 0}
    var v8 int32 = 427
    var v9 int32 = 42
    if (0 > 0) {
        var v10 int = 32768
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int16 = 0
            switch v1 {
            case "a":
                v1 = v1
            default:
                var v13 uint64 = 672
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        _ = len(v2)
        switch (v1 + v1) {
        case "hello":
            if false {
                var v14 float64 = float64(-2.25)
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 uint32 = 2147483647
                    var v17 bool = false
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v18 int8 = -1
                _ = v14
                _ = v18
            }
        case "\n\t":
            v4 = (528 <= 0)
        }
        v1 = ((("hello" + "\n\t") + v1) + "foo bar")
    }
    if (float32((0.0 + 1e10)) == 0.0) {
        _ = len(v3)
        switch -9223372036854775808 {
        case 9223372036854775807:
            v19 := GenS2{F0: 0, F1: 2, F2: 1}
            _ = v19
        case 127:
            switch "a" {
            case "\n\t":
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            case "a":
                v20 := GenS3{F0: -1, F1: "a", F2: 32767}
                _ = v20
            default:
                var v21 int = 0
                _ = v21
            }
        case 991:
            fmt.Println(v9, v1)
        }
        var v22 int8 = -128
        _ = v22
    } else {
        for fi23 := 0; fi23 < 5; fi23++ {
            fmt.Println(fi23, v1)
            v2 = append(v2, 1)
            _ = fi23
        }
    }
    var v24 bool = (1 != 32767)
    var v25 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
