package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 string }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 float64 = (1.5 + 3.14159)
    v5 := v3[-178]
    var v6 int64 = 669
    v1 += -2.25
    switch "foo bar" {
    case "a":
        fmt.Println(v1)
    case "":
        v7 := GenS1{F0: 128}
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int = 32767
            _ = fi8
            _ = v9
        }
    }
    if (false && false) {
        v5 += (v5 >> ((687) & 63))
        v1 = v1
    } else {
        v4 = v4
    }
    v10 := []int{0, 32767, -1}
    switch 32768 {
    case -65:
        var v11 uint64 = 1
        _ = v11
    case -9223372036854775808:
        v4 = (-2.25 - v4)
    }
    v12 := GenS2{F0: 2147483647}
    if false {
        switch "hello" {
        case "":
            if true {
                v13 := GenS1{F0: 0}
                var v14 int = 9223372036854775807
                _ = v13
                _ = v14
            }
        default:
            v6 += (v6 >> ((65535) & 63))
        }
        if false {
            _ = len(v10)
            var v15 bool = false
            var v16 string = ""
            _ = v15
            _ = v16
        } else {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            switch ("foo bar" + ("" + (("hello" + "hello") + "\n\t"))) {
            case "":
                v17 := GenS3{F0: 65, F1: ("a" + ((("foo bar" + "a") + "") + "foo bar"))}
                _ = v17
            }
        }
        switch -107 {
        case 42:
            var v18 uint64 = 2
            _ = v18
        case 1:
            var v19 int64 = v6
            _ = v19
        }
    }
    v5 -= v5
    _ = len(v10)
    var v20 int16 = 757
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        fmt.Println(fi21)
        fmt.Println(v4, v1)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
