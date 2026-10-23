package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int16) int8 {
    v4 := GenS1{F0: -1}
    v5 := []int{127, 127, 425, 874}
    if false {
        v6 := GenS1{F0: 9223372036854775807}
        if true {
            if true {
                switch -713 {
                case 128:
                    var v7 float64 = (-2.25 / (1e10 - 1.5))
                    _ = v7
                default:
                    fmt.Println(p0)
                }
                p0 += p0
            }
            switch 42 {
            case 547:
                var v8 uint64 = 622
                _ = v8
            case 2147483647:
                switch (("foo bar" + "") + ((((((("\n\t" + "") + "a") + "\n\t") + ("a" + ("hello" + ("foo bar" + ("hello" + "foo bar"))))) + ("" + "a")) + "hello") + "hello")) {
                case "foo bar":
                    switch ((("foo bar" + "hello") + "\n\t") + "") {
                    case "hello":
                        var v9 int = 1
                        _ = v9
                    case "":
                        p0 = (p0 & p0)
                    }
                case "\n\t":
                    fmt.Println(p0)
                default:
                    switch "" {
                    case "a":
                        var v10 int32 = 989
                        _ = v10
                    case "\n\t":
                        v11 := GenS1{F0: 1}
                        _ = v11
                    case "":
                        var v12 int8 = -128
                        _ = v12
                    }
                }
            case -830:
                p0 = (^(p0))
            default:
                switch ("a" + "hello") {
                case "\n\t":
                    for fi13 := 0; fi13 < 19; fi13++ {
                        var v14 bool = ((!((fi13 | (fi13 % (((fi13 / (((fi13 - -453)) | 1))) | 1))) > (fi13 >> ((65535) & 63)))) && false)
                        v14 = (255 < 65535)
                        _ = fi13
                        _ = v14
                    }
                }
            }
            switch ((("a" + "foo bar") + ("a" + "a")) + ("hello" + ((("\n\t" + "foo bar") + ("foo bar" + ("" + "a"))) + "\n\t"))) {
            case "a":
                fmt.Println(p0)
            case "":
                var v15 bool = false
                _ = v15
            default:
                p0 -= 1000
            }
        } else {
            v16 := GenS2{F0: -1}
            _ = v16
        }
        p0 -= (p0 % ((1000) | 1))
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return 25
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v17 := []int{9223372036854775807, -2147483648}
    delete(v2, 20)
    var v18 uint64 = 537
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS2{F0: -1}
        _ = fi19
        _ = v20
    }
    v18 = v18
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 float32 = -2.25
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 9223372036854775807}
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := v2[994]
        var v26 uint32 = 1
        switch (v25 + "hello") {
        case "foo bar":
            v26 = v26
        case "\n\t":
            fmt.Println(v18, v25, v25)
        default:
            _ = len(v2)
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    v18 += v18
    v18 -= 42
    var v27 int8 = -30
    switch "hello" {
    case "foo bar":
        var v28 int = 245
        _ = v28
    }
    _ = len(v2)
    _ = len(v17)
    for fi29 := 0; fi29 < 16; fi29++ {
        v18 = 42
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v27
    _ = v3
}
