package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 int32, p1 string) int8 {
    if (100 > 46) {
        var v6 int64 = 625
        fmt.Println(p1, p1)
        var v7 uint = 1
        _ = v6
        _ = v7
    }
    p1 = p1
    switch p1 {
    case "a":
        v8 := []int{32767, 32767}
        _ = v8
    default:
        v9 := []int{-916, 128, -956}
        _ = v9
    }
    var v10 bool = (192 <= 93)
    for fi11 := 0; fi11 < 20; fi11++ {
        p1 = p1
        v12 := GenS1{F0: 0}
        for fi13 := 0; fi13 < 1; fi13++ {
            if (1.5 <= -2.25) {
                switch (-(fi11)) {
                case -518:
                    var v14 uint16 = 0
                    _ = v14
                default:
                    switch (fi13 - (fi13 - fi13)) {
                    case 4294967295:
                        v15 := []int{-735, -9223372036854775808, fi11, 9223372036854775807}
                        _ = v15
                    case 32768:
                        var v16 int32 = (p0 | p0)
                        _ = v16
                    case 529:
                        var v17 int8 = 92
                        _ = v17
                    }
                }
            } else {
                switch fi11 {
                case 606:
                    v18 := GenS1{F0: 786}
                    _ = v18
                }
                fmt.Println(fi11, fi11, fi11)
                p1 = p1
            }
            v10 = ((0.1 + float64((((3.14159 * (0.1 / 1.5)) * 1e10) * 3.14159))) == -2.25)
            var v19 bool = v10
            _ = fi13
            _ = v19
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return -82
}
func genFn1() int {
    var v20 uint64 = 893
    if true {
        var v21 int8 = -122
        var v22 int64 = 42
        v23 := GenS2{F0: 2, F1: -1}
        _ = v21
        _ = v22
        _ = v23
    }
    v20 += 42
    v24 := []int{-111}
    _ = v20
    _ = v24
    return -103
}

func main() {
    v1 := []string{("\n\t" + ((("\n\t" + "hello") + ("hello" + "foo bar")) + "foo bar"))}
    var v2 int16 = 1000
    var v3 uint16 = 312
    v4 := []string{"foo bar"}
    var v5 uint64 = 0
    v25 := GenS3{F0: "", F1: v2}
    v26 := GenS4{F0: "\n\t"}
    v3 = (^(v3))
    v2 = v2
    fmt.Println(v5)
    fmt.Println(v2)
    fmt.Println(v3, v2)
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS4{F0: "\n\t"}
        _ = fi27
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
