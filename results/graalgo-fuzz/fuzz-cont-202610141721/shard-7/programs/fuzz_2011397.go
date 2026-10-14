package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int16; F2 string }

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]bool)
    var v3 int32 = -2147483648
    switch -9223372036854775808 {
    case -1:
        switch 32768 {
        case 32767:
            v4 := GenS1{F0: 142}
            _ = v4
        case -392:
            v1 += 128
        default:
            for fi5 := 0; fi5 < 2; fi5++ {
                v6 := []int{32768, fi5}
                var v7 int64 = -352
                _ = fi5
                _ = v6
                _ = v7
            }
        }
    default:
        var v8 uint8 = (v1 * (^(v1)))
        _ = v8
    }
    if (-1 > -913) {
        v3 = (v3 * (v3 >> ((2) & 63)))
        var v9 int16 = 1000
        switch -9223372036854775808 {
        case 127:
            v3 -= v3
        case 1:
            v2[582] = false
        default:
            fmt.Println(v1)
        }
        _ = v9
    }
    var v10 int64 = -862
    v3 = 42
    var v11 float32 = (-2.25 + 0.0)
    var v12 int64 = -9223372036854775808
    v13 := GenS2{F0: -527, F1: int16(v10), F2: "\n\t"}
    v14 := GenS2{F0: -159, F1: 42, F2: ("" + "a")}
    v12 += -1
    if true {
        var v15 float32 = v11
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint16 = 955
            v3 = (-(v3))
            v11 = v11
            _ = fi17
            _ = v18
        }
        var v19 int = 128
        var v20 uint64 = 18446744073709551615
        _ = fi16
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 127}
    var v22 bool = (!(915 >= 255))
    if (v22 && (!(0 != 42))) {
        v23 := []int{-552}
        var v24 uint = 255
        _ = v23
        _ = v24
    } else {
        fmt.Println(v22, v3, v22)
        fmt.Println(v11, v1, v11)
        v11 += v11
    }
    fmt.Println(v10)
    switch ("\n\t" + ("" + "a")) {
    case "foo bar":
        switch 9223372036854775807 {
        case 207:
            v25 := GenS1{F0: v1}
            _ = v25
        }
    case "\n\t":
        var v26 float32 = (float32(0.0) - -2.25)
        _ = v26
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        v28 := GenS2{F0: -748, F1: -32768, F2: ("" + ((("" + (("foo bar" + ("a" + "")) + "a")) + "hello") + "\n\t"))}
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
