package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int32 = -1
    var v2 uint = 1
    v3 := make(map[int]int)
    var v4 float32 = ((0.1 * -2.25) - 0.1)
    var v5 bool = true
    var v6 int8 = 100
    fmt.Println(v5, v2)
    v7 := GenS1{F0: -1}
    var v8 int = 514
    var v9 bool = (false && (42 >= 167))
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint = v2
        fmt.Println(v4)
        if v5 {
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS1{F0: (v6 / ((27) | 1))}
                _ = fi12
                _ = v13
            }
            var v14 uint8 = 6
            _ = v14
        } else {
            v11 = (^(v11))
            v5 = (255 > (v11 - (v2 << (((v11 - (v2 | (^(v2))))) & 63))))
            fmt.Println(v6, v2, v6)
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v1, v5, v5)
    for fi15 := 0; fi15 < 8; fi15++ {
        switch 127 {
        case 128:
            v16 := []int{(v8 * fi15), fi15}
            _ = v16
        default:
            if (!v5) {
                fmt.Println(v6, v9)
                v17 := GenS1{F0: v6}
                var v18 int16 = 0
                _ = v17
                _ = v18
            } else {
                fmt.Println(v6)
            }
        }
        _ = len(v3)
        v19 := GenS1{F0: v6}
        _ = fi15
        _ = v19
    }
    var v20 bool = v9
    var v21 int16 = int16(v2)
    v5 = (!((!false) || (((-552 <= (v21 >> ((v2) & 63))) || (!(42 != 0))) || (int64(v6) < 613))))
    if v20 {
        var v22 uint64 = 18446744073709551615
        _ = v22
    } else {
        v23 := GenS1{F0: (^(v6))}
        v20 = ((v1 << (((v2 >> ((65535) & 63))) & 63)) == v1)
        _ = v23
    }
    delete(v3, 32768)
    var v24 int = -9223372036854775808
    v25 := GenS1{F0: 92}
    v26 := []int{v8, (v24 ^ v24)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
