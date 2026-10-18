package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    var v1 uint8 = 42
    var v2 string = "a"
    v3 := []uint32{454}
    _ = len(v3)
    fmt.Println(v1, v1)
    switch "hello" {
    case "hello":
        var v4 uint = 4294967295
        _ = v4
    case "a":
        for fi5 := 0; fi5 < 15; fi5++ {
            fmt.Println(v2, fi5, v2)
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 int8 = 15
                fmt.Println(v1, v7)
                _ = fi6
                _ = v7
            }
            fmt.Println(v2, v2)
            _ = fi5
        }
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: v1}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 bool = (!false)
        for fi12 := 0; fi12 < 9; fi12++ {
            v3 = append(v3, 65535)
            v11 = (949 >= (fi12 | (fi12 >> ((996) & 63))))
            v11 = true
            _ = fi12
        }
        v13 := GenS2{F0: -40, F1: 32767}
        _ = fi10
        _ = v11
        _ = v13
    }
    var v14 uint64 = 2
    fmt.Println(v14)
    for fi15 := 0; fi15 < 17; fi15++ {
        if false {
            _ = len(v3)
        } else {
            v14 = (v14 & (v14 / ((v14) | 1)))
            v1 = v1
            switch "\n\t" {
            case "hello":
                v14 -= (v14 * v14)
            case "":
                _ = len(v3)
            case "foo bar":
                fmt.Println(v14)
            default:
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 float32 = -2.25
                    if len(v3) > 0 {
                        v3[0] = 65535
                    }
                    _ = fi16
                    _ = v17
                }
            }
        }
        fmt.Println(v1)
        _ = fi15
    }
    switch 127 {
    case 779:
        var v18 float32 = (0.1 / (1e10 - 3.14159))
        _ = v18
    case 127:
        var v19 uint8 = v1
        _ = v19
    default:
        fmt.Println(v14, v14)
    }
    var v20 string = "hello"
    if len(v3) > 0 {
        v3[0] = uint32(v1)
    }
    var v21 string = ("a" + v2)
    var v22 int16 = 32767
    v20 = v21
    fmt.Println(v1)
    switch -1 {
    case 42:
        for fi23 := 0; fi23 < 3; fi23++ {
            switch "hello" {
            case "a":
                var v24 string = v2
                _ = v24
            default:
                v3 = append(v3, 1)
            }
            if true {
                for fi25 := 0; fi25 < 13; fi25++ {
                    fmt.Println(v22)
                    fmt.Println(v1, v14)
                    _ = fi25
                }
            } else {
                var v26 int64 = 9223372036854775807
                _ = v26
            }
            switch fi23 {
            case 128:
                v2 = v2
            case 0:
                v27 := GenS2{F0: 1, F1: 1}
                _ = v27
            case 970:
                if true {
                    var v28 int8 = -1
                    fmt.Println(fi23, fi23)
                    _ = v28
                }
            }
            _ = fi23
        }
    default:
        fmt.Println(v22)
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        _ = len(v3)
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
