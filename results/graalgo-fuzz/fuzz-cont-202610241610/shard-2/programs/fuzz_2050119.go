package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint16 }
func genFn0() uint8 {
    var v5 int16 = 0
    fmt.Println(v5)
    v5 += v5
    _ = v5
    return 204
}
func genFn1() int8 {
    var v6 uint16 = 0
    var v7 uint = 572
    _ = v6
    _ = v7
    return 100
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    _ = len(v4)
    v4[(("\n\t" + "hello") + ((("hello" + "") + "foo bar") + (("hello" + (("foo bar" + "hello") + "a")) + (("hello" + (("\n\t" + "foo bar") + ("a" + "foo bar"))) + "hello"))))] = 2147483647
    _ = len(v2)
    var v8 int64 = 9223372036854775807
    if false {
        var v9 uint = 154
        v10 := v3[0]
        _ = v10
        _ = v9
    }
    switch "hello" {
    case "hello":
        v8 += v8
    }
    var v11 int32 = -463
    switch "foo bar" {
    case "\n\t":
        var v12 float64 = 0.0
        _ = v12
    case "":
        v13 := []int{125, 4294967295}
        _ = v13
    case "hello":
        if (383 < 622) {
            var v14 float32 = 1.5
            v4[("\n\t" + "\n\t")] = 32768
            if (true && true) {
                v15 := []int{-9223372036854775808}
                _ = v15
            } else {
                var v16 uint8 = genFn0()
                _ = v16
            }
            _ = v14
        } else {
            var v17 int64 = 0
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v8, v8)
            }
            _ = v17
        }
    default:
        v18 := GenS1{F0: (!((v8 >> ((255) & 63)) > 0)), F1: 42, F2: 2}
        _ = v18
    }
    v8 = v8
    var v19 float64 = 1.5
    var v20 uint8 = 42
    var v21 float64 = float64((float32((0.0 - v19)) / -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v8
}
