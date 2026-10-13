package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int }

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    v3 := []int{-958, 2147483647}
    v4 := []uint32{946, 349, 65535, 550}
    v5 := make(map[int]uint32)
    if v2 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint16 = 116
        _ = fi6
        _ = v7
    }
    if true {
        if v2 {
            var v8 int = 451
            _ = v8
        } else {
            switch 127 {
            case -2147483648:
                if (!(!true)) {
                    var v9 uint = 902
                    _ = v9
                }
            default:
                var v10 uint16 = 0
                _ = v10
            }
        }
        _ = len(v3)
    } else {
        if v2 {
            var v11 int8 = -128
            var v12 int64 = -563
            _ = v11
            _ = v12
        } else {
            if len(v3) > 0 {
                v3[0] = -664
            }
            fmt.Println(v2)
        }
    }
    var v13 string = "foo bar"
    v2 = (-593 != 0)
    if (!v2) {
        v13 = "foo bar"
    } else {
        v2 = v2
        delete(v1, v13)
    }
    v14 := GenS1{F0: 42}
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 string = v13
        _ = fi15
        _ = v16
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = false
    var v17 uint = 576
    v18 := GenS2{F0: 2, F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
