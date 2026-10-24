package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int; F1 int8; F2 int }
func genFn0(p0 uint64) uint32 {
    var v4 int = -368
    var v5 float32 = (3.14159 * 0.1)
    v5 = v5
    _ = p0
    _ = v4
    _ = v5
    return 752
}

func main() {
    var v1 float32 = float32(((float64(3.14159) + -2.25) + (0.0 + ((0.1 * 0.0) + 0.0))))
    var v2 uint32 = 0
    v3 := make(map[int]int)
    v2 -= (^(v2))
    switch "\n\t" {
    case "hello":
        delete(v3, -856)
    case "\n\t":
        v6 := GenS1{F0: 1, F1: 127, F2: (v2 + 2147483648)}
        _ = v6
    case "a":
        v3[int(v2)] = -314
    }
    v7 := []int{594, 1, 9223372036854775807, 1}
    var v8 string = ""
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS1{F0: 9223372036854775807, F1: 1, F2: (v2 / ((42) | 1))}
        fmt.Println(v8, v1)
        var v11 bool = (-40 == 100)
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int16 = -1
    fmt.Println(v8)
    var v13 uint = 463
    v14 := GenS1{F0: 1, F1: 0, F2: (^(v2))}
    fmt.Println(v2, v1, v8)
    for fi15 := 0; fi15 < 13; fi15++ {
        v2 = v2
        v13 -= (v13 + v13)
        _ = fi15
    }
    v16 := GenS2{F0: 1, F1: 0}
    var v17 uint16 = uint16(v2)
    v18 := GenS3{F0: -2147483648, F1: 78, F2: 9223372036854775807}
    switch (("hello" + "\n\t") + "a") {
    case "":
        switch ("" + v8) {
        case "hello":
            v2 -= 1
        default:
            if ((!(2 < 0)) && (true && (true || (uint64(v13) <= 2)))) {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v12 = v12
            } else {
                fmt.Println(v17)
                v13 = v13
                var v19 uint16 = (v17 * 32767)
                _ = v19
            }
        }
    case "\n\t":
        if (!false) {
            for fi20 := 0; fi20 < 16; fi20++ {
                v8 = "\n\t"
                var v21 int8 = -111
                _ = fi20
                _ = v21
            }
            v7 = append(v7, 128)
            if ((v12 & (v12 + (v12 % ((-32768) | 1)))) <= (-(v12))) {
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 int8 = 0
                    v24 := GenS1{F0: -9223372036854775808, F1: 218, F2: (v2 - v2)}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
        } else {
            for fi25 := 0; fi25 < 2; fi25++ {
                v17 -= (v17 * 119)
                v7 = append(v7, (fi25 - fi25))
                _ = fi25
            }
            switch v8 {
            case "":
                if (100 != -98) {
                    var v26 string = "foo bar"
                    var v27 int64 = -1
                    _ = v26
                    _ = v27
                }
            case "hello":
                v28 := GenS1{F0: 0, F1: 78, F2: v2}
                _ = v28
            }
        }
    default:
        for fi29 := 0; fi29 < 13; fi29++ {
            for fi30 := 0; fi30 < 14; fi30++ {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi30
            }
            v8 = "\n\t"
            _ = fi29
        }
    }
    fmt.Println(v17, v13, v8)
    fmt.Println(v8, v17)
    switch -1 {
    case -709:
        v17 = v17
    case -1:
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
