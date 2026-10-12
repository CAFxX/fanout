package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0() string {
    for fi5 := 0; fi5 < 1; fi5++ {
        switch (fi5 / ((fi5) | 1)) {
        case 0:
            fmt.Println(fi5)
        case 9223372036854775807:
            fmt.Println(fi5)
        default:
            v6 := GenS1{F0: 375}
            _ = v6
        }
        _ = fi5
    }
    v7 := []int{}
    _ = v7
    return ""
}
func genFn1() uint {
    var v8 uint16 = 1
    v9 := GenS1{F0: 600}
    switch 128 {
    case 32768:
        v10 := GenS1{F0: 470}
        _ = v10
    case 127:
        fmt.Println(v8)
    }
    v8 = 328
    v11 := []int{}
    _ = v11
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 int = 2147483647
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    var v4 bool = true
    if (genFn1() <= genFn1()) {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        var v12 uint = 1
        _ = v12
    } else {
        if len(v2) > 0 {
            v2[0] = 65535
        }
        switch (v1 & -2147483648) {
        case 4294967295:
            v4 = v4
        case 665:
            fmt.Println(v4, v1)
        }
        switch v1 {
        case -413:
            v4 = (false || (65535 > 1))
        default:
            if v4 {
                var v13 int64 = 9223372036854775807
                var v14 uint = genFn1()
                if v4 {
                    switch "foo bar" {
                    case "":
                        var v15 bool = (v4 || v4)
                        _ = v15
                    case "\n\t":
                        v16 := GenS2{F0: -510, F1: 114}
                        _ = v16
                    case "hello":
                        var v17 uint32 = 618
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 8; fi18++ {
                        fmt.Println(v1, fi18)
                        delete(v3, (v1 % (((^(fi18))) | 1)))
                        _ = fi18
                    }
                    delete(v3, (v1 | (^(v1))))
                } else {
                    if ((247 < 127) && true) {
                        var v19 int64 = v13
                        v20 := GenS1{F0: 1}
                        if false {
                            v21 := GenS1{F0: uint64(v14)}
                            _ = v21
                        } else {
                            v14 -= v14
                            if (-16 >= -128) {
                                fmt.Println(v4, v4)
                                v4 = (572 <= uint32(v14))
                                if ((v13 % ((169) | 1)) != (-(v19))) {
                                    var v22 uint64 = 511
                                    _ = v22
                                } else {
                                    var v23 string = ("\n\t" + "\n\t")
                                    v24 := GenS1{F0: 2}
                                    v4 = (((v23 + "foo bar") + (v23 + ("a" + (v23 + "foo bar")))) == (v23 + v23))
                                    _ = v23
                                    _ = v24
                                }
                            }
                            if ((v14 + 65535) != (^(v14))) {
                                var v25 bool = (-32768 < 1000)
                                fmt.Println(v13, v19, v14)
                                if (1 < -492) {
                                    var v26 int8 = 100
                                    v19 = (v13 * (v19 & v19))
                                    v14 = (v14 + (^(v14)))
                                    _ = v26
                                }
                                _ = v25
                            } else {
                                if (1e10 <= float32(((((0.1 * 0.0) - 1.5) - 0.1) / 3.14159))) {
                                    if v4 {
                                        v1 -= v1
                                        var v27 string = "foo bar"
                                        _ = v27
                                    } else {
                                        v2 = append(v2, 773)
                                        var v28 float64 = -2.25
                                        _ = v28
                                    }
                                    var v29 int16 = int16(v14)
                                    _ = v29
                                } else {
                                    v30 := GenS1{F0: 0}
                                    _ = v30
                                }
                                for fi31 := 0; fi31 < 9; fi31++ {
                                    var v32 uint16 = 698
                                    _ = fi31
                                    _ = v32
                                }
                            }
                        }
                        _ = v19
                        _ = v20
                    }
                }
                _ = v13
                _ = v14
            }
        }
    }
    v4 = (-1 > 32767)
    for fi33 := 0; fi33 < 10; fi33++ {
        v4 = (!v4)
        fmt.Println(v1, v4)
        _ = fi33
    }
    fmt.Println(v1, v1)
    var v34 int8 = 100
    v1 = 42
    switch (-(v1)) {
    case 2147483647:
        v34 = v34
    case -9223372036854775808:
        for fi35 := 0; fi35 < 8; fi35++ {
            v36 := GenS1{F0: 4294967295}
            _ = fi35
            _ = v36
        }
    }
    var v37 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v4
}
