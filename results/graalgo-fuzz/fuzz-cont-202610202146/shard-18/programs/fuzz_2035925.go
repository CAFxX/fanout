package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 127
    var v3 int64 = 42
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    var v6 uint16 = 42
    v6 = v6
    v7 := []int{-246, 435}
    var v8 uint32 = 42
    var v9 float64 = float64(float32(float64((-2.25 / -2.25))))
    var v10 int32 = 42
    v11 := GenS1{F0: 128}
    if true {
        v12 := GenS1{F0: 0}
        _ = v12
    }
    switch "foo bar" {
    case "foo bar":
        for fi13 := 0; fi13 < 1; fi13++ {
            v2 = (v2 << ((65535) & 63))
            for fi14 := 0; fi14 < 13; fi14++ {
                v8 = v8
                var v15 uint64 = 1
                v10 -= v10
                _ = fi14
                _ = v15
            }
            var v16 int64 = v3
            _ = fi13
            _ = v16
        }
    case "hello":
        switch 1 {
        case 42:
            if true {
                v17 := GenS1{F0: 4294967295}
                var v18 uint = 18446744073709551615
                fmt.Println(v18)
                _ = v17
                _ = v18
            }
        case 9223372036854775807:
            fmt.Println(v3, v8)
        case -2147483648:
            v19 := GenS1{F0: -2147483648}
            _ = v19
        }
    default:
        var v20 bool = (7 <= 65535)
        _ = v20
    }
    switch -9223372036854775808 {
    case 128:
        var v21 float64 = v9
        _ = v21
    case 2147483647:
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int8 = 1
            fmt.Println(v6)
            _ = fi22
            _ = v23
        }
    case 854:
        var v24 string = ("a" + (("foo bar" + "a") + "\n\t"))
        _ = v24
    default:
        var v25 float32 = 0.0
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
