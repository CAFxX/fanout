package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 uint32, p1 uint32) uint16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        p0 = p0
        if true {
            v6 := []int{fi5}
            fmt.Println(p0)
            v7 := GenS1{F0: 1, F1: 1}
            _ = v6
            _ = v7
        }
        fmt.Println(fi5)
        _ = fi5
    }
    var v8 int32 = 0
    if false {
        fmt.Println(v8)
        v9 := []int{1, -2147483648}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v8
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    v3 := []int{128, 4294967295, 9223372036854775807, 127}
    var v4 int32 = 0
    var v10 string = (("\n\t" + "foo bar") + (("a" + (("\n\t" + "hello") + "\n\t")) + "\n\t"))
    var v11 int64 = -315
    if false {
        v12 := GenS1{F0: 423, F1: -1}
        var v13 float32 = (1.5 / 3.14159)
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 0, F1: v4}
        _ = v14
    }
    fmt.Println(v10)
    v10 = (v10 + "")
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v11)
            _ = fi16
        }
        if len(v3) > 0 {
            v3[0] = fi15
        }
        _ = fi15
    }
    fmt.Println(v10)
    v17 := GenS1{F0: 476, F1: v4}
    v10 = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
