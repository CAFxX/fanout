package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint64) int8 {
    p0 = 728
    var v6 string = "a"
    var v7 int = -9223372036854775808
    v7 = v7
    if (v7 != 128) {
        v8 := []int{v7}
        var v9 uint64 = (^(p0))
        v10 := GenS1{F0: 1}
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := []int{-9223372036854775808, 42, -195}
    var v2 int32 = 813
    var v3 int16 = 1
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    v2 += int32(v3)
    var v11 uint64 = 42
    var v12 int16 = 32767
    if ((((((float32((float64(((float32(3.14159) / 3.14159) + (3.14159 * 0.0))) - 0.0)) - (1e10 / 1.0)) + float32(0.0)) / 3.14159) / 1e10) - 0.1) > 3.14159) {
        var v13 int32 = v2
        var v14 int8 = 111
        var v15 uint32 = 2
        _ = v13
        _ = v14
        _ = v15
    }
    if len(v1) > 0 {
        v1[0] = -310
    }
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            if (-339 > (fi16 | (fi17 / ((fi17) | 1)))) {
                var v18 int = fi17
                fmt.Println(fi17)
                var v19 int32 = 42
                var v20 float32 = (3.14159 + (1.5 - ((float32((float64(0.0) / 1.0)) * ((-2.25 + 3.14159) + 1e10)) - (0.0 * 1.5))))
                _ = v18
                _ = v19
                _ = v20
            } else {
                v3 = 750
            }
            v4["a"] = ("foo bar" + (("hello" + ("" + ("" + "foo bar"))) + "a"))
            v2 += 42
            fmt.Println(fi17)
            _ = fi17
        }
        _ = fi16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
