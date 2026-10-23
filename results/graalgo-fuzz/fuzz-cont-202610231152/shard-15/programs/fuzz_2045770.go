package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int32 }
func genFn0(p0 int64) uint64 {
    var v7 int16 = 42
    fmt.Println(p0, p0)
    var v8 uint16 = 151
    for fi9 := 0; fi9 < 4; fi9++ {
        if false {
            p0 = -9223372036854775808
            v10 := []int{-9223372036854775808, (fi9 + (fi9 & 643))}
            _ = v10
        }
        fmt.Println(fi9, fi9)
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 uint, p1 uint32, p2 bool) string {
    p0 = (p0 - 42)
    v11 := []int{465, -1, 477}
    var v12 uint64 = 2
    v12 = v12
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return "hello"
}

func main() {
    var v1 int32 = 1
    v2 := make(map[int]uint32)
    var v3 uint64 = 42
    v4 := []uint32{}
    v5 := make(map[uint32]bool)
    v6 := []string{"\n\t", ""}
    v6 = append(v6, ("" + "foo bar"))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v13 int16 = 42
    if (-1 > -108) {
        v14 := v5[2147483648]
        v13 = (v13 + (v13 - v13))
        _ = v14
    }
    var v15 int16 = (v13 % ((int16(v3)) | 1))
    if false {
        v3 = genFn0(-550)
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 int8 = 42
            delete(v5, 42)
            _ = fi16
            _ = v17
        }
        _ = len(v6)
    } else {
        if false {
            v3 += v3
            v3 -= (v3 * (v3 << ((18446744073709551615) & 63)))
            fmt.Println(v1)
        }
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        var v18 uint32 = 2
        _ = v18
    }
    var v19 uint32 = 0
    var v20 uint8 = 6
    for fi21 := 0; fi21 < 3; fi21++ {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v15, v13)
        case "":
            fmt.Println(v20, v1)
        case "a":
            var v22 int32 = 0
            _ = v22
        }
        v23 := v5[v19]
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 uint8 = 0
            _ = fi24
            _ = v25
        }
        _ = fi21
        _ = v23
    }
    fmt.Println(v1)
    var v26 int64 = -1
    v27 := v5[(v19 * (v19 | 65535))]
    switch (("foo bar" + ("foo bar" + "foo bar")) + ("" + (("\n\t" + "hello") + ("a" + "")))) {
    case "foo bar":
        v28 := GenS1{F0: v13, F1: 42, F2: 2147483647}
        _ = v28
    default:
        for fi29 := 0; fi29 < 6; fi29++ {
            _ = len(v2)
            switch "" {
            case "foo bar":
                var v30 float64 = ((((((1e10 + 3.14159) / 1.5) + (-2.25 * float64(1e10))) / 1.0) * 1.5) / -2.25)
                _ = v30
            case "hello":
                if v27 {
                    v31 := GenS1{F0: (v13 / ((v15) | 1)), F1: 948, F2: v1}
                    _ = v31
                } else {
                    var v32 string = ("hello" + (("foo bar" + "\n\t") + (((("foo bar" + "") + ("hello" + ("\n\t" + "hello"))) + "a") + ("foo bar" + ((("a" + "hello") + "a") + "\n\t")))))
                    if v27 {
                        var v33 uint64 = genFn0(v26)
                        _ = v33
                    } else {
                        switch v32 {
                        case "\n\t":
                            if true {
                                var v34 float32 = (3.14159 + 1.5)
                                _ = len(v2)
                                _ = v34
                            }
                        case "foo bar":
                            fmt.Println(v1, v27)
                        case "":
                            var v35 int16 = 42
                            _ = v35
                        default:
                            v36 := GenS1{F0: 943, F1: 643, F2: (v1 / ((876) | 1))}
                            _ = v36
                        }
                        v27 = v27
                    }
                    _ = v32
                }
            case "\n\t":
                var v37 uint64 = (v3 % ((1) | 1))
                _ = v37
            }
            _ = fi29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
