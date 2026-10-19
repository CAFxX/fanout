package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint64 }

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 uint16 = 472
    v5 := []uint32{}
    v1 = -294
    if ((782 > 2147483647) && false) {
        v6 := GenS1{F0: 1, F1: 2147483648}
        fmt.Println(v1)
        var v7 int64 = (v1 & -983)
        _ = v6
        _ = v7
    }
    if (true && true) {
        if (2147483647 >= -564) {
            var v8 int8 = 100
            fmt.Println(v4)
            _ = v8
        } else {
            fmt.Println(v1, v4)
            var v9 uint32 = 2147483648
            _ = v9
        }
    } else {
        if len(v5) > 0 {
            v5[0] = 2147483648
        }
    }
    v4 -= 65535
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v10 uint32 = 2147483648
    v4 = (v4 ^ (v4 + v4))
    switch 720 {
    case 9223372036854775807:
        v1 -= v1
    case -1:
        var v11 uint32 = 65535
        _ = v11
    case -9223372036854775808:
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 string = ""
            _ = fi12
            _ = v13
        }
    }
    v1 += v1
    v10 = 226
    v4 += v4
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS2{F0: v4}
        if true {
            v3[(("" + ("hello" + ("" + ("hello" + "")))) + "")] = true
            var v16 int = 106
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        v10 = 42
        for fi18 := 0; fi18 < 6; fi18++ {
            if false {
                v10 += v10
                switch "foo bar" {
                case "\n\t":
                    v4 = 32767
                case "a":
                    var v19 int = (fi18 - (fi17 ^ (fi18 - (fi17 ^ 2147483647))))
                    _ = v19
                case "hello":
                    v4 = (v4 & v4)
                default:
                    if true {
                        v20 := GenS3{F0: -669, F1: 18446744073709551615}
                        _ = v20
                    }
                }
            } else {
                fmt.Println(fi18)
                v21 := GenS2{F0: 253}
                var v22 bool = ((("" + "a") + "hello") == "foo bar")
                _ = v21
                _ = v22
            }
            switch "" {
            case "\n\t":
                fmt.Println(fi18, v4, v1)
            case "hello":
                var v23 int64 = v1
                v5 = append(v5, v10)
                switch "a" {
                case "foo bar":
                    _ = len(v2)
                case "":
                    v1 -= 433
                default:
                    var v24 string = "a"
                    _ = v24
                }
                _ = v23
            }
            v25 := GenS1{F0: 2147483647, F1: 42}
            _ = fi18
            _ = v25
        }
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 uint16 = (v4 / (((v4 | (v4 / (((v4 * 204)) | 1)))) | 1))
            _ = fi26
            _ = v27
        }
        _ = fi17
    }
    v10 = v10
    v28 := GenS3{F0: -1, F1: 206}
    switch -405 {
    case -2147483648:
        fmt.Println(v4, v10, v1)
    case 127:
        var v29 uint8 = 2
        _ = v29
    default:
        for fi30 := 0; fi30 < 4; fi30++ {
            v31 := GenS1{F0: 0, F1: 1}
            fmt.Println(v4, v4)
            _ = fi30
            _ = v31
        }
    }
    var v32 int64 = (v1 >> ((255) & 63))
    v33 := GenS3{F0: 0, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
