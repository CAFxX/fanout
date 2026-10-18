package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
func genFn0() int32 {
    var v6 float32 = 1.5
    var v7 float32 = v6
    v8 := GenS1{F0: 1}
    v9 := []int{517, 42}
    var v10 int16 = -580
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int) int32 {
    var v11 int64 = 1
    var v12 uint64 = 0
    v13 := []int{p0}
    fmt.Println(v11, p0)
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 uint16 = 841
    v2 := make(map[int]bool)
    var v3 int64 = 0
    var v4 int16 = 32767
    var v5 int8 = 1
    var v14 uint64 = 18446744073709551615
    v15 := []int{295, 800, 242, 128}
    v2[32767] = true
    var v16 uint = 65535
    v17 := GenS2{F0: v4}
    v4 = v4
    if (!(((3.14159 / (1e10 - -2.25)) + 3.14159) > (1e10 - 0.0))) {
        _ = len(v2)
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int8 = (v5 | v5)
            fmt.Println(fi18, v19, v3)
            _ = fi18
            _ = v19
        }
    }
    v5 = v5
    var v20 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
