package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint32 }
func genFn0(p0 uint32) int16 {
    var v5 float64 = ((0.1 / ((1.5 / 0.1) * float64(0.1))) + ((0.1 - 0.1) + (1.5 / 3.14159)))
    fmt.Println(p0, v5)
    p0 = p0
    _ = p0
    _ = v5
    return -32768
}
func genFn1(p0 string, p1 bool) string {
    fmt.Println(p1)
    v6 := []int{38, 42}
    if (-1 == 1) {
        var v7 uint8 = 42
        var v8 int16 = -32768
        _ = v7
        _ = v8
    } else {
        var v9 int64 = -759
        if p1 {
            var v10 int = 0
            switch p0 {
            case "":
                if (1e10 == 1.5) {
                    v6 = append(v6, 32767)
                }
            default:
                var v11 string = (p0 + p0)
                _ = v11
            }
            fmt.Println(v9)
            _ = v10
        }
        if p1 {
            v9 += v9
            v12 := GenS1{F0: -2147483648, F1: 127}
            v13 := GenS2{F0: 1, F1: 138}
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 13; fi14++ {
                v9 = v9
                _ = fi14
            }
            if p1 {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 int8 = 42
    var v2 bool = true
    var v3 bool = (-2.25 <= (float64(0.1) - 1e10))
    var v4 uint = 156
    v4 = v4
    var v15 bool = (v2 || (v3 || v3))
    if (!false) {
        for fi16 := 0; fi16 < 2; fi16++ {
            if (-2.25 == -2.25) {
                fmt.Println(v3, v15)
            }
            for fi17 := 0; fi17 < 2; fi17++ {
                switch "a" {
                case "\n\t":
                    v3 = (("hello" + "") <= (("\n\t" + "") + "\n\t"))
                case "a":
                    var v18 uint8 = 128
                    _ = v18
                default:
                    if v2 {
                        fmt.Println(v4, v3)
                        var v19 uint = v4
                        _ = v19
                    } else {
                        var v20 int32 = -2147483648
                        var v21 string = ("a" + "\n\t")
                        _ = v20
                        _ = v21
                    }
                }
                v4 += v4
                _ = fi17
            }
            _ = fi16
        }
        v22 := []int{2147483647, 128}
        switch "" {
        case "a":
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        case "":
            _ = len(v22)
        default:
            var v23 uint16 = 255
            _ = v23
        }
        _ = v22
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        fmt.Println(v1, v2, v3)
        _ = fi24
    }
    var v25 float32 = 0.1
    v3 = (!(1 <= 0))
    v26 := GenS3{F0: 2, F1: 84, F2: 569}
    fmt.Println(v15, v2)
    for fi27 := 0; fi27 < 17; fi27++ {
        v4 -= v4
        _ = fi27
    }
    v28 := []int{0, -103, -169}
    for fi29 := 0; fi29 < 14; fi29++ {
        var v30 int = fi29
        _ = fi29
        _ = v30
    }
    v25 -= v25
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 int = fi31
        for fi33 := 0; fi33 < 7; fi33++ {
            switch "hello" {
            case "\n\t":
                fmt.Println(v1, fi33, v1)
            case "hello":
                var v34 float32 = 1e10
                v1 = -72
                fmt.Println(v34, v32)
                _ = v34
            }
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    var v35 bool = v15
    fmt.Println(v3, v4)
    fmt.Println(v2, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v35
    _ = v4
}
