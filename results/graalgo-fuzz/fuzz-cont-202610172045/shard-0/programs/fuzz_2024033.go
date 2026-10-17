package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint32; F2 uint64 }

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]bool)
    var v3 int16 = -740
    v4 := make(map[string]bool)
    v5 := []uint32{1, 2147483648}
    v6 := GenS1{F0: -1, F1: 32767}
    v1 = v1
    if (false || false) {
        switch "a" {
        case "foo bar":
            var v7 bool = true
            _ = v7
        case "\n\t":
            v8 := GenS1{F0: -39, F1: 32767}
            _ = v8
        }
    }
    _ = len(v4)
    for fi9 := 0; fi9 < 19; fi9++ {
        delete(v4, "a")
        _ = fi9
    }
    if false {
        v10 := GenS1{F0: 1, F1: (v3 * v3)}
        var v11 float32 = (1.5 + 1e10)
        _ = v10
        _ = v11
    } else {
        v12 := v2["foo bar"]
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS2{F0: 65535}
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    var v15 string = "foo bar"
    if true {
        var v16 uint = uint(v1)
        v17 := GenS2{F0: (v16 >> ((v16) & 63))}
        v18 := GenS2{F0: 18446744073709551615}
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 42, F1: 1}
    v20 := GenS3{F0: 4294967295, F1: 1, F2: 4294967295}
    fmt.Println(v15, v1)
    v21 := GenS2{F0: 255}
    var v22 int8 = -104
    v22 -= v22
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
