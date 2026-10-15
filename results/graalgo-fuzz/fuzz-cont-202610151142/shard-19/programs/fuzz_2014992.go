package main

import (
    "fmt"
)


func main() {
    var v1 int32 = 373
    var v2 int64 = 0
    v3 := []int{int(v1)}
    v4 := make(map[int]uint32)
    var v5 int32 = v1
    _ = len(v3)
    var v6 uint16 = 65535
    var v7 string = ("\n\t" + ("a" + "\n\t"))
    v3 = append(v3, -9223372036854775808)
    v8 := v4[127]
    _ = len(v4)
    v7 = "hello"
    v3 = append(v3, 9223372036854775807)
    fmt.Println(v5, v6, v1)
    var v9 string = v7
    var v10 uint16 = 32767
    for fi11 := 0; fi11 < 7; fi11++ {
        v5 += v5
        _ = fi11
    }
    fmt.Println(v6)
    var v12 uint8 = 0
    v8 -= v8
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
