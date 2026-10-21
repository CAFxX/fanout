package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint; F1 int8; F2 int }
type GenS3 struct { F0 int8; F1 uint32 }

func main() {
    v1 := []string{("" + (("" + "hello") + "a")), "hello", "a", "a"}
    var v2 int = -419
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v5 uint64 = 1
    if (!(42 > 32767)) {
        var v6 int = v2
        var v7 uint32 = 4294967295
        _ = v6
        _ = v7
    } else {
        var v8 int32 = 0
        switch (^(v2)) {
        case 127:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            if true {
                var v10 float64 = float64((-2.25 / 1.5))
                var v11 int = (fi9 / ((fi9) | 1))
                v12 := GenS1{F0: 776, F1: true, F2: 2147483648}
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        _ = v8
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v3[42] = v2
    _ = len(v1)
    if true {
        v13 := GenS1{F0: 2, F1: true, F2: 65535}
        _ = v13
    }
    if true {
        if (!(!("foo bar" > ("foo bar" + "foo bar")))) {
            if false {
                var v14 uint32 = 32
                _ = v14
            }
        } else {
            v2 = 4294967295
        }
        var v15 uint32 = 444
        if (float64(3.14159) == float64(1.5)) {
            if (v2 <= (-(v2))) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v16 := GenS2{F0: 375, F1: 42, F2: 226}
                switch (v2 + (v2 | -1)) {
                case 9223372036854775807:
                    for fi17 := 0; fi17 < 18; fi17++ {
                        var v18 bool = false
                        switch (-(fi17)) {
                        case 4294967295:
                            var v19 uint = 18446744073709551615
                            _ = v19
                        default:
                            var v20 int = (-(v2))
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 18; fi21++ {
                            switch ("hello" + "foo bar") {
                            case "":
                                var v22 bool = (!true)
                                _ = v22
                            case "hello":
                                if v18 {
                                    if len(v1) > 0 {
                                        v1[0] = ("hello" + "foo bar")
                                    }
                                    fmt.Println(v2)
                                }
                            case "\n\t":
                                fmt.Println(v15, v5)
                            }
                            v18 = false
                            _ = fi21
                        }
                        _ = fi17
                        _ = v18
                    }
                case -864:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = v16
            }
            v15 = 666
        } else {
            v5 += (v5 * (v5 + (v5 + (v5 >> ((255) & 63)))))
        }
        _ = v15
    }
    v5 = v5
    var v23 int64 = -1
    v24 := GenS3{F0: -128, F1: 872}
    if (-2147483648 > -2147483648) {
        switch ("a" + "a") {
        case "hello":
            var v25 int32 = 42
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 uint = 18446744073709551615
        var v28 float32 = float32((0.0 - 1.5))
        v27 -= (v27 - (v27 + (v27 * (v27 | (v27 / ((0) | 1))))))
        _ = fi26
        _ = v27
        _ = v28
    }
    v5 += (v5 % (((^(v5))) | 1))
    for fi29 := 0; fi29 < 14; fi29++ {
        fmt.Println(v23, fi29)
        if true {
            var v30 int16 = 10
            _ = v30
        } else {
            fmt.Println(v23)
        }
        _ = fi29
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        _ = len(v4)
        fmt.Println(fi31, v2)
        _ = fi31
    }
    fmt.Println(v23, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
