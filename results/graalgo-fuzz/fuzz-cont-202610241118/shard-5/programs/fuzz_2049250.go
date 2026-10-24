package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint; F2 int }
func genFn0(p0 uint8, p1 string, p2 uint16) uint32 {
    v6 := []int{-738, 32767, -818, 32768}
    fmt.Println(p1, p1)
    var v7 float64 = 1e10
    fmt.Println(p1, p1)
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int32 = 441
        for fi10 := 0; fi10 < 6; fi10++ {
            if (false || false) {
                _ = len(v6)
            } else {
                _ = len(v6)
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 923
}
func genFn1() uint16 {
    v11 := []int{42, -888}
    v12 := GenS1{F0: -1}
    var v13 uint32 = 65535
    _ = v11
    _ = v12
    _ = v13
    return 99
}

func main() {
    var v1 string = ("a" + "hello")
    v2 := make(map[string]string)
    var v3 int = 128
    v4 := []string{(v1 + v1)}
    v5 := make(map[uint32]int)
    _ = len(v4)
    v3 = v3
    v3 = v3
    v14 := GenS2{F0: "", F1: 0, F2: (v3 % ((v3) | 1))}
    if (!false) {
        fmt.Println(v1)
    } else {
        _ = len(v5)
        v15 := GenS2{F0: v1, F1: 2, F2: v3}
        var v16 uint = 0
        _ = v15
        _ = v16
    }
    if false {
        v3 -= (v3 * v3)
        if false {
            v4 = append(v4, (v1 + (v1 + "\n\t")))
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 uint = 455
            _ = fi17
            _ = v18
        }
    }
    if ((0.0 / 1.0) == 1e10) {
        var v19 int = v3
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        v1 = (v1 + v1)
        _ = fi20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        if len(v4) > 0 {
            v4[0] = v1
        }
        _ = fi21
    }
    var v22 int32 = 42
    delete(v5, 787)
    if (619 >= 1000) {
        var v23 float64 = 0.0
        _ = len(v4)
        _ = v23
    }
    v1 = (v1 + ("hello" + v1))
    fmt.Println(v22)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
