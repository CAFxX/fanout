package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int; F1 int32; F2 int16 }
func genFn0(p0 int, p1 int32, p2 int16) uint {
    var v7 uint64 = 2
    var v8 int16 = (p2 ^ p2)
    switch "foo bar" {
    case "\n\t":
        switch ("hello" + "foo bar") {
        case "":
            var v9 bool = false
            _ = v9
        default:
            v10 := []int{}
            _ = v10
        }
    case "":
        v11 := GenS1{F0: p1, F1: 127}
        _ = v11
    }
    v12 := GenS2{F0: v7, F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 878
    var v4 uint8 = 47
    v5 := []int{0, int(v3), -1, 4294967295}
    v6 := make(map[string]int)
    fmt.Println(v3)
    fmt.Println(v4, v3)
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS2{F0: 42, F1: 939}
        _ = fi13
        _ = v14
    }
    v4 = 42
    v4 += v4
    v15 := GenS2{F0: 0, F1: 2147483648}
    v3 += (v3 * 913)
    v4 += 2
    fmt.Println(v4, v4)
    for fi16 := 0; fi16 < 10; fi16++ {
        v17 := GenS3{F0: fi16, F1: 184, F2: 0}
        v4 = (v4 >> ((0) & 63))
        _ = fi16
        _ = v17
    }
    switch "hello" {
    case "":
        var v18 int16 = 963
        _ = v18
    case "hello":
        if len(v5) > 0 {
            v5[0] = 32768
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
