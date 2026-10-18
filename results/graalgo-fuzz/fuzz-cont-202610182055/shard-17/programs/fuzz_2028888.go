package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int8; F1 string; F2 uint8 }
func genFn0(p0 uint8, p1 uint16, p2 int64) uint32 {
    switch "hello" {
    case "\n\t":
        p2 = (p2 * p2)
    case "hello":
        var v4 int16 = -403
        _ = v4
    case "foo bar":
        if false {
            v5 := GenS1{F0: p1, F1: 32767}
            _ = v5
        } else {
            var v6 uint32 = 267
            v7 := GenS2{F0: 42, F1: (p2 * 42)}
            _ = v6
            _ = v7
        }
    default:
        v8 := GenS1{F0: p1, F1: 32767}
        _ = v8
    }
    v9 := []int{4294967295}
    var v10 string = ""
    _ = len(v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 42
}
func genFn1() uint32 {
    switch -2147483648 {
    case -9223372036854775808:
        var v11 int8 = -1
        _ = v11
    }
    var v12 uint32 = 264
    _ = v12
    return (v12 >> ((uint(v12)) & 63))
}

func main() {
    v1 := []int{9223372036854775807}
    var v2 uint8 = 127
    var v3 int16 = 572
    v3 = v3
    v13 := GenS1{F0: 605, F1: -1}
    v14 := GenS3{F0: 100, F1: "", F2: (v2 * (^(v2)))}
    v15 := GenS1{F0: 255, F1: -947}
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2)
    case "a":
        fmt.Println(v3, v3)
    default:
        for fi16 := 0; fi16 < 15; fi16++ {
            if ((((float32((-2.25 - 3.14159)) / ((1.5 * 1.5) / float32((((float64(float32((float64(-2.25) * 1e10))) * 0.1) / 0.1) / (1.5 - 0.1))))) + (1e10 / 1.0)) / (float32(float64(1.5)) + -2.25)) <= (1e10 / 1.0)) {
                if false {
                    v17 := GenS2{F0: 1, F1: 0}
                    v2 = (v2 | v2)
                    _ = v17
                }
                var v18 uint8 = v2
                _ = v18
            } else {
                var v19 bool = false
                _ = len(v1)
                _ = v19
            }
            v20 := GenS1{F0: 501, F1: (v3 / (((v3 / (((v3 - v3)) | 1))) | 1))}
            _ = fi16
            _ = v20
        }
    }
    var v21 uint = 4294967295
    v21 += v21
    var v22 bool = (65535 < 1)
    var v23 uint16 = 255
    v3 = (v3 % ((0) | 1))
    v22 = false
    var v24 float32 = float32(3.14159)
    if len(v1) > 0 {
        v1[0] = 772
    }
    switch ("foo bar" + "") {
    case "hello":
        v1 = append(v1, 42)
    case "":
        if (v21 == 0) {
            switch "a" {
            case "":
                var v25 string = "a"
                _ = v25
            }
            fmt.Println(v21)
        }
    case "\n\t":
        switch "foo bar" {
        case "hello":
            v3 = v3
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
