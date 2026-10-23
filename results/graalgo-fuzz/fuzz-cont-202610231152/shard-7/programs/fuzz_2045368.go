package main

import (
    "fmt"
)


func main() {
    v1 := []int{-293, 122, -2147483648, 1}
    v2 := make(map[string]uint32)
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]int)
    var v5 uint32 = 2
    if true {
        v1 = append(v1, 128)
        fmt.Println(v5)
    }
    switch 53 {
    case -1:
        fmt.Println(v3, v5)
    case 128:
        var v6 float32 = 1e10
        _ = v6
    default:
        fmt.Println(v3, v3)
    }
    _ = len(v2)
    v3 -= v3
    v5 -= v5
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v5, v3)
        if (((((32767 > 65535) || false) && true) && false) || false) {
            v5 = (v5 / (((v5 - v5)) | 1))
            fmt.Println(v3, fi7)
            var v8 int16 = -898
            _ = v8
        } else {
            if (-1 >= -425) {
                v5 -= v5
                if (!(!true)) {
                    for fi9 := 0; fi9 < 5; fi9++ {
                        var v10 uint8 = 126
                        _ = fi9
                        _ = v10
                    }
                }
            }
            if false {
                var v11 uint64 = 2
                v5 = (v5 + 4294967295)
                _ = v11
            }
        }
        _ = fi7
    }
    fmt.Println(v3, v5)
    fmt.Println(v5, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
