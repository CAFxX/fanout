package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 196
    var v2 uint = 0
    var v3 string = "hello"
    v4 := make(map[int]uint32)
    var v5 uint8 = v1
    for fi6 := 0; fi6 < 15; fi6++ {
        _ = len(v4)
        _ = fi6
    }
    v5 = (v1 << ((v2) & 63))
    v5 = 128
    var v7 uint16 = 0
    var v8 int = 0
    var v9 float32 = ((0.1 - float32((3.14159 / float64(0.1)))) / (float32(((-2.25 - 3.14159) * 0.1)) / 3.14159))
    var v10 int16 = -385
    v10 -= v10
    var v11 int64 = 9223372036854775807
    fmt.Println(v10)
    var v12 int8 = 61
    if (!(42 < 65535)) {
        v13 := []int{v8, (v8 * -1), (v8 + v8)}
        _ = v13
    }
    if (((2 <= 0) && false) || (42 < (v10 >> ((v2) & 63)))) {
        var v14 uint8 = v1
        fmt.Println(v12, v7, v5)
        fmt.Println(v1)
        _ = v14
    }
    var v15 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
