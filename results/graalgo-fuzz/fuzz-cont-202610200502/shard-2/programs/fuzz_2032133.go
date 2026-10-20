package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint = 18446744073709551615
    var v2 string = "a"
    var v3 int = 9223372036854775807
    v4 := make(map[uint32]uint32)
    v5 := []int{(v3 >> (((v1 >> ((1) & 63))) & 63)), 2147483647, v3}
    var v6 uint = (^(v1))
    v5 = append(v5, v3)
    if false {
        v1 = (v6 | v1)
        v7 := v4[65535]
        fmt.Println(v6, v6, v1)
        _ = v7
    } else {
        v2 = ""
    }
    var v8 float64 = 3.14159
    fmt.Println(v3, v2)
    v9 := GenS1{F0: ("a" + "foo bar")}
    var v10 int8 = -24
    switch v3 {
    case 9223372036854775807:
        for fi11 := 0; fi11 < 20; fi11++ {
            if false {
                v2 = (((v2 + (v2 + (v2 + ((v2 + ("\n\t" + "hello")) + "hello")))) + v2) + v2)
            }
            _ = fi11
        }
    default:
        v6 += (v1 % (((v6 - v6)) | 1))
    }
    var v12 int16 = 42
    v1 = v6
    fmt.Println(v2, v8)
    var v13 uint16 = 1
    var v14 int64 = 0
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint64 = 18446744073709551615
        var v17 int64 = (v14 | -452)
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
