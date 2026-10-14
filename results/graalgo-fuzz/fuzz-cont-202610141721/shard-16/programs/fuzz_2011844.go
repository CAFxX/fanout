package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int8; F1 int8; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 1
    var v3 uint = 795
    v4 := []uint32{121}
    v5 := make(map[string]uint32)
    var v6 int8 = 42
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float32 = 1e10
        if ((fi7 >> ((v3) & 63)) < -927) {
            v9 := v5["foo bar"]
            _ = v9
        }
        if ((v3 ^ v3) < (^(v3))) {
            for fi10 := 0; fi10 < 3; fi10++ {
                _ = len(v1)
                _ = fi10
            }
            v11 := GenS1{F0: 2, F1: v3}
            if (42 < 65) {
                v12 := GenS1{F0: 1, F1: (v3 ^ (^(v3)))}
                v3 += 18446744073709551615
                var v13 int16 = -32768
                _ = v12
                _ = v13
            }
            _ = v11
        } else {
            for fi14 := 0; fi14 < 5; fi14++ {
                if true {
                    v2 = 2147483647
                }
                _ = len(v1)
                v8 = ((v8 * v8) + (v8 - v8))
                _ = fi14
            }
        }
        _ = fi7
        _ = v8
    }
    var v15 uint16 = 38
    if (!(false || false)) {
        v6 -= (v6 - (v6 & (^(v6))))
    } else {
        if true {
            v16 := GenS2{F0: v6, F1: 127, F2: 1}
            _ = v16
        } else {
            fmt.Println(v6, v2, v3)
            fmt.Println(v6)
        }
        var v17 string = "hello"
        _ = v17
    }
    var v18 int16 = -32768
    switch int(v6) {
    case 952:
        switch 389 {
        case -31:
            var v19 float64 = ((0.0 * float64(-2.25)) + float64(0.0))
            _ = v19
        case 934:
            if true {
                v20 := GenS2{F0: (v6 * (v6 & 0)), F1: v6, F2: -399}
                v3 = (v3 % ((42) | 1))
                _ = v20
            } else {
                fmt.Println(v18)
                fmt.Println(v3, v15)
            }
        }
    case -905:
        v6 -= (v6 + v6)
    case 226:
        var v21 bool = false
        _ = v21
    }
    var v22 string = "\n\t"
    switch ("foo bar" + (v22 + "a")) {
    case "a":
        v6 += 100
    }
    if (true && (!false)) {
        var v23 uint64 = 42
        _ = v23
    } else {
        fmt.Println(v22, v15, v3)
        delete(v1, -947)
        v24 := GenS1{F0: 4294967295, F1: (^(v3))}
        _ = v24
    }
    var v25 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
