package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint32 }
type GenS2 struct { F0 int; F1 uint64; F2 int }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int32 = -716
    v2 := make(map[string]string)
    v3 := []string{"hello", (("" + "\n\t") + ((("a" + "hello") + ("\n\t" + ("a" + (("a" + ("hello" + "foo bar")) + "\n\t")))) + ("hello" + "foo bar"))), ("hello" + ("" + ("" + ""))), ("" + ("\n\t" + "hello"))}
    var v4 float32 = float32(float64(3.14159))
    v5 := GenS1{F0: -789, F1: 100, F2: 2147483647}
    v4 = v4
    var v6 uint32 = 171
    if (32767 > 849) {
        v7 := GenS2{F0: 1, F1: 42, F2: -672}
        _ = v7
    }
    v1 = v1
    v6 = (v6 / ((4294967295) | 1))
    v2[((("" + "foo bar") + "\n\t") + "hello")] = ""
    v8 := GenS3{F0: 42}
    var v9 uint16 = 255
    for fi10 := 0; fi10 < 20; fi10++ {
        switch "hello" {
        case "a":
            v1 += v1
        case "hello":
            var v11 int16 = -32768
            _ = v11
        case "\n\t":
            v3 = append(v3, "\n\t")
        default:
            v3 = append(v3, (("a" + "\n\t") + ""))
        }
        _ = fi10
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 bool = ((v4 / v4) < v4)
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 int32 = (v1 * v1)
            _ = fi14
            _ = v15
        }
        v13 = (!(758 > 55))
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
