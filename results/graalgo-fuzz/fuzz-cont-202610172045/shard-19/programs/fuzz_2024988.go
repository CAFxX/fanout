package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint8, p1 int64) uint64 {
    if false {
        fmt.Println(p0, p0)
        var v4 int64 = p1
        _ = v4
    } else {
        v5 := GenS1{F0: p0}
        var v6 uint = 236
        if false {
            if true {
                fmt.Println(v6, p0, p0)
                if false {
                    switch -319 {
                    case 1:
                        for fi7 := 0; fi7 < 3; fi7++ {
                            for fi8 := 0; fi8 < 20; fi8++ {
                                v9 := []int{fi8, 127, fi8, 710}
                                _ = fi8
                                _ = v9
                            }
                            _ = fi7
                        }
                    case -1:
                        v10 := []int{42, 1, -2147483648, 0}
                        _ = v10
                    case 4294967295:
                        for fi11 := 0; fi11 < 8; fi11++ {
                            var v12 int32 = 268
                            var v13 int32 = (-(v12))
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    default:
                        v6 = 255
                    }
                    if false {
                        switch 4294967295 {
                        case -265:
                            if true {
                                if false {
                                    if false {
                                        var v14 bool = (-128 < -1)
                                        _ = v14
                                    }
                                    v6 -= v6
                                } else {
                                    fmt.Println(p1)
                                    switch 9223372036854775807 {
                                    case 32767:
                                        if ((1e10 + ((3.14159 / 1.0) - (0.0 + -2.25))) != 1.5) {
                                            for fi15 := 0; fi15 < 17; fi15++ {
                                                var v16 bool = true
                                                _ = fi15
                                                _ = v16
                                            }
                                        }
                                    case 2147483647:
                                        v6 -= (v6 / ((v6) | 1))
                                    case -9223372036854775808:
                                        p1 -= (p1 << (((v6 % (((^(v6))) | 1))) & 63))
                                    }
                                }
                                var v17 bool = true
                                _ = v17
                            } else {
                                v6 -= 1
                            }
                        case 447:
                            p0 += p0
                        default:
                            v18 := []int{9223372036854775807, 740}
                            _ = v18
                        }
                    }
                    var v19 int = 32767
                    _ = v19
                }
            }
        }
        _ = v5
        _ = v6
    }
    p1 = p1
    p1 += (p1 | p1)
    if false {
        fmt.Println(p1)
        v20 := GenS1{F0: p0}
        _ = v20
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint64, p1 uint) int {
    switch 2147483647 {
    case 1:
        if (2147483647 <= 42) {
            p1 = (p1 * p1)
            fmt.Println(p1, p0)
            v21 := []int{2147483647, 4294967295, 32767}
            _ = v21
        }
    default:
        v22 := GenS1{F0: 1}
        _ = v22
    }
    if (!false) {
        var v23 uint = (p1 >> ((1) & 63))
        _ = v23
    }
    switch "a" {
    case "hello":
        var v24 uint16 = 42
        _ = v24
    case "a":
        if (false && true) {
            for fi25 := 0; fi25 < 14; fi25++ {
                for fi26 := 0; fi26 < 11; fi26++ {
                    fmt.Println(p0)
                    _ = fi26
                }
                var v27 uint8 = 127
                var v28 float32 = 1.5
                _ = fi25
                _ = v27
                _ = v28
            }
            p0 = (p0 / ((genFn0(0, -662)) | 1))
            fmt.Println(p1, p0)
        } else {
            v29 := []int{42, 0, -9223372036854775808, 32767}
            _ = v29
        }
    case "foo bar":
        fmt.Println(p1)
    default:
        p0 += (p0 >> ((p1) & 63))
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    var v1 int8 = 100
    var v2 bool = true
    v3 := []string{("" + "hello"), "", ("foo bar" + "a"), (("\n\t" + "a") + "foo bar")}
    v2 = (v2 || false)
    for fi30 := 0; fi30 < 19; fi30++ {
        switch (fi30 >> ((255) & 63)) {
        case 2147483647:
            v2 = (!((true && v2) && v2))
        case 127:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        if len(v3) > 0 {
            v3[0] = ""
        }
        _ = fi30
    }
    v31 := GenS1{F0: 127}
    v3 = append(v3, ("\n\t" + ("" + ("foo bar" + "hello"))))
    v1 -= v1
    v2 = true
    v32 := GenS1{F0: 0}
    var v33 int32 = 2147483647
    v1 -= (v1 - -93)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
