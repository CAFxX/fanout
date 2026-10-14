package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }

func main() {
    v1 := []uint32{2147483647}
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 2
    _ = len(v1)
    if true {
        var v4 float64 = (1.5 * 0.0)
        v2 += v2
        v3 = 42
        _ = v4
    }
    v3 -= (v3 << ((18446744073709551615) & 63))
    var v5 int = 127
    var v6 int16 = -1
    var v7 uint = 2
    var v8 uint8 = 129
    var v9 string = "hello"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!(!true)) {
        if ((true && true) || true) {
            fmt.Println(v3, v6)
        }
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 bool = (!(4294967295 != 192))
            if v12 {
                v13 := GenS1{F0: (v3 - (v3 & 0)), F1: fi11}
                _ = v13
            }
            switch (fi11 >> ((v7) & 63)) {
            case -9223372036854775808:
                v7 += (v7 ^ (v7 % ((0) | 1)))
            case -155:
                v2 = v2
            }
            _ = fi11
            _ = v12
        }
        v14 := GenS1{F0: 0, F1: (fi10 + 32768)}
        v8 = (v8 * (v8 & (v8 * (v8 + v8))))
        _ = fi10
        _ = v14
    }
    if ((false && true) || true) {
        v5 -= v5
        if (true && false) {
            v6 += (v6 & v6)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v15 := GenS1{F0: (v3 ^ v3), F1: 937}
            _ = v15
        } else {
            fmt.Println(v9)
            var v16 int = (v5 / ((-9223372036854775808) | 1))
            v17 := GenS1{F0: v3, F1: 32767}
            _ = v16
            _ = v17
        }
    } else {
        if true {
            if (-128 != 110) {
                fmt.Println(v5)
            }
        } else {
            var v18 string = (v9 + v9)
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 uint64 = v3
        switch (v9 + "foo bar") {
        case "":
            v5 = -2147483648
        case "\n\t":
            for fi21 := 0; fi21 < 20; fi21++ {
                v22 := GenS1{F0: 18446744073709551615, F1: fi19}
                switch 326 {
                case 4294967295:
                    if true {
                        fmt.Println(v3, v8, v2)
                    } else {
                        if true {
                            switch int(v3) {
                            case -354:
                                v2 = -1
                            }
                            v1 = append(v1, 877)
                            v6 += -32768
                        } else {
                            var v23 string = v9
                            _ = v23
                        }
                        var v24 float32 = 0.0
                        v8 += 255
                        fmt.Println(v20, v20, v3)
                        _ = v24
                    }
                default:
                    v9 = v9
                    v3 -= (v3 & v3)
                    v25 := GenS1{F0: v3, F1: 32768}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if false {
                        v1 = append(v1, 728)
                        v7 = 781
                        var v26 uint32 = 4294967295
                        _ = v26
                    } else {
                        v27 := GenS1{F0: 2, F1: (-(v5))}
                        var v28 float32 = (-2.25 - (1e10 - (-2.25 + 1.5)))
                        var v29 uint64 = v20
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    v20 = (^(v20))
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
        case "foo bar":
            for fi30 := 0; fi30 < 2; fi30++ {
                fmt.Println(fi19, v8, v2)
                _ = fi30
            }
        default:
            var v31 float32 = 1e10
            _ = v31
        }
        _ = fi19
        _ = v20
    }
    fmt.Println(v9, v8)
    var v32 uint = (v7 & (v7 / ((v7) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
