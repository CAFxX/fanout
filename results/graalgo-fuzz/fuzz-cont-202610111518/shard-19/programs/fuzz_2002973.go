package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 2
    var v3 uint = 18446744073709551615
    var v4 uint8 = 32
    fmt.Println(v3, v4)
    var v5 int16 = 386
    v6 := []int{}
    v5 = (v5 >> (((^(v3))) & 63))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if true {
        v5 = 42
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 679}
            fmt.Println(v3, v5)
            _ = fi7
            _ = v8
        }
        v1 = 0.0
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        if true {
            var v10 string = (("a" + "hello") + (("a" + "hello") + (("foo bar" + "foo bar") + "hello")))
            v11 := GenS1{F0: -1}
            var v12 float64 = v1
            _ = v10
            _ = v11
            _ = v12
        }
        if (false && (2 > 22)) {
            v2 = v2
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            var v13 uint8 = v4
            v13 = v13
            _ = v13
        }
        _ = fi9
    }
    switch "" {
    case "hello":
        if true {
            var v14 bool = ((0.0 + v1) >= ((-2.25 - (v1 + (float64((3.14159 + float32(v1))) + v1))) - -2.25))
            _ = v14
        } else {
            var v15 bool = true
            v3 += v3
            v1 = float64(1e10)
            _ = v15
        }
    case "\n\t":
        v16 := GenS1{F0: -1}
        _ = v16
    case "foo bar":
        var v17 string = (((("" + "a") + "foo bar") + "hello") + ("a" + "foo bar"))
        _ = v17
    default:
        var v18 uint32 = 0
        _ = v18
    }
    var v19 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
