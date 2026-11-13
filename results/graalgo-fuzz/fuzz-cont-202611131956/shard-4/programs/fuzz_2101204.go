package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint64 = 1
    var v3 string = (("\n\t" + "\n\t") + "hello")
    var v4 int64 = -9223372036854775808
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: 1}
        var v7 int64 = v4
        _ = fi5
        _ = v6
        _ = v7
    }
    v3 = ""
    var v8 int8 = 1
    fmt.Println(v8, v8, v1)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 float32 = ((3.14159 * ((0.1 - 1.5) + (0.1 * 3.14159))) / 1.5)
        v2 += v2
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 197}
    var v12 uint32 = 65535
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
