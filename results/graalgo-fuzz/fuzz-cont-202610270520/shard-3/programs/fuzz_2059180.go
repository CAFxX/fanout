package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int8; F1 int8; F2 string }
func genFn0(p0 uint64, p1 uint8, p2 int64) int {
    switch ("a" + "hello") {
    case "":
        p1 = 128
    case "\n\t":
        if (true && ((101 >= 539) && false)) {
            fmt.Println(p2, p1)
            var v4 int = 9223372036854775807
            v5 := []int{(v4 * 127), v4, (v4 << ((626) & 63)), 4294967295}
            _ = v4
            _ = v5
        }
    case "hello":
        if (1 < 0) {
            var v6 int32 = -1
            var v7 int16 = 42
            v8 := GenS1{F0: 2, F1: 1}
            _ = v6
            _ = v7
            _ = v8
        }
    default:
        switch 0 {
        case -2147483648:
            var v9 float64 = 0.0
            _ = v9
        case 446:
            switch 9223372036854775807 {
            case -9223372036854775808:
                p2 = (-(p2))
            default:
                if (p1 < (^(p1))) {
                    p1 += (p1 - 2)
                }
            }
        case 0:
            var v10 int8 = 0
            _ = v10
        default:
            switch "" {
            case "foo bar":
                p2 = int64(p1)
            default:
                var v11 int16 = 0
                _ = v11
            }
        }
    }
    if true {
        var v12 uint64 = p0
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 531
}

func main() {
    var v1 int = -924
    var v2 int16 = -684
    v3 := make(map[string]int)
    var v13 bool = true
    v3["a"] = 2147483647
    switch "a" {
    case "":
        if v13 {
            if (18446744073709551615 < 760) {
                var v14 uint8 = 2
                _ = v14
            }
        }
    case "foo bar":
        v15 := v3["\n\t"]
        _ = v15
    default:
        v2 = v2
    }
    fmt.Println(v2, v13, v13)
    v3[("hello" + "\n\t")] = (v1 & genFn0(42, 0, -354))
    v2 = 42
    v16 := []int{(v1 & genFn0(442, 0, 42)), (v1 ^ v1)}
    if (255 < 0) {
        switch v1 {
        case 42:
            var v17 string = "\n\t"
            _ = v17
        default:
            if true {
                v18 := v3["\n\t"]
                _ = v18
            }
        }
        v19 := GenS1{F0: 108, F1: 2}
        v20 := GenS2{F0: 4294967295, F1: 2, F2: 2147483647}
        _ = v19
        _ = v20
    }
    var v21 uint = 42
    v22 := GenS1{F0: 255, F1: 439}
    if v13 {
        var v23 string = ""
        for fi24 := 0; fi24 < 8; fi24++ {
            switch genFn0(430, 251, 0) {
            case -508:
                v25 := GenS3{F0: 1, F1: -60, F2: v23}
                _ = v25
            case 4294967295:
                var v26 int8 = 42
                _ = v26
            default:
                var v27 string = v23
                _ = v27
            }
            var v28 int8 = -113
            _ = fi24
            _ = v28
        }
        v16 = append(v16, (v1 | genFn0(840, 128, 1)))
        _ = v23
    }
    _ = len(v16)
    fmt.Println(v13)
    if v13 {
        for fi29 := 0; fi29 < 10; fi29++ {
            v3[("" + ("a" + ""))] = (^(v1))
            var v30 uint32 = 2
            if len(v16) > 0 {
                v16[0] = genFn0(0, 255, -1)
            }
            _ = fi29
            _ = v30
        }
        var v31 int32 = -1
        _ = v31
    } else {
        var v32 string = ""
        v21 = v21
        v32 = ""
        _ = v32
    }
    if false {
        var v33 uint32 = 641
        var v34 string = "hello"
        fmt.Println(v34, v2)
        _ = v33
        _ = v34
    } else {
        v21 += (v21 / (((v21 + 2)) | 1))
    }
    v13 = (!(882 == -2147483648))
    if (((("" + ("a" + "hello")) + "foo bar") == "hello") && false) {
        delete(v3, (((("\n\t" + "\n\t") + ("foo bar" + "foo bar")) + "") + "hello"))
        var v35 float32 = 0.0
        v13 = v13
        _ = v35
    } else {
        var v36 string = (("hello" + "hello") + ((("foo bar" + "a") + ("\n\t" + (("foo bar" + "a") + "foo bar"))) + "a"))
        _ = v36
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        v2 -= (v2 + -1)
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
