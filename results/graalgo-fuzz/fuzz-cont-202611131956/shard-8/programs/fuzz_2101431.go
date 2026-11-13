package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int16; F2 int8 }
type GenS4 struct { F0 int16; F1 int32 }
func genFn0(p0 int16) uint8 {
    var v5 string = ""
    var v6 int32 = -220
    var v7 uint16 = 798
    v8 := []int{42}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 58
}
func genFn1(p0 string, p1 uint8) bool {
    v9 := GenS1{F0: 606, F1: -2147483648}
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int = fi10
        var v12 int32 = 1
        var v13 int8 = 28
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := []int{32767, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v14
    _ = v9
    return (!false)
}

func main() {
    var v1 int = 2147483647
    var v2 int32 = -2147483648
    v3 := []string{("" + (("hello" + "hello") + "\n\t")), "hello", ("a" + "foo bar")}
    var v4 int16 = -13
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v15 int16 = v4
    v16 := GenS2{F0: 2}
    switch v1 {
    case 2147483647:
        v17 := GenS2{F0: 652}
        _ = v17
    case 491:
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 string = ""
            v15 -= (v4 % (((v4 << ((18446744073709551615) & 63))) | 1))
            v19 = v19
            _ = fi18
            _ = v19
        }
    default:
        v20 := GenS1{F0: 2, F1: 2147483647}
        _ = v20
    }
    switch "\n\t" {
    case "hello":
        v4 += v15
    }
    v15 -= 1000
    fmt.Println(v1, v4)
    switch (("a" + "\n\t") + "hello") {
    case "foo bar":
        fmt.Println(v2, v15, v4)
    case "hello":
        for fi21 := 0; fi21 < 8; fi21++ {
            v2 = v2
            switch "hello" {
            case "hello":
                v22 := GenS3{F0: 42, F1: v15, F2: 100}
                _ = v22
            case "\n\t":
                var v23 string = "hello"
                _ = v23
            }
            _ = fi21
        }
    case "a":
        var v24 int = -655
        _ = v24
    }
    switch v1 {
    case -217:
        fmt.Println(v4, v4)
    case 32767:
        fmt.Println(v1, v2)
    case -831:
        var v25 int = (v1 + (^(v1)))
        _ = v25
    }
    var v26 int8 = -119
    if (false || true) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v15 += v4
    v27 := GenS4{F0: (^(v4)), F1: v2}
    switch (^(v1)) {
    case 32768:
        var v28 int64 = 42
        _ = v28
    case -1:
        for fi29 := 0; fi29 < 2; fi29++ {
            v2 = (v2 % (((v2 >> ((65535) & 63))) | 1))
            for fi30 := 0; fi30 < 13; fi30++ {
                fmt.Println(v1, v26, v15)
                var v31 int64 = 9223372036854775807
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
    case 127:
        var v32 uint32 = 4294967295
        _ = v32
    }
    fmt.Println(v26, v15, v4)
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 float32 = 0.1
        v15 = v15
        var v35 float64 = 3.14159
        _ = fi33
        _ = v34
        _ = v35
    }
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
