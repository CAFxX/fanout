package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint16, p1 bool, p2 int32) uint8 {
    p0 += (p0 / ((p0) | 1))
    p2 -= p2
    fmt.Println(p1, p0)
    if p1 {
        var v4 bool = false
        v5 := []int{}
        fmt.Println(p1, v4)
        _ = v4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 235
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = (("hello" + ((("hello" + "") + "\n\t") + "")) + "")
    v3 := []uint32{}
    var v6 string = v2
    for fi7 := 0; fi7 < 7; fi7++ {
        v6 = v6
        var v8 uint = 255
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        if false {
            var v10 string = v6
            if false {
                v2 = ("a" + "foo bar")
                v11 := GenS1{F0: 455, F1: 4294967295}
                _ = v11
            } else {
                var v12 uint = 823
                v13 := GenS1{F0: -921, F1: 18446744073709551615}
                v2 = v6
                _ = v12
                _ = v13
            }
            _ = v10
        } else {
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint64 = 0
                fmt.Println(fi14, fi14, fi14)
                v2 = "\n\t"
                _ = fi14
                _ = v15
            }
            fmt.Println(fi9)
            var v16 float64 = (-2.25 + ((1.5 / 0.1) / 0.1))
            _ = v16
        }
        _ = len(v1)
        _ = fi9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v17 := GenS1{F0: 51, F1: 0}
    var v18 int16 = 32767
    v2 = (v2 + v6)
    v19 := GenS1{F0: (v18 >> ((65535) & 63)), F1: 210}
    v18 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
