package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int64 }
func genFn0() int32 {
    var v7 uint64 = 18446744073709551615
    var v8 int16 = 32767
    v8 -= -748
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int32, p1 string) int8 {
    fmt.Println(p0, p0)
    p0 = (p0 - (p0 / ((370) | 1)))
    var v9 int16 = 42
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint = 1
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 89
}

func main() {
    v1 := []int{32768}
    var v2 int16 = -141
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]uint32)
    var v5 string = ("" + "\n\t")
    var v6 string = v5
    _ = len(v4)
    var v12 string = v6
    var v13 uint32 = 1
    switch v6 {
    case "hello":
        fmt.Println(v12, v5, v2)
    case "\n\t":
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(v13, v6)
            for fi15 := 0; fi15 < 9; fi15++ {
                v2 = (-(v2))
                v12 = "a"
                v16 := GenS1{F0: 1, F1: 255}
                _ = fi15
                _ = v16
            }
            fmt.Println(v12, v13, v13)
            _ = fi14
        }
    case "a":
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 string = ""
            switch -9223372036854775808 {
            case -2147483648:
                v19 := GenS2{F0: 9223372036854775807}
                _ = v19
            case 127:
                switch -243 {
                case 0:
                    var v20 bool = (((-2.25 < -2.25) || true) && ((1 > -509) || false))
                    _ = v20
                case 425:
                    if false {
                        var v21 int8 = 47
                        _ = v21
                    }
                default:
                    switch ("" + "hello") {
                    case "\n\t":
                        v3 = 18446744073709551615
                    }
                }
            case 1:
                var v22 float64 = float64(3.14159)
                _ = v22
            default:
                fmt.Println(v12, v5, v13)
            }
            v23 := GenS1{F0: 1, F1: 103}
            _ = fi17
            _ = v18
            _ = v23
        }
    default:
        v2 += (v2 >> ((96) & 63))
    }
    v24 := GenS1{F0: 42, F1: 2}
    delete(v4, 844)
    if true {
        var v25 uint = 65535
        switch "hello" {
        case "hello":
            if false {
                v2 = 1
                for fi26 := 0; fi26 < 10; fi26++ {
                    if true {
                        fmt.Println(fi26)
                        fmt.Println(v13, v2)
                        var v27 string = "\n\t"
                        _ = v27
                    }
                    _ = fi26
                }
            } else {
                if len(v1) > 0 {
                    v1[0] = -1
                }
                v5 = (v5 + v5)
            }
        case "\n\t":
            var v28 string = ((v6 + "\n\t") + v12)
            _ = v28
        case "foo bar":
            fmt.Println(v3, v5, v3)
        default:
            var v29 uint32 = v13
            _ = v29
        }
        _ = v25
    } else {
        v13 -= v13
        var v30 uint8 = 1
        _ = v30
    }
    if (0.0 == 3.14159) {
        switch "a" {
        case "a":
            var v31 float64 = 1.5
            _ = v31
        default:
            for fi32 := 0; fi32 < 20; fi32++ {
                for fi33 := 0; fi33 < 20; fi33++ {
                    switch (fi32 + 0) {
                    case 253:
                        var v34 bool = false
                        _ = v34
                    case 4294967295:
                        v35 := v4[9223372036854775807]
                        _ = v35
                    }
                    _ = fi33
                }
                v36 := v4[fi32]
                _ = fi32
                _ = v36
            }
        }
        v6 = ""
    } else {
        v13 += v13
        fmt.Println(v6)
        v12 = "\n\t"
    }
    v5 = v12
    switch v5 {
    case "a":
        v4[32768] = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
