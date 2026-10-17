package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 uint8; F1 uint16; F2 int }
type GenS3 struct { F0 int8; F1 uint }
func genFn0(p0 uint32, p1 int8) int8 {
    var v5 bool = (32768 < 127)
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 float32 = 0.1
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return (-(p1))
}

func main() {
    var v1 float32 = 3.14159
    var v2 float64 = ((1e10 / (((float64(float32(3.14159)) + 1e10) - (0.1 / 1.5)) - 0.0)) - float64(v1))
    var v3 uint16 = 681
    var v4 bool = (-2147483648 == -2147483648)
    v8 := GenS1{F0: 1, F1: -32768, F2: 32768}
    var v9 uint64 = 2
    fmt.Println(v4)
    v9 = v9
    fmt.Println(v1, v1)
    var v10 string = ""
    v11 := []int{}
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(v4, v9)
        switch "hello" {
        case "a":
            switch "hello" {
            case "":
                v2 = v2
            case "hello":
                if v4 {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        var v14 float32 = (v1 + (0.1 - (((((v1 - float32(v2)) / v1) - v1) / ((v1 * (0.0 - 0.0)) + (1.5 + (v1 / 0.1)))) / v1)))
                        v2 += -2.25
                        fmt.Println(v1, fi13)
                        _ = fi13
                        _ = v14
                    }
                    v2 = (float64(-2.25) - ((v2 * float64(v1)) + float64(1.5)))
                } else {
                    var v15 int32 = -1
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v17 := GenS2{F0: 217, F1: v3, F2: int(v15)}
                        fmt.Println(fi12, v2)
                        if ((v9 - 0) >= (v9 ^ v9)) {
                            v3 = 32767
                        } else {
                            v18 := GenS1{F0: 710, F1: 0, F2: -326}
                            v4 = true
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            _ = v18
                        }
                        _ = fi16
                        _ = v17
                    }
                    if len(v11) > 0 {
                        v11[0] = fi12
                    }
                    _ = v15
                }
            default:
                switch (v10 + "") {
                case "\n\t":
                    for fi19 := 0; fi19 < 3; fi19++ {
                        v2 = float64(v1)
                        _ = fi19
                    }
                case "foo bar":
                    v20 := GenS3{F0: 1, F1: 65535}
                    _ = v20
                case "hello":
                    _ = len(v11)
                default:
                    v10 = v10
                }
            }
        case "hello":
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS1{F0: 0, F1: -1, F2: (fi21 * (fi21 + fi21))}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                var v23 int64 = 0
                _ = fi21
                _ = v22
                _ = v23
            }
        case "":
            if len(v11) > 0 {
                v11[0] = fi12
            }
        default:
            v2 = v2
        }
        _ = fi12
    }
    v4 = (true || true)
    if v4 {
        var v24 int = 1
        _ = v24
    }
    v11 = append(v11, 72)
    v11 = append(v11, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
