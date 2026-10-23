package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int }
type GenS2 struct { F0 int; F1 int }

func main() {
    var v1 int = 409
    var v2 int32 = 42
    var v3 uint16 = 301
    var v4 float32 = ((-2.25 * float32(1.5)) - ((-2.25 * (float32(1e10) * 3.14159)) + float32((-2.25 - 1e10))))
    v5 := GenS1{F0: -1, F1: 65535, F2: v1}
    fmt.Println(v4, v3)
    for fi6 := 0; fi6 < 4; fi6++ {
        v2 += v2
        _ = fi6
    }
    v7 := GenS1{F0: 9223372036854775807, F1: 0, F2: v1}
    v8 := GenS1{F0: -1, F1: 1, F2: (v1 | v1)}
    switch "\n\t" {
    case "":
        var v9 string = ""
        _ = v9
    case "foo bar":
        v10 := GenS1{F0: 42, F1: 265, F2: (^(v1))}
        _ = v10
    case "hello":
        var v11 uint8 = 156
        _ = v11
    default:
        if (true && (true || true)) {
            var v12 int16 = 593
            _ = v12
        } else {
            v4 = 1.5
        }
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        if true {
            var v14 int32 = 42
            v2 = (-(v2))
            _ = v14
        } else {
            if true {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v3 -= 644
                    _ = fi15
                }
                var v16 int8 = 80
                var v17 int16 = int16(v2)
                _ = v16
                _ = v17
            } else {
                v18 := []int{(fi13 - v1), (v1 / (((fi13 / ((-9223372036854775808) | 1))) | 1)), -2147483648}
                var v19 uint16 = 255
                if (!false) {
                    switch v1 {
                    case 830:
                        if len(v18) > 0 {
                            v18[0] = -9223372036854775808
                        }
                    case 0:
                        if (true || (1 > -886)) {
                            v18 = append(v18, 128)
                            switch fi13 {
                            case -2147483648:
                                for fi20 := 0; fi20 < 9; fi20++ {
                                    switch -2147483648 {
                                    case -407:
                                        v2 = v2
                                    case 128:
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                    case 9223372036854775807:
                                        var v21 string = "foo bar"
                                        _ = v21
                                    default:
                                        var v22 int = (fi20 >> ((449) & 63))
                                        _ = v22
                                    }
                                    var v23 bool = (374 < 4294967295)
                                    _ = fi20
                                    _ = v23
                                }
                            case 128:
                                for fi24 := 0; fi24 < 3; fi24++ {
                                    switch "\n\t" {
                                    case "foo bar":
                                        switch "hello" {
                                        case "a":
                                            fmt.Println(v19, fi24)
                                        case "foo bar":
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                        default:
                                            if (false || (97 != 255)) {
                                                v1 = (v1 | (fi13 % ((fi24) | 1)))
                                                v19 = (v19 & (v3 + 32767))
                                            }
                                        }
                                        var v25 float64 = (1.5 - 0.0)
                                        if len(v18) > 0 {
                                            v18[0] = 32768
                                        }
                                        _ = v25
                                    case "\n\t":
                                        v2 = -2147483648
                                    default:
                                        v4 = 0.0
                                    }
                                    _ = fi24
                                }
                            case 0:
                                var v26 uint32 = 65535
                                _ = v26
                            default:
                                if false {
                                    var v27 int16 = 42
                                    _ = v27
                                } else {
                                    var v28 int16 = -1
                                    if true {
                                        switch fi13 {
                                        case -1:
                                            v29 := GenS2{F0: (fi13 * -882), F1: fi13}
                                            _ = v29
                                        case 4294967295:
                                            for fi30 := 0; fi30 < 15; fi30++ {
                                                var v31 float64 = ((0.1 + float64(0.0)) * -2.25)
                                                _ = fi30
                                                _ = v31
                                            }
                                        default:
                                            fmt.Println(v28)
                                        }
                                        v18 = append(v18, int(v28))
                                    } else {
                                        var v32 uint8 = 191
                                        _ = len(v18)
                                        if len(v18) > 0 {
                                            v18[0] = v1
                                        }
                                        _ = v32
                                    }
                                    _ = v28
                                }
                            }
                            var v33 int = (v1 + v1)
                            _ = v33
                        }
                    }
                } else {
                    var v34 string = "hello"
                    _ = v34
                }
                _ = v18
                _ = v19
            }
            if true {
                v4 -= -2.25
                var v35 int = (v1 & (fi13 | fi13))
                _ = v35
            } else {
                v36 := []int{fi13, v1, 128, (-(fi13))}
                v36 = append(v36, 0)
                for fi37 := 0; fi37 < 10; fi37++ {
                    switch "a" {
                    case "foo bar":
                        v4 = v4
                    case "":
                        if false {
                            if false {
                                if len(v36) >= 2 {
                                    v36 = v36[1:len(v36)]
                                }
                                var v38 float32 = v4
                                v4 -= 1e10
                                _ = v38
                            }
                            v4 -= float32(((float64(v4) - 3.14159) * float64(v4)))
                        }
                    }
                    _ = fi37
                }
                _ = v36
            }
        }
        _ = fi13
    }
    v2 -= v2
    var v39 uint = 18446744073709551615
    v2 -= v2
    fmt.Println(v1)
    v1 = v1
    switch 795 {
    case 0:
        if false {
            fmt.Println(v39)
            v40 := []int{}
            fmt.Println(v1, v39, v1)
            _ = v40
        } else {
            v41 := GenS1{F0: -9223372036854775808, F1: v39, F2: 548}
            _ = v41
        }
    case -9223372036854775808:
        v3 -= 255
    }
    v42 := []int{(v1 ^ 4294967295), v1}
    if true {
        for fi43 := 0; fi43 < 19; fi43++ {
            if len(v42) >= 2 {
                v42 = v42[1:len(v42)]
            }
            _ = fi43
        }
    }
    switch v1 {
    case 170:
        v42 = append(v42, v1)
    default:
        if len(v42) >= 2 {
            v42 = v42[1:len(v42)]
        }
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v4, v3)
    }
    switch 128 {
    case 128:
        fmt.Println(v4, v39, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v7
    _ = v8
}
