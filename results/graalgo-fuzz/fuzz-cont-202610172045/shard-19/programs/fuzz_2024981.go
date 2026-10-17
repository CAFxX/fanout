package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{74, 42, 1}
    var v2 uint32 = 650
    var v3 uint32 = v2
    var v4 string = "foo bar"
    v1 = append(v1, (v2 ^ (v2 << ((2) & 63))))
    v1 = append(v1, v2)
    v1 = append(v1, 0)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v5 uint16 = 42
    switch (v4 + (v4 + v4)) {
    case "foo bar":
        v2 = 2147483647
    case "hello":
        var v6 float64 = (0.1 + 0.1)
        _ = v6
    case "\n\t":
        for fi7 := 0; fi7 < 8; fi7++ {
            if (!true) {
                for fi8 := 0; fi8 < 16; fi8++ {
                    v1 = append(v1, v2)
                    var v9 float32 = 0.0
                    if (-128 > -96) {
                        v2 -= v2
                        _ = len(v1)
                    } else {
                        v3 -= v3
                        fmt.Println(v4, v2)
                    }
                    _ = fi8
                    _ = v9
                }
                fmt.Println(v5, v2, v2)
                v3 -= 2147483648
            }
            switch (fi7 & 32768) {
            case 42:
                var v10 bool = (1 < 1)
                _ = v10
            case 0:
                var v11 uint32 = v3
                _ = v11
            }
            for fi12 := 0; fi12 < 17; fi12++ {
                v3 += v3
                v5 = v5
                v2 += v3
                _ = fi12
            }
            _ = fi7
        }
    }
    var v13 int64 = 1
    if len(v1) > 0 {
        v1[0] = (v2 >> ((0) & 63))
    }
    switch -1 {
    case 127:
        v13 = v13
    default:
        v5 += (v5 / ((32767) | 1))
    }
    if true {
        v5 += v5
    } else {
        v4 = ((v4 + v4) + (v4 + v4))
        v2 -= (^(v2))
        var v14 bool = true
        _ = v14
    }
    fmt.Println(v2, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
