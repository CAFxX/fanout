package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 int) int {
    v7 := []int{(^(p0)), (p0 | 128), p0, p0}
    var v8 uint = 0
    _ = p0
    _ = v7
    _ = v8
    return 779
}
func genFn1(p0 int16, p1 uint, p2 string) int16 {
    fmt.Println(p0, p0, p0)
    var v9 int64 = -405
    p1 += 43
    switch p2 {
    case "a":
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: (v9 ^ (v9 >> (((p1 - p1)) & 63)))}
            var v12 uint16 = 172
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := []int{}
                v12 += 2
                _ = fi13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    case "hello":
        v15 := GenS1{F0: (v9 / ((v9) | 1))}
        _ = v15
    }
    if false {
        var v16 uint8 = 0
        fmt.Println(p1, v16)
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (-(p0))
}

func main() {
    var v1 uint8 = 128
    v2 := []int{-1, -697, 2147483647, 873}
    v3 := make(map[uint32]int)
    var v4 uint16 = 2
    v5 := make(map[string]bool)
    v6 := make(map[uint32]string)
    for fi17 := 0; fi17 < 14; fi17++ {
        v2 = append(v2, (fi17 - (-(fi17))))
        v18 := GenS1{F0: -1}
        if false {
            fmt.Println(v4, fi17)
            var v19 uint8 = v1
            v20 := v3[42]
            _ = v19
            _ = v20
        } else {
            v21 := GenS2{F0: 65535, F1: 1, F2: (42 == 32767)}
            for fi22 := 0; fi22 < 10; fi22++ {
                switch "a" {
                case "\n\t":
                    v6[65535] = ""
                case "hello":
                    v23 := GenS3{F0: false}
                    _ = v23
                case "foo bar":
                    v24 := GenS4{F0: 100}
                    var v25 float64 = (0.0 * float64(float32(-2.25)))
                    if false {
                        var v26 int = -393
                        _ = v26
                    }
                    _ = v24
                    _ = v25
                default:
                    _ = len(v3)
                }
                var v27 int = (^(fi22))
                _ = fi22
                _ = v27
            }
            fmt.Println(v1)
            _ = v21
        }
        _ = fi17
        _ = v18
    }
    if (868 < 42) {
        var v28 int8 = 127
        v28 = v28
        _ = v28
    } else {
        v4 = (v4 / (((v4 % (((v4 ^ 2)) | 1))) | 1))
        for fi29 := 0; fi29 < 12; fi29++ {
            fmt.Println(v4)
            _ = fi29
        }
    }
    var v30 uint16 = v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v31 := GenS2{F0: 364, F1: 504, F2: (genFn0(-858) != 12)}
    if true {
        var v32 int32 = 623
        fmt.Println(v1, v4, v1)
        _ = v32
    }
    var v33 int32 = 120
    var v34 float32 = 0.0
    v30 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
