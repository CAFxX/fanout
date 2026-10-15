package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 bool) int16 {
    switch 1 {
    case -834:
        if (!true) {
            v5 := []int{-1, 32767, -396, 949}
            var v6 int = -1
            _ = v5
            _ = v6
        } else {
            p0 = (((0.0 / 0.1) - 0.0) != 1.5)
            var v7 float32 = (((float32(3.14159) + (-2.25 * 0.1)) + ((1e10 / (0.1 / -2.25)) / 1.5)) - ((float32(float64((0.0 + (1.5 + (3.14159 * (1.5 * 3.14159)))))) + -2.25) * float32(0.1)))
            p0 = true
            _ = v7
        }
    }
    var v8 float64 = -2.25
    for fi9 := 0; fi9 < 13; fi9++ {
        switch (("foo bar" + "\n\t") + "") {
        case "foo bar":
            v10 := []int{0}
            _ = v10
        case "hello":
            var v11 bool = (2147483647 >= -2147483648)
            _ = v11
        }
        _ = fi9
    }
    _ = p0
    _ = v8
    return -737
}
func genFn1(p0 uint32) int64 {
    switch "a" {
    case "hello":
        switch ("hello" + ((("\n\t" + "\n\t") + "\n\t") + ("\n\t" + (("hello" + "hello") + ("" + (("hello" + (("\n\t" + "hello") + "a")) + ("foo bar" + (("a" + (((((("\n\t" + "hello") + "foo bar") + "a") + "\n\t") + (("" + "\n\t") + "")) + ("hello" + "foo bar"))) + ("a" + "foo bar"))))))))) {
        case "":
            p0 = p0
        }
    }
    p0 -= (p0 - p0)
    _ = p0
    return -539
}

func main() {
    var v1 float64 = 0.1
    v2 := []string{("\n\t" + "foo bar"), ("foo bar" + ""), (("\n\t" + "\n\t") + "foo bar"), "a"}
    v3 := make(map[string]string)
    var v4 uint32 = 42
    v1 = v1
    v4 = 850
    v12 := GenS1{F0: 42}
    if false {
        v13 := v3[(("" + ("a" + "foo bar")) + "hello")]
        var v14 int64 = 1
        if (0 == -1) {
            var v15 uint64 = 42
            v14 -= (v14 - 726)
            v1 = (0.1 / 1.0)
            _ = v15
        } else {
            _ = len(v3)
            var v16 bool = (2 < uint16(v4))
            var v17 float64 = (v1 * float64(((0.0 - float32(v1)) * 1e10)))
            _ = v16
            _ = v17
        }
        _ = v13
        _ = v14
    } else {
        v1 = ((v1 * 0.0) - v1)
        var v18 float32 = -2.25
        v19 := GenS2{F0: 42, F1: uint64(v4)}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 int16 = 1
        var v22 float64 = 0.0
        _ = fi20
        _ = v21
        _ = v22
    }
    v4 -= v4
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v23 uint = uint(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
