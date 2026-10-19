package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int16 = 1000
    var v2 float64 = 0.0
    var v3 float32 = -2.25
    var v4 int16 = v1
    var v5 bool = (735 > 2147483647)
    var v6 int64 = 1
    var v7 uint8 = 115
    var v8 float64 = 1.5
    switch -719 {
    case 829:
        v9 := GenS1{F0: 0, F1: ("hello" + "foo bar")}
        _ = v9
    case -2147483648:
        fmt.Println(v8, v4, v3)
    case 0:
        if v5 {
            v10 := []int{783, 623, 802, 1}
            v11 := GenS2{F0: 754, F1: uint32(v7)}
            _ = v10
            _ = v11
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 20; fi13++ {
                v1 = (v4 / (((v4 - 0)) | 1))
                v14 := []int{}
                _ = fi13
                _ = v14
            }
            var v15 int32 = int32(fi12)
            v15 = v15
            _ = fi12
            _ = v15
        }
    case "":
        fmt.Println(v3, v7)
    default:
        for fi16 := 0; fi16 < 8; fi16++ {
            v6 += 0
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 float64 = v8
                var v19 string = "hello"
                _ = fi17
                _ = v18
                _ = v19
            }
            fmt.Println(v4, v4, v4)
            _ = fi16
        }
    }
    v20 := GenS3{F0: -92}
    if true {
        fmt.Println(v6, v6)
    }
    fmt.Println(v2)
    var v21 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
