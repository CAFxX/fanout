package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 121
    var v3 float32 = 0.0
    var v4 int = 32768
    v3 += (v3 - float32((3.14159 + 0.1)))
    if true {
        var v5 uint32 = 635
        v2 -= uint8(v5)
        _ = v5
    }
    var v6 int8 = 89
    switch -2147483648 {
    case 32768:
        v3 = v3
    }
    var v7 uint8 = v2
    var v8 int = 127
    v9 := v1[0]
    var v10 uint = 18446744073709551615
    var v11 int16 = -868
    var v12 uint32 = 2147483648
    fmt.Println(v8, v2)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int = int(v7)
        _ = fi13
        _ = v14
    }
    fmt.Println(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
