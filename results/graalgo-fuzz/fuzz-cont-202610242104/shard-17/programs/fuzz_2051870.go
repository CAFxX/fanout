package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }

func main() {
    var v1 bool = true
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    var v4 string = "\n\t"
    var v5 uint16 = 35
    v4 = v4
    v5 -= 378
    var v6 int64 = -1
    v4 = (v4 + "")
    var v7 int16 = 1
    var v8 float32 = 1e10
    v1 = (v1 && (v6 > (v6 << ((42) & 63))))
    var v9 uint64 = 1
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: "", F1: (fi10 + fi10)}
        if (v1 || (4294967295 > 65535)) {
            var v12 string = (v4 + (v4 + (v4 + "")))
            var v13 int16 = (^(v7))
            _ = v12
            _ = v13
        } else {
            fmt.Println(fi10, fi10, v5)
        }
        _ = fi10
        _ = v11
    }
    v1 = (2 == 65535)
    if (v1 || v1) {
        switch v4 {
        case "":
            switch "a" {
            case "a":
                for fi14 := 0; fi14 < 11; fi14++ {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 string = v4
                        _ = fi15
                        _ = v16
                    }
                    _ = fi14
                }
            case "foo bar":
                switch 32768 {
                case -9223372036854775808:
                    var v17 uint8 = 110
                    _ = v17
                case 951:
                    v4 = (v4 + (v4 + ""))
                default:
                    fmt.Println(v7, v8, v5)
                }
            case "hello":
                var v18 uint = uint(v9)
                _ = v18
            }
        default:
            if (-2147483648 < 805) {
                fmt.Println(v5, v6)
                var v19 int16 = v7
                switch 9223372036854775807 {
                case 4294967295:
                    v20 := GenS1{F0: v4, F1: -296}
                    _ = v20
                default:
                    fmt.Println(v9, v8)
                }
                _ = v19
            } else {
                var v21 uint16 = v5
                v8 -= ((v8 + (v8 + 1e10)) / v8)
                _ = v21
            }
        }
        if (v1 && v1) {
            for fi22 := 0; fi22 < 13; fi22++ {
                switch (fi22 | fi22) {
                case -546:
                    v5 += (v5 >> ((65535) & 63))
                }
                v5 = v5
                _ = fi22
            }
            var v23 int8 = 1
            _ = v23
        }
    } else {
        if (416 < 0) {
            fmt.Println(v4, v4)
        }
    }
    if ((!false) && (-484 > -2147483648)) {
        var v24 float64 = 1.5
        var v25 int16 = v7
        var v26 uint = 508
        _ = v24
        _ = v25
        _ = v26
    }
    v1 = false
    var v27 uint16 = v5
    var v28 uint16 = 0
    var v29 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
