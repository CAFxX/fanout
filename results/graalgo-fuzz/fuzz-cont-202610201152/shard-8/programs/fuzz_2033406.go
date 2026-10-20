package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int64; F1 uint32; F2 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 int8; F2 int }
func genFn0(p0 string) int16 {
    v5 := GenS1{F0: 21, F1: 127, F2: p0}
    switch (("a" + ((p0 + "") + (p0 + "hello"))) + p0) {
    case "hello":
        if (2 <= 18446744073709551615) {
            if true {
                if (true || (-104 < 100)) {
                    v6 := []int{-1}
                    for fi7 := 0; fi7 < 14; fi7++ {
                        p0 = p0
                        var v8 float64 = (1.5 + (3.14159 + (0.1 * 0.0)))
                        _ = fi7
                        _ = v8
                    }
                    var v9 int32 = -2147483648
                    _ = v6
                    _ = v9
                } else {
                    var v10 uint64 = 0
                    v11 := GenS1{F0: -861, F1: 0, F2: (p0 + "")}
                    var v12 uint = 233
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
            switch p0 {
            case "foo bar":
                v13 := []int{-9223372036854775808, 32768, 9223372036854775807}
                _ = v13
            case "":
                v14 := GenS2{F0: 65535}
                _ = v14
            }
        }
    }
    v15 := GenS3{F0: 1, F1: 4294967295}
    switch "hello" {
    case "":
        v16 := []int{-345}
        _ = v16
    }
    _ = p0
    _ = v15
    _ = v5
    return -65
}
func genFn1(p0 uint8, p1 uint64) uint8 {
    p1 -= (^(p1))
    fmt.Println(p0, p1)
    v17 := GenS4{F0: -790, F1: 2147483647, F2: (("foo bar" + "\n\t") < "a")}
    var v18 int16 = genFn0("\n\t")
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return p0
}

func main() {
    var v1 float32 = -2.25
    var v2 int32 = 42
    var v3 uint8 = 127
    var v4 uint32 = 2
    v19 := GenS5{F0: ""}
    var v20 uint = uint(v3)
    var v21 uint = (v20 - v20)
    switch -9223372036854775808 {
    case -508:
        v22 := []int{9223372036854775807, 1}
        _ = v22
    case 0:
        var v23 int32 = (v2 * (v2 * v2))
        _ = v23
    case 809:
        if false {
            fmt.Println(v4)
        }
    default:
        var v24 uint16 = 497
        _ = v24
    }
    v25 := []int{}
    if ((127 < 127) && (851 != -9223372036854775808)) {
        fmt.Println(v2, v2, v2)
        fmt.Println(v20, v20)
        for fi26 := 0; fi26 < 15; fi26++ {
            _ = len(v25)
            var v27 int8 = 127
            _ = fi26
            _ = v27
        }
    }
    switch "foo bar" {
    case "\n\t":
        for fi28 := 0; fi28 < 6; fi28++ {
            v21 += v21
            var v29 float32 = -2.25
            _ = fi28
            _ = v29
        }
    case "a":
        if len(v25) > 0 {
            v25[0] = 4294967295
        }
    case "hello":
        if false {
            fmt.Println(v4)
        } else {
            if len(v25) > 0 {
                v25[0] = 42
            }
        }
    }
    v20 = uint(v3)
    for fi30 := 0; fi30 < 5; fi30++ {
        v21 -= v20
        v31 := GenS1{F0: 1, F1: 42, F2: "foo bar"}
        _ = fi30
        _ = v31
    }
    v32 := GenS3{F0: 341, F1: v20}
    if true {
        v3 += (^(v3))
    }
    switch 32768 {
    case -1:
        _ = len(v25)
    case 2147483647:
        v33 := GenS4{F0: -329, F1: uint32(v3), F2: (!false)}
        _ = v33
    case -663:
        fmt.Println(v21)
    default:
        var v34 uint8 = 1
        _ = v34
    }
    switch 2147483647 {
    case -241:
        fmt.Println(v20)
    case -9223372036854775808:
        v35 := GenS6{F0: v3, F1: 1, F2: 734}
        _ = v35
    case 127:
        for fi36 := 0; fi36 < 16; fi36++ {
            v4 += (v4 << ((v21) & 63))
            _ = fi36
        }
    default:
        var v37 float32 = v1
        _ = v37
    }
    v1 = (float32(((-2.25 - float64(0.0)) * 0.1)) - ((-2.25 * v1) - ((((v1 + v1) * float32(0.1)) / v1) - 0.1)))
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 int8 = 0
        switch fi38 {
        case 269:
            var v40 uint64 = 188
            _ = v40
        case -2147483648:
            v39 -= v39
        case -536:
            for fi41 := 0; fi41 < 19; fi41++ {
                v4 = (v4 ^ 2)
                _ = len(v25)
                _ = fi41
            }
        }
        _ = fi38
        _ = v39
    }
    v42 := GenS6{F0: 97, F1: 108, F2: 127}
    _ = len(v25)
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v42
}
