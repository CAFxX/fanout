package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int, p1 int64, p2 int32) uint32 {
    v4 := GenS1{F0: 42}
    p2 += p2
    v5 := []int{p0, p0}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]string)
    v3 := []int{int(v1), 42, -9223372036854775808}
    if false {
        _ = len(v2)
        var v6 float64 = (1e10 + float64(0.0))
        v2[9223372036854775807] = "\n\t"
        _ = v6
    }
    _ = len(v2)
    v2[50] = "hello"
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            v3 = append(v3, (fi8 | (fi7 ^ fi7)))
            fmt.Println(fi7, fi8, fi8)
            v1 = (^(v1))
            _ = fi8
        }
        var v9 int8 = 1
        _ = fi7
        _ = v9
    }
    var v10 float32 = 1e10
    v1 += (v1 & 237)
    if (128 != 1) {
        _ = len(v2)
    } else {
        var v11 string = (("foo bar" + "") + ("hello" + "\n\t"))
        v1 += (v1 << ((0) & 63))
        v1 = v1
        _ = v11
    }
    delete(v2, -557)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
