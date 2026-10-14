package main

import (
    "fmt"
)


func main() {
    v1 := []int{127}
    var v2 int = 4294967295
    v3 := []uint32{2147483648, 1, 760, 0}
    var v4 int8 = 1
    var v5 uint16 = 788
    v5 = 302
    var v6 float32 = 0.1
    v2 += (v2 + v2)
    v6 = 1.5
    if (!true) {
        if true {
            var v7 int64 = -9223372036854775808
            v5 = 644
            var v8 float32 = 1e10
            _ = v7
            _ = v8
        } else {
            var v9 int16 = 451
            switch "a" {
            case "\n\t":
                var v10 float32 = 1.5
                _ = v10
            case "hello":
                _ = len(v3)
            }
            var v11 uint16 = v5
            _ = v11
            _ = v9
        }
    }
    v4 -= (v4 * int8(v5))
    var v12 uint16 = (v5 << ((997) & 63))
    _ = len(v3)
    var v13 float64 = -2.25
    v5 = 341
    var v14 uint32 = 512
    v13 = (float64(1.5) * ((3.14159 + ((0.0 / v13) - v13)) + v13))
    switch ("hello" + "\n\t") {
    case "a":
        if true {
            if true {
                for fi15 := 0; fi15 < 4; fi15++ {
                    if len(v1) > 0 {
                        v1[0] = (fi15 ^ (fi15 & fi15))
                    }
                    _ = fi15
                }
            }
        }
    default:
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v13, v14)
        default:
            v3 = append(v3, 65535)
        }
    }
    var v16 int8 = (^(v4))
    if (-2147483648 > 96) {
        fmt.Println(v4, v16)
        if len(v1) > 0 {
            v1[0] = (v2 << ((65535) & 63))
        }
        if len(v1) > 0 {
            v1[0] = (v2 << ((65535) & 63))
        }
    }
    fmt.Println(v6)
    v14 = v14
    var v17 int64 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
