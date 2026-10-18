package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int32, p1 int64) int32 {
    p0 -= p0
    if ((p0 ^ -1) == p0) {
        var v7 string = "foo bar"
        var v8 int32 = (p0 >> ((885) & 63))
        _ = v7
        _ = v8
    } else {
        fmt.Println(p0, p1)
        var v9 int32 = (p0 + p0)
        v10 := []int{-1, 42, 0}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 int8 = -128
    var v2 uint16 = 1
    v3 := make(map[int]string)
    var v4 uint8 = 2
    var v5 uint16 = v2
    v6 := make(map[string]int)
    var v11 uint32 = 42
    delete(v6, ("hello" + "a"))
    fmt.Println(v5, v5, v1)
    v12 := GenS1{F0: 2}
    var v13 float64 = 1e10
    if false {
        var v14 int16 = 0
        v15 := []int{-9223372036854775808, 9223372036854775807, 1}
        _ = v14
        _ = v15
    }
    fmt.Println(v2, v4)
    if true {
        if ((^(v11)) <= (v11 << ((255) & 63))) {
            v2 += (^(v2))
        } else {
            var v16 uint8 = (v4 | v4)
            for fi17 := 0; fi17 < 4; fi17++ {
                v5 = 65535
                v2 = v5
                _ = fi17
            }
            _ = v16
        }
    }
    v18 := []int{32767, -265, -686, 9223372036854775807}
    if false {
        v6[("foo bar" + "foo bar")] = 127
    } else {
        fmt.Println(v13)
        v5 += v2
        var v19 int64 = 9223372036854775807
        _ = v19
    }
    v13 += 3.14159
    var v20 bool = (true && true)
    v5 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
