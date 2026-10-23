package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16, p1 int8, p2 int) int32 {
    for fi5 := 0; fi5 < 9; fi5++ {
        p2 += (fi5 | (fi5 / ((fi5) | 1)))
        _ = fi5
    }
    fmt.Println(p2, p2)
    v6 := []int{}
    if false {
        switch "foo bar" {
        case "hello":
            v7 := GenS1{F0: 568}
            _ = v7
        case "\n\t":
            var v8 uint16 = 2
            _ = v8
        default:
            fmt.Println(p0)
        }
        var v9 int16 = (p0 % ((-32768) | 1))
        _ = v9
    } else {
        var v10 uint8 = 233
        v10 = v10
        v11 := GenS1{F0: 153}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -173
}

func main() {
    v1 := []string{(("\n\t" + "\n\t") + "foo bar"), "", "hello"}
    var v2 int32 = 1
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    fmt.Println(v2)
    var v12 int16 = 0
    if true {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 int64 = 1
            v12 += (v12 - v12)
            v2 = v2
            _ = fi13
            _ = v14
        }
        v2 += (v2 * v2)
        fmt.Println(v12, v12)
    } else {
        var v15 uint64 = 2
        v2 -= v2
        _ = v15
    }
    var v16 uint16 = 42
    v16 += 0
    fmt.Println(v12, v2)
    v17 := GenS1{F0: 4294967295}
    v12 += -1
    v12 += (^(v12))
    var v18 float64 = (0.1 + 0.1)
    v18 = 1e10
    if ((v16 | (v16 & v16)) >= v16) {
        switch 0 {
        case -9223372036854775808:
            switch "\n\t" {
            case "foo bar":
                if false {
                    v2 = (v2 & (v2 ^ (^(v2))))
                } else {
                    var v19 float32 = 3.14159
                    var v20 uint16 = (v16 * 2)
                    _ = v19
                    _ = v20
                }
            case "a":
                v21 := GenS2{F0: 75}
                _ = v21
            }
        case 32767:
            v22 := GenS2{F0: 1}
            _ = v22
        default:
            fmt.Println(v18, v16)
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 int = fi23
            _ = fi23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        var v26 int = fi25
        var v27 float64 = (1e10 - (0.0 * (float64(0.0) + v18)))
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint64 = 1
    for fi29 := 0; fi29 < 16; fi29++ {
        v16 += v16
        var v30 int16 = v12
        var v31 uint = 4294967295
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println(v16)
    var v32 int64 = -532
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
}
