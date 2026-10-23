package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint64; F2 int64 }

func main() {
    v1 := []int{627}
    var v2 uint32 = 2
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: 128, F1: (980 == 4294967295), F2: 2147483647}
    if true {
        v2 += (v2 / ((42) | 1))
        for fi5 := 0; fi5 < 15; fi5++ {
            v2 -= v2
            _ = fi5
        }
    } else {
        v3[(^(v2))] = ("\n\t" + "hello")
        switch "\n\t" {
        case "hello":
            v6 := GenS2{F0: 128}
            _ = v6
        default:
            switch "foo bar" {
            case "a":
                v2 -= 1
            case "foo bar":
                v7 := GenS1{F0: 127, F1: true, F2: 973}
                _ = v7
            case "hello":
                fmt.Println(v2)
            }
        }
    }
    v2 -= v2
    v8 := GenS2{F0: 255}
    if (false && false) {
        for fi9 := 0; fi9 < 3; fi9++ {
            switch (fi9 / (((fi9 << ((969) & 63))) | 1)) {
            case 9223372036854775807:
                var v10 float32 = (-2.25 + 1e10)
                _ = v10
            case 42:
                var v11 int64 = -1
                _ = v11
            }
            v3[v2] = "hello"
            v1 = append(v1, fi9)
            _ = fi9
        }
    } else {
        v2 -= v2
    }
    if false {
        switch 32767 {
        case 127:
            for fi12 := 0; fi12 < 1; fi12++ {
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 uint = 42
                    delete(v3, v2)
                    _ = fi13
                    _ = v14
                }
                v2 += v2
                if len(v1) > 0 {
                    v1[0] = (fi12 ^ (fi12 / ((fi12) | 1)))
                }
                _ = fi12
            }
        default:
            for fi15 := 0; fi15 < 7; fi15++ {
                if true {
                    var v16 float64 = 3.14159
                    _ = v16
                }
                switch (("a" + ((("hello" + "foo bar") + (("a" + (("" + ("" + "foo bar")) + "foo bar")) + "hello")) + ("foo bar" + ""))) + "") {
                case "":
                    switch "a" {
                    case "a":
                        v2 -= v2
                    default:
                        _ = len(v1)
                    }
                }
                _ = fi15
            }
        }
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            fmt.Println(v2, v2)
            v2 = v2
            v18 := GenS3{F0: 100, F1: 1}
            _ = fi17
            _ = v18
        }
        fmt.Println(v2)
    }
    switch "hello" {
    case "\n\t":
        var v19 uint64 = 53
        _ = v19
    }
    v2 = (v2 ^ 1)
    if false {
        if true {
            v2 += v2
            for fi20 := 0; fi20 < 9; fi20++ {
                v2 -= (v2 / ((16) | 1))
                for fi21 := 0; fi21 < 16; fi21++ {
                    v22 := GenS3{F0: 125, F1: 65535}
                    v2 = 2147483647
                    var v23 uint16 = 815
                    fmt.Println(v23, fi21)
                    if ((true || (42 != 42)) || (32767 > 388)) {
                        v24 := GenS4{F0: 42, F1: 42, F2: -16}
                        v2 -= v2
                        v23 = v23
                        _ = v24
                    }
                    var v25 bool = true
                    var v26 uint16 = (v23 ^ v23)
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v25
                    _ = v26
                }
                v2 -= 708
                _ = fi20
            }
            if (true && true) {
                if (0.1 != -2.25) {
                    var v27 string = "hello"
                    var v28 int8 = 100
                    _ = v27
                    _ = v28
                }
                var v29 int8 = -125
                _ = len(v1)
                _ = v29
            }
        }
    } else {
        fmt.Println(v2)
        fmt.Println(v2)
        var v30 int64 = 339
        _ = v30
    }
    if len(v1) > 0 {
        v1[0] = -1
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        v32 := GenS4{F0: -1, F1: 0, F2: int64(v2)}
        fmt.Println(v2, v2)
        v2 -= (v2 | 4294967295)
        _ = fi31
        _ = v32
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v2 = (v2 | v2)
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
