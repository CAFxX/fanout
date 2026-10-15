package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
func genFn0() int8 {
    v4 := GenS1{F0: 42, F1: 1}
    switch 0 {
    case 42:
        v5 := []int{-2147483648, 1, 42}
        _ = v5
    case -522:
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := []int{fi6, (-(fi6))}
            _ = len(v7)
            _ = fi6
            _ = v7
        }
    case 0:
        for fi8 := 0; fi8 < 18; fi8++ {
            switch "\n\t" {
            case "hello":
                switch fi8 {
                case 32767:
                    fmt.Println(fi8)
                case 680:
                    var v9 int32 = -1
                    _ = v9
                }
            case "":
                var v10 int16 = 386
                _ = v10
            }
            var v11 float32 = ((1.5 + 0.1) * (0.1 - (1.5 + (0.0 * 0.0))))
            var v12 string = ("hello" + "foo bar")
            _ = fi8
            _ = v11
            _ = v12
        }
    }
    var v13 int16 = 202
    _ = v13
    _ = v4
    return -115
}

func main() {
    var v1 int32 = -1
    v2 := []string{"\n\t"}
    var v3 float64 = 1e10
    v3 = (v3 - 0.1)
    v1 -= v1
    v3 = ((0.0 * v3) / v3)
    var v14 uint64 = 2
    v3 += -2.25
    v2 = append(v2, "a")
    var v15 uint64 = v14
    var v16 string = "foo bar"
    switch 780 {
    case -2147483648:
        fmt.Println(v15, v16)
    }
    var v17 uint32 = 50
    var v18 int8 = 0
    if len(v2) > 0 {
        v2[0] = v16
    }
    v18 = v18
    var v19 float64 = 1.5
    fmt.Println(v18, v18)
    v17 = 42
    var v20 int = 32767
    v21 := GenS1{F0: v18, F1: (v14 << ((0) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
