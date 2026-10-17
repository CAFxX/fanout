package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint32; F2 string }
type GenS4 struct { F0 uint; F1 uint16; F2 int8 }
func genFn0(p0 int16) int32 {
    if false {
        v7 := GenS1{F0: false, F1: (p0 / ((p0) | 1))}
        if true {
            v8 := []int{2147483647, -593}
            p0 -= 0
            _ = v8
        }
        switch -2147483648 {
        case -710:
            p0 -= p0
        case 42:
            p0 += p0
        default:
            var v9 int8 = 19
            _ = v9
        }
        _ = v7
    }
    if false {
        switch "foo bar" {
        case "\n\t":
            var v10 uint64 = 2
            _ = v10
        case "a":
            v11 := GenS1{F0: false, F1: 32767}
            _ = v11
        case "hello":
            fmt.Println(p0)
        }
        var v12 int32 = 42
        _ = v12
    }
    switch 416 {
    case 127:
        p0 = (p0 - p0)
    case 202:
        p0 = 1000
    }
    _ = p0
    return -472
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{42}
    v3 := make(map[int]string)
    var v4 int16 = -1
    v5 := []int{32768, -713, 2147483647, 128}
    v6 := []uint32{2147483648, 2147483647, 2147483648}
    for fi13 := 0; fi13 < 3; fi13++ {
        switch "a" {
        case "hello":
            v4 -= -1
        case "foo bar":
            switch fi13 {
            case -490:
                var v14 uint = 923
                _ = v14
            case 4294967295:
                var v15 float64 = (-2.25 * -2.25)
                _ = v15
            case -1:
                fmt.Println(v4, v4)
            }
        case "":
            var v16 int64 = 9223372036854775807
            _ = v16
        }
        var v17 uint8 = 209
        fmt.Println(v17, v17)
        _ = fi13
        _ = v17
    }
    v4 = (v4 >> ((191) & 63))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch 42 {
    case 0:
        v18 := GenS1{F0: true, F1: v4}
        _ = v18
    case 9223372036854775807:
        if (4294967295 == 370) {
            v19 := GenS1{F0: true, F1: (v4 >> ((396) & 63))}
            fmt.Println(v4)
            if len(v6) > 0 {
                v6[0] = 65535
            }
            _ = v19
        } else {
            switch "hello" {
            case "foo bar":
                v20 := GenS2{F0: true}
                _ = v20
            case "a":
                _ = len(v3)
            default:
                for fi21 := 0; fi21 < 7; fi21++ {
                    if ((-(v4)) < 0) {
                        v22 := GenS3{F0: true, F1: 2, F2: ("foo bar" + "")}
                        _ = v22
                    }
                    fmt.Println(fi21, v4)
                    var v23 uint32 = 65535
                    _ = fi21
                    _ = v23
                }
            }
            v3[2147483647] = ("foo bar" + "a")
        }
    case -202:
        v24 := GenS4{F0: 18446744073709551615, F1: 47, F2: -32}
        _ = v24
    }
    fmt.Println(v4)
    var v25 int64 = 9223372036854775807
    fmt.Println(v25)
    if (-128 >= -1) {
        v25 += v25
    }
    v25 -= (-(v25))
    for fi26 := 0; fi26 < 1; fi26++ {
        switch "" {
        case "foo bar":
            var v27 int64 = (v25 - (v25 << ((1) & 63)))
            _ = v27
        case "hello":
            fmt.Println(v25, v4, fi26)
        case "a":
            switch "hello" {
            case "foo bar":
                var v28 uint = 488
                _ = v28
            case "hello":
                var v29 uint32 = 249
                _ = v29
            case "a":
                for fi30 := 0; fi30 < 1; fi30++ {
                    delete(v3, (fi26 | -2147483648))
                    v25 += (v25 >> ((255) & 63))
                    _ = fi30
                }
            }
        }
        _ = fi26
    }
    switch "" {
    case "\n\t":
        v3[274] = "hello"
    case "hello":
        if false {
            v31 := GenS2{F0: (!false)}
            for fi32 := 0; fi32 < 14; fi32++ {
                for fi33 := 0; fi33 < 14; fi33++ {
                    v34 := GenS4{F0: 0, F1: 255, F2: 127}
                    delete(v1, (fi32 ^ fi33))
                    _ = fi33
                    _ = v34
                }
                _ = fi32
            }
            _ = v31
        }
    }
    var v35 int8 = -107
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
