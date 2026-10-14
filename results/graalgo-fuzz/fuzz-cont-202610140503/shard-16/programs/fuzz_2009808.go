package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint; F1 string; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 uint64; F1 uint32; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = float32((-2.25 * float64(3.14159)))
    var v3 float64 = 1e10
    if (171 < 1000) {
        for fi4 := 0; fi4 < 17; fi4++ {
            for fi5 := 0; fi5 < 16; fi5++ {
                v2 += (0.0 * v2)
                switch "" {
                case "hello":
                    if (!(!(647 <= 0))) {
                        v6 := []int{(^(fi5))}
                        _ = v6
                    }
                default:
                    v3 = v3
                }
                v7 := GenS1{F0: 18446744073709551615}
                _ = fi5
                _ = v7
            }
            v8 := GenS1{F0: 817}
            _ = fi4
            _ = v8
        }
    } else {
        v3 -= v3
        var v9 int32 = 646
        var v10 int16 = int16(v9)
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int32 = 42
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS2{F0: 42, F1: 2}
            _ = fi13
            _ = v14
        }
        fmt.Println(v3)
        _ = fi11
        _ = v12
    }
    v15 := []int{352, 9223372036854775807, -38}
    fmt.Println(v2)
    v16 := GenS2{F0: 255, F1: 0}
    var v17 uint16 = 0
    v18 := GenS1{F0: 4294967295}
    switch "hello" {
    case "foo bar":
        if false {
            switch 1 {
            case -2147483648:
                if (1 != 0) {
                    var v19 int32 = 1
                    v20 := GenS3{F0: 65535, F1: 141}
                    _ = v19
                    _ = v20
                } else {
                    switch 1 {
                    case 276:
                        switch 42 {
                        case -611:
                            var v21 float64 = ((-2.25 - 0.1) - float64((3.14159 * 1.5)))
                            _ = v21
                        case 128:
                            for fi22 := 0; fi22 < 14; fi22++ {
                                var v23 uint16 = (v17 % ((v17) | 1))
                                v24 := GenS4{F0: 1, F1: "a", F2: 147}
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        case 4294967295:
                            v1["foo bar"] = 2147483648
                        }
                    case 127:
                        _ = len(v1)
                    case -1:
                        v25 := GenS5{F0: ("hello" >= "\n\t"), F1: 255}
                        _ = v25
                    default:
                        v2 += v2
                    }
                    for fi26 := 0; fi26 < 4; fi26++ {
                        if (!(!(!((63 != 1) || false)))) {
                            if false {
                                if true {
                                    fmt.Println(v3, v3)
                                    var v27 int64 = 9223372036854775807
                                    var v28 int64 = v27
                                    _ = v27
                                    _ = v28
                                }
                            }
                            for fi29 := 0; fi29 < 13; fi29++ {
                                var v30 uint = 528
                                _ = fi29
                                _ = v30
                            }
                        } else {
                            if (42 == 1) {
                                var v31 int = fi26
                                for fi32 := 0; fi32 < 1; fi32++ {
                                    fmt.Println(fi26)
                                    var v33 uint16 = v17
                                    fmt.Println(fi26, v3)
                                    _ = fi32
                                    _ = v33
                                }
                                _ = v31
                            } else {
                                v2 = 1e10
                                fmt.Println(v2, fi26, fi26)
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    fmt.Println(v3, fi26, fi34)
                                    var v35 uint8 = 128
                                    switch -486 {
                                    case -558:
                                        if true {
                                            fmt.Println(fi26, v3)
                                            v36 := GenS5{F0: true, F1: (v35 & v35)}
                                            var v37 uint8 = (v35 + v35)
                                            _ = v36
                                            _ = v37
                                        } else {
                                            v2 -= (v2 * (float32(v3) * 1e10))
                                        }
                                    }
                                    _ = fi34
                                    _ = v35
                                }
                            }
                            switch "a" {
                            case "":
                                var v38 int16 = 1
                                _ = v38
                            case "hello":
                                var v39 int16 = 1000
                                _ = v39
                            case "a":
                                var v40 int = (fi26 / ((-353) | 1))
                                _ = v40
                            default:
                                var v41 uint64 = 4294967295
                                _ = v41
                            }
                            v17 -= (^(v17))
                        }
                        _ = fi26
                    }
                    if len(v15) > 0 {
                        v15[0] = 42
                    }
                }
            case 1:
                if true {
                    var v42 float64 = v3
                    v43 := GenS2{F0: 4294967295, F1: uint32(v17)}
                    _ = v42
                    _ = v43
                } else {
                    v44 := GenS6{F0: 42, F1: 1, F2: 42}
                    if len(v15) > 0 {
                        v15[0] = 32767
                    }
                    _ = v44
                }
            }
        } else {
            v2 += (3.14159 / v2)
        }
    case "\n\t":
        _ = len(v15)
    }
    var v45 uint = 255
    if true {
        switch ("\n\t" + "hello") {
        case "hello":
            v2 = v2
        case "a":
            var v46 int64 = 42
            _ = v46
        case "foo bar":
            v2 = (0.1 - v2)
        }
    } else {
        fmt.Println(v3)
    }
    fmt.Println(v2, v3, v3)
    v45 = v45
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v45
}
