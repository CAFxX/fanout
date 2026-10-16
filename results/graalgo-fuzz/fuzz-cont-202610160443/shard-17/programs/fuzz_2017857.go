package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int8 }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    v1 := []uint32{181}
    var v2 string = "foo bar"
    var v3 float64 = 0.0
    v4 := make(map[string]int)
    var v5 int64 = -1
    for fi6 := 0; fi6 < 11; fi6++ {
        v4[(v2 + v2)] = 127
        fmt.Println(v2)
        _ = fi6
    }
    v7 := GenS1{F0: 42, F1: true, F2: 100}
    v5 = v5
    var v8 float32 = float32(1.5)
    if false {
        v5 += -9223372036854775808
        switch int(v5) {
        case -41:
            var v9 uint32 = 4294967295
            _ = v9
        }
    }
    v2 = v2
    if (-9223372036854775808 > 42) {
        var v10 float32 = (v8 / (0.1 / (0.1 + v8)))
        switch v2 {
        case "hello":
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 uint = 318
                var v13 uint16 = 42
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            v2 = v2
        }
        _ = v10
    } else {
        v14 := GenS1{F0: -32768, F1: true, F2: 0}
        v3 = 0.0
        var v15 bool = false
        _ = v14
        _ = v15
    }
    var v16 float32 = (v8 * 3.14159)
    if (false || (42 > 0)) {
        for fi17 := 0; fi17 < 5; fi17++ {
            switch v2 {
            case "a":
                v18 := GenS2{F0: -782, F1: (255 == 42)}
                _ = v18
            case "\n\t":
                v19 := v4[v2]
                _ = v19
            }
            var v20 uint64 = 611
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(v2)
                v2 = v2
                v20 += v20
                var v22 string = ((v2 + "hello") + v2)
                _ = fi21
                _ = v22
            }
            _ = fi17
            _ = v20
        }
        v23 := GenS2{F0: -2147483648, F1: true}
        for fi24 := 0; fi24 < 19; fi24++ {
            if (!(32767 > 551)) {
                var v25 float32 = v8
                if false {
                    var v26 uint64 = 977
                    var v27 uint8 = 0
                    _ = v26
                    _ = v27
                } else {
                    v25 -= v16
                }
                var v28 uint32 = 42
                _ = v25
                _ = v28
            }
            _ = fi24
        }
        _ = v23
    } else {
        var v29 int = -1
        if false {
            v2 = "foo bar"
            switch v29 {
            case -998:
                v16 -= v16
            case 127:
                var v30 uint8 = 2
                _ = v30
            case -1:
                _ = len(v1)
            }
        } else {
            switch "\n\t" {
            case "a":
                _ = len(v1)
            }
        }
        _ = v29
    }
    v16 = ((1e10 / (-2.25 * (v8 + (0.0 + 1e10)))) / (0.0 + (v16 - (1e10 - v8))))
    for fi31 := 0; fi31 < 18; fi31++ {
        v2 = v2
        _ = fi31
    }
    switch -673 {
    case 2147483647:
        v16 -= float32(float64(v8))
    default:
        var v32 uint16 = 270
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
