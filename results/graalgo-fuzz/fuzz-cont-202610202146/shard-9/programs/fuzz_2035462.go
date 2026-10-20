package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int8, p1 string, p2 int) int {
    v4 := GenS1{F0: 128}
    v5 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v6 string = (("hello" + "") + "")
    v7 := []int{genFn0(111, v6, 32767)}
    var v8 int8 = -1
    v8 += v8
    v8 = v8
    v6 = v6
    var v9 int = 32767
    if true {
        v9 += (-(v9))
    } else {
        var v10 float32 = float32((1e10 - 3.14159))
        fmt.Println(v1)
        _ = v10
    }
    v3[1] = (("foo bar" + "hello") + v6)
    var v11 int8 = (v8 ^ -128)
    var v12 float64 = 0.1
    v3[1] = ""
    for fi13 := 0; fi13 < 4; fi13++ {
        v12 += float64(0.0)
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v14 := GenS1{F0: 127}
        _ = fi13
        _ = v14
    }
    v8 = (v11 % ((127) | 1))
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 20; fi16++ {
            v17 := v3[2147483648]
            v1 = v1
            v2[v6] = (-(fi16))
            _ = fi16
            _ = v17
        }
        v1 = (v1 + v1)
        _ = fi15
    }
    var v18 int8 = -1
    v9 -= v9
    v19 := GenS1{F0: 4}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
