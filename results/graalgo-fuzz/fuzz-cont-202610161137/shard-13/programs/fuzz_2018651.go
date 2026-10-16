package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 uint16) string {
    p1 += p1
    v5 := GenS1{F0: 831, F1: 255}
    var v6 int8 = 6
    switch "" {
    case "\n\t":
        fmt.Println(p0)
    }
    var v7 float32 = ((1e10 + (1.5 + (1e10 * (-2.25 + float32(1e10))))) + (3.14159 - (-2.25 + -2.25)))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    var v1 uint = 42
    var v2 float64 = ((((1e10 / 0.1) * 0.1) - 0.0) - 0.0)
    var v3 int16 = 0
    var v4 bool = true
    v3 -= (^(v3))
    fmt.Println(v1)
    switch "a" {
    case "foo bar":
        fmt.Println(v3)
    case "\n\t":
        switch -918 {
        case 417:
            v3 = v3
        case 128:
            v8 := GenS2{F0: 2147483648}
            _ = v8
        default:
            switch 32767 {
            case 2147483647:
                var v9 uint64 = 1
                _ = v9
            case 331:
                v1 -= (v1 & (v1 % ((v1) | 1)))
            }
        }
    default:
        v10 := []int{}
        _ = v10
    }
    v1 = 2
    var v11 bool = v4
    v12 := []int{}
    if v4 {
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int16 = (v3 - 0)
                var v16 float32 = 1.5
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = len(v12)
            var v17 uint8 = 127
            _ = fi13
            _ = v17
        }
        switch "hello" {
        case "foo bar":
            for fi18 := 0; fi18 < 20; fi18++ {
                if false {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v20 := GenS2{F0: 42}
                        switch (fi19 << ((v1) & 63)) {
                        case 998:
                            switch -134 {
                            case 0:
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                        case 4294967295:
                            v2 = 1.5
                        }
                        v21 := GenS2{F0: 2}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    if v4 {
                        v2 = 1.5
                    }
                    v22 := GenS3{F0: "\n\t"}
                    _ = v22
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    v3 = 1000
                    _ = fi23
                }
                v24 := GenS2{F0: 2147483647}
                _ = fi18
                _ = v24
            }
        case "a":
            var v25 int8 = -128
            _ = v25
        case "hello":
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
    } else {
        v26 := GenS1{F0: -2147483648, F1: (v1 - 65535)}
        _ = v26
    }
    var v27 uint16 = uint16(v1)
    v1 = 266
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
