package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint64 = 0
    var v2 int16 = -696
    v3 := []uint32{2147483648}
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    if len(v3) > 0 {
        v3[0] = 363
    }
    v6 := GenS1{F0: 255, F1: -128}
    v2 -= -179
    switch 32767 {
    case 32767:
        fmt.Println(v2)
    }
    v2 -= 42
    var v7 float32 = (0.1 - 1.5)
    var v8 bool = false
    for fi9 := 0; fi9 < 2; fi9++ {
        if (2 >= 2) {
            v10 := GenS1{F0: 0, F1: 107}
            var v11 int16 = (v2 | v2)
            _ = v10
            _ = v11
        } else {
            v12 := v4[(-(fi9))]
            if false {
                delete(v5, (fi9 >> ((42) & 63)))
                var v13 int16 = v2
                _ = v13
            }
            switch 15 {
            case -9223372036854775808:
                var v14 uint32 = 513
                _ = v14
            }
            _ = v12
        }
        v15 := GenS2{F0: 1}
        for fi16 := 0; fi16 < 3; fi16++ {
            v7 += (0.0 * (0.1 + v7))
            fmt.Println(v2, fi16, v2)
            _ = fi16
        }
        _ = fi9
        _ = v15
    }
    var v17 int16 = -564
    _ = len(v3)
    var v18 uint16 = 65535
    var v19 float32 = float32((float64(float32((0.1 * 1.5))) + 1e10))
    var v20 float64 = -2.25
    _ = len(v4)
    switch "hello" {
    case "hello":
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS2{F0: 1}
            _ = fi21
            _ = v22
        }
    case "\n\t":
        v1 = (^(v1))
    case "a":
        var v23 int8 = -1
        _ = v23
    }
    if (((v17 - v17) == (v17 << ((188) & 63))) && true) {
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 uint8 = 42
            _ = fi24
            _ = v25
        }
    } else {
        if len(v3) > 0 {
            v3[0] = 2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
