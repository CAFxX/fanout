package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint }
func genFn0(p0 uint8, p1 string, p2 uint) uint64 {
    var v6 int32 = 0
    var v7 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 921
}

func main() {
    v1 := []string{}
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 bool = false
    var v5 uint = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 = (v5 << ((255) & 63))
    v4 = (127 < 100)
    var v8 bool = (42 <= 1)
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 uint32 = 0
            switch (fi9 - fi10) {
            case 0:
                v3[(fi10 - (fi9 % ((fi9) | 1)))] = int(v11)
            case -1:
                switch (("a" + "") + "a") {
                case "hello":
                    var v12 string = "a"
                    _ = v12
                case "a":
                    fmt.Println(v8)
                case "\n\t":
                    switch "foo bar" {
                    case "\n\t":
                        var v13 int8 = -94
                        _ = v13
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "a":
                        var v14 float64 = 1.5
                        _ = v14
                    }
                default:
                    var v15 int32 = -846
                    _ = v15
                }
            default:
                _ = len(v3)
            }
            _ = fi10
            _ = v11
        }
        var v16 uint = 65535
        v8 = v8
        _ = fi9
        _ = v16
    }
    if false {
        if v8 {
            fmt.Println(v4, v8)
        }
        switch "hello" {
        case "\n\t":
            var v17 uint16 = 65535
            _ = v17
        default:
            fmt.Println(v5, v8, v5)
        }
    } else {
        v8 = (42 < -9223372036854775808)
        var v18 int8 = -1
        var v19 uint16 = 220
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(v5)
        if (42 <= 0) {
            var v21 float32 = 1e10
            v5 = (^(v5))
            _ = v21
        } else {
            for fi22 := 0; fi22 < 10; fi22++ {
                if true {
                    fmt.Println(fi22, v5, v5)
                    var v23 int8 = 42
                    var v24 uint64 = genFn0(172, "hello", (v5 << (((v5 % ((v5) | 1))) & 63)))
                    var v25 int64 = 0
                    _ = v23
                    _ = v24
                    _ = v25
                }
                var v26 uint64 = 42
                _ = fi22
                _ = v26
            }
        }
        v27 := GenS1{F0: 1, F1: -889, F2: 0}
        _ = fi20
        _ = v27
    }
    v4 = v4
    var v28 int = 42
    v5 += (^(v5))
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 uint = v5
        _ = fi29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
