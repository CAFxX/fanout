package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int16 }
type GenS2 struct { F0 string }
func genFn0() int64 {
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := GenS1{F0: 1, F1: -149, F2: -574}
        fmt.Println(fi4)
        _ = fi4
        _ = v5
    }
    if (42 < 2) {
        var v6 bool = true
        v6 = (32767 > 0)
        _ = v6
    } else {
        v7 := []int{-945}
        _ = v7
    }
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 27
    var v2 int = -940
    var v3 string = "a"
    v8 := []int{}
    _ = len(v8)
    for fi9 := 0; fi9 < 3; fi9++ {
        if ((float32(0.1) <= 0.0) && (4294967295 == uint(v1))) {
            for fi10 := 0; fi10 < 6; fi10++ {
                v3 = v3
                _ = fi10
            }
            v2 = -9223372036854775808
        }
        fmt.Println(v2, v1)
        v11 := GenS1{F0: 1, F1: -1, F2: -899}
        _ = fi9
        _ = v11
    }
    v2 = (v2 / ((9223372036854775807) | 1))
    switch (v2 % ((-397) | 1)) {
    case 0:
        var v12 uint64 = 544
        _ = v12
    case 127:
        var v13 int = 9223372036854775807
        _ = v13
    default:
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int64 = 9223372036854775807
            var v16 uint32 = 315
            var v17 string = (v3 + v3)
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    }
    if true {
        if (!(32767 != 1)) {
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 string = "hello"
                v3 = "a"
                v2 = (-(v2))
                _ = fi18
                _ = v19
            }
            var v20 float32 = float32(float64(float32(float64(((1.5 - 1e10) / -2.25)))))
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 uint = 1
                v2 += (v2 ^ v2)
                _ = fi21
                _ = v22
            }
            _ = v20
        }
    }
    if (false || (1e10 > float32(0.1))) {
        _ = len(v8)
        v1 -= 32767
        switch (v3 + v3) {
        case "hello":
            v1 += v1
        }
    } else {
        if true {
            switch v3 {
            case "hello":
                var v23 float32 = (1.5 * 3.14159)
                _ = v23
            case "foo bar":
                v24 := GenS2{F0: v3}
                _ = v24
            }
            v25 := GenS2{F0: (v3 + "hello")}
            _ = v25
        }
        var v26 float64 = ((1.5 - (float64(0.0) - 1e10)) - (-2.25 / 1.0))
        for fi27 := 0; fi27 < 5; fi27++ {
            fmt.Println(fi27)
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi27
        }
        _ = v26
    }
    v1 = v1
    var v28 int32 = 0
    for fi29 := 0; fi29 < 7; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            v1 = (v1 - v1)
            fmt.Println(v2)
            _ = fi30
        }
        _ = fi29
    }
    v28 = (v28 & 1)
    if (!true) {
        var v31 float32 = 0.1
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 int = -9223372036854775808
            _ = fi32
            _ = v33
        }
        _ = v31
    } else {
        for fi34 := 0; fi34 < 11; fi34++ {
            for fi35 := 0; fi35 < 15; fi35++ {
                v1 = v1
                var v36 int16 = 0
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    }
    v3 = ((v3 + (v3 + v3)) + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v8
}
