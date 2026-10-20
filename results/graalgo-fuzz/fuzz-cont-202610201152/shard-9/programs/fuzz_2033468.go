package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint16; F1 int8; F2 int }
func genFn0() uint {
    v7 := GenS1{F0: 2, F1: -35}
    var v8 uint = 2
    fmt.Println(v8)
    _ = v7
    _ = v8
    return (v8 | 0)
}
func genFn1(p0 uint16) uint {
    v9 := []int{-467, -599, 32768}
    if true {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    } else {
        fmt.Println(p0)
        for fi10 := 0; fi10 < 16; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                if (((-2.25 + (((0.1 + 0.1) - 3.14159) * (3.14159 + 1e10))) / (3.14159 - (((1.5 + 3.14159) * 1e10) + -2.25))) > 1e10) {
                    p0 -= p0
                    p0 = (p0 / (((^(p0))) | 1))
                } else {
                    if len(v9) > 0 {
                        v9[0] = fi11
                    }
                    _ = len(v9)
                    var v12 int32 = -2147483648
                    _ = v12
                }
                p0 -= p0
                _ = fi11
            }
            _ = fi10
        }
    }
    var v13 int16 = 0
    var v14 uint16 = (p0 + (p0 | 42))
    var v15 int64 = 0
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    _ = v9
    return 511
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 int64 = -1
    v5 := make(map[string]string)
    var v6 int32 = v1
    var v16 uint32 = 240
    fmt.Println(v1, v1)
    if false {
        v3[362] = 65535
        fmt.Println(v6, v4, v16)
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            if (0 == 119) {
                fmt.Println(v1)
                var v19 int64 = (v4 % ((int64(v1)) | 1))
                _ = len(v5)
                _ = v19
            } else {
                v20 := GenS1{F0: 65535, F1: 59}
                _ = v20
            }
            var v21 int16 = 614
            switch ("hello" + "") {
            case "foo bar":
                v22 := GenS2{F0: -32768, F1: 74, F2: 127}
                _ = v22
            default:
                v23 := GenS2{F0: v21, F1: 128, F2: 191}
                _ = v23
            }
            _ = fi18
            _ = v21
        }
        if false {
            v16 = (v16 | 2147483648)
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int16 = 1
                _ = fi24
                _ = v25
            }
            fmt.Println(v6, v1, v4)
        } else {
            switch "hello" {
            case "foo bar":
                v26 := []int{fi17, (fi17 % ((9223372036854775807) | 1)), fi17}
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 19; fi27++ {
            var v28 int = int(v16)
            v6 = 42
            var v29 uint8 = 42
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi17
    }
    v30 := []int{}
    v31 := GenS1{F0: 1, F1: 127}
    if len(v30) > 0 {
        v30[0] = 128
    }
    v4 = (v4 + v4)
    v32 := GenS3{F0: (!(2147483648 >= 0)), F1: v6}
    var v33 float32 = (-2.25 + (((float32(float64(float32((((1.5 + 3.14159) - 0.1) * -2.25)))) / 3.14159) / float32(-2.25)) / 1.0))
    for fi34 := 0; fi34 < 5; fi34++ {
        v4 = (v4 % (((v4 | (v4 | v4))) | 1))
        if ((v33 - v33) != v33) {
            v35 := GenS1{F0: (v16 + v16), F1: 127}
            v36 := GenS4{F0: 32767, F1: 127, F2: (^(fi34))}
            fmt.Println(v4, v4, v33)
            _ = v35
            _ = v36
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
