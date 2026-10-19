package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint; F1 int8; F2 uint16 }
func genFn0(p0 int, p1 int64, p2 uint32) int64 {
    v5 := GenS1{F0: 42, F1: false}
    p1 = 9223372036854775807
    var v6 float32 = 0.0
    v7 := []int{p0, p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 944
}
func genFn1(p0 uint64) int8 {
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint16 = 65535
        p0 = p0
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint32 = 65535
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v12 := []int{}
    var v13 int16 = -32768
    v13 = (v13 % ((-471) | 1))
    _ = p0
    _ = v12
    _ = v13
    return 59
}

func main() {
    var v1 int16 = -464
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 87
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(fi14)
        fmt.Println(fi14, v1)
        _ = fi14
    }
    var v15 bool = false
    v4 = 255
    fmt.Println(v15, v15, v15)
    v16 := GenS1{F0: 1, F1: (1e10 == (0.0 / 1.0))}
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 bool = true
        if ((!(2 > 438)) && (v18 || v15)) {
            v19 := []int{(fi17 << ((0) & 63)), fi17, -399}
            _ = v19
        } else {
            var v20 string = "foo bar"
            v21 := GenS1{F0: 0, F1: false}
            _ = v20
            _ = v21
        }
        _ = fi17
        _ = v18
    }
    v1 = 42
    if true {
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS2{F0: 1, F1: (v4 << ((uint(v4)) & 63)), F2: 0}
            _ = fi22
            _ = v23
        }
        fmt.Println(v1)
        for fi24 := 0; fi24 < 20; fi24++ {
            for fi25 := 0; fi25 < 18; fi25++ {
                fmt.Println(fi24)
                var v26 bool = (!(genFn0(128, 42, 928) != 1))
                _ = fi25
                _ = v26
            }
            var v27 uint = 1
            v28 := GenS2{F0: 18446744073709551615, F1: (v4 & 52), F2: v1}
            _ = fi24
            _ = v27
            _ = v28
        }
    }
    switch 32767 {
    case 9223372036854775807:
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 uint16 = 1
            _ = fi29
            _ = v30
        }
    case -1:
        for fi31 := 0; fi31 < 20; fi31++ {
            v15 = (-248 < v1)
            v32 := GenS1{F0: 0, F1: true}
            _ = fi31
            _ = v32
        }
    case -783:
        v33 := GenS3{F0: 2, F1: 1}
        _ = v33
    default:
        v4 = (^(v4))
    }
    fmt.Println(v15, v1, v15)
    v1 += v1
    var v34 float64 = float64(-2.25)
    for fi35 := 0; fi35 < 2; fi35++ {
        var v36 bool = true
        v37 := GenS1{F0: 2147483647, F1: (0.0 > 1e10)}
        _ = fi35
        _ = v36
        _ = v37
    }
    if v15 {
        switch -630 {
        case 32767:
            for fi38 := 0; fi38 < 3; fi38++ {
                v34 = v34
                v39 := GenS4{F0: 0, F1: -1, F2: 32767}
                switch fi38 {
                case 128:
                    var v40 float32 = (1.5 * 0.0)
                    _ = v40
                default:
                    var v41 uint32 = uint32(v4)
                    _ = v41
                }
                _ = fi38
                _ = v39
            }
        }
        var v42 float64 = (v34 / float64((-2.25 / ((0.1 * 3.14159) + 0.1))))
        _ = v42
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
