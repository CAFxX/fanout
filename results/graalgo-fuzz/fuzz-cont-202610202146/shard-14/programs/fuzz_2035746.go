package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint8 }

func main() {
    var v1 float64 = 1e10
    var v2 float32 = ((3.14159 - float32(0.0)) / float32(1e10))
    var v3 uint = 65535
    var v4 uint = (v3 - (v3 & (v3 & v3)))
    var v5 float64 = ((3.14159 / -2.25) * v1)
    var v6 int8 = 100
    v7 := GenS1{F0: v4, F1: 18446744073709551615, F2: -1}
    v8 := GenS1{F0: v4, F1: 42, F2: 0}
    switch "foo bar" {
    case "\n\t":
        if (65535 >= 0) {
            switch "" {
            case "hello":
                var v9 bool = (481 != 2)
                _ = v9
            case "a":
                var v10 uint = (v4 & 18446744073709551615)
                _ = v10
            case "\n\t":
                if true {
                    v11 := GenS1{F0: (v3 + v4), F1: 2, F2: 517}
                    if (true || ((true && false) && (v6 == 127))) {
                        v3 += 18446744073709551615
                        switch 128 {
                        case 642:
                            fmt.Println(v2, v5)
                        case 1:
                            v12 := GenS2{F0: 72}
                            _ = v12
                        }
                        var v13 uint8 = 2
                        _ = v13
                    }
                    _ = v11
                } else {
                    v4 -= (v4 / ((69) | 1))
                    v14 := []int{42, -93}
                    _ = v14
                }
            default:
                switch (("\n\t" + "\n\t") + (("hello" + (("" + (("hello" + "\n\t") + (("" + "") + "\n\t"))) + "a")) + "")) {
                case "":
                    var v15 uint = v4
                    _ = v15
                default:
                    switch "foo bar" {
                    case "a":
                        for fi16 := 0; fi16 < 17; fi16++ {
                            v17 := GenS3{F0: 42, F1: v4, F2: 82}
                            var v18 uint32 = 959
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    default:
                        var v19 int = -9223372036854775808
                        _ = v19
                    }
                }
            }
        }
    case "a":
        for fi20 := 0; fi20 < 20; fi20++ {
            v5 -= 0.0
            _ = fi20
        }
    case "hello":
        switch 0 {
        case 753:
            v21 := []int{42, -68, -859, 42}
            _ = v21
        default:
            if false {
                v6 += -89
            } else {
                for fi22 := 0; fi22 < 17; fi22++ {
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v24 := GenS1{F0: v4, F1: 2, F2: -1}
                        _ = fi23
                        _ = v24
                    }
                    v2 += 1.5
                    _ = fi22
                }
            }
        }
    }
    fmt.Println(v4, v6)
    v25 := []int{}
    _ = len(v25)
    v6 = v6
    switch "hello" {
    case "\n\t":
        var v26 bool = true
        _ = v26
    case "hello":
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
    }
    v4 -= (v3 / (((^(v4))) | 1))
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
