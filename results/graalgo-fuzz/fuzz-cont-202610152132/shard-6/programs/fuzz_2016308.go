package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 bool, p1 bool) int32 {
    if (p0 && p1) {
        var v5 int = 0
        _ = v5
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 2}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 string = ((("foo bar" + ("\n\t" + "\n\t")) + "foo bar") + ("foo bar" + "a"))
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := []int{-443}
    v1 = v1
    v1 = ("" + "a")
    var v8 uint64 = 1
    v1 = ((v1 + (v1 + v1)) + v1)
    var v9 bool = (32767 > 1)
    fmt.Println(v9, v9, v8)
    v10 := GenS1{F0: v8}
    var v11 bool = true
    v9 = (!false)
    v8 += 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
