package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 bool }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int16 }

func main() {
    var v1 float32 = 0.0
    var v2 float64 = (3.14159 / 3.14159)
    v3 := make(map[string]uint32)
    var v4 bool = (!false)
    var v5 float64 = v2
    v6 := GenS1{F0: 877, F1: v4, F2: (592 > 255)}
    v7 := GenS1{F0: 0, F1: (2 < 659), F2: (!(127 < 32767))}
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(v5, fi8)
        v4 = (!true)
        v2 = v5
        _ = fi8
    }
    if v4 {
        v9 := GenS2{F0: "", F1: 2147483647}
        if ((1.5 > 0.1) && (2 <= 408)) {
            v10 := GenS2{F0: ("\n\t" + (("foo bar" + "hello") + ("hello" + "\n\t"))), F1: -2147483648}
            _ = v10
        } else {
            v1 = v1
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(v2)
            var v12 float64 = (0.1 + v5)
            _ = fi11
            _ = v12
        }
        _ = v9
    } else {
        fmt.Println(v2, v2)
    }
    if (v4 && (18446744073709551615 > 18446744073709551615)) {
        var v13 bool = true
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS3{F0: -142, F1: 2}
        _ = fi14
        _ = v15
    }
    v16 := GenS4{F0: 655}
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int32 = 1
        var v19 float32 = v1
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        v4 = ((1e10 < -2.25) && v4)
        for fi21 := 0; fi21 < 1; fi21++ {
            v2 = v2
            var v22 uint8 = 127
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v23 := []int{-2147483648, 358, 42, 0}
    v23 = append(v23, 1)
    var v24 uint16 = 65535
    for fi25 := 0; fi25 < 3; fi25++ {
        if (431 <= 4294967295) {
            v2 += v2
            if true {
                var v26 uint = 1
                _ = v26
            }
            v24 += (v24 / ((166) | 1))
        }
        v5 -= (v2 + 1.5)
        _ = fi25
    }
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
