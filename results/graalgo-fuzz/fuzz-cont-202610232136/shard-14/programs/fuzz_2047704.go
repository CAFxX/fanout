package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16; F1 int16 }
func genFn0(p0 bool, p1 int) int8 {
    p0 = (100 >= -1)
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: -1}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0, p0)
    switch (("" + ("" + "foo bar")) + (("a" + ("\n\t" + "\n\t")) + ("a" + ("hello" + ((("\n\t" + ("foo bar" + ("hello" + ""))) + ("hello" + (("a" + "\n\t") + "\n\t"))) + ((((("" + ("hello" + "")) + ("hello" + "hello")) + "") + "hello") + "a")))))) {
    case "\n\t":
        var v8 float32 = ((1e10 - (-2.25 + (-2.25 - (-2.25 * (1e10 * (0.0 / 1.0)))))) - 1.5)
        _ = v8
    case "":
        var v9 int32 = 42
        _ = v9
    case "foo bar":
        fmt.Println(p0, p0)
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int32) int64 {
    for fi10 := 0; fi10 < 10; fi10++ {
        p0 = (^(p0))
        v11 := []int{4294967295}
        _ = fi10
        _ = v11
    }
    p0 -= p0
    p0 = p0
    _ = p0
    return 9223372036854775807
}

func main() {
    var v1 int8 = -74
    var v2 float32 = 3.14159
    var v3 uint32 = 2
    var v4 int = -9223372036854775808
    var v5 uint = 42
    if false {
        for fi12 := 0; fi12 < 9; fi12++ {
            if (169 != 1) {
                v13 := GenS2{F0: 9223372036854775807, F1: 65535, F2: 100}
                var v14 int32 = 2147483647
                _ = v13
                _ = v14
            }
            v4 = fi12
            v1 = (v1 - v1)
            _ = fi12
        }
        var v15 uint8 = 42
        var v16 uint = v5
        _ = v15
        _ = v16
    } else {
        switch v4 {
        case -820:
            if false {
                var v17 int64 = 5
                _ = v17
            }
        default:
            v4 = -9223372036854775808
        }
        fmt.Println(v2)
        var v18 int16 = 481
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 uint16 = 65535
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        v1 += int8(v5)
        _ = fi21
    }
    var v22 int16 = -205
    for fi23 := 0; fi23 < 18; fi23++ {
        switch "foo bar" {
        case "a":
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int64 = 42
                _ = fi24
                _ = v25
            }
        case "hello":
            fmt.Println(v22, v2)
        case "\n\t":
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 uint8 = 0
                _ = fi26
                _ = v27
            }
        }
        var v28 int32 = 1
        _ = fi23
        _ = v28
    }
    v29 := GenS3{F0: (v5 - v5)}
    v2 -= float32(0.0)
    v30 := []int{(^(v4)), 128, v4, v4}
    for fi31 := 0; fi31 < 17; fi31++ {
        for fi32 := 0; fi32 < 14; fi32++ {
            fmt.Println(v2, fi32, v4)
            v33 := GenS4{F0: (v3 % ((uint32(v5)) | 1)), F1: 419}
            _ = fi32
            _ = v33
        }
        v34 := GenS5{F0: 32767}
        _ = fi31
        _ = v34
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        fmt.Println(v1, v22, v5)
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        _ = fi35
    }
    v36 := GenS2{F0: 9223372036854775807, F1: uint16(v5), F2: v1}
    v2 = v2
    var v37 int16 = -1
    v4 = v4
    v2 -= float32(1.5)
    if (false && (2147483648 >= (v3 * (v3 >> (((v5 << ((v5) & 63))) & 63))))) {
        var v38 float64 = 3.14159
        _ = v38
    } else {
        _ = len(v30)
        var v39 float64 = (0.0 - 1e10)
        v40 := GenS6{F0: 507, F1: -32768}
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
