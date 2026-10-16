package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 bool }

func main() {
    var v1 uint8 = 2
    var v2 float64 = 1.5
    var v3 float64 = (v2 - float64(0.1))
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    var v6 uint = 541
    v7 := GenS1{F0: 100}
    for fi8 := 0; fi8 < 15; fi8++ {
        switch (("" + "") + "") {
        case "\n\t":
            var v9 int = 595
            _ = v9
        case "":
            v6 = 1
        case "a":
            v2 = float64(-2.25)
        default:
            fmt.Println(v1, v3)
        }
        var v10 uint16 = 255
        v11 := []int{int(v10), 127, (fi8 >> (((^(v6))) & 63))}
        _ = fi8
        _ = v10
        _ = v11
    }
    var v12 int32 = int32(v6)
    if (!true) {
        if true {
            v12 = (v12 % ((-1) | 1))
        } else {
            var v13 uint64 = 1
            _ = v13
        }
    }
    v1 = v1
    if false {
        if false {
            v3 = v2
            fmt.Println(v2, v1)
            var v14 int16 = 32767
            _ = v14
        }
        switch 1 {
        case 810:
            var v15 int32 = (^(v12))
            _ = v15
        default:
            v16 := GenS1{F0: -1}
            _ = v16
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            v3 = (-2.25 * 0.0)
            var v18 float64 = v2
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS2{F0: 645, F1: (!((("hello" + "\n\t") + "") < "foo bar"))}
        v5[""] = "\n\t"
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
