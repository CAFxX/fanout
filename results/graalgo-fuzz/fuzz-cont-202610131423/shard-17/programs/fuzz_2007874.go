package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int, p1 int16, p2 uint64) int {
    var v7 uint16 = 2
    v8 := []int{p0, p0, (^(p0)), -2147483648}
    var v9 uint = 42
    for fi10 := 0; fi10 < 9; fi10++ {
        p0 -= p0
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p0 % (((p0 - (p0 >> (((^(v9))) & 63)))) | 1))
}

func main() {
    var v1 int = 4294967295
    var v2 uint32 = 1
    var v3 uint64 = 18446744073709551615
    v4 := []string{}
    v5 := make(map[string]bool)
    var v6 int16 = 1
    v11 := GenS1{F0: 168, F1: v1, F2: 119}
    var v12 float32 = 1e10
    for fi13 := 0; fi13 < 11; fi13++ {
        v2 -= v2
        delete(v5, "hello")
        v14 := GenS1{F0: 65535, F1: (v1 % ((genFn0(v1, v6, (v3 ^ (v3 - v3)))) | 1)), F2: -79}
        _ = fi13
        _ = v14
    }
    delete(v5, "hello")
    _ = len(v4)
    v12 = 0.1
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint64 = 2
        switch "" {
        case "a":
            var v17 bool = false
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    if (!(42 < 4294967295)) {
        var v18 uint32 = v2
        v4 = append(v4, "hello")
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v6 = 538
        if false {
            fmt.Println(v2)
            for fi20 := 0; fi20 < 19; fi20++ {
                switch 0 {
                case 2147483647:
                    _ = len(v5)
                case 320:
                    v21 := v5["\n\t"]
                    _ = v21
                case 42:
                    var v22 bool = (v3 == 508)
                    var v23 uint16 = 255
                    _ = v22
                    _ = v23
                default:
                    v24 := GenS2{F0: 2, F1: 255}
                    _ = v24
                }
                _ = fi20
            }
            fmt.Println(v1, v6)
        }
        v25 := GenS3{F0: 65535}
        _ = fi19
        _ = v25
    }
    if true {
        var v26 int64 = -1
        _ = v26
    } else {
        var v27 int32 = -1
        _ = v27
    }
    var v28 int64 = 42
    switch "a" {
    case "hello":
        if false {
            v29 := GenS2{F0: 17, F1: 32767}
            v6 += 42
            var v30 int16 = 32767
            _ = v29
            _ = v30
        }
    default:
        v1 += (v1 * (-(v1)))
    }
    v12 -= v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
