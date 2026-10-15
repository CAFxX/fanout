package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{128}
    var v3 uint64 = 42
    v4 := []uint32{1, 2}
    var v5 float64 = 3.14159
    var v6 int32 = 590
    if ((!(true || false)) && true) {
        v6 += int32(v3)
        var v7 uint64 = (v3 % (((v3 - v3)) | 1))
        _ = v7
    } else {
        var v8 int32 = v6
        _ = v8
    }
    switch "foo bar" {
    case "":
        var v9 int16 = 1
        _ = v9
    case "foo bar":
        v5 -= float64(1.5)
    case "a":
        _ = len(v4)
    default:
        if false {
            v5 += (v5 / v5)
        }
    }
    if false {
        var v10 bool = false
        v3 += (v3 % ((v3) | 1))
        if v10 {
            v5 = v5
            var v11 uint64 = 855
            var v12 float64 = (1e10 * v5)
            _ = v11
            _ = v12
        } else {
            fmt.Println(v5, v6)
            v6 -= (v6 % (((v6 ^ -1)) | 1))
            _ = len(v2)
        }
        _ = v10
    }
    var v13 int = 42
    if (!true) {
        var v14 string = ""
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v5 += v5
        _ = fi15
    }
    v6 = (v6 << ((307) & 63))
    var v16 int32 = v6
    v17 := GenS1{F0: v16}
    v6 = v6
    v16 = (v16 % ((v16) | 1))
    v18 := v1[""]
    v3 += (v3 ^ 34)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
