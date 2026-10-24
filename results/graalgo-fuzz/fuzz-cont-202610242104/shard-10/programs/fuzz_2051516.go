package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 string }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 float32 = 1e10
    var v2 int8 = 100
    var v3 uint16 = 581
    var v4 int8 = v2
    var v5 uint16 = (v3 & (v3 << ((255) & 63)))
    var v6 uint8 = 35
    fmt.Println(v3)
    v7 := []int{84}
    v2 -= v4
    if false {
        switch ("hello" + "foo bar") {
        case "\n\t":
            fmt.Println(v4, v2, v3)
        }
        var v8 float32 = float32(-2.25)
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            if false {
                v3 = (v3 | v5)
                var v11 bool = false
                _ = v11
            }
            switch fi9 {
            case -289:
                v12 := GenS1{F0: (127 > v6), F1: 281, F2: 98}
                _ = v12
            case -601:
                if true {
                    var v13 int32 = 2147483647
                    _ = v13
                }
            default:
                var v14 uint8 = 127
                var v15 float64 = (float64(float32((float64(0.1) * 1e10))) * -2.25)
                fmt.Println(fi10, v5)
                if len(v7) > 0 {
                    v7[0] = (^(fi10))
                }
                v15 += v15
                _ = v14
                _ = v15
            }
            _ = fi10
        }
        _ = fi9
    }
    v6 -= (v6 ^ (v6 & (v6 / ((v6) | 1))))
    var v16 uint = uint(v6)
    fmt.Println(v16, v6, v6)
    v2 += (v2 * (v4 + (v4 % ((117) | 1))))
    v7 = append(v7, -2147483648)
    if true {
        var v17 uint32 = 4294967295
        for fi18 := 0; fi18 < 7; fi18++ {
            v6 += 42
            switch "a" {
            case "\n\t":
                v19 := GenS2{F0: true, F1: -9223372036854775808, F2: ""}
                _ = v19
            case "foo bar":
                fmt.Println(v6)
            case "hello":
                var v20 int64 = 845
                _ = v20
            default:
                v21 := GenS3{F0: 1, F1: 348}
                _ = v21
            }
            v22 := GenS1{F0: (float64(v1) == 3.14159), F1: v16, F2: 1}
            _ = fi18
            _ = v22
        }
        fmt.Println(v3, v4, v3)
        _ = v17
    }
    fmt.Println(v1, v3, v3)
    var v23 uint16 = v3
    v24 := GenS4{F0: 304}
    if false {
        var v25 uint64 = 881
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := GenS2{F0: false, F1: -839, F2: "a"}
        fmt.Println(fi26, v23)
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        v6 += 255
        if true {
            var v29 int64 = 114
            var v30 uint = v16
            _ = v29
            _ = v30
        } else {
            for fi31 := 0; fi31 < 11; fi31++ {
                if len(v7) > 0 {
                    v7[0] = (fi31 % ((fi28) | 1))
                }
                v23 += (v3 | v3)
                v32 := GenS4{F0: 470}
                _ = fi31
                _ = v32
            }
            v1 = (-2.25 * float32(-2.25))
        }
        if ((-(v4)) > (^(v2))) {
            var v33 int = (fi28 * (fi28 % ((fi28) | 1)))
            fmt.Println(v16)
            v34 := GenS4{F0: 2147483647}
            _ = v33
            _ = v34
        }
        _ = fi28
    }
    v35 := GenS4{F0: 280}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
