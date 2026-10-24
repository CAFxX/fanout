package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 int }
func genFn0(p0 bool) int {
    p0 = p0
    switch 42 {
    case -2147483648:
        p0 = (p0 && p0)
    default:
        if true {
            var v4 int16 = -95
            v5 := []int{}
            p0 = (783 <= 563)
            _ = v4
            _ = v5
        } else {
            fmt.Println(p0)
            p0 = false
        }
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        if true {
            p0 = (22 != 222)
        } else {
            fmt.Println(fi6, fi6)
            if false {
                var v7 uint16 = 598
                var v8 uint = 255
                _ = v7
                _ = v8
            }
        }
        _ = fi6
    }
    if p0 {
        if p0 {
            fmt.Println(p0)
        } else {
            v9 := []int{401, -9223372036854775808, 32767, 2147483647}
            _ = v9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(p0, p0)
            var v11 uint8 = 128
            fmt.Println(v11, v11, p0)
            _ = fi10
            _ = v11
        }
    } else {
        p0 = p0
        v12 := []int{}
        _ = v12
    }
    v13 := GenS1{F0: 247, F1: 1, F2: 0}
    _ = p0
    _ = v13
    return -1
}

func main() {
    v1 := []int{-1, 32768, -1}
    var v2 uint8 = 0
    var v3 int16 = -32768
    v3 -= v3
    if (-1 == 1) {
        for fi14 := 0; fi14 < 16; fi14++ {
            switch (-(fi14)) {
            case 4294967295:
                fmt.Println(v3)
            case 128:
                var v15 float32 = float32(-2.25)
                _ = v15
            }
            for fi16 := 0; fi16 < 11; fi16++ {
                _ = len(v1)
                var v17 uint = 18446744073709551615
                _ = fi16
                _ = v17
            }
            var v18 float32 = 1.5
            _ = fi14
            _ = v18
        }
        v19 := GenS2{F0: (^(v3)), F1: true, F2: 32767}
        var v20 float64 = float64(1e10)
        _ = v19
        _ = v20
    }
    var v21 int = 128
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 uint16 = 255
        var v24 uint8 = v2
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 uint = 18446744073709551615
    v25 += 65535
    v25 += (v25 | (^(v25)))
    v2 = 127
    var v26 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
}
