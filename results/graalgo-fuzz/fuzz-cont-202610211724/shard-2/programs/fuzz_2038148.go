package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string) uint8 {
    var v4 int = 127
    v5 := GenS1{F0: v4}
    var v6 int16 = 0
    var v7 uint16 = 1
    var v8 uint32 = 42
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-2147483648}
    v3 := make(map[int]string)
    var v9 string = ("foo bar" + "hello")
    v10 := GenS1{F0: -469}
    v11 := GenS2{F0: 0}
    var v12 int16 = 13
    v9 = v9
    var v13 int64 = -25
    v12 = (v12 | -594)
    var v14 bool = false
    v9 = (v9 + v9)
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: (-(fi15))}
        v17 := GenS2{F0: 65535}
        if len(v2) > 0 {
            v2[0] = fi15
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
