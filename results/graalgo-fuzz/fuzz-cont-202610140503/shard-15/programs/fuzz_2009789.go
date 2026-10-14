package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []int{0}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 255
    var v4 uint8 = v3
    v5 := GenS1{F0: -1}
    v1 = append(v1, 128)
    if true {
        if (18446744073709551615 < 4294967295) {
            var v6 uint8 = 2
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        if false {
            fmt.Println(v4, v4)
            v8 := GenS1{F0: 1}
            if true {
                v4 += (^(v4))
            } else {
                var v9 int = (fi7 | (-(fi7)))
                _ = v9
            }
            _ = v8
        }
        v10 := v2[2147483648]
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        fmt.Println(v3, v3)
        v4 = 83
        _ = fi11
    }
    var v12 float32 = float32(0.1)
    fmt.Println(v12, v12, v4)
    if true {
        v2[42] = 0
        var v13 int = 9223372036854775807
        if (!true) {
            if (((1.5 / float32((3.14159 / 3.14159))) - v12) > v12) {
                v3 -= (v4 / (((v3 ^ v4)) | 1))
                v3 = (v3 * 128)
            }
        }
        _ = v13
    }
    v14 := GenS1{F0: -368}
    fmt.Println(v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
