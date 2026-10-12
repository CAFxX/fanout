package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    var v3 uint = 0
    v4 := GenS1{F0: (("" + "foo bar") + (("\n\t" + "\n\t") + "")), F1: 721, F2: 1}
    delete(v1, ((("" + "hello") + "hello") + ("foo bar" + (("\n\t" + ("foo bar" + "hello")) + "hello"))))
    switch 42 {
    case -855:
        v3 += (v3 >> ((517) & 63))
    }
    var v5 int64 = 0
    fmt.Println(v5)
    if false {
        fmt.Println(v5, v5)
    }
    v3 -= (v3 >> ((383) & 63))
    fmt.Println(v3)
    var v6 float32 = (0.0 + 1e10)
    if (!(!false)) {
        if true {
            for fi7 := 0; fi7 < 14; fi7++ {
                fmt.Println(v5, fi7)
                _ = fi7
            }
        }
    }
    _ = len(v1)
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: ("\n\t" + ("\n\t" + "a")), F1: 65535, F2: v5}
        switch "hello" {
        case "hello":
            v6 = (0.0 - v6)
        default:
            v6 = (v6 - 3.14159)
        }
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: "hello", F1: 1, F2: v5}
        _ = fi10
        _ = v11
    }
    v5 = (v5 * (v5 ^ v5))
    var v12 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
