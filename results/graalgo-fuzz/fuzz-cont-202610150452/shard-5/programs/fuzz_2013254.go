package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 bool }
type GenS2 struct { F0 string; F1 uint16; F2 int16 }
type GenS3 struct { F0 int32; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = false
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v5 int = 0
    var v6 int = -503
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint8 = 42
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        if v2 {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 int = v6
                v12 := GenS1{F0: 903, F1: 53, F2: v2}
                v13 := v1[-1]
                v14 := []int{}
                fmt.Println(v2, v2, v11)
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
                _ = v14
            }
            v2 = v2
        }
        _ = fi9
    }
    var v15 uint16 = 65535
    var v16 int16 = 1
    fmt.Println(v15, v15)
    v17 := GenS2{F0: "foo bar", F1: v15, F2: v16}
    var v18 int32 = -2147483648
    v2 = (false && false)
    switch v5 {
    case -1:
        var v19 uint8 = 73
        _ = v19
    case 1:
        switch v5 {
        case -1:
            var v20 int = (^(v5))
            _ = v20
        case 32768:
            for fi21 := 0; fi21 < 10; fi21++ {
                v22 := GenS3{F0: (v18 | (v18 * 42)), F1: 127}
                _ = fi21
                _ = v22
            }
        case 254:
            var v23 uint32 = 753
            _ = v23
        default:
            fmt.Println(v2)
        }
    case 0:
        if v2 {
            v24 := GenS1{F0: 65535, F1: 42, F2: false}
            for fi25 := 0; fi25 < 12; fi25++ {
                var v26 bool = v2
                _ = fi25
                _ = v26
            }
            var v27 int64 = 1
            _ = v24
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 4; fi28++ {
            switch (("\n\t" + "foo bar") + "a") {
            case "":
                for fi29 := 0; fi29 < 15; fi29++ {
                    v5 -= (v5 << ((18446744073709551615) & 63))
                    v30 := []int{}
                    _ = fi29
                    _ = v30
                }
            case "\n\t":
                var v31 uint8 = 127
                _ = v31
            default:
                fmt.Println(v15)
            }
            _ = fi28
        }
    }
    if v2 {
        v32 := GenS2{F0: "\n\t", F1: 255, F2: v16}
        _ = v32
    }
    var v33 uint64 = 32
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
