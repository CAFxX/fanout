package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 string; F1 uint; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 1
    var v3 int = 0
    var v4 bool = true
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int8 = int8(v3)
        _ = fi5
        _ = v6
    }
    if (!(4294967295 >= 643)) {
        v3 = (-(v3))
    }
    v7 := []int{v3, v3, -223, v3}
    v4 = v4
    var v8 int8 = int8(v2)
    if v4 {
        switch (((("" + "") + "") + "hello") + "a") {
        case "":
            var v9 bool = v4
            _ = v9
        default:
            v4 = v4
        }
        v4 = false
        fmt.Println(v8, v4, v4)
    } else {
        var v10 int8 = (v8 ^ (^(v8)))
        var v11 float32 = float32(1e10)
        switch ("a" + ("\n\t" + "\n\t")) {
        case "a":
            _ = len(v1)
        }
        _ = v10
        _ = v11
    }
    if v4 {
        if (166 > -1) {
            fmt.Println(v3, v2, v4)
        }
        var v12 uint = 439
        var v13 bool = v4
        _ = v12
        _ = v13
    }
    var v14 uint32 = 4294967295
    v7 = append(v7, (v3 % ((v3) | 1)))
    delete(v1, 65535)
    switch "foo bar" {
    case "hello":
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 uint32 = (v14 + v14)
            if (!(v2 == v2)) {
                v17 := GenS1{F0: 1, F1: 1}
                _ = v17
            } else {
                var v18 float64 = (0.0 - 1e10)
                switch (fi15 / ((0) | 1)) {
                case 127:
                    switch (v3 * fi15) {
                    case 128:
                        if true {
                            v18 = v18
                        }
                    case 805:
                        fmt.Println(v14, v2, v3)
                    }
                }
                _ = v18
            }
            _ = fi15
            _ = v16
        }
    case "\n\t":
        if ((42 == 418) && (18446744073709551615 > 255)) {
            v8 = 0
        } else {
            v4 = (v4 && (v4 || (v8 == 42)))
        }
    case "":
        v4 = (v4 || false)
    default:
        v19 := GenS1{F0: 1, F1: 1}
        _ = v19
    }
    v20 := GenS2{F0: "", F1: 952, F2: -1}
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 uint64 = 2
        fmt.Println(v8, v4, v8)
        var v23 uint16 = (v2 / (((^(v2))) | 1))
        _ = fi21
        _ = v22
        _ = v23
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
