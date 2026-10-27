package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 3.14159
    var v3 bool = ((!true) && (2147483647 < 32768))
    var v4 int32 = 0
    fmt.Println(v4, v4, v3)
    if v3 {
        switch 42 {
        case 804:
            if ((4294967295 > 2) || v3) {
                var v5 int16 = 42
                for fi6 := 0; fi6 < 12; fi6++ {
                    for fi7 := 0; fi7 < 4; fi7++ {
                        v8 := []int{(fi6 | (fi6 ^ fi7))}
                        var v9 string = "a"
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    _ = fi6
                }
                v10 := GenS1{F0: "hello", F1: 42}
                _ = v10
                _ = v5
            }
        default:
            fmt.Println(v4, v3)
        }
        switch 128 {
        case -288:
            var v11 uint8 = 255
            _ = v11
        default:
            if true {
                v3 = v3
                fmt.Println(v2, v3)
            }
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 int32 = -594
            v14 := []int{fi12, (fi12 / (((fi12 / ((833) | 1))) | 1)), (-(fi12)), (fi12 - (^(fi12)))}
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint16 = 365
                _ = fi15
                _ = v16
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        v2 -= 0.0
        v2 = (-2.25 + (v2 - 0.1))
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := GenS2{F0: v4}
        _ = fi17
        _ = v18
    }
    fmt.Println(v3)
    v19 := GenS3{F0: 32767, F1: 127}
    var v20 float64 = 0.0
    var v21 uint16 = 627
    fmt.Println(v21, v2, v21)
    if (233 != -691) {
        if (-1 >= 42) {
            var v22 uint64 = 0
            _ = v22
        }
    }
    v23 := GenS3{F0: -869, F1: 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
}
