package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint = 65535
    var v2 float64 = 3.14159
    var v3 uint64 = 915
    var v4 bool = (v3 != (v3 & v3))
    v5 := make(map[int]int)
    v4 = (-769 < int64(v1))
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: v1}
        var v8 uint32 = 2147483648
        var v9 uint = v1
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v3 += 18446744073709551615
    fmt.Println(v3, v4)
    v10 := GenS1{F0: v1}
    v11 := []int{32768, 200}
    if (-32768 >= 1) {
        var v12 uint = (v1 & (^(v1)))
        _ = v12
    } else {
        v1 = (v1 ^ (v1 * v1))
        var v13 string = ("foo bar" + "\n\t")
        var v14 uint = 4294967295
        _ = v13
        _ = v14
    }
    var v15 uint64 = v3
    v15 += v15
    var v16 int64 = 0
    if true {
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 int8 = -1
            _ = fi17
            _ = v18
        }
        v3 = v3
    }
    v16 = v16
    fmt.Println(v2)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
