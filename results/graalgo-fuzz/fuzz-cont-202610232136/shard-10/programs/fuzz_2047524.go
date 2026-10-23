package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 int32; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint) uint64 {
    var v4 int64 = -1
    v4 += (v4 ^ int64(p0))
    v5 := GenS1{F0: 65535, F1: 4294967295}
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 bool = (42 < 42)
    v2 := []string{}
    var v3 float32 = 3.14159
    if true {
        for fi6 := 0; fi6 < 19; fi6++ {
            if len(v2) > 0 {
                v2[0] = ("\n\t" + "\n\t")
            }
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := GenS1{F0: 622, F1: 65535}
                var v9 bool = ((!(41 <= 0)) || false)
                switch fi6 {
                case 32768:
                    switch (fi7 ^ fi7) {
                    case -9223372036854775808:
                        switch "foo bar" {
                        case "foo bar":
                            v1 = (77 >= 65535)
                        case "a":
                            var v10 uint8 = 128
                            _ = v10
                        }
                    default:
                        if v9 {
                            fmt.Println(fi7)
                        } else {
                            v9 = ((42 < 360) && (-180 < 1))
                            switch 803 {
                            case 128:
                                var v11 int32 = 42
                                _ = v11
                            case 618:
                                v1 = true
                            case -953:
                                v3 -= v3
                            default:
                                v1 = (42 == -9223372036854775808)
                            }
                            fmt.Println(fi6)
                        }
                    }
                default:
                    fmt.Println(v1)
                }
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
        v12 := GenS2{F0: "a", F1: 42, F2: 321}
        var v13 float32 = 0.0
        _ = v12
        _ = v13
    } else {
        var v14 uint64 = 9
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := GenS1{F0: 444, F1: 42}
            v1 = v1
            _ = fi15
            _ = v16
        }
        var v17 int32 = 562
        _ = v14
        _ = v17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        if v1 {
            v3 -= 0.0
        } else {
            v1 = (100 == -87)
            v19 := GenS1{F0: 180, F1: 18446744073709551615}
            switch (^(fi18)) {
            case -2147483648:
                v20 := GenS3{F0: v1, F1: 382, F2: 255}
                _ = v20
            case 9223372036854775807:
                var v21 uint16 = 965
                _ = v21
            default:
                var v22 uint32 = 2
                _ = v22
            }
            _ = v19
        }
        v1 = v1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi18
    }
    switch "foo bar" {
    case "\n\t":
        v23 := GenS1{F0: 65535, F1: 2}
        _ = v23
    }
    var v24 int32 = 161
    v25 := GenS4{F0: genFn0(1)}
    v3 -= v3
    for fi26 := 0; fi26 < 1; fi26++ {
        v24 -= v24
        v2 = append(v2, ("\n\t" + ("foo bar" + ("foo bar" + (("foo bar" + "") + "foo bar")))))
        v27 := GenS5{F0: v1}
        _ = fi26
        _ = v27
    }
    v28 := GenS4{F0: 577}
    var v29 float64 = 0.0
    var v30 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
