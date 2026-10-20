package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 bool }
func genFn0(p0 bool) uint64 {
    p0 = (false || p0)
    v5 := []int{-9223372036854775808, 4294967295, 32767}
    var v6 bool = (!false)
    var v7 int = 128
    for fi8 := 0; fi8 < 5; fi8++ {
        v6 = v6
        fmt.Println(v6, p0, v6)
        for fi9 := 0; fi9 < 13; fi9++ {
            fmt.Println(v6)
            _ = fi9
        }
        _ = fi8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 118
}
func genFn1() int16 {
    var v10 float32 = -2.25
    var v11 uint16 = 1
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 bool = (1 <= 4294967295)
        _ = fi12
        _ = v13
    }
    _ = v10
    _ = v11
    return -499
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 127
    var v3 int64 = 0
    v4 := make(map[string]uint32)
    switch 0 {
    case -677:
        var v14 uint32 = 390
        _ = v14
    case 9223372036854775807:
        if false {
            v15 := GenS1{F0: 42, F1: 1, F2: false}
            fmt.Println(v3, v3)
            _ = v15
        }
    }
    fmt.Println(v2, v2)
    if false {
        v16 := []int{}
        v2 = (v2 | 95)
        _ = v16
    }
    v17 := GenS1{F0: -1, F1: 2147483647, F2: false}
    _ = len(v1)
    if true {
        if (2 < 510) {
            v3 += (v3 - (v3 + (v3 >> ((1) & 63))))
        }
        v2 += (v2 | 1)
        if true {
            switch "\n\t" {
            case "hello":
                var v18 uint8 = v2
                _ = v18
            case "a":
                if true {
                    v4[""] = 2147483647
                }
            case "":
                v3 = 1
            default:
                delete(v4, ((("" + "a") + (("" + ("\n\t" + "\n\t")) + "\n\t")) + ("a" + ("\n\t" + "a"))))
            }
            if false {
                fmt.Println(v2, v3)
                var v19 int8 = 103
                _ = v19
            }
            switch (((("a" + "\n\t") + "a") + ("hello" + "hello")) + "") {
            case "foo bar":
                if (!(!true)) {
                    for fi20 := 0; fi20 < 11; fi20++ {
                        switch 127 {
                        case 128:
                            if true {
                                var v21 int64 = (v3 - 448)
                                v3 -= (v21 & v21)
                                _ = v21
                            } else {
                                var v22 float64 = 0.0
                                _ = v22
                            }
                        case -2147483648:
                            v23 := []int{(fi20 >> ((uint(v2)) & 63))}
                            _ = v23
                        default:
                            var v24 uint64 = 42
                            _ = v24
                        }
                        v2 = v2
                        fmt.Println(v2)
                        _ = fi20
                    }
                }
            }
        }
    }
    v25 := v1[2]
    if true {
        if true {
            v2 += 2
            v2 -= (v2 * (v2 % (((v2 >> ((42) & 63))) | 1)))
        } else {
            v26 := []int{}
            v3 += (^(v3))
            _ = v26
        }
        switch ("a" + (("foo bar" + "foo bar") + "a")) {
        case "\n\t":
            v3 = v3
        case "":
            _ = len(v4)
        case "a":
            var v27 int16 = 42
            _ = v27
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
