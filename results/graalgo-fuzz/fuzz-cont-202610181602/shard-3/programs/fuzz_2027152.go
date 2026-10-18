package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int32; F1 int8 }

func main() {
    var v1 int32 = -1
    var v2 uint64 = 4294967295
    var v3 int16 = int16(v1)
    var v4 int32 = v1
    v5 := make(map[string]uint32)
    v6 := make(map[string]bool)
    if (("\n\t" != "foo bar") || (v4 == v1)) {
        var v7 int16 = (v3 >> ((65535) & 63))
        if (255 >= 0) {
            _ = len(v5)
            v8 := GenS1{F0: 25}
            fmt.Println(v1, v1)
            _ = v8
        } else {
            if false {
                v3 = -32768
                for fi9 := 0; fi9 < 19; fi9++ {
                    switch (fi9 + fi9) {
                    case -441:
                        var v10 int8 = -1
                        _ = v10
                    case 127:
                        v1 += (v4 | (v4 | 0))
                    case 2147483647:
                        v11 := []int{fi9, (fi9 % (((^(fi9))) | 1))}
                        _ = v11
                    }
                    var v12 int64 = 1
                    _ = fi9
                    _ = v12
                }
            } else {
                v13 := GenS2{F0: 127, F1: true}
                var v14 float64 = (0.0 - -2.25)
                _ = v13
                _ = v14
            }
        }
        var v15 uint32 = 0
        _ = v15
        _ = v7
    } else {
        for fi16 := 0; fi16 < 17; fi16++ {
            v4 = v4
            v17 := GenS3{F0: v1, F1: -128}
            var v18 float32 = 1.5
            _ = fi16
            _ = v17
            _ = v18
        }
        v3 = 1000
        fmt.Println(v2, v4)
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        v6[("\n\t" + "\n\t")] = (42 == 1)
        switch "" {
        case "foo bar":
            v5[""] = 0
        case "\n\t":
            fmt.Println(fi19, v3, v1)
        case "a":
            for fi20 := 0; fi20 < 9; fi20++ {
                v1 = (^(v4))
                _ = fi20
            }
        default:
            _ = len(v6)
        }
        _ = fi19
    }
    var v21 float64 = 1.5
    var v22 int64 = 1
    fmt.Println(v21, v21)
    var v23 int = -2147483648
    v24 := []int{v23}
    v25 := v5["a"]
    var v26 uint16 = 2
    v24 = append(v24, v23)
    var v27 uint64 = (v2 + (v2 % ((v2) | 1)))
    _ = len(v6)
    fmt.Println(v4, v21)
    switch "foo bar" {
    case "hello":
        v2 = (v2 * (v2 ^ (v27 + v27)))
    default:
        var v28 int = (v23 / ((v23) | 1))
        _ = v28
    }
    fmt.Println(v27, v23, v27)
    if false {
        if true {
            v3 = (v3 % (((v3 >> ((2) & 63))) | 1))
            for fi29 := 0; fi29 < 12; fi29++ {
                var v30 uint = 18446744073709551615
                v22 = -1
                _ = fi29
                _ = v30
            }
            fmt.Println(v21, v3)
        }
        var v31 float64 = v21
        _ = v31
    } else {
        v25 -= v25
        for fi32 := 0; fi32 < 15; fi32++ {
            _ = len(v6)
            _ = fi32
        }
        _ = len(v24)
    }
    switch "foo bar" {
    case "\n\t":
        for fi33 := 0; fi33 < 18; fi33++ {
            v21 = v21
            for fi34 := 0; fi34 < 6; fi34++ {
                switch (fi33 % ((fi33) | 1)) {
                case 0:
                    fmt.Println(v25)
                case 32767:
                    v35 := GenS3{F0: (v4 & v4), F1: 127}
                    _ = v35
                case 42:
                    switch -420 {
                    case 128:
                        v36 := GenS1{F0: 90}
                        _ = v36
                    }
                }
                v22 -= (v22 + v22)
                _ = fi34
            }
            var v37 uint64 = v27
            _ = fi33
            _ = v37
        }
    default:
        var v38 float32 = 0.0
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
