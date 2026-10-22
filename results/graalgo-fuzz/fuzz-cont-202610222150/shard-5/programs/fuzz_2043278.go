package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int; F2 int8 }
func genFn0(p0 bool) uint8 {
    if false {
        v5 := []int{-238, -354, -751, 32768}
        switch "a" {
        case "":
            p0 = false
        case "foo bar":
            p0 = false
        default:
            var v6 uint = 0
            _ = v6
        }
        _ = v5
    }
    switch ("foo bar" + (("\n\t" + "\n\t") + "\n\t")) {
    case "\n\t":
        if p0 {
            v7 := GenS1{F0: p0}
            p0 = false
            if true {
                var v8 float32 = (0.1 / 1e10)
                v9 := GenS1{F0: (1 > 602)}
                _ = v8
                _ = v9
            }
            _ = v7
        }
    case "":
        p0 = (false && (98 != 237))
    case "foo bar":
        var v10 string = "foo bar"
        _ = v10
    default:
        v11 := []int{128, 127}
        _ = v11
    }
    p0 = false
    if p0 {
        v12 := GenS2{F0: -154}
        _ = v12
    } else {
        var v13 int = 0
        _ = v13
    }
    var v14 bool = false
    _ = p0
    _ = v14
    return 172
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    var v3 float32 = 1e10
    var v4 int16 = 42
    var v15 string = (("\n\t" + "") + "\n\t")
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := GenS3{F0: 751, F1: fi16, F2: 42}
        _ = fi16
        _ = v17
    }
    switch 128 {
    case 965:
        var v18 float32 = 3.14159
        _ = v18
    default:
        switch v15 {
        case "":
            if false {
                if (!(359 != 42)) {
                    fmt.Println(v4, v4, v15)
                } else {
                    switch 42 {
                    case 32768:
                        var v19 int64 = 0
                        _ = v19
                    }
                }
                fmt.Println(v15, v4, v3)
            }
        }
    }
    v15 = v15
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 float64 = (1e10 - 0.0)
        _ = fi20
        _ = v21
    }
    var v22 bool = true
    v3 += (((1e10 * ((v3 * v3) - v3)) * 1e10) + v3)
    v23 := []int{42}
    v4 = v4
    var v24 uint = 2
    for fi25 := 0; fi25 < 8; fi25++ {
        v4 = v4
        for fi26 := 0; fi26 < 7; fi26++ {
            _ = len(v1)
            _ = fi26
        }
        v4 = -32768
        _ = fi25
    }
    var v27 int8 = 123
    fmt.Println(v27, v15, v22)
    var v28 int = -981
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
