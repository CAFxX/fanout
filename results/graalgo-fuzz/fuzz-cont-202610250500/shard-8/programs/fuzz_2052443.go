package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int8 = -1
    var v2 bool = ((v1 / ((v1) | 1)) != v1)
    v3 := make(map[string]int)
    var v4 int32 = -255
    var v5 int16 = 32767
    v6 := make(map[int]string)
    v7 := GenS1{F0: 65535, F1: v4}
    var v8 float32 = 1e10
    if (false || false) {
        switch "" {
        case "":
            switch "foo bar" {
            case "a":
                var v9 int64 = 596
                _ = v9
            default:
                v10 := GenS2{F0: 1}
                _ = v10
            }
        case "hello":
            var v11 uint64 = 20
            _ = v11
        }
    }
    switch 42 {
    case 876:
        v8 = -2.25
    case -2147483648:
        var v12 uint16 = 42
        _ = v12
    }
    var v13 string = ("" + "")
    v13 = ((((v13 + v13) + (v13 + v13)) + v13) + v13)
    v5 = v5
    v6[-360] = (v13 + "a")
    var v14 int8 = v1
    v15 := GenS1{F0: 666, F1: v4}
    v16 := []int{}
    v4 = (v4 ^ (v4 * v4))
    for fi17 := 0; fi17 < 18; fi17++ {
        v14 = (^(v14))
        _ = fi17
    }
    var v18 uint32 = 836
    switch 2147483647 {
    case 1:
        var v19 uint8 = 128
        _ = v19
    case 32767:
        v18 = 4294967295
    }
    var v20 float32 = -2.25
    v21 := GenS2{F0: int64(v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
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
