package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 int32 }
type GenS4 struct { F0 int64; F1 int32; F2 string }
func genFn0(p0 uint32) bool {
    p0 -= 2147483648
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint16 = 1
        _ = fi5
        _ = v6
    }
    p0 = 2
    _ = p0
    return false
}
func genFn1(p0 string, p1 int64, p2 int64) int64 {
    var v7 string = p0
    if true {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 & p1)
}

func main() {
    v1 := []int{}
    var v2 int8 = 42
    var v3 uint16 = 65535
    v4 := make(map[string]int)
    switch "\n\t" {
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "\n\t":
        var v8 uint32 = 2147483647
        _ = v8
    case "hello":
        if (true || true) {
            _ = len(v1)
            v2 += (v2 - v2)
        } else {
            v3 = (^(v3))
            if (true && false) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v9 int16 = 1000
                _ = v9
            } else {
                if true {
                    if false {
                        var v10 int32 = 42
                        v4["hello"] = 800
                        fmt.Println(v2, v2, v2)
                        _ = v10
                    }
                    _ = len(v4)
                } else {
                    if false {
                        v2 = v2
                    }
                    v3 = (v3 / ((1) | 1))
                    v11 := GenS1{F0: false}
                    _ = v11
                }
                v12 := GenS2{F0: 755, F1: 398, F2: 2}
                v3 += (v3 * 2)
                _ = v12
            }
            fmt.Println(v3)
        }
    default:
        fmt.Println(v2, v3)
    }
    fmt.Println(v3)
    var v13 int32 = 0
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS2{F0: 2, F1: 42, F2: 42}
        v16 := GenS1{F0: true}
        _ = len(v1)
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int32 = (v13 % ((v13) | 1))
        v19 := GenS2{F0: 42, F1: 0, F2: 1}
        _ = fi17
        _ = v18
        _ = v19
    }
    if ((^(v2)) <= v2) {
        v20 := GenS3{F0: 128, F1: 504, F2: 1}
        _ = v20
    }
    fmt.Println(v13, v2)
    if (255 < 255) {
        v2 -= (v2 * v2)
        v2 = (v2 * (v2 * -61))
    }
    v21 := GenS3{F0: 0, F1: -884, F2: 42}
    var v22 int64 = genFn1("", -9223372036854775808, -1)
    var v23 uint32 = 0
    var v24 uint8 = 137
    v24 = (^(v24))
    switch "a" {
    case "":
        v25 := GenS4{F0: (^(v22)), F1: v13, F2: ((("foo bar" + "hello") + ("\n\t" + ("a" + ""))) + "foo bar")}
        _ = v25
    }
    v2 = v2
    fmt.Println(v2, v2, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v26 int = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
