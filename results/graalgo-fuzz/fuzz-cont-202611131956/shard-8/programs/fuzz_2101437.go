package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint32; F2 string }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 int8) int16 {
    v4 := GenS1{F0: 2}
    var v5 uint = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -802
}

func main() {
    v1 := []int{-931, 4294967295, 26}
    var v2 string = ("hello" + "hello")
    v3 := make(map[uint32]bool)
    v2 = (v2 + "")
    if false {
        var v6 uint32 = 1
        var v7 uint64 = 594
        _ = v6
        _ = v7
    } else {
        switch -2147483648 {
        case 4294967295:
            v8 := v3[465]
            _ = v8
        case 0:
            if false {
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 int32 = int32(fi9)
                    _ = fi9
                    _ = v10
                }
                v2 = (v2 + "")
            } else {
                if true {
                    var v11 bool = (0 >= 2)
                    _ = v11
                }
                for fi12 := 0; fi12 < 2; fi12++ {
                    v3[42] = true
                    _ = fi12
                }
            }
        }
    }
    v13 := GenS2{F0: 1, F1: 2, F2: v2}
    for fi14 := 0; fi14 < 12; fi14++ {
        if ((0 != -128) || (127 != 42)) {
            v2 = v2
        }
        _ = fi14
    }
    v2 = v2
    for fi15 := 0; fi15 < 7; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            if true {
                fmt.Println(fi16, v2)
            } else {
                v17 := GenS3{F0: 255, F1: 9}
                fmt.Println(fi15)
                fmt.Println(fi15, v2, fi15)
                _ = v17
            }
            fmt.Println(v2, fi16, fi16)
            v18 := GenS4{F0: v2}
            _ = fi16
            _ = v18
        }
        _ = fi15
    }
    v2 = (v2 + v2)
    v3[2147483648] = ((809 == 2) || (((3.14159 * 1e10) >= ((1e10 * 0.1) - 1.5)) || (0 < 255)))
    var v19 float32 = float32(-2.25)
    v2 = v2
    var v20 int16 = -278
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
