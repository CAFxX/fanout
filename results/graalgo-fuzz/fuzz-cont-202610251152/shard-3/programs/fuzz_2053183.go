package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 bool; F2 bool }

func main() {
    v1 := make(map[int]bool)
    var v2 int = -9223372036854775808
    var v3 bool = false
    var v4 int = -1
    v2 += v4
    v3 = v3
    var v5 uint8 = 2
    v1[128] = ((!(2 >= 32767)) && true)
    v2 += -956
    switch (v2 / ((32767) | 1)) {
    case 9223372036854775807:
        v3 = (!true)
    default:
        v6 := GenS1{F0: 2}
        _ = v6
    }
    v5 = (v5 << ((0) & 63))
    var v7 uint = 677
    if v3 {
        switch "\n\t" {
        case "a":
            var v8 int32 = 0
            _ = v8
        case "":
            fmt.Println(v4, v4)
        case "\n\t":
            if false {
                v9 := v1[(v4 >> ((v7) & 63))]
                if (v2 >= (v4 * -361)) {
                    var v10 float64 = 0.1
                    for fi11 := 0; fi11 < 14; fi11++ {
                        var v12 string = ("hello" + (("a" + "\n\t") + "\n\t"))
                        v13 := []int{2147483647, (v2 / ((fi11) | 1))}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                } else {
                    delete(v1, v2)
                    v7 = (v7 * (v7 % (((^(v7))) | 1)))
                    v14 := GenS1{F0: 2}
                    _ = v14
                }
                _ = v9
            }
        }
        v15 := GenS2{F0: 676, F1: v3, F2: (1.5 <= 0.1)}
        fmt.Println(v7)
        _ = v15
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := GenS1{F0: 65535}
            _ = fi16
            _ = v17
        }
    }
    var v18 bool = (1 >= 42)
    var v19 uint32 = 2
    var v20 int = v4
    fmt.Println(v7, v4)
    v18 = v18
    v21 := []int{v2}
    if len(v21) > 0 {
        v21[0] = (v4 + (-(v2)))
    }
    if len(v21) > 0 {
        v21[0] = 110
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 uint = (v7 / ((v7) | 1))
        fmt.Println(v19, v19, v20)
        _ = fi22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
