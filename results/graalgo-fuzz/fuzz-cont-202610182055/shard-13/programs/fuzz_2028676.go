package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 int16 = -45
    var v4 float32 = float32((3.14159 - (1e10 / 1e10)))
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(v3)
        for fi6 := 0; fi6 < 2; fi6++ {
            v4 = (((v4 * v4) + float32(-2.25)) / ((v4 + v4) * 0.0))
            _ = fi6
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            v4 -= (float32(((0.0 * (((1.5 - 0.0) / 3.14159) / ((0.1 / ((1.5 - (((0.1 * 1.5) * 3.14159) + (0.1 / 0.1))) / (0.1 * float64(float32(-2.25))))) / float64(0.1)))) + 1e10)) / 1.5)
            v8 := GenS1{F0: 1, F1: -765, F2: 1}
            v9 := GenS1{F0: 42, F1: -437, F2: -123}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
    }
    var v10 int8 = -13
    fmt.Println(v3)
    var v11 float64 = (1.5 + 3.14159)
    switch 1 {
    case -58:
        var v12 uint8 = 0
        _ = v12
    case -36:
        var v13 uint = 610
        _ = v13
    case 408:
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint8 = 1
            switch (-(fi14)) {
            case 128:
                switch ("foo bar" + (("foo bar" + "foo bar") + ("\n\t" + "foo bar"))) {
                case "a":
                    v16 := []int{1, fi14}
                    _ = v16
                }
            case 32767:
                switch (("hello" + ("foo bar" + ("" + "foo bar"))) + "\n\t") {
                case "a":
                    var v17 float32 = (v4 / (((float32(float64((float32(v11) + v4))) * v4) * (v4 / float32(1e10))) - v4))
                    _ = v17
                }
            }
            var v18 float64 = 1e10
            _ = fi14
            _ = v15
            _ = v18
        }
    default:
        fmt.Println(v10, v4, v11)
    }
    v10 += v10
    var v19 string = ("\n\t" + "foo bar")
    if (2 != 1) {
        var v20 int = 1
        for fi21 := 0; fi21 < 7; fi21++ {
            if false {
                v3 = v3
                v22 := GenS2{F0: 1, F1: true}
                _ = v22
            } else {
                v11 += v11
            }
            _ = fi21
        }
        _ = v20
    } else {
        v23 := GenS1{F0: 0, F1: int64(v3), F2: (v10 >> ((1) & 63))}
        var v24 uint64 = 1
        var v25 int8 = int8(v3)
        _ = v23
        _ = v24
        _ = v25
    }
    v1[128] = 2147483647
    v26 := []int{32767, -2147483648}
    fmt.Println(v19, v19)
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    var v27 float64 = v11
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := GenS3{F0: (-(fi28)), F1: 1}
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 int8 = v10
            v26 = append(v26, fi28)
            v31 -= (v31 & v31)
            _ = fi30
            _ = v31
        }
        for fi32 := 0; fi32 < 16; fi32++ {
            fmt.Println(v10, v10, v3)
            _ = fi32
        }
        _ = fi28
        _ = v29
    }
    switch 969 {
    case 32767:
        var v33 uint32 = 42
        _ = v33
    default:
        v34 := GenS4{F0: v10}
        _ = v34
    }
    var v35 bool = (!false)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v4
}
