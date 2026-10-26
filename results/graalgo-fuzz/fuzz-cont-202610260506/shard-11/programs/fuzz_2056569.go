package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }

func main() {
    var v1 string = "foo bar"
    var v2 float64 = -2.25
    var v3 uint64 = 42
    var v4 float32 = (0.1 + 3.14159)
    v5 := make(map[string]string)
    var v6 uint32 = 502
    v7 := []int{1, 32767}
    switch -9223372036854775808 {
    case 470:
        fmt.Println(v4)
    case 1:
        for fi8 := 0; fi8 < 1; fi8++ {
            v7 = append(v7, -3)
            v9 := GenS1{F0: -1, F1: (v6 - 0), F2: -1}
            if true {
                v6 = (^(v6))
            } else {
                var v10 int8 = -108
                _ = v10
            }
            _ = fi8
            _ = v9
        }
    }
    fmt.Println(v6, v6, v6)
    for fi11 := 0; fi11 < 20; fi11++ {
        if true {
            fmt.Println(v2)
            v3 += 4294967295
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            if (!(9223372036854775807 <= 1)) {
                fmt.Println(v4, v1)
            }
            _ = fi12
        }
        _ = fi11
    }
    v7 = append(v7, -485)
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS1{F0: 100, F1: (v6 << ((4294967295) & 63)), F2: -642}
        _ = fi13
        _ = v14
    }
    var v15 int8 = 86
    var v16 uint8 = 128
    for fi17 := 0; fi17 < 18; fi17++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi17
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
