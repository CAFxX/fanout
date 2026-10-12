package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool, p1 int64) int32 {
    var v4 uint64 = 0
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    return -1
}

func main() {
    v1 := []string{"hello", "\n\t", "a", "a"}
    v2 := make(map[int]uint32)
    var v3 bool = (127 <= -52)
    v3 = (128 <= 128)
    fmt.Println(v3)
    v5 := GenS1{F0: 0, F1: 32767}
    v6 := GenS2{F0: -1, F1: 2, F2: 804}
    fmt.Println(v3)
    v3 = ((v3 || (!false)) || v3)
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 bool = v3
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 0
    v10 := GenS3{F0: 32767}
    v3 = v3
    v9 -= 127
    for fi11 := 0; fi11 < 6; fi11++ {
        v2[(fi11 + fi11)] = 2147483647
        _ = fi11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int32 = -2147483648
        _ = fi12
        _ = v13
    }
    if (genFn0((!true), 42) <= 2147483647) {
        fmt.Println(v9, v9)
    } else {
        if ((42 <= 0) || false) {
            fmt.Println(v3, v3)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            delete(v2, 4294967295)
        }
    }
    switch "foo bar" {
    case "":
        v2[32767] = 42
    default:
        var v14 float32 = float32(0.1)
        _ = v14
    }
    v3 = (true && (-32768 != 65))
    var v15 uint32 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
