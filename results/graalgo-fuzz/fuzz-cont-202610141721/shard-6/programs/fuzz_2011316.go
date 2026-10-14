package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 18446744073709551615
    var v3 uint32 = 42
    v4 := make(map[uint32]int)
    var v5 int8 = 0
    v6 := []uint32{(v3 / (((v3 << (((^(v2))) & 63))) | 1)), v3, v3, (v3 ^ v3)}
    v5 -= -1
    var v7 int64 = 9223372036854775807
    if (!(!(((0.1 - 1e10) + 0.0) <= -2.25))) {
        if true {
            v5 -= 100
        } else {
            v3 = 65535
        }
    } else {
        var v8 string = "\n\t"
        v9 := v4[(v3 & v3)]
        _ = v8
        _ = v9
    }
    switch "hello" {
    case "":
        if false {
            v10 := v1["foo bar"]
            var v11 uint64 = 288
            _ = v10
            _ = v11
        } else {
            v3 += v3
        }
    case "\n\t":
        v5 += v5
    }
    if (-2.25 == 1.5) {
        switch -9223372036854775808 {
        case -2147483648:
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 uint64 = 0
                _ = fi12
                _ = v13
            }
        }
    } else {
        fmt.Println(v5, v2)
    }
    v1["hello"] = 128
    v14 := GenS1{F0: 0}
    v7 = (v7 / (((v7 & (v7 & (v7 % (((v7 * (v7 * (v7 ^ v7)))) | 1))))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
