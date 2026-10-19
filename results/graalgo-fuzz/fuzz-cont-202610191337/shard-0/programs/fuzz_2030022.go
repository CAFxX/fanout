package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 string; F1 int8 }

func main() {
    var v1 int32 = 0
    var v2 bool = (v1 >= (v1 & v1))
    var v3 int = 42
    var v4 uint8 = 128
    var v5 uint = 42
    var v6 uint = v5
    v7 := []int{v3}
    fmt.Println(v4, v4, v5)
    _ = len(v7)
    fmt.Println(v1, v6, v4)
    v8 := GenS1{F0: -534, F1: 209}
    v1 += (v1 & v1)
    if (v2 && v2) {
        if true {
            if true {
                v3 -= v3
                for fi9 := 0; fi9 < 8; fi9++ {
                    var v10 float32 = (3.14159 - 0.0)
                    var v11 int64 = 327
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                for fi12 := 0; fi12 < 13; fi12++ {
                    for fi13 := 0; fi13 < 5; fi13++ {
                        fmt.Println(v2, fi12)
                        switch ("hello" + (("\n\t" + "foo bar") + "a")) {
                        case "\n\t":
                            v5 = 0
                        }
                        var v14 int8 = 127
                        _ = fi13
                        _ = v14
                    }
                    var v15 int64 = 0
                    v16 := GenS1{F0: -13, F1: 1000}
                    _ = fi12
                    _ = v15
                    _ = v16
                }
            }
        } else {
            v4 += 42
            if ((!true) || v2) {
                var v17 int8 = 0
                _ = v17
            } else {
                for fi18 := 0; fi18 < 17; fi18++ {
                    switch "hello" {
                    case "foo bar":
                        var v19 int64 = 1
                        _ = v19
                    case "a":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    _ = fi18
                }
                v1 = (^(v1))
                if ((v1 / ((v1) | 1)) >= v1) {
                    if len(v7) > 0 {
                        v7[0] = v3
                    }
                    if (!v2) {
                        fmt.Println(v5, v2, v5)
                        _ = len(v7)
                    }
                }
            }
        }
        v7 = append(v7, 140)
    } else {
        v2 = (v2 && false)
        var v20 float32 = 0.1
        switch v3 {
        case 32767:
            v4 = (v4 << (((v5 % (((v5 >> (((v5 + (^(v6)))) & 63))) | 1))) & 63))
        case 4294967295:
            v4 += v4
        default:
            switch v3 {
            case 9223372036854775807:
                var v21 float64 = (1.5 * 0.0)
                _ = v21
            default:
                for fi22 := 0; fi22 < 10; fi22++ {
                    v23 := GenS2{F0: "a", F1: -49}
                    _ = fi22
                    _ = v23
                }
            }
        }
        _ = v20
    }
    fmt.Println(v2, v2, v3)
    var v24 uint8 = 42
    fmt.Println(v2, v5)
    v24 += v24
    if (v2 || false) {
        var v25 bool = true
        var v26 float64 = (1e10 / 3.14159)
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        v28 := GenS1{F0: -699, F1: 1000}
        v4 = 255
        _ = fi27
        _ = v28
    }
    switch (v3 & v3) {
    case 9223372036854775807:
        var v29 int64 = 42
        _ = v29
    case 128:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    default:
        if v2 {
            var v30 int64 = 42
            v3 = v3
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
