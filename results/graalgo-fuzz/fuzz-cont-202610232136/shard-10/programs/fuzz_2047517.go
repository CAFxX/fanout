package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int8; F1 bool; F2 int }
type GenS3 struct { F0 int16; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int; F1 bool }

func main() {
    var v1 uint32 = 0
    var v2 int16 = 1000
    var v3 float32 = 1.5
    v4 := make(map[string]uint32)
    v5 := []string{"\n\t", ("hello" + ("\n\t" + ("" + "a"))), ("hello" + "a"), ("" + "")}
    v6 := GenS1{F0: 45, F1: 9223372036854775807}
    v7 := GenS2{F0: 1, F1: (((v1 ^ (^(v1))) <= v1) && ((!true) && true)), F2: 0}
    var v8 uint32 = v1
    if false {
        var v9 int32 = 42
        v5 = append(v5, "a")
        _ = v9
    }
    var v10 uint8 = 255
    fmt.Println(v3)
    v11 := GenS1{F0: v10, F1: 0}
    if true {
        if (1 >= -2147483648) {
            v12 := GenS1{F0: (v10 - v10), F1: 9223372036854775807}
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            fmt.Println(v8, v1)
            _ = v12
        }
        switch ("" + "foo bar") {
        case "":
            for fi13 := 0; fi13 < 13; fi13++ {
                fmt.Println(fi13, v1)
                fmt.Println(v3, v10, v8)
                fmt.Println(v2)
                _ = fi13
            }
        case "\n\t":
            var v14 uint8 = v10
            _ = v14
        case "a":
            var v15 uint32 = 2147483647
            _ = v15
        default:
            _ = len(v4)
        }
        if ((!true) && (!((-2.25 * 1e10) != (0.1 + -2.25)))) {
            if false {
                for fi16 := 0; fi16 < 11; fi16++ {
                    v17 := GenS3{F0: v2, F1: uint32(v10), F2: 100}
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 2; fi18++ {
                    if true {
                        v8 = v1
                        v19 := GenS1{F0: 51, F1: 667}
                        for fi20 := 0; fi20 < 10; fi20++ {
                            v8 = 0
                            _ = fi20
                        }
                        _ = v19
                    } else {
                        v21 := GenS4{F0: 2, F1: (v2 - (v2 / (((^(v2))) | 1)))}
                        v22 := GenS2{F0: 127, F1: (true && (!(!(0 != 42)))), F2: fi18}
                        _ = v21
                        _ = v22
                    }
                    if (127 >= 1) {
                        if true {
                            v23 := GenS1{F0: v10, F1: 0}
                            _ = v23
                        }
                        _ = len(v4)
                        for fi24 := 0; fi24 < 7; fi24++ {
                            var v25 int8 = 116
                            _ = fi24
                            _ = v25
                        }
                    }
                    v1 += (v8 + 65535)
                    _ = fi18
                }
            } else {
                fmt.Println(v3, v1)
            }
        }
    } else {
        var v26 uint16 = 48
        fmt.Println(v2, v3)
        _ = v26
    }
    v27 := GenS5{F0: -1, F1: ((^(v10)) >= 42)}
    var v28 uint32 = (v1 | uint32(v10))
    v5 = append(v5, "hello")
    var v29 float32 = 1e10
    var v30 uint8 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
