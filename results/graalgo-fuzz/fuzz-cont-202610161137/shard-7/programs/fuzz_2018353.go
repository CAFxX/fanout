package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 32767
    var v2 float64 = (((float64(1.5) + float64(-2.25)) + float64(0.1)) * 3.14159)
    v3 := make(map[uint32]int)
    var v4 int8 = 14
    var v5 int64 = int64(v1)
    for fi6 := 0; fi6 < 2; fi6++ {
        v4 += (v4 << ((298) & 63))
        _ = fi6
    }
    fmt.Println(v1)
    var v7 int16 = 32767
    var v8 uint8 = 255
    var v9 uint32 = 2147483647
    v10 := []int{2147483647}
    v8 = uint8(v1)
    fmt.Println(v5, v8, v4)
    for fi11 := 0; fi11 < 8; fi11++ {
        v2 += float64(1e10)
        if (42 <= 1) {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            switch (^(fi11)) {
            case -577:
                var v12 string = ("a" + "hello")
                _ = v12
            }
        }
        _ = fi11
    }
    v1 -= (v1 / ((v1) | 1))
    for fi13 := 0; fi13 < 10; fi13++ {
        if (((v9 << ((142) & 63)) >= v9) && false) {
            v2 = (float64((1e10 * 0.1)) * v2)
        }
        v7 -= (v7 / ((v7) | 1))
        var v14 int64 = 0
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v2 = (v2 / (v2 / (0.1 * (v2 * v2))))
        _ = fi15
    }
    v8 -= (v8 | (v8 + v8))
    if len(v10) > 0 {
        v10[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
