package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 0
    var v2 uint32 = 2147483647
    var v3 float32 = 0.1
    v4 := make(map[string]string)
    var v5 float32 = v3
    var v6 int8 = 127
    switch 32767 {
    case 4294967295:
        if true {
            fmt.Println(v5, v3, v6)
        }
    case 128:
        var v7 string = ""
        _ = v7
    case 127:
        var v8 int64 = 9223372036854775807
        _ = v8
    }
    v5 -= (-2.25 - v3)
    var v9 float32 = float32((float64(v5) - -2.25))
    v4["a"] = "a"
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            v2 -= 2
            _ = fi11
        }
        var v12 int = 127
        var v13 uint8 = 255
        _ = fi10
        _ = v12
        _ = v13
    }
    fmt.Println(v2)
    switch 127 {
    case -1:
        fmt.Println(v2)
    default:
        fmt.Println(v6, v5, v5)
    }
    switch 9223372036854775807 {
    case -1:
        var v14 uint64 = 18446744073709551615
        _ = v14
    case -2147483648:
        var v15 int64 = 0
        _ = v15
    default:
        switch 9223372036854775807 {
        case 4294967295:
            v16 := []int{}
            _ = v16
        default:
            var v17 uint16 = 405
            _ = v17
        }
    }
    var v18 uint32 = v2
    v3 += ((((v3 + v3) - (v3 + float32(((1.5 - (float64(v5) + float64(3.14159))) / (float64(v3) + float64(v5)))))) - -2.25) + v9)
    _ = len(v4)
    var v19 string = ""
    var v20 int = 4294967295
    var v21 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
