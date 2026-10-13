package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint32 = 42
    v2 := make(map[int]uint32)
    var v3 uint64 = 4294967295
    var v4 int8 = 42
    for fi5 := 0; fi5 < 3; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    if true {
        var v6 uint64 = (v3 * v3)
        for fi7 := 0; fi7 < 4; fi7++ {
            switch "hello" {
            case "\n\t":
                fmt.Println(v6, v3, v6)
            case "a":
                delete(v2, fi7)
            default:
                var v8 int32 = 275
                _ = v8
            }
            _ = fi7
        }
        v6 -= (^(v6))
        _ = v6
    } else {
        v9 := GenS1{F0: 0}
        fmt.Println(v4)
        _ = len(v2)
        _ = v9
    }
    var v10 uint16 = 255
    var v11 uint32 = (v1 / (((v1 + v1)) | 1))
    fmt.Println(v3, v3)
    v12 := []int{806, 0}
    var v13 uint32 = v11
    var v14 uint64 = v3
    if (!(0 <= (v13 >> ((42) & 63)))) {
        delete(v2, 32767)
        var v15 int16 = 0
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
