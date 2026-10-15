package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 int16; F2 int8 }

func main() {
    var v1 string = ""
    var v2 float32 = 0.0
    v3 := make(map[string]string)
    var v4 int16 = 32767
    var v5 float32 = v2
    var v6 int8 = 100
    if (!(2 != 443)) {
        var v7 uint64 = 130
        if (237 <= 779) {
            _ = len(v3)
        }
        fmt.Println(v6, v7, v1)
        _ = v7
    } else {
        var v8 uint8 = 0
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := GenS1{F0: 665}
            _ = fi9
            _ = v10
        }
        if ((v4 << ((328) & 63)) <= v4) {
            fmt.Println(v5, v4, v8)
        } else {
            switch 32767 {
            case 1:
                v3[("" + v1)] = v1
            case -1:
                v11 := []int{1, -9223372036854775808, 32767}
                _ = v11
            }
        }
        _ = v8
    }
    switch v1 {
    case "hello":
        for fi12 := 0; fi12 < 20; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                v4 += (v4 & 1000)
                v14 := GenS1{F0: 4294967295}
                var v15 uint32 = 0
                fmt.Println(v4, v4)
                v16 := GenS2{F0: -32768, F1: v4, F2: v6}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            v17 := GenS1{F0: 0}
            v2 += (v5 * v5)
            _ = fi12
            _ = v17
        }
    case "\n\t":
        for fi18 := 0; fi18 < 4; fi18++ {
            v4 = (v4 >> ((18446744073709551615) & 63))
            var v19 float32 = v2
            _ = fi18
            _ = v19
        }
    case "":
        if false {
            switch v1 {
            case "":
                var v20 uint = 42
                _ = v20
            case "a":
                v21 := v3[v1]
                _ = v21
            }
        }
    default:
        if (!(0 == 187)) {
            v4 = 235
        } else {
            for fi22 := 0; fi22 < 11; fi22++ {
                fmt.Println(fi22)
                _ = fi22
            }
        }
    }
    v23 := []int{127, 1}
    var v24 uint32 = 238
    var v25 float64 = (0.1 + 0.0)
    v26 := GenS2{F0: (^(v4)), F1: v4, F2: (v6 - v6)}
    var v27 float64 = -2.25
    _ = len(v3)
    _ = len(v23)
    v25 = (3.14159 / float64((v2 * v5)))
    v5 = 1e10
    var v28 uint8 = 128
    fmt.Println(v1, v28)
    if ((v28 - 128) > v28) {
        v6 = v6
        for fi29 := 0; fi29 < 15; fi29++ {
            switch fi29 {
            case 9223372036854775807:
                v28 = 127
            case -9223372036854775808:
                var v30 bool = false
                _ = v30
            case 4294967295:
                v1 = ""
            }
            var v31 float64 = v25
            _ = fi29
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        fmt.Println(v24, v27, v1)
        v33 := GenS1{F0: 2}
        v5 -= (v5 + float32((v25 / 1.0)))
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
