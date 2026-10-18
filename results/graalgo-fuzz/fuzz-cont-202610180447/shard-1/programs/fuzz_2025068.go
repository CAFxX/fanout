package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int; F1 int8; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint64) uint8 {
    var v7 int = 42
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float32 = (((3.14159 + 3.14159) / 1.0) * -2.25)
        v10 := []int{-912}
        var v11 float32 = (v9 / (((-2.25 / v9) + -2.25) / (v9 - v9)))
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 int64, p2 int32) int64 {
    switch 128 {
    case 2147483647:
        fmt.Println(p2, p2)
    case -1:
        var v12 float64 = (0.0 / 0.1)
        _ = v12
    }
    if false {
        v13 := GenS1{F0: p1, F1: -94}
        var v14 uint16 = 0
        v14 = (^(v14))
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 128, F1: 24, F2: 1}
    var v16 int8 = 0
    switch 4294967295 {
    case -2147483648:
        switch "" {
        case "\n\t":
            var v17 uint8 = 42
            _ = v17
        case "foo bar":
            switch 32768 {
            case 2147483647:
                var v18 int64 = -917
                _ = v18
            case 635:
                var v19 uint16 = 255
                _ = v19
            }
        default:
            p0 = 18446744073709551615
        }
    default:
        if (2 == uint64(p0)) {
            p0 = p0
            fmt.Println(p2)
        } else {
            var v20 int16 = 161
            _ = v20
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return (^(p1))
}

func main() {
    var v1 float32 = 1.5
    var v2 float32 = (v1 - v1)
    v3 := make(map[uint32]int)
    var v4 uint16 = 255
    var v5 uint8 = 42
    var v6 uint = 2
    if ((1e10 + -2.25) <= 0.1) {
        var v21 uint = (v6 | (v6 & 45))
        switch (((("" + "a") + ("\n\t" + "")) + ("" + ("a" + "hello"))) + ("hello" + ("a" + ""))) {
        case "hello":
            v22 := []int{-9223372036854775808, 42, -312}
            _ = v22
        case "":
            v4 = 2
        case "a":
            v2 -= v1
        }
        _ = v21
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 uint64 = 18446744073709551615
        for fi25 := 0; fi25 < 3; fi25++ {
            switch int(v6) {
            case 210:
                var v26 int = fi25
                _ = v26
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    v27 := v3[2147483648]
    fmt.Println(v6, v2)
    for fi28 := 0; fi28 < 15; fi28++ {
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 string = ("hello" + ("hello" + ((("hello" + ("foo bar" + ("\n\t" + "a"))) + (("foo bar" + ("foo bar" + ("foo bar" + (("a" + ("\n\t" + "a")) + ("" + ("a" + (("foo bar" + "hello") + "a"))))))) + "hello")) + "hello")))
            v6 += 1
            var v31 uint64 = 727
            var v32 int16 = 0
            v6 -= 807
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        if false {
            v33 := GenS3{F0: "\n\t"}
            _ = v33
        }
        _ = fi28
    }
    v34 := GenS1{F0: 42, F1: 0}
    v35 := GenS4{F0: 2147483647}
    delete(v3, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
