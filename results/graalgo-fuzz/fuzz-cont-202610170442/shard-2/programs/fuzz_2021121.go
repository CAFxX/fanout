package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
func genFn0(p0 int32, p1 int32) int64 {
    v6 := []int{}
    v7 := GenS1{F0: 9223372036854775807, F1: "hello", F2: 0}
    p0 += (^(p0))
    v8 := GenS1{F0: 95, F1: "foo bar", F2: 1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int, p1 bool) int32 {
    p0 = 128
    p1 = (58 >= 2)
    if (p1 && p1) {
        var v9 int16 = -1
        _ = v9
    } else {
        switch p0 {
        case 694:
            fmt.Println(p0, p0)
        default:
            v10 := GenS1{F0: 1, F1: ("\n\t" + ("\n\t" + ("\n\t" + ("" + ("\n\t" + "hello"))))), F2: 128}
            _ = v10
        }
    }
    switch p0 {
    case -789:
        var v11 uint = 575
        _ = v11
    case 575:
        if true {
            if (0 >= 2147483647) {
                switch (p0 + p0) {
                case -9223372036854775808:
                    fmt.Println(p1, p0)
                default:
                    var v12 uint8 = 0
                    _ = v12
                }
            } else {
                v13 := []int{(-(p0)), 9223372036854775807}
                var v14 int8 = -1
                _ = v13
                _ = v14
            }
            p0 -= p0
        }
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    v1 := []int{916, -836}
    v2 := make(map[string]bool)
    var v3 bool = false
    var v4 float64 = (-2.25 - float64(1.5))
    v5 := make(map[string]int)
    fmt.Println(v3, v4)
    var v15 bool = (!(0 >= 2))
    v16 := GenS1{F0: 0, F1: "foo bar", F2: 1}
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS2{F0: -1}
        _ = fi17
        _ = v18
    }
    v19 := GenS3{F0: "hello"}
    fmt.Println(v15, v4, v15)
    v20 := GenS2{F0: 887}
    v15 = true
    v21 := GenS3{F0: "a"}
    switch 1 {
    case 0:
        fmt.Println(v4)
    case 127:
        for fi22 := 0; fi22 < 13; fi22++ {
            v4 = v4
            _ = fi22
        }
    default:
        fmt.Println(v3, v15, v3)
    }
    delete(v2, ("hello" + (("\n\t" + "a") + "")))
    var v23 uint = 182
    v23 = v23
    v23 += (v23 - v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
