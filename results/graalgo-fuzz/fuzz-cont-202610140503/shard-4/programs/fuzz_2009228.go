package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 string; F1 uint }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{128, 9223372036854775807}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    var v5 int32 = 2147483647
    var v6 uint32 = 890
    var v7 float64 = (1.5 - 3.14159)
    for fi8 := 0; fi8 < 2; fi8++ {
        if ((fi8 >> ((18446744073709551615) & 63)) < 127) {
            v1 = 0
        } else {
            v9 := GenS1{F0: fi8}
            if false {
                for fi10 := 0; fi10 < 20; fi10++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (!true) {
                        v11 := GenS1{F0: -2147483648}
                        _ = v11
                    } else {
                        v2 = append(v2, (^(fi8)))
                        var v12 float64 = (float64(1.5) * v7)
                        if ((v7 + (v12 + v7)) >= v12) {
                            var v13 string = ("hello" + "\n\t")
                            v7 += (float64((0.0 - 3.14159)) - (v12 * (0.0 + v7)))
                            if true {
                                if true {
                                    switch "" {
                                    case "foo bar":
                                        if true {
                                            fmt.Println(v12, fi10, v1)
                                        } else {
                                            v5 = int32(v1)
                                            fmt.Println(v12, v12, fi10)
                                        }
                                    case "hello":
                                        var v14 int32 = (v5 % ((int32(fi10)) | 1))
                                        _ = v14
                                    }
                                }
                                v12 += float64(((((0.0 - 3.14159) - (3.14159 + 0.1)) - 1.5) * ((0.0 - 0.0) + ((float32(v7) - 1.5) + float32(v7)))))
                            } else {
                                var v15 int64 = 9223372036854775807
                                v2 = append(v2, fi8)
                                switch fi8 {
                                case 127:
                                    var v16 string = "foo bar"
                                    _ = v16
                                }
                                _ = v15
                            }
                            _ = v13
                        }
                        v7 = (v7 / v12)
                        v7 -= 1e10
                        _ = v12
                    }
                    _ = fi10
                }
            } else {
                var v17 float32 = 3.14159
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 bool = (2 > 127)
                    switch fi18 {
                    case -2147483648:
                        v2 = append(v2, fi8)
                    case -9223372036854775808:
                        if ((v6 >> ((65535) & 63)) == (^(v6))) {
                            v20 := GenS1{F0: 576}
                            _ = v20
                        } else {
                            if len(v2) > 0 {
                                v2[0] = fi18
                            }
                            var v21 int32 = 1
                            var v22 int32 = v5
                            fmt.Println(fi18, v1, v19)
                            _ = v21
                            _ = v22
                        }
                    default:
                        v7 += (((v7 + (v7 - ((v7 + v7) - (((v7 * (1e10 * (v7 * v7))) - (v7 - 3.14159)) + (-2.25 / -2.25))))) - 0.0) - v7)
                    }
                    v19 = v19
                    _ = fi18
                    _ = v19
                }
                switch "foo bar" {
                case "a":
                    _ = len(v2)
                case "\n\t":
                    if (0 < 0) {
                        v17 = -2.25
                        v5 -= v5
                        v1 -= v1
                    } else {
                        var v23 string = ("a" + ("" + ("" + "a")))
                        v23 = v23
                        v17 += 0.1
                        _ = v23
                    }
                case "foo bar":
                    v17 += v17
                default:
                    if true {
                        v1 -= 9223372036854775807
                    } else {
                        if (!(!(2 == 2))) {
                            var v24 uint8 = 2
                            v7 = v7
                            _ = v24
                        }
                        v25 := GenS2{F0: (v1 >> ((683) & 63)), F1: 2}
                        var v26 uint = 465
                        _ = v25
                        _ = v26
                    }
                }
                _ = v17
            }
            v7 -= 0.0
            _ = v9
        }
        _ = fi8
    }
    var v27 int32 = v5
    v2 = append(v2, 127)
    switch "" {
    case "":
        for fi28 := 0; fi28 < 19; fi28++ {
            if true {
                if true {
                    v29 := GenS2{F0: (v1 % ((v1) | 1)), F1: 0}
                    v27 = 541
                    var v30 uint64 = 42
                    _ = v29
                    _ = v30
                } else {
                    v31 := GenS1{F0: fi28}
                    fmt.Println(v1)
                    v32 := v3[v6]
                    _ = v31
                    _ = v32
                }
                fmt.Println(fi28)
                delete(v3, v6)
            }
            v6 = v6
            var v33 uint = 186
            _ = fi28
            _ = v33
        }
    case "a":
        _ = len(v4)
    default:
        switch "a" {
        case "\n\t":
            for fi34 := 0; fi34 < 18; fi34++ {
                v6 -= 391
                if true {
                    v35 := GenS3{F0: "", F1: 2}
                    _ = len(v3)
                    _ = v35
                } else {
                    v36 := GenS1{F0: (fi34 * (fi34 ^ fi34))}
                    _ = v36
                }
                for fi37 := 0; fi37 < 20; fi37++ {
                    fmt.Println(v6, fi34)
                    var v38 uint8 = 42
                    _ = fi37
                    _ = v38
                }
                _ = fi34
            }
        case "hello":
            v27 = v27
        }
    }
    v2 = append(v2, -956)
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    if (42 > 100) {
        var v39 uint32 = (^(v6))
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
