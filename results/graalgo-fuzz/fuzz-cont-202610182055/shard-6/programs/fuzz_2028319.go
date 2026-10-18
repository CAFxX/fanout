package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 string }

func main() {
    var v1 float64 = ((1.5 * (-2.25 + 3.14159)) + float64(float32(-2.25)))
    var v2 uint32 = 586
    v3 := []uint32{v2, (v2 % (((v2 % ((v2) | 1))) | 1)), v2, v2}
    var v4 uint16 = 605
    var v5 int16 = 32767
    var v6 uint = 65535
    if true {
        for fi7 := 0; fi7 < 17; fi7++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi7
        }
        var v8 uint16 = v4
        var v9 uint = v6
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int16 = 32767
            _ = len(v3)
            _ = fi11
            _ = v12
        }
        var v13 int8 = 100
        var v14 uint16 = v4
        _ = fi10
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS1{F0: ("" + "foo bar"), F1: 1, F2: v2}
        var v17 uint8 = 42
        fmt.Println(v17)
        _ = fi15
        _ = v16
        _ = v17
    }
    if len(v3) > 0 {
        v3[0] = v2
    }
    v2 -= v2
    var v18 uint8 = 255
    switch "hello" {
    case "hello":
        v18 = v18
    default:
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(v1, v1)
            v20 := GenS2{F0: 9223372036854775807, F1: v6, F2: 637}
            if false {
                v6 = 0
                fmt.Println(v2, fi19, v1)
            }
            _ = fi19
            _ = v20
        }
    }
    var v21 bool = false
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 string = ""
        _ = fi22
        _ = v23
    }
    if true {
        v5 += (v5 | -1)
        if (!true) {
            fmt.Println(v2)
            v6 = (v6 % (((v6 - (v6 | (v6 & (v6 / ((1) | 1)))))) | 1))
            v24 := GenS3{F0: 42, F1: 42, F2: ("a" + "\n\t")}
            _ = v24
        } else {
            var v25 uint16 = v4
            var v26 int64 = 0
            var v27 int32 = 1
            _ = v25
            _ = v26
            _ = v27
        }
    }
    var v28 int16 = v5
    switch "\n\t" {
    case "":
        var v29 uint64 = 42
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
