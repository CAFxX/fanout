package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
func genFn0(p0 int8, p1 uint, p2 int64) int16 {
    var v4 uint8 = 127
    var v5 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1000
}

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{2147483648, 4294967295, 2147483648}
    v3 := []string{("hello" + ("foo bar" + "foo bar")), ("foo bar" + ((("" + "") + "a") + "\n\t"))}
    var v6 uint8 = 0
    v7 := GenS1{F0: "hello", F1: 1}
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(v6, fi9, fi8)
            if len(v2) > 0 {
                v2[0] = 42
            }
            v10 := GenS1{F0: ("foo bar" + ("foo bar" + ("" + "a"))), F1: 42}
            _ = fi9
            _ = v10
        }
        if true {
            var v11 string = "a"
            for fi12 := 0; fi12 < 17; fi12++ {
                _ = len(v3)
                if false {
                    fmt.Println(v1, v1, fi12)
                    var v13 int8 = 127
                    fmt.Println(fi8, v6)
                    v11 = "hello"
                    _ = v13
                } else {
                    var v14 int8 = 0
                    var v15 int = fi12
                    v1 += (0.0 * v1)
                    _ = v14
                    _ = v15
                }
                v1 = (1.5 * 1.5)
                _ = fi12
            }
            switch v11 {
            case "\n\t":
                var v16 uint16 = 761
                _ = v16
            default:
                var v17 int = fi8
                _ = v17
            }
            _ = v11
        } else {
            if (true || (false || (1 < -73))) {
                _ = len(v2)
                var v18 float32 = 1.5
                _ = v18
            }
            if false {
                v6 = (v6 ^ (^(v6)))
            }
            v1 = (v1 + (v1 - v1))
        }
        v1 += v1
        _ = fi8
    }
    if true {
        if (2147483647 != 2) {
            v19 := GenS1{F0: (("" + "\n\t") + "hello"), F1: 342}
            _ = v19
        }
    }
    if (-238 > 1) {
        v20 := GenS1{F0: ((("hello" + ("hello" + ("foo bar" + "a"))) + "foo bar") + "a"), F1: 42}
        _ = v20
    }
    var v21 uint = 0
    for fi22 := 0; fi22 < 1; fi22++ {
        _ = len(v3)
        fmt.Println(fi22, v21, v1)
        _ = fi22
    }
    v6 -= v6
    var v23 int64 = 9223372036854775807
    for fi24 := 0; fi24 < 6; fi24++ {
        _ = len(v2)
        switch ((("" + "a") + ("a" + "\n\t")) + "\n\t") {
        case "\n\t":
            v1 = (float32(1e10) * (v1 / float32(-2.25)))
        default:
            fmt.Println(v1, v23, v23)
        }
        _ = fi24
    }
    var v25 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v25
    _ = v3
    _ = v6
    _ = v7
}
