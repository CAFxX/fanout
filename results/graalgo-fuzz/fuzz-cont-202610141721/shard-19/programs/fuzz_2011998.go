package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 int, p1 int) uint16 {
    if true {
        switch "hello" {
        case "foo bar":
            var v6 int16 = 0
            _ = v6
        default:
            var v7 int32 = -18
            _ = v7
        }
        v8 := []int{}
        fmt.Println(p0)
        _ = v8
    } else {
        v9 := GenS1{F0: 65535}
        p0 += (p1 | p1)
        v10 := GenS1{F0: 2}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int64 = 0
            p1 = (fi11 - 42)
            _ = fi12
            _ = v13
        }
        var v14 int32 = -872
        _ = fi11
        _ = v14
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []uint32{}
    v2 := []string{("a" + "foo bar")}
    var v3 int16 = -639
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]bool)
    var v15 int64 = 729
    var v16 string = "a"
    switch 9223372036854775807 {
    case 9223372036854775807:
        v1 = append(v1, 4294967295)
    case 595:
        v17 := GenS1{F0: 1}
        _ = v17
    case 32767:
        fmt.Println(v16)
    }
    v18 := GenS2{F0: 60}
    var v19 int = 128
    var v20 string = "hello"
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := GenS3{F0: -1, F1: (863 >= 1)}
        _ = fi21
        _ = v22
    }
    var v23 float32 = (3.14159 - float32((1e10 / 0.1)))
    var v24 float32 = (v23 + (v23 * ((v23 - v23) / ((((v23 * (v23 - (((v23 / 1.5) * v23) - v23))) + 0.1) / v23) * (v23 + ((v23 - v23) + v23))))))
    var v25 uint32 = 65535
    for fi26 := 0; fi26 < 14; fi26++ {
        fmt.Println(v23)
        if (((v20 + "foo bar") + v20) == (v20 + v16)) {
            var v27 int8 = 42
            _ = v27
        }
        _ = fi26
    }
    v25 -= (^(v25))
    v28 := GenS2{F0: 42}
    switch 32767 {
    case 833:
        switch (v16 + (((v16 + "hello") + ("" + (((v20 + ("foo bar" + (((((v16 + (((v16 + v20) + v20) + v16)) + v20) + "\n\t") + v16) + "hello"))) + "a") + v20))) + v20)) {
        case "hello":
            if ((110 != 42) && false) {
                var v29 string = (("a" + v16) + ((v20 + "") + "a"))
                if true {
                    var v30 uint16 = 32767
                    _ = v30
                }
                _ = v29
            } else {
                var v31 float32 = 3.14159
                var v32 int32 = 42
                _ = v31
                _ = v32
            }
        case "a":
            fmt.Println(v25, v19)
        }
    case 127:
        var v33 uint16 = 65535
        _ = v33
    case 0:
        _ = len(v4)
    }
    switch -897 {
    case 218:
        var v34 uint = 48
        _ = v34
    default:
        fmt.Println(v15, v25)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
