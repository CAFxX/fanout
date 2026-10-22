package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 1
    v3 := make(map[int]string)
    var v4 uint16 = 0
    v5 := []int{v2, v2}
    switch v2 {
    case 128:
        if true {
            var v6 float32 = ((1e10 - ((((float32(((((1e10 / 1.0) * 0.0) - ((((0.1 - float64(-2.25)) - 1.5) - 1e10) + -2.25)) - 0.0)) + ((((3.14159 - ((-2.25 / 3.14159) + (3.14159 - -2.25))) / (1e10 + 1.5)) + 1e10) / 3.14159)) - (0.0 + -2.25)) / (-2.25 - float32(((float64(0.1) * ((0.1 * ((1e10 * -2.25) / 1.0)) + 3.14159)) * 3.14159)))) / -2.25)) / (-2.25 - (0.0 - ((-2.25 / (((-2.25 * 1e10) / 1.0) + -2.25)) + -2.25))))
            _ = v6
        }
    default:
        switch (v2 | v2) {
        case 2147483647:
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 string = (("foo bar" + "hello") + "")
                v9 := GenS1{F0: 0, F1: v8, F2: 1}
                if false {
                    if true {
                        var v10 bool = true
                        v11 := GenS1{F0: 42, F1: "hello", F2: 518}
                        switch (fi7 << ((0) & 63)) {
                        case 128:
                            v8 = "\n\t"
                        case -24:
                            for fi12 := 0; fi12 < 12; fi12++ {
                                if (1e10 >= 0.1) {
                                    var v13 float64 = (3.14159 + (1.5 * float64((3.14159 + float32(float64(-2.25))))))
                                    _ = v13
                                }
                                var v14 uint64 = 0
                                _ = fi12
                                _ = v14
                            }
                        }
                        _ = v10
                        _ = v11
                    }
                    if (!true) {
                        fmt.Println(fi7, v2)
                    }
                } else {
                    if (false || true) {
                        var v15 int64 = 42
                        _ = v15
                    }
                    switch fi7 {
                    case 1:
                        v8 = v8
                    }
                    var v16 uint8 = 255
                    _ = v16
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        case -1:
            var v17 int16 = 0
            _ = v17
        case 32767:
            var v18 float64 = ((3.14159 / 1e10) + (-2.25 / ((((1.5 / -2.25) / ((3.14159 * 3.14159) + 1e10)) * float64(1.5)) - 0.0)))
            _ = v18
        default:
            if len(v5) > 0 {
                v5[0] = v2
            }
        }
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        if len(v5) > 0 {
            v5[0] = 9223372036854775807
        }
        _ = len(v3)
        _ = fi19
    }
    v20 := GenS1{F0: 1, F1: ("" + "foo bar"), F2: 655}
    if ((1 >= 18446744073709551615) || (0.1 >= 1.5)) {
        v4 = v4
    }
    v21 := GenS1{F0: 1, F1: "a", F2: 4294967295}
    for fi22 := 0; fi22 < 4; fi22++ {
        if true {
            if false {
                fmt.Println(fi22, v2, v2)
            }
            var v23 uint64 = 1
            for fi24 := 0; fi24 < 20; fi24++ {
                v25 := GenS2{F0: -963}
                _ = fi24
                _ = v25
            }
            _ = v23
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi22
    }
    var v26 string = "hello"
    var v27 int64 = 414
    switch v26 {
    case "hello":
        var v28 string = ""
        _ = v28
    case "foo bar":
        var v29 int32 = 1
        _ = v29
    case "a":
        var v30 uint64 = 4294967295
        _ = v30
    }
    v27 = 715
    var v31 int16 = -317
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
