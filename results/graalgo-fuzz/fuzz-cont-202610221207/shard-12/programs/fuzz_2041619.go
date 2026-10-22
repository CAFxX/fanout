package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int8; F2 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 string = (("a" + "foo bar") + ("foo bar" + "hello"))
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    var v5 int = 9223372036854775807
    v6 := make(map[int]bool)
    v3[(v2 + (v2 + v2))] = (-(v5))
    var v7 int64 = -1
    switch 2147483647 {
    case -1:
        v8 := GenS1{F0: 18446744073709551615, F1: 10}
        _ = v8
    case 1:
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int16 = 1
            _ = fi9
            _ = v10
        }
    case 4294967295:
        v5 += (v5 | (v5 % ((1) | 1)))
    default:
        var v11 int32 = 42
        _ = v11
    }
    switch -2147483648 {
    case 32767:
        if (2 < 42) {
            v2 = v2
        }
    case 128:
        v12 := GenS2{F0: 1, F1: false, F2: 127}
        _ = v12
    case -9223372036854775808:
        v5 = v5
    default:
        if (255 <= 4294967295) {
            for fi13 := 0; fi13 < 18; fi13++ {
                for fi14 := 0; fi14 < 9; fi14++ {
                    v2 = ((v2 + "a") + v2)
                    _ = fi14
                }
                switch 763 {
                case -796:
                    var v15 uint32 = 4294967295
                    _ = v15
                default:
                    v2 = ("a" + v2)
                }
                _ = fi13
            }
        }
    }
    switch "" {
    case "hello":
        for fi16 := 0; fi16 < 10; fi16++ {
            switch v5 {
            case 128:
                var v17 int8 = -35
                _ = v17
            case -802:
                fmt.Println(fi16, fi16)
            case 792:
                fmt.Println(v5, v7)
            default:
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 bool = (976 == 2)
                    if (4294967295 != 1) {
                        var v20 uint32 = 432
                        _ = v20
                    } else {
                        var v21 bool = false
                        _ = v21
                    }
                    v22 := []int{2147483647, fi16, fi16, (fi16 ^ 9223372036854775807)}
                    if (2 <= 65535) {
                        var v23 int32 = -2147483648
                        _ = v23
                    } else {
                        if (!v19) {
                            v24 := v4[981]
                            _ = v24
                        }
                        var v25 int = fi16
                        v5 -= 4294967295
                        _ = v25
                    }
                    v5 = 2147483647
                    if v19 {
                        if (!v19) {
                            switch (v5 / ((-9223372036854775808) | 1)) {
                            case 0:
                                var v26 float64 = (0.1 * ((1e10 - (1.5 * ((0.0 * float64(-2.25)) + (1.5 - (((((1.5 - (float64(((1.5 * 1.5) - 1e10)) + 0.1)) * (0.0 / -2.25)) / 1e10) * ((1e10 + 0.1) + 1e10)) - -2.25))))) * (3.14159 - 1.5)))
                                _ = v26
                            case 9223372036854775807:
                                var v27 int8 = 107
                                _ = v27
                            case 127:
                                _ = len(v1)
                            }
                            v7 = (-(v7))
                            var v28 int16 = 42
                            _ = v28
                        } else {
                            v5 = 32767
                        }
                        v19 = false
                        v29 := GenS3{F0: (v7 | v7), F1: (v7 / ((v7) | 1))}
                        _ = v29
                    } else {
                        v30 := GenS1{F0: 18446744073709551615, F1: 42}
                        var v31 float32 = 3.14159
                        if true {
                            v19 = v19
                            switch (v2 + "foo bar") {
                            case "hello":
                                switch "a" {
                                case "a":
                                    v7 += -885
                                default:
                                    fmt.Println(v2, fi18, v31)
                                }
                            case "\n\t":
                                fmt.Println(fi18, v2, v7)
                            default:
                                v2 = v2
                            }
                            switch v2 {
                            case "a":
                                v19 = (42 >= -76)
                            case "foo bar":
                                if (-128 == -128) {
                                    fmt.Println(v19)
                                }
                            }
                        } else {
                            switch (("foo bar" + ((v2 + v2) + "")) + (((v2 + v2) + v2) + v2)) {
                            case "foo bar":
                                if false {
                                    v1[-1] = (v2 + v2)
                                    fmt.Println(v19, v7, v7)
                                }
                            }
                            var v32 float32 = 0.1
                            switch ("hello" + v2) {
                            case "hello":
                                fmt.Println(v31)
                            case "\n\t":
                                var v33 bool = ((!v19) && (42 > 208))
                                _ = v33
                            case "foo bar":
                                switch v2 {
                                case "a":
                                    var v34 int32 = -917
                                    _ = v34
                                case "\n\t":
                                    v35 := GenS1{F0: 42, F1: 128}
                                    _ = v35
                                case "foo bar":
                                    v31 -= v31
                                }
                            }
                            switch "" {
                            case "a":
                                v2 = "a"
                            case "":
                                v36 := GenS4{F0: 240}
                                _ = v36
                            default:
                                v37 := v4[814]
                                _ = v37
                            }
                            _ = v32
                        }
                        _ = v30
                        _ = v31
                    }
                    v38 := GenS5{F0: fi18, F1: 100, F2: 198}
                    _ = fi18
                    _ = v19
                    _ = v22
                    _ = v38
                }
            }
            for fi39 := 0; fi39 < 20; fi39++ {
                v40 := GenS3{F0: 0, F1: (v7 >> ((326) & 63))}
                _ = fi39
                _ = v40
            }
            _ = fi16
        }
    default:
        v2 = ("hello" + v2)
    }
    var v41 uint64 = 0
    var v42 int32 = 2147483647
    _ = len(v6)
    if true {
        fmt.Println(v42)
    }
    var v43 uint64 = 18446744073709551615
    v41 -= (v43 * v43)
    if (65535 >= 607) {
        fmt.Println(v2, v7, v41)
    } else {
        v7 += (^(v7))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
