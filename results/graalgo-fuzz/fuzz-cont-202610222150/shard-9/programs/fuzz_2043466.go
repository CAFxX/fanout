package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int16 }

func main() {
    v1 := []uint32{1, 2147483648, 1, 4294967295}
    v2 := make(map[uint32]bool)
    var v3 uint = 1
    var v4 float32 = float32(0.1)
    v5 := make(map[int]string)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        if false {
            fmt.Println(v4, v3)
            if (!(!false)) {
                fmt.Println(v3)
            }
            fmt.Println(fi6, v3)
        }
        _ = fi6
    }
    switch 32767 {
    case -763:
        v3 = 65535
    case 4294967295:
        var v7 string = "hello"
        _ = v7
    case 128:
        v3 = (^(v3))
    }
    v8 := GenS1{F0: -1}
    _ = len(v2)
    _ = len(v5)
    switch "\n\t" {
    case "foo bar":
        switch ("a" + "a") {
        case "foo bar":
            v9 := GenS2{F0: 42, F1: -32768}
            _ = v9
        case "hello":
            v4 -= (v4 - (float32(3.14159) + 1e10))
        case "":
            v10 := GenS1{F0: -128}
            _ = v10
        }
    }
    v11 := GenS1{F0: 42}
    fmt.Println(v3, v4)
    switch "a" {
    case "foo bar":
        if true {
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 int64 = 371
                delete(v2, 0)
                var v14 bool = false
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 float32 = (((v4 * 0.1) - ((float32(((((-2.25 - 0.1) * 3.14159) / (((float64(v4) / 1.0) - 1.5) * -2.25)) * 1.5)) - v4) + -2.25)) + (0.0 * v4))
            _ = v15
        }
    }
    var v16 int32 = 2147483647
    if true {
        var v17 float64 = 1e10
        _ = v17
    }
    v3 = 2
    var v18 uint32 = 42
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 int8 = 0
        _ = fi19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
