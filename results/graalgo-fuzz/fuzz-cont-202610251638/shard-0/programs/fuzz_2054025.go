package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint) int8 {
    var v4 float64 = 0.0
    var v5 string = ""
    v5 = ("a" + ("foo bar" + ""))
    v6 := GenS1{F0: 0, F1: 1000, F2: false}
    if ((false && (737 == -9223372036854775808)) || false) {
        var v7 uint8 = 127
        v5 = v5
        _ = v7
    } else {
        var v8 float64 = -2.25
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1() string {
    var v9 string = "\n\t"
    if false {
        var v10 uint16 = 42
        _ = v10
    }
    _ = v9
    return "foo bar"
}

func main() {
    var v1 float32 = float32(-2.25)
    var v2 int8 = -71
    v3 := make(map[uint32]bool)
    switch 32768 {
    case 4294967295:
        var v11 int64 = 42
        _ = v11
    case 42:
        switch 2147483647 {
        case 722:
            switch -2147483648 {
            case 128:
                var v12 bool = false
                _ = v12
            case -1:
                v1 -= 3.14159
            case -443:
                v3[2147483647] = (true && false)
            default:
                fmt.Println(v2, v2)
            }
        }
    case 418:
        if true {
            for fi13 := 0; fi13 < 1; fi13++ {
                v1 = (1.5 + (v1 / v1))
                switch fi13 {
                case 1:
                    v14 := GenS2{F0: 42}
                    _ = v14
                }
                _ = fi13
            }
        } else {
            v2 = v2
            var v15 bool = (1.5 < 1.5)
            var v16 int16 = -367
            _ = v15
            _ = v16
        }
    }
    v17 := GenS2{F0: 42}
    v2 = (-(v2))
    v2 = v2
    var v18 int32 = 42
    v19 := []int{42, -1, 32768}
    var v20 uint8 = 81
    var v21 float64 = (0.0 * 1e10)
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 int32 = (v18 & (v18 >> ((42) & 63)))
        for fi24 := 0; fi24 < 13; fi24++ {
            if (-1 > 32767) {
                var v25 int64 = -280
                var v26 float64 = 0.0
                var v27 float32 = 1.5
                _ = v25
                _ = v26
                _ = v27
            }
            var v28 uint8 = (v20 ^ v20)
            delete(v3, 2)
            _ = fi24
            _ = v28
        }
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := GenS3{F0: 879}
            _ = fi29
            _ = v30
        }
        _ = fi22
        _ = v23
    }
    if len(v19) > 0 {
        v19[0] = 32768
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
