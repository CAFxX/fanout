package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int }
type GenS2 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint16) uint64 {
    v5 := []int{2147483647, 0, 0, 42}
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            p1 = p1
            var v8 int = (fi6 << ((496) & 63))
            switch ((("\n\t" + "hello") + "") + "\n\t") {
            case "\n\t":
                switch (("\n\t" + "a") + ("hello" + "\n\t")) {
                case "hello":
                    p1 += p1
                default:
                    p1 = 464
                }
            case "foo bar":
                fmt.Println(v8, fi7)
            case "hello":
                switch "hello" {
                case "":
                    var v9 int16 = 1000
                    _ = v9
                case "foo bar":
                    fmt.Println(p1)
                case "\n\t":
                    var v10 float64 = (0.0 + 1e10)
                    _ = v10
                }
            default:
                v5 = append(v5, fi6)
            }
            _ = fi7
            _ = v8
        }
        if p0 {
            var v11 int16 = -1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v11
        } else {
            p0 = (255 >= 237)
            for fi12 := 0; fi12 < 17; fi12++ {
                switch "" {
                case "hello":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                var v13 uint = 4294967295
                v14 := GenS1{F0: "\n\t", F1: ("hello" + (("a" + ("foo bar" + ("\n\t" + ("hello" + "")))) + ("hello" + "foo bar"))), F2: fi12}
                _ = fi12
                _ = v13
                _ = v14
            }
            p1 -= 426
        }
        _ = fi6
    }
    _ = len(v5)
    p1 = (^(p1))
    var v15 uint = 4294967295
    _ = p0
    _ = p1
    _ = v15
    _ = v5
    return 2
}
func genFn1(p0 int64) int16 {
    switch "" {
    case "foo bar":
        p0 -= p0
    case "\n\t":
        p0 += (p0 * (p0 - p0))
    case "a":
        var v16 uint = 65535
        _ = v16
    }
    var v17 uint64 = 0
    v18 := GenS1{F0: "\n\t", F1: (("foo bar" + "a") + "a"), F2: 128}
    p0 += p0
    _ = p0
    _ = v17
    _ = v18
    return 0
}

func main() {
    var v1 int32 = 2147483647
    var v2 string = "\n\t"
    var v3 bool = ((!((193 <= 255) || false)) && (!(8 == 100)))
    var v4 int16 = -1
    v19 := []int{32767, 4294967295, 2147483647, -9223372036854775808}
    switch v2 {
    case "":
        fmt.Println(v1, v1)
    case "\n\t":
        v20 := GenS1{F0: v2, F1: (v2 + v2), F2: 4294967295}
        _ = v20
    case "foo bar":
        if len(v19) > 0 {
            v19[0] = 0
        }
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    if (!true) {
        for fi21 := 0; fi21 < 10; fi21++ {
            v2 = "a"
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 int8 = -113
                if ((float64((1e10 * 1e10)) * float64(float32((0.0 * (float64(float32(-2.25)) + float64(1.5)))))) == 0.0) {
                    v4 -= genFn1(42)
                    var v24 string = v2
                    _ = v24
                } else {
                    var v25 float32 = float32((-2.25 / (((-2.25 / 1e10) + ((-2.25 + ((0.0 / float64((0.1 * -2.25))) / float64((-2.25 + (0.0 - (float32(0.0) + 1e10)))))) - 1e10)) + ((1.5 - 1.5) * 0.0))))
                    _ = v25
                }
                var v26 uint8 = 2
                _ = fi22
                _ = v23
                _ = v26
            }
            _ = fi21
        }
        switch v2 {
        case "\n\t":
            v27 := GenS1{F0: v2, F1: v2, F2: -9223372036854775808}
            _ = v27
        case "":
            var v28 int16 = (v4 + (v4 / ((v4) | 1)))
            _ = v28
        }
        fmt.Println(v4)
    }
    v4 = (-(v4))
    for fi29 := 0; fi29 < 3; fi29++ {
        v30 := GenS2{F0: 33}
        _ = fi29
        _ = v30
    }
    v1 = (v1 | (v1 + (v1 % ((v1) | 1))))
    v4 -= (-(v4))
    var v31 string = ("foo bar" + v2)
    v1 = v1
    var v32 uint16 = 0
    var v33 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
