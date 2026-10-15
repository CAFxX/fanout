package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    v3 := []string{("foo bar" + ("\n\t" + "\n\t")), ("hello" + ("hello" + ("a" + "foo bar"))), "a"}
    var v4 float64 = float64(3.14159)
    switch "\n\t" {
    case "a":
        for fi5 := 0; fi5 < 3; fi5++ {
            fmt.Println(fi5)
            fmt.Println(v4)
            _ = fi5
        }
    case "hello":
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 uint16 = 559
            _ = fi6
            _ = v7
        }
    }
    if false {
        fmt.Println(v4)
        v8 := GenS1{F0: 2147483647}
        _ = v8
    } else {
        fmt.Println(v4)
        v9 := GenS1{F0: 42}
        _ = v9
    }
    v10 := GenS2{F0: 1, F1: 216, F2: 0}
    var v11 uint64 = 2
    for fi12 := 0; fi12 < 9; fi12++ {
        switch "" {
        case "hello":
            for fi13 := 0; fi13 < 3; fi13++ {
                v11 = v11
                if false {
                    v14 := GenS2{F0: -359, F1: -2147483648, F2: v11}
                    if true {
                        v4 -= 1e10
                        var v15 uint64 = (v11 * 743)
                        _ = v15
                    }
                    switch fi13 {
                    case 32768:
                        fmt.Println(v4, v4)
                    case 2147483647:
                        switch (^(fi13)) {
                        case 43:
                            if false {
                                v16 := GenS1{F0: 0}
                                var v17 uint16 = 643
                                var v18 int64 = -9223372036854775808
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        case -1:
                            v4 += (0.1 + ((v4 / (((v4 / v4) - 3.14159) + float64(0.0))) / (float64(1.5) * ((v4 - v4) * (((v4 - v4) * v4) * (v4 * (((v4 * (1.5 + (v4 - v4))) * 1.5) - float64((-2.25 * 0.0)))))))))
                        case 32767:
                            var v19 float64 = ((v4 / v4) / v4)
                            _ = v19
                        }
                    case 1:
                        var v20 uint32 = 689
                        _ = v20
                    default:
                        fmt.Println(fi13, v11)
                    }
                    _ = v14
                }
                _ = fi13
            }
        default:
            switch (fi12 * fi12) {
            case 4294967295:
                var v21 int16 = 1
                _ = v21
            case -9223372036854775808:
                var v22 string = "hello"
                _ = v22
            case 765:
                switch (-(fi12)) {
                case -2147483648:
                    var v23 int8 = 0
                    _ = v23
                case -9223372036854775808:
                    for fi24 := 0; fi24 < 2; fi24++ {
                        v4 -= ((v4 * v4) / float64(-2.25))
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v11 -= v11
                        } else {
                            v11 += v11
                            v3 = append(v3, "hello")
                            var v25 int = 42
                            _ = v25
                        }
                        if true {
                            _ = len(v3)
                            v26 := GenS3{F0: 127}
                            _ = v26
                        }
                        _ = fi24
                    }
                default:
                    switch "" {
                    case "a":
                        for fi27 := 0; fi27 < 13; fi27++ {
                            fmt.Println(v11, fi27)
                            if false {
                                if (false && true) {
                                    v4 = v4
                                    if true {
                                        v28 := GenS4{F0: 0, F1: 42, F2: 2}
                                        v11 += v11
                                        _ = v28
                                    }
                                }
                                var v29 int64 = 9223372036854775807
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v29
                            } else {
                                v30 := GenS1{F0: 658}
                                if false {
                                    v4 = (float64((3.14159 / 1e10)) - 1.5)
                                    var v31 uint32 = 4294967295
                                    _ = v31
                                } else {
                                    v4 -= 0.1
                                }
                                switch 9223372036854775807 {
                                case 42:
                                    var v32 int16 = -32768
                                    _ = v32
                                case 9223372036854775807:
                                    v11 -= v11
                                default:
                                    if false {
                                        v33 := GenS5{F0: 613}
                                        _ = v33
                                    } else {
                                        var v34 uint64 = v11
                                        _ = v34
                                    }
                                }
                                v4 = -2.25
                                var v35 int = fi12
                                var v36 int32 = -1
                                _ = v30
                                _ = v35
                                _ = v36
                            }
                            var v37 float64 = v4
                            _ = fi27
                            _ = v37
                        }
                    case "hello":
                        fmt.Println(v11, fi12)
                    case "\n\t":
                        v38 := GenS1{F0: 2}
                        _ = v38
                    }
                }
            default:
                var v39 float64 = float64(((3.14159 - (1e10 - 0.1)) - (3.14159 * (((0.1 - float32(float64(float32((0.0 / float64((-2.25 / 3.14159))))))) * 1.5) / (3.14159 / 0.1)))))
                _ = v39
            }
        }
        _ = fi12
    }
    if false {
        v40 := GenS3{F0: 2}
        if true {
            var v41 bool = ((-9223372036854775808 != 808) && (0.1 == (1e10 * -2.25)))
            _ = v41
        }
        _ = v40
    } else {
        fmt.Println(v11, v11)
        v11 = v11
    }
    _ = len(v1)
    v11 = v11
    var v42 uint64 = (v11 >> ((1) & 63))
    if true {
        fmt.Println(v11)
        switch 0 {
        case 32767:
            v4 = ((v4 - v4) * v4)
        case -715:
            v43 := GenS6{F0: ""}
            _ = v43
        }
    } else {
        for fi44 := 0; fi44 < 2; fi44++ {
            v42 += v42
            _ = fi44
        }
        if true {
            v42 = 285
        } else {
            v11 = v11
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
