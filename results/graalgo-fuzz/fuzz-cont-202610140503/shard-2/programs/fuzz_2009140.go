package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 string) uint {
    switch p0 {
    case "foo bar":
        var v4 bool = (false || true)
        _ = v4
    case "":
        for fi5 := 0; fi5 < 10; fi5++ {
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 int = 894
                v7 = v7
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    case "hello":
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: 42, F1: 255, F2: 42}
            if false {
                p0 = "a"
                var v10 uint32 = 649
                _ = v10
            }
            _ = fi8
            _ = v9
        }
    }
    v11 := []int{0, 127, 9223372036854775807}
    _ = p0
    _ = v11
    return 560
}

func main() {
    var v1 bool = false
    var v2 uint16 = 835
    var v3 int8 = -102
    var v12 uint = 29
    var v13 uint64 = 2
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int8 = (v3 % ((0) | 1))
        v16 := GenS1{F0: 1, F1: v2, F2: 1}
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        if false {
            var v17 int16 = 42
            switch "a" {
            case "":
                if ((!((v13 * (v13 - (v13 + v13))) == 28)) && ((-(v17)) >= v17)) {
                    var v18 int8 = (^(v3))
                    var v19 uint16 = (^(v2))
                    _ = v18
                    _ = v19
                }
            case "foo bar":
                var v20 uint16 = (v2 % (((v2 / (((^(v2))) | 1))) | 1))
                _ = v20
            default:
                v2 += 0
            }
            for fi21 := 0; fi21 < 15; fi21++ {
                switch 1 {
                case -9223372036854775808:
                    switch fi21 {
                    case 1:
                        fmt.Println(v1)
                    case 128:
                        v22 := GenS2{F0: v13}
                        _ = v22
                    }
                case 4294967295:
                    v17 -= (v17 ^ int16(fi21))
                default:
                    v23 := GenS1{F0: 2147483647, F1: (v2 >> (((v12 - genFn0("a"))) & 63)), F2: 128}
                    _ = v23
                }
                var v24 bool = true
                fmt.Println(v1)
                _ = fi21
                _ = v24
            }
            _ = v17
        }
    } else {
        fmt.Println(v2)
        var v25 uint64 = v13
        _ = v25
    }
    v12 += (v12 / ((2) | 1))
    v13 = uint64(v12)
    v26 := GenS1{F0: 740, F1: 713, F2: 42}
    if true {
        v27 := GenS3{F0: v2}
        _ = v27
    }
    var v28 float64 = float64(0.1)
    for fi29 := 0; fi29 < 2; fi29++ {
        v28 -= float64(1.5)
        var v30 float32 = 0.0
        v31 := GenS3{F0: v2}
        _ = fi29
        _ = v30
        _ = v31
    }
    if true {
        v32 := []int{-539}
        fmt.Println(v13)
        fmt.Println(v1, v1, v3)
        _ = v32
    } else {
        v33 := GenS1{F0: -158, F1: (v2 & v2), F2: 42}
        _ = v33
    }
    v3 += 100
    if v1 {
        v12 = 824
        var v34 uint16 = (v2 ^ (v2 + 42))
        _ = v34
    } else {
        var v35 int = 9223372036854775807
        fmt.Println(v35, v35)
        _ = v35
    }
    v3 = int8(v13)
    v3 = v3
    v36 := []int{153}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v36
}
