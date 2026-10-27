package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 string }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 int32, p1 int64, p2 uint32) int64 {
    v4 := []int{-828}
    p0 -= p0
    switch 594 {
    case 0:
        fmt.Println(p0)
    default:
        switch (("hello" + ("" + "hello")) + "foo bar") {
        case "foo bar":
            v4 = append(v4, 2147483647)
        default:
            p1 -= p1
        }
    }
    var v5 int8 = 42
    var v6 uint8 = 220
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return -925
}

func main() {
    var v1 int8 = 0
    var v2 uint = 1
    v3 := make(map[uint32]bool)
    v2 = (v2 | v2)
    var v7 int32 = -2147483648
    fmt.Println(v7, v2, v7)
    var v8 int8 = v1
    switch "foo bar" {
    case "a":
        v1 += v8
    case "":
        v9 := []int{32768}
        _ = v9
    }
    switch 128 {
    case 127:
        v2 = 1
    }
    if false {
        v10 := []int{32768}
        _ = v10
    } else {
        v11 := GenS1{F0: (^(v7)), F1: -222, F2: "foo bar"}
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint32 = 2147483647
            if (366 == 2) {
                var v14 int16 = -32768
                switch "foo bar" {
                case "":
                    var v15 uint8 = 72
                    _ = v15
                }
                fmt.Println(fi12)
                _ = v14
            }
            v16 := GenS1{F0: v7, F1: v7, F2: ("\n\t" + ("hello" + "\n\t"))}
            _ = fi12
            _ = v13
            _ = v16
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int = (fi17 * 32767)
            _ = fi17
            _ = v18
        }
        switch 32768 {
        case 32768:
            v19 := GenS2{F0: "", F1: 186}
            _ = v19
        case -546:
            if (!(!true)) {
                switch 32767 {
                case 127:
                    v2 = v2
                case 32768:
                    if (!false) {
                        if false {
                            switch ("\n\t" + "foo bar") {
                            case "":
                                fmt.Println(v2)
                            case "foo bar":
                                v20 := GenS1{F0: v7, F1: v7, F2: "hello"}
                                _ = v20
                            case "hello":
                                v7 = (v7 % ((v7) | 1))
                            }
                        }
                        _ = len(v3)
                    }
                default:
                    var v21 string = ("a" + "foo bar")
                    _ = v21
                }
                v22 := []int{}
                _ = len(v3)
                _ = v22
            } else {
                switch "hello" {
                case "\n\t":
                    var v23 int = 710
                    _ = v23
                default:
                    if false {
                        for fi24 := 0; fi24 < 6; fi24++ {
                            v8 += v8
                            var v25 int8 = 42
                            _ = fi24
                            _ = v25
                        }
                        v26 := []int{}
                        _ = v26
                    } else {
                        v1 += 100
                        v27 := v3[805]
                        v2 = 505
                        _ = v27
                    }
                }
                var v28 uint = v2
                _ = v28
            }
        default:
            var v29 float32 = (1e10 * 0.0)
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
