package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint32; F2 int64 }

func main() {
    var v1 float32 = 0.1
    v2 := []string{(("" + "foo bar") + ("hello" + "a")), ("hello" + "")}
    v3 := make(map[int]bool)
    v4 := []string{("\n\t" + ""), "\n\t"}
    v5 := make(map[uint32]int)
    if true {
        fmt.Println(v1)
        if false {
            switch ("" + ("foo bar" + ("a" + "foo bar"))) {
            case "a":
                var v6 string = ("hello" + ("a" + "foo bar"))
                _ = v6
            default:
                if (((253 != -1) || true) || false) {
                    for fi7 := 0; fi7 < 17; fi7++ {
                        var v8 uint64 = 0
                        var v9 uint32 = 4294967295
                        var v10 int8 = 38
                        _ = fi7
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                }
            }
            v11 := GenS1{F0: false}
            _ = v11
        }
    } else {
        _ = len(v3)
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS1{F0: true}
            _ = fi12
            _ = v13
        }
    }
    fmt.Println(v1)
    fmt.Println(v1)
    switch "foo bar" {
    case "a":
        fmt.Println(v1)
    default:
        var v14 uint32 = 4294967295
        _ = v14
    }
    switch -9223372036854775808 {
    case -562:
        fmt.Println(v1)
    default:
        v1 = 0.1
    }
    var v15 uint16 = 0
    var v16 uint = 2
    if (v16 >= v16) {
        v17 := GenS2{F0: 1, F1: 1, F2: 1}
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
