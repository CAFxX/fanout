package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int, p1 string) int32 {
    fmt.Println(p0)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 float64 = float64((0.0 * float32(0.0)))
    var v2 bool = (!true)
    var v3 int8 = 100
    v4 := make(map[int]int)
    v5 := []string{("hello" + ""), "hello"}
    v6 := GenS1{F0: 298, F1: 371}
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(v2, v1, v2)
        for fi8 := 0; fi8 < 2; fi8++ {
            if ((fi7 + (fi7 & (fi8 % ((fi8) | 1)))) >= (^(fi8))) {
                var v9 uint64 = 0
                var v10 uint32 = 217
                _ = v10
                _ = v9
            }
            var v11 string = "hello"
            v12 := GenS1{F0: 1, F1: 839}
            if v2 {
                _ = len(v5)
                fmt.Println(v3)
            } else {
                if len(v5) > 0 {
                    v5[0] = ""
                }
                v13 := v4[(fi8 * 0)]
                _ = v13
            }
            _ = fi8
            _ = v11
            _ = v12
        }
        v3 -= (v3 - v3)
        _ = fi7
    }
    v14 := v4[-489]
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 int32 = -937
        _ = fi15
        _ = v16
    }
    var v17 int8 = v3
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if v2 {
        fmt.Println(v14, v2, v3)
        if true {
            v18 := GenS1{F0: 434, F1: 0}
            v2 = v2
            _ = v18
        }
        var v19 uint = 4294967295
        _ = v19
    } else {
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 float32 = -2.25
            switch 128 {
            case -1:
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 float32 = float32(((3.14159 * 1.5) + v1))
                    if true {
                        v14 += fi22
                    }
                    _ = fi22
                    _ = v23
                }
            case -750:
                var v24 int16 = 42
                _ = v24
            case 127:
                var v25 string = "hello"
                _ = v25
            }
            for fi26 := 0; fi26 < 2; fi26++ {
                var v27 uint8 = 2
                _ = fi26
                _ = v27
            }
            _ = fi20
            _ = v21
        }
        v14 = (^(v14))
    }
    var v28 int64 = 762
    _ = len(v4)
    switch "\n\t" {
    case "a":
        for fi29 := 0; fi29 < 6; fi29++ {
            v3 += 1
            v5 = append(v5, "\n\t")
            _ = fi29
        }
    }
    v17 -= v3
    var v30 int = (v14 - 0)
    v31 := GenS2{F0: 2}
    var v32 bool = (!(!(v2 && v2)))
    var v33 float32 = 0.0
    var v34 bool = true
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v35 uint = 209
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
