package main

import (
    "fmt"
)

func genFn0(p0 int32) uint8 {
    var v4 string = ""
    v5 := []int{-9223372036854775808}
    v4 = v4
    fmt.Println(v4)
    v4 = v4
    _ = p0
    _ = v4
    _ = v5
    return 2
}

func main() {
    var v1 uint64 = 222
    v2 := []int{158, 84}
    var v3 float32 = 1.5
    v1 = v1
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 string = "foo bar"
        _ = fi6
        _ = v7
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = 213
    var v8 uint32 = 0
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v3 = v3
    v3 = (1.5 - v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
