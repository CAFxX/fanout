package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint16) uint16 {
    fmt.Println(p0)
    var v4 uint64 = 998
    _ = p0
    _ = v4
    return p0
}

func main() {
    var v1 uint16 = 2
    var v2 int64 = 9223372036854775807
    v3 := []string{""}
    _ = len(v3)
    v5 := GenS1{F0: 4294967295}
    if true {
        if false {
            v3 = append(v3, "")
        } else {
            for fi6 := 0; fi6 < 20; fi6++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi6
            }
        }
        if true {
            var v7 int16 = 1
            v1 += v1
            _ = v7
        }
    } else {
        if true {
            v8 := GenS1{F0: 2}
            var v9 uint = 1
            _ = v8
            _ = v9
        } else {
            fmt.Println(v1, v2)
            v10 := GenS1{F0: 588}
            var v11 uint64 = 4294967295
            _ = v10
            _ = v11
        }
        v2 = (v2 / (((v2 % ((319) | 1))) | 1))
    }
    v12 := GenS1{F0: 1}
    v13 := GenS1{F0: 0}
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            fmt.Println(v2, fi14)
            v1 += 228
            v2 = v2
            _ = fi15
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi14
    }
    if (4294967295 < 18446744073709551615) {
        v1 = genFn0(65535)
    }
    v1 -= v1
    if false {
        if true {
            if (!true) {
                if false {
                    if true {
                        if (float64(1.5) == (((0.0 / 1e10) - -2.25) / 3.14159)) {
                            switch -456 {
                            case 1:
                                v2 += (v2 - v2)
                            case -9223372036854775808:
                                var v16 int64 = v2
                                _ = v16
                            default:
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                            }
                            switch "" {
                            case "a":
                                for fi17 := 0; fi17 < 20; fi17++ {
                                    v1 = (v1 >> ((2) & 63))
                                    _ = fi17
                                }
                            }
                            if false {
                                v2 += 9223372036854775807
                            }
                        }
                    }
                }
                if false {
                    var v18 int16 = -421
                    for fi19 := 0; fi19 < 9; fi19++ {
                        if true {
                            for fi20 := 0; fi20 < 20; fi20++ {
                                fmt.Println(v18, v18, v18)
                                if true {
                                    fmt.Println(v2, v2)
                                } else {
                                    var v21 int8 = 0
                                    v3 = append(v3, "\n\t")
                                    _ = v21
                                }
                                _ = fi20
                            }
                        } else {
                            v22 := GenS1{F0: 65535}
                            _ = v22
                        }
                        _ = fi19
                    }
                    _ = v18
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
            }
        } else {
            switch ("" + "foo bar") {
            case "\n\t":
                var v23 bool = false
                _ = v23
            case "a":
                if len(v3) > 0 {
                    v3[0] = ""
                }
            case "":
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := GenS2{F0: 255}
                    v3 = append(v3, "a")
                    v1 = 416
                    _ = fi24
                    _ = v25
                }
            default:
                var v26 uint32 = 2147483648
                _ = v26
            }
            v1 -= (^(v1))
            var v27 uint16 = (^(v1))
            _ = v27
        }
        fmt.Println(v1, v1)
    } else {
        var v28 float32 = 0.0
        _ = v28
    }
    fmt.Println(v2, v2)
    var v29 int64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v5
}
