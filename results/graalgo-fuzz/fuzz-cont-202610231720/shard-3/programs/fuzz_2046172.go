package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32; F1 string; F2 int }
type GenS3 struct { F0 int16; F1 bool }
func genFn0(p0 int, p1 string) uint32 {
    var v5 int16 = -1
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(p0)
        if false {
            switch (p0 << ((255) & 63)) {
            case 4294967295:
                v7 := GenS1{F0: -128, F1: false}
                _ = v7
            }
            v8 := []int{fi6, (p0 << ((724) & 63)), 42, p0}
            _ = v8
        } else {
            v5 = 1
        }
        _ = fi6
    }
    fmt.Println(p0)
    if false {
        var v9 int64 = 0
        var v10 int64 = (v9 << ((2) & 63))
        v10 += (v10 / ((42) | 1))
        _ = v10
        _ = v9
    } else {
        if (127 == 0) {
            if (1.5 > 0.0) {
                if ((p1 + "hello") >= "\n\t") {
                    fmt.Println(p1, v5, p0)
                    var v11 bool = true
                    _ = v11
                }
            }
            var v12 bool = true
            _ = v12
        }
        if false {
            v13 := []int{p0, 32767, -845}
            _ = v13
        }
        v14 := GenS1{F0: 121, F1: true}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1() int16 {
    for fi15 := 0; fi15 < 18; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    var v16 float32 = -2.25
    v16 += (((v16 - v16) - 1e10) / (v16 + v16))
    if false {
        v17 := GenS1{F0: -6, F1: (!(128 > 127))}
        fmt.Println(v16)
        _ = v17
    }
    _ = v16
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := []uint32{4294967295, 994}
    v4 := make(map[int]int)
    var v18 uint16 = 32767
    if len(v3) > 0 {
        v3[0] = 66
    }
    var v19 uint32 = 2
    var v20 float64 = -2.25
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    fmt.Println(v20)
    switch 9223372036854775807 {
    case 0:
        if false {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v20)
            case "hello":
                v21 := GenS1{F0: -128, F1: ((((float32(v20) / -2.25) - 0.0) / 0.1) == 3.14159)}
                _ = v21
            default:
                v18 -= (v18 ^ 42)
            }
            v22 := GenS1{F0: 100, F1: true}
            v23 := GenS2{F0: 42, F1: "foo bar", F2: 32768}
            _ = v22
            _ = v23
        }
    default:
        var v24 uint16 = (v18 / ((v18) | 1))
        _ = v24
    }
    v25 := GenS3{F0: 42, F1: (!(0 > 2))}
    v20 = (((v20 + (v20 + v20)) + v20) - (((0.1 / float64(-2.25)) + v20) / 1.5))
    v18 += 2
    if false {
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 string = "foo bar"
            switch int(v18) {
            case 42:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case -995:
                v27 = v27
            default:
                var v28 uint = 0
                _ = v28
            }
            _ = fi26
            _ = v27
        }
    }
    v18 -= 361
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
}
