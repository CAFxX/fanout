package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int; F1 uint; F2 bool }
func genFn0(p0 uint64) int8 {
    fmt.Println(p0)
    if (416 < 423) {
        var v7 uint16 = 65535
        _ = v7
    } else {
        switch "\n\t" {
        case "foo bar":
            if (!true) {
                var v8 uint16 = 2
                var v9 float64 = 3.14159
                v10 := GenS1{F0: 4294967295, F1: p0, F2: 2}
                _ = v10
                _ = v8
                _ = v9
            }
        case "a":
            fmt.Println(p0)
        }
        p0 = p0
    }
    p0 -= p0
    _ = p0
    return -1
}

func main() {
    var v1 int32 = 615
    var v2 uint = 255
    var v3 int64 = 535
    v4 := make(map[uint32]int)
    var v5 uint64 = 1
    var v6 uint = v2
    v6 += v2
    var v11 uint64 = v5
    fmt.Println(v3, v5, v11)
    fmt.Println(v1, v11, v11)
    v4[4294967295] = -940
    v11 = v5
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            v2 -= uint(v11)
            delete(v4, 733)
            var v14 uint16 = 1
            _ = fi13
            _ = v14
        }
        var v15 float32 = float32((-2.25 * 0.0))
        _ = fi12
        _ = v15
    }
    var v16 uint64 = (v11 << ((601) & 63))
    var v17 int32 = (^(v1))
    switch -9223372036854775808 {
    case 4294967295:
        var v18 int16 = 438
        _ = v18
    case 32768:
        v19 := v4[752]
        _ = v19
    case 127:
        switch -1 {
        case 42:
            v3 += -1
        }
    default:
        switch 4294967295 {
        case 420:
            switch 42 {
            case 4294967295:
                if true {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v21 := GenS2{F0: (fi20 * fi20), F1: 755, F2: false}
                        var v22 int64 = (v3 * 9223372036854775807)
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v3 -= (v3 << ((v6) & 63))
                } else {
                    var v23 uint = 1
                    _ = v23
                }
            default:
                v24 := GenS2{F0: 4294967295, F1: (v6 - 173), F2: (!((v3 ^ v3) > (v3 & v3)))}
                _ = v24
            }
        }
    }
    v25 := []int{}
    for fi26 := 0; fi26 < 15; fi26++ {
        v4[65535] = (fi26 * fi26)
        v27 := GenS1{F0: 585, F1: 571, F2: 421}
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
