package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 bool) string {
    for fi5 := 0; fi5 < 5; fi5++ {
        p0 = p0
        _ = fi5
    }
    var v6 int = -478
    _ = p0
    _ = v6
    return (("foo bar" + "") + "hello")
}
func genFn1(p0 uint8) int32 {
    v7 := []int{-878, -1, 0}
    v8 := GenS1{F0: 0}
    _ = p0
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    var v1 float64 = 1e10
    var v2 int = -772
    v3 := []string{"hello", ""}
    var v4 int8 = 42
    for fi9 := 0; fi9 < 10; fi9++ {
        _ = len(v3)
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(v2, fi10, fi10)
            _ = fi10
        }
        _ = fi9
    }
    v11 := GenS1{F0: -1}
    if false {
        var v12 uint8 = 170
        _ = v12
    }
    v2 += (v2 >> ((42) & 63))
    for fi13 := 0; fi13 < 19; fi13++ {
        v3 = append(v3, "foo bar")
        fmt.Println(fi13, fi13, v4)
        _ = fi13
    }
    v4 -= v4
    for fi14 := 0; fi14 < 5; fi14++ {
        if len(v3) > 0 {
            v3[0] = "a"
        }
        _ = fi14
    }
    v1 -= 0.0
    var v15 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
