package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint8; F2 int8 }

func main() {
    var v1 int32 = 1
    v2 := make(map[string]string)
    var v3 string = "a"
    switch v3 {
    case "foo bar":
        if (0 < 9223372036854775807) {
            switch 224 {
            case -9223372036854775808:
                v4 := GenS1{F0: 9223372036854775807, F1: -9223372036854775808}
                _ = v4
            case 738:
                var v5 float32 = (((3.14159 / ((float32(1e10) * -2.25) + float32(0.0))) / 1.0) / 1.0)
                _ = v5
            case 4294967295:
                for fi6 := 0; fi6 < 15; fi6++ {
                    v7 := GenS1{F0: 0, F1: 429}
                    if (true || (!false)) {
                        fmt.Println(fi6)
                        var v8 uint = 1
                        v9 := GenS1{F0: -515, F1: 9223372036854775807}
                        _ = v8
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 13; fi10++ {
                        var v11 string = v3
                        _ = fi10
                        _ = v11
                    }
                    _ = fi6
                    _ = v7
                }
            }
            v2["a"] = v3
        } else {
            v3 = "\n\t"
        }
    case "a":
        fmt.Println(v1, v1)
    }
    v3 = v3
    v12 := GenS1{F0: 42, F1: 9223372036854775807}
    for fi13 := 0; fi13 < 18; fi13++ {
        v3 = v3
        var v14 float64 = (1e10 + 0.1)
        _ = fi13
        _ = v14
    }
    v15 := []int{}
    switch -249 {
    case 4294967295:
        v16 := GenS1{F0: 160, F1: 80}
        _ = v16
    case -9223372036854775808:
        var v17 int32 = (-(v1))
        _ = v17
    case 2147483647:
        if (!true) {
            v3 = "\n\t"
            if (0 <= -864) {
                v3 = v3
                if true {
                    v18 := GenS2{F0: 128}
                    _ = v18
                }
            }
            v19 := GenS3{F0: v3, F1: 39, F2: -79}
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            switch v3 {
            case "hello":
                var v22 string = v3
                _ = v22
            case "foo bar":
                v15 = append(v15, (fi21 % (((fi20 >> ((317) & 63))) | 1)))
            default:
                v23 := GenS1{F0: 738, F1: 42}
                _ = v23
            }
            _ = fi21
        }
        var v24 uint = 667
        _ = fi20
        _ = v24
    }
    if false {
        switch v3 {
        case "a":
            delete(v2, ("foo bar" + v3))
        case "foo bar":
            switch v3 {
            case "hello":
                if (!(!true)) {
                    var v25 int = 0
                    _ = v25
                }
            case "a":
                fmt.Println(v1, v1)
            case "foo bar":
                for fi26 := 0; fi26 < 18; fi26++ {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        delete(v2, v3)
                        var v28 int64 = -259
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
            default:
                if (!true) {
                    var v29 int = 0
                    v30 := GenS2{F0: v29}
                    _ = v29
                    _ = v30
                }
            }
        default:
            var v31 uint32 = 2
            _ = v31
        }
    } else {
        v1 = (v1 | 0)
    }
    var v32 float64 = (3.14159 + 0.0)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v32
}
