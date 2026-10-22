package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int16 }
func genFn0() uint {
    if (((("\n\t" + "") + (((("foo bar" + ("hello" + "")) + "foo bar") + ("" + "hello")) + ("" + "foo bar"))) == "hello") || false) {
        v5 := GenS1{F0: -9223372036854775808, F1: -84}
        if (1000 < -32768) {
            var v6 string = (("\n\t" + "a") + ("foo bar" + "a"))
            v6 = v6
            var v7 int = 585
            _ = v6
            _ = v7
        }
        switch -710 {
        case -101:
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 string = "foo bar"
                v9 = ""
                _ = fi8
                _ = v9
            }
        }
        _ = v5
    } else {
        var v10 uint8 = 42
        _ = v10
    }
    v11 := GenS2{F0: 0}
    if false {
        v12 := GenS3{F0: 9223372036854775807, F1: 611, F2: -128}
        for fi13 := 0; fi13 < 4; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 int8 = -128
                v15 += (v15 / ((-29) | 1))
                v16 := GenS2{F0: 211}
                _ = fi14
                _ = v15
                _ = v16
            }
            if (1.5 > 3.14159) {
                v17 := GenS4{F0: 1, F1: 660, F2: 16}
                _ = v17
            }
            v18 := []int{(fi13 % ((fi13) | 1))}
            _ = fi13
            _ = v18
        }
        _ = v12
    } else {
        var v19 uint32 = 2
        _ = v19
    }
    if (!((1 < 2) || false)) {
        var v20 int16 = 341
        _ = v20
    }
    _ = v11
    return 85
}
func genFn1() uint64 {
    var v21 bool = (1 < 42)
    if v21 {
        var v22 uint16 = 32767
        _ = v22
    }
    var v23 int = 1
    v24 := []int{v23, 32768, v23}
    _ = v21
    _ = v23
    _ = v24
    return 613
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 0
    var v3 uint32 = 4294967295
    v4 := []uint32{}
    for fi25 := 0; fi25 < 17; fi25++ {
        _ = len(v4)
        switch "\n\t" {
        case "":
            if len(v4) > 0 {
                v4[0] = v3
            }
        case "\n\t":
            switch ("" + "foo bar") {
            case "a":
                var v26 int32 = 1
                _ = v26
            }
        }
        v27 := v1[(v3 >> ((65535) & 63))]
        _ = fi25
        _ = v27
    }
    v3 = v3
    var v28 int32 = 2147483647
    for fi29 := 0; fi29 < 19; fi29++ {
        fmt.Println(v28, v3, v28)
        _ = fi29
    }
    v28 -= v28
    var v30 int64 = 956
    var v31 uint64 = genFn1()
    v3 = (^(v3))
    v4 = append(v4, uint32(v31))
    var v32 string = "foo bar"
    var v33 uint16 = uint16(v31)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
