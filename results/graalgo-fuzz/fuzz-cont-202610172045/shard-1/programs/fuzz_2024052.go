package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0(p0 uint32) string {
    var v5 int8 = 0
    var v6 int = 32768
    _ = p0
    _ = v5
    _ = v6
    return ""
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{32768}
    var v3 uint = 42
    v4 := []int{471, -1, 837, 1}
    v3 += 52
    for fi7 := 0; fi7 < 10; fi7++ {
        if false {
            fmt.Println(v3, v3)
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(fi8, v3)
            _ = fi8
        }
        var v9 int32 = 2147483647
        _ = fi7
        _ = v9
    }
    v3 = (v3 % (((v3 * v3)) | 1))
    var v10 uint = (v3 >> ((v3) & 63))
    fmt.Println(v3, v10)
    if len(v2) > 0 {
        v2[0] = 42
    }
    if (0 < 830) {
        var v11 uint32 = uint32(v3)
        _ = v11
    } else {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v12 float32 = (0.0 / 3.14159)
            _ = v12
        }
    }
    v13 := GenS1{F0: "\n\t", F1: 1}
    v14 := v1["\n\t"]
    v15 := GenS1{F0: v14, F1: 42}
    if true {
        var v16 uint64 = 154
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int16 = 42
            var v19 float64 = -2.25
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 uint = (^(v3))
        _ = v16
        _ = v20
    } else {
        v10 = v10
    }
    if (!false) {
        v14 = v14
        var v21 int8 = 100
        if false {
            var v22 uint8 = 255
            _ = v22
        }
        _ = v21
    }
    v23 := GenS1{F0: "hello", F1: -56}
    v24 := GenS1{F0: v14, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
