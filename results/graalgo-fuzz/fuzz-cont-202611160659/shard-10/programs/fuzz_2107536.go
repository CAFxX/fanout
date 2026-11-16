package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
type GenS2 struct { F0 string; F1 bool; F2 uint }

func main() {
    var v1 string = "hello"
    v2 := []string{("hello" + ("hello" + v1)), "a", v1}
    var v3 float32 = -2.25
    if ("foo bar" > ((("hello" + v1) + v1) + "foo bar")) {
        v1 = v1
        if false {
            v1 = (v1 + v1)
            v1 = v1
        }
        var v4 uint = 27
        _ = v4
    } else {
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 int64 = -9223372036854775808
            if (true && false) {
                v3 = float32(0.1)
                var v7 int32 = 1
                var v8 int64 = v6
                _ = v7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
        if (!true) {
            for fi9 := 0; fi9 < 12; fi9++ {
                v3 = v3
                _ = fi9
            }
        }
    }
    v10 := GenS1{F0: "\n\t", F1: 42, F2: 0}
    v3 += (v3 * v3)
    var v11 float32 = float32(0.1)
    var v12 uint = 42
    if false {
        if true {
            v13 := GenS1{F0: "foo bar", F1: 32767, F2: 2147483647}
            var v14 float64 = (0.1 + (-2.25 / (-2.25 + (0.0 / (float64(v3) / -2.25)))))
            _ = v13
            _ = v14
        }
    }
    switch v1 {
    case "a":
        switch 127 {
        case 700:
            v15 := GenS1{F0: (v1 + ("\n\t" + (v1 + "a"))), F1: 42, F2: 0}
            _ = v15
        }
    case "":
        v1 = v1
    }
    if true {
        switch v1 {
        case "hello":
            v16 := GenS1{F0: "foo bar", F1: 0, F2: 107}
            _ = v16
        case "\n\t":
            fmt.Println(v12)
        default:
            v3 += -2.25
        }
        switch 1 {
        case 32768:
            v3 = 3.14159
        default:
            var v17 int16 = -771
            _ = v17
        }
        v1 = (((v1 + ("\n\t" + "\n\t")) + "") + "a")
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v11 = (3.14159 / (v3 + 1e10))
        v19 := GenS2{F0: v1, F1: true, F2: 945}
        _ = fi18
        _ = v19
    }
    v1 = "hello"
    var v20 bool = false
    switch "foo bar" {
    case "hello":
        switch ((("" + "") + "foo bar") + "") {
        case "":
            if len(v2) > 0 {
                v2[0] = v1
            }
        case "a":
            var v21 int8 = -68
            _ = v21
        case "foo bar":
            v3 = v11
        default:
            fmt.Println(v3)
        }
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = v1
        }
    case "foo bar":
        var v22 int16 = 0
        _ = v22
    default:
        for fi23 := 0; fi23 < 11; fi23++ {
            v3 -= (v3 * float32(-2.25))
            var v24 uint8 = 1
            _ = fi23
            _ = v24
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
}
