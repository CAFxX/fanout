package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 79
    v3 := make(map[uint32]uint32)
    var v4 float32 = 0.0
    if true {
        if (!(42 >= 1)) {
            v4 = v4
            v2 = (v2 | (v2 - (v2 ^ v2)))
            v2 = (v2 % ((1) | 1))
        }
        v5 := []int{2147483647, 32767, 4294967295, -9223372036854775808}
        _ = v5
    } else {
        if (!true) {
            v4 += float32(0.0)
        }
    }
    fmt.Println(v2, v4)
    switch "" {
    case "\n\t":
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: 65535, F1: 2}
            var v8 int32 = -165
            var v9 int = fi6
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    default:
        if true {
            for fi10 := 0; fi10 < 20; fi10++ {
                v4 -= v4
                _ = fi10
            }
        }
    }
    v11 := GenS2{F0: 32767, F1: 0}
    v12 := GenS2{F0: 42, F1: 1}
    delete(v3, 2147483648)
    var v13 float64 = 0.1
    v14 := []int{}
    v13 -= v13
    for fi15 := 0; fi15 < 8; fi15++ {
        if (-342 > -494) {
            var v16 int = fi15
            _ = v16
        } else {
            fmt.Println(v13, fi15)
            v4 = 0.1
            var v17 bool = (((801 != -32768) && (!(!(true || (v2 != (v2 - v2)))))) && false)
            _ = v17
        }
        if ((fi15 <= (fi15 | (fi15 % ((-1) | 1)))) && (!((42 > 229) || true))) {
            switch "foo bar" {
            case "\n\t":
                v2 -= 128
            case "a":
                var v18 int64 = -341
                _ = v18
            case "foo bar":
                fmt.Println(v13, v2)
            default:
                var v19 string = ("\n\t" + "hello")
                _ = v19
            }
        }
        _ = fi15
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 string = "a"
            _ = fi21
            _ = v22
        }
        v4 = v4
        var v23 uint16 = 1
        _ = fi20
        _ = v23
    }
    var v24 uint = 18446744073709551615
    var v25 bool = (!false)
    for fi26 := 0; fi26 < 3; fi26++ {
        switch "a" {
        case "":
            for fi27 := 0; fi27 < 14; fi27++ {
                v28 := GenS3{F0: 1000}
                v1[781] = false
                switch "foo bar" {
                case "\n\t":
                    v2 = 127
                case "hello":
                    fmt.Println(v13)
                case "a":
                    v3[0] = 4294967295
                default:
                    v14 = append(v14, 42)
                    v14 = append(v14, fi27)
                    _ = len(v14)
                }
                v4 -= v4
                fmt.Println(v2)
                _ = fi27
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            v30 := GenS3{F0: -208}
            _ = fi29
            _ = v30
        }
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
