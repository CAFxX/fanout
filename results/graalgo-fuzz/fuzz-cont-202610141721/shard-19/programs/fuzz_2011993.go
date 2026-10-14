package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int64 }

func main() {
    v1 := []uint32{}
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    var v4 float64 = 1.5
    v5 := make(map[string]uint32)
    if true {
        var v6 int = 4294967295
        _ = v6
    } else {
        v4 = (1.5 * (v4 - 1.5))
        fmt.Println(v4)
    }
    v4 += v4
    v4 = float64((0.0 + 1e10))
    var v7 uint = 63
    switch ((((((("foo bar" + ("\n\t" + (("hello" + ((((("a" + ("hello" + "hello")) + "a") + "") + ("foo bar" + "")) + "foo bar")) + "foo bar"))) + "foo bar") + "foo bar") + "\n\t") + "hello") + (("" + ("a" + "")) + ("hello" + "\n\t"))) + "a") {
    case "hello":
        var v8 uint32 = 42
        _ = v8
    default:
        v9 := GenS1{F0: "\n\t", F1: 127, F2: 777}
        _ = v9
    }
    v4 = 0.0
    v4 += float64(1.5)
    switch 128 {
    case 18:
        var v10 uint8 = 2
        _ = v10
    case -926:
        var v11 uint = (v7 & v7)
        _ = v11
    case 4294967295:
        v12 := GenS1{F0: ("" + "foo bar"), F1: 1, F2: 1}
        _ = v12
    default:
        var v13 float64 = v4
        _ = v13
    }
    var v14 int64 = 1
    var v15 int32 = -1
    if (true && (1 < 976)) {
        switch (("" + "a") + "\n\t") {
        case "\n\t":
            v16 := v5["foo bar"]
            _ = v16
        case "":
            if false {
                if (("hello" + ("foo bar" + (("foo bar" + "\n\t") + "hello"))) >= "a") {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        var v18 int64 = 9223372036854775807
                        for fi19 := 0; fi19 < 14; fi19++ {
                            var v20 int16 = -1
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                            _ = fi19
                            _ = v20
                        }
                        _ = fi17
                        _ = v18
                    }
                }
            }
        case "foo bar":
            v3[0] = false
        default:
            v15 = v15
        }
        var v21 int = -9223372036854775808
        _ = len(v2)
        _ = v21
    }
    if (!(!true)) {
        v14 += v14
        var v22 uint = v7
        _ = len(v3)
        _ = v22
    } else {
        v23 := v3[65535]
        v23 = false
        _ = v23
    }
    var v24 uint64 = 621
    if len(v1) > 0 {
        v1[0] = 42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
