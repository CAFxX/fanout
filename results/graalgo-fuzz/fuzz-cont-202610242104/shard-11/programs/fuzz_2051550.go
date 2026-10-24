package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 bool; F1 int16; F2 uint8 }
func genFn0() int32 {
    var v7 float64 = float64(-2.25)
    var v8 float64 = v7
    var v9 uint8 = 143
    v8 = float64(((0.0 * float32(0.0)) * (1e10 + 0.0)))
    v10 := GenS1{F0: "\n\t", F1: -639, F2: 1}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 int16) int32 {
    if false {
        v11 := []int{}
        switch "foo bar" {
        case "hello":
            var v12 float32 = 1.5
            _ = v12
        case "foo bar":
            var v13 uint8 = 229
            _ = v13
        case "a":
            var v14 uint16 = 916
            _ = v14
        }
        _ = v11
    }
    p0 = (p0 << ((1) & 63))
    if (517 >= 65535) {
        v15 := GenS1{F0: "foo bar", F1: -557, F2: (p0 / (((p0 ^ (p0 & (p0 & -1)))) | 1))}
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS1{F0: "", F1: 684, F2: (p0 + -1)}
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = -407
    v3 := []uint32{2}
    v4 := []string{("\n\t" + "hello"), "foo bar", "\n\t", ("a" + ("hello" + ""))}
    v5 := make(map[string]int)
    var v6 int32 = -298
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 bool = (0.0 >= 0.1)
        for fi20 := 0; fi20 < 13; fi20++ {
            if (!(!(1 >= -9223372036854775808))) {
                switch "hello" {
                case "a":
                    v19 = v19
                }
                var v21 uint8 = 212
                fmt.Println(fi18, fi20, v19)
                _ = len(v4)
                switch ("hello" + ("foo bar" + "a")) {
                case "hello":
                    v19 = false
                default:
                    switch (("a" + ("" + ("hello" + ("\n\t" + "a")))) + "foo bar") {
                    case "foo bar":
                        var v22 uint = 18446744073709551615
                        _ = v22
                    }
                }
                _ = v21
            } else {
                fmt.Println(fi20, fi18)
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    if (166 == 255) {
        var v23 int64 = 9223372036854775807
        for fi24 := 0; fi24 < 10; fi24++ {
            if (326 != 1) {
                v23 -= 9223372036854775807
            }
            v2 = (-(fi24))
            var v25 float64 = 1.5
            _ = fi24
            _ = v25
        }
        var v26 int32 = (v6 * (v6 / ((v6) | 1)))
        _ = v23
        _ = v26
    }
    switch v2 {
    case -462:
        var v27 int8 = 42
        _ = v27
    case 0:
        for fi28 := 0; fi28 < 5; fi28++ {
            if true {
                v1[(fi28 << ((629) & 63))] = true
            }
            _ = fi28
        }
    case 164:
        v29 := GenS2{F0: -1, F1: 1000}
        _ = v29
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        fmt.Println(v6, fi30, fi30)
        v6 = 0
        _ = fi30
    }
    var v31 int8 = -1
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 int = v2
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 uint64 = 42
        var v36 int8 = -128
        v37 := GenS3{F0: false, F1: 32767, F2: 0}
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
