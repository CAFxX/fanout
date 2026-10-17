package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 string }
func genFn0(p0 uint16, p1 int) int64 {
    fmt.Println(p0)
    switch p1 {
    case -9223372036854775808:
        p0 = 42
    default:
        var v6 int = 272
        _ = v6
    }
    p1 += p1
    if true {
        var v7 uint64 = 42
        if (892 != p0) {
            fmt.Println(p0)
            fmt.Println(v7, p1, v7)
            v8 := GenS1{F0: 16, F1: "hello"}
            _ = v8
        }
        _ = v7
    }
    var v9 int64 = int64(p1)
    _ = p0
    _ = p1
    _ = v9
    return v9
}
func genFn1(p0 uint) uint {
    v10 := []int{-2147483648, 42}
    v10 = append(v10, 42)
    var v11 float32 = 3.14159
    var v12 int32 = 1
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return (p0 * (p0 * (p0 - (p0 % (((^(p0))) | 1)))))
}

func main() {
    var v1 uint = 186
    var v2 float64 = float64(1e10)
    var v3 uint64 = 4294967295
    var v4 bool = false
    v5 := make(map[string]string)
    var v13 uint32 = 2
    var v14 int16 = 32767
    if false {
        v15 := []int{}
        var v16 int64 = 490
        fmt.Println(v1, v14, v2)
        _ = v15
        _ = v16
    } else {
        var v17 string = (("" + ((((("hello" + "") + ((("hello" + "a") + ("" + "a")) + "hello")) + "a") + ("\n\t" + "a")) + "hello")) + "hello")
        _ = v17
    }
    if (!(-9223372036854775808 == 0)) {
        switch "" {
        case "foo bar":
            for fi18 := 0; fi18 < 5; fi18++ {
                v13 += 0
                _ = fi18
            }
        case "a":
            if v4 {
                v19 := []int{-2147483648, -281, 840}
                v1 = 255
                var v20 int32 = -192
                _ = v19
                _ = v20
            }
        case "":
            v2 += (v2 - v2)
        }
        v14 = -957
        v1 += 255
    } else {
        v21 := GenS1{F0: 55, F1: ""}
        _ = v21
    }
    fmt.Println(v13)
    switch "\n\t" {
    case "hello":
        v22 := []int{-13, 213}
        _ = v22
    case "\n\t":
        switch "a" {
        case "":
            v2 += (v2 - (1e10 * v2))
        }
    }
    v23 := GenS1{F0: 162, F1: "hello"}
    v1 += 864
    for fi24 := 0; fi24 < 1; fi24++ {
        switch fi24 {
        case -662:
            var v25 int32 = -58
            _ = v25
        default:
            v26 := GenS2{F0: ("hello" + (("" + "") + "a"))}
            _ = v26
        }
        _ = fi24
    }
    var v27 float32 = 0.1
    v28 := GenS2{F0: "foo bar"}
    switch (((("" + ("a" + ("foo bar" + "hello"))) + "") + "foo bar") + "") {
    case "foo bar":
        for fi29 := 0; fi29 < 19; fi29++ {
            v30 := GenS2{F0: ("a" + "\n\t")}
            _ = fi29
            _ = v30
        }
    case "a":
        var v31 int8 = 42
        _ = v31
    default:
        var v32 uint16 = 32767
        _ = v32
    }
    v13 -= v13
    v33 := GenS2{F0: "a"}
    v4 = v4
    v14 += 42
    for fi34 := 0; fi34 < 6; fi34++ {
        v35 := GenS2{F0: "\n\t"}
        _ = fi34
        _ = v35
    }
    v36 := GenS1{F0: 127, F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v5
}
