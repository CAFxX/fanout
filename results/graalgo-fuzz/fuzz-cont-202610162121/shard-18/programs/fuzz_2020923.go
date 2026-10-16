package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 string }
func genFn0() int16 {
    var v4 uint16 = 720
    if false {
        v5 := []int{128}
        _ = v5
    } else {
        var v6 uint16 = (v4 + v4)
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 uint = 0
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        _ = v6
    }
    v4 = 586
    _ = v4
    return -1
}

func main() {
    v1 := []string{"", "\n\t", "hello", "hello"}
    var v2 float64 = (1e10 - 0.1)
    var v3 int64 = 9223372036854775807
    var v10 string = ""
    fmt.Println(v3, v10)
    if ((!false) && false) {
        v2 -= v2
        switch 32767 {
        case 32768:
            switch ((v10 + (v10 + v10)) + "") {
            case "hello":
                fmt.Println(v3)
            default:
                if false {
                    v10 = ""
                }
            }
        case 302:
            v3 += (-(v3))
        case 32767:
            v11 := GenS1{F0: 0, F1: 1, F2: v10}
            _ = v11
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = len(v1)
    }
    v10 = v10
    var v12 uint = 2
    for fi13 := 0; fi13 < 16; fi13++ {
        _ = len(v1)
        fmt.Println(v2, fi13, v12)
        fmt.Println(v10, v10, v10)
        _ = fi13
    }
    if true {
        _ = len(v1)
        var v14 int8 = 0
        _ = v14
    }
    var v15 string = v10
    v12 = (v12 ^ v12)
    v10 = v10
    for fi16 := 0; fi16 < 17; fi16++ {
        if true {
            var v17 uint32 = 2147483648
            _ = v17
        }
        v12 += (v12 << ((v12) & 63))
        _ = fi16
    }
    fmt.Println(v15, v15, v15)
    if (4294967295 < 2) {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v15, v3)
        default:
            var v18 int64 = (^(v3))
            _ = v18
        }
    } else {
        fmt.Println(v2, v2)
        if true {
            for fi19 := 0; fi19 < 19; fi19++ {
                if (false && true) {
                    v20 := GenS1{F0: 124, F1: -92, F2: v15}
                    if (float64(1e10) <= v2) {
                        fmt.Println(v12)
                        switch ("hello" + (v15 + v10)) {
                        case "":
                            if false {
                                if true {
                                    var v21 uint = 333
                                    _ = v21
                                } else {
                                    v15 = ("foo bar" + (((v10 + v10) + v15) + v15))
                                }
                            }
                        case "a":
                            switch 42 {
                            case -39:
                                var v22 bool = ((-846 == 505) && (true && ((v12 != (v12 / (((v12 / ((v12) | 1))) | 1))) && (4294967295 <= 2))))
                                _ = v22
                            case 42:
                                for fi23 := 0; fi23 < 19; fi23++ {
                                    switch (fi23 >> ((v12) & 63)) {
                                    case 32767:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case -2147483648:
                                        var v24 bool = (32767 >= genFn0())
                                        _ = v24
                                    case 4294967295:
                                        var v25 float64 = (v2 / ((v2 / ((v2 - (float64(float32(v2)) * v2)) * (((v2 + 0.1) * v2) - 1e10))) - float64((1e10 * 0.1))))
                                        _ = v25
                                    default:
                                        if len(v1) > 0 {
                                            v1[0] = v10
                                        }
                                    }
                                    v26 := GenS1{F0: 2, F1: 42, F2: "a"}
                                    _ = fi23
                                    _ = v26
                                }
                            default:
                                fmt.Println(v3, v3)
                            }
                        }
                    }
                    _ = v20
                } else {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        var v28 bool = false
                        _ = fi27
                        _ = v28
                    }
                }
                _ = fi19
            }
        } else {
            if false {
                var v29 uint = v12
                fmt.Println(v15)
                if len(v1) > 0 {
                    v1[0] = v15
                }
                _ = v29
            }
            var v30 float64 = float64((1.5 * 0.0))
            _ = v30
        }
        var v31 bool = false
        _ = v31
    }
    v32 := GenS1{F0: 65535, F1: 17, F2: "hello"}
    for fi33 := 0; fi33 < 15; fi33++ {
        fmt.Println(v15, v3, v12)
        v34 := GenS1{F0: 42, F1: 127, F2: v10}
        _ = fi33
        _ = v34
    }
    fmt.Println(v15, v2)
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v32
}
