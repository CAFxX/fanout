package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int }
func genFn0() uint8 {
    v4 := []int{}
    switch "foo bar" {
    case "\n\t":
        var v5 int32 = 249
        _ = v5
    case "a":
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint32 = 1
            v7 = 2147483647
            var v8 float32 = 1.5
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: (!(1.5 <= -2.25)), F1: 1}
    _ = v4
    _ = v9
    return 127
}
func genFn1(p0 int, p1 string) int16 {
    p1 = (p1 + p1)
    fmt.Println(p0, p0)
    switch p1 {
    case "hello":
        p1 = "\n\t"
    case "foo bar":
        var v10 uint64 = 1
        _ = v10
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []uint32{394}
    v2 := []int{-531, 4294967295}
    var v3 float64 = 3.14159
    if ("a" >= "foo bar") {
        v3 = v3
    } else {
        switch "" {
        case "hello":
            var v11 int16 = 1
            _ = v11
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if true {
        if len(v1) > 0 {
            v1[0] = 376
        }
    }
    v1 = append(v1, 296)
    var v12 int = 282
    _ = len(v1)
    v3 = (v3 / v3)
    if len(v1) > 0 {
        v1[0] = 1
    }
    var v13 float64 = (v3 / 1.0)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!true) {
        switch (v12 & (v12 >> ((65535) & 63))) {
        case 114:
            v14 := GenS1{F0: ((0.0 * 1e10) >= (-2.25 / 3.14159)), F1: -1}
            _ = v14
        case 32768:
            v15 := GenS1{F0: false, F1: 32767}
            _ = v15
        case 1:
            v16 := GenS2{F0: 1}
            _ = v16
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int16 = 567
            v3 = ((v3 / v13) * (v3 * -2.25))
            _ = fi17
            _ = v18
        }
        var v19 int64 = -521
        _ = v19
    }
    v12 = v12
    var v20 float64 = 1.5
    var v21 uint8 = 255
    v21 = v21
    fmt.Println(v21, v13)
    v20 += v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
