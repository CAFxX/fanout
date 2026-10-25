package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 int16 }
func genFn0() uint {
    v5 := GenS1{F0: -2147483648, F1: 1, F2: 2}
    var v6 int64 = 0
    var v7 int = 0
    fmt.Println(v6)
    _ = v5
    _ = v6
    _ = v7
    return 112
}

func main() {
    var v1 bool = (false && false)
    var v2 int = 32767
    v3 := make(map[string]bool)
    var v4 string = "hello"
    var v8 uint32 = 2
    for fi9 := 0; fi9 < 14; fi9++ {
        if (9223372036854775807 > 0) {
            fmt.Println(v4, v2, v4)
            v10 := GenS2{F0: -686}
            var v11 int64 = 0
            _ = v10
            _ = v11
        } else {
            v8 -= (v8 * (v8 % ((750) | 1)))
        }
        switch (fi9 - v2) {
        case -1:
            v2 -= fi9
        }
        _ = fi9
    }
    v8 = v8
    switch -482 {
    case -1:
        var v12 int16 = -405
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        if v1 {
            var v14 float32 = 1.5
            _ = v14
        }
        switch v4 {
        case "a":
            fmt.Println(v8)
        case "":
            v15 := GenS1{F0: 42, F1: 1, F2: 128}
            _ = v15
        }
        _ = fi13
    }
    v3[v4] = v1
    v16 := []int{}
    if v1 {
        _ = len(v16)
    } else {
        if v1 {
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 int32 = 1
                _ = fi17
                _ = v18
            }
            v16 = append(v16, 9223372036854775807)
        } else {
            _ = len(v16)
        }
        fmt.Println(v4, v8, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
