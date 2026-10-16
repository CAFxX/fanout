package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16) uint64 {
    var v5 float64 = 0.0
    if true {
        for fi6 := 0; fi6 < 12; fi6++ {
            switch (^(fi6)) {
            case 4294967295:
                if ((p0 >> ((42) & 63)) <= 1) {
                    if (0 >= 255) {
                        if true {
                            var v7 uint64 = 1
                            v8 := []int{}
                            _ = v7
                            _ = v8
                        }
                        p0 -= p0
                    }
                    var v9 int64 = 372
                    _ = v9
                }
            }
            fmt.Println(v5)
            v10 := GenS1{F0: 0, F1: fi6}
            _ = fi6
            _ = v10
        }
    }
    v5 += ((0.0 / v5) * (1.5 - (float64(1.5) / (3.14159 + v5))))
    _ = p0
    _ = v5
    return 525
}
func genFn1(p0 int8, p1 int) uint32 {
    var v11 float32 = float32(((float64(1.5) - 1.5) - 1.5))
    if false {
        var v12 uint8 = 181
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 string = "a"
            v11 = float32(1.5)
            var v15 uint = 42
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: 2147483647}
        _ = v12
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        p0 -= p0
        _ = fi17
    }
    v18 := []int{int(p0)}
    var v19 int8 = p0
    _ = p0
    _ = p1
    _ = v11
    _ = v18
    _ = v19
    return 747
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int = 128
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    var v20 string = (("hello" + "") + ("a" + (("hello" + "foo bar") + "")))
    for fi21 := 0; fi21 < 5; fi21++ {
        fmt.Println(v1, fi21)
        v1 -= (v1 >> ((18446744073709551615) & 63))
        v22 := []int{}
        _ = fi21
        _ = v22
    }
    fmt.Println(v20)
    v2 = v2
    v2 += v2
    fmt.Println(v20)
    v1 -= 4294967295
    v23 := GenS1{F0: -49, F1: (v2 / ((707) | 1))}
    var v24 uint16 = 1
    switch (v2 >> ((736) & 63)) {
    case 127:
        v2 += int(v24)
    case 42:
        v25 := []int{}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
