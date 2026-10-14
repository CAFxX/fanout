package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint }
func genFn0() int8 {
    if ("" < "\n\t") {
        var v4 int8 = 59
        var v5 int32 = 302
        _ = v4
        _ = v5
    } else {
        var v6 int32 = -2147483648
        v7 := []int{-9223372036854775808}
        _ = v6
        _ = v7
    }
    v8 := []int{1, -526}
    var v9 float32 = 0.0
    v10 := GenS1{F0: 1}
    _ = v10
    _ = v8
    _ = v9
    return 100
}
func genFn1(p0 int) int32 {
    fmt.Println(p0)
    var v11 bool = true
    _ = p0
    _ = v11
    return 710
}

func main() {
    var v1 int64 = -529
    var v2 bool = (false && (42 < -16))
    var v3 int16 = 0
    var v12 float32 = (-2.25 * 1.5)
    v3 = -349
    v2 = (true || false)
    v13 := GenS1{F0: 0}
    v14 := GenS2{F0: 21, F1: 18446744073709551615}
    switch "hello" {
    case "a":
        v15 := []int{1, -2147483648}
        _ = v15
    case "":
        fmt.Println(v2)
    case "\n\t":
        v16 := GenS1{F0: 237}
        _ = v16
    }
    if ((4294967295 >= 255) || (true || v2)) {
        fmt.Println(v2, v2, v3)
        fmt.Println(v2, v1)
        switch "foo bar" {
        case "foo bar":
            var v17 string = "hello"
            _ = v17
        case "a":
            v2 = (!v2)
        default:
            fmt.Println(v12, v2, v1)
        }
    } else {
        v18 := GenS2{F0: 42, F1: 690}
        var v19 uint64 = 0
        _ = v18
        _ = v19
    }
    var v20 uint8 = 122
    for fi21 := 0; fi21 < 15; fi21++ {
        v3 += v3
        _ = fi21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int = 1
        if (genFn1((fi22 / ((823) | 1))) > -2147483648) {
            for fi24 := 0; fi24 < 4; fi24++ {
                v25 := []int{(fi24 << ((65535) & 63))}
                var v26 int8 = -128
                v27 := GenS3{F0: (v1 / ((v1) | 1)), F1: (v26 ^ (v26 >> ((286) & 63)))}
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            v2 = (0 >= 0)
        } else {
            v1 += 9223372036854775807
            var v28 uint = 65535
            _ = v28
        }
        v23 = (^(fi22))
        _ = fi22
        _ = v23
    }
    var v29 float64 = 0.0
    if true {
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 uint = 42
            _ = fi30
            _ = v31
        }
    } else {
        v32 := []int{68, 128}
        v1 -= 9223372036854775807
        var v33 string = ("foo bar" + ("a" + ""))
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        fmt.Println(v29)
        v35 := GenS4{F0: 0, F1: 2147483647, F2: 295}
        v36 := []int{(fi34 - fi34), (-(fi34))}
        _ = fi34
        _ = v35
        _ = v36
    }
    v1 = 42
    v1 -= (v1 + (v1 % (((v1 << ((65535) & 63))) | 1)))
    v37 := GenS3{F0: (^(v1)), F1: 100}
    if (65535 > 32767) {
        v1 -= v1
        switch (((("foo bar" + ("" + ("" + ""))) + ("a" + "foo bar")) + ("\n\t" + "foo bar")) + "foo bar") {
        case "\n\t":
            for fi38 := 0; fi38 < 2; fi38++ {
                fmt.Println(v3)
                _ = fi38
            }
        default:
            var v39 int64 = v1
            _ = v39
        }
    }
    var v40 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v37
    _ = v40
}
