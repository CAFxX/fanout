package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 900
    var v3 float64 = ((0.0 - (((0.1 * 0.0) + 1e10) * 1.5)) - 0.1)
    if (true && true) {
        for fi4 := 0; fi4 < 3; fi4++ {
            v5 := GenS1{F0: 42, F1: 834}
            switch ("" + (((("\n\t" + "a") + "foo bar") + "foo bar") + ("foo bar" + "\n\t"))) {
            case "foo bar":
                if false {
                    fmt.Println(v2)
                    v2 = (^(v2))
                }
            case "":
                for fi6 := 0; fi6 < 16; fi6++ {
                    var v7 int = 2147483647
                    fmt.Println(v3, fi4)
                    v3 = v3
                    _ = fi6
                    _ = v7
                }
            }
            _ = fi4
            _ = v5
        }
    } else {
        if false {
            fmt.Println(v2, v2)
        }
        delete(v1, "foo bar")
        if (0 == 0) {
            var v8 uint16 = 42
            switch 329 {
            case -1:
                fmt.Println(v2, v2, v3)
            case 32767:
                var v9 int32 = -1
                _ = v9
            default:
                var v10 uint16 = v8
                _ = v10
            }
            _ = v8
        }
    }
    var v11 float64 = v3
    var v12 int8 = -128
    var v13 int64 = v2
    v13 += v2
    var v14 uint = 4294967295
    for fi15 := 0; fi15 < 20; fi15++ {
        v12 = v12
        v13 -= v2
        _ = fi15
    }
    var v16 uint64 = 897
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
}
