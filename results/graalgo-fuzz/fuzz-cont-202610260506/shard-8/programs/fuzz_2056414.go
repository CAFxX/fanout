package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int8; F1 uint32; F2 int16 }

func main() {
    var v1 int8 = -54
    var v2 int8 = 1
    var v3 int32 = 1
    v4 := make(map[string]uint32)
    if false {
        fmt.Println(v2, v1)
        v5 := GenS1{F0: 0}
        v4["a"] = 1
        _ = v5
    }
    switch -681 {
    case -743:
        v6 := GenS1{F0: 0}
        _ = v6
    case 0:
        switch 32767 {
        case 42:
            var v7 string = "a"
            _ = v7
        case 1:
            v8 := []int{956, -674, 4294967295}
            _ = v8
        case 36:
            switch "hello" {
            case "a":
                v1 += v1
            default:
                fmt.Println(v2, v1)
            }
        }
    case -694:
        v1 += (v1 | 127)
    default:
        fmt.Println(v2)
    }
    var v9 int16 = 1000
    v3 = (v3 >> ((4294967295) & 63))
    v4[("" + (((("hello" + "a") + "foo bar") + ("a" + "\n\t")) + "\n\t"))] = 447
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "hello" {
        case "foo bar":
            fmt.Println(v2, v3, v3)
        case "":
            v11 := GenS1{F0: 738}
            _ = v11
        }
        fmt.Println(v1)
        _ = fi10
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(fi12, v3)
        v13 := GenS1{F0: 32767}
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS1{F0: 32767}
            v3 -= v3
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v16 := GenS2{F0: 198, F1: 65535}
    if true {
        if false {
            fmt.Println(v1, v1)
            v17 := GenS3{F0: v1, F1: 840, F2: -1}
            _ = v17
        }
    }
    v18 := GenS3{F0: v1, F1: 645, F2: (v9 & 520)}
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS3{F0: v2, F1: 931, F2: (v9 >> ((0) & 63))}
        _ = fi19
        _ = v20
    }
    switch "a" {
    case "hello":
        var v21 int16 = v9
        _ = v21
    }
    var v22 float32 = (-2.25 - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
