package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int; F2 uint32 }
func genFn0(p0 int64, p1 bool, p2 string) uint8 {
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int = 295
        var v7 float64 = 3.14159
        switch p2 {
        case "\n\t":
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 int8 = -82
                var v10 string = (p2 + "\n\t")
                if false {
                    var v11 int8 = (v9 % ((v9) | 1))
                    _ = v11
                } else {
                    var v12 int = (fi8 & fi8)
                    v13 := []int{(v12 | (v12 ^ (fi5 + (v12 & 1)))), 4294967295}
                    fmt.Println(fi5)
                    _ = v12
                    _ = v13
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v14 int = 42
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 string = ("foo bar" + p2)
        switch (v14 & 2147483647) {
        case 9223372036854775807:
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := []int{v14}
                v16 = v16
                _ = fi17
                _ = v18
            }
        default:
            v19 := GenS1{F0: p1}
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 255
}
func genFn1(p0 uint16, p1 int8, p2 int32) int16 {
    v20 := GenS1{F0: false}
    var v21 uint32 = 1
    p1 += int8(v21)
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 42
}

func main() {
    v1 := []int{1}
    var v2 int16 = -891
    v3 := []int{-692}
    v4 := []string{""}
    var v22 int = 6
    if (296 != 65535) {
        fmt.Println(v22, v2)
        if len(v1) > 0 {
            v1[0] = (^(v22))
        }
        var v23 uint64 = 1
        _ = v23
    }
    v2 += 42
    fmt.Println(v22, v22)
    if true {
        v24 := GenS2{F0: genFn1(255, 100, 42), F1: 0, F2: 2147483647}
        var v25 int8 = 127
        var v26 int8 = v25
        _ = v24
        _ = v25
        _ = v26
    }
    fmt.Println(v2, v22)
    if (42 == 1) {
        var v27 uint8 = genFn0(-131, false, ("foo bar" + ("a" + "a")))
        _ = v27
    }
    var v28 bool = (1 < 4294967295)
    v28 = ((v22 * (-(v22))) != v22)
    switch "\n\t" {
    case "":
        switch "hello" {
        case "foo bar":
            var v29 uint64 = 4294967295
            _ = v29
        case "\n\t":
            v2 += genFn1(0, -128, 563)
        }
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        fmt.Println(v28, fi30, fi30)
        v28 = (!(v28 || false))
        v22 = v22
        _ = fi30
    }
    if (!v28) {
        v28 = v28
    } else {
        v22 -= (v22 * (v22 - int(v2)))
    }
    v2 += genFn1(2, 42, 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
}
