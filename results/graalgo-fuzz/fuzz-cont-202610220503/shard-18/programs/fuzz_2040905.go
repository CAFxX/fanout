package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    var v1 string = (("hello" + (((("a" + "") + "\n\t") + "hello") + ("" + "hello"))) + ((((("hello" + "foo bar") + "foo bar") + "\n\t") + "\n\t") + "foo bar"))
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 uint32 = 2
    var v5 int = 1
    fmt.Println(v1, v4, v1)
    v6 := GenS1{F0: 32767}
    v5 = v5
    v7 := []int{(v5 | -2147483648), (-(v5)), v5}
    switch v1 {
    case "hello":
        var v8 int64 = -239
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        if (((1e10 / 1.5) - -2.25) != 0.1) {
            v10 := GenS1{F0: 453}
            v11 := GenS1{F0: 2}
            _ = v10
            _ = v11
        } else {
            v5 = v5
        }
        v4 += 2
        _ = fi9
    }
    fmt.Println(v1)
    v12 := GenS1{F0: 65535}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v13 := GenS2{F0: 187, F1: 964}
    switch (("foo bar" + ((v1 + "\n\t") + "hello")) + v1) {
    case "a":
        if false {
            fmt.Println(v4)
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint32 = (^(v4))
                var v16 uint16 = 0
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            v5 = (v5 ^ (v5 | (-(v5))))
            switch (("" + "a") + ("a" + (v1 + "hello"))) {
            case "":
                if len(v7) > 0 {
                    v7[0] = 1
                }
            case "a":
                v17 := GenS2{F0: 42, F1: 32767}
                _ = v17
            default:
                if true {
                    v18 := GenS1{F0: 65535}
                    var v19 int = v5
                    for fi20 := 0; fi20 < 13; fi20++ {
                        for fi21 := 0; fi21 < 11; fi21++ {
                            _ = len(v7)
                            _ = fi21
                        }
                        var v22 int32 = 42
                        _ = fi20
                        _ = v22
                    }
                    _ = v18
                    _ = v19
                } else {
                    fmt.Println(v1, v4, v1)
                }
            }
            v23 := GenS2{F0: 65535, F1: -1}
            _ = v23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
