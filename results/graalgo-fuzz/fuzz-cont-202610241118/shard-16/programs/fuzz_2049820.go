package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint }
type GenS2 struct { F0 string }
func genFn0(p0 uint32, p1 int32) int32 {
    p0 -= (p0 ^ 1)
    var v5 uint = 1
    _ = p0
    _ = p1
    _ = v5
    return (p1 & p1)
}

func main() {
    var v1 uint8 = 42
    var v2 uint8 = 127
    var v3 int = -1
    v4 := make(map[uint32]int)
    var v6 string = "\n\t"
    fmt.Println(v2, v3)
    v7 := []int{}
    v1 += (^(v1))
    fmt.Println(v1, v2)
    if (v3 != (v3 - (v3 - v3))) {
        if false {
            var v8 uint64 = 4294967295
            _ = v8
        }
    }
    switch v6 {
    case "\n\t":
        var v9 float32 = -2.25
        _ = v9
    default:
        v3 -= (v3 * v3)
    }
    switch ("hello" + "") {
    case "\n\t":
        if (825 > int16(v3)) {
            v2 = v1
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 uint = 0
                v3 = fi10
                fmt.Println(fi10, v11, v1)
                _ = fi10
                _ = v11
            }
            v7 = append(v7, (^(v3)))
        }
    case "hello":
        v12 := GenS1{F0: -945, F1: -469, F2: 0}
        _ = v12
    case "a":
        if false {
            fmt.Println(v3, v1, v6)
            if true {
                var v13 int16 = 42
                _ = v13
            }
            var v14 uint32 = 369
            _ = v14
        } else {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
    }
    if (v6 < "foo bar") {
        v1 = v2
    } else {
        switch "\n\t" {
        case "":
            var v15 string = "hello"
            _ = v15
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            v3 = (^(v3))
            v17 := GenS1{F0: 1, F1: 1000, F2: 65535}
            _ = fi16
            _ = v17
        }
    }
    switch v6 {
    case "":
        switch "\n\t" {
        case "foo bar":
            if true {
                v6 = ("foo bar" + (("foo bar" + v6) + (v6 + ("" + "a"))))
                var v18 int16 = -1
                if false {
                    v7 = append(v7, v3)
                    var v19 int8 = 127
                    v3 += v3
                    _ = v19
                }
                _ = v18
            } else {
                fmt.Println(v6, v6)
            }
        case "\n\t":
            if (-1 > 0) {
                v20 := GenS1{F0: 1, F1: 231, F2: 0}
                switch v3 {
                case 127:
                    for fi21 := 0; fi21 < 19; fi21++ {
                        var v22 uint8 = (v2 / ((0) | 1))
                        _ = fi21
                        _ = v22
                    }
                }
                _ = v20
            } else {
                var v23 int8 = 16
                v24 := GenS2{F0: ""}
                _ = v23
                _ = v24
            }
        case "hello":
            var v25 float32 = -2.25
            _ = v25
        default:
            v26 := GenS1{F0: 9223372036854775807, F1: 32767, F2: 18446744073709551615}
            _ = v26
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v27 uint16 = 468
    v28 := GenS1{F0: -721, F1: -1, F2: 516}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
