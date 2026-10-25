package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int32 }
func genFn0() int64 {
    var v4 int8 = -57
    v4 = v4
    fmt.Println(v4)
    v5 := GenS1{F0: v4}
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 string = "foo bar"
    var v2 int16 = 42
    v3 := []int{-2147483648}
    v1 = ""
    fmt.Println(v2, v1)
    if false {
        var v6 int = -756
        _ = v6
    }
    var v7 int8 = 107
    v8 := GenS2{F0: 2147483647, F1: 242}
    var v9 bool = (!(!false))
    if len(v3) > 0 {
        v3[0] = 576
    }
    fmt.Println(v7)
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(v2, fi11, v9)
            v1 = "a"
            var v12 bool = v9
            _ = fi11
            _ = v12
        }
        var v13 int32 = -2147483648
        fmt.Println(v9, v7)
        _ = fi10
        _ = v13
    }
    v14 := GenS1{F0: v7}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
