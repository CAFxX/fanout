package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int16 }

func main() {
    var v1 int32 = 212
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 float64 = (-2.25 * (0.0 - 0.0))
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            switch ("a" + ((("hello" + "a") + "a") + ((("" + "foo bar") + ("" + "\n\t")) + "hello"))) {
            case "\n\t":
                fmt.Println(v4, fi5, v4)
            case "a":
                v1 -= 0
            }
            _ = fi6
        }
        _ = fi5
    }
    var v7 int64 = 42
    v4 += v4
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: 1}
        v10 := GenS1{F0: 0}
        var v11 int8 = -111
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    if false {
        for fi12 := 0; fi12 < 4; fi12++ {
            fmt.Println(fi12, v1, v1)
            _ = fi12
        }
        switch "\n\t" {
        case "\n\t":
            v4 -= v4
        case "foo bar":
            switch "a" {
            case "":
                var v13 int16 = -659
                _ = v13
            case "a":
                for fi14 := 0; fi14 < 4; fi14++ {
                    if false {
                        v7 = (v7 - v7)
                    }
                    _ = fi14
                }
            default:
                fmt.Println(v4, v1)
            }
        case "":
            v15 := []int{836}
            _ = v15
        }
        fmt.Println(v7, v4, v4)
    }
    v7 += (v7 >> ((1) & 63))
    for fi16 := 0; fi16 < 13; fi16++ {
        delete(v3, "foo bar")
        if (!true) {
            v17 := GenS2{F0: 42}
            v7 -= v7
            _ = v17
        }
        _ = fi16
    }
    v18 := GenS3{F0: (v1 % ((-2147483648) | 1)), F1: 1}
    var v19 int64 = 0
    fmt.Println(v7, v4)
    v20 := GenS2{F0: 304}
    var v21 bool = true
    fmt.Println(v1, v19, v21)
    v22 := v3["\n\t"]
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 int16 = 806
        v22 = -1
        _ = fi23
        _ = v24
    }
    var v25 uint8 = 76
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
