package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint16, p1 uint16) int {
    var v5 int8 = -1
    var v6 string = "\n\t"
    v7 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 string) int {
    v8 := []int{127, genFn0(2, 563)}
    if false {
        var v9 uint = 452
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS1{F0: -32768}
            v12 := GenS2{F0: 1000, F1: 1}
            var v13 int = fi10
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        fmt.Println(p0, v9)
        _ = v9
    } else {
        switch genFn0(32767, 42) {
        case 677:
            v14 := GenS1{F0: 0}
            _ = v14
        case 875:
            p0 = ((p0 + p0) + (p0 + p0))
        }
        var v15 uint8 = 2
        var v16 int32 = 42
        _ = v15
        _ = v16
    }
    _ = p0
    _ = v8
    return 32768
}

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]bool)
    var v3 int16 = 288
    var v4 float64 = ((0.0 - ((1.5 - 1e10) / -2.25)) / 3.14159)
    v3 = v3
    for fi17 := 0; fi17 < 11; fi17++ {
        v3 = -32768
        _ = fi17
    }
    if true {
        for fi18 := 0; fi18 < 20; fi18++ {
            v1 += v1
            _ = fi18
        }
        var v19 uint32 = 240
        var v20 int32 = v1
        _ = v19
        _ = v20
    }
    v4 -= (float64(0.1) / v4)
    v4 -= (0.0 - (v4 - (v4 / 1.0)))
    v4 += v4
    v2[65535] = (true || false)
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            switch "a" {
            case "a":
                var v23 int16 = (^(v3))
                _ = v23
            }
            _ = fi22
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            fmt.Println(fi21)
            _ = fi24
        }
        _ = fi21
    }
    v3 = v3
    var v25 uint = 42
    if (!false) {
        v26 := GenS2{F0: v3, F1: 81}
        _ = v26
    } else {
        fmt.Println(v4, v25)
    }
    v1 -= v1
    delete(v2, 517)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
