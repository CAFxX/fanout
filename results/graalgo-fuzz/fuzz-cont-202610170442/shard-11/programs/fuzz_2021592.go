package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 string; F1 int64; F2 int32 }
func genFn0(p0 int64, p1 int16, p2 uint8) int16 {
    v6 := GenS1{F0: -653}
    var v7 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}
func genFn1(p0 uint32) int16 {
    var v8 uint64 = 74
    var v9 uint8 = 42
    v10 := GenS2{F0: "\n\t", F1: v9}
    var v11 int8 = -51
    var v12 uint16 = 233
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint16 = 69
    var v3 uint32 = 42
    var v4 int8 = 101
    var v5 uint32 = 2
    switch -1 {
    case -403:
        v13 := GenS2{F0: "\n\t", F1: 127}
        _ = v13
    default:
        v3 -= (v3 % ((v5) | 1))
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := []int{}
        var v16 int64 = 1
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v3, v2)
    v17 := GenS1{F0: -437}
    v18 := GenS2{F0: (("hello" + "a") + ("a" + "hello")), F1: 76}
    if ((!(!(((v1 < v1) && false) || true))) && ((v1 - 1) > v1)) {
        v19 := []int{127, 127}
        _ = v19
    }
    switch ("a" + "a") {
    case "\n\t":
        var v20 int = 943
        _ = v20
    case "hello":
        for fi21 := 0; fi21 < 11; fi21++ {
            if (float64(1e10) == 0.1) {
                v22 := GenS3{F0: ("" + "a"), F1: -967, F2: 238}
                v1 = v1
                _ = v22
            }
            fmt.Println(fi21, v5)
            _ = fi21
        }
    case "":
        var v23 uint = v1
        _ = v23
    default:
        v2 += 2
    }
    switch ((("hello" + "foo bar") + "a") + "\n\t") {
    case "":
        var v24 string = "hello"
        _ = v24
    case "foo bar":
        for fi25 := 0; fi25 < 8; fi25++ {
            v26 := GenS1{F0: -524}
            var v27 int = (fi25 / ((128) | 1))
            _ = fi25
            _ = v26
            _ = v27
        }
    case "\n\t":
        fmt.Println(v1, v5, v1)
    }
    if true {
        var v28 string = ""
        _ = v28
    } else {
        v4 = 105
    }
    var v29 uint8 = 1
    switch "" {
    case "a":
        if ((!(1 != (v4 % ((1) | 1)))) && (42 < -1)) {
            v30 := []int{}
            for fi31 := 0; fi31 < 3; fi31++ {
                v32 := GenS2{F0: (("a" + "\n\t") + (("\n\t" + ("hello" + "a")) + ("foo bar" + "a"))), F1: v29}
                _ = fi31
                _ = v32
            }
            _ = v30
        }
    default:
        switch "\n\t" {
        case "foo bar":
            v29 -= v29
        case "hello":
            for fi33 := 0; fi33 < 3; fi33++ {
                var v34 int64 = -1
                _ = fi33
                _ = v34
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
