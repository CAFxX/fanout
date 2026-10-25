package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int8; F1 int; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]uint32)
    var v3 string = "hello"
    var v4 int32 = 2147483647
    v5 := GenS1{F0: 42, F1: 100, F2: 591}
    if (1.5 == -2.25) {
        v3 = v3
        v6 := GenS2{F0: 9223372036854775807}
        _ = v6
    }
    var v7 float32 = 3.14159
    switch v3 {
    case "":
        fmt.Println(v3)
    case "\n\t":
        var v8 bool = false
        _ = v8
    case "hello":
        switch "hello" {
        case "\n\t":
            var v9 float64 = 0.0
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 6; fi10++ {
            if true {
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 uint16 = 32767
                    v12 -= 803
                    _ = fi11
                    _ = v12
                }
            } else {
                fmt.Println(v4)
                v4 -= v4
            }
            _ = fi10
        }
    }
    switch "hello" {
    case "\n\t":
        switch 2147483647 {
        case 331:
            switch 895 {
            case -688:
                _ = len(v2)
            case 32768:
                for fi13 := 0; fi13 < 1; fi13++ {
                    v14 := GenS1{F0: 124, F1: 100, F2: 18446744073709551615}
                    v3 = (("" + (((v3 + v3) + "") + "")) + "a")
                    if ((true || (708 <= -1)) && false) {
                        v3 = v3
                    } else {
                        var v15 int16 = 32767
                        v15 -= (v15 & (^(v15)))
                        for fi16 := 0; fi16 < 18; fi16++ {
                            if (42 < 4294967295) {
                                v17 := GenS2{F0: 683}
                                _ = v17
                            } else {
                                var v18 uint16 = 42
                                v19 := GenS1{F0: (v4 >> ((42) & 63)), F1: 58, F2: uint(v18)}
                                _ = len(v1)
                                _ = v18
                                _ = v19
                            }
                            if true {
                                var v20 string = v3
                                var v21 uint = 413
                                _ = v20
                                _ = v21
                            }
                            _ = fi16
                        }
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
            case -9223372036854775808:
                for fi22 := 0; fi22 < 19; fi22++ {
                    v23 := GenS3{F0: 1, F1: 0}
                    if true {
                        if true {
                            switch v3 {
                            case "a":
                                var v24 int = fi22
                                _ = v24
                            case "foo bar":
                                if true {
                                    v7 = 0.0
                                } else {
                                    for fi25 := 0; fi25 < 5; fi25++ {
                                        v26 := GenS2{F0: -1}
                                        fmt.Println(v7, v4)
                                        _ = fi25
                                        _ = v26
                                    }
                                    v4 = 1
                                }
                            case "":
                                for fi27 := 0; fi27 < 6; fi27++ {
                                    v28 := GenS4{F0: 127, F1: -9223372036854775808, F2: 379}
                                    var v29 string = v3
                                    v30 := []int{fi22, -193}
                                    var v31 float64 = ((-2.25 - (0.0 * -2.25)) * 0.0)
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                            default:
                                v32 := GenS1{F0: (v4 / (((v4 / (((v4 / ((42) | 1))) | 1))) | 1)), F1: -128, F2: 2}
                                _ = v32
                            }
                        } else {
                            for fi33 := 0; fi33 < 18; fi33++ {
                                v34 := GenS5{F0: 0, F1: 42, F2: 1}
                                v35 := GenS3{F0: 1, F1: 970}
                                switch fi22 {
                                case 0:
                                    v3 = "\n\t"
                                }
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        }
                        if (!false) {
                            v7 = v7
                            var v36 int = (fi22 / ((-9223372036854775808) | 1))
                            _ = v36
                        }
                        switch fi22 {
                        case 4294967295:
                            if true {
                                if (0 == 0) {
                                    v37 := []int{32768, -9223372036854775808}
                                    _ = v37
                                } else {
                                    var v38 float64 = 0.1
                                    _ = v38
                                }
                                var v39 float32 = 3.14159
                                var v40 bool = (true && (v4 == 42))
                                _ = v39
                                _ = v40
                            }
                        }
                    }
                    for fi41 := 0; fi41 < 10; fi41++ {
                        v42 := GenS3{F0: 18446744073709551615, F1: 4294967295}
                        var v43 float32 = (v7 + 1e10)
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    _ = fi22
                    _ = v23
                }
            }
        case -274:
            var v44 uint64 = 995
            _ = v44
        }
    case "a":
        v4 -= 502
    }
    switch "" {
    case "foo bar":
        var v45 float32 = v7
        _ = v45
    case "":
        var v46 uint64 = 4294967295
        _ = v46
    default:
        v4 = v4
    }
    v4 += v4
    v47 := GenS3{F0: 18446744073709551615, F1: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v7
}
