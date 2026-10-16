package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int16; F2 string }
func genFn0() uint64 {
    var v4 uint = 65535
    switch 2147483647 {
    case 127:
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 uint = (v4 | 1)
            _ = fi5
            _ = v6
        }
    case 8:
        v7 := GenS1{F0: "hello", F1: -312}
        _ = v7
    case 219:
        v8 := []int{-766, -2147483648, 128}
        _ = v8
    }
    v4 += v4
    var v9 uint16 = 2
    _ = v4
    _ = v9
    return 287
}
func genFn1() bool {
    var v10 int32 = 1
    v11 := []int{}
    v10 -= (v10 - v10)
    _ = v10
    _ = v11
    return true
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1.5
    v3 := []uint32{}
    var v12 float64 = float64(3.14159)
    v1[-2147483648] = ("hello" + "a")
    for fi13 := 0; fi13 < 7; fi13++ {
        v12 -= ((v12 * ((3.14159 + float64(0.1)) / (1.5 - float64(0.1)))) - 0.0)
        _ = fi13
    }
    fmt.Println(v2)
    switch 31 {
    case -1:
        var v14 int32 = -1
        _ = v14
    case 127:
        v15 := GenS2{F0: 255}
        _ = v15
    case -2147483648:
        v16 := GenS2{F0: 128}
        _ = v16
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 9223372036854775807 {
    case 519:
        if true {
            var v17 int64 = -176
            fmt.Println(v12, v12)
            var v18 int8 = 0
            _ = v17
            _ = v18
        } else {
            switch -9223372036854775808 {
            case 32768:
                v12 = v12
            default:
                var v19 int16 = -1
                _ = v19
            }
        }
    case -2147483648:
        if (!(!false)) {
            fmt.Println(v2, v2)
            for fi20 := 0; fi20 < 11; fi20++ {
                v2 -= float64(1.5)
                fmt.Println(fi20, v2, v12)
                _ = fi20
            }
            var v21 uint16 = 42
            _ = v21
        } else {
            v2 = float64(0.0)
            v2 = v12
            var v22 uint16 = 0
            _ = v22
        }
    }
    if false {
        switch "hello" {
        case "hello":
            var v23 uint64 = 42
            _ = v23
        default:
            v24 := GenS2{F0: 2}
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        v12 -= v2
        _ = fi25
    }
    v26 := GenS3{F0: 32768, F1: -1, F2: "\n\t"}
    v2 = 0.1
    v27 := GenS1{F0: ((("" + "a") + "foo bar") + ""), F1: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
