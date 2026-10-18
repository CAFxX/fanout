package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint32; F1 int16; F2 int8 }
func genFn0(p0 int64, p1 string, p2 int16) string {
    fmt.Println(p1)
    v6 := []int{42}
    if (-32 > 0) {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (("foo bar" + ("foo bar" + "")) + p1)
}
func genFn1(p0 uint32, p1 bool) string {
    switch 32767 {
    case 128:
        switch (("\n\t" + "a") + ("" + ("\n\t" + "a"))) {
        case "hello":
            if p1 {
                fmt.Println(p0, p1)
            }
        case "a":
            var v7 uint32 = (p0 & p0)
            _ = v7
        }
    case -703:
        var v8 uint32 = p0
        _ = v8
    default:
        var v9 bool = (-1 == -473)
        _ = v9
    }
    p0 += (p0 & 2)
    var v10 int8 = 0
    p1 = p1
    p0 = 82
    _ = p0
    _ = p1
    _ = v10
    return (((("hello" + "hello") + "foo bar") + (("" + ("a" + "a")) + ("\n\t" + ""))) + ("hello" + (("\n\t" + ("hello" + "a")) + "a")))
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "\n\t"
    v3 := []uint32{}
    var v4 uint8 = 2
    var v5 int = -9223372036854775808
    v11 := GenS1{F0: 255, F1: v5, F2: 2}
    v12 := GenS1{F0: (v4 & v4), F1: v5, F2: 18446744073709551615}
    switch (v2 + v2) {
    case "a":
        v4 += (v4 ^ v4)
    }
    fmt.Println(v5)
    v13 := GenS1{F0: v4, F1: v5, F2: 42}
    switch v2 {
    case "\n\t":
        switch v2 {
        case "":
            fmt.Println(v2, v5, v2)
        case "foo bar":
            v5 = (v5 & v5)
        case "a":
            switch v5 {
            case -937:
                var v14 int16 = 1000
                _ = v14
            case -2147483648:
                v4 += (v4 << ((0) & 63))
            case -9223372036854775808:
                switch v2 {
                case "":
                    var v15 int8 = 123
                    _ = v15
                case "hello":
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v17 := v1[(v5 | fi16)]
                        for fi18 := 0; fi18 < 7; fi18++ {
                            _ = len(v3)
                            v19 := GenS2{F0: v2, F1: (v4 >> ((18446744073709551615) & 63))}
                            _ = fi18
                            _ = v19
                        }
                        var v20 uint = 4294967295
                        _ = fi16
                        _ = v17
                        _ = v20
                    }
                default:
                    var v21 string = "hello"
                    _ = v21
                }
            }
        default:
            v22 := GenS3{F0: 1, F1: -199, F2: 221}
            _ = v22
        }
    case "a":
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS4{F0: -9223372036854775808, F1: int8(fi23)}
            _ = fi23
            _ = v24
        }
    default:
        _ = len(v1)
    }
    var v25 int64 = 42
    var v26 int8 = 13
    if true {
        v27 := GenS5{F0: 0, F1: 0, F2: v26}
        switch (v5 ^ v5) {
        case 32767:
            var v28 float32 = 1.5
            _ = v28
        case 52:
            v29 := GenS1{F0: (v4 - v4), F1: v5, F2: 2}
            _ = v29
        case 32768:
            for fi30 := 0; fi30 < 20; fi30++ {
                v31 := GenS1{F0: (^(v4)), F1: v5, F2: 1}
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
                _ = fi30
                _ = v31
            }
        }
        for fi32 := 0; fi32 < 7; fi32++ {
            _ = len(v1)
            _ = fi32
        }
        _ = v27
    }
    v26 += v26
    var v33 uint16 = 424
    if true {
        var v34 bool = ("foo bar" < ("a" + "hello"))
        _ = v34
    } else {
        var v35 int32 = -2147483648
        v4 += 42
        _ = v35
    }
    var v36 int64 = (v25 & (^(v25)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v5
}
