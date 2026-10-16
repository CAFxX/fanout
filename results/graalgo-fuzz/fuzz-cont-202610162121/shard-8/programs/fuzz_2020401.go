package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 uint32, p1 bool) int32 {
    var v5 bool = true
    fmt.Println(v5, v5)
    fmt.Println(v5)
    _ = p0
    _ = p1
    _ = v5
    return -452
}
func genFn1() uint16 {
    var v6 float64 = ((3.14159 * (((((-2.25 + 0.0) / (0.1 * 0.1)) + 1e10) + 1.5) - 0.0)) + ((0.1 - -2.25) - 1.5))
    v6 += (v6 + (v6 + 1e10))
    var v7 string = "a"
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: -128, F1: 1}
        v6 -= -2.25
        _ = fi8
        _ = v9
    }
    v7 = v7
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    v4 := []uint32{2}
    switch ("foo bar" + ("" + ("a" + "foo bar"))) {
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        var v10 bool = (!(110 <= -1))
        _ = v10
    case "":
        var v11 int32 = -1
        _ = v11
    default:
        switch 0 {
        case -431:
            switch "\n\t" {
            case "\n\t":
                if (!true) {
                    v1 -= (v1 ^ v1)
                }
            }
        case 32767:
            if true {
                var v12 uint8 = (v1 % (((^(v1))) | 1))
                v1 -= 184
                _ = v12
            }
        case 128:
            if (!((2147483647 > -2147483648) || (127 != 127))) {
                if (!true) {
                    var v13 int32 = genFn0(65535, (!(false || ("foo bar" < "foo bar"))))
                    _ = v13
                }
                var v14 uint64 = 2
                _ = v14
            }
        default:
            var v15 string = ("\n\t" + (("foo bar" + ("hello" + "a")) + ("a" + ("hello" + ("a" + "foo bar")))))
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 uint16 = 65535
        if true {
            var v18 int16 = 1000
            switch fi16 {
            case 128:
                v19 := GenS2{F0: -128, F1: -128}
                _ = v19
            case 32768:
                var v20 int32 = -767
                _ = v20
            case -1:
                v4 = append(v4, 2147483647)
            default:
                var v21 uint32 = 408
                _ = v21
            }
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    if (3.14159 <= float64(-2.25)) {
        switch "" {
        case "\n\t":
            for fi22 := 0; fi22 < 11; fi22++ {
                v23 := GenS2{F0: 100, F1: 100}
                if ("" == "a") {
                    v1 = 98
                }
                _ = fi22
                _ = v23
            }
        case "a":
            for fi24 := 0; fi24 < 17; fi24++ {
                v1 = (v1 % ((255) | 1))
                _ = fi24
            }
        case "foo bar":
            fmt.Println(v1)
        }
        fmt.Println(v1)
    }
    fmt.Println(v1)
    var v25 uint32 = 0
    switch -1 {
    case 451:
        if (!(!true)) {
            fmt.Println(v25, v1)
        }
    case -9223372036854775808:
        var v26 int16 = 1000
        _ = v26
    case 0:
        var v27 bool = false
        _ = v27
    default:
        _ = len(v4)
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        v25 = 187
        _ = fi28
    }
    v1 += v1
    switch ("a" + ("\n\t" + "hello")) {
    case "foo bar":
        v29 := GenS2{F0: -1, F1: 1}
        _ = v29
    case "\n\t":
        var v30 int8 = -128
        _ = v30
    default:
        switch (("hello" + (("\n\t" + (((("" + "a") + "a") + "") + ("foo bar" + "foo bar"))) + "a")) + ("a" + ("\n\t" + "foo bar"))) {
        case "foo bar":
            v31 := GenS1{F0: 81, F1: int64(v1)}
            _ = v31
        case "\n\t":
            fmt.Println(v1)
        }
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        fmt.Println(fi32)
        _ = fi32
    }
    fmt.Println(v25)
    var v33 uint = 0
    v25 = (v25 << ((656) & 63))
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v4
}
