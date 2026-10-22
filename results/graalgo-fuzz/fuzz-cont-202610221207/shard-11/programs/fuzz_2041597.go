package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 bool; F1 int64; F2 bool }

func main() {
    var v1 int = 676
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    var v6 uint32 = 4294967295
    switch ("foo bar" + "hello") {
    case "\n\t":
        var v7 float64 = -2.25
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint64 = 308
            var v10 float32 = (0.1 / float32(1e10))
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v1, fi11)
        _ = fi11
    }
    if (("a" + ("\n\t" + (("foo bar" + "\n\t") + "hello"))) > "foo bar") {
        var v12 uint32 = 2
        fmt.Println(v12, v12, v1)
        _ = v12
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 string = "\n\t"
            v1 += fi13
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            v1 = (^(v1))
            var v16 uint32 = v6
            v17 := GenS1{F0: 46, F1: 2}
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := GenS1{F0: 42, F1: 0}
        _ = v18
    }
    v6 = v6
    switch "a" {
    case "hello":
        if true {
            for fi19 := 0; fi19 < 4; fi19++ {
                v1 = 32768
                switch 32768 {
                case 32767:
                    var v20 float64 = 0.0
                    _ = v20
                case -1:
                    v21 := GenS2{F0: true, F1: 42, F2: false}
                    _ = v21
                }
                _ = fi19
            }
        }
    case "":
        v22 := GenS1{F0: 127, F1: 42}
        _ = v22
    }
    if false {
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 float64 = 0.0
            v1 = (v1 ^ 867)
            _ = fi23
            _ = v24
        }
    }
    v2[(-(v1))] = (v1 % (((v1 + (^(v1)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
