package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int32; F1 int; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 bool) int32 {
    var v5 int32 = 220
    p0 = (!(0 >= 9223372036854775807))
    if p0 {
        if p0 {
            p0 = true
            var v6 int64 = -100
            _ = v6
        } else {
            v7 := GenS1{F0: 1000, F1: "foo bar"}
            var v8 int16 = -1
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = v5
    return v5
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]bool)
    var v3 int16 = 1000
    var v4 int8 = 0
    var v9 int8 = (v4 / ((v4) | 1))
    v3 = v3
    v10 := GenS1{F0: v3, F1: ("foo bar" + "foo bar")}
    v3 = (v3 - (^(v3)))
    v1 -= (^(v1))
    switch 911 {
    case 32768:
        fmt.Println(v4)
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            delete(v2, "hello")
            _ = fi12
        }
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v1, v3, v9)
        case "hello":
            v1 += 42
        case "":
            v3 -= (v3 & (-(v3)))
        }
        v3 = 42
        _ = fi11
    }
    switch 9223372036854775807 {
    case -815:
        fmt.Println(v9)
    case 128:
        v4 -= (v9 ^ v4)
    case 0:
        v4 += v4
    default:
        v13 := GenS2{F0: 1, F1: -1, F2: false}
        _ = v13
    }
    v14 := GenS3{F0: 667}
    fmt.Println(v1)
    v4 -= (v9 ^ (-(v9)))
    fmt.Println(v9, v3)
    var v15 int8 = v9
    v3 -= v3
    switch 2147483647 {
    case 590:
        v16 := GenS4{F0: 640}
        _ = v16
    case -9223372036854775808:
        if false {
            switch 1 {
            case 127:
                var v17 uint = 0
                _ = v17
            default:
                var v18 string = (((("hello" + (("hello" + "foo bar") + ("hello" + ("\n\t" + ("hello" + "\n\t"))))) + "a") + ("foo bar" + "a")) + "hello")
                _ = v18
            }
        } else {
            for fi19 := 0; fi19 < 2; fi19++ {
                v15 = (v4 * (v15 ^ (v15 >> ((216) & 63))))
                _ = fi19
            }
        }
    }
    v20 := GenS2{F0: -2147483648, F1: 42, F2: true}
    switch "\n\t" {
    case "":
        var v21 int = 4294967295
        _ = v21
    case "\n\t":
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := []int{337}
            v3 -= (v3 - (v3 - v3))
            v24 := GenS1{F0: (v3 | 483), F1: "a"}
            _ = fi22
            _ = v23
            _ = v24
        }
    default:
        fmt.Println(v9, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
