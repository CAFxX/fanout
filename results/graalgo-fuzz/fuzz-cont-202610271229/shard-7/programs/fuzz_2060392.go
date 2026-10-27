package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 bool }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 42
    v3 := make(map[int]string)
    var v4 uint = 1
    v5 := GenS1{F0: 2147483647}
    switch "foo bar" {
    case "hello":
        var v6 float64 = (-2.25 + float64((1.5 + (3.14159 * (1.5 / 3.14159)))))
        _ = v6
    case "":
        var v7 string = "\n\t"
        _ = v7
    default:
        v8 := []int{32767, 0, -9223372036854775808}
        _ = v8
    }
    if true {
        if true {
            for fi9 := 0; fi9 < 14; fi9++ {
                switch (fi9 << (((v4 + (v4 << ((v4) & 63)))) & 63)) {
                case 185:
                    v4 += (v4 >> ((2) & 63))
                default:
                    switch 2147483647 {
                    case 127:
                        var v10 int = fi9
                        _ = v10
                    default:
                        var v11 int32 = 1
                        _ = v11
                    }
                }
                _ = fi9
            }
            var v12 bool = (-782 > -988)
            _ = v12
        }
        fmt.Println(v2)
        if false {
            switch "a" {
            case "a":
                var v13 int16 = -32768
                _ = v13
            case "foo bar":
                v14 := GenS2{F0: 4294967295, F1: 2, F2: (2 > 42)}
                _ = v14
            case "hello":
                var v15 uint = v4
                _ = v15
            }
            v2 -= (^(v2))
            switch ("a" + "foo bar") {
            case "":
                v2 -= v2
            default:
                var v16 uint32 = 525
                _ = v16
            }
        }
    } else {
        v2 -= v2
    }
    var v17 uint16 = 1
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS1{F0: 4294967295}
        var v20 int16 = (v2 + (-(v2)))
        v17 += 65535
        _ = fi18
        _ = v19
        _ = v20
    }
    switch -321 {
    case 127:
        var v21 int16 = (v2 | v2)
        _ = v21
    default:
        var v22 string = "foo bar"
        _ = v22
    }
    if false {
        v23 := GenS3{F0: 42}
        switch 4294967295 {
        case 127:
            if true {
                var v24 float64 = 3.14159
                _ = v24
            }
        default:
            fmt.Println(v4)
        }
        _ = v23
    }
    fmt.Println(v4, v2, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
