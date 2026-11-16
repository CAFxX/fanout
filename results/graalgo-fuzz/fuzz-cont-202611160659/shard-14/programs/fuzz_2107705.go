package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint; F1 uint; F2 bool }
type GenS3 struct { F0 uint32 }
func genFn0() string {
    v6 := []int{32768}
    var v7 int64 = 434
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    v1 := []string{"", "", "hello"}
    var v2 uint = 309
    var v3 uint8 = 128
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    var v8 uint16 = 249
    v8 += v8
    switch "" {
    case "foo bar":
        v1 = append(v1, "a")
    case "a":
        for fi9 := 0; fi9 < 19; fi9++ {
            switch "hello" {
            case "":
                var v10 int = fi9
                _ = v10
            case "hello":
                v11 := v5[2147483647]
                _ = v11
            case "a":
                v8 = (v8 + (^(v8)))
            }
            var v12 uint64 = 4294967295
            v12 -= (^(v12))
            _ = fi9
            _ = v12
        }
    case "hello":
        delete(v5, 2)
    default:
        v13 := GenS1{F0: -2147483648, F1: 128}
        _ = v13
    }
    var v14 uint64 = 1
    v2 = (v2 % (((v2 - v2)) | 1))
    if false {
        fmt.Println(v3, v14)
    } else {
        for fi15 := 0; fi15 < 10; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                fmt.Println(v14)
                var v17 int32 = 2147483647
                var v18 uint32 = 2147483648
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi15
        }
        switch "foo bar" {
        case "a":
            v19 := GenS2{F0: v2, F1: uint(v3), F2: (65535 >= 1)}
            _ = v19
        case "foo bar":
            if (!false) {
                var v20 int16 = 42
                switch ("foo bar" + "foo bar") {
                case "a":
                    var v21 bool = true
                    _ = v21
                }
                fmt.Println(v14, v14)
                _ = v20
            } else {
                v22 := GenS3{F0: 4294967295}
                _ = v22
            }
        }
    }
    switch "" {
    case "a":
        if ((1000 > -1) || (42 == 32767)) {
            var v23 uint = v2
            v23 += (v23 / ((42) | 1))
            _ = v23
        } else {
            v14 = (v14 % ((4294967295) | 1))
            var v24 uint = (v2 | (^(v2)))
            _ = v24
        }
    case "hello":
        var v25 uint16 = (v8 & (v8 / (((v8 + 274)) | 1)))
        _ = v25
    }
    fmt.Println(v14)
    var v26 float64 = (0.0 + float64(float32(float64(0.1))))
    for fi27 := 0; fi27 < 15; fi27++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi27
    }
    var v28 uint32 = 2147483647
    var v29 uint = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
