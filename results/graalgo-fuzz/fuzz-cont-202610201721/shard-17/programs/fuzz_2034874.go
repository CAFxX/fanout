package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8) int64 {
    v7 := GenS1{F0: 1, F1: 42, F2: 255}
    var v8 int = 9223372036854775807
    _ = p0
    _ = v7
    _ = v8
    return 404
}

func main() {
    v1 := []string{}
    v2 := []uint32{42, 0, 65535, 42}
    var v3 string = "\n\t"
    var v4 uint64 = 18446744073709551615
    v5 := []uint32{2147483647, 4294967295}
    v6 := make(map[string]uint32)
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v3)
    if false {
        var v9 int8 = -122
        _ = v9
    } else {
        var v10 int32 = 810
        _ = v10
    }
    var v11 uint = 2
    var v12 int16 = 0
    for fi13 := 0; fi13 < 4; fi13++ {
        if false {
            v3 = (v3 + v3)
            v14 := GenS2{F0: -1}
            _ = v14
        } else {
            switch "foo bar" {
            case "":
                v3 = ("" + "")
            default:
                fmt.Println(v12)
            }
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS1{F0: 2, F1: v11, F2: (v11 | v11)}
            var v17 uint64 = 4294967295
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    if (0 <= 2147483647) {
        delete(v6, v3)
    }
    v18 := GenS2{F0: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
