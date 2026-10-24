package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 uint; F2 string }
type GenS5 struct { F0 int16; F1 string; F2 int16 }
func genFn0() uint64 {
    var v4 int32 = -2147483648
    var v5 float32 = 1e10
    if (-69 > 42) {
        var v6 int16 = -32768
        var v7 uint16 = 32767
        v4 += (v4 ^ v4)
        _ = v6
        _ = v7
    } else {
        v8 := []int{340, 563, 42, 92}
        v4 = v4
        if false {
            v4 = (-(v4))
            v8 = append(v8, 32767)
            fmt.Println(v5, v4)
        }
        _ = v8
    }
    v5 += ((v5 * float32((float64(float32(1e10)) * (1.5 * (1e10 / (3.14159 / 0.1)))))) + (1e10 * (v5 / (1.5 + v5))))
    v5 += v5
    _ = v4
    _ = v5
    return 670
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = 1
    var v3 int8 = 1
    if (!((501 == 1) || (42 > 9223372036854775807))) {
        v9 := GenS1{F0: 733}
        v2 += v2
        _ = v9
    }
    v10 := []int{-208, -1, 818}
    v1 = genFn0()
    v11 := GenS1{F0: -995}
    for fi12 := 0; fi12 < 14; fi12++ {
        if true {
            if false {
                switch int(v1) {
                case -351:
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 uint = 82
                        var v15 int64 = v2
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                case 42:
                    var v16 uint = 1
                    _ = v16
                }
            } else {
                var v17 bool = false
                v2 += v2
                _ = v17
            }
            if false {
                for fi18 := 0; fi18 < 4; fi18++ {
                    v10 = append(v10, 127)
                    v10 = append(v10, 700)
                    fmt.Println(v2, fi18)
                    _ = fi18
                }
                v19 := GenS1{F0: (fi12 | (fi12 + 32767))}
                var v20 float64 = (0.1 + 1.5)
                _ = v19
                _ = v20
            }
            var v21 float64 = 0.1
            _ = v21
        } else {
            var v22 uint8 = 128
            v2 += 0
            if ((0.0 > float32((3.14159 - 1.5))) && false) {
                switch (("\n\t" + ("foo bar" + "hello")) + "a") {
                case "a":
                    if true {
                        var v23 bool = false
                        switch "a" {
                        case "hello":
                            var v24 int32 = 1
                            _ = v24
                        case "":
                            v22 = v22
                        default:
                            if false {
                                for fi25 := 0; fi25 < 4; fi25++ {
                                    v1 += (v1 + v1)
                                    _ = fi25
                                }
                            } else {
                                v23 = true
                            }
                        }
                        for fi26 := 0; fi26 < 14; fi26++ {
                            fmt.Println(fi26, v22)
                            v10 = append(v10, fi26)
                            v1 += v1
                            v27 := GenS2{F0: fi12, F1: -1, F2: (v22 << ((65535) & 63))}
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = fi26
                            _ = v27
                        }
                        _ = v23
                    }
                case "":
                    v22 = v22
                default:
                    fmt.Println(v22, v2, v2)
                }
            } else {
                fmt.Println(v2, v22)
                switch "" {
                case "foo bar":
                    v10 = append(v10, -609)
                }
            }
            _ = v22
        }
        for fi28 := 0; fi28 < 1; fi28++ {
            if true {
                if (true && (v3 > v3)) {
                    if false {
                        fmt.Println(v2)
                        fmt.Println(fi28)
                        if true {
                            var v29 int16 = 1
                            var v30 uint8 = 128
                            _ = v29
                            _ = v30
                        }
                    } else {
                        if ((!false) && (0.1 > 0.0)) {
                            v1 += (v1 * (v1 * 18446744073709551615))
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            var v31 uint = 65535
                            _ = v31
                        }
                    }
                    v32 := GenS3{F0: v1}
                    _ = v32
                } else {
                    if false {
                        var v33 int = 55
                        v34 := GenS3{F0: 247}
                        _ = v33
                        _ = v34
                    }
                    var v35 int = 32768
                    if true {
                        _ = len(v10)
                        v35 += (^(fi12))
                    }
                    if (127 < 255) {
                        switch 559 {
                        case 32767:
                            v2 = v2
                        default:
                            fmt.Println(v35, v3, fi28)
                        }
                    }
                    _ = v35
                }
                if (-2.25 > (0.1 * 0.0)) {
                    var v36 int32 = 5
                    var v37 uint16 = 1
                    _ = v36
                    _ = v37
                } else {
                    v38 := GenS4{F0: 73, F1: 1, F2: "foo bar"}
                    _ = v38
                }
                fmt.Println(v3)
            }
            var v39 int64 = 9223372036854775807
            fmt.Println(fi28)
            _ = fi28
            _ = v39
        }
        _ = fi12
    }
    var v40 uint32 = 634
    v2 = (^(v2))
    var v41 uint = 63
    for fi42 := 0; fi42 < 20; fi42++ {
        for fi43 := 0; fi43 < 2; fi43++ {
            v41 -= v41
            _ = fi43
        }
        v44 := GenS5{F0: 294, F1: ("foo bar" + "foo bar"), F2: 977}
        _ = len(v10)
        _ = fi42
        _ = v44
    }
    v45 := GenS5{F0: 32767, F1: "hello", F2: 1}
    fmt.Println(v3, v1, v3)
    fmt.Println(v1)
    v2 = (v2 % (((v2 * (v2 - v2))) | 1))
    switch ("\n\t" + "\n\t") {
    case "":
        v10 = append(v10, 1)
    case "a":
        v46 := GenS2{F0: 32767, F1: 826, F2: 42}
        _ = v46
    default:
        v2 = (-(v2))
    }
    switch "" {
    case "":
        if len(v10) > 0 {
            v10[0] = 9223372036854775807
        }
    }
    v1 += (v1 % (((v1 / ((uint64(v40)) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v45
}
