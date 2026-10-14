package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint = 1
    var v2 float64 = (float64(-2.25) / float64((0.1 + 0.1)))
    var v3 uint32 = 2147483647
    var v4 bool = ((false || ("foo bar" < "")) && false)
    switch (("" + "\n\t") + "\n\t") {
    case "hello":
        if (-60 >= 88) {
            for fi5 := 0; fi5 < 11; fi5++ {
                v4 = (false && v4)
                _ = fi5
            }
            v1 = v1
        }
    case "foo bar":
        if ((true || (((v2 + (v2 - (v2 + ((v2 * ((-2.25 + v2) * ((v2 / ((-2.25 + (v2 * v2)) - ((((v2 * (v2 + 0.0)) / ((v2 * v2) - 1.5)) / (((v2 * 0.0) + v2) * 0.0)) - 3.14159))) * ((v2 + -2.25) + (v2 - float64(0.0)))))) * v2)))) / (v2 / (1e10 - (((v2 * ((3.14159 - (((-2.25 * v2) + (float64((1.5 * 0.1)) - v2)) - 1.5)) + v2)) * (v2 + float64(1e10))) - 3.14159)))) >= v2)) && false) {
            v2 += v2
        }
    }
    v6 := GenS1{F0: 1, F1: v3, F2: 64}
    v2 += v2
    for fi7 := 0; fi7 < 20; fi7++ {
        switch (fi7 << ((v1) & 63)) {
        case 42:
            var v8 uint16 = 268
            _ = v8
        case -1:
            var v9 int64 = -492
            _ = v9
        case -2147483648:
            v1 = 4294967295
        default:
            switch "foo bar" {
            case "foo bar":
                if true {
                    for fi10 := 0; fi10 < 5; fi10++ {
                        fmt.Println(v1, v3, v3)
                        v2 -= 0.0
                        var v11 uint8 = 233
                        v3 = 42
                        _ = fi10
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v13 := []int{(fi7 >> (((v1 << ((v1) & 63))) & 63)), 0, (fi7 & 9223372036854775807)}
                        _ = fi12
                        _ = v13
                    }
                }
            }
        }
        _ = fi7
    }
    v14 := []int{-826}
    var v15 int = -2147483648
    v4 = (v4 || (315 > 455))
    switch (-(v15)) {
    case 365:
        var v16 int32 = 42
        _ = v16
    case 127:
        if false {
            var v17 uint64 = 1
            _ = v17
        } else {
            var v18 int8 = 42
            _ = v18
        }
    }
    var v19 float32 = 1e10
    fmt.Println(v19)
    if ((-1 != 0) && v4) {
        v4 = v4
    } else {
        v20 := GenS1{F0: -207, F1: v3, F2: -400}
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 uint = 0
            _ = fi22
            _ = v23
        }
        var v24 int16 = -1
        _ = fi21
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        for fi26 := 0; fi26 < 2; fi26++ {
            v19 -= v19
            v27 := GenS2{F0: 0}
            _ = fi26
            _ = v27
        }
        v28 := GenS1{F0: 1, F1: (v3 >> ((v1) & 63)), F2: 1}
        var v29 int64 = -1
        _ = fi25
        _ = v28
        _ = v29
    }
    var v30 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v6
}
