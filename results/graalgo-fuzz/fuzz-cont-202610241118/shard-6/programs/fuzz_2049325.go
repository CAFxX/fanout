package main

import (
    "fmt"
)


func main() {
    var v1 string = ""
    var v2 uint16 = 234
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 uint64 = 2
    var v6 int8 = 0
    v7 := []int{}
    if (127 <= (v6 / ((v6) | 1))) {
        var v8 int32 = -2147483648
        _ = v8
    }
    delete(v3, 456)
    _ = len(v3)
    var v9 uint16 = v2
    var v10 uint8 = 127
    fmt.Println(v5, v1)
    switch 127 {
    case 985:
        var v11 int8 = v6
        _ = v11
    default:
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 14; fi13++ {
                v1 = ""
                fmt.Println(v6, v9)
                _ = fi13
            }
            var v14 uint16 = (v2 + (^(v2)))
            v5 = v5
            _ = fi12
            _ = v14
        }
    }
    v1 = (v1 + v1)
    v5 = (^(v5))
    v1 = "a"
    for fi15 := 0; fi15 < 6; fi15++ {
        v10 -= v10
        var v16 int64 = -715
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v5 += v5
        switch 1 {
        case 42:
            v4[(fi17 << ((0) & 63))] = (-(fi17))
        case 128:
            if false {
                fmt.Println(v9)
                var v18 int8 = v6
                var v19 uint64 = v5
                _ = v18
                _ = v19
            }
        }
        _ = fi17
    }
    delete(v4, 32767)
    var v20 float64 = 1e10
    _ = len(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
