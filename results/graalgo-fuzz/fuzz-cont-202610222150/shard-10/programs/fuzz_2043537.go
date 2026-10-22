package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []uint32{2, 2147483647}
    var v2 float64 = 0.0
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]uint32)
    var v5 float32 = 3.14159
    v5 += v5
    var v6 int8 = -128
    v5 = (v5 / 3.14159)
    var v7 int = 32768
    switch "foo bar" {
    case "a":
        var v8 uint = 109
        _ = v8
    case "hello":
        v6 += 0
    }
    var v9 int64 = 42
    if (true && ((1.5 + 0.0) > 3.14159)) {
        var v10 float64 = v2
        v11 := GenS1{F0: v6}
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 bool = (!true)
            fmt.Println(v3, v9, v6)
            if v13 {
                if v13 {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        var v15 float32 = 3.14159
                        var v16 float32 = v5
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
                v6 += -128
            }
            _ = fi12
            _ = v13
        }
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        v9 = (v9 + v3)
        v1 = append(v1, 42)
        fmt.Println(v3, v2)
        _ = fi17
    }
    v5 = 3.14159
    var v18 uint16 = 576
    if ((!false) && (!true)) {
        for fi19 := 0; fi19 < 13; fi19++ {
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 uint = 89
                _ = len(v4)
                switch "hello" {
                case "":
                    fmt.Println(v3)
                default:
                    fmt.Println(v3, v18, fi19)
                }
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: v6}
            _ = fi19
            _ = v22
        }
        v2 -= v2
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v7 = (fi23 | fi23)
        _ = fi23
    }
    switch (v7 % ((v7) | 1)) {
    case 4294967295:
        v3 += v3
    default:
        var v24 int8 = (v6 - v6)
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        if true {
            _ = len(v1)
        }
        v2 = v2
        var v26 float64 = v2
        _ = fi25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
