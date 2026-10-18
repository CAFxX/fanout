package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int16; F1 int; F2 bool }
func genFn0(p0 int32, p1 uint32) int8 {
    p0 -= (^(p0))
    v5 := GenS1{F0: "a", F1: (("foo bar" + "a") + ("\n\t" + "foo bar"))}
    v6 := []int{1, int(p1)}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 100
}

func main() {
    var v1 uint32 = 42
    var v2 int = 32767
    var v3 float64 = (0.0 / ((float64(((float32((0.1 + 1e10)) - -2.25) - -2.25)) + -2.25) / 1.0))
    v4 := make(map[string]bool)
    delete(v4, "a")
    var v7 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            switch (fi9 / ((fi8) | 1)) {
            case -985:
                switch (fi9 | (fi8 << ((18446744073709551615) & 63))) {
                case 127:
                    var v10 uint64 = 4294967295
                    _ = v10
                case 32768:
                    if false {
                        switch "\n\t" {
                        case "\n\t":
                            _ = len(v4)
                        case "foo bar":
                            v3 -= v3
                        case "hello":
                            v11 := GenS1{F0: "hello", F1: ""}
                            _ = v11
                        default:
                            v2 = fi8
                        }
                    } else {
                        fmt.Println(fi9)
                    }
                }
            case 0:
                if (17 >= 0) {
                    fmt.Println(fi8, v2, v7)
                    var v12 int16 = 754
                    v13 := GenS1{F0: "foo bar", F1: ("\n\t" + "foo bar")}
                    _ = v12
                    _ = v13
                }
            case -1:
                var v14 bool = (("a" + ("\n\t" + "a")) != ("a" + "\n\t"))
                _ = v14
            default:
                v15 := GenS1{F0: "foo bar", F1: ""}
                _ = v15
            }
            v16 := []int{(fi8 >> ((uint(v1)) & 63)), fi8, 734, 32767}
            _ = fi9
            _ = v16
        }
        var v17 int8 = -128
        _ = fi8
        _ = v17
    }
    v18 := []int{v2}
    switch "" {
    case "foo bar":
        delete(v4, "")
    case "hello":
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(v1, v1, v2)
            for fi20 := 0; fi20 < 18; fi20++ {
                if len(v18) > 0 {
                    v18[0] = 128
                }
                if true {
                    var v21 uint8 = 128
                    if false {
                        var v22 uint = 1
                        _ = v22
                    } else {
                        var v23 uint64 = 683
                        var v24 int64 = (^(v7))
                        _ = v23
                        _ = v24
                    }
                    _ = v21
                }
                _ = fi20
            }
            _ = fi19
        }
    case "\n\t":
        v2 = v2
    }
    v25 := GenS2{F0: 32767, F1: 127, F2: false}
    if (738 >= 255) {
        v1 = v1
        if true {
            var v26 int8 = -33
            var v27 string = "hello"
            _ = v26
            _ = v27
        } else {
            var v28 float64 = 1.5
            _ = v28
        }
    } else {
        fmt.Println(v2, v3)
    }
    switch (v2 << ((534) & 63)) {
    case 1:
        v3 = 3.14159
    case -787:
        v29 := GenS1{F0: "hello", F1: "a"}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
