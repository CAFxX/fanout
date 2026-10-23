package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint64 }

func main() {
    var v1 int16 = 32767
    var v2 int8 = 1
    v3 := make(map[uint32]uint32)
    switch "hello" {
    case "":
        v1 = 1
    default:
        if true {
            if true {
                fmt.Println(v1)
            }
        } else {
            var v4 uint32 = 2147483647
            fmt.Println(v1, v1, v4)
            _ = v4
        }
    }
    v2 -= (v2 % (((v2 | v2)) | 1))
    switch "hello" {
    case "foo bar":
        v1 = (v1 << ((18446744073709551615) & 63))
    case "a":
        var v5 uint16 = 42
        _ = v5
    case "\n\t":
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint = 255
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 float64 = 1e10
                var v10 int64 = -681
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    v11 := GenS1{F0: -225}
    fmt.Println(v2, v2)
    var v12 uint8 = 60
    v13 := GenS2{F0: 2147483647, F1: v2, F2: 2}
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := []int{(fi14 / ((364) | 1))}
        _ = fi14
        _ = v15
    }
    fmt.Println(v1, v2, v12)
    fmt.Println(v1, v1, v12)
    v16 := GenS1{F0: 928}
    v2 = (v2 / (((v2 & v2)) | 1))
    v17 := GenS2{F0: 2147483647, F1: (v2 % (((v2 << ((42) & 63))) | 1)), F2: 726}
    v18 := GenS3{F0: 42, F1: 629}
    var v19 int = -9223372036854775808
    v20 := v3[uint32(v12)]
    var v21 string = ("foo bar" + ("hello" + "foo bar"))
    v19 = (v19 | 42)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
