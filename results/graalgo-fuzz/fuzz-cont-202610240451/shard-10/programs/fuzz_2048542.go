package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 int64; F1 int8; F2 int32 }
type GenS3 struct { F0 uint; F1 int64; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 string; F1 int64 }
func genFn0() bool {
    v6 := GenS1{F0: 542, F1: 61, F2: "foo bar"}
    switch 2147483647 {
    case 127:
        var v7 uint8 = 252
        _ = v7
    case 1:
        var v8 int = -836
        _ = v8
    default:
        var v9 int8 = -85
        _ = v9
    }
    var v10 int32 = 1
    if false {
        v11 := GenS2{F0: 709, F1: 127, F2: 103}
        _ = v11
    } else {
        v12 := GenS2{F0: -599, F1: 100, F2: v10}
        _ = v12
    }
    v13 := []int{-816, -806}
    _ = v10
    _ = v13
    _ = v6
    return ((!true) && false)
}
func genFn1() uint {
    v14 := []int{32767, 127}
    switch (("hello" + "a") + "a") {
    case "":
        v15 := GenS3{F0: 4294967295, F1: -9223372036854775808, F2: 0}
        _ = v15
    case "hello":
        var v16 int64 = -9223372036854775808
        _ = v16
    case "a":
        switch "" {
        case "hello":
            switch "" {
            case "":
                if (false || (true || true)) {
                    var v17 int64 = 1
                    _ = v17
                } else {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    var v18 float32 = (0.1 * 1e10)
                    v19 := GenS4{F0: 0}
                    _ = v18
                    _ = v19
                }
            default:
                var v20 uint32 = 1
                _ = v20
            }
        }
    default:
        var v21 float32 = -2.25
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS1{F0: 680, F1: 42, F2: "\n\t"}
        _ = fi22
        _ = v23
    }
    _ = v14
    return 0
}

func main() {
    v1 := []string{"a", (("foo bar" + ("\n\t" + ("a" + ""))) + "foo bar"), "a", ("" + ("foo bar" + ""))}
    v2 := make(map[int]uint32)
    var v3 int8 = 1
    var v4 int = 4294967295
    var v5 bool = (!(-28 != -32768))
    for fi24 := 0; fi24 < 10; fi24++ {
        fmt.Println(v5)
        if v5 {
            v3 = (v3 - v3)
            v25 := GenS5{F0: v3, F1: 0}
            for fi26 := 0; fi26 < 19; fi26++ {
                fmt.Println(v5)
                fmt.Println(fi24, fi26)
                _ = fi26
            }
            _ = v25
        } else {
            if (!false) {
                v4 -= 4294967295
            }
            var v27 int8 = 0
            _ = v27
        }
        var v28 int = (fi24 * fi24)
        _ = fi24
        _ = v28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 uint = 65535
        v4 = fi29
        for fi31 := 0; fi31 < 14; fi31++ {
            fmt.Println(v30)
            v32 := GenS6{F0: 0}
            v3 += 1
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    if (-2.25 <= float32(0.0)) {
        for fi33 := 0; fi33 < 7; fi33++ {
            var v34 int = -797
            _ = fi33
            _ = v34
        }
    } else {
        v4 += (v4 + v4)
    }
    fmt.Println(v4)
    v2[(-(v4))] = 0
    fmt.Println(v4)
    switch "" {
    case "foo bar":
        v35 := GenS7{F0: 1, F1: 2147483647}
        _ = v35
    case "\n\t":
        switch v4 {
        case 127:
            v36 := GenS5{F0: v3, F1: 750}
            _ = v36
        case 32768:
            v1 = append(v1, "")
        }
    case "a":
        var v37 string = "hello"
        _ = v37
    default:
        v38 := GenS8{F0: "foo bar", F1: 1}
        _ = v38
    }
    v5 = v5
    if v5 {
        var v39 int = (v4 / ((v4) | 1))
        v40 := GenS6{F0: 942}
        _ = v39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
