package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint32; F2 int64 }
func genFn0(p0 uint8) uint16 {
    p0 += p0
    if false {
        p0 = 118
    }
    var v6 uint64 = 1
    _ = p0
    _ = v6
    return 32767
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 float32 = ((1.5 + -2.25) * 0.0)
    switch "\n\t" {
    case "foo bar":
        v7 := GenS1{F0: genFn0(255)}
        _ = v7
    }
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 20; fi9++ {
        delete(v3, (fi9 ^ -356))
        v10 := GenS1{F0: 32767}
        var v11 int = -2147483648
        _ = fi9
        _ = v10
        _ = v11
    }
    if (-1 <= -101) {
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(fi12, v5)
            _ = fi12
        }
        var v13 int = -365
        v3[682] = (0 != 70)
        _ = v13
    } else {
        v3[-9223372036854775808] = true
    }
    var v14 uint64 = 2
    v14 -= (^(v14))
    if true {
        v15 := GenS2{F0: (true || false), F1: 65535, F2: -1}
        v16 := v4[2147483647]
        _ = v15
        _ = v16
    } else {
        fmt.Println(v1, v14, v1)
        v17 := []int{}
        v1 = 1e10
        _ = v17
    }
    v18 := v2[32767]
    var v19 string = "a"
    for fi20 := 0; fi20 < 14; fi20++ {
        v21 := []int{0, fi20}
        if len(v21) > 0 {
            v21[0] = fi20
        }
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if (-1 == -1) {
            switch v19 {
            case "\n\t":
                switch (("hello" + "") + v19) {
                case "a":
                    v14 += (^(v14))
                case "hello":
                    v23 := GenS2{F0: false, F1: 504, F2: 1}
                    _ = v23
                default:
                    v24 := GenS2{F0: false, F1: 413, F2: 1}
                    _ = v24
                }
            case "a":
                v1 += (v1 / 1.5)
            default:
                fmt.Println(fi22, v19)
            }
            var v25 bool = ("" <= "a")
            _ = v25
        } else {
            v1 = v1
        }
        _ = fi22
    }
    var v26 int32 = -2147483648
    v27 := []int{-254, 812, 2147483647, -804}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
