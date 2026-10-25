package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int16; F2 bool }
type GenS4 struct { F0 uint8; F1 uint; F2 int }
func genFn0(p0 uint32, p1 int64, p2 uint64) int16 {
    p1 = 42
    v7 := []int{-788, 32768}
    if true {
        p2 = (p2 * p2)
        for fi8 := 0; fi8 < 9; fi8++ {
            if len(v7) > 0 {
                v7[0] = 9223372036854775807
            }
            _ = fi8
        }
        v9 := GenS1{F0: 1}
        _ = v9
    } else {
        var v10 string = ("foo bar" + ("" + "\n\t"))
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = v10
    }
    if len(v7) > 0 {
        v7[0] = -693
    }
    v11 := GenS2{F0: p0, F1: 1, F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return -695
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 1.5
    var v3 int16 = -1
    v4 := []string{"hello", "hello"}
    var v5 string = "hello"
    var v6 uint16 = 0
    fmt.Println(v3, v6, v5)
    for fi12 := 0; fi12 < 19; fi12++ {
        v3 = (v3 << ((2) & 63))
        switch fi12 {
        case 277:
            v4 = append(v4, ("foo bar" + v5))
        case 128:
            switch fi12 {
            case 9223372036854775807:
                v13 := GenS3{F0: 2, F1: v3, F2: true}
                _ = v13
            case 0:
                v6 -= v6
            case 32768:
                v14 := GenS4{F0: 224, F1: 401, F2: (fi12 / ((fi12) | 1))}
                _ = v14
            default:
                var v15 bool = (953 == 144)
                _ = v15
            }
        }
        fmt.Println(fi12)
        _ = fi12
    }
    v6 = (v6 << ((690) & 63))
    fmt.Println(v2, v5, v5)
    v6 = (^(v6))
    if false {
        v5 = v5
    }
    v2 -= v2
    if true {
        switch -9223372036854775808 {
        case 643:
            switch (v5 + ((v5 + "hello") + "foo bar")) {
            case "a":
                var v16 uint32 = 4294967295
                _ = v16
            case "foo bar":
                v17 := GenS1{F0: 1}
                _ = v17
            case "hello":
                v6 = (v6 % ((v6) | 1))
            }
        case 996:
            v18 := GenS4{F0: 0, F1: 4294967295, F2: -679}
            _ = v18
        case -2147483648:
            v3 += -32768
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            if (0.0 > -2.25) {
                switch "" {
                case "foo bar":
                    v20 := GenS2{F0: 65535, F1: 0, F2: 1}
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 3; fi21++ {
                        fmt.Println(v3)
                        var v22 uint = 255
                        fmt.Println(v3)
                        _ = fi21
                        _ = v22
                    }
                }
            }
            _ = fi19
        }
    }
    var v23 uint16 = 255
    var v24 uint64 = 480
    var v25 uint32 = 2147483648
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 int16 = v3
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        if true {
            switch v5 {
            case "foo bar":
                var v29 uint32 = v25
                _ = v29
            case "hello":
                var v30 int16 = v3
                _ = v30
            }
        }
        fmt.Println(v3)
        _ = fi28
    }
    switch "hello" {
    case "\n\t":
        v31 := GenS3{F0: 105, F1: genFn0(v25, 1, 929), F2: (uint(v25) <= 255)}
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
