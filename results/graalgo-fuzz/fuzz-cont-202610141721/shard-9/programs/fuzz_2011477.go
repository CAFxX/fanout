package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 string }

func main() {
    var v1 uint32 = 65535
    var v2 string = ("a" + "")
    var v3 uint16 = 37
    v4 := GenS1{F0: v3}
    if true {
        switch "\n\t" {
        case "foo bar":
            var v5 int = 0
            _ = v5
        default:
            v6 := []int{0}
            _ = v6
        }
        fmt.Println(v2)
    } else {
        fmt.Println(v3, v3, v1)
    }
    var v7 uint32 = (v1 + v1)
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint8 = 1
        v10 := []int{fi8, (fi8 << ((1) & 63)), (fi8 | (^(fi8))), 370}
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = (v7 | 2)
    v7 += 2147483647
    for fi11 := 0; fi11 < 11; fi11++ {
        if (true && (9223372036854775807 != 0)) {
            switch 1 {
            case -9223372036854775808:
                var v12 bool = (!(!(!false)))
                _ = v12
            }
        }
        var v13 int16 = 972
        _ = fi11
        _ = v13
    }
    fmt.Println(v7, v3, v2)
    v14 := GenS1{F0: uint16(v1)}
    if ("a" >= v2) {
        var v15 int32 = -184
        _ = v15
    }
    v3 = 32767
    switch int(v7) {
    case 9223372036854775807:
        var v16 float32 = (float32(3.14159) * 0.0)
        _ = v16
    case 686:
        v7 -= (v7 ^ (v1 * 4294967295))
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := GenS2{F0: 0, F1: "a"}
        v1 = v7
        _ = fi17
        _ = v18
    }
    if (true || false) {
        v1 += v1
    }
    if (32768 <= 0) {
        switch v2 {
        case "":
            if (true || (-9223372036854775808 == 0)) {
                v1 = v7
                fmt.Println(v7, v7, v7)
            } else {
                var v19 uint8 = 171
                if (!false) {
                    v2 = (v2 + (v2 + "\n\t"))
                    var v20 bool = false
                    switch -1 {
                    case -2147483648:
                        if (((v19 >> ((0) & 63)) <= v19) && ((!(false || v20)) || v20)) {
                            for fi21 := 0; fi21 < 18; fi21++ {
                                v19 = (^(v19))
                                v19 = 0
                                _ = fi21
                            }
                            if (4294967295 <= 4294967295) {
                                v1 = 42
                                var v22 int = 232
                                _ = v22
                            } else {
                                var v23 int64 = 659
                                v3 = (v3 + (v3 * 42))
                                _ = v23
                            }
                        }
                    case -647:
                        v7 = v1
                    case 1:
                        for fi24 := 0; fi24 < 17; fi24++ {
                            var v25 float64 = float64(1.5)
                            var v26 bool = (!(true || v20))
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = v20
                } else {
                    switch "foo bar" {
                    case "hello":
                        var v27 int64 = 0
                        _ = v27
                    }
                }
                _ = v19
            }
        case "\n\t":
            v28 := []int{32767, 4294967295}
            _ = v28
        case "a":
            var v29 int64 = -9223372036854775808
            _ = v29
        default:
            var v30 bool = true
            _ = v30
        }
    }
    if (true || false) {
        v31 := []int{127}
        v7 += v7
        var v32 uint16 = 630
        _ = v31
        _ = v32
    }
    v7 += 243
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
