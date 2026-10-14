package main

import (
    "fmt"
)

func genFn0() int {
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int = (fi6 * fi7)
            fmt.Println(fi7)
            var v9 int16 = 32767
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 uint32 = 4294967295
        _ = fi6
        _ = v10
    }
    var v11 float64 = 0.1
    v11 = (0.1 * (-2.25 / 0.1))
    _ = v11
    return -575
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 512
    var v3 uint16 = 1
    v4 := []string{"foo bar", (("hello" + "") + "\n\t"), "", "hello"}
    var v5 uint32 = 2147483647
    fmt.Println(v3, v2, v1)
    v2 = (v2 & (v2 * (v2 >> (((^(v2))) & 63))))
    if (float32(0.0) <= (float32(1e10) * ((3.14159 + 1e10) / float32((1.5 / ((0.1 / 1.0) / -2.25)))))) {
        _ = len(v4)
    } else {
        fmt.Println(v1, v5, v2)
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        v3 -= (v3 & (v3 | v3))
        fmt.Println(v1, fi12, fi12)
        _ = fi12
    }
    var v13 uint8 = 127
    fmt.Println(v5)
    var v14 string = ""
    _ = len(v4)
    var v15 uint32 = 722
    v13 += (v13 + (v13 - v13))
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint64 = 42
        _ = fi16
        _ = v17
    }
    switch 4294967295 {
    case -248:
        v13 = 255
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
