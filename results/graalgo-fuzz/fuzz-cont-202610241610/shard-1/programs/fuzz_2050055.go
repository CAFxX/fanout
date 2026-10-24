package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint32 }
type GenS2 struct { F0 bool; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 uint16; F1 uint32; F2 int64 }
func genFn0(p0 bool, p1 int64) int {
    p0 = (p1 >= 9223372036854775807)
    p0 = false
    var v6 uint16 = 793
    var v7 bool = (4294967295 != 0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 float32 = 0.1
    var v2 uint64 = 149
    var v3 float64 = 1e10
    var v4 float64 = v3
    var v5 int16 = -1
    var v8 float32 = v1
    v1 = (v8 - (3.14159 - 0.1))
    v1 -= v8
    v9 := GenS1{F0: "a", F1: v2, F2: 638}
    v10 := []int{-9223372036854775808}
    var v11 int8 = 1
    switch 9223372036854775807 {
    case 2147483647:
        if false {
            v2 += 42
            v11 += 100
            if (42 <= 4294967295) {
                var v12 int16 = (v5 - v5)
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 int64 = 1
                    for fi15 := 0; fi15 < 18; fi15++ {
                        var v16 uint16 = 549
                        v11 += (^(v11))
                        fmt.Println(v11, v4, v2)
                        _ = fi15
                        _ = v16
                    }
                    var v17 uint8 = 2
                    _ = fi13
                    _ = v14
                    _ = v17
                }
                _ = v12
            } else {
                v8 += v8
                v4 += v3
            }
        } else {
            var v18 uint64 = (v2 & v2)
            for fi19 := 0; fi19 < 13; fi19++ {
                if (!true) {
                    for fi20 := 0; fi20 < 16; fi20++ {
                        v21 := GenS2{F0: (-1 >= 375), F1: 9223372036854775807, F2: 42}
                        if false {
                            v10 = append(v10, 9223372036854775807)
                            var v22 int = -3
                            _ = v22
                        } else {
                            v23 := GenS3{F0: 128, F1: (v18 + (^(v18))), F2: 42}
                            _ = v23
                        }
                        _ = fi20
                        _ = v21
                    }
                } else {
                    fmt.Println(v11, fi19, v4)
                    switch (("foo bar" + "") + "foo bar") {
                    case "\n\t":
                        var v24 uint32 = 1
                        _ = v24
                    case "":
                        if false {
                            switch (fi19 << ((2) & 63)) {
                            case 250:
                                if (false || true) {
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    fmt.Println(fi19, fi19)
                                } else {
                                    var v25 int = fi19
                                    if true {
                                        for fi26 := 0; fi26 < 14; fi26++ {
                                            var v27 int16 = 32767
                                            _ = fi26
                                            _ = v27
                                        }
                                        for fi28 := 0; fi28 < 11; fi28++ {
                                            _ = len(v10)
                                            _ = fi28
                                        }
                                    } else {
                                        v8 = v1
                                        v29 := GenS4{F0: 982}
                                        if (!(!(true && (false && false)))) {
                                            v30 := GenS2{F0: (4294967295 != 363), F1: 9223372036854775807, F2: 1}
                                            _ = v30
                                        } else {
                                            for fi31 := 0; fi31 < 5; fi31++ {
                                                v10 = append(v10, 9223372036854775807)
                                                if len(v10) > 0 {
                                                    v10[0] = (fi31 % ((genFn0(((v11 + (v11 + (-(v11)))) <= 127), 19)) | 1))
                                                }
                                                _ = fi31
                                            }
                                        }
                                        _ = v29
                                    }
                                    v32 := GenS2{F0: (v18 > (v2 - v18)), F1: -296, F2: 128}
                                    _ = v25
                                    _ = v32
                                }
                            case 2147483647:
                                v10 = append(v10, fi19)
                            case 128:
                                for fi33 := 0; fi33 < 3; fi33++ {
                                    var v34 int16 = v5
                                    _ = fi33
                                    _ = v34
                                }
                            default:
                                v10 = append(v10, 2147483647)
                            }
                            v11 = v11
                            var v35 uint = 18446744073709551615
                            _ = v35
                        } else {
                            if len(v10) > 0 {
                                v10[0] = 32767
                            }
                            switch fi19 {
                            case 42:
                                switch -56 {
                                case -2147483648:
                                    fmt.Println(v3, v8)
                                }
                            case 32768:
                                if true {
                                    if len(v10) > 0 {
                                        v10[0] = (fi19 - fi19)
                                    }
                                    v5 = -351
                                    var v36 string = "hello"
                                    _ = v36
                                }
                            default:
                                var v37 int8 = 127
                                _ = v37
                            }
                        }
                    case "foo bar":
                        if (((v4 - (v3 + v3)) >= float64((v1 * ((v1 / v1) * (((v1 * ((float32(3.14159) - float32(v4)) - (1.5 * -2.25))) + v8) / 3.14159))))) || (128 != 139)) {
                            v11 -= (v11 * v11)
                            v2 += v18
                        }
                    }
                    v38 := GenS5{F0: (v11 & v11), F1: (v2 ^ v2)}
                    _ = v38
                }
                if true {
                    v4 += float64(0.0)
                } else {
                    for fi39 := 0; fi39 < 18; fi39++ {
                        var v40 float32 = ((v8 + v8) / float32(v3))
                        v41 := GenS4{F0: -9223372036854775808}
                        var v42 uint32 = 2147483648
                        _ = fi39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    fmt.Println(fi19)
                    switch "hello" {
                    case "foo bar":
                        v3 = float64((((v8 * 1e10) / v1) * float32(3.14159)))
                    case "\n\t":
                        v43 := GenS1{F0: "hello", F1: (v2 / (((v18 << ((417) & 63))) | 1)), F2: 162}
                        _ = v43
                    }
                }
                _ = fi19
            }
            _ = v18
        }
    default:
        var v44 uint8 = 222
        _ = v44
    }
    var v45 bool = true
    v4 += v4
    v4 = v4
    v4 = (float64(3.14159) - (((1.5 * ((((((-2.25 + v3) * -2.25) / 1.5) * v3) * 1e10) - v4)) / float64(((v8 * float32(-2.25)) - ((((v8 - v1) - ((1.5 - v8) * v8)) + v1) - 1e10)))) - v3))
    for fi46 := 0; fi46 < 8; fi46++ {
        switch "foo bar" {
        case "hello":
            for fi47 := 0; fi47 < 19; fi47++ {
                var v48 int32 = -1
                var v49 int8 = -1
                _ = fi47
                _ = v48
                _ = v49
            }
        }
        var v50 float64 = -2.25
        _ = fi46
        _ = v50
    }
    v45 = v45
    v51 := GenS6{F0: 42, F1: 1, F2: int64(v11)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v51
    _ = v8
    _ = v9
}
