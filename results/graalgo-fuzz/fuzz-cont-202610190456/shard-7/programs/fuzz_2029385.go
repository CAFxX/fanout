package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 string }
type GenS4 struct { F0 uint; F1 string; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int32 }
func genFn0(p0 uint64) bool {
    for fi6 := 0; fi6 < 13; fi6++ {
        p0 = (p0 ^ p0)
        fmt.Println(p0)
        var v7 uint16 = 32767
        _ = fi6
        _ = v7
    }
    p0 -= 94
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint = 65535
        _ = fi8
        _ = v9
    }
    _ = p0
    return ((p0 & p0) < (p0 + p0))
}
func genFn1(p0 bool, p1 uint8, p2 uint32) uint8 {
    var v10 uint64 = 702
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS1{F0: (p1 / ((1) | 1))}
        _ = fi11
        _ = v12
    }
    var v13 int = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    return (^(p1))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    var v4 uint = 4294967295
    v5 := make(map[string]int)
    if false {
        v4 += (v4 % ((341) | 1))
        v14 := []int{-1000}
        _ = v14
    }
    v15 := GenS2{F0: 778, F1: 552}
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int64 = 9223372036854775807
        _ = fi16
        _ = v17
    }
    if (false && true) {
        switch 351 {
        case 128:
            var v18 uint8 = 255
            _ = v18
        case 1:
            switch -643 {
            case -9223372036854775808:
                var v19 uint8 = 42
                _ = v19
            }
        case -434:
            v4 = (^(v4))
        }
        var v20 float64 = (0.1 * 3.14159)
        _ = v20
    } else {
        v21 := GenS1{F0: 128}
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := GenS3{F0: 1, F1: 69, F2: ("hello" + "a")}
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := v3[""]
            for fi26 := 0; fi26 < 8; fi26++ {
                fmt.Println(fi24, fi26, v4)
                if true {
                    v27 := GenS4{F0: v4, F1: "hello", F2: 100}
                    var v28 float32 = 0.1
                    _ = len(v2)
                    fmt.Println(fi24, fi24)
                    var v29 uint64 = uint64(v4)
                    _ = v27
                    _ = v28
                    _ = v29
                }
                v25 = fi24
                _ = fi26
            }
            v30 := GenS1{F0: 52}
            _ = fi24
            _ = v25
            _ = v30
        }
        _ = v21
    }
    var v31 float64 = (1.5 * -2.25)
    v32 := GenS4{F0: v4, F1: ("foo bar" + ("hello" + ("hello" + ""))), F2: 126}
    fmt.Println(v4, v31)
    v31 = (v31 * float64((3.14159 / 1.5)))
    var v33 uint16 = 255
    v4 = v4
    v33 += (v33 | (v33 + uint16(v4)))
    if true {
        v4 += v4
        v34 := []int{}
        v35 := GenS5{F0: -32768}
        _ = v34
        _ = v35
    } else {
        v4 = v4
    }
    if (true && false) {
        var v36 uint32 = 806
        _ = v36
    } else {
        fmt.Println(v4)
        v37 := GenS4{F0: v4, F1: "a", F2: -11}
        switch 0 {
        case 4294967295:
            v38 := GenS6{F0: 404, F1: 674}
            _ = v38
        }
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
