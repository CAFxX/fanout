package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int8 = 0
    var v2 uint8 = 255
    v3 := []uint32{uint32(v2), 42, 2147483647}
    var v4 uint = 457
    if true {
        fmt.Println(v1, v4)
    }
    for fi5 := 0; fi5 < 8; fi5++ {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        _ = fi5
    }
    fmt.Println(v1, v2)
    switch -2147483648 {
    case 948:
        if true {
            v4 = 4294967295
        }
    case 127:
        var v6 int64 = -555
        _ = v6
    case 0:
        var v7 float32 = 0.1
        _ = v7
    default:
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    var v8 string = "foo bar"
    v9 := GenS1{F0: 42}
    var v10 int16 = -236
    if true {
        v11 := GenS1{F0: 42}
        v10 = (v10 ^ (v10 << (((^(v4))) & 63)))
        _ = v11
    } else {
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS2{F0: (2 >= 2)}
            _ = fi12
            _ = v13
        }
    }
    var v14 string = (v8 + "foo bar")
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            v3 = append(v3, uint32(v2))
            if (!((3.14159 / float32(3.14159)) >= 1e10)) {
                v3 = append(v3, 2)
            }
            fmt.Println(v2, fi16, v8)
            _ = fi16
        }
        v17 := GenS2{F0: false}
        var v18 uint64 = 18446744073709551615
        _ = fi15
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
