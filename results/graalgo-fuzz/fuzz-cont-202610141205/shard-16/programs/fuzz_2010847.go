package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int32; F1 int8; F2 uint }
type GenS3 struct { F0 uint32; F1 bool; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int32; F2 int64 }
func genFn0(p0 bool, p1 uint64) uint64 {
    if true {
        p1 = (^(p1))
        for fi5 := 0; fi5 < 7; fi5++ {
            switch (fi5 | fi5) {
            case 42:
                p0 = p0
            case -1:
                v6 := []int{9223372036854775807, fi5}
                _ = v6
            case 518:
                switch 932 {
                case -112:
                    if ((fi5 - (fi5 ^ 42)) < (fi5 / ((-1) | 1))) {
                        var v7 int64 = -1
                        _ = v7
                    } else {
                        var v8 int32 = 1
                        _ = v8
                    }
                default:
                    p1 = 4294967295
                }
            default:
                for fi9 := 0; fi9 < 9; fi9++ {
                    switch "hello" {
                    case "hello":
                        v10 := []int{fi9}
                        _ = v10
                    case "a":
                        fmt.Println(fi5, fi9)
                    default:
                        var v11 float64 = 1.5
                        _ = v11
                    }
                    _ = fi9
                }
            }
            for fi12 := 0; fi12 < 17; fi12++ {
                p0 = (774 != 391)
                _ = fi12
            }
            fmt.Println(p0, p1, p0)
            _ = fi5
        }
    }
    v13 := GenS1{F0: (32767 == 32767), F1: (42 < 0)}
    switch "" {
    case "hello":
        p1 = 1
    case "\n\t":
        v14 := GenS2{F0: 2147483647, F1: 1, F2: 1}
        _ = v14
    case "a":
        switch 127 {
        case 32767:
            v15 := GenS2{F0: 687, F1: 100, F2: 4294967295}
            _ = v15
        default:
            p1 = (^(p1))
        }
    }
    _ = p0
    _ = p1
    _ = v13
    return p1
}
func genFn1(p0 uint64, p1 int) uint32 {
    v16 := []int{(p1 >> ((4) & 63))}
    p0 -= 42
    p1 = 42
    if ((-2.25 * ((((float64(1e10) / 1.0) * ((1.5 + -2.25) * 0.0)) + -2.25) + ((1.5 + 0.1) - 1.5))) > 1.5) {
        switch "a" {
        case "":
            p1 = 32768
        default:
            v17 := GenS1{F0: (false || true), F1: (0 == 9223372036854775807)}
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v16
    return 0
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]int)
    var v3 bool = (4294967295 < -643)
    var v4 int8 = 100
    v18 := []int{2147483647, 0}
    v4 -= int8(v1)
    if (v3 && true) {
        switch 128 {
        case 1:
            v19 := GenS3{F0: 4294967295, F1: ((-1 > -1) || ((v3 && v3) && (0.1 < (1.5 - (((0.1 - float64(0.0)) / 1.0) + 3.14159))))), F2: ((100 <= 42) && (v3 && v3))}
            _ = v19
        }
    } else {
        fmt.Println(v4, v4, v3)
    }
    delete(v2, 2)
    switch ("\n\t" + "foo bar") {
    case "hello":
        fmt.Println(v1, v1)
    default:
        var v20 uint32 = 1
        _ = v20
    }
    if (!(true && (4294967295 >= 1))) {
        v21 := GenS4{F0: ("a" + "hello")}
        _ = v21
    } else {
        v22 := GenS5{F0: (1 >= (^(v4))), F1: 625, F2: 0}
        _ = v22
    }
    _ = len(v2)
    v23 := GenS2{F0: 42, F1: v4, F2: 353}
    v24 := v2[1]
    _ = len(v2)
    var v25 bool = v3
    fmt.Println(v25, v25)
    _ = len(v2)
    for fi26 := 0; fi26 < 19; fi26++ {
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 string = ("foo bar" + ("\n\t" + "hello"))
            v24 += v24
            var v29 string = ((v28 + ("" + v28)) + "a")
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
