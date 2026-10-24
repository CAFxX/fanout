package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS2 struct { F0 int64; F1 uint }

func main() {
    var v1 int16 = 42
    var v2 float32 = 0.0
    v3 := []int{89, 1}
    for fi4 := 0; fi4 < 7; fi4++ {
        if false {
            var v5 uint = 65535
            _ = v5
        } else {
            var v6 uint64 = 2
            _ = v6
        }
        switch "a" {
        case "\n\t":
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 uint32 = 42
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v9 uint64 = 4294967295
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 4; fi10++ {
            v1 -= v1
            switch (("foo bar" + ("hello" + "hello")) + "\n\t") {
            case "":
                _ = len(v3)
            case "a":
                var v11 int = (fi10 + fi4)
                _ = v11
            case "foo bar":
                switch (fi10 ^ (fi10 >> ((65535) & 63))) {
                case -1:
                    v1 = (v1 / (((-(v1))) | 1))
                default:
                    v1 = v1
                }
            }
            if false {
                var v12 string = "a"
                _ = v12
            }
            _ = fi10
        }
        _ = fi4
    }
    switch 2147483647 {
    case -2147483648:
        v2 = v2
    case 32767:
        v13 := GenS1{F0: 2147483647, F1: 1, F2: 100}
        _ = v13
    case 32768:
        if (1 == (v1 ^ v1)) {
            v1 -= v1
            if len(v3) > 0 {
                v3[0] = -450
            }
            v14 := GenS1{F0: 948, F1: 452, F2: 122}
            _ = v14
        }
    default:
        var v15 float64 = (0.0 * 0.0)
        _ = v15
    }
    _ = len(v3)
    fmt.Println(v1, v1)
    var v16 int64 = 9223372036854775807
    switch -221 {
    case 114:
        _ = len(v3)
    case 4294967295:
        var v17 int64 = v16
        _ = v17
    default:
        v18 := GenS1{F0: 2, F1: 255, F2: -128}
        _ = v18
    }
    fmt.Println(v16, v2, v2)
    v16 -= (v16 * int64(v1))
    v3 = append(v3, -1)
    if true {
        var v19 uint = 0
        v19 = v19
        if ((!true) && (0 < (-(v1)))) {
            var v20 uint64 = 956
            var v21 uint8 = 1
            var v22 uint8 = 170
            _ = v20
            _ = v21
            _ = v22
        } else {
            if (v2 >= ((v2 + ((3.14159 - v2) + -2.25)) / (((v2 + 0.0) - v2) / v2))) {
                fmt.Println(v2, v16, v2)
            }
            v23 := GenS1{F0: 65535, F1: 255, F2: -128}
            var v24 int16 = -1
            _ = v23
            _ = v24
        }
        _ = v19
    }
    if false {
        for fi25 := 0; fi25 < 16; fi25++ {
            fmt.Println(v2)
            if len(v3) > 0 {
                v3[0] = (fi25 * (fi25 + fi25))
            }
            var v26 uint32 = 1
            _ = fi25
            _ = v26
        }
        fmt.Println(v2, v1, v1)
        v2 = v2
    } else {
        if true {
            var v27 int = 2147483647
            if len(v3) > 0 {
                v3[0] = (v27 + v27)
            }
            for fi28 := 0; fi28 < 15; fi28++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v29 := GenS2{F0: -9223372036854775808, F1: 0}
                _ = fi28
                _ = v29
            }
            _ = v27
        }
        var v30 int = 1
        if true {
            for fi31 := 0; fi31 < 20; fi31++ {
                v1 = v1
                var v32 bool = (false || true)
                _ = fi31
                _ = v32
            }
            v2 = 0.0
            var v33 float64 = (0.1 / (0.1 * -2.25))
            _ = v33
        } else {
            var v34 int32 = 303
            _ = v34
        }
        _ = v30
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        switch (fi35 % ((fi35) | 1)) {
        case -441:
            fmt.Println(fi35, v16)
        default:
            v16 = -9223372036854775808
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
}
