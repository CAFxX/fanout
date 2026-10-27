package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
func genFn0(p0 uint16, p1 uint32) string {
    var v4 bool = ((0 > 278) || true)
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: 271, F1: 877}
        switch "a" {
        case "a":
            var v7 int16 = 0
            _ = v7
        }
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            v8 := []int{32768, fi5, (fi5 >> ((1) & 63))}
            _ = v8
        case "hello":
            var v9 bool = (p0 >= p0)
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    p0 += (p0 & (p0 << ((18446744073709551615) & 63)))
    v10 := []int{388}
    v4 = (14 >= -614)
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    return (("" + "hello") + ("a" + "a"))
}
func genFn1() uint16 {
    var v11 string = ("" + "hello")
    v11 = "\n\t"
    var v12 string = "hello"
    v11 = v11
    v13 := []int{128, -259, -943, -9223372036854775808}
    _ = v11
    _ = v12
    _ = v13
    return 32767
}

func main() {
    var v1 int = 32768
    var v2 int64 = 42
    var v3 int8 = -128
    v14 := []int{v1, v1, (v1 >> ((2) & 63)), -596}
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS1{F0: 1, F1: 255}
            if len(v14) > 0 {
                v14[0] = (v1 ^ fi15)
            }
            _ = fi16
            _ = v17
        }
        v2 = (-(v2))
        for fi18 := 0; fi18 < 10; fi18++ {
            if len(v14) > 0 {
                v14[0] = 32768
            }
            var v19 bool = (v2 <= (v2 & (v2 ^ v2)))
            switch 118 {
            case -317:
                var v20 uint = 467
                _ = v20
            default:
                var v21 int64 = (v2 % ((int64(v3)) | 1))
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        _ = fi15
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS1{F0: 1, F1: 697}
        _ = fi22
        _ = v23
    }
    v1 = 4294967295
    v24 := GenS1{F0: 18446744073709551615, F1: 4294967295}
    var v25 int = -2147483648
    var v26 float64 = 3.14159
    v25 = 4294967295
    switch (("hello" + "a") + ("a" + "a")) {
    case "foo bar":
        _ = len(v14)
    default:
        var v27 int64 = 9223372036854775807
        _ = v27
    }
    switch "foo bar" {
    case "hello":
        v2 += int64(v3)
    case "\n\t":
        if (0 >= 0) {
            for fi28 := 0; fi28 < 17; fi28++ {
                v29 := GenS1{F0: 2, F1: 42}
                _ = fi28
                _ = v29
            }
            fmt.Println(v3, v3, v26)
        }
    }
    if true {
        v26 -= v26
    } else {
        fmt.Println(v26, v1)
        if len(v14) > 0 {
            v14[0] = (v1 ^ (v25 - v25))
        }
    }
    if true {
        var v30 uint = 131
        switch "hello" {
        case "hello":
            v3 += (v3 ^ v3)
        case "\n\t":
            switch 1 {
            case -9223372036854775808:
                v3 = v3
            case -305:
                fmt.Println(v2, v25)
            case 128:
                v31 := GenS1{F0: 42, F1: v30}
                _ = v31
            default:
                var v32 bool = false
                _ = v32
            }
        }
        if true {
            v25 = (v25 * 310)
            v30 = (v30 + (v30 % (((v30 | v30)) | 1)))
            v2 += (v2 * (-(v2)))
        }
        _ = v30
    } else {
        var v33 uint64 = 18446744073709551615
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 bool = false
            v14 = append(v14, (v1 << ((0) & 63)))
            if len(v14) > 0 {
                v14[0] = fi34
            }
            _ = fi34
            _ = v35
        }
        _ = v33
    }
    if (!false) {
        var v36 uint16 = 486
        fmt.Println(v25, v2)
        _ = len(v14)
        _ = v36
    }
    var v37 int64 = v2
    var v38 int = (-(v1))
    if ((42 >= 1) || (1e10 < 1.5)) {
        v37 -= v2
    }
    switch (("\n\t" + (("" + "") + "")) + "hello") {
    case "foo bar":
        var v39 uint8 = 31
        _ = v39
    case "hello":
        switch -9223372036854775808 {
        case -734:
            var v40 bool = (987 <= 580)
            _ = v40
        default:
            var v41 int = 4294967295
            _ = v41
        }
    }
    v42 := GenS1{F0: 880, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v42
}
