package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 int8) uint64 {
    var v4 int8 = (p1 / ((p1) | 1))
    p1 = (p1 ^ p0)
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 int32 = -2147483648
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    fmt.Println(p0, p1)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 float64 = (3.14159 / (1.5 - ((3.14159 / float64(-2.25)) - float64(float32(0.1)))))
    var v2 int8 = -57
    v3 := make(map[int]bool)
    var v8 string = "\n\t"
    var v9 uint64 = 1
    v10 := GenS1{F0: -760, F1: 574, F2: 764}
    v11 := GenS2{F0: v8}
    v8 = ("a" + v8)
    v12 := GenS1{F0: -2147483648, F1: 693, F2: 1}
    v2 -= (v2 ^ (v2 / ((0) | 1)))
    switch 2147483647 {
    case 42:
        v13 := GenS2{F0: v8}
        _ = v13
    }
    v14 := []int{0, 32768}
    for fi15 := 0; fi15 < 14; fi15++ {
        switch fi15 {
        case 1:
            var v16 uint8 = 42
            _ = v16
        case 127:
            v17 := GenS1{F0: 42, F1: 65535, F2: -2147483648}
            _ = v17
        case 32767:
            if false {
                v14 = append(v14, (fi15 ^ fi15))
                switch -2147483648 {
                case 1:
                    for fi18 := 0; fi18 < 13; fi18++ {
                        v19 := v3[(fi15 ^ (-(fi18)))]
                        if len(v14) > 0 {
                            v14[0] = fi15
                        }
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = fi18
                        _ = v19
                    }
                }
                for fi20 := 0; fi20 < 1; fi20++ {
                    _ = len(v3)
                    _ = fi20
                }
            }
        }
        var v21 int16 = -1
        _ = fi15
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
