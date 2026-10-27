package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = false
    var v3 bool = (false || (2 < 733))
    var v4 int16 = 42
    var v5 int64 = 1
    var v6 int = 433
    v7 := []int{(v6 % ((-123) | 1)), v6}
    var v8 uint16 = 65535
    v4 = (^(v4))
    var v9 uint8 = 1
    v9 = (^(v9))
    if false {
        switch v6 {
        case 128:
            for fi10 := 0; fi10 < 9; fi10++ {
                for fi11 := 0; fi11 < 8; fi11++ {
                    v12 := GenS1{F0: ((v4 << ((0) & 63)) < int16(v8)), F1: (v9 + v9), F2: 42}
                    v13 := GenS2{F0: 245}
                    v14 := GenS1{F0: v2, F1: (v9 * v9), F2: 42}
                    _ = len(v1)
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi10
            }
        case -842:
            for fi15 := 0; fi15 < 11; fi15++ {
                v7 = append(v7, (v6 << ((255) & 63)))
                v6 = 2147483647
                var v16 float32 = 1e10
                _ = fi15
                _ = v16
            }
        }
        var v17 uint64 = 4294967295
        _ = v17
    } else {
        var v18 float32 = 0.1
        _ = v18
    }
    fmt.Println(v6, v4)
    var v19 int64 = 362
    switch (v6 >> ((67) & 63)) {
    case 128:
        if v2 {
            fmt.Println(v8, v3)
        }
    case -9223372036854775808:
        v7 = append(v7, 4294967295)
    }
    var v20 uint32 = 91
    var v21 uint = 42
    var v22 int16 = v4
    v22 += 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
