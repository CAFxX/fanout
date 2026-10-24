package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int = 1
    v2 := []string{"\n\t", ""}
    var v3 float32 = -2.25
    var v4 uint16 = 315
    var v5 int32 = 1
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: 0}
    if ("foo bar" < "") {
        switch "a" {
        case "\n\t":
            v8 := GenS2{F0: 611, F1: 0, F2: 459}
            _ = v8
        case "":
            v3 = (float32(((float64(v3) - 1.5) / 1.5)) + (((float32(1.5) * float32(1.5)) * v3) - (float32(0.1) * float32(0.1))))
        default:
            fmt.Println(v5)
        }
        var v9 int = v1
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        _ = len(v6)
        var v11 bool = (158 > -32768)
        if v11 {
            _ = len(v6)
        }
        _ = fi10
        _ = v11
    }
    var v12 int32 = (v5 - 615)
    if true {
        for fi13 := 0; fi13 < 14; fi13++ {
            if (!(true && ((!true) || (!false)))) {
                if ("" < ((((("" + "foo bar") + ("foo bar" + ("\n\t" + ("a" + "hello")))) + ((("foo bar" + (("\n\t" + "hello") + ("hello" + ("hello" + "a")))) + ("a" + "a")) + "")) + "\n\t") + "\n\t")) {
                    v6["\n\t"] = 213
                }
                var v14 int8 = 127
                v14 = v14
                _ = v14
            }
            _ = fi13
        }
        fmt.Println(v5, v3)
        v6[""] = 4294967295
    }
    var v15 uint = 817
    v6["hello"] = 0
    for fi16 := 0; fi16 < 4; fi16++ {
        v6["\n\t"] = 346
        v17 := GenS3{F0: v4}
        _ = fi16
        _ = v17
    }
    v1 = (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
