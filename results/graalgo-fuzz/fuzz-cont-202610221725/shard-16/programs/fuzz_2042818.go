package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int8; F2 string }

func main() {
    v1 := []string{"\n\t", "foo bar"}
    var v2 string = "foo bar"
    v3 := make(map[uint32]string)
    v4 := make(map[int]bool)
    var v5 uint8 = 127
    var v6 uint64 = 900
    if true {
        var v7 float32 = 1.5
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int16 = -176
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v10 := GenS1{F0: -1, F1: 752, F2: 87}
    if false {
        var v11 bool = true
        var v12 uint = 65535
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v2, v5)
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 int16 = 329
            delete(v4, fi14)
            fmt.Println(v2, v6, v6)
            v15 = v15
            fmt.Println(v15)
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v5 += v5
    var v16 int8 = 98
    var v17 int = 4294967295
    if (false && false) {
        var v18 uint8 = (v5 - v5)
        switch v2 {
        case "\n\t":
            var v19 bool = false
            _ = v19
        default:
            if ((266 != 42) && (1000 != 526)) {
                for fi20 := 0; fi20 < 14; fi20++ {
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v22 := GenS2{F0: 810, F1: 81, F2: ("foo bar" + "a")}
                        v3[0] = v2
                        _ = fi21
                        _ = v22
                    }
                    switch (v2 + "foo bar") {
                    case "hello":
                        var v23 string = v2
                        _ = v23
                    case "a":
                        var v24 int64 = 0
                        _ = v24
                    case "foo bar":
                        v5 += (^(v18))
                    }
                    switch v2 {
                    case "\n\t":
                        v6 -= (v6 & (v6 * v6))
                    }
                    _ = fi20
                }
                v16 -= int8(v6)
                var v25 int32 = 0
                _ = v25
            }
        }
        v26 := GenS2{F0: 32767, F1: (v16 / ((v16) | 1)), F2: "foo bar"}
        _ = v18
        _ = v26
    } else {
        if (!true) {
            for fi27 := 0; fi27 < 1; fi27++ {
                fmt.Println(v16, v17)
                v5 = 128
                _ = fi27
            }
            var v28 float32 = 0.0
            _ = v28
        }
        if false {
            v1 = append(v1, v2)
            v17 = 127
        }
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 bool = (false || false)
        v4[(v17 & v17)] = v30
        _ = fi29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
