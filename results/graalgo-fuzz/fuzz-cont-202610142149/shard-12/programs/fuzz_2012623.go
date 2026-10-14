package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 bool, p1 uint) uint16 {
    var v4 uint8 = 0
    p1 = 4294967295
    fmt.Println(p0, v4)
    p1 += p1
    fmt.Println(p1, v4, p0)
    _ = p0
    _ = p1
    _ = v4
    return 481
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint16 = 65535
    v3 := []uint32{394}
    fmt.Println(v2, v1)
    var v5 string = ""
    v2 = 668
    v2 += genFn0(false, 18446744073709551615)
    v5 = (v5 + v5)
    var v6 int8 = 127
    var v7 uint32 = 875
    if (v1 == 42) {
        if len(v3) > 0 {
            v3[0] = v7
        }
    }
    if false {
        v5 = v5
        if ((1.5 * (3.14159 - (((3.14159 - ((1e10 * (0.0 - ((3.14159 / -2.25) / -2.25))) / 1e10)) - -2.25) / 1e10))) >= (0.0 * -2.25)) {
            for fi8 := 0; fi8 < 9; fi8++ {
                v1 = (v1 / (((^(v1))) | 1))
                if false {
                    for fi9 := 0; fi9 < 15; fi9++ {
                        var v10 uint32 = v7
                        var v11 uint16 = (^(v2))
                        var v12 int64 = 957
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                } else {
                    v13 := GenS1{F0: -586}
                    if false {
                        for fi14 := 0; fi14 < 16; fi14++ {
                            switch (v5 + "hello") {
                            case "foo bar":
                                var v15 uint32 = 461
                                _ = v15
                            case "\n\t":
                                var v16 uint64 = 42
                                _ = v16
                            case "hello":
                                if (!true) {
                                    v17 := GenS2{F0: v1}
                                    _ = len(v3)
                                    _ = v17
                                }
                            default:
                                if (v5 == ((v5 + "a") + (v5 + v5))) {
                                    _ = len(v3)
                                    if false {
                                        v18 := GenS2{F0: 2147483647}
                                        fmt.Println(v6)
                                        var v19 int32 = v1
                                        _ = v18
                                        _ = v19
                                    } else {
                                        if true {
                                            v6 = (v6 & (v6 & (v6 ^ (v6 * v6))))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            fmt.Println(v6, v5, v5)
                                            var v20 bool = (0.0 == 1e10)
                                            v6 -= -128
                                            v21 := GenS1{F0: 0}
                                            v22 := GenS3{F0: (-(fi14))}
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                }
                            }
                            _ = fi14
                        }
                    } else {
                        v5 = v5
                    }
                    _ = v13
                }
                _ = fi8
            }
            switch 127 {
            case 0:
                for fi23 := 0; fi23 < 3; fi23++ {
                    v5 = "\n\t"
                    v6 = v6
                    _ = fi23
                }
            case 347:
                v6 = -50
            case -495:
                var v24 string = (v5 + v5)
                _ = v24
            default:
                var v25 uint16 = v2
                _ = v25
            }
            if false {
                for fi26 := 0; fi26 < 20; fi26++ {
                    v27 := GenS2{F0: v1}
                    var v28 int64 = -156
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v3 = append(v3, v7)
                fmt.Println(v2)
            }
        } else {
            fmt.Println(v7, v6)
            v29 := GenS1{F0: 0}
            _ = v29
        }
    }
    v30 := GenS2{F0: -893}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v5
    _ = v6
    _ = v7
}
