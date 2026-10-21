package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }
type GenS2 struct { F0 int8; F1 string; F2 bool }
func genFn0(p0 int64, p1 int) int64 {
    var v5 int8 = 42
    p1 = p1
    v6 := GenS1{F0: 1, F1: 318, F2: 2147483647}
    fmt.Println(p1)
    if false {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int8 = v5
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 768, F1: 42, F2: -1}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(fi10, fi10)
            p1 = p1
            _ = fi10
        }
        fmt.Println(v5, p1)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 float64 = (-2.25 + (1e10 - -2.25))
    var v2 int16 = -32768
    var v3 float64 = (0.1 - v1)
    var v4 int64 = 759
    v3 += -2.25
    switch "hello" {
    case "a":
        switch ("hello" + ("a" + "a")) {
        case "\n\t":
            var v11 float32 = ((1.5 / 1.0) * (0.0 / 1.0))
            _ = v11
        case "a":
            v3 -= v1
        case "":
            if true {
                for fi12 := 0; fi12 < 17; fi12++ {
                    switch "hello" {
                    case "hello":
                        var v13 int16 = (v2 | (v2 >> ((967) & 63)))
                        _ = v13
                    case "foo bar":
                        v2 += v2
                    }
                    v14 := []int{}
                    _ = fi12
                    _ = v14
                }
            }
        default:
            v4 += (v4 >> ((0) & 63))
        }
    case "":
        v15 := []int{127, 0, 0}
        _ = v15
    case "hello":
        v16 := []int{127, 32768, -218}
        _ = v16
    default:
        var v17 uint16 = 42
        _ = v17
    }
    v18 := []int{297, 32767}
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := GenS1{F0: 2, F1: 18446744073709551615, F2: 1}
        if len(v18) > 0 {
            v18[0] = fi19
        }
        _ = fi19
        _ = v20
    }
    if (!((v2 % (((v2 ^ v2)) | 1)) < 1000)) {
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := GenS1{F0: 42, F1: 255, F2: 65}
            switch (fi21 ^ (^(fi21))) {
            case 4294967295:
                v3 -= (((v1 / (((float64(((0.0 * 1e10) + 0.0)) / 1.5) - v1) / 3.14159)) / v3) * (v1 / v3))
            case 833:
                var v23 uint8 = 174
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        var v24 int = 1
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        _ = v24
    }
    v25 := GenS2{F0: 100, F1: ("foo bar" + "a"), F2: true}
    var v26 int8 = 127
    if len(v18) > 0 {
        v18[0] = 32767
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        if len(v18) > 0 {
            v18[0] = fi27
        }
        var v28 uint64 = 238
        _ = fi27
        _ = v28
    }
    var v29 uint8 = 26
    var v30 uint8 = 128
    _ = len(v18)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
