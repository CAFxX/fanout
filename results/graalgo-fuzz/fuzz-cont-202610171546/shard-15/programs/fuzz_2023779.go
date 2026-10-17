package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 0
    var v3 float64 = 1e10
    var v4 uint32 = 1
    v5 := []uint32{(v4 * (v4 % ((v4) | 1))), v4, 1, v4}
    var v6 int64 = -1
    switch "a" {
    case "\n\t":
        var v7 int64 = 0
        _ = v7
    case "":
        v8 := GenS1{F0: 42, F1: -439}
        _ = v8
    case "a":
        v3 = v3
    default:
        switch ("a" + (("\n\t" + "\n\t") + "hello")) {
        case "hello":
            var v9 int = 128
            _ = v9
        case "\n\t":
            var v10 float64 = v3
            _ = v10
        case "a":
            if len(v5) > 0 {
                v5[0] = 0
            }
        }
    }
    var v11 int16 = -32768
    v6 = (v6 % (((v6 * v6)) | 1))
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS1{F0: 757, F1: -1}
        fmt.Println(v3, v3, v2)
        _ = fi12
        _ = v13
    }
    switch ((("" + (((((("foo bar" + "") + (("foo bar" + "hello") + "a")) + "\n\t") + "foo bar") + ((("a" + "hello") + ("foo bar" + (((("a" + "a") + ("foo bar" + (("foo bar" + "\n\t") + "hello"))) + "") + "foo bar"))) + "hello")) + ("a" + ("hello" + ("\n\t" + "foo bar"))))) + "") + "a") {
    case "hello":
        var v14 int16 = 32767
        _ = v14
    case "\n\t":
        for fi15 := 0; fi15 < 5; fi15++ {
            v4 = v4
            _ = len(v1)
            var v16 int64 = v6
            _ = fi15
            _ = v16
        }
    default:
        var v17 uint32 = (v4 / (((^(v4))) | 1))
        _ = v17
    }
    fmt.Println(v11, v2)
    _ = len(v1)
    v6 += (v6 | (-(v6)))
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := GenS1{F0: 0, F1: 384}
            var v21 int32 = -2147483648
            _ = fi19
            _ = v20
            _ = v21
        }
        fmt.Println(v2)
        _ = fi18
    }
    delete(v1, 32767)
    switch "foo bar" {
    case "foo bar":
        var v22 uint = 2
        _ = v22
    case "hello":
        for fi23 := 0; fi23 < 7; fi23++ {
            v2 = v2
            v6 -= v6
            _ = fi23
        }
    default:
        switch int(v6) {
        case 1:
            v24 := GenS2{F0: v2}
            _ = v24
        default:
            var v25 string = "\n\t"
            _ = v25
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
