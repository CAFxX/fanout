package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 1.5
    v3 := make(map[int]string)
    var v4 uint16 = 908
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int = (fi5 ^ -660)
        if true {
            v2 -= v2
        }
        _ = fi5
        _ = v6
    }
    var v7 int16 = 1
    fmt.Println(v2, v7)
    var v8 int32 = 209
    v9 := []int{1}
    if false {
        _ = len(v1)
        switch 397 {
        case 32767:
            for fi10 := 0; fi10 < 17; fi10++ {
                for fi11 := 0; fi11 < 17; fi11++ {
                    if (!(0 < -1)) {
                        v12 := GenS1{F0: 42, F1: 65535, F2: "hello"}
                        _ = v12
                    }
                    var v13 int = (fi11 | 0)
                    _ = fi11
                    _ = v13
                }
                v2 = -2.25
                _ = fi10
            }
        }
        v7 = -724
    } else {
        var v14 int64 = 42
        _ = v14
    }
    var v15 int64 = -1
    fmt.Println(v15, v4, v15)
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS2{F0: 0}
        if true {
            v2 = v2
        } else {
            v8 = -271
            v18 := GenS1{F0: 1000, F1: 2147483647, F2: ("\n\t" + "hello")}
            _ = v18
        }
        if true {
            v4 += (v4 ^ (v4 & (^(v4))))
        } else {
            if (0.1 != v2) {
                v2 = v2
                v19 := GenS3{F0: v8, F1: 42}
                _ = v19
            } else {
                v2 -= 0.1
                var v20 uint16 = (^(v4))
                v4 -= (v4 + 65535)
                _ = v20
            }
        }
        _ = fi16
        _ = v17
    }
    if true {
        var v21 int64 = -1
        _ = v21
    } else {
        for fi22 := 0; fi22 < 14; fi22++ {
            v7 = v7
            _ = fi22
        }
        if len(v9) > 0 {
            v9[0] = 868
        }
    }
    v8 = (-(v8))
    v7 += (-(v7))
    v8 -= v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
