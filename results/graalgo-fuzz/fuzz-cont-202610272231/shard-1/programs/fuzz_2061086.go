package main

import (
    "fmt"
)


func main() {
    var v1 float32 = 3.14159
    var v2 uint16 = 32767
    v3 := make(map[string]uint32)
    v4 := []int{803, -9223372036854775808, 32767, 2147483647}
    var v5 int16 = -32768
    var v6 uint8 = 128
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    if true {
        v1 = float32(-2.25)
        if true {
            if len(v4) > 0 {
                v4[0] = 32767
            }
        }
    } else {
        _ = len(v3)
        v3[("\n\t" + ((("foo bar" + "") + "foo bar") + ("a" + "")))] = 2
        var v8 int64 = -1
        _ = v8
    }
    v5 = v5
    switch 1 {
    case 32767:
        var v9 uint8 = (v6 << ((65535) & 63))
        _ = v9
    case 128:
        _ = len(v4)
    case 1:
        var v10 int32 = 909
        _ = v10
    default:
        _ = len(v3)
    }
    if true {
        v5 += 590
    }
    v2 += (^(v2))
    fmt.Println(v5)
    var v11 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
