package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint64) uint8 {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := []int{}
        p0 -= 42
        if (2 < 65535) {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 uint64 = p0
                _ = fi8
                _ = v9
            }
            if true {
                var v10 string = ((("a" + "foo bar") + "foo bar") + ("" + "hello"))
                _ = v10
            }
        } else {
            _ = len(v7)
        }
        _ = fi6
        _ = v7
    }
    if (false || (1.5 <= ((0.1 + -2.25) - 3.14159))) {
        fmt.Println(p0)
        var v11 string = "\n\t"
        var v12 int16 = 0
        _ = v11
        _ = v12
    }
    v13 := []int{32768, 0}
    var v14 uint32 = 1
    _ = p0
    _ = v13
    _ = v14
    return uint8(p0)
}

func main() {
    var v1 float32 = float32(0.1)
    var v2 int32 = 2147483647
    v3 := make(map[uint32]string)
    var v4 float64 = 0.1
    v5 := []string{"a"}
    v1 = ((v1 * (v1 / float32(v4))) / v1)
    v15 := GenS1{F0: -9223372036854775808, F1: 2}
    for fi16 := 0; fi16 < 15; fi16++ {
        if true {
            var v17 uint32 = 2147483648
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int8 = 127
                _ = len(v5)
                var v20 int16 = 165
                v21 := GenS2{F0: 42, F1: 42}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            _ = v17
        }
        _ = fi16
    }
    v1 -= 1.5
    v22 := GenS2{F0: 18446744073709551615, F1: 42}
    var v23 int = 32767
    var v24 int16 = -1
    switch ("foo bar" + "") {
    case "":
        var v25 int = 315
        _ = v25
    case "a":
        v23 -= v23
    case "\n\t":
        _ = len(v3)
    }
    switch (v23 - (v23 & v23)) {
    case -1:
        v26 := GenS2{F0: 2, F1: -9223372036854775808}
        _ = v26
    }
    v2 = (v2 + (v2 & (^(v2))))
    switch 0 {
    case 31:
        var v27 string = "\n\t"
        _ = v27
    case 289:
        v23 -= v23
    case 431:
        v23 = 128
    }
    var v28 string = "hello"
    v28 = v28
    var v29 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
