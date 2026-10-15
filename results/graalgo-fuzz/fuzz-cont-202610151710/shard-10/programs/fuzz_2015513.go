package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16, p1 int) uint {
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int32 = 931
        v7 := []int{p1, (fi5 ^ 561), p1}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p1, p1)
    v8 := GenS1{F0: 234}
    _ = p0
    _ = p1
    _ = v8
    return 4294967295
}

func main() {
    v1 := []int{374, -307, -1}
    v2 := make(map[int]string)
    var v3 float64 = (3.14159 / 1.0)
    v4 := make(map[int]string)
    switch -507 {
    case 1:
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(fi9, fi9)
            var v10 uint16 = 65535
            _ = fi9
            _ = v10
        }
    case 4294967295:
        switch -708 {
        case 15:
            fmt.Println(v3)
        case 110:
            if false {
                switch 447 {
                case -133:
                    fmt.Println(v3)
                case 1:
                    v3 += 1.5
                default:
                    var v11 float32 = (3.14159 * 0.1)
                    _ = v11
                }
                v12 := GenS2{F0: -479}
                _ = v12
            }
        default:
            v3 -= ((v3 + float64(0.1)) * v3)
        }
    case 490:
        fmt.Println(v3)
    default:
        delete(v2, 127)
    }
    v13 := GenS2{F0: 42}
    v3 = (v3 + v3)
    var v14 uint32 = 42
    v14 += 2
    var v15 uint64 = 4294967295
    var v16 int = 32767
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v3 = (v3 / (v3 / (1.5 - v3)))
    switch (^(v16)) {
    case 32768:
        v3 = 0.1
    case -230:
        var v17 int32 = 624
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
