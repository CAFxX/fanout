package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 string; F2 int64 }
func genFn0(p0 int64) bool {
    switch ((("hello" + ("hello" + "\n\t")) + "hello") + "hello") {
    case "\n\t":
        v6 := GenS1{F0: false, F1: 954}
        _ = v6
    }
    p0 += (p0 | 1)
    var v7 uint = 967
    _ = p0
    _ = v7
    return false
}
func genFn1(p0 int32, p1 int64) int8 {
    fmt.Println(p0, p0)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 26
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]bool)
    var v3 uint16 = 32767
    var v4 uint8 = 235
    v5 := []string{("foo bar" + "foo bar"), "foo bar", "", ("a" + "\n\t")}
    var v8 bool = true
    switch 32767 {
    case 128:
        v9 := GenS1{F0: (0 >= 42), F1: -1}
        _ = v9
    case -2147483648:
        var v10 uint16 = v3
        _ = v10
    }
    _ = len(v5)
    v3 = 582
    v11 := GenS1{F0: (!v8), F1: 880}
    if (!(v4 > v4)) {
        fmt.Println(v8, v4)
        var v12 uint64 = 833
        _ = v12
    }
    switch -1 {
    case 127:
        if v8 {
            switch "\n\t" {
            case "hello":
                var v13 int8 = genFn1(0, 568)
                _ = v13
            default:
                _ = len(v2)
            }
        } else {
            switch "a" {
            case "\n\t":
                var v14 int32 = 0
                _ = v14
            case "":
                _ = len(v1)
            default:
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := GenS2{F0: 0}
                    var v17 int = -2147483648
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            v18 := GenS2{F0: 659}
            _ = v18
        }
    case -646:
        var v19 int32 = 435
        _ = v19
    }
    if ((!v8) && false) {
        v2[-671] = v8
    } else {
        fmt.Println(v4)
        v5 = append(v5, "foo bar")
    }
    v8 = (!v8)
    switch "hello" {
    case "":
        var v20 int8 = 80
        _ = v20
    case "a":
        v21 := GenS1{F0: v8, F1: 2147483647}
        _ = v21
    }
    var v22 uint32 = 2147483648
    v8 = v8
    var v23 int16 = 1000
    v8 = v8
    if (!(("hello" != "a") || v8)) {
        fmt.Println(v22, v22)
    } else {
        v24 := GenS3{F0: 44, F1: "foo bar", F2: 1}
        v23 = v23
        var v25 uint8 = (v4 >> ((255) & 63))
        _ = v24
        _ = v25
    }
    v22 -= (v22 | 2147483648)
    if (837 != 0) {
        var v26 int = 32767
        _ = v26
    } else {
        var v27 float64 = 1e10
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
