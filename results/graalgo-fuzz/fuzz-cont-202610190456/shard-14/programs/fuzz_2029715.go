package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 int64 }

func main() {
    var v1 int64 = -503
    var v2 float32 = 3.14159
    v3 := make(map[int]int)
    var v4 int64 = 9223372036854775807
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 2147483648}
        v7 := GenS2{F0: 39, F1: 127, F2: -234}
        var v8 int32 = -1
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    fmt.Println(v4, v1)
    v9 := GenS3{F0: 2147483647, F1: true}
    fmt.Println(v1, v1)
    v10 := GenS4{F0: 18446744073709551615}
    var v11 uint64 = 108
    v12 := GenS3{F0: 42, F1: true}
    v13 := v3[32767]
    v14 := GenS1{F0: 0}
    v15 := GenS5{F0: 788, F1: -1}
    for fi16 := 0; fi16 < 15; fi16++ {
        switch "hello" {
        case "hello":
            var v17 bool = ("" > "foo bar")
            _ = v17
        case "foo bar":
            fmt.Println(v4, v1)
        case "\n\t":
            switch "\n\t" {
            case "hello":
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := []int{}
                    v1 += v4
                    v13 = (v13 % (((fi18 ^ -522)) | 1))
                    var v20 float32 = v2
                    v2 = 0.1
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case "":
                switch (^(v13)) {
                case 42:
                    var v21 int32 = -551
                    _ = v21
                case -809:
                    for fi22 := 0; fi22 < 11; fi22++ {
                        fmt.Println(v2, v11, v11)
                        v23 := []int{4294967295, (fi22 / ((v13) | 1)), fi16}
                        v4 = v4
                        _ = fi22
                        _ = v23
                    }
                default:
                    v24 := []int{(v13 >> ((1) & 63)), 0, v13}
                    _ = v24
                }
            case "\n\t":
                if (2 > 255) {
                    var v25 string = ""
                    _ = v25
                }
            default:
                var v26 float32 = (((v2 + 0.0) * float32((-2.25 + -2.25))) + 1e10)
                _ = v26
            }
        }
        var v27 uint32 = 65535
        v4 -= 0
        _ = fi16
        _ = v27
    }
    var v28 float64 = (1e10 * 1e10)
    var v29 uint16 = 1
    v30 := []int{4294967295, -1, (^(v13))}
    _ = len(v30)
    var v31 uint64 = uint64(v29)
    v4 = (v1 >> ((852) & 63))
    var v32 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v9
}
