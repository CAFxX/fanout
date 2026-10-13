package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = -1
    v3 := []uint32{}
    var v4 uint = 18446744073709551615
    v5 := make(map[string]int)
    var v6 int32 = v1
    var v7 float32 = ((0.0 / 1e10) / -2.25)
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            if len(v3) > 0 {
                v3[0] = 1
            }
            _ = fi9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            v1 = v1
            _ = len(v5)
            v11 := GenS1{F0: "\n\t"}
            v7 = ((1.5 - float32(1e10)) + float32(0.1))
            switch "hello" {
            case "\n\t":
                fmt.Println(fi10)
            default:
                _ = len(v5)
            }
            _ = fi10
            _ = v11
        }
        _ = fi8
    }
    v2 = v2
    v6 = (v1 ^ (^(v6)))
    var v12 float32 = v7
    var v13 uint16 = 0
    if (!(128 == 886)) {
        if (((-2.25 + -2.25) - 1e10) <= 1e10) {
            var v14 uint32 = 1
            _ = v14
        }
    } else {
        fmt.Println(v6, v13, v7)
        var v15 uint16 = (v13 % ((v13) | 1))
        _ = v15
    }
    var v16 int32 = (v6 & v1)
    var v17 string = (("\n\t" + "\n\t") + "")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
