package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
func genFn0() uint {
    var v7 int64 = 356
    var v8 float32 = 3.14159
    v8 -= 1.5
    for fi9 := 0; fi9 < 6; fi9++ {
        if (255 == 213) {
            var v10 string = ((("" + "foo bar") + ("" + "hello")) + "foo bar")
            var v11 uint = 42
            fmt.Println(v10)
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    v12 := GenS1{F0: 18, F1: ""}
    _ = v12
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint8, p1 uint16, p2 int) uint16 {
    var v13 string = ("hello" + "foo bar")
    v13 = v13
    fmt.Println(v13, p1, p1)
    v14 := GenS1{F0: (p0 * (p0 * (^(p0)))), F1: v13}
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int16 = -32768
        v17 := []int{}
        var v18 uint8 = 127
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return p1
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]int)
    v3 := []string{"\n\t", "", ("\n\t" + ""), ""}
    var v4 uint8 = 152
    var v5 uint32 = 695
    var v6 uint = 276
    v4 -= v4
    if false {
        var v19 int = 127
        switch 1 {
        case -9223372036854775808:
            v6 = 2
        }
        _ = v19
    } else {
        fmt.Println(v5, v1, v1)
    }
    if true {
        if (-32768 < 963) {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 uint16 = 549
                _ = fi20
                _ = v21
            }
            v4 += v4
            v5 = v5
        }
        v22 := GenS1{F0: (v4 & (v4 / (((v4 - v4)) | 1))), F1: ("\n\t" + "foo bar")}
        _ = v22
    }
    v23 := GenS1{F0: 2, F1: "hello"}
    var v24 int = -530
    if true {
        switch "" {
        case "\n\t":
            _ = len(v2)
        case "a":
            fmt.Println(v6, v4)
        case "hello":
            var v25 uint8 = (v4 + (v4 << ((genFn0()) & 63)))
            _ = v25
        }
        fmt.Println(v5, v1)
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 uint64 = 0
            var v28 uint8 = v4
            _ = fi26
            _ = v27
            _ = v28
        }
    } else {
        for fi29 := 0; fi29 < 17; fi29++ {
            v1 -= v1
            _ = fi29
        }
    }
    if (!(true && false)) {
        var v30 string = ("\n\t" + "a")
        _ = v30
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        v4 -= 42
        _ = fi31
    }
    v5 = (v5 ^ v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
