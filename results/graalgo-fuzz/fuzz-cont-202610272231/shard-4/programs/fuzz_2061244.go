package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int; F1 int32; F2 int8 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 uint; F1 int64; F2 uint }
func genFn0() bool {
    var v7 float64 = ((3.14159 / 1.5) + 0.0)
    var v8 bool = true
    _ = v7
    _ = v8
    return (true && false)
}
func genFn1(p0 uint8) int32 {
    for fi9 := 0; fi9 < 5; fi9++ {
        switch -768 {
        case 810:
            p0 -= 42
        default:
            if true {
                v10 := []int{fi9, (fi9 * fi9)}
                var v11 uint32 = 689
                var v12 int16 = -800
                _ = v10
                _ = v11
                _ = v12
            } else {
                fmt.Println(fi9, fi9)
            }
        }
        v13 := GenS1{F0: 0, F1: 805}
        _ = fi9
        _ = v13
    }
    if (2 < 587) {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int16 = 42
            p0 += p0
            _ = fi14
            _ = v15
        }
        switch -238 {
        case -617:
            for fi16 := 0; fi16 < 3; fi16++ {
                p0 -= (p0 & p0)
                _ = fi16
            }
        case -1:
            if (128 == 127) {
                switch "\n\t" {
                case "foo bar":
                    p0 -= p0
                case "a":
                    p0 += (p0 & p0)
                case "hello":
                    v17 := GenS2{F0: "foo bar", F1: "a"}
                    _ = v17
                }
                fmt.Println(p0)
            }
        default:
            v18 := GenS3{F0: 9223372036854775807, F1: 0, F2: 1}
            _ = v18
        }
    } else {
        switch "\n\t" {
        case "foo bar":
            if false {
                p0 += p0
                if false {
                    var v19 uint = 0
                    _ = v19
                } else {
                    var v20 bool = false
                    var v21 int = 32768
                    p0 -= p0
                    _ = v20
                    _ = v21
                }
            }
        case "":
            var v22 bool = true
            _ = v22
        }
    }
    _ = p0
    return 2147483647
}

func main() {
    v1 := []string{"a", "foo bar", (("\n\t" + ("foo bar" + "hello")) + "hello")}
    v2 := make(map[uint32]bool)
    v3 := []string{"foo bar"}
    var v4 uint32 = 2147483648
    var v5 uint64 = 2
    var v6 uint = uint(v4)
    _ = len(v2)
    v23 := GenS4{F0: 0, F1: 42}
    if (v5 > v5) {
        v24 := GenS5{F0: 118, F1: 42, F2: v6}
        _ = v24
    }
    var v25 uint8 = 43
    var v26 string = "hello"
    var v27 float64 = 0.0
    for fi28 := 0; fi28 < 6; fi28++ {
        fmt.Println(v5, v25, v26)
        fmt.Println(v5, v26, v27)
        v29 := GenS1{F0: 0, F1: -296}
        _ = fi28
        _ = v29
    }
    var v30 uint16 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
