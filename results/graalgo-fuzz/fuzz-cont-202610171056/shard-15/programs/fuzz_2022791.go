package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint32, p1 int32) uint16 {
    var v7 float32 = 1e10
    fmt.Println(p0, p1)
    v8 := GenS1{F0: "hello"}
    v9 := []int{}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 117
}

func main() {
    var v1 uint8 = 255
    v2 := []string{"a", "hello", "hello", "hello"}
    var v3 uint8 = v1
    var v4 int = -2147483648
    var v5 uint32 = 2
    v6 := make(map[uint32]int)
    if false {
        switch "" {
        case "":
            v3 += v1
        case "a":
            var v10 uint16 = 467
            _ = v10
        }
        var v11 string = "hello"
        _ = v11
    }
    v4 = 766
    var v12 int32 = 1
    v13 := GenS1{F0: "a"}
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(fi14, v4, v3)
        _ = fi14
    }
    var v15 int64 = 74
    v3 = (v1 / (((^(v1))) | 1))
    for fi16 := 0; fi16 < 10; fi16++ {
        v15 -= -1
        _ = fi16
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
