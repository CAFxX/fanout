package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 string }

func main() {
    var v1 bool = (false && false)
    var v2 float64 = float64((1.5 * (-2.25 * 1.5)))
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    var v5 int32 = 911
    v6 := make(map[uint32]string)
    for fi7 := 0; fi7 < 7; fi7++ {
        if (0.0 == v2) {
            v8 := GenS1{F0: 1, F1: 1}
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 float32 = float32(v2)
                v1 = true
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 uint = 872
                if false {
                    v1 = false
                } else {
                    v13 := GenS2{F0: v5, F1: v12}
                    var v14 float32 = 0.1
                    if true {
                        var v15 uint = (^(v12))
                        switch (fi11 - (fi7 - 283)) {
                        case 32767:
                            v15 = (v15 * (v12 * v12))
                        case 125:
                            if (("" != "hello") && false) {
                                v1 = true
                                v12 -= (v12 | 255)
                                v5 -= v5
                            }
                        case 4294967295:
                            v16 := []int{(fi7 % ((fi7) | 1)), (fi11 ^ fi11)}
                            _ = v16
                        default:
                            v14 = (v14 * v14)
                        }
                        _ = v15
                    }
                    _ = v13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
        }
        _ = fi7
    }
    v17 := GenS3{F0: ""}
    fmt.Println(v2)
    var v18 float32 = float32(v2)
    fmt.Println(v2, v1)
    v18 += v18
    v19 := []int{}
    var v20 uint = 282
    v21 := GenS3{F0: (("\n\t" + "foo bar") + ("hello" + "a"))}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
