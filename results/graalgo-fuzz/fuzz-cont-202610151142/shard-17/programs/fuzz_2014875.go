package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint8) int64 {
    switch 42 {
    case 128:
        switch "hello" {
        case "":
            p0 = (p0 & 0)
        case "hello":
            p0 += 128
        }
    case -412:
        v7 := GenS1{F0: 557, F1: 2}
        _ = v7
    }
    p0 = 85
    var v8 float64 = (((float64(0.0) * 0.0) / float64(((((3.14159 + 0.1) + 0.0) * -2.25) - 0.1))) + 1e10)
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint64 = 1
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := []string{("" + (("a" + ("a" + (((("\n\t" + ("foo bar" + "a")) + "") + ("hello" + ("a" + "foo bar"))) + ("foo bar" + (("foo bar" + "") + (("foo bar" + "foo bar") + "")))))) + ("foo bar" + ("" + ((("a" + "foo bar") + "") + "")))))}
    var v2 bool = (false && true)
    v3 := make(map[int]uint32)
    var v4 float32 = 1.5
    v5 := []string{("" + ("hello" + "a")), ("a" + "hello"), "foo bar", ("hello" + "foo bar")}
    v6 := make(map[uint32]uint32)
    if true {
        if ((((3.14159 - float64(v4)) + (0.1 - -2.25)) + 3.14159) >= (0.0 / ((0.1 - 0.0) / ((-2.25 + 1.5) + 3.14159)))) {
            var v11 uint32 = 2147483647
            fmt.Println(v4, v4, v4)
            _ = v11
        } else {
            var v12 int32 = 0
            v12 = (v12 % (((-(v12))) | 1))
            _ = v12
        }
    } else {
        fmt.Println(v4)
    }
    v4 += v4
    v1 = append(v1, "hello")
    var v13 uint16 = 42
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 int16 = 1000
        v13 = v13
        _ = fi14
        _ = v15
    }
    if (!false) {
        fmt.Println(v4)
        v2 = ((!(true || true)) || (false && ((32767 > 9223372036854775807) && v2)))
        var v16 int64 = 9223372036854775807
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 uint8 = 30
        var v19 uint64 = 0
        v2 = (v18 != v18)
        _ = fi17
        _ = v18
        _ = v19
    }
    v2 = (!(314 > 815))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
