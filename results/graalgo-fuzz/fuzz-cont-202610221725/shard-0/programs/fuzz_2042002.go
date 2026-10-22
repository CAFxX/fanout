package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 2
    v3 := []string{"a", ("a" + "foo bar"), "a", "\n\t"}
    v4 := []uint32{2147483648, 1, 1, 0}
    var v5 uint = v2
    var v6 int = -607
    v5 -= 0
    fmt.Println(v6, v6)
    fmt.Println(v6)
    fmt.Println(v2, v2, v6)
    v6 -= v6
    var v7 uint32 = 4294967295
    switch (("foo bar" + "hello") + "a") {
    case "hello":
        delete(v1, (v6 - (v6 << (((v5 << ((v5) & 63))) & 63))))
    case "a":
        switch (v6 % (((v6 * v6)) | 1)) {
        case 42:
            var v8 int8 = -1
            _ = v8
        default:
            _ = len(v4)
        }
    }
    switch "foo bar" {
    case "":
        v9 := GenS1{F0: false}
        _ = v9
    case "\n\t":
        var v10 float32 = -2.25
        _ = v10
    case "a":
        var v11 int32 = 774
        _ = v11
    default:
        var v12 int = v6
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v6 += (fi13 >> (((v2 << ((0) & 63))) & 63))
        delete(v1, 2147483647)
        fmt.Println(v2, v2)
        _ = fi13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        switch fi14 {
        case -1:
            var v15 int16 = int16(fi14)
            _ = v15
        case -9223372036854775808:
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 uint64 = 4294967295
                var v18 int8 = -80
                v7 += v7
                _ = fi16
                _ = v17
                _ = v18
            }
        default:
            v7 = (v7 + 2147483648)
        }
        v6 = v6
        _ = fi14
    }
    v2 += v2
    fmt.Println(v5, v7)
    fmt.Println(v6, v6, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
