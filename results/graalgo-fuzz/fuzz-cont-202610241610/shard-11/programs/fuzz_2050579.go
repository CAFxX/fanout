package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 bool }

func main() {
    var v1 uint8 = 255
    var v2 uint8 = v1
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 19; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 int64 = 1
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v1 -= (v2 & (v1 / (((v1 ^ v2)) | 1)))
    v7 := GenS1{F0: 376, F1: 80, F2: true}
    v1 = (v1 >> ((18446744073709551615) & 63))
    var v8 int32 = 1
    v9 := GenS1{F0: -2147483648, F1: -39, F2: (false && true)}
    var v10 uint = 42
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(v1, v8, fi11)
        var v12 int8 = 100
        _ = fi11
        _ = v12
    }
    v10 = v10
    var v13 string = "foo bar"
    var v14 uint8 = (^(v2))
    for fi15 := 0; fi15 < 17; fi15++ {
        if true {
            v16 := []int{fi15, fi15, (-(fi15))}
            v17 := GenS1{F0: (fi15 * fi15), F1: 67, F2: ((((1.5 * (1.5 / 1e10)) / 0.1) - 1.5) >= float64(1.5))}
            _ = v16
            _ = v17
        }
        fmt.Println(v13, v13)
        v18 := GenS1{F0: fi15, F1: 1, F2: false}
        _ = fi15
        _ = v18
    }
    v10 += (^(v10))
    fmt.Println(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
