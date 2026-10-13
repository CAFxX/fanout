package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint32 = 2
    var v2 uint8 = 38
    v3 := make(map[int]string)
    var v4 int8 = 127
    v4 = 100
    fmt.Println(v4, v2)
    for fi5 := 0; fi5 < 1; fi5++ {
        if false {
            var v6 int64 = 9223372036854775807
            _ = v6
        } else {
            v7 := GenS1{F0: 0}
            _ = v7
        }
        switch (fi5 ^ (fi5 * (^(fi5)))) {
        case 980:
            var v8 uint32 = (v1 / ((2147483647) | 1))
            _ = v8
        case -15:
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 string = "\n\t"
                switch "\n\t" {
                case "hello":
                    v11 := GenS1{F0: 42}
                    _ = v11
                case "":
                    switch (fi9 % (((fi5 * 641)) | 1)) {
                    case 32767:
                        var v12 uint64 = 669
                        _ = v12
                    }
                case "foo bar":
                    if (!false) {
                        fmt.Println(v10)
                    }
                }
                if (1 > (fi5 >> ((uint(v1)) & 63))) {
                    switch (v10 + v10) {
                    case "\n\t":
                        switch 128 {
                        case 630:
                            switch ((v10 + "hello") + v10) {
                            case "a":
                                v4 = v4
                            }
                            v13 := GenS2{F0: (fi5 | int(v1))}
                            _ = v13
                        }
                    case "hello":
                        v4 = v4
                    case "a":
                        v14 := GenS1{F0: 336}
                        _ = v14
                    default:
                        v1 = v1
                    }
                } else {
                    switch "foo bar" {
                    case "a":
                        v15 := GenS2{F0: (fi5 & fi5)}
                        if true {
                            v16 := GenS3{F0: 42, F1: 1}
                            var v17 int64 = 884
                            _ = v16
                            _ = v17
                        }
                        _ = v15
                    case "hello":
                        if (407 < -2147483648) {
                            var v18 uint16 = 160
                            var v19 int16 = 0
                            v20 := GenS2{F0: fi5}
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            v21 := GenS4{F0: -9223372036854775808, F1: int32(v1)}
                            v1 = v1
                            _ = v21
                        }
                    case "\n\t":
                        v1 -= (v1 & 4294967295)
                        var v22 int32 = 1
                        _ = v22
                    default:
                        var v23 float32 = 1.5
                        _ = v23
                    }
                }
                _ = fi9
                _ = v10
            }
        }
        var v24 uint16 = 0
        _ = fi5
        _ = v24
    }
    v1 += v1
    if false {
        v2 = (v2 - v2)
    } else {
        v4 -= (v4 ^ v4)
    }
    var v25 int = 637
    switch (v25 ^ v25) {
    case 0:
        if true {
            delete(v3, v25)
        } else {
            delete(v3, (^(v25)))
            v25 += v25
            v2 -= (^(v2))
        }
    case -2147483648:
        v26 := []int{(v25 >> ((1) & 63)), int(v4), 655, v25}
        _ = v26
    default:
        var v27 float64 = 1.5
        _ = v27
    }
    v25 = (v25 % (((v25 ^ (v25 * v25))) | 1))
    v28 := []int{}
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
