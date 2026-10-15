package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{}
    var v3 int32 = -207
    v4 := []int{-209, -545}
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 19; fi6++ {
        if (32767 <= 1000) {
            var v7 int8 = 122
            _ = v7
        }
        delete(v1, ((("" + ((("" + "") + "a") + "")) + (("foo bar" + "foo bar") + "")) + "foo bar"))
        _ = fi6
    }
    var v8 uint32 = 42
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 float64 = (1e10 / 0.1)
        v10 += (v10 + (((0.0 + (((v10 + v10) + float64(3.14159)) - (float64(float32(1.5)) * (v10 + (v10 / float64((1e10 + 1e10))))))) - (-2.25 - (v10 + ((float64(1e10) + 0.0) / v10)))) * 1e10))
        v11 := GenS1{F0: 1}
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v8, v8)
    for fi12 := 0; fi12 < 13; fi12++ {
        v3 = -1
        v13 := GenS1{F0: 1}
        var v14 uint32 = (v8 | v8)
        _ = fi12
        _ = v13
        _ = v14
    }
    v4 = append(v4, 127)
    v15 := GenS1{F0: 771}
    v5[0] = "\n\t"
    if false {
        fmt.Println(v8)
        for fi16 := 0; fi16 < 7; fi16++ {
            v8 += 65535
            fmt.Println(v8)
            fmt.Println(v3, fi16)
            _ = fi16
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
