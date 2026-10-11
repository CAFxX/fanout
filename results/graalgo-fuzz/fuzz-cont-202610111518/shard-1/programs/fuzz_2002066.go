package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
func genFn0(p0 string) int16 {
    v5 := GenS1{F0: 42}
    v6 := []int{32768, 42}
    var v7 bool = (false || ((0 > -315) || true))
    v8 := GenS2{F0: -433}
    for fi9 := 0; fi9 < 20; fi9++ {
        v7 = v7
        _ = len(v6)
        _ = fi9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint32 = 165
    v2 := make(map[string]uint32)
    var v3 bool = false
    var v4 int64 = -9223372036854775808
    var v10 int = 743
    delete(v2, "")
    v1 = (^(v1))
    v11 := GenS3{F0: v10}
    v12 := []int{(v10 >> ((4294967295) & 63))}
    var v13 uint = 2
    v10 = (v10 % (((-(v10))) | 1))
    fmt.Println(v4, v10, v10)
    v14 := GenS4{F0: -32768}
    v13 -= 255
    fmt.Println(v3)
    var v15 int32 = 0
    v16 := GenS4{F0: genFn0("\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
