package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int8; F2 uint8 }
type GenS3 struct { F0 int8 }
func genFn0() int64 {
    v5 := GenS1{F0: false}
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 float64 = float64(float32(1.5))
        switch ("" + "\n\t") {
        case "a":
            var v8 bool = false
            _ = v8
        default:
            v9 := []int{}
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(fi11, fi11)
            }
            _ = fi11
        }
        _ = fi10
    }
    v12 := GenS2{F0: -9223372036854775808, F1: 1, F2: 2}
    _ = v12
    _ = v5
    return -82
}
func genFn1(p0 int, p1 uint64, p2 uint8) int16 {
    p0 = (p0 ^ 680)
    p0 = (p0 / (((p0 / ((325) | 1))) | 1))
    if (true || ((1 <= 2) && true)) {
        p2 -= (p2 | 0)
        fmt.Println(p2)
        p2 -= (^(p2))
    }
    v13 := []int{(-(p0))}
    if false {
        var v14 int64 = -9223372036854775808
        v14 -= v14
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 uint64 = 42
    v4 := make(map[int]bool)
    fmt.Println(v3)
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(fi15, fi15)
        _ = fi15
    }
    v3 = (v3 - (v3 ^ v3))
    fmt.Println(v3)
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 9; fi17++ {
            switch fi16 {
            case 479:
                if (789 == 939) {
                    v3 += (v3 + (v3 >> ((4294967295) & 63)))
                }
            case -2147483648:
                var v18 float32 = (0.1 + ((((-2.25 - 1e10) * float32(1.5)) - (1.5 + float32(float64(0.0)))) + 0.0))
                _ = v18
            default:
                var v19 uint32 = 0
                _ = v19
            }
            _ = fi17
        }
        _ = fi16
    }
    if (true && ((0 <= 18446744073709551615) || true)) {
        v3 -= v3
        switch 998 {
        case 0:
            if (!(false || false)) {
                v3 -= 2
                v3 += 1
            }
        case 4294967295:
            for fi20 := 0; fi20 < 12; fi20++ {
                for fi21 := 0; fi21 < 9; fi21++ {
                    v3 = 858
                    v22 := []int{699, fi20}
                    _ = fi21
                    _ = v22
                }
                v23 := GenS2{F0: 0, F1: -42, F2: 255}
                _ = fi20
                _ = v23
            }
        default:
            var v24 string = "\n\t"
            _ = v24
        }
        var v25 int64 = genFn0()
        _ = v25
    } else {
        _ = len(v2)
        switch ((("hello" + ("foo bar" + ((("a" + ((("a" + "hello") + "foo bar") + "hello")) + ("foo bar" + "")) + "\n\t"))) + "hello") + ("foo bar" + "")) {
        case "foo bar":
            var v26 bool = (!(18446744073709551615 != 0))
            _ = v26
        case "hello":
            fmt.Println(v3)
        }
    }
    if true {
        v27 := GenS3{F0: 0}
        _ = v27
    } else {
        var v28 int32 = -766
        _ = v28
    }
    if (true || true) {
        v29 := GenS3{F0: 28}
        _ = v29
    }
    v30 := GenS1{F0: true}
    var v31 int64 = -9223372036854775808
    for fi32 := 0; fi32 < 15; fi32++ {
        switch "foo bar" {
        case "\n\t":
            delete(v1, "\n\t")
        case "":
            switch (fi32 - (fi32 % ((-388) | 1))) {
            case -1:
                v33 := GenS2{F0: 9223372036854775807, F1: 100, F2: 42}
                _ = v33
            }
        case "hello":
            var v34 string = "foo bar"
            _ = v34
        }
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
