package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32; F1 int64; F2 uint64 }
func genFn0() int {
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint = 0
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    switch ("" + ("a" + "")) {
    case "":
        var v10 int = -146
        _ = v10
    case "\n\t":
        var v11 uint = 1
        _ = v11
    case "foo bar":
        if (!false) {
            var v12 uint64 = 1
            _ = v12
        } else {
            for fi13 := 0; fi13 < 3; fi13++ {
                v9 = append(v9, fi13)
                switch fi13 {
                case 1:
                    for fi14 := 0; fi14 < 11; fi14++ {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        var v15 int32 = 1
                        _ = fi14
                        _ = v15
                    }
                default:
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
                _ = fi13
            }
            var v16 float32 = 0.0
            for fi17 := 0; fi17 < 9; fi17++ {
                v16 = -2.25
                _ = fi17
            }
            _ = v16
        }
    }
    _ = v9
    return 128
}
func genFn1(p0 int) int64 {
    p0 = (-(p0))
    if true {
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 uint = 0
            _ = fi18
            _ = v19
        }
        v20 := []int{p0, p0, (-(p0))}
        p0 = p0
        _ = v20
    }
    _ = p0
    return 0
}

func main() {
    var v1 int64 = 1
    var v2 uint32 = 65535
    var v3 float64 = 3.14159
    var v4 int8 = -91
    v5 := make(map[uint32]int)
    v6 := []uint32{v2, (v2 ^ (v2 / ((697) | 1)))}
    switch (("" + "") + "a") {
    case "\n\t":
        if (794 <= 1) {
            v21 := v5[v2]
            _ = v21
        }
    case "hello":
        var v22 int = 9223372036854775807
        _ = v22
    }
    v23 := GenS1{F0: 255, F1: -100}
    if (!true) {
        for fi24 := 0; fi24 < 10; fi24++ {
            for fi25 := 0; fi25 < 19; fi25++ {
                fmt.Println(v4)
                v3 -= (v3 / v3)
                v26 := GenS1{F0: 744, F1: 1}
                _ = fi25
                _ = v26
            }
            var v27 int32 = 42
            _ = fi24
            _ = v27
        }
    }
    var v28 int16 = -658
    v29 := GenS2{F0: -1, F1: v1, F2: 491}
    var v30 float32 = float32(v3)
    var v31 uint16 = 724
    v31 = (v31 & 972)
    for fi32 := 0; fi32 < 12; fi32++ {
        v3 = ((v3 / v3) / 1e10)
        _ = fi32
    }
    v3 = v3
    if true {
        var v33 float32 = (((float32(v3) + 0.1) + 0.0) / v30)
        for fi34 := 0; fi34 < 14; fi34++ {
            fmt.Println(v4, fi34, v31)
            fmt.Println(v28, v30, v30)
            _ = fi34
        }
        _ = v33
    } else {
        v30 -= v30
        var v35 int16 = v28
        fmt.Println(v30)
        _ = v35
    }
    var v36 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
