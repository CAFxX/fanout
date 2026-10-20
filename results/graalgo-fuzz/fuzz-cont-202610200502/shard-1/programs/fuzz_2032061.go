package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 string; F2 int8 }
func genFn0() uint32 {
    var v4 float64 = 1.5
    v5 := GenS1{F0: 275, F1: 1, F2: false}
    fmt.Println(v4)
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    var v1 float64 = 1e10
    v2 := []int{231, 128}
    var v3 uint16 = 32767
    for fi6 := 0; fi6 < 13; fi6++ {
        v1 += 0.1
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int8 = 127
            _ = fi7
            _ = v8
        }
        v3 += v3
        _ = fi6
    }
    if true {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 bool = false
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 uint8 = 127
                v2 = append(v2, fi9)
                _ = fi11
                _ = v12
            }
            if v10 {
                fmt.Println(v10, v3)
            } else {
                v13 := GenS2{F0: genFn0()}
                for fi14 := 0; fi14 < 20; fi14++ {
                    if (v3 != v3) {
                        v3 = (^(v3))
                        var v15 uint64 = 0
                        _ = v15
                    } else {
                        var v16 int16 = -1
                        _ = v16
                    }
                    _ = fi14
                }
                _ = v13
            }
            _ = fi9
            _ = v10
        }
        if (-1 != 591) {
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS1{F0: 896, F1: 42, F2: (fi17 < (fi17 - (fi17 >> ((255) & 63))))}
                fmt.Println(v3)
                v3 = (v3 & v3)
                _ = fi17
                _ = v18
            }
            if false {
                var v19 int32 = 42
                var v20 uint16 = v3
                _ = v19
                _ = v20
            }
        } else {
            var v21 int = -389
            v22 := GenS1{F0: 1, F1: 18446744073709551615, F2: false}
            for fi23 := 0; fi23 < 1; fi23++ {
                v3 = (v3 << ((419) & 63))
                switch (fi23 & fi23) {
                case 127:
                    fmt.Println(v3, v3, v21)
                }
                v21 = (fi23 ^ 4294967295)
                _ = fi23
            }
            _ = v21
            _ = v22
        }
    }
    fmt.Println(v1, v3)
    var v24 int32 = -915
    fmt.Println(v3)
    v24 -= (v24 / ((1) | 1))
    switch ("hello" + "foo bar") {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = -228
        }
    case "hello":
        var v25 uint32 = 834
        _ = v25
    case "":
        switch 9223372036854775807 {
        case -9223372036854775808:
            var v26 int16 = 32767
            _ = v26
        case 4294967295:
            v1 = 3.14159
        case 32767:
            v2 = append(v2, -1)
        default:
            v27 := GenS3{F0: 1, F1: ("" + "foo bar"), F2: 100}
            _ = v27
        }
    default:
        if (18446744073709551615 > 4294967295) {
            v24 += 2147483647
            for fi28 := 0; fi28 < 4; fi28++ {
                v29 := GenS3{F0: 32767, F1: "", F2: 100}
                _ = fi28
                _ = v29
            }
            v3 = v3
        }
    }
    v1 += 0.0
    switch 4294967295 {
    case -2147483648:
        v1 = v1
    case 32768:
        var v30 int = 42
        _ = v30
    }
    var v31 uint16 = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v32 int = 4294967295
    if len(v2) > 0 {
        v2[0] = v32
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v32
}
