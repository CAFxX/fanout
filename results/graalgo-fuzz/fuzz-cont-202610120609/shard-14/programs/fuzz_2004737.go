package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 string; F1 int32; F2 int16 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int; F2 int }

func main() {
    var v1 float64 = (3.14159 * 0.1)
    var v2 int = -310
    v3 := make(map[int]uint32)
    var v4 float32 = 1e10
    var v5 float64 = 0.0
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: 100, F1: 2, F2: 0}
    if true {
        v8 := GenS1{F0: 42, F1: 0, F2: 9223372036854775807}
        v9 := GenS2{F0: ("foo bar" + ("foo bar" + "foo bar")), F1: 0, F2: -1}
        _ = v8
        _ = v9
    } else {
        var v10 uint8 = 128
        _ = v10
    }
    v1 = v1
    if ("a" <= ((("" + "foo bar") + "a") + "")) {
        var v11 uint8 = 0
        fmt.Println(v4)
        _ = v11
    } else {
        switch "foo bar" {
        case "hello":
            var v12 string = "foo bar"
            _ = v12
        case "foo bar":
            switch "hello" {
            case "":
                v13 := GenS3{F0: true, F1: 42, F2: 1}
                _ = v13
            case "\n\t":
                v5 = (float64(v4) / -2.25)
            default:
                if false {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        var v15 uint = 890
                        v5 -= (v1 / float64(-2.25))
                        v5 -= (v1 * (v1 / v5))
                        _ = fi14
                        _ = v15
                    }
                    var v16 uint8 = 42
                    v2 -= (v2 | (v2 & 9223372036854775807))
                    _ = v16
                } else {
                    v5 = (v1 + float64(float32((v1 + (1e10 - v5)))))
                    v2 += 0
                }
            }
        }
        v17 := []int{}
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v5 += v1
        if true {
            v1 = (1e10 - -2.25)
        }
        _ = fi18
    }
    v19 := GenS4{F0: 255, F1: v2, F2: 127}
    switch v2 {
    case 2147483647:
        var v20 int8 = -99
        _ = v20
    case 32768:
        v21 := GenS4{F0: 174, F1: (v2 - v2), F2: v2}
        _ = v21
    }
    var v22 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
