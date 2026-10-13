package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 bool }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int16; F1 int16; F2 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1000
    var v3 uint64 = 1
    var v4 int16 = (v2 + (v2 / ((687) | 1)))
    switch "" {
    case "hello":
        if true {
            for fi5 := 0; fi5 < 15; fi5++ {
                for fi6 := 0; fi6 < 5; fi6++ {
                    v7 := GenS1{F0: 76, F1: -1, F2: false}
                    fmt.Println(v2)
                    fmt.Println(v3, v4)
                    _ = fi6
                    _ = v7
                }
                var v8 int64 = 42
                var v9 int = 127
                _ = fi5
                _ = v8
                _ = v9
            }
        }
    }
    v3 += v3
    v2 = -238
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 uint64 = v3
        if (fi10 != 4294967295) {
            var v12 int8 = 37
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v3)
    v3 += (v3 / ((2) | 1))
    var v13 int = 0
    switch "foo bar" {
    case "hello":
        for fi14 := 0; fi14 < 20; fi14++ {
            if (-741 == -1) {
                v15 := GenS2{F0: true, F1: 255, F2: 486}
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := GenS3{F0: (v4 * v2), F1: (v2 % ((v4) | 1)), F2: 18446744073709551615}
                    _ = fi16
                    _ = v17
                }
                v18 := GenS4{F0: 0}
                _ = v15
                _ = v18
            }
            _ = fi14
        }
    }
    if (true || ((!true) && true)) {
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 int32 = 0
            for fi21 := 0; fi21 < 20; fi21++ {
                v22 := v1[v13]
                v23 := GenS2{F0: true, F1: 255, F2: 32767}
                var v24 float64 = float64(1.5)
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            _ = fi19
            _ = v20
        }
        v4 = 1
        for fi25 := 0; fi25 < 12; fi25++ {
            v26 := GenS3{F0: -54, F1: v2, F2: 60}
            _ = fi25
            _ = v26
        }
    }
    v27 := []int{(v13 >> ((1) & 63))}
    switch "hello" {
    case "\n\t":
        v13 = v13
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(v4, v4)
    case "":
        if false {
            fmt.Println(v3)
        } else {
            var v28 uint32 = 2147483647
            _ = v28
        }
    case "hello":
        if (0.0 > 0.0) {
            v29 := GenS5{F0: 0, F1: 1, F2: 352}
            _ = v29
        } else {
            v3 += 18446744073709551615
            v30 := GenS1{F0: 1, F1: 62, F2: true}
            v1[(-(v13))] = (v13 | -698)
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 uint8 = 42
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 int8 = 0
            _ = fi33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    if false {
        var v35 uint64 = v3
        _ = v35
    } else {
        v3 += (^(v3))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
