package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 bool; F1 int; F2 uint16 }
func genFn0() int64 {
    var v4 int = 32768
    v5 := GenS1{F0: 1}
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 int16 = 510
    v2 := []int{}
    v3 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 6; fi6++ {
        switch "foo bar" {
        case "":
            if (42 <= 9223372036854775807) {
                fmt.Println(fi6, fi6)
            } else {
                var v7 int64 = -1
                v7 += 9223372036854775807
                if true {
                    v7 = v7
                    if (!false) {
                        for fi8 := 0; fi8 < 3; fi8++ {
                            var v9 bool = true
                            _ = fi8
                            _ = v9
                        }
                    }
                    v1 += (v1 << ((255) & 63))
                }
                _ = v7
            }
        default:
            for fi10 := 0; fi10 < 10; fi10++ {
                v3[0] = true
                v11 := GenS1{F0: 4294967295}
                var v12 uint32 = 2147483647
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = fi6
    }
    var v13 int16 = v1
    v14 := GenS2{F0: 18446744073709551615, F1: 255}
    fmt.Println(v1, v1)
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 float32 = -2.25
        _ = fi15
        _ = v16
    }
    if true {
        v1 = 42
    }
    fmt.Println(v13)
    for fi17 := 0; fi17 < 3; fi17++ {
        if len(v2) > 0 {
            v2[0] = 9223372036854775807
        }
        _ = fi17
    }
    var v18 int64 = genFn0()
    v13 = 0
    fmt.Println(v13, v18, v13)
    var v19 uint8 = 121
    switch -623 {
    case 1:
        switch 42 {
        case 4294967295:
            var v20 uint8 = (v19 | v19)
            _ = v20
        case 32767:
            for fi21 := 0; fi21 < 14; fi21++ {
                if false {
                    v22 := GenS3{F0: 127}
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v13 += v1
                        v24 := GenS1{F0: 216}
                        if false {
                            if false {
                                fmt.Println(v13, v18)
                                v1 += v13
                            }
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                v18 = (v18 ^ v18)
                            case "a":
                                if (("hello" + "hello") <= "\n\t") {
                                    v25 := GenS3{F0: (v19 + v19)}
                                    _ = v25
                                }
                            case "hello":
                                v1 += v13
                            }
                            var v26 float32 = float32(0.0)
                            var v27 uint16 = 42
                            var v28 float64 = (1.5 - (1e10 * 1.5))
                            v28 += v28
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        _ = fi23
                        _ = v24
                    }
                    v19 += 2
                    _ = v22
                } else {
                    v29 := GenS3{F0: v19}
                    for fi30 := 0; fi30 < 18; fi30++ {
                        v31 := GenS1{F0: 2}
                        switch ("\n\t" + "") {
                        case "\n\t":
                            v13 = v13
                        }
                        _ = fi30
                        _ = v31
                    }
                    _ = v29
                }
                if (-290 <= -23) {
                    if ((2147483647 == 0) && false) {
                        for fi32 := 0; fi32 < 18; fi32++ {
                            v33 := GenS4{F0: 2, F1: 42}
                            fmt.Println(v18)
                            _ = fi32
                            _ = v33
                        }
                    }
                    var v34 int = -54
                    _ = v34
                } else {
                    var v35 int16 = (v1 * (v1 << ((1) & 63)))
                    v13 = v35
                    _ = v35
                }
                _ = fi21
            }
        }
    case 1000:
        v36 := GenS5{F0: true, F1: 109, F2: 604}
        _ = v36
    case 2147483647:
        v37 := v3[333]
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
