package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 string = "a"
    var v2 uint32 = 417
    v3 := make(map[string]int)
    var v4 bool = (((1e10 * 1.5) == (1.5 * 0.1)) || true)
    var v5 uint8 = 2
    v3[v1] = 865
    var v6 float32 = ((0.1 + (((-2.25 + 1e10) / 1.5) + 3.14159)) * 3.14159)
    v7 := v3[(v1 + v1)]
    if v4 {
        var v8 int16 = -32768
        v9 := GenS1{F0: (v7 | v7), F1: 768, F2: v7}
        for fi10 := 0; fi10 < 13; fi10++ {
            for fi11 := 0; fi11 < 11; fi11++ {
                if false {
                    v12 := []int{}
                    v4 = ((528 != (v2 * v2)) && v4)
                    _ = v12
                }
                if v4 {
                    v2 -= v2
                    v13 := GenS1{F0: fi10, F1: 18446744073709551615, F2: fi11}
                    _ = v13
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        v14 := GenS2{F0: (v8 & (^(v8))), F1: (v4 && (!v4))}
                        _ = v14
                    default:
                        if v4 {
                            v15 := GenS3{F0: (v8 + v8), F1: v4}
                            v7 = int(v8)
                            _ = v15
                        }
                    }
                    var v16 uint8 = (v5 / (((v5 + v5)) | 1))
                    var v17 int64 = 42
                    v6 = v6
                    v5 = v16
                    v3["a"] = (fi10 ^ v7)
                    fmt.Println(v4, v6, fi10)
                    if v4 {
                        var v18 string = v1
                        switch (fi11 % ((v7) | 1)) {
                        case 2147483647:
                            if v4 {
                                fmt.Println(v4)
                            }
                        }
                        _ = v18
                    } else {
                        var v19 int16 = int16(v2)
                        var v20 int32 = -1
                        _ = v19
                        _ = v20
                    }
                    v5 += v5
                    _ = v16
                    _ = v17
                }
                v6 += 1e10
                _ = fi11
            }
            switch v7 {
            case -9223372036854775808:
                switch (((v1 + "hello") + (v1 + ((("a" + v1) + ((("hello" + v1) + ("hello" + ((v1 + v1) + ""))) + ("hello" + v1))) + (v1 + "a")))) + v1) {
                case "a":
                    switch (fi10 + v7) {
                    case -9223372036854775808:
                        if (100 >= 1) {
                            var v21 uint64 = 764
                            v3[v1] = 32767
                            _ = v21
                        }
                    }
                case "hello":
                    for fi22 := 0; fi22 < 11; fi22++ {
                        fmt.Println(v4, v8)
                        if (!v4) {
                            var v23 string = v1
                            v5 = (v5 ^ (v5 << ((42) & 63)))
                            v24 := v3[v1]
                            v3[((("foo bar" + v23) + v23) + v1)] = (fi22 / ((v7) | 1))
                            var v25 float64 = ((0.1 - ((1e10 / (0.0 + -2.25)) / 1e10)) * (1.5 / (3.14159 * 1e10)))
                            v6 -= v6
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        v5 = v5
                        _ = fi22
                    }
                case "\n\t":
                    v26 := GenS3{F0: v8, F1: v4}
                    _ = v26
                }
            case -752:
                v27 := GenS4{F0: (v5 + v5), F1: (v5 >> ((0) & 63))}
                _ = v27
            case 32768:
                v28 := []int{}
                _ = v28
            default:
                if v4 {
                    switch (v7 ^ v7) {
                    case -9223372036854775808:
                        for fi29 := 0; fi29 < 3; fi29++ {
                            fmt.Println(v2, v2)
                            switch (v1 + v1) {
                            case "\n\t":
                                v30 := GenS5{F0: uint8(v2)}
                                switch v1 {
                                case "\n\t":
                                    v1 = v1
                                case "":
                                    switch ((v1 + ("a" + ((v1 + (((v1 + (("foo bar" + "a") + "a")) + v1) + "")) + ((((v1 + "hello") + "foo bar") + (v1 + "foo bar")) + (v1 + (v1 + v1)))))) + v1) {
                                    case "a":
                                        if (v1 != ((("a" + v1) + v1) + v1)) {
                                            switch "a" {
                                            case "\n\t":
                                                var v31 int8 = -62
                                                _ = v31
                                            case "hello":
                                                var v32 uint8 = v5
                                                _ = v32
                                            default:
                                                v33 := GenS6{F0: 32767, F1: (v2 * (v2 & (v2 << ((278) & 63))))}
                                                _ = v33
                                            }
                                        }
                                    case "\n\t":
                                        var v34 float64 = 3.14159
                                        _ = v34
                                    }
                                case "foo bar":
                                    var v35 uint32 = v2
                                    _ = v35
                                }
                                fmt.Println(v5, v4, fi10)
                                _ = v30
                            case "":
                                if (!(42 < 748)) {
                                    if true {
                                        v8 = (v8 << ((uint(v2)) & 63))
                                        if (0 >= 2) {
                                            v36 := []int{(fi10 & v7), fi10}
                                            if (!(((!((v2 << ((410) & 63)) >= 2147483647)) && ("\n\t" > "\n\t")) || false)) {
                                                _ = len(v3)
                                                if len(v36) >= 2 {
                                                    v36 = v36[1:len(v36)]
                                                }
                                            }
                                            _ = v36
                                        }
                                        var v37 uint64 = 4294967295
                                        switch -2147483648 {
                                        case 128:
                                            _ = len(v3)
                                        case 1:
                                            if (v7 == (v7 & (-(v7)))) {
                                                v2 = (v2 << ((647) & 63))
                                                v2 = (v2 - (v2 << ((830) & 63)))
                                            } else {
                                                switch (^(fi29)) {
                                                case 217:
                                                    var v38 bool = (v4 && (-1 == 1))
                                                    _ = v38
                                                case 32768:
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v39 float32 = float32(0.1)
                                                        _ = v39
                                                    }
                                                }
                                            }
                                        case 127:
                                            v40 := GenS4{F0: 114, F1: (^(v5))}
                                            _ = v40
                                        default:
                                            switch fi29 {
                                            case 2147483647:
                                                v6 = 1.5
                                            default:
                                                v41 := v3["\n\t"]
                                                _ = v41
                                            }
                                        }
                                        v42 := []int{}
                                        _ = v37
                                        _ = v42
                                    } else {
                                        v2 -= (v2 | 4294967295)
                                    }
                                }
                            case "hello":
                                v43 := GenS1{F0: fi29, F1: 65535, F2: fi29}
                                _ = v43
                            }
                            _ = fi29
                        }
                    default:
                        var v44 int = v7
                        _ = v44
                    }
                } else {
                    if (!v4) {
                        var v45 uint32 = v2
                        for fi46 := 0; fi46 < 13; fi46++ {
                            var v47 int8 = 42
                            v48 := GenS3{F0: (v8 % ((v8) | 1)), F1: (!true)}
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                        _ = v45
                    }
                }
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    v4 = (!((true || v4) && false))
    v4 = v4
    var v49 int16 = -1
    for fi50 := 0; fi50 < 16; fi50++ {
        v49 -= v49
        _ = fi50
    }
    _ = len(v3)
    v51 := GenS2{F0: (v49 >> ((18446744073709551615) & 63)), F1: true}
    v52 := GenS4{F0: (v5 * (v5 & 178)), F1: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v51
    _ = v52
    _ = v6
    _ = v7
}
