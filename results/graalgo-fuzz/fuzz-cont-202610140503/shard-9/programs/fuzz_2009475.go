package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint16 }

func main() {
    v1 := []string{}
    var v2 uint16 = 42
    var v3 uint = 65535
    v4 := []int{9223372036854775807, 255, -180, 685}
    var v5 int64 = 42
    var v6 uint8 = 127
    switch ("a" + ("\n\t" + "hello")) {
    case "\n\t":
        if true {
            for fi7 := 0; fi7 < 19; fi7++ {
                _ = len(v1)
                _ = fi7
            }
        } else {
            var v8 string = "\n\t"
            v8 = v8
            if false {
                fmt.Println(v6)
                var v9 int32 = 42
                _ = v9
            }
            _ = v8
        }
    case "a":
        v5 = v5
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v10 uint64 = 42
        _ = v10
    }
    v5 += (-(v5))
    var v11 int = 143
    for fi12 := 0; fi12 < 20; fi12++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi12
    }
    v13 := GenS1{F0: v3, F1: -27, F2: v2}
    var v14 float64 = (1.5 + (1e10 - -2.25))
    _ = len(v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v14 = -2.25
    var v15 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
