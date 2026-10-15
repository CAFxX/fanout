package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint64 }
func genFn0(p0 int16, p1 uint8) int32 {
    var v5 float32 = 1e10
    for fi6 := 0; fi6 < 4; fi6++ {
        p1 -= 1
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    var v3 uint8 = 1
    v4 := make(map[int]int)
    if (633 == 0) {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 uint16 = 255
            _ = fi7
            _ = v8
        }
    }
    v3 += (v3 * v3)
    switch 127 {
    case 4294967295:
        v9 := []int{1, 128}
        _ = v9
    case 299:
        for fi10 := 0; fi10 < 20; fi10++ {
            for fi11 := 0; fi11 < 10; fi11++ {
                if true {
                    v12 := GenS1{F0: 65535, F1: 2, F2: 1}
                    _ = v12
                }
                _ = fi11
            }
            v13 := GenS1{F0: 130, F1: 1, F2: 42}
            _ = fi10
            _ = v13
        }
    default:
        delete(v4, -172)
    }
    if true {
        var v14 uint16 = 2
        _ = v14
    }
    v3 -= v3
    _ = len(v2)
    fmt.Println(v3)
    v15 := []int{0}
    switch "foo bar" {
    case "a":
        v3 += v3
    default:
        var v16 uint64 = 334
        _ = v16
    }
    v3 = v3
    var v17 uint32 = 500
    var v18 uint32 = (v17 & v17)
    var v19 uint16 = 0
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 bool = true
        _ = fi20
        _ = v21
    }
    var v22 string = ("hello" + "")
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 int = 32767
            _ = fi24
            _ = v25
        }
        v3 = (v3 + 127)
        for fi26 := 0; fi26 < 20; fi26++ {
            v18 -= uint32(v3)
            v27 := GenS1{F0: 982, F1: 124, F2: 0}
            if len(v15) > 0 {
                v15[0] = fi26
            }
            _ = fi26
            _ = v27
        }
        _ = fi23
    }
    switch v22 {
    case "foo bar":
        var v28 uint8 = (v3 - 127)
        _ = v28
    case "\n\t":
        _ = len(v15)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
