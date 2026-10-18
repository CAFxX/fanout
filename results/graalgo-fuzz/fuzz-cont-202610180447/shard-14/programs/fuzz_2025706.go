package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 1
    var v3 bool = true
    var v4 uint8 = 127
    var v5 uint32 = 4294967295
    var v6 float64 = 3.14159
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{(^(fi7)), (fi7 & fi7), (fi7 + fi7)}
        switch "foo bar" {
        case "\n\t":
            switch "hello" {
            case "":
                var v9 uint8 = v4
                _ = v9
            default:
                var v10 uint = 502
                _ = v10
            }
        default:
            fmt.Println(v4, v5)
        }
        if false {
            fmt.Println(v5, v3)
            var v11 int16 = -32768
            v2 = (v2 << ((394) & 63))
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    var v12 uint32 = 65535
    var v13 int16 = -1
    v1[4294967295] = 9223372036854775807
    if (false || (9223372036854775807 <= -9223372036854775808)) {
        v12 = v5
    } else {
        var v14 float32 = float32(float64((float32(v6) * (3.14159 * (0.0 * 1e10)))))
        v1[2147483647] = 750
        fmt.Println(v3)
        _ = v14
    }
    v13 = (v13 % (((v13 * 0)) | 1))
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint8 = 1
        fmt.Println(v16, v13)
        v17 := GenS1{F0: 65535}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint32 = 65535
    var v19 uint64 = 18446744073709551615
    if (v19 <= (v19 ^ 1)) {
        var v20 string = "hello"
        v13 -= v13
        v6 = float64(1e10)
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v3 = v3
        fmt.Println(v18, v13, v13)
        _ = fi21
    }
    var v22 uint8 = (v4 * (v4 / ((v4) | 1)))
    v23 := GenS2{F0: 2147483647, F1: 1, F2: v12}
    v18 += uint32(v22)
    switch "a" {
    case "":
        var v24 int = 439
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
