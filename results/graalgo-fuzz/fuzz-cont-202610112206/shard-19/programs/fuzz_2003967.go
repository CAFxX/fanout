package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    v1 := []uint32{}
    var v2 int64 = 1
    var v3 uint = 18446744073709551615
    var v4 uint = (^(v3))
    var v5 int32 = 42
    var v6 int32 = v5
    var v7 float32 = -2.25
    v8 := GenS1{F0: (v3 - v4), F1: 2147483648}
    var v9 int8 = 47
    _ = len(v1)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint8 = 212
        _ = fi10
        _ = v11
    }
    fmt.Println(v3)
    fmt.Println(v4)
    v12 := GenS2{F0: "hello", F1: -9223372036854775808}
    for fi13 := 0; fi13 < 14; fi13++ {
        if (!(!true)) {
            fmt.Println(v3, v5, v6)
        } else {
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 bool = false
                v2 += v2
                _ = fi14
                _ = v15
            }
            var v16 int = fi13
            _ = v16
        }
        if false {
            fmt.Println(v3, v5, v2)
        } else {
            var v17 float64 = 3.14159
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 int32 = -1
                _ = fi18
                _ = v19
            }
            switch (fi13 & fi13) {
            case 32768:
                for fi20 := 0; fi20 < 18; fi20++ {
                    var v21 bool = ((^(v5)) != v5)
                    v17 = (-2.25 + float64(3.14159))
                    _ = fi20
                    _ = v21
                }
            }
            _ = v17
        }
        fmt.Println(v3, v7)
        _ = fi13
    }
    v22 := GenS2{F0: "foo bar", F1: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
