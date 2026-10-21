package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }

func main() {
    var v1 int64 = -1
    var v2 uint16 = 310
    v3 := []string{"\n\t", ("\n\t" + ("foo bar" + "foo bar")), "\n\t", (("a" + "a") + "")}
    var v4 int16 = -881
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 float32 = -2.25
        fmt.Println(v1, fi5)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: ("foo bar" + ("foo bar" + "hello")), F1: 389}
    for fi8 := 0; fi8 < 17; fi8++ {
        v1 = v1
        _ = fi8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 += (v1 * (v1 - v1))
    switch ("\n\t" + "hello") {
    case "":
        switch "\n\t" {
        case "a":
            if true {
                var v9 int16 = (v4 / ((-934) | 1))
                _ = v9
            }
        case "":
            v2 = v2
        case "\n\t":
            var v10 float64 = 0.0
            _ = v10
        }
    }
    v4 -= (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
