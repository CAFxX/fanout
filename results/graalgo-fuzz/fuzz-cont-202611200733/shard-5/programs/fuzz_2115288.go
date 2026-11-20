package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int }
func genFn0() uint32 {
    var v4 uint = 2
    if (65535 > 2147483648) {
        v4 -= (^(v4))
    }
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 float32 = (-2.25 - float32((3.14159 / ((3.14159 - 3.14159) - float64(-2.25)))))
        v6 -= 0.0
        _ = fi5
        _ = v6
    }
    fmt.Println(v4)
    var v7 uint8 = 159
    _ = v4
    _ = v7
    return 0
}

func main() {
    var v1 float64 = (float64((0.0 + 1.5)) + ((-2.25 * 3.14159) * 0.1))
    var v2 float64 = (0.1 / 0.1)
    var v3 float32 = (1.5 / (0.1 + 3.14159))
    v2 = 3.14159
    fmt.Println(v3, v3)
    v8 := []int{32767, -915, 128}
    v1 -= ((v1 * v1) / v1)
    var v9 bool = (127 <= 100)
    for fi10 := 0; fi10 < 19; fi10++ {
        v2 = (float64((0.0 + v3)) + 1.5)
        v9 = (127 != 89)
        var v11 int = (fi10 * fi10)
        _ = fi10
        _ = v11
    }
    v2 += v1
    v12 := GenS1{F0: 45}
    var v13 bool = (1 > -32768)
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int32 = 42
        var v16 uint32 = 4294967295
        switch ("" + "foo bar") {
        case "":
            switch fi14 {
            case -32:
                if v13 {
                    _ = len(v8)
                }
            case 114:
                var v17 bool = (42 <= -1)
                _ = v17
            default:
                var v18 float64 = ((3.14159 * 1e10) / (v2 - v2))
                _ = v18
            }
        case "foo bar":
            var v19 int = (fi14 & (-(fi14)))
            _ = v19
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v20 bool = false
    var v21 uint = 65535
    switch "\n\t" {
    case "a":
        for fi22 := 0; fi22 < 9; fi22++ {
            if true {
                v23 := GenS2{F0: 2, F1: -56}
                _ = v23
            } else {
                for fi24 := 0; fi24 < 17; fi24++ {
                    fmt.Println(v21)
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v25 string = "foo bar"
                    if len(v8) > 0 {
                        v8[0] = fi24
                    }
                    _ = fi24
                    _ = v25
                }
                v9 = (!v20)
            }
            _ = fi22
        }
    case "foo bar":
        var v26 int = 127
        _ = v26
    case "":
        var v27 float32 = 3.14159
        _ = v27
    default:
        v21 -= 267
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        v29 := GenS1{F0: 201}
        fmt.Println(v2)
        v2 = ((v1 * float64((1.5 * 3.14159))) - v2)
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        v2 -= v1
        _ = len(v8)
        v31 := GenS3{F0: fi30}
        _ = fi30
        _ = v31
    }
    v9 = v13
    for fi32 := 0; fi32 < 13; fi32++ {
        if (float64(v3) == 1.5) {
            v2 = 3.14159
            for fi33 := 0; fi33 < 13; fi33++ {
                var v34 int16 = 42
                var v35 bool = v9
                fmt.Println(v2, v2)
                _ = fi33
                _ = v34
                _ = v35
            }
        } else {
            var v36 int = (fi32 % ((fi32) | 1))
            _ = v36
        }
        var v37 int64 = 1
        _ = fi32
        _ = v37
    }
    for fi38 := 0; fi38 < 6; fi38++ {
        v1 -= ((v1 / v1) - v2)
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v8
    _ = v9
}
