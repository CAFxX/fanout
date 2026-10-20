package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint64 }
func genFn0() uint16 {
    var v5 int = -1
    fmt.Println(v5)
    switch v5 {
    case 32767:
        v5 -= (v5 & -522)
    case 406:
        v5 = v5
    }
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: 4294967295}
    _ = v5
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int64, p1 uint64, p2 int32) int32 {
    fmt.Println(p0, p2)
    v8 := []int{}
    p2 += p2
    var v9 uint8 = 6
    for fi10 := 0; fi10 < 15; fi10++ {
        p1 = p1
        fmt.Println(p2)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p2
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 bool = (true && false)
    var v3 float32 = ((1e10 * 0.0) - 3.14159)
    v4 := []int{-244, 1}
    var v11 int64 = v1
    v3 += (3.14159 * 0.1)
    switch ("" + "\n\t") {
    case "hello":
        switch 128 {
        case -948:
            for fi12 := 0; fi12 < 12; fi12++ {
                fmt.Println(v3, v11, fi12)
                _ = fi12
            }
        }
    case "a":
        switch "a" {
        case "\n\t":
            switch "hello" {
            case "":
                switch "a" {
                case "a":
                    var v13 int32 = 425
                    _ = v13
                }
            case "hello":
                var v14 float64 = ((0.0 - 1.5) * 0.1)
                _ = v14
            case "foo bar":
                v15 := GenS2{F0: -9223372036854775808, F1: v1, F2: 557}
                _ = v15
            default:
                if (179 <= 2) {
                    fmt.Println(v3)
                } else {
                    v11 = (v11 ^ (v11 | (v11 ^ (^(v11)))))
                    _ = len(v4)
                }
            }
        case "":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    default:
        var v16 float64 = (0.1 / 0.1)
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        if (127 != 42) {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 int8 = -1
                if ((v11 % ((1) | 1)) < -1) {
                    var v20 int = (^(fi17))
                    var v21 int8 = 42
                    v22 := GenS3{F0: 1}
                    _ = v20
                    _ = v21
                    _ = v22
                }
                v2 = v2
                _ = fi18
                _ = v19
            }
            if (!(!(18446744073709551615 > 209))) {
                v2 = false
                v11 += (^(v11))
            }
        } else {
            var v23 float32 = float32((1e10 / 3.14159))
            switch "hello" {
            case "hello":
                var v24 int = fi17
                _ = v24
            case "":
                v25 := GenS1{F0: 167}
                _ = v25
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v23
        }
        var v26 string = "a"
        _ = fi17
        _ = v26
    }
    var v27 float32 = float32(0.0)
    var v28 float64 = -2.25
    var v29 int = -60
    switch "" {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            var v30 float64 = v28
            _ = v30
        case "a":
            var v31 int = (^(v29))
            _ = v31
        default:
            if (("a" == ("\n\t" + "hello")) || true) {
                if v2 {
                    v32 := GenS4{F0: 1, F1: 243}
                    _ = v32
                }
                for fi33 := 0; fi33 < 9; fi33++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi33
                }
                for fi34 := 0; fi34 < 17; fi34++ {
                    v3 -= ((1e10 / v27) * (v3 * 1e10))
                    _ = len(v4)
                    _ = fi34
                }
            }
        }
    case "foo bar":
        var v35 float32 = v27
        _ = v35
    }
    if v2 {
        var v36 uint16 = 123
        _ = v36
    } else {
        var v37 float32 = 3.14159
        for fi38 := 0; fi38 < 20; fi38++ {
            for fi39 := 0; fi39 < 2; fi39++ {
                v11 = (v1 / ((v11) | 1))
                fmt.Println(v11)
                _ = fi39
            }
            _ = fi38
        }
        _ = v37
    }
    _ = len(v4)
    fmt.Println(v1, v28, v29)
    fmt.Println(v2, v11, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
