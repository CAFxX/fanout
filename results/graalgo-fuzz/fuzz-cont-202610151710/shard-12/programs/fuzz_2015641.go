package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
type GenS2 struct { F0 string; F1 uint; F2 uint64 }
func genFn0() int8 {
    if false {
        switch "\n\t" {
        case "foo bar":
            for fi5 := 0; fi5 < 14; fi5++ {
                v6 := GenS1{F0: "hello", F1: (fi5 != (fi5 - fi5)), F2: -1}
                if true {
                    for fi7 := 0; fi7 < 13; fi7++ {
                        v8 := []int{}
                        switch 128 {
                        case -9223372036854775808:
                            var v9 uint8 = 2
                            _ = v9
                        case 420:
                            var v10 uint = 4294967295
                            _ = v10
                        }
                        _ = fi7
                        _ = v8
                    }
                    var v11 uint = 42
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v13 := GenS1{F0: "foo bar", F1: (!true), F2: (fi12 * 806)}
                        var v14 float64 = 0.1
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                }
                _ = fi5
                _ = v6
            }
        }
        v15 := []int{}
        switch "a" {
        case "hello":
            var v16 bool = (!(true && true))
            _ = v16
        case "":
            v15 = append(v15, 32768)
        }
        _ = v15
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int = fi17
        var v19 uint32 = 42
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 uint32 = 342
        _ = fi20
        _ = v21
    }
    var v22 int32 = -768
    v23 := GenS1{F0: "", F1: true, F2: -1}
    _ = v22
    _ = v23
    return 100
}
func genFn1() string {
    var v24 float64 = 1.5
    var v25 uint64 = 956
    switch 1 {
    case 1:
        if ((v24 + float64((-2.25 * (3.14159 * 1e10)))) < 1e10) {
            var v26 int8 = 63
            _ = v26
        } else {
            fmt.Println(v24, v25)
        }
    default:
        var v27 uint = 2
        _ = v27
    }
    _ = v24
    _ = v25
    return "a"
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 42
    var v4 int8 = -128
    for fi28 := 0; fi28 < 10; fi28++ {
        if (!(((^(v1)) == v1) && false)) {
            var v29 uint8 = 128
            if false {
                v30 := v2[949]
                _ = v30
            } else {
                if false {
                    v3 -= v3
                    v2[uint32(v3)] = 2147483647
                } else {
                    v31 := []int{(fi28 - (-(fi28))), (fi28 + 389), (fi28 / ((42) | 1)), fi28}
                    v31 = append(v31, (fi28 + (fi28 ^ fi28)))
                    _ = v31
                }
            }
            _ = v29
        }
        v1 = (v1 / ((v1) | 1))
        v4 += -1
        _ = fi28
    }
    v32 := GenS1{F0: "", F1: ((177 == -703) || false), F2: -9223372036854775808}
    v33 := []int{-707, -872, 268, 9223372036854775807}
    v4 += -1
    v1 = v1
    var v34 int8 = (v4 - (-(v4)))
    var v35 float64 = (1e10 - 3.14159)
    v1 = (^(v1))
    switch ("" + (("" + ("" + "a")) + "")) {
    case "a":
        var v36 uint8 = 2
        _ = v36
    case "":
        _ = len(v2)
    default:
        v37 := GenS2{F0: "a", F1: (v1 - v1), F2: v3}
        _ = v37
    }
    if false {
        var v38 float32 = 1.5
        fmt.Println(v3, v1)
        _ = v38
    } else {
        v35 += (1.5 - (0.0 * v35))
    }
    v39 := GenS1{F0: ("" + ((("" + (("a" + "foo bar") + (("\n\t" + "hello") + "hello"))) + ("foo bar" + "a")) + ("foo bar" + "hello"))), F1: true, F2: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v4
}
