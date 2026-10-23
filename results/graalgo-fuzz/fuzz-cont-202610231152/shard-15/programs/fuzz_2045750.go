package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint = 255
    v2 := make(map[string]string)
    v3 := []uint32{65535, 4294967295}
    var v4 int64 = 1
    v4 = v4
    if ((v4 - v4) == 42) {
        v5 := GenS1{F0: -553}
        _ = v5
    } else {
        switch 127 {
        case 1:
            fmt.Println(v1)
        case -1:
            v4 -= v4
        case -12:
            v4 = (v4 - v4)
        default:
            var v6 float32 = float32((0.1 / 1e10))
            _ = v6
        }
    }
    var v7 int16 = 42
    var v8 uint8 = 42
    var v9 int8 = 110
    var v10 float64 = -2.25
    var v11 uint64 = 0
    v3 = append(v3, 65535)
    var v12 int32 = -2147483648
    var v13 uint8 = 2
    for fi14 := 0; fi14 < 15; fi14++ {
        switch (fi14 >> (((v1 | v1)) & 63)) {
        case 0:
            for fi15 := 0; fi15 < 17; fi15++ {
                v10 += (v10 + v10)
                fmt.Println(v7, v13, fi14)
                v10 += (v10 + v10)
                _ = fi15
            }
        }
        _ = fi14
    }
    v16 := v2[""]
    if true {
        switch -9223372036854775808 {
        case 0:
            switch 526 {
            case -1:
                v1 -= (v1 << ((42) & 63))
            case 780:
                v3 = append(v3, 494)
            default:
                v1 = uint(v11)
            }
        default:
            var v17 int8 = (v9 ^ v9)
            _ = v17
        }
        v13 = v8
    }
    var v18 string = "foo bar"
    var v19 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
