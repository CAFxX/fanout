package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
func genFn0(p0 string) int8 {
    if (!false) {
        var v4 uint8 = 42
        _ = v4
    } else {
        fmt.Println(p0)
        switch p0 {
        case "":
            fmt.Println(p0)
        case "hello":
            switch "foo bar" {
            case "foo bar":
                for fi5 := 0; fi5 < 1; fi5++ {
                    var v6 string = ("hello" + (p0 + "\n\t"))
                    var v7 float32 = 0.0
                    fmt.Println(p0, p0, v6)
                    _ = fi5
                    _ = v6
                    _ = v7
                }
            }
        case "a":
            var v8 uint8 = 5
            _ = v8
        default:
            var v9 bool = true
            _ = v9
        }
        var v10 uint16 = 620
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        if false {
            for fi12 := 0; fi12 < 16; fi12++ {
                p0 = p0
                _ = fi12
            }
            fmt.Println(p0)
        }
        p0 = (((p0 + "a") + "\n\t") + p0)
        _ = fi11
    }
    _ = p0
    return -1
}
func genFn1(p0 string, p1 int32, p2 bool) uint16 {
    v13 := []int{-986, 2147483647}
    _ = len(v13)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 42
}

func main() {
    var v1 int64 = -96
    var v2 int = -1
    var v3 uint = 914
    var v14 int64 = int64(v2)
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: 0, F1: genFn1("\n\t", 0, true)}
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(v14)
            v3 = 4294967295
            v18 := []int{-2147483648}
            _ = fi17
            _ = v18
        }
        fmt.Println(v2)
        _ = fi15
        _ = v16
    }
    switch "hello" {
    case "a":
        v19 := GenS1{F0: 127, F1: 790}
        _ = v19
    default:
        var v20 float32 = 1.5
        _ = v20
    }
    v1 = (v1 - v1)
    var v21 uint16 = 0
    v21 -= genFn1((("hello" + ("" + "\n\t")) + "a"), -2147483648, false)
    if ((float32(float64(3.14159)) / 1.0) > -2.25) {
        v22 := []int{-1, 1, -34}
        _ = v22
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v14, v2, v1)
    case "foo bar":
        var v23 uint32 = 0
        _ = v23
    case "hello":
        v3 = v3
    default:
        var v24 int = v2
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
}
