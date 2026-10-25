package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16; F1 int16; F2 bool }
func genFn0(p0 uint64, p1 int16, p2 int) int32 {
    var v6 uint16 = 65535
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 742
}
func genFn1(p0 uint32, p1 int32) int64 {
    var v7 float64 = float64((0.1 + (0.1 / (0.1 - 0.0))))
    v8 := []int{0, 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -642
}

func main() {
    var v1 string = "foo bar"
    var v2 bool = true
    v3 := make(map[uint32]int)
    var v4 int8 = -1
    var v5 float32 = 0.0
    v1 = v1
    var v9 uint16 = 32767
    for fi10 := 0; fi10 < 17; fi10++ {
        switch v1 {
        case "hello":
            v4 = 127
        default:
            v2 = (3 != 94)
        }
        _ = fi10
    }
    v1 = v1
    v5 += v5
    v2 = (!(v2 && (!false)))
    var v11 uint64 = 481
    v2 = (128 <= 94)
    if (v2 && ((v4 + -81) != 102)) {
        v12 := GenS1{F0: (v9 << ((224) & 63))}
        for fi13 := 0; fi13 < 9; fi13++ {
            if false {
                v4 = -20
                v14 := GenS2{F0: 4294967295, F1: v4, F2: 1}
                for fi15 := 0; fi15 < 10; fi15++ {
                    switch fi13 {
                    case -83:
                        var v16 int64 = 192
                        _ = v16
                    default:
                        v5 -= v5
                    }
                    _ = fi15
                }
                _ = v14
            } else {
                v4 = 0
                var v17 float64 = float64(v5)
                var v18 int32 = 42
                _ = v17
                _ = v18
            }
            if false {
                fmt.Println(v11, v4)
                var v19 string = "\n\t"
                for fi20 := 0; fi20 < 1; fi20++ {
                    var v21 int32 = 15
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                v22 := GenS1{F0: (v9 + v9)}
                _ = v22
            }
            _ = fi13
        }
        if (((v11 + (v11 - 1)) == uint64(v9)) && v2) {
            v23 := GenS3{F0: v9, F1: 0, F2: v2}
            _ = v23
        } else {
            for fi24 := 0; fi24 < 14; fi24++ {
                switch (^(fi24)) {
                case 32767:
                    v1 = v1
                default:
                    _ = len(v3)
                }
                _ = fi24
            }
            fmt.Println(v1)
            fmt.Println(v11, v9)
        }
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
