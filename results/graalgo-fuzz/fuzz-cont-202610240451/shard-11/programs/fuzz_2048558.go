package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int64; F2 int8 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int8 = -1
    v2 := []string{}
    v3 := make(map[string]uint32)
    v4 := []int{}
    var v5 uint32 = 2147483648
    v6 := GenS1{F0: 2}
    fmt.Println(v5)
    v4 = append(v4, -2147483648)
    var v7 int8 = (^(v1))
    switch 9223372036854775807 {
    case 127:
        switch ("a" + ((("a" + ((("a" + "a") + ("hello" + ("\n\t" + ""))) + "hello")) + "\n\t") + "foo bar")) {
        case "":
            var v8 bool = false
            _ = v8
        case "\n\t":
            v5 += (v5 | 4294967295)
        }
    case 686:
        for fi9 := 0; fi9 < 2; fi9++ {
            if len(v4) > 0 {
                v4[0] = fi9
            }
            for fi10 := 0; fi10 < 19; fi10++ {
                if (196 == 128) {
                    v11 := GenS1{F0: 18446744073709551615}
                    v12 := GenS2{F0: 255}
                    v1 -= (v1 * 1)
                    var v13 int16 = 32767
                    if false {
                        v14 := GenS3{F0: 255}
                        var v15 bool = (0 > 2)
                        var v16 int32 = 2147483647
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        v13 = v13
                        var v17 int16 = v13
                        _ = v17
                    }
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    v18 := v3["a"]
                    fmt.Println(fi10)
                    v19 := v3["\n\t"]
                    _ = v18
                    _ = v19
                }
                v1 = (^(v1))
                _ = fi10
            }
            _ = fi9
        }
    case 131:
        for fi20 := 0; fi20 < 6; fi20++ {
            for fi21 := 0; fi21 < 5; fi21++ {
                v22 := GenS4{F0: v1, F1: -9223372036854775808, F2: 0}
                fmt.Println(v7)
                _ = fi21
                _ = v22
            }
            v5 = v5
            for fi23 := 0; fi23 < 14; fi23++ {
                v24 := GenS2{F0: uint8(v5)}
                _ = fi23
                _ = v24
            }
            _ = fi20
        }
    }
    if (1e10 < 0.0) {
        var v25 bool = false
        var v26 int16 = 0
        fmt.Println(v7)
        _ = v25
        _ = v26
    } else {
        var v27 int64 = -906
        _ = v27
    }
    fmt.Println(v1)
    var v28 bool = (!((-2.25 + 0.1) <= 0.0))
    v3[("hello" + ((("hello" + "hello") + "\n\t") + ""))] = 2
    v28 = (v28 || (!true))
    if v28 {
        var v29 int8 = v7
        v3[("foo bar" + "hello")] = (^(v5))
        v4 = append(v4, 0)
        _ = v29
    } else {
        v28 = (759 != 32767)
    }
    var v30 float32 = ((3.14159 + (-2.25 + (0.0 * -2.25))) + (1e10 - -2.25))
    var v31 uint16 = 143
    if (v31 == 1) {
        v5 -= (v5 >> ((1) & 63))
    } else {
        v30 -= (v30 - float32(0.0))
    }
    var v32 uint8 = 42
    switch "a" {
    case "a":
        var v33 int8 = 100
        _ = v33
    case "":
        var v34 int32 = 593
        _ = v34
    case "foo bar":
        for fi35 := 0; fi35 < 9; fi35++ {
            if (v30 < v30) {
                _ = len(v4)
                delete(v3, "hello")
                v36 := GenS5{F0: v31}
                _ = v36
            }
            v5 -= (v5 % ((v5) | 1))
            for fi37 := 0; fi37 < 8; fi37++ {
                fmt.Println(fi35, v7, v32)
                v30 += (1.5 / v30)
                _ = fi37
            }
            _ = fi35
        }
    default:
        v28 = (86 >= 656)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
