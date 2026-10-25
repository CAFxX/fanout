package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := []string{}
    v2 := []int{2147483647}
    v3 := make(map[int]bool)
    v4 := []uint32{1, 4294967295}
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: 1}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 := GenS1{F0: -318}
    switch "\n\t" {
    case "hello":
        var v9 float32 = (0.0 / ((0.0 / 3.14159) + -2.25))
        _ = v9
    case "":
        var v10 int64 = 997
        _ = v10
    default:
        v1 = append(v1, "a")
    }
    var v11 uint16 = 255
    v12 := GenS1{F0: -1}
    var v13 uint32 = 689
    for fi14 := 0; fi14 < 12; fi14++ {
        v13 = v13
        for fi15 := 0; fi15 < 7; fi15++ {
            delete(v6, 0)
            var v16 uint32 = v13
            switch ("\n\t" + "foo bar") {
            case "a":
                v11 = (^(v11))
            case "":
                var v17 int8 = int8(v16)
                _ = v17
            case "hello":
                v18 := GenS1{F0: 2147483647}
                _ = v18
            default:
                v19 := GenS1{F0: 42}
                _ = v19
            }
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    _ = len(v2)
    v20 := GenS1{F0: -1}
    delete(v6, v13)
    switch "a" {
    case "hello":
        v13 -= (v13 << ((4294967295) & 63))
    case "a":
        var v21 int8 = 100
        _ = v21
    }
    var v22 uint = 4294967295
    delete(v3, -339)
    v11 = (v11 + (v11 % (((v11 + v11)) | 1)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6[(v13 | v13)] = (v13 - (v13 << ((v22) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
