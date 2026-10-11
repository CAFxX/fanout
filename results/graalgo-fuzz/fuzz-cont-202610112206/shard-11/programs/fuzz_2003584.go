package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int32; F2 int8 }
func genFn0(p0 uint8) int8 {
    p0 -= p0
    switch 1 {
    case 32768:
        for fi4 := 0; fi4 < 4; fi4++ {
            p0 = p0
            switch fi4 {
            case 879:
                if true {
                    p0 += 0
                } else {
                    switch "foo bar" {
                    case "":
                        p0 -= 127
                    }
                }
            default:
                var v5 int64 = 9223372036854775807
                _ = v5
            }
            _ = fi4
        }
    default:
        var v6 bool = false
        _ = v6
    }
    _ = p0
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 1
    var v3 bool = true
    v3 = v3
    v7 := GenS1{F0: 186, F1: 2147483647}
    switch 9223372036854775807 {
    case 127:
        v3 = (42 < 32767)
    case -9223372036854775808:
        fmt.Println(v3)
    }
    v3 = false
    var v8 uint8 = uint8(v2)
    v9 := []int{9223372036854775807, -63, 32768}
    v10 := GenS1{F0: 978, F1: 42}
    v8 -= 42
    v11 := GenS2{F0: v8, F1: 0, F2: -71}
    if (0.0 == 1e10) {
        v12 := GenS2{F0: (v8 - v8), F1: 2147483647, F2: 1}
        _ = v12
    }
    v2 += (v2 & (v2 << ((42) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
