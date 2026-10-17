package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint }
func genFn0() uint {
    if true {
        if true {
            var v4 uint16 = 32767
            v5 := []int{-2147483648, -80, -284, 127}
            if len(v5) > 0 {
                v5[0] = 1
            }
            _ = v4
            _ = v5
        }
        if (!(!(!true))) {
            var v6 float32 = 3.14159
            _ = v6
        } else {
            var v7 float64 = -2.25
            v7 -= 3.14159
            _ = v7
        }
    }
    var v8 float64 = (0.1 * 1.5)
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := []int{127, fi9}
        _ = fi9
        _ = v10
    }
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 0
    var v3 uint16 = 255
    v11 := GenS1{F0: (128 >= 42)}
    v12 := []int{32768, 9223372036854775807}
    for fi13 := 0; fi13 < 17; fi13++ {
        if (0.1 < 0.1) {
            switch "a" {
            case "a":
                fmt.Println(fi13, v2, v2)
            case "hello":
                v2 -= 865
            }
            v14 := GenS2{F0: (v3 ^ 65535)}
            v3 += v3
            _ = v14
        } else {
            fmt.Println(fi13)
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            fmt.Println(fi13)
            _ = fi15
        }
        fmt.Println(v3, v2, v3)
        _ = fi13
    }
    var v16 bool = false
    if false {
        var v17 uint8 = 2
        if (977 < 1) {
            if (false && v16) {
                var v18 int32 = -762
                v18 -= v18
                _ = v18
            } else {
                v19 := v1[-525]
                _ = v19
            }
            v20 := v1[-2147483648]
            var v21 uint8 = v17
            _ = v20
            _ = v21
        }
        switch -633 {
        case 9223372036854775807:
            fmt.Println(v3, v2)
        default:
            v22 := GenS2{F0: 65535}
            _ = v22
        }
        _ = v17
    }
    v23 := GenS3{F0: 127, F1: 32767}
    var v24 int = 128
    var v25 float64 = 1e10
    _ = len(v12)
    for fi26 := 0; fi26 < 4; fi26++ {
        _ = len(v12)
        _ = fi26
    }
    v16 = v16
    if (((-2.25 / 1.0) == float32(v25)) && (18446744073709551615 < 0)) {
        if (2 == 255) {
            switch ("" + "foo bar") {
            case "\n\t":
                var v27 int8 = 28
                _ = v27
            case "":
                v3 = v3
            default:
                var v28 uint16 = v3
                _ = v28
            }
            var v29 int32 = 2147483647
            var v30 int32 = v29
            _ = v29
            _ = v30
        }
        fmt.Println(v25, v24)
        var v31 bool = (!(-774 <= -52))
        _ = v31
    }
    if false {
        var v32 int8 = 0
        if v16 {
            switch ("hello" + "\n\t") {
            case "foo bar":
                var v33 uint64 = 0
                _ = v33
            }
            if v16 {
                v1[-378] = v24
                v2 = (^(v2))
                v34 := GenS4{F0: 1}
                _ = v34
            } else {
                v16 = (v3 >= (v3 & v3))
                fmt.Println(v24, v32, v32)
            }
            v25 += v25
        }
        _ = v32
    } else {
        for fi35 := 0; fi35 < 17; fi35++ {
            v25 -= v25
            _ = fi35
        }
        for fi36 := 0; fi36 < 11; fi36++ {
            fmt.Println(v2, v2, v16)
            _ = fi36
        }
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        for fi38 := 0; fi38 < 20; fi38++ {
            v39 := GenS1{F0: (int(v3) != 32768)}
            fmt.Println(fi37)
            _ = fi38
            _ = v39
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
