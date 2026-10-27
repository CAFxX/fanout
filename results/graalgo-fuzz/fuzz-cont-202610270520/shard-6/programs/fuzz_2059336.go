package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int32 }
func genFn0(p0 int32) int32 {
    if false {
        v5 := []int{-1, 583, 0, -9223372036854775808}
        switch "a" {
        case "\n\t":
            fmt.Println(p0)
        case "":
            var v6 bool = (true || true)
            _ = v6
        case "hello":
            var v7 uint16 = 138
            _ = v7
        default:
            p0 = (p0 | p0)
        }
        var v8 uint = 1
        _ = v5
        _ = v8
    }
    var v9 int8 = -1
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int = (fi10 - (fi10 / ((-259) | 1)))
        p0 -= p0
        v9 -= 100
        _ = fi10
        _ = v11
    }
    switch -807 {
    case 1:
        v9 += v9
    case 128:
        p0 = (p0 - p0)
    }
    _ = p0
    _ = v9
    return int32(v9)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    var v3 int16 = 95
    var v4 uint16 = 32767
    v12 := []int{127}
    v13 := v1[32768]
    var v14 uint8 = 127
    _ = len(v12)
    v15 := GenS1{F0: -105}
    if (4294967295 > 323) {
        for fi16 := 0; fi16 < 10; fi16++ {
            v14 -= v14
            if true {
                v12 = append(v12, (fi16 ^ fi16))
                var v17 uint8 = 1
                v17 += v14
                _ = v17
            }
            switch "foo bar" {
            case "\n\t":
                var v18 uint64 = 1
                _ = v18
            case "hello":
                v19 := GenS1{F0: -1}
                _ = v19
            }
            _ = fi16
        }
    } else {
        _ = len(v12)
        var v20 int32 = 223
        var v21 uint8 = 1
        _ = v20
        _ = v21
    }
    switch 4294967295 {
    case 42:
        v22 := GenS2{F0: v13, F1: v14, F2: 0}
        _ = v22
    default:
        fmt.Println(v13, v2)
    }
    var v23 string = ((((("" + "\n\t") + "") + "foo bar") + "\n\t") + "")
    v24 := v1[2147483647]
    fmt.Println(v24)
    var v25 bool = (uint16(v14) <= (v4 ^ (v4 + (^(v4)))))
    fmt.Println(v2, v23, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
