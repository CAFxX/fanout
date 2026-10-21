package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint32, p2 int32) int16 {
    var v5 int16 = 694
    var v6 bool = false
    fmt.Println(p0, v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (v5 - 1000)
}
func genFn1(p0 bool, p1 int64, p2 int16) int8 {
    var v7 int = -773
    v7 -= (v7 + 32768)
    var v8 uint16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -128
}

func main() {
    v1 := []uint32{2147483648, 462}
    var v2 uint8 = 2
    var v3 uint16 = 0
    var v4 uint8 = 128
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(v2, v4)
        v10 := GenS1{F0: (false || false), F1: -128, F2: fi9}
        v3 = 249
        _ = fi9
        _ = v10
    }
    _ = len(v1)
    for fi11 := 0; fi11 < 18; fi11++ {
        switch (fi11 - 9223372036854775807) {
        case 32768:
            if len(v1) > 0 {
                v1[0] = 206
            }
        case 128:
            v4 = 127
        case -65:
            v3 -= 1
        default:
            if false {
                var v12 uint = 0
                var v13 float32 = (float32(-2.25) * -2.25)
                _ = v12
                _ = v13
            }
        }
        _ = fi11
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v3 -= 0
    var v14 int8 = 25
    v15 := GenS2{F0: 2147483647}
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 uint32 = 546
            switch "" {
            case "":
                v1 = append(v1, 2147483648)
            case "foo bar":
                if false {
                    if len(v1) > 0 {
                        v1[0] = v18
                    }
                }
            }
            _ = fi17
            _ = v18
        }
        v1 = append(v1, uint32(v2))
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 float32 = 1.5
            _ = fi19
            _ = v20
        }
        _ = fi16
    }
    var v21 int64 = -40
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v14, v2)
    var v22 uint16 = v3
    var v23 int64 = (-(v21))
    fmt.Println(v21, v14, v22)
    for fi24 := 0; fi24 < 3; fi24++ {
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 int = fi24
            fmt.Println(v3, v14)
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    v27 := GenS2{F0: 442}
    for fi28 := 0; fi28 < 6; fi28++ {
        if len(v1) > 0 {
            v1[0] = 0
        }
        _ = fi28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
}
