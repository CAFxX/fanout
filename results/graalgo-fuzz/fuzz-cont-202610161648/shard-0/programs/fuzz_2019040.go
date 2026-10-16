package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint, p2 int16) uint64 {
    p1 = 1
    p0 -= -539
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    var v1 uint16 = 1
    var v2 int8 = 42
    var v3 int16 = 886
    var v4 bool = true
    var v5 int16 = v3
    v7 := []int{858, 195}
    var v8 uint8 = 42
    var v9 int32 = 2147483647
    fmt.Println(v4, v8, v3)
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 bool = (v3 <= int16(fi10))
        _ = fi10
        _ = v11
    }
    v3 -= (v5 | v3)
    var v12 uint16 = (^(v1))
    for fi13 := 0; fi13 < 5; fi13++ {
        fmt.Println(v12, v4, fi13)
        fmt.Println(v1, v1, v8)
        _ = fi13
    }
    if v4 {
        fmt.Println(v2, v12)
    } else {
        for fi14 := 0; fi14 < 10; fi14++ {
            switch (("hello" + "") + "foo bar") {
            case "foo bar":
                var v15 uint32 = 895
                _ = v15
            case "a":
                fmt.Println(v12, v2)
            case "\n\t":
                var v16 uint16 = (v1 - 0)
                _ = v16
            }
            _ = fi14
        }
        fmt.Println(v8)
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint16 = (^(v12))
        v8 -= v8
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: v8}
    var v20 bool = ((!v4) || true)
    var v21 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
