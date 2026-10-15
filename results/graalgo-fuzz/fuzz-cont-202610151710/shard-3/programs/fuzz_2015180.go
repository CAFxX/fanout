package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 string; F2 int16 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 string; F1 uint32 }
func genFn0() uint64 {
    if false {
        var v7 int32 = 1
        switch -303 {
        case 389:
            v7 = 0
        case -873:
            var v8 uint = 848
            _ = v8
        case 32767:
            fmt.Println(v7)
        default:
            var v9 uint8 = 50
            _ = v9
        }
        _ = v7
    }
    if (true || (!false)) {
        v10 := GenS1{F0: 2, F1: 0}
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS2{F0: 42}
        if (9223372036854775807 < -1) {
            var v13 int = (^(fi11))
            var v14 int16 = 0
            var v15 uint16 = 272
            _ = v13
            _ = v14
            _ = v15
        }
        fmt.Println(fi11)
        _ = fi11
        _ = v12
    }
    var v16 int64 = -1
    var v17 float64 = -2.25
    _ = v16
    _ = v17
    return 2
}

func main() {
    v1 := []string{"", "foo bar"}
    var v2 uint32 = 867
    v3 := make(map[uint32]string)
    var v4 float64 = ((-2.25 + 1e10) + 3.14159)
    var v5 string = "a"
    var v6 uint16 = 717
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 uint8 = 0
        _ = fi18
        _ = v19
    }
    if (18446744073709551615 < 155) {
        v2 += (^(v2))
    }
    v4 -= v4
    v20 := GenS3{F0: 65535, F1: "hello", F2: -976}
    for fi21 := 0; fi21 < 5; fi21++ {
        v5 = v5
        if len(v1) > 0 {
            v1[0] = v5
        }
        var v22 bool = true
        _ = fi21
        _ = v22
    }
    v6 = v6
    fmt.Println(v2, v4, v6)
    if (614 > 4294967295) {
        if (0 != 999) {
            for fi23 := 0; fi23 < 10; fi23++ {
                v5 = (v5 + v5)
                fmt.Println(v4, v5, v4)
                _ = fi23
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                fmt.Println(v2)
                var v25 uint64 = 0
                _ = fi24
                _ = v25
            }
            v26 := GenS3{F0: v2, F1: "foo bar", F2: 1}
            _ = v26
        }
        v4 += v4
        var v27 float64 = v4
        _ = v27
    } else {
        v4 = v4
        v28 := GenS4{F0: ((((v4 != (3.14159 / (v4 / v4))) || (!true)) || ((0.0 + 3.14159) != ((v4 * 3.14159) - v4))) && true), F1: 32767}
        var v29 float64 = v4
        _ = v28
        _ = v29
    }
    if true {
        var v30 uint8 = 0
        fmt.Println(v4)
        _ = v30
    }
    var v31 float32 = 1e10
    if (!(((!true) || (1 == -9223372036854775808)) && true)) {
        v32 := v3[v2]
        v2 -= (^(v2))
        switch 42 {
        case -823:
            if (0 != -38) {
                switch 279 {
                case 32767:
                    if false {
                        fmt.Println(v5, v4)
                    }
                }
                for fi33 := 0; fi33 < 2; fi33++ {
                    v34 := GenS5{F0: 42, F1: (v32 + "a")}
                    if false {
                        v35 := GenS2{F0: 9223372036854775807}
                        switch (fi33 * fi33) {
                        case 0:
                            var v36 int = fi33
                            _ = v36
                        case 54:
                            v2 = v2
                        case -2147483648:
                            var v37 uint16 = (^(v6))
                            _ = v37
                        default:
                            switch v5 {
                            case "\n\t":
                                delete(v3, v2)
                            default:
                                _ = len(v1)
                            }
                        }
                        var v38 int16 = -1
                        _ = v35
                        _ = v38
                    } else {
                        v39 := GenS4{F0: false, F1: -32768}
                        _ = v39
                    }
                    v31 = (v31 * -2.25)
                    _ = fi33
                    _ = v34
                }
                for fi40 := 0; fi40 < 12; fi40++ {
                    v6 -= (v6 / ((90) | 1))
                    _ = fi40
                }
            } else {
                switch 32768 {
                case 881:
                    v2 -= (v2 | 65535)
                case 809:
                    v31 += v31
                case -693:
                    switch 128 {
                    case 4294967295:
                        _ = len(v3)
                    case 32768:
                        var v41 float64 = v4
                        _ = v41
                    }
                default:
                    v42 := GenS6{F0: v5, F1: (v2 / (((v2 * (^(v2)))) | 1))}
                    _ = v42
                }
            }
        default:
            var v43 bool = true
            _ = v43
        }
        _ = v32
    } else {
        if false {
            _ = len(v3)
        }
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            v45 := GenS6{F0: "\n\t", F1: v2}
            fmt.Println(v2)
            _ = v45
        }
        fmt.Println(v31, fi44, v2)
        _ = fi44
    }
    var v46 uint16 = (v6 % ((v6) | 1))
    _ = len(v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
