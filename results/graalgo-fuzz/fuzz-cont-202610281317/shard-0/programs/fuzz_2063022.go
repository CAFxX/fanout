package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int; F1 int32; F2 int32 }
type GenS6 struct { F0 int32; F1 uint16 }

func main() {
    v1 := []int{605, 127}
    var v2 int8 = -103
    var v3 int = 32767
    if (float32(((1e10 / 1.0) / -2.25)) <= (float32(1.5) / 0.1)) {
        switch ("foo bar" + "foo bar") {
        case "":
            var v4 int8 = v2
            _ = v4
        }
        _ = len(v1)
    }
    fmt.Println(v3, v2)
    var v5 int32 = 42
    if true {
        v1 = append(v1, (-(v3)))
        switch v3 {
        case 82:
            switch "a" {
            case "":
                if (!true) {
                    switch "hello" {
                    case "foo bar":
                        if true {
                            v6 := GenS1{F0: 0}
                            if (false && false) {
                                v7 := GenS1{F0: -9223372036854775808}
                                _ = v7
                            }
                            var v8 bool = false
                            _ = v6
                            _ = v8
                        }
                    }
                }
            case "foo bar":
                v2 += v2
            case "\n\t":
                v9 := GenS1{F0: -9223372036854775808}
                _ = v9
            }
        case 45:
            if true {
                if ((-(v5)) == v5) {
                    v3 = (v3 ^ v3)
                    var v10 int64 = 967
                    _ = v10
                }
            }
        case 1:
            v1 = append(v1, (v3 << ((18446744073709551615) & 63)))
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint32 = 0
            if ((-(v3)) <= v3) {
                var v13 uint = 757
                for fi14 := 0; fi14 < 5; fi14++ {
                    v15 := GenS2{F0: 2, F1: (v3 + 352), F2: 43}
                    if ((0.0 + 0.0) < -2.25) {
                        var v16 float32 = 3.14159
                        _ = v16
                    }
                    var v17 uint32 = (v12 * v12)
                    v17 -= (v12 | v17)
                    v1 = append(v1, 128)
                    var v18 bool = false
                    _ = fi14
                    _ = v15
                    _ = v17
                    _ = v18
                }
                if (535 < 255) {
                    fmt.Println(v3, v5, v13)
                }
                _ = v13
            } else {
                if true {
                    v3 += fi11
                    for fi19 := 0; fi19 < 4; fi19++ {
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(fi19)
                        case "hello":
                            v12 += (^(v12))
                        case "a":
                            if len(v1) > 0 {
                                v1[0] = fi11
                            }
                        default:
                            switch "" {
                            case "":
                                if false {
                                    var v20 uint64 = 2
                                    var v21 uint = 1
                                    _ = v20
                                    _ = v21
                                } else {
                                    switch (v3 - (fi19 << ((18446744073709551615) & 63))) {
                                    case 140:
                                        var v22 int = 609
                                        _ = v22
                                    case 9223372036854775807:
                                        if (!false) {
                                            if len(v1) > 0 {
                                                v1[0] = (fi11 / ((172) | 1))
                                            }
                                            _ = len(v1)
                                            v12 += 65535
                                        } else {
                                            if true {
                                                v23 := GenS1{F0: 42}
                                                v24 := GenS1{F0: 259}
                                                _ = v23
                                                _ = v24
                                            } else {
                                                _ = len(v1)
                                                v25 := GenS2{F0: 128, F1: (v3 / ((v3) | 1)), F2: 1}
                                                v26 := GenS3{F0: (-(fi19)), F1: 255}
                                                _ = v25
                                                _ = v26
                                            }
                                        }
                                    case 679:
                                        v5 += (v5 << ((0) & 63))
                                    default:
                                        if (1.5 >= (1.5 * 0.0)) {
                                            var v27 uint32 = (v12 << ((392) & 63))
                                            var v28 uint = 4294967295
                                            _ = v27
                                            _ = v28
                                        } else {
                                            var v29 uint32 = (v12 & 844)
                                            _ = v29
                                        }
                                    }
                                }
                            case "foo bar":
                                if (!((v5 - v5) <= (v5 | v5))) {
                                    var v30 int8 = v2
                                    fmt.Println(v12)
                                    _ = v30
                                } else {
                                    v12 += v12
                                    if (1000 == 1000) {
                                        _ = len(v1)
                                        v31 := GenS3{F0: 232, F1: 172}
                                        v32 := GenS4{F0: v12, F1: v2}
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v33 uint8 = 121
                                        v34 := GenS1{F0: 9223372036854775807}
                                        _ = v33
                                        _ = v34
                                    }
                                }
                            }
                        }
                        v5 += 2147483647
                        v35 := GenS5{F0: (fi19 | (-(fi11))), F1: 42, F2: v5}
                        _ = fi19
                        _ = v35
                    }
                }
            }
            _ = fi11
            _ = v12
        }
    }
    v36 := GenS6{F0: -127, F1: 2}
    v1 = append(v1, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v37 := GenS5{F0: (^(v3)), F1: v5, F2: (^(v5))}
    fmt.Println(v5, v3)
    v1 = append(v1, v3)
    v3 = v3
    v38 := GenS3{F0: v3, F1: 949}
    var v39 uint8 = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v5
}
