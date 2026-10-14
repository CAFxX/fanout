package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 int32, p1 uint32) uint {
    v7 := []int{4294967295}
    if false {
        p1 -= (p1 | 308)
        switch 2147483647 {
        case 922:
            var v8 int32 = p0
            _ = v8
        }
    }
    p1 += (p1 - p1)
    p1 += (p1 << ((42) & 63))
    _ = p0
    _ = p1
    _ = v7
    return 297
}
func genFn1() uint8 {
    var v9 uint = 4294967295
    var v10 int64 = 486
    fmt.Println(v10)
    fmt.Println(v9)
    _ = v10
    _ = v9
    return 16
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 42
    v3 := make(map[int]int)
    var v4 string = ("\n\t" + "hello")
    var v5 uint32 = 42
    var v6 uint32 = (v5 - v5)
    _ = len(v1)
    fmt.Println(v4, v4)
    v11 := GenS1{F0: 255, F1: v4}
    var v12 bool = false
    v12 = true
    if v12 {
        if (true && true) {
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 int64 = (v2 / (((v2 >> ((2) & 63))) | 1))
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint64 = 608
                for fi17 := 0; fi17 < 7; fi17++ {
                    v18 := GenS1{F0: 42, F1: v4}
                    fmt.Println(v2)
                    _ = fi17
                    _ = v18
                }
                v4 = (("" + v4) + v4)
                _ = fi15
                _ = v16
            }
        } else {
            var v19 uint8 = 128
            _ = v19
        }
        switch 873 {
        case 1:
            fmt.Println(v6)
        default:
            var v20 float32 = 1e10
            _ = v20
        }
        if (0 != 501) {
            if v12 {
                var v21 uint32 = v5
                if (128 >= 4294967295) {
                    if false {
                        v22 := []int{}
                        v23 := GenS1{F0: 974, F1: ("foo bar" + ("hello" + v4))}
                        var v24 uint64 = 601
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                } else {
                    for fi25 := 0; fi25 < 7; fi25++ {
                        v3[fi25] = (fi25 + (fi25 | fi25))
                        _ = fi25
                    }
                    v26 := []int{-278, -550}
                    switch int(v2) {
                    case 128:
                        v27 := GenS1{F0: 65535, F1: "a"}
                        _ = v27
                    }
                    _ = v26
                }
                if (float32(0.0) <= -2.25) {
                    var v28 int8 = 1
                    _ = v28
                }
                _ = v21
            } else {
                v5 = (v5 * (v6 / (((v6 / ((v6) | 1))) | 1)))
                v29 := GenS1{F0: 2, F1: v4}
                _ = v29
            }
        }
    }
    v6 -= v5
    v6 = (v5 - 637)
    v2 = 1
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
