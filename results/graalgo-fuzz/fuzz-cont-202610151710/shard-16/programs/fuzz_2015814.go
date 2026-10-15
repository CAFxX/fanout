package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 152
    var v4 string = "hello"
    var v5 int64 = 42
    for fi6 := 0; fi6 < 6; fi6++ {
        v3 = v3
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: -1, F1: (32768 > (fi7 & (fi7 ^ fi6)))}
            v9 := GenS2{F0: 0, F1: 42}
            var v10 int16 = -812
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 uint16 = v3
            fmt.Println(v12, v12)
            var v13 uint = 912
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi6
    }
    v14 := v2[0]
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    fmt.Println(v5, v14)
    switch -9223372036854775808 {
    case -1:
        var v15 uint = 65535
        _ = v15
    case -2147483648:
        var v16 bool = true
        _ = v16
    }
    delete(v2, 2)
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 float64 = (float64(3.14159) * float64(((((0.1 / 1.5) - 3.14159) + float32(0.1)) / (-2.25 + 0.1))))
        v4 = v4
        _ = fi17
        _ = v18
    }
    var v19 int8 = 127
    var v20 uint8 = uint8(v3)
    _ = len(v1)
    if false {
        v2[v14] = v14
        switch "a" {
        case "foo bar":
            v21 := GenS3{F0: v5}
            _ = v21
        case "\n\t":
            fmt.Println(v5, v5, v5)
        }
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        if true {
            for fi23 := 0; fi23 < 16; fi23++ {
                v20 = v20
                v2[(v14 - uint32(v3))] = (^(v14))
                _ = fi23
            }
        } else {
            v1 = append(v1, (fi22 & (fi22 >> ((805) & 63))))
            var v24 uint32 = v14
            _ = v24
        }
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 float64 = (3.14159 + 0.1)
            var v27 int = fi22
            var v28 uint64 = 18446744073709551615
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        v4 = v4
        _ = fi22
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    var v29 uint32 = v14
    _ = len(v2)
    var v30 uint = 0
    if false {
        var v31 int64 = (v5 | 42)
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
