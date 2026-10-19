package main

import (
    "fmt"
)

func genFn0(p0 int32) bool {
    var v7 uint8 = 127
    for fi8 := 0; fi8 < 15; fi8++ {
        p0 = -411
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint = 504
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return true
}

func main() {
    var v1 uint8 = 1
    var v2 int8 = 100
    v3 := []int{127, 32768}
    var v4 uint = 897
    v5 := []string{"hello"}
    var v6 int32 = 888
    var v11 uint8 = (v1 % ((v1) | 1))
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int8 = (v2 % (((v2 + v2)) | 1))
        _ = fi12
        _ = v13
    }
    fmt.Println(v11)
    var v14 int = 2147483647
    var v15 uint8 = 148
    v15 -= v11
    v4 = 255
    v1 += (v1 | 2)
    fmt.Println(v1)
    v11 += v15
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
