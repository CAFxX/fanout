package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int8, p1 uint64) uint8 {
    p1 += (p1 ^ (p1 - 4294967295))
    var v7 uint8 = 127
    _ = p0
    _ = p1
    _ = v7
    return v7
}

func main() {
    var v1 uint64 = 779
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    v4 := []string{"foo bar", ("hello" + "foo bar")}
    v5 := make(map[int]uint32)
    var v6 int64 = 393
    if (!true) {
        _ = len(v3)
    }
    var v8 int64 = 1
    var v9 uint8 = 42
    var v10 bool = true
    fmt.Println(v8, v8, v9)
    switch -1 {
    case -9223372036854775808:
        v11 := GenS1{F0: 255}
        _ = v11
    case 4294967295:
        v12 := GenS1{F0: 15}
        _ = v12
    }
    v13 := GenS1{F0: 1}
    delete(v3, "\n\t")
    if (v10 || (v10 || ((!v10) && (65535 < 879)))) {
        var v14 uint64 = uint64(v9)
        v15 := v3["hello"]
        _ = v14
        _ = v15
    } else {
        var v16 uint = 42
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v4 = append(v4, "hello")
        _ = fi17
    }
    _ = len(v2)
    fmt.Println(v9, v10, v9)
    v18 := GenS1{F0: 649}
    v19 := GenS1{F0: 897}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
