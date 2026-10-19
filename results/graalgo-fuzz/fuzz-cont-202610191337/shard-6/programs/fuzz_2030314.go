package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint8; F1 int32; F2 uint }
func genFn0(p0 uint, p1 uint32, p2 bool) uint16 {
    var v4 int16 = -32768
    p0 -= p0
    switch "\n\t" {
    case "":
        v5 := GenS1{F0: uint8(p0), F1: -3}
        _ = v5
    case "hello":
        var v6 float64 = (((0.0 / 1.0) - 0.0) - ((-2.25 * ((1e10 * (((0.0 / (0.0 + 1e10)) + 0.1) * (1.5 + ((float64(0.0) - -2.25) + -2.25)))) + 0.0)) * (1.5 * 1.5)))
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32767
}

func main() {
    var v1 int64 = -744
    var v2 float64 = (((3.14159 + 0.0) / 0.1) * 0.0)
    v3 := make(map[uint32]uint32)
    v7 := GenS1{F0: 2, F1: 2147483647}
    switch -9223372036854775808 {
    case -285:
        switch -1 {
        case 127:
            fmt.Println(v1, v2)
        case 128:
            v2 = ((1e10 - v2) - float64(0.1))
        default:
            v2 = (v2 - v2)
        }
    default:
        v8 := GenS2{F0: 32767}
        _ = v8
    }
    var v9 float64 = 0.1
    for fi10 := 0; fi10 < 4; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 float64 = ((1e10 - 0.0) + 1e10)
            var v13 uint = 883
            var v14 bool = true
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int32 = 594
        _ = fi15
        _ = v16
    }
    if (237 == 2) {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint64 = 968
            _ = fi17
            _ = v18
        }
    }
    switch (((("" + "hello") + "hello") + (("a" + ("a" + "")) + "a")) + "\n\t") {
    case "foo bar":
        var v19 int32 = 871
        _ = v19
    case "\n\t":
        v1 += v1
    }
    v20 := GenS3{F0: 0, F1: (("hello" + "hello") + "a")}
    var v21 string = ""
    switch "a" {
    case "foo bar":
        if (65535 >= 65535) {
            v22 := v3[65535]
            _ = v22
        }
    case "a":
        v23 := []int{}
        _ = v23
    case "\n\t":
        switch -456 {
        case -525:
            var v24 int = -9223372036854775808
            _ = v24
        case -403:
            v25 := []int{128}
            _ = v25
        case 1:
            for fi26 := 0; fi26 < 7; fi26++ {
                v1 = v1
                _ = fi26
            }
        default:
            var v27 float32 = -2.25
            _ = v27
        }
    default:
        v9 += 3.14159
    }
    if (((-2147483648 > 2147483647) && (v21 >= (v21 + v21))) && (128 != 0)) {
        v28 := []int{42}
        switch -89 {
        case -9223372036854775808:
            if false {
                var v29 int16 = -1
                v21 = ""
                v1 -= (v1 + (v1 << ((255) & 63)))
                _ = v29
            }
        }
        _ = v28
    } else {
        _ = len(v3)
    }
    var v30 bool = ((!false) && true)
    v30 = (false && (!true))
    v2 = v2
    if true {
        v21 = ""
        switch v21 {
        case "":
            v31 := GenS2{F0: 32767}
            _ = v31
        }
        v30 = (v30 && (!v30))
    } else {
        if v30 {
            v32 := GenS4{F0: 255, F1: 885, F2: 42}
            _ = v32
        }
        var v33 uint = 0
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v7
    _ = v9
}
