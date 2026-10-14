package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint = 255
    v2 := make(map[int]int)
    var v3 float32 = 1e10
    v4 := []string{}
    if ((42 >= 974) && ((true && true) || ((^(v1)) > v1))) {
        for fi5 := 0; fi5 < 1; fi5++ {
            fmt.Println(v1, v1)
            _ = fi5
        }
    }
    var v6 uint64 = 290
    v1 = v1
    v6 -= 2
    if false {
        v1 -= 18446744073709551615
        v6 = v6
    }
    if true {
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(v3)
            _ = fi7
        }
        v3 -= v3
    } else {
        var v8 uint64 = v6
        _ = v8
    }
    switch ((("foo bar" + "hello") + "foo bar") + "a") {
    case "\n\t":
        v4 = append(v4, "\n\t")
    case "":
        for fi9 := 0; fi9 < 3; fi9++ {
            switch "\n\t" {
            case "hello":
                v1 = (v1 ^ v1)
            case "foo bar":
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := GenS1{F0: 18446744073709551615}
                    _ = fi10
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 2; fi12++ {
                    switch "a" {
                    case "hello":
                        v6 += 377
                    }
                    v3 -= 0.0
                    _ = fi12
                }
            }
            v4 = append(v4, "")
            switch fi9 {
            case 0:
                _ = len(v4)
            case 42:
                if false {
                    var v13 int16 = -962
                    _ = v13
                }
            }
            _ = fi9
        }
    }
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    var v14 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
