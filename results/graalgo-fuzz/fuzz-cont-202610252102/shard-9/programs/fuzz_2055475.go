package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0() uint8 {
    var v4 float64 = (-2.25 / 1.0)
    switch "\n\t" {
    case "hello":
        var v5 uint = 2
        _ = v5
    default:
        var v6 int32 = 215
        _ = v6
    }
    switch 128 {
    case 2147483647:
        var v7 string = (("foo bar" + "a") + (("hello" + "") + ((("a" + "\n\t") + "") + "")))
        _ = v7
    case 42:
        var v8 uint = 2
        _ = v8
    }
    fmt.Println(v4)
    _ = v4
    return 226
}
func genFn1() uint16 {
    if true {
        switch "\n\t" {
        case "":
            var v9 float64 = ((0.0 * 0.0) * 0.0)
            _ = v9
        case "\n\t":
            switch 2147483647 {
            case -922:
                for fi10 := 0; fi10 < 19; fi10++ {
                    var v11 int = (fi10 % (((^(fi10))) | 1))
                    v12 := []int{fi10, -860, fi10}
                    if false {
                        var v13 string = (("" + ("" + "hello")) + "a")
                        switch "\n\t" {
                        case "foo bar":
                            for fi14 := 0; fi14 < 9; fi14++ {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                v13 = v13
                                var v15 uint64 = 241
                                _ = fi14
                                _ = v15
                            }
                        case "a":
                            for fi16 := 0; fi16 < 18; fi16++ {
                                switch (fi16 << ((63) & 63)) {
                                case 4294967295:
                                    var v17 uint = 2
                                    _ = v17
                                case 32767:
                                    v11 -= 32767
                                default:
                                    if (-2.25 < (((((1e10 / 1.0) * (float64(0.0) * 1e10)) / 1e10) - 1.5) / (3.14159 / 1.5))) {
                                        if true {
                                            fmt.Println(v13)
                                            switch "hello" {
                                            case "a":
                                                if (!false) {
                                                    var v18 int32 = 0
                                                    var v19 uint64 = 4294967295
                                                    _ = v18
                                                    _ = v19
                                                }
                                            }
                                        } else {
                                            _ = len(v12)
                                            switch v11 {
                                            case 0:
                                                fmt.Println(v11, fi10, v13)
                                            }
                                            if false {
                                                v20 := GenS1{F0: 1, F1: false}
                                                var v21 uint64 = 364
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                                _ = v20
                                                _ = v21
                                            } else {
                                                v11 = (^(fi10))
                                            }
                                        }
                                    }
                                }
                                var v22 bool = (127 == 1)
                                _ = fi16
                                _ = v22
                            }
                        default:
                            v13 = v13
                        }
                        _ = v13
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case 760:
                var v23 int64 = -766
                _ = v23
            case 4294967295:
                var v24 uint = 84
                _ = v24
            }
        }
    }
    v25 := []int{-9223372036854775808, 42, -9223372036854775808}
    _ = v25
    return 385
}

func main() {
    var v1 uint = 558
    v2 := []uint32{}
    var v3 float64 = 0.0
    fmt.Println(v1)
    v3 = v3
    if (!true) {
        _ = len(v2)
    } else {
        if false {
            switch 4294967295 {
            case -9223372036854775808:
                for fi26 := 0; fi26 < 16; fi26++ {
                    v3 -= v3
                    var v27 bool = (("\n\t" + ("foo bar" + ("\n\t" + ("\n\t" + "\n\t")))) <= "foo bar")
                    _ = fi26
                    _ = v27
                }
            case 32768:
                var v28 int32 = 2147483647
                _ = v28
            case 128:
                var v29 uint32 = uint32(v1)
                _ = v29
            }
            v30 := GenS2{F0: 42, F1: 128}
            _ = v30
        }
        if len(v2) > 0 {
            v2[0] = 935
        }
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        if (!false) {
            for fi32 := 0; fi32 < 3; fi32++ {
                v1 = 42
                _ = len(v2)
                _ = fi32
            }
            v1 -= (v1 - (v1 | (v1 + 1)))
        }
        fmt.Println(v1, fi31)
        _ = len(v2)
        _ = fi31
    }
    v1 -= (v1 * v1)
    fmt.Println(v3, v1)
    for fi33 := 0; fi33 < 2; fi33++ {
        v1 = 42
        _ = fi33
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        if true {
            v1 = 604
            for fi35 := 0; fi35 < 15; fi35++ {
                if len(v2) > 0 {
                    v2[0] = 858
                }
                v3 = (-2.25 - -2.25)
                _ = fi35
            }
        }
        _ = fi34
    }
    var v36 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
}
