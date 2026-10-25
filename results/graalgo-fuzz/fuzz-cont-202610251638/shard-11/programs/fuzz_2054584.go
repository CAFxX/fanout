package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int8; F2 bool }
func genFn0(p0 int64, p1 bool) int64 {
    p0 += -1
    fmt.Println(p1)
    fmt.Println(p0, p0)
    var v7 int8 = 0
    _ = p0
    _ = p1
    _ = v7
    return -680
}

func main() {
    var v1 int8 = 18
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    var v4 uint8 = 128
    v5 := []string{"\n\t", "foo bar"}
    var v6 uint64 = 1
    fmt.Println(v6)
    v8 := GenS1{F0: (false || false), F1: 877}
    v1 = v1
    var v9 uint16 = uint16(v6)
    v6 = v6
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v6 = v6
            if len(v5) > 0 {
                v5[0] = ""
            }
            fmt.Println(v1, v4)
            v12 := GenS2{F0: (v4 >> ((947) & 63))}
            _ = fi11
            _ = v12
        }
        if false {
            var v13 float32 = (0.0 * 0.0)
            _ = v13
        }
        _ = fi10
    }
    if false {
        v14 := GenS3{F0: 42, F1: 27, F2: false}
        v15 := GenS1{F0: ((-128 > v1) && true), F1: 796}
        _ = v14
        _ = v15
    } else {
        var v16 uint = 2
        _ = v16
    }
    if (!false) {
        v17 := GenS3{F0: 707, F1: v1, F2: ((-2.25 + (3.14159 + (float32(1.5) / (1.5 * ((0.0 * (0.0 * float32(-2.25))) + 3.14159))))) < 0.1)}
        v1 = 127
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        _ = v17
    } else {
        var v18 int16 = -32768
        if false {
            v9 = v9
        }
        _ = v18
    }
    var v19 uint32 = 2
    fmt.Println(v1, v19, v6)
    for fi20 := 0; fi20 < 11; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            v4 = 2
            v22 := GenS1{F0: true, F1: 162}
            v4 = v4
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 int = 127
    fmt.Println(v23)
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
