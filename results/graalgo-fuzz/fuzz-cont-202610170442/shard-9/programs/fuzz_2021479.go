package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 0
    var v2 string = ("hello" + ("foo bar" + "a"))
    var v3 int8 = 118
    v4 := make(map[string]uint32)
    var v5 int16 = v1
    var v6 float32 = 1.5
    var v7 bool = true
    fmt.Println(v7, v1, v6)
    var v8 int64 = -9223372036854775808
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 float32 = ((v6 - ((v6 / float32((-2.25 + (float64(float32(3.14159)) + (3.14159 + ((1.5 / 3.14159) - 3.14159)))))) - 1e10)) / (float32(float64(v6)) * 3.14159))
        var v11 uint64 = 4294967295
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v8 = v8
        _ = fi12
    }
    v3 += 127
    var v13 uint32 = 4294967295
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
