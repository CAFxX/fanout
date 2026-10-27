package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint8 }

func main() {
    var v1 int8 = -128
    var v2 uint32 = 2
    var v3 uint64 = 258
    v4 := make(map[string]string)
    var v5 int16 = -1
    fmt.Println(v1)
    if true {
        if (v5 != (v5 + 1000)) {
            v1 += (^(v1))
            var v6 uint8 = 255
            var v7 int8 = 42
            _ = v6
            _ = v7
        } else {
            if false {
                v8 := GenS1{F0: 739}
                v3 -= v3
                var v9 bool = (false || (!true))
                _ = v8
                _ = v9
            }
            var v10 float32 = (-2.25 / (((((0.0 / 1e10) + 1.5) * 3.14159) / (1e10 + 1e10)) - float32(-2.25)))
            _ = v10
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 float64 = 0.0
            v5 -= v5
            v1 = v1
            _ = fi11
            _ = v12
        }
    }
    v1 += -128
    switch "\n\t" {
    case "foo bar":
        v3 = (v3 % (((^(v3))) | 1))
    case "hello":
        switch 127 {
        case 32767:
            switch 127 {
            case -807:
                var v13 uint64 = v3
                _ = v13
            }
        }
    }
    v3 = 4294967295
    v1 += v1
    v4[""] = "foo bar"
    var v14 uint32 = (v2 * v2)
    v14 = v2
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 bool = false
        _ = fi15
        _ = v16
    }
    v4[(("a" + "\n\t") + "")] = ("foo bar" + "hello")
    if ("a" != ("\n\t" + "\n\t")) {
        var v17 uint16 = 0
        _ = v17
    } else {
        if false {
            switch "hello" {
            case "a":
                var v18 uint8 = 152
                _ = v18
            case "hello":
                switch 436 {
                case 390:
                    v1 = v1
                }
            }
        }
        var v19 uint16 = uint16(v3)
        var v20 int64 = 577
        _ = v19
        _ = v20
    }
    fmt.Println(v14, v2, v1)
    v21 := []int{694, 0}
    v22 := GenS2{F0: ("\n\t" + (("" + "foo bar") + "a")), F1: 127}
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 int16 = v5
        _ = len(v4)
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 255}
    v26 := GenS1{F0: 489}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
