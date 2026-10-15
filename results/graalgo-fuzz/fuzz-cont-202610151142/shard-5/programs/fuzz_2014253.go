package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 string; F1 int32; F2 uint }
func genFn0() uint {
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(fi5)
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(fi5, fi6)
            var v7 bool = (2 != 912)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v8 bool = true
    var v9 float64 = ((-2.25 / 1.0) + (float64(0.1) + -2.25))
    if (100 == 1) {
        v10 := GenS1{F0: 1, F1: (("a" + ("a" + "a")) + "foo bar")}
        v11 := GenS2{F0: "\n\t", F1: 1, F2: 419}
        v12 := []int{4294967295}
        _ = v10
        _ = v11
        _ = v12
    }
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int16, p1 int16) uint32 {
    p0 += p1
    p0 += (-(p1))
    p1 -= (-(p1))
    switch ("\n\t" + ("" + "foo bar")) {
    case "":
        var v13 float64 = -2.25
        _ = v13
    default:
        v14 := GenS2{F0: "", F1: 0, F2: genFn0()}
        _ = v14
    }
    var v15 uint32 = 134
    _ = p0
    _ = p1
    _ = v15
    return 709
}

func main() {
    var v1 int16 = 582
    v2 := []int{-913, -308}
    var v3 int = 399
    v4 := make(map[uint32]string)
    switch (v3 & v3) {
    case -660:
        v3 += (^(v3))
    default:
        v16 := GenS1{F0: 74, F1: ("hello" + "\n\t")}
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        if true {
            if true {
                var v18 uint8 = 127
                var v19 uint = 0
                _ = v18
                _ = v19
            }
        }
        _ = fi17
    }
    v4[0] = ""
    v3 += 4294967295
    _ = len(v2)
    var v20 int32 = -240
    var v21 int = (^(v3))
    fmt.Println(v3)
    v20 += (-(v20))
    v4[0] = ("a" + "")
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 int64 = 0
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
