package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint8, p1 uint16, p2 int16) int32 {
    p1 += 2
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 uint64 = 42
        var v6 int64 = -9223372036854775808
        _ = fi4
        _ = v5
        _ = v6
    }
    if (!(!true)) {
        p1 = (p1 >> ((544) & 63))
        var v7 uint32 = 0
        p2 -= (^(p2))
        _ = v7
    } else {
        if true {
            p1 = (p1 >> ((255) & 63))
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int64 = 42
    var v2 int32 = 42
    v3 := make(map[string]string)
    if true {
        delete(v3, "\n\t")
    }
    v8 := []int{127, 2147483647, 1}
    var v9 bool = false
    if (v9 || (!(-2147483648 < 490))) {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        v10 := GenS1{F0: ""}
        _ = v10
    }
    var v11 bool = false
    var v12 uint = 0
    v1 = v1
    v9 = false
    var v13 int8 = 100
    fmt.Println(v2, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
