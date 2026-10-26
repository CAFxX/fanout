package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint }
func genFn0(p0 int16, p1 int32, p2 int64) string {
    if (((p0 % ((p0) | 1)) > p0) && true) {
        var v7 int64 = 9223372036854775807
        _ = v7
    }
    fmt.Println(p2, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = ((-9223372036854775808 < 42) && true)
    var v3 int16 = 42
    var v4 bool = true
    var v5 int32 = -1
    var v6 int8 = 100
    var v8 string = "hello"
    v9 := []int{}
    _ = len(v9)
    v6 = v6
    if true {
        v6 -= v6
        for fi10 := 0; fi10 < 18; fi10++ {
            switch fi10 {
            case 1:
                var v11 float32 = (((0.1 * 1e10) - (((((float32(1.5) + ((3.14159 * (((0.0 * 1.5) + 0.1) / ((1e10 - 0.1) / (float32((1.5 * 1e10)) / -2.25)))) / ((0.0 + 1e10) + ((-2.25 - 0.1) - 0.0)))) * (float32(0.1) * 1.5)) / 1e10) + 0.1) - 1.5)) - float32(((1e10 + 3.14159) * -2.25)))
                _ = v11
            case 127:
                var v12 uint8 = 255
                _ = v12
            case 32768:
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := GenS1{F0: v5}
                    var v15 float64 = (-2.25 * ((3.14159 - float64(1.5)) / 0.1))
                    v16 := GenS1{F0: 958}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            default:
                _ = len(v1)
            }
            if v4 {
                v6 += (v6 + (v6 % (((v6 / ((v6) | 1))) | 1)))
                v3 = 67
                if v4 {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        _ = fi17
                    }
                    _ = len(v1)
                    var v18 uint64 = 1
                    _ = v18
                }
            } else {
                delete(v1, v8)
                for fi19 := 0; fi19 < 10; fi19++ {
                    switch v8 {
                    case "":
                        v20 := GenS2{F0: 1, F1: 0, F2: 1}
                        _ = v20
                    case "a":
                        fmt.Println(fi19, v6)
                    }
                    v6 = (v6 | -128)
                    var v21 bool = (!(false || (!(280 != -404))))
                    _ = fi19
                    _ = v21
                }
                fmt.Println(v8)
            }
            _ = fi10
        }
    } else {
        v5 -= (v5 >> ((65535) & 63))
    }
    v22 := GenS2{F0: 65535, F1: 1, F2: 1}
    v23 := GenS3{F0: 1, F1: 42}
    for fi24 := 0; fi24 < 8; fi24++ {
        fmt.Println(v8)
        var v25 int16 = v3
        v5 = 0
        _ = fi24
        _ = v25
    }
    _ = len(v9)
    for fi26 := 0; fi26 < 20; fi26++ {
        for fi27 := 0; fi27 < 7; fi27++ {
            if v2 {
                fmt.Println(v2)
            }
            var v28 int32 = v5
            _ = fi27
            _ = v28
        }
        v1[v8] = (-(fi26))
        _ = fi26
    }
    v29 := GenS1{F0: v5}
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS4{F0: 95}
        var v32 string = v8
        _ = fi30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        v3 = v3
        _ = fi33
    }
    switch -9223372036854775808 {
    case 32767:
        var v34 uint64 = 554
        _ = v34
    case -3:
        for fi35 := 0; fi35 < 12; fi35++ {
            v3 -= -269
            fmt.Println(v4, fi35)
            _ = fi35
        }
    }
    _ = len(v9)
    v9 = append(v9, -507)
    var v36 uint16 = 846
    v37 := GenS2{F0: 0, F1: v36, F2: -784}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
