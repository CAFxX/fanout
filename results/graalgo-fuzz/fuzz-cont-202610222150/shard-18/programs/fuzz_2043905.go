package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 bool; F1 uint }

func main() {
    v1 := []uint32{313, 4294967295, 454}
    var v2 float64 = (1e10 / 1.0)
    var v3 int8 = 100
    var v4 float64 = (3.14159 + v2)
    var v5 int16 = -32768
    var v6 int32 = -705
    switch ("\n\t" + "") {
    case "a":
        if false {
            _ = len(v1)
            for fi7 := 0; fi7 < 6; fi7++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v5 -= -1
                v5 -= v5
                _ = fi7
            }
            for fi8 := 0; fi8 < 19; fi8++ {
                v6 -= 353
                if false {
                    if true {
                        v3 += 103
                    } else {
                        v9 := GenS1{F0: 736, F1: 780, F2: 2147483647}
                        _ = v9
                    }
                    v6 = -1
                } else {
                    switch "hello" {
                    case "hello":
                        if len(v1) > 0 {
                            v1[0] = 240
                        }
                    case "foo bar":
                        v10 := GenS1{F0: 694, F1: 18446744073709551615, F2: 2147483647}
                        _ = v10
                    }
                    v6 = v6
                }
                if false {
                    var v11 uint16 = 255
                    v12 := GenS2{F0: 29, F1: (v6 ^ v6)}
                    _ = v11
                    _ = v12
                } else {
                    v4 = v2
                }
                _ = fi8
            }
        } else {
            var v13 uint64 = 42
            _ = len(v1)
            var v14 string = ("\n\t" + ("hello" + ("" + "\n\t")))
            _ = v13
            _ = v14
        }
    case "":
        var v15 int8 = 110
        _ = v15
    }
    var v16 float64 = 1.5
    v6 -= (v6 % ((2147483647) | 1))
    if (v6 > v6) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v17 := GenS3{F0: (339 > 18446744073709551615), F1: 18446744073709551615}
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        _ = v17
    }
    v18 := GenS2{F0: 202, F1: v6}
    _ = len(v1)
    var v19 int = 1
    if (v5 < (v5 << ((0) & 63))) {
        if true {
            var v20 float64 = 1e10
            var v21 uint32 = 0
            _ = v20
            _ = v21
        }
    }
    if (98 == 1) {
        fmt.Println(v16)
        if (!true) {
            v16 = v4
            fmt.Println(v3)
            fmt.Println(v3, v3, v2)
        } else {
            if ((!(v5 < 1000)) || true) {
                fmt.Println(v16, v19)
                if (true && true) {
                    v6 -= 1
                    v4 = v16
                    if (!false) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v2 -= -2.25
                    }
                } else {
                    var v22 int = (v19 % (((^(v19))) | 1))
                    v16 = v4
                    v6 = v6
                    _ = v22
                }
                var v23 uint8 = 128
                _ = v23
            }
            var v24 uint8 = 42
            _ = v24
        }
    } else {
        v6 -= v6
        var v25 int16 = 0
        var v26 float32 = 1.5
        _ = v25
        _ = v26
    }
    v4 = 0.0
    fmt.Println(v5)
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS1{F0: 826, F1: 1, F2: (v6 + (v6 / ((v6) | 1)))}
        v29 := GenS3{F0: (!true), F1: 993}
        _ = fi27
        _ = v28
        _ = v29
    }
    v5 += (v5 / (((v5 >> ((560) & 63))) | 1))
    var v30 uint16 = 255
    v31 := GenS2{F0: 0, F1: 670}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
