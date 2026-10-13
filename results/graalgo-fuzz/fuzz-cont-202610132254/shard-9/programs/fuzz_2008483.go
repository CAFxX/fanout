package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool) int64 {
    var v6 uint64 = 2
    v7 := GenS1{F0: 1, F1: 1}
    _ = p0
    _ = v6
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 bool, p1 uint32, p2 int8) int64 {
    p1 = 1
    p0 = true
    var v8 int64 = 9223372036854775807
    v9 := []int{}
    if (!p0) {
        var v10 uint = 646
        if p0 {
            if len(v9) > 0 {
                v9[0] = 873
            }
            fmt.Println(p2, p2, p0)
        }
        _ = v10
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            if len(v9) > 0 {
                v9[0] = fi11
            }
            if (!(127 < 2)) {
                _ = len(v9)
                fmt.Println(fi11)
                p2 += 0
            } else {
                switch "a" {
                case "a":
                    for fi12 := 0; fi12 < 5; fi12++ {
                        switch fi11 {
                        case 551:
                            v13 := GenS2{F0: -1}
                            _ = v13
                        }
                        _ = fi12
                    }
                }
                var v14 uint64 = 815
                var v15 float64 = -2.25
                _ = v14
                _ = v15
            }
            var v16 uint32 = 500
            _ = fi11
            _ = v16
        }
        v17 := GenS2{F0: 0}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 430
    var v3 uint16 = 2
    v4 := make(map[int]bool)
    var v5 uint32 = 4294967295
    var v18 uint32 = v5
    var v19 int64 = 202
    if ((true && false) || (32767 <= -32768)) {
        var v20 float32 = -2.25
        var v21 uint = v2
        _ = v20
        _ = v21
    } else {
        v19 += v19
        var v22 uint32 = 0
        v23 := GenS2{F0: 0}
        _ = v22
        _ = v23
    }
    fmt.Println(v3)
    var v24 uint32 = (v5 * v18)
    v25 := GenS1{F0: int16(v2), F1: 0}
    switch 608 {
    case 9223372036854775807:
        var v26 uint64 = uint64(v2)
        _ = v26
    default:
        v27 := GenS2{F0: 2147483647}
        _ = v27
    }
    var v28 float32 = 0.0
    var v29 uint = 556
    var v30 uint32 = 767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
