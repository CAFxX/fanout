package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint32 }
func genFn0(p0 bool) int8 {
    p0 = (!p0)
    var v6 int32 = 42
    v7 := GenS1{F0: 294}
    var v8 string = (("a" + "foo bar") + ((("hello" + "a") + "") + "a"))
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 100
}

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 42
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    var v9 int16 = -32768
    if true {
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(v2)
            _ = len(v5)
            _ = fi10
        }
        _ = len(v4)
        v1 = float32(-2.25)
    }
    v11 := GenS1{F0: -482}
    v2 += v2
    v12 := GenS1{F0: 42}
    delete(v5, (("hello" + "foo bar") + "foo bar"))
    v2 += (v2 + v2)
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 int64 = 42
            _ = fi14
            _ = v15
        }
        var v16 int16 = -293
        _ = fi13
        _ = v16
    }
    if true {
        delete(v5, ("" + "hello"))
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 float32 = v1
            var v19 float32 = v18
            v18 = float32(1e10)
            _ = fi17
            _ = v18
            _ = v19
        }
        v2 = 562
    } else {
        var v20 int32 = 42
        _ = v20
    }
    fmt.Println(v2, v9)
    if false {
        switch "a" {
        case "\n\t":
            for fi21 := 0; fi21 < 3; fi21++ {
                v2 -= (v2 * v2)
                v22 := GenS2{F0: 2}
                _ = fi21
                _ = v22
            }
        default:
            var v23 float32 = v1
            _ = v23
        }
    } else {
        fmt.Println(v9, v1)
        v24 := GenS2{F0: v2}
        _ = v24
    }
    fmt.Println(v2)
    v25 := GenS3{F0: -2147483648, F1: 168}
    v26 := GenS3{F0: 127, F1: 65535}
    for fi27 := 0; fi27 < 11; fi27++ {
        _ = len(v5)
        var v28 int = fi27
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        v5[(((("foo bar" + "foo bar") + "foo bar") + (("" + (("" + ("hello" + "a")) + ("hello" + (("a" + "a") + ((("a" + "") + ("foo bar" + ("" + "a"))) + ("\n\t" + "")))))) + ("a" + "\n\t"))) + "\n\t")] = "hello"
        _ = fi29
    }
    fmt.Println(v9, v9, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
