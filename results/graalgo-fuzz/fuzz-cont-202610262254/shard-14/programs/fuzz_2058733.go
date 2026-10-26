package main

import (
    "fmt"
)


func main() {
    var v1 string = "a"
    var v2 uint = 4294967295
    v3 := make(map[int]uint32)
    var v4 uint = v2
    var v5 uint8 = 1
    v4 -= v2
    v2 = v4
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 uint32 = 65535
            delete(v3, (fi7 | fi7))
            _ = fi7
            _ = v8
        }
        fmt.Println(v5, v2, v5)
        v2 = (v2 & (v4 % ((uint(v5)) | 1)))
        _ = fi6
    }
    v5 -= (^(v5))
    v1 = ("a" + (v1 + ("\n\t" + v1)))
    v4 -= v4
    switch -460 {
    case 127:
        switch 42 {
        case 1:
            v5 += v5
        default:
            var v9 int8 = 100
            _ = v9
        }
    case 128:
        fmt.Println(v5, v2, v1)
    default:
        var v10 uint = 18446744073709551615
        _ = v10
    }
    var v11 string = (v1 + (v1 + v1))
    v5 -= 128
    switch 42 {
    case -9223372036854775808:
        v11 = v11
    default:
        v12 := v3[-9223372036854775808]
        _ = v12
    }
    v5 += v5
    if false {
        fmt.Println(v1, v4, v4)
        for fi13 := 0; fi13 < 8; fi13++ {
            v5 = v5
            _ = fi13
        }
        fmt.Println(v2, v5, v1)
    }
    var v14 int16 = 585
    if (false && false) {
        fmt.Println(v14, v4, v2)
    }
    var v15 uint8 = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
