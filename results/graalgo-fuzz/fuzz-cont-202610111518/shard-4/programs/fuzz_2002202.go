package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []string{""}
    var v2 bool = (168 != 1)
    v3 := []string{}
    v4 := make(map[string]uint32)
    var v5 int32 = -1
    fmt.Println(v2)
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 float32 = 0.1
        v5 += v5
        switch "a" {
        case "\n\t":
            if (!false) {
                var v8 int = fi6
                fmt.Println(v7, fi6, v7)
                var v9 int64 = 499
                _ = v8
                _ = v9
            }
        default:
            var v10 int = (fi6 - fi6)
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v5, v5)
    switch "" {
    case "":
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 int64 = -9223372036854775808
            _ = fi11
            _ = v12
        }
    case "hello":
        fmt.Println(v2)
    default:
        var v13 int32 = (v5 ^ v5)
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        delete(v4, "")
        var v15 int64 = 9223372036854775807
        _ = fi14
        _ = v15
    }
    if false {
        var v16 int64 = 1
        _ = v16
    } else {
        v17 := GenS1{F0: 524}
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int8 = 100
        fmt.Println(fi18, v2, fi18)
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        v5 = 718
        fmt.Println(v2, v2, fi20)
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := GenS1{F0: 2}
            var v23 int8 = 1
            v3 = append(v3, "a")
            var v24 uint16 = 185
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi20
    }
    v25 := GenS1{F0: 0}
    fmt.Println(v5, v2)
    switch 2147483647 {
    case -9223372036854775808:
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 float32 = 1e10
            v28 := GenS1{F0: 1}
            v2 = v2
            _ = fi26
            _ = v27
            _ = v28
        }
    case -815:
        v29 := GenS1{F0: 2}
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
