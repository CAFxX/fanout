package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint64 }

func main() {
    var v1 int16 = 492
    v2 := make(map[string]uint32)
    var v3 uint32 = 505
    var v4 uint32 = (v3 | 227)
    var v5 int64 = -983
    v3 = v3
    delete(v2, "")
    var v6 uint64 = 4294967295
    v7 := v2["\n\t"]
    var v8 uint64 = v6
    v8 = v6
    v8 += (v8 | v8)
    var v9 uint16 = 1
    switch 128 {
    case -185:
        if (65535 > 1) {
            switch ("foo bar" + ("foo bar" + "a")) {
            case "a":
                switch 0 {
                case -26:
                    if false {
                        v10 := []int{9223372036854775807}
                        _ = v10
                    }
                case -2147483648:
                    if true {
                        switch 9223372036854775807 {
                        case 2147483647:
                            if (64 <= 2) {
                                if true {
                                    fmt.Println(v1)
                                    fmt.Println(v8, v3)
                                } else {
                                    for fi11 := 0; fi11 < 2; fi11++ {
                                        if (!false) {
                                            for fi12 := 0; fi12 < 16; fi12++ {
                                                v13 := v2[("a" + "")]
                                                var v14 int = fi11
                                                _ = fi12
                                                _ = v13
                                                _ = v14
                                            }
                                        }
                                        if false {
                                            v4 -= (^(v4))
                                            for fi15 := 0; fi15 < 6; fi15++ {
                                                v3 = (v4 & uint32(v6))
                                                var v16 bool = (((float64(3.14159) + ((-2.25 - 1e10) + 0.0)) != float64(((((1.5 * (((-2.25 + float32(0.0)) - float32(((-2.25 + -2.25) / 1.0))) * 0.1)) + -2.25) * 0.0) + (0.1 + (float32(3.14159) - 1.5))))) || ((0.0 * 1.5) >= (1e10 / 0.1)))
                                                var v17 int8 = 127
                                                _ = fi15
                                                _ = v16
                                                _ = v17
                                            }
                                            var v18 bool = (!((0 >= v9) || ((!(true && (42 < uint(v6)))) && (true && false))))
                                            _ = v18
                                        }
                                        v19 := GenS1{F0: (v5 | v5), F1: 18446744073709551615, F2: 1}
                                        _ = fi11
                                        _ = v19
                                    }
                                    v20 := GenS1{F0: (v5 & (v5 << ((409) & 63))), F1: (^(v6)), F2: 0}
                                    v7 -= (^(v7))
                                    _ = v20
                                }
                                for fi21 := 0; fi21 < 14; fi21++ {
                                    v7 = (^(v4))
                                    v22 := GenS2{F0: (v9 / ((v9) | 1)), F1: v5, F2: v1}
                                    _ = fi21
                                    _ = v22
                                }
                            }
                        }
                        delete(v2, "hello")
                    }
                }
            case "\n\t":
                for fi23 := 0; fi23 < 3; fi23++ {
                    v24 := GenS3{F0: v1}
                    v25 := GenS3{F0: (v1 & (v1 % ((v1) | 1)))}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            case "hello":
                for fi26 := 0; fi26 < 5; fi26++ {
                    _ = len(v2)
                    v1 += v1
                    v27 := GenS4{F0: 18446744073709551615}
                    _ = fi26
                    _ = v27
                }
            default:
                if true {
                    v1 = v1
                    fmt.Println(v1)
                } else {
                    var v28 uint8 = 127
                    if false {
                        var v29 int64 = 0
                        v3 += (v3 & v7)
                        v30 := GenS5{F0: false}
                        _ = v29
                        _ = v30
                    }
                    _ = v28
                }
            }
        } else {
            switch ((("foo bar" + "a") + "") + "foo bar") {
            case "\n\t":
                for fi31 := 0; fi31 < 20; fi31++ {
                    for fi32 := 0; fi32 < 8; fi32++ {
                        if (!false) {
                            v33 := GenS2{F0: (v9 / ((v9) | 1)), F1: -1, F2: v1}
                            fmt.Println(v1, v3, v1)
                            var v34 uint32 = (v7 << ((255) & 63))
                            _ = v33
                            _ = v34
                        }
                        var v35 uint32 = v4
                        var v36 int64 = v5
                        _ = fi32
                        _ = v35
                        _ = v36
                    }
                    v37 := GenS3{F0: v1}
                    if (true && true) {
                        if true {
                            switch "" {
                            case "":
                                v38 := []int{642, fi31, (fi31 % (((fi31 + 9223372036854775807)) | 1)), -1}
                                _ = v38
                            case "a":
                                v3 = v7
                            case "foo bar":
                                v1 -= (v1 + v1)
                            }
                            for fi39 := 0; fi39 < 17; fi39++ {
                                v2["hello"] = (v4 + uint32(v9))
                                switch (fi39 & int(v3)) {
                                case -9223372036854775808:
                                    fmt.Println(v7, v3, fi39)
                                }
                                var v40 float64 = (((1.5 + ((float64(1.5) - 1e10) * (1.5 * 3.14159))) - 0.1) * 0.0)
                                _ = fi39
                                _ = v40
                            }
                            v41 := []int{(fi31 >> ((65535) & 63)), -2147483648, (-(fi31)), fi31}
                            _ = v41
                        }
                    } else {
                        for fi42 := 0; fi42 < 4; fi42++ {
                            if false {
                                var v43 string = (("hello" + "hello") + "")
                                _ = len(v2)
                                fmt.Println(v6, v1)
                                _ = v43
                            }
                            var v44 float32 = 1e10
                            _ = fi42
                            _ = v44
                        }
                    }
                    _ = fi31
                    _ = v37
                }
            }
            v45 := v2["hello"]
            v1 += (v1 >> ((0) & 63))
            _ = v45
        }
    case 714:
        var v46 int32 = 42
        _ = v46
    default:
        for fi47 := 0; fi47 < 17; fi47++ {
            v48 := GenS6{F0: -1}
            _ = fi47
            _ = v48
        }
    }
    for fi49 := 0; fi49 < 4; fi49++ {
        v4 = 1
        var v50 int8 = 42
        v1 = (v1 / (((v1 >> ((743) & 63))) | 1))
        _ = fi49
        _ = v50
    }
    var v51 int64 = int64(v9)
    for fi52 := 0; fi52 < 1; fi52++ {
        fmt.Println(v9, v51)
        v53 := GenS7{F0: false, F1: 127}
        _ = fi52
        _ = v53
    }
    fmt.Println(v8, v4, v4)
    for fi54 := 0; fi54 < 17; fi54++ {
        fmt.Println(v1, v8)
        _ = fi54
    }
    v55 := GenS8{F0: (v5 >> ((311) & 63)), F1: v8}
    var v56 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v55
    _ = v56
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
