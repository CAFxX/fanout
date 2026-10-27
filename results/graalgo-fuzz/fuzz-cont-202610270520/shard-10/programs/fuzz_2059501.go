package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
func genFn0(p0 string, p1 int64) uint {
    p1 = (p1 + 9223372036854775807)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := []int{722, 0}
    var v2 int32 = 42
    var v3 int = 0
    var v4 uint8 = 233
    v5 := make(map[uint32]bool)
    var v6 int32 = v2
    for fi7 := 0; fi7 < 20; fi7++ {
        v2 = v6
        _ = fi7
    }
    fmt.Println(v6, v2)
    v8 := GenS1{F0: true}
    v4 -= 127
    for fi9 := 0; fi9 < 13; fi9++ {
        delete(v5, 77)
        _ = fi9
    }
    v5[2] = (float32(float64(1e10)) > 1.5)
    for fi10 := 0; fi10 < 12; fi10++ {
        switch -616 {
        case -282:
            fmt.Println(fi10, v4)
        case -1:
            switch 32768 {
            case 1:
                v3 += 9223372036854775807
            }
        default:
            var v11 uint16 = 1
            _ = v11
        }
        v12 := GenS2{F0: (fi10 - (^(fi10)))}
        fmt.Println(v6)
        _ = fi10
        _ = v12
    }
    _ = len(v1)
    var v13 float64 = 0.1
    switch v3 {
    case 2147483647:
        var v14 uint16 = 32767
        _ = v14
    case -17:
        v6 -= (-(v6))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
