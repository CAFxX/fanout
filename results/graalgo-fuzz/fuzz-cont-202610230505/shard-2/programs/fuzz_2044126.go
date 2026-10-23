package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint }
func genFn0() int64 {
    var v6 int16 = 1
    var v7 int = 32768
    _ = v6
    _ = v7
    return int64(v7)
}
func genFn1(p0 int) int32 {
    var v8 uint = 0
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v8, p0, fi9)
        for fi10 := 0; fi10 < 9; fi10++ {
            fmt.Println(p0, p0)
            var v11 float64 = 1e10
            _ = fi10
            _ = v11
        }
        v12 := []int{fi9, (p0 & fi9), p0, 2147483647}
        _ = fi9
        _ = v12
    }
    if (65535 >= 617) {
        v13 := []int{p0}
        _ = v13
    } else {
        if false {
            v14 := []int{p0, 128, p0}
            for fi15 := 0; fi15 < 14; fi15++ {
                for fi16 := 0; fi16 < 8; fi16++ {
                    switch fi16 {
                    case 0:
                        p0 -= -816
                    case -178:
                        v8 = (^(v8))
                    case -125:
                        var v17 float64 = (0.0 + ((3.14159 / 1.0) * (0.0 - -2.25)))
                        _ = v17
                    default:
                        v8 = v8
                    }
                    switch fi16 {
                    case 32767:
                        v18 := GenS1{F0: (v8 | (v8 << ((v8) & 63))), F1: fi15}
                        _ = v18
                    case -2147483648:
                        _ = len(v14)
                    case 443:
                        var v19 float32 = -2.25
                        _ = v19
                    default:
                        v20 := GenS2{F0: 339}
                        _ = v20
                    }
                    p0 = -9223372036854775808
                    _ = fi16
                }
                if len(v14) > 0 {
                    v14[0] = (fi15 >> ((v8) & 63))
                }
                _ = fi15
            }
            _ = v14
        }
        var v21 bool = false
        _ = v21
    }
    _ = p0
    _ = v8
    return 937
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 577
    v3 := make(map[uint32]bool)
    var v4 float64 = -2.25
    var v5 float64 = (v4 + ((v4 * ((v4 / v4) + -2.25)) / (v4 + float64(0.0))))
    v4 = 1e10
    v22 := v1["hello"]
    v4 = v4
    v23 := GenS2{F0: -230}
    if ((-2147483648 <= 146) && (!true)) {
        v24 := []int{v22, 0}
        _ = v24
    }
    if true {
        v25 := GenS3{F0: 0}
        v4 = 1e10
        v26 := GenS4{F0: (v22 * v22), F1: v22}
        _ = v25
        _ = v26
    } else {
        v27 := GenS3{F0: 42}
        var v28 uint = 65535
        v29 := []int{(v22 | v22)}
        _ = v27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := GenS5{F0: 18446744073709551615}
        _ = fi30
        _ = v31
    }
    v5 = 0.1
    _ = len(v3)
    var v32 float32 = -2.25
    switch "a" {
    case "a":
        var v33 uint16 = 139
        _ = v33
    case "hello":
        v34 := GenS1{F0: 255, F1: v22}
        _ = v34
    case "foo bar":
        fmt.Println(v5)
    }
    v5 = ((v5 - -2.25) * (float64(v32) / 3.14159))
    v5 = float64(((1e10 - v32) - float32((v4 / float64((float32(((v5 - v5) + 0.1)) - (v32 * v32)))))))
    for fi35 := 0; fi35 < 7; fi35++ {
        fmt.Println(v4, v4)
        _ = fi35
    }
    var v36 int16 = -517
    fmt.Println(v2, v5, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
}
