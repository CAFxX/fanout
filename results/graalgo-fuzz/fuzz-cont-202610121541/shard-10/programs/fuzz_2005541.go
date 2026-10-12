package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint64 }
type GenS2 struct { F0 int; F1 uint32; F2 uint8 }

func main() {
    var v1 int = 127
    var v2 uint8 = 0
    v3 := []string{"", "hello", (("foo bar" + "hello") + ("a" + ("\n\t" + ("hello" + "a"))))}
    var v4 int16 = 422
    var v5 uint64 = 2
    switch (v1 % ((v1) | 1)) {
    case 32767:
        v3 = append(v3, "foo bar")
    default:
        for fi6 := 0; fi6 < 15; fi6++ {
            if (v5 > 1) {
                fmt.Println(v1, v2)
                fmt.Println(v4, v1, v1)
                var v7 uint8 = v2
                _ = v7
            } else {
                switch (^(fi6)) {
                case -346:
                    var v8 int = fi6
                    _ = v8
                case 128:
                    for fi9 := 0; fi9 < 7; fi9++ {
                        v10 := GenS1{F0: true, F1: true, F2: v5}
                        _ = fi9
                        _ = v10
                    }
                default:
                    switch (fi6 | fi6) {
                    case 4294967295:
                        if true {
                            for fi11 := 0; fi11 < 12; fi11++ {
                                if false {
                                    v12 := GenS1{F0: false, F1: true, F2: 583}
                                    _ = v12
                                }
                                _ = fi11
                            }
                        }
                    }
                }
            }
            fmt.Println(v5)
            _ = fi6
        }
    }
    var v13 uint32 = 2147483647
    var v14 string = "foo bar"
    fmt.Println(v14)
    v2 -= (v2 & uint8(v13))
    var v15 string = (v14 + "hello")
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 float64 = float64(1.5)
        var v18 uint32 = 544
        for fi19 := 0; fi19 < 5; fi19++ {
            v5 -= (v5 + (^(v5)))
            v17 += (((v17 * v17) + 3.14159) * float64((1.5 + 0.1)))
            if true {
                fmt.Println(v13)
                v20 := GenS1{F0: false, F1: true, F2: v5}
                var v21 string = v15
                _ = v20
                _ = v21
            }
            _ = fi19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    var v22 int8 = 1
    var v23 string = ("" + v14)
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 uint16 = 2
        v26 := GenS1{F0: true, F1: (!(!(9223372036854775807 <= -9223372036854775808))), F2: v5}
        if (uint32(v25) >= v13) {
            var v27 int16 = (v4 / (((-(v4))) | 1))
            _ = v27
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        v3 = append(v3, v23)
        _ = fi28
    }
    var v29 uint16 = 42
    var v30 int16 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        fmt.Println(v2, v22, v4)
    } else {
        v3 = append(v3, v15)
        fmt.Println(v1, v5, v4)
    }
    if true {
        v31 := GenS2{F0: (-(v1)), F1: (v13 / ((0) | 1)), F2: uint8(v13)}
        v32 := GenS2{F0: v1, F1: (v13 % (((v13 * (v13 & 880))) | 1)), F2: v2}
        _ = v31
        _ = v32
    }
    v3 = append(v3, "hello")
    v5 -= uint64(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
