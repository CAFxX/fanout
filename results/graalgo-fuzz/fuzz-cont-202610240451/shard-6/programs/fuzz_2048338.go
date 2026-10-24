package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int }
type GenS2 struct { F0 int32; F1 uint32; F2 uint }
type GenS3 struct { F0 string; F1 uint8; F2 uint32 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint64 }

func main() {
    var v1 bool = false
    var v2 uint16 = 581
    v3 := []int{-9223372036854775808, -1}
    v2 = v2
    v4 := GenS1{F0: 739, F1: -935, F2: 127}
    if true {
        var v5 uint = 4294967295
        _ = v5
    }
    v6 := GenS1{F0: 22, F1: -676, F2: 0}
    for fi7 := 0; fi7 < 1; fi7++ {
        v2 -= v2
        var v8 float32 = 1e10
        _ = fi7
        _ = v8
    }
    v1 = v1
    var v9 float32 = 1e10
    switch -1 {
    case 128:
        if (v1 || true) {
            var v10 string = "\n\t"
            _ = v10
        } else {
            v1 = (v1 && v1)
            fmt.Println(v1, v1)
            v1 = true
        }
    case 32767:
        for fi11 := 0; fi11 < 17; fi11++ {
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS2{F0: 42, F1: 0, F2: 18446744073709551615}
                var v14 uint8 = 139
                v14 += v14
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    case -165:
        var v15 string = "a"
        _ = v15
    default:
        fmt.Println(v2, v2, v1)
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        if v1 {
            switch int(v2) {
            case 552:
                fmt.Println(v9, v2, fi16)
            case 9223372036854775807:
                v1 = v1
            case 504:
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 int64 = 639
                    if len(v3) > 0 {
                        v3[0] = -2147483648
                    }
                    _ = fi17
                    _ = v18
                }
            }
            if ((-2.25 >= (3.14159 * ((-2.25 + 1e10) + 0.0))) && (v1 && (!v1))) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v9, v1)
            }
            v3 = append(v3, (fi16 / ((fi16) | 1)))
        } else {
            switch 4294967295 {
            case 2147483647:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                var v19 uint = 42
                _ = v19
            }
            for fi20 := 0; fi20 < 15; fi20++ {
                v3 = append(v3, fi20)
                if v1 {
                    if len(v3) > 0 {
                        v3[0] = (fi20 - (fi16 | fi20))
                    }
                }
                v21 := GenS3{F0: ("hello" + "a"), F1: 255, F2: 596}
                _ = fi20
                _ = v21
            }
        }
        v22 := GenS1{F0: 42, F1: -403, F2: fi16}
        var v23 uint = 4294967295
        _ = fi16
        _ = v22
        _ = v23
    }
    fmt.Println(v2, v2, v9)
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 float64 = 0.0
        v26 := GenS4{F0: "", F1: -128}
        _ = fi24
        _ = v25
        _ = v26
    }
    v2 += (v2 | v2)
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 5; fi28++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v1 = true
            _ = fi28
        }
        v29 := GenS2{F0: 1, F1: 204, F2: 65535}
        v30 := GenS5{F0: 4294967295, F1: 4294967295, F2: 915}
        _ = fi27
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
