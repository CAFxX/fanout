package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 103
    var v3 int8 = 1
    v4 := []uint32{42, 669, 2147483648}
    var v5 uint32 = 65535
    v4 = append(v4, (v5 & (^(v5))))
    v6 := GenS1{F0: 4294967295}
    if false {
        var v7 uint8 = 78
        v8 := GenS1{F0: 1}
        for fi9 := 0; fi9 < 18; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                v7 = v2
                _ = fi10
            }
            var v11 uint8 = (v7 - 128)
            var v12 float32 = 1.5
            _ = fi9
            _ = v11
            _ = v12
        }
        _ = v7
        _ = v8
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            if true {
                v14 := GenS1{F0: 42}
                if false {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v5 -= (v5 | (v5 * (v5 ^ v5)))
                        _ = fi15
                    }
                    v3 += (^(v3))
                }
                _ = v14
            }
            _ = fi13
        }
        v1[9223372036854775807] = ("a" + "a")
        if false {
            var v16 uint8 = v2
            _ = v16
        } else {
            v2 -= v2
        }
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 float64 = 0.1
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: 729, F1: "hello"}
    if false {
        v2 -= v2
        v20 := v1[-534]
        _ = v20
    } else {
        switch "\n\t" {
        case "":
            v2 = 24
        }
        v21 := GenS1{F0: 42}
        _ = v21
    }
    v3 = 127
    if true {
        fmt.Println(v2, v3)
        switch -146 {
        case 4294967295:
            if false {
                var v22 bool = true
                _ = v22
            }
        case -1:
            if true {
                v5 = (v5 ^ 4294967295)
                v2 = (v2 * 160)
                var v23 uint32 = v5
                _ = v23
            }
        case -836:
            switch -507 {
            case 32768:
                v24 := GenS1{F0: uint(v2)}
                _ = v24
            case -668:
                var v25 int8 = (-(v3))
                _ = v25
            case 32767:
                if (1 >= 303) {
                    for fi26 := 0; fi26 < 13; fi26++ {
                        v27 := GenS2{F0: 428, F1: "a"}
                        _ = fi26
                        _ = v27
                    }
                    _ = len(v1)
                    var v28 uint8 = v2
                    _ = v28
                }
            default:
                fmt.Println(v2)
            }
        default:
            v2 -= (v2 + v2)
        }
        v5 = v5
    } else {
        switch "foo bar" {
        case "\n\t":
            switch "a" {
            case "foo bar":
                fmt.Println(v3, v2, v5)
            default:
                v5 = v5
            }
        case "hello":
            var v29 uint8 = v2
            _ = v29
        }
        var v30 uint32 = (^(v5))
        fmt.Println(v5, v5)
        _ = v30
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        if (true && false) {
            v3 -= (v3 ^ int8(v2))
            fmt.Println(fi31, v2)
        }
        _ = fi31
    }
    fmt.Println(v2, v5, v3)
    if true {
        if (true || false) {
            var v32 int32 = 514
            _ = v32
        }
        v5 = v5
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        var v34 int64 = int64(v3)
        _ = fi33
        _ = v34
    }
    if false {
        var v35 uint64 = 0
        _ = v35
    } else {
        fmt.Println(v3, v2)
    }
    var v36 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
