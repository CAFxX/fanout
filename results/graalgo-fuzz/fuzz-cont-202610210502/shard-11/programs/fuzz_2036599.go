package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint) uint64 {
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint16 = 719
        _ = fi6
        _ = v7
    }
    if false {
        if false {
            fmt.Println(p0)
            v8 := GenS1{F0: -32768}
            _ = v8
        } else {
            var v9 int32 = 1
            p0 -= p0
            p0 += p0
            _ = v9
        }
    }
    var v10 string = ""
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: 32767}
            var v14 float32 = 1.5
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := []int{(fi11 + 9223372036854775807)}
        var v16 int8 = -1
        _ = fi11
        _ = v15
        _ = v16
    }
    _ = p0
    _ = v10
    return 512
}
func genFn1(p0 uint, p1 uint) uint64 {
    if true {
        var v17 int32 = 42
        _ = v17
    } else {
        fmt.Println(p1, p0)
    }
    v18 := []int{-348, 2147483647}
    var v19 string = "foo bar"
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 0
}

func main() {
    var v1 int16 = 42
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    var v5 uint8 = 1
    v1 += v1
    switch 127 {
    case 4294967295:
        if true {
            switch 1 {
            case -1:
                var v20 uint32 = 392
                _ = v20
            default:
                v21 := v3[931]
                _ = v21
            }
            v1 += v1
            v22 := GenS1{F0: (v1 / ((55) | 1))}
            _ = v22
        }
    case 326:
        v1 += v1
    case 128:
        v5 += 2
    default:
        for fi23 := 0; fi23 < 6; fi23++ {
            fmt.Println(v1, v5, fi23)
            v24 := []int{(fi23 & fi23), (fi23 / (((-(fi23))) | 1)), 0, (^(fi23))}
            v25 := GenS1{F0: -876}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    switch ("hello" + ("" + "")) {
    case "foo bar":
        v26 := GenS2{F0: v1}
        _ = v26
    }
    fmt.Println(v1, v1)
    v5 = v5
    v1 = v1
    for fi27 := 0; fi27 < 16; fi27++ {
        v1 = (-(v1))
        _ = fi27
    }
    v28 := GenS2{F0: (v1 << ((1) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
