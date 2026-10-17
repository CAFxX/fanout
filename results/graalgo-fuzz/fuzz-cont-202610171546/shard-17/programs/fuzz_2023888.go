package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int; F1 uint; F2 int }
func genFn0(p0 int64, p1 uint16, p2 uint16) int64 {
    switch 32768 {
    case -9223372036854775808:
        p1 += (p2 - p2)
    case 42:
        var v6 int8 = -128
        _ = v6
    }
    p2 -= 1
    var v7 int32 = -675
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 string = "a"
    v2 := []int{127, -482}
    v3 := make(map[uint32]bool)
    var v4 float64 = 3.14159
    v5 := make(map[string]uint32)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v8 uint = 255
    var v9 float32 = -2.25
    v5[(("" + v1) + (((v1 + v1) + ((v1 + (v1 + v1)) + v1)) + ("\n\t" + v1)))] = 47
    v4 -= v4
    switch ("\n\t" + (v1 + v1)) {
    case "":
        v4 += 3.14159
    case "hello":
        v10 := GenS1{F0: (v1 + (v1 + v1)), F1: ((v1 + v1) + v1)}
        _ = v10
    default:
        var v11 uint8 = 127
        _ = v11
    }
    var v12 bool = (587 < -540)
    for fi13 := 0; fi13 < 16; fi13++ {
        v8 -= v8
        fmt.Println(v4)
        if (v12 || (v1 <= "a")) {
            var v14 int8 = 42
            _ = v14
        } else {
            v2 = append(v2, fi13)
            var v15 uint32 = 583
            _ = v15
        }
        _ = fi13
    }
    var v16 int = 32767
    switch ("" + "a") {
    case "hello":
        switch "\n\t" {
        case "":
            switch (v1 + "foo bar") {
            case "hello":
                for fi17 := 0; fi17 < 20; fi17++ {
                    _ = len(v2)
                    fmt.Println(v8)
                    _ = fi17
                }
            case "":
                for fi18 := 0; fi18 < 15; fi18++ {
                    if v12 {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            var v20 float32 = 0.1
                            _ = fi19
                            _ = v20
                        }
                        if ((v16 / (((v16 & fi18)) | 1)) == (v16 * fi18)) {
                            v8 = v8
                            var v21 int8 = 127
                            _ = v21
                        } else {
                            var v22 uint16 = 42
                            _ = v22
                        }
                        if ((-94 != 42) && (0 >= 0)) {
                            v23 := GenS2{F0: v16, F1: (v8 / (((v8 * (v8 * (^(v8))))) | 1)), F2: fi18}
                            v1 = "\n\t"
                            _ = v23
                        }
                    } else {
                        fmt.Println(v12, v8, v9)
                        for fi24 := 0; fi24 < 16; fi24++ {
                            if (!false) {
                                fmt.Println(fi24, v16, v16)
                            }
                            var v25 int8 = 100
                            fmt.Println(v25)
                            _ = fi24
                            _ = v25
                        }
                    }
                    _ = fi18
                }
            }
        }
    case "a":
        var v26 float64 = v4
        _ = v26
    default:
        switch v1 {
        case "":
            var v27 uint8 = 44
            _ = v27
        case "\n\t":
            v12 = v12
        case "a":
            var v28 float64 = (v4 - -2.25)
            _ = v28
        default:
            v4 -= (((v4 * (1e10 - v4)) - ((1e10 - 3.14159) / ((float64(((((-2.25 * 1.5) * v9) / float32(((v4 + (v4 - ((0.1 * v4) / float64((float32(-2.25) + v9))))) + ((float64(v9) / (v4 / v4)) / v4)))) / v9)) * v4) / float64(1e10)))) + v4)
        }
    }
    v4 += ((v4 + v4) * -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
