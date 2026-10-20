package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 bool }
func genFn0(p0 string, p1 bool, p2 int) uint8 {
    fmt.Println(p2, p0)
    v6 := []int{709, (-(p2))}
    switch p2 {
    case 128:
        v7 := GenS1{F0: p2}
        _ = v7
    default:
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 18
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 uint32 = 646
    var v5 int8 = 47
    v8 := GenS2{F0: 1, F1: 1}
    switch 192 {
    case 127:
        if true {
            var v9 int64 = -1
            _ = v9
        } else {
            for fi10 := 0; fi10 < 13; fi10++ {
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := GenS1{F0: (fi11 ^ (fi11 - fi10))}
                    delete(v2, (^(fi10)))
                    v13 := GenS1{F0: 32768}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
            fmt.Println(v4, v1, v5)
            v14 := GenS1{F0: 4294967295}
            _ = v14
        }
    case 2147483647:
        v4 -= v4
    default:
        v1 -= -2.25
    }
    v4 -= (v4 | v4)
    delete(v3, -259)
    v15 := v2[4294967295]
    var v16 string = "foo bar"
    switch 42 {
    case -2147483648:
        v17 := GenS3{F0: -1}
        _ = v17
    case 744:
        v5 = (v5 * (v5 | v5))
    default:
        for fi18 := 0; fi18 < 17; fi18++ {
            delete(v2, fi18)
            _ = fi18
        }
    }
    v19 := GenS4{F0: -2147483648, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
