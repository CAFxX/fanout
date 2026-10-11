package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    var v3 int16 = 0
    v4 := make(map[int]uint32)
    fmt.Println(v3)
    v3 = (v3 + v3)
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{fi5, fi5, (fi5 >> ((4294967295) & 63)), 9223372036854775807}
        var v7 bool = false
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int = -1
    switch (v8 | (-(v8))) {
    case -9223372036854775808:
        var v9 uint16 = 402
        _ = v9
    case 4294967295:
        var v10 int64 = 143
        _ = v10
    default:
        var v11 float64 = 3.14159
        _ = v11
    }
    if (false && false) {
        switch ("\n\t" + (((("a" + "\n\t") + "foo bar") + "hello") + "foo bar")) {
        case "\n\t":
            if ((v8 / (((v8 / (((v8 & 42)) | 1))) | 1)) <= int(v2)) {
                v12 := GenS1{F0: (v2 | (v2 - v2))}
                v13 := GenS1{F0: v2}
                _ = v12
                _ = v13
            }
        default:
            if ((!true) || (false || true)) {
                switch 1 {
                case 2147483647:
                    v8 = 127
                }
            } else {
                v2 += v2
                v2 = 178
                v14 := []int{0}
                _ = v14
            }
        }
    }
    var v15 bool = (!false)
    v15 = (-1 > -173)
    if (false || (65535 >= 65535)) {
        var v16 int = (v8 ^ (v8 * v8))
        var v17 int16 = -32768
        fmt.Println(v16, v3)
        _ = v16
        _ = v17
    }
    fmt.Println(v2, v8, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
