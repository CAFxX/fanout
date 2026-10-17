package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 int64 }

func main() {
    var v1 uint8 = 186
    var v2 uint8 = 128
    var v3 int8 = 42
    var v4 float64 = (1e10 * -2.25)
    var v5 uint16 = 645
    var v6 uint32 = 2147483647
    if false {
        var v7 int32 = 386
        if false {
            v8 := []int{}
            var v9 float32 = (0.1 - 1e10)
            var v10 uint = 591
            _ = v10
            _ = v8
            _ = v9
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: 192, F1: 0}
            var v13 uint16 = (v5 ^ (^(v5)))
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v7
    }
    var v14 float64 = (v4 / float64(float32((-2.25 / 1.0))))
    switch "foo bar" {
    case "foo bar":
        v15 := GenS2{F0: (v14 <= float64(-2.25)), F1: 1, F2: 42}
        _ = v15
    case "a":
        switch "\n\t" {
        case "":
            v3 = (v3 % (((v3 + (-(v3)))) | 1))
        }
    }
    v5 += v5
    if true {
        v2 += v1
    } else {
        if (569 > 758) {
            v16 := []int{-551, 543, -686, 9223372036854775807}
            _ = v16
        } else {
            var v17 int = 1
            _ = v17
        }
    }
    v2 = v2
    if false {
        if false {
            var v18 int16 = 600
            for fi19 := 0; fi19 < 4; fi19++ {
                v6 = v6
                v3 = v3
                v14 += float64(0.1)
                _ = fi19
            }
            v18 -= int16(v3)
            _ = v18
        } else {
            v6 += v6
            fmt.Println(v1)
            fmt.Println(v6)
        }
        v4 -= v4
    } else {
        v20 := GenS1{F0: 9223372036854775807, F1: v3}
        _ = v20
    }
    switch 1 {
    case 4294967295:
        v21 := []int{int(v3), int(v2), 440}
        _ = v21
    case -1:
        fmt.Println(v4, v5)
    case -901:
        v22 := []int{}
        _ = v22
    default:
        switch "hello" {
        case "\n\t":
            var v23 uint16 = v5
            _ = v23
        case "foo bar":
            switch "foo bar" {
            case "hello":
                var v24 float64 = v14
                _ = v24
            case "a":
                var v25 int64 = 1
                _ = v25
            case "\n\t":
                var v26 float32 = 0.1
                _ = v26
            }
        }
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
