package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint) uint {
    var v4 int32 = 42
    var v5 int16 = 0
    if false {
        v6 := []int{2147483647, 133, 32767, 784}
        v7 := GenS1{F0: 1, F1: 0, F2: 15}
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    return p0
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]string)
    v3 := []int{}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v8 := GenS1{F0: 2147483647, F1: 1, F2: -33}
    delete(v2, 284)
    v9 := GenS2{F0: -9223372036854775808}
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v1, v1)
        var v11 bool = (true && (false && true))
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 uint64 = 481
            v14 := GenS3{F0: "", F1: 0}
            v11 = v11
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 float32 = 1e10
        switch (-(fi15)) {
        case 652:
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS4{F0: v1}
                v2[2] = "foo bar"
                if ((("foo bar" + ("foo bar" + "a")) <= "") && (fi17 > (fi15 ^ 32768))) {
                    v19 := v2[2147483647]
                    v16 = v16
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
        case 1:
            if true {
                var v20 uint = 655
                _ = v20
            }
        }
        v16 -= float32((1.5 * ((float64(1e10) - float64(((v16 / ((v16 * (v16 - v16)) + (v16 + float32(1e10)))) * v16))) - 0.1)))
        _ = fi15
        _ = v16
    }
    v21 := GenS5{F0: 485}
    v1 -= v1
    switch 0 {
    case -293:
        v1 += v1
    case -352:
        for fi22 := 0; fi22 < 15; fi22++ {
            fmt.Println(fi22, v1)
            v1 -= 490
            _ = fi22
        }
    default:
        fmt.Println(v1)
    }
    if false {
        fmt.Println(v1)
        var v23 int64 = 42
        _ = v23
    } else {
        v1 -= v1
    }
    v24 := GenS6{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v8
    _ = v9
}
