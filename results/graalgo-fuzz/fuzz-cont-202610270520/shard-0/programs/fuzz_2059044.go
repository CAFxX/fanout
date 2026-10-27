package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 1e10
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    var v5 int16 = -1
    for fi6 := 0; fi6 < 2; fi6++ {
        switch -2147483648 {
        case 4294967295:
            switch fi6 {
            case 1:
                v2 += (0.0 * (3.14159 - float64(3.14159)))
            }
        default:
            var v7 float32 = float32(0.1)
            _ = v7
        }
        var v8 uint16 = 42
        switch "" {
        case "\n\t":
            var v9 int32 = 2147483647
            _ = v9
        case "hello":
            v10 := GenS1{F0: -1}
            _ = v10
        case "a":
            v11 := GenS1{F0: v5}
            _ = v11
        }
        _ = fi6
        _ = v8
    }
    v12 := []int{128}
    var v13 int32 = -2147483648
    delete(v3, "")
    fmt.Println(v5)
    if false {
        var v14 uint32 = 4294967295
        _ = len(v12)
        _ = v14
    }
    v13 = v13
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 int8 = 42
        v13 -= v13
        _ = fi15
        _ = v16
    }
    v5 = v5
    for fi17 := 0; fi17 < 16; fi17++ {
        v13 = (v13 ^ v13)
        var v18 int64 = 42
        _ = fi17
        _ = v18
    }
    v1[-651] = true
    v19 := GenS1{F0: (v5 % (((v5 * v5)) | 1))}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    fmt.Println(v5, v2)
    v2 += v2
    v13 = v13
    var v20 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
