package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
func genFn0(p0 uint8, p1 uint) int8 {
    if (!(!false)) {
        if (-1 < -1) {
            p0 -= (^(p0))
            var v5 int = -9223372036854775808
            for fi6 := 0; fi6 < 15; fi6++ {
                v7 := []int{(fi6 * 9223372036854775807), fi6}
                _ = len(v7)
                _ = fi6
                _ = v7
            }
            _ = v5
        } else {
            var v8 int = 127
            v9 := GenS1{F0: 32767}
            _ = v8
            _ = v9
        }
    }
    fmt.Println(p1)
    for fi10 := 0; fi10 < 12; fi10++ {
        switch fi10 {
        case 32768:
            v11 := GenS1{F0: 588}
            _ = v11
        }
        p1 = p1
        _ = fi10
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 int16 = 0
    v2 := []int{9223372036854775807, 0}
    v3 := make(map[string]int)
    var v4 bool = true
    fmt.Println(v1)
    switch 127 {
    case -2147483648:
        var v12 string = (("a" + "a") + "\n\t")
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v4 = v4
        _ = fi13
    }
    var v14 int = -601
    v15 := GenS2{F0: 42, F1: v14, F2: 42}
    switch "foo bar" {
    case "\n\t":
        if true {
            switch v14 {
            case 128:
                fmt.Println(v1, v4)
            case 2147483647:
                var v16 float64 = 0.1
                _ = v16
            default:
                var v17 float32 = 1e10
                _ = v17
            }
            switch "\n\t" {
            case "a":
                v18 := GenS3{F0: v14, F1: (-(v14))}
                _ = v18
            }
            var v19 float32 = 1.5
            _ = v19
        }
    case "hello":
        var v20 float64 = -2.25
        _ = v20
    }
    v4 = (18446744073709551615 < 4294967295)
    var v21 uint = 4294967295
    fmt.Println(v21)
    v22 := v3[((("" + (("foo bar" + "a") + "a")) + "foo bar") + (("hello" + "a") + "\n\t"))]
    var v23 uint = (^(v21))
    var v24 bool = (0 != 2147483647)
    fmt.Println(v24, v21)
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := v3[""]
        var v27 float32 = 1.5
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        v4 = v4
        if (v4 && (65 <= 128)) {
            fmt.Println(v1, fi28, fi28)
            if len(v2) > 0 {
                v2[0] = (^(fi28))
            }
            if len(v2) > 0 {
                v2[0] = -1
            }
        } else {
            v29 := GenS4{F0: 386}
            switch "foo bar" {
            case "":
                var v30 float32 = 1e10
                _ = v30
            }
            v31 := GenS5{F0: (v14 & 0)}
            _ = v29
            _ = v31
        }
        var v32 uint32 = 2147483647
        _ = fi28
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
