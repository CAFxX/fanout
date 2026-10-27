package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = ((1.5 + ((0.0 - 1.5) / 1.5)) * 1.5)
    var v3 int32 = 0
    v3 = -933
    switch ((("foo bar" + (("\n\t" + "") + ((("" + "foo bar") + "hello") + "hello"))) + "") + "hello") {
    case "\n\t":
        v4 := GenS1{F0: ("foo bar" + "")}
        _ = v4
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int8 = int8(fi5)
        _ = fi5
        _ = v6
    }
    fmt.Println(v2)
    v3 += v3
    if (-1 > 127) {
        var v7 int16 = 1
        _ = v7
    } else {
        if ((2 != 67) && false) {
            var v8 bool = (474 == 4294967295)
            v2 = (v2 - float64(float32(1.5)))
            _ = v8
        } else {
            v2 += -2.25
        }
    }
    switch "\n\t" {
    case "":
        v3 = (^(v3))
    case "hello":
        v9 := []int{}
        _ = v9
    case "foo bar":
        v10 := []int{-502, -673}
        _ = v10
    }
    if false {
        v11 := GenS2{F0: "foo bar", F1: -9223372036854775808}
        switch "foo bar" {
        case "foo bar":
            var v12 bool = (!(18446744073709551615 >= 0))
            _ = v12
        case "a":
            var v13 bool = (137 > 716)
            _ = v13
        default:
            if false {
                v14 := GenS3{F0: 127, F1: (v3 - (v3 >> ((65535) & 63)))}
                v2 = -2.25
                _ = v14
            }
        }
        _ = v11
    }
    var v15 int8 = int8(v3)
    fmt.Println(v2, v3, v3)
    v2 += 3.14159
    v2 += (float64(((3.14159 - (3.14159 * (0.1 + (1e10 * (3.14159 * (3.14159 / 1.0)))))) * 1.5)) - v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
}
