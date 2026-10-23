package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint16, p2 uint32) bool {
    p1 = 808
    v4 := GenS1{F0: 42}
    v5 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{""}
    var v3 uint32 = 788
    fmt.Println(v3)
    v3 -= v3
    var v6 uint8 = 154
    var v7 uint64 = 536
    v8 := GenS1{F0: -1}
    v9 := GenS1{F0: 2147483647}
    if (!true) {
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS2{F0: 1}
            fmt.Println(fi10, v3)
            v12 := GenS3{F0: v7}
            _ = fi10
            _ = v11
            _ = v12
        }
        v3 = v3
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            switch fi13 {
            case 2147483647:
                v7 -= (v7 | (v7 & (^(v7))))
            case -752:
                v6 = v6
            case -9223372036854775808:
                fmt.Println(v3, v7)
            }
            v14 := GenS2{F0: int64(fi13)}
            var v15 uint16 = 31
            _ = fi13
            _ = v14
            _ = v15
        }
        switch "" {
        case "":
            switch 2147483647 {
            case 127:
                var v16 uint = 0
                _ = v16
            case -1:
                v7 = v7
            }
        }
    }
    fmt.Println(v7)
    v6 -= v6
    if (0 < 87) {
        var v17 uint64 = v7
        fmt.Println(v17, v17)
        var v18 float32 = 0.0
        _ = v17
        _ = v18
    }
    if true {
        switch -714 {
        case 32767:
            v6 = 2
        default:
            var v19 uint64 = 858
            _ = v19
        }
        var v20 int32 = 1
        _ = v20
    }
    var v21 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
