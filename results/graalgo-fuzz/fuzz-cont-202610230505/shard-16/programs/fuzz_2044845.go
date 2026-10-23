package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 bool; F1 int64; F2 int16 }
func genFn0(p0 int8) uint8 {
    var v5 int64 = 1
    var v6 uint64 = 588
    var v7 string = ("foo bar" + "foo bar")
    p0 += 100
    if false {
        v8 := []int{9223372036854775807, 853, 32768, 0}
        v9 := GenS1{F0: (^(v6))}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint, p1 int32, p2 int16) int64 {
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            p1 = (p1 >> ((p0) & 63))
            if true {
                fmt.Println(p0)
                p1 -= (p1 % ((p1) | 1))
            }
            var v12 uint8 = 112
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    var v13 string = ("\n\t" + "")
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 829
}

func main() {
    var v1 float64 = float64(float32(0.0))
    var v2 string = "foo bar"
    var v3 uint16 = 862
    var v4 uint = 42
    v14 := GenS2{F0: 127, F1: -1}
    var v15 float64 = -2.25
    var v16 float64 = float64(1e10)
    if (!true) {
        v17 := []int{90, 0, 438, -344}
        _ = v17
    }
    var v18 uint16 = (v3 - v3)
    v19 := GenS3{F0: true, F1: 42, F2: -394}
    v20 := []int{9223372036854775807, 2147483647}
    var v21 string = (v2 + ("a" + v2))
    var v22 uint16 = v3
    v3 -= v18
    var v23 float64 = float64(3.14159)
    if false {
        if (!(!true)) {
            fmt.Println(v4, v4)
        } else {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            v20 = append(v20, 440)
        }
        switch 9223372036854775807 {
        case -652:
            var v24 uint64 = 127
            _ = v24
        default:
            v4 = v4
        }
    } else {
        var v25 int8 = -1
        _ = v25
    }
    v3 += (v18 | 1)
    v2 = ("hello" + "a")
    fmt.Println(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
