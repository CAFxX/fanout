package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint64) int8 {
    p0 = p0
    p0 = 1
    var v4 string = ("foo bar" + "a")
    if false {
        v5 := []int{}
        _ = v5
    }
    _ = p0
    _ = v4
    return 36
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 int8 = -80
    v6 := GenS1{F0: true, F1: 42}
    switch ("a" + ((("foo bar" + "foo bar") + ((((("\n\t" + "\n\t") + "hello") + ((("\n\t" + (("foo bar" + "foo bar") + "hello")) + ("a" + "\n\t")) + "")) + (((("" + ("" + ("" + ("" + (("foo bar" + "foo bar") + "\n\t"))))) + (("hello" + (("hello" + (("\n\t" + "hello") + ("hello" + ("" + (("a" + "") + "\n\t"))))) + "")) + "foo bar")) + "a") + "a")) + "\n\t")) + "hello")) {
    case "\n\t":
        switch "\n\t" {
        case "a":
            fmt.Println(v3)
        case "\n\t":
            v3 -= v3
        case "foo bar":
            fmt.Println(v3)
        default:
            v3 -= (v3 / (((v3 ^ v3)) | 1))
        }
    }
    var v7 float32 = 1e10
    v8 := []int{}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v9 := GenS2{F0: 0, F1: 2147483647}
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint = 444
        v12 := GenS3{F0: -23}
        _ = fi10
        _ = v11
        _ = v12
    }
    v7 -= ((v7 - -2.25) * 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
