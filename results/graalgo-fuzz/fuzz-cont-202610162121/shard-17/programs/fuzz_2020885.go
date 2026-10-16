package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]string)
    var v3 int = 128
    var v4 uint8 = 2
    if false {
        v1 = "foo bar"
        var v5 int64 = -9223372036854775808
        _ = v5
    } else {
        if (!(-9223372036854775808 > 1)) {
            fmt.Println(v4, v4, v3)
            v6 := GenS1{F0: v3, F1: v3, F2: (v4 != (v4 >> ((0) & 63)))}
            _ = v6
        } else {
            fmt.Println(v3, v3, v3)
        }
        if false {
            fmt.Println(v4)
            switch v3 {
            case 128:
                switch (v3 | (v3 % (((v3 & 42)) | 1))) {
                case -9223372036854775808:
                    for fi7 := 0; fi7 < 8; fi7++ {
                        v8 := GenS2{F0: -1}
                        _ = fi7
                        _ = v8
                    }
                case 9223372036854775807:
                    v9 := GenS3{F0: true, F1: -2147483648}
                    _ = v9
                }
            case 4294967295:
                var v10 uint32 = 870
                _ = v10
            }
            v4 = 127
        } else {
            fmt.Println(v1, v1)
            v4 = v4
            switch v1 {
            case "\n\t":
                v2[2147483648] = v1
            case "hello":
                v3 += -9223372036854775808
            default:
                switch "hello" {
                case "foo bar":
                    v11 := []int{}
                    _ = v11
                }
            }
        }
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := []int{v3, v3, (^(fi12)), -2147483648}
        _ = fi12
        _ = v13
    }
    v14 := []int{}
    var v15 string = ""
    var v16 int8 = 127
    var v17 uint8 = 1
    if ((255 > 0) && false) {
        v18 := GenS3{F0: (!true), F1: 2147483647}
        _ = v18
    } else {
        v19 := GenS1{F0: v3, F1: (v3 << ((65535) & 63)), F2: false}
        _ = v19
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
