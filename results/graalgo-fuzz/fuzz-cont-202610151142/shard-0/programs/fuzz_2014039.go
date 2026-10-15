package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    v1 = (-(v1))
    var v4 uint8 = 127
    v1 -= v1
    v5 := GenS1{F0: "foo bar", F1: "foo bar", F2: "\n\t"}
    delete(v2, -894)
    v6 := GenS1{F0: ("foo bar" + ""), F1: ("\n\t" + "hello"), F2: ""}
    fmt.Println(v1, v1)
    v7 := GenS1{F0: (("" + "a") + ""), F1: ("" + ("\n\t" + "hello")), F2: "foo bar"}
    v8 := []int{}
    v9 := GenS1{F0: "foo bar", F1: (("foo bar" + "") + ("foo bar" + "foo bar")), F2: ("" + "a")}
    fmt.Println(v4, v4)
    for fi10 := 0; fi10 < 13; fi10++ {
        if true {
            switch "foo bar" {
            case "foo bar":
                var v11 float64 = float64(3.14159)
                _ = v11
            case "hello":
                fmt.Println(v1, fi10, v4)
            }
            v1 = -651
            v4 -= (v4 * v4)
        }
        switch "foo bar" {
        case "foo bar":
            var v12 float32 = (1e10 + (3.14159 * (1e10 / 1.5)))
            _ = v12
        case "a":
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 uint16 = 32767
                _ = fi13
                _ = v14
            }
        }
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
