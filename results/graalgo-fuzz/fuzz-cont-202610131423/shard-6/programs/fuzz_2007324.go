package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint; F1 int; F2 int8 }
type GenS4 struct { F0 int32; F1 uint64 }
func genFn0(p0 int, p1 int32, p2 uint16) uint16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        switch ("hello" + ("foo bar" + ("" + "a"))) {
        case "hello":
            var v8 uint = 65535
            _ = v8
        case "\n\t":
            v9 := GenS1{F0: 42, F1: true}
            _ = v9
        }
        var v10 uint64 = 248
        switch (p0 >> ((2) & 63)) {
        case 9223372036854775807:
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 uint8 = 201
                v13 := GenS2{F0: (^(p1)), F1: 0}
                _ = fi11
                _ = v12
                _ = v13
            }
        case 799:
            v14 := GenS2{F0: 326, F1: 9223372036854775807}
            _ = v14
        case 42:
            if true {
                var v15 int = fi7
                _ = v15
            }
        }
        _ = fi7
        _ = v10
    }
    v16 := GenS3{F0: 4294967295, F1: p0, F2: -1}
    v17 := []int{}
    for fi18 := 0; fi18 < 5; fi18++ {
        fmt.Println(p1)
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return (^(p2))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = (float64(0.1) / (-2.25 / 1e10))
    var v3 int32 = -687
    var v4 uint64 = 23
    v5 := make(map[string]string)
    v6 := make(map[string]uint32)
    var v19 int64 = -1
    switch "foo bar" {
    case "a":
        v4 = (^(v4))
    }
    switch ((("\n\t" + "foo bar") + "\n\t") + "\n\t") {
    case "":
        if false {
            var v20 float64 = 0.0
            _ = v20
        }
    case "foo bar":
        fmt.Println(v3)
    case "hello":
        v21 := GenS2{F0: -1, F1: (v19 - (v19 / ((v19) | 1)))}
        _ = v21
    }
    if false {
        var v22 int64 = 1
        _ = v22
    }
    v23 := GenS3{F0: 4294967295, F1: 128, F2: 100}
    fmt.Println(v2, v3)
    v24 := []int{}
    switch "a" {
    case "\n\t":
        var v25 int16 = -338
        _ = v25
    case "a":
        if true {
            if true {
                v2 += float64(-2.25)
            } else {
                v3 -= v3
                v4 -= (v4 << ((4294967295) & 63))
            }
            v26 := GenS4{F0: 511, F1: 58}
            _ = v26
        } else {
            v27 := GenS4{F0: v3, F1: v4}
            _ = v27
        }
    default:
        switch 0 {
        case -78:
            var v28 uint8 = 1
            _ = v28
        case 741:
            for fi29 := 0; fi29 < 4; fi29++ {
                v2 += v2
                var v30 float32 = (1.5 * 0.0)
                _ = fi29
                _ = v30
            }
        case -1:
            v31 := GenS3{F0: 18446744073709551615, F1: 392, F2: 1}
            _ = v31
        default:
            v19 += v19
        }
    }
    var v32 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
