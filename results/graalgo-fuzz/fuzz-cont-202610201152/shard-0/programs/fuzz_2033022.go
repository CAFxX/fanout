package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 uint32; F2 uint }

func main() {
    v1 := []string{(("a" + ((("a" + "") + "\n\t") + "hello")) + ("foo bar" + ""))}
    var v2 int64 = -9223372036854775808
    var v3 float32 = 1e10
    v3 += -2.25
    v4 := GenS1{F0: 685, F1: 1}
    v1 = append(v1, ("foo bar" + ("\n\t" + "foo bar")))
    if (987 == 296) {
        switch 4294967295 {
        case 1:
            if (false && (!(!(!true)))) {
                if true {
                    var v5 uint = 65535
                    for fi6 := 0; fi6 < 9; fi6++ {
                        if false {
                            v7 := GenS2{F0: 4294967295, F1: 559, F2: 732}
                            _ = v7
                        } else {
                            switch "hello" {
                            case "foo bar":
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                            }
                            for fi8 := 0; fi8 < 20; fi8++ {
                                v5 = v5
                                v3 = v3
                                v2 = (v2 * (v2 ^ v2))
                                if false {
                                    var v9 uint64 = 1
                                    if false {
                                        v1 = append(v1, "")
                                    } else {
                                        v3 = ((v3 * v3) - (((((((v3 - 1e10) + (v3 + ((3.14159 - 1.5) * (0.0 + float32(0.1))))) / v3) - 0.1) + v3) - (3.14159 / v3)) * -2.25))
                                    }
                                    _ = v9
                                } else {
                                    v10 := GenS2{F0: 4294967295, F1: 18446744073709551615, F2: 2147483648}
                                    _ = v10
                                }
                                _ = fi8
                            }
                        }
                        for fi11 := 0; fi11 < 17; fi11++ {
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = len(v1)
                                var v12 float32 = v3
                                _ = v12
                            } else {
                                v13 := GenS1{F0: v5, F1: 1}
                                _ = v13
                            }
                            _ = fi11
                        }
                        for fi14 := 0; fi14 < 11; fi14++ {
                            v15 := GenS3{F0: true}
                            _ = fi14
                            _ = v15
                        }
                        _ = fi6
                    }
                    v5 = 65535
                    _ = v5
                }
            }
        default:
            var v16 uint8 = 0
            _ = v16
        }
        var v17 uint32 = 683
        fmt.Println(v17, v3)
        _ = v17
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi18 := 0; fi18 < 4; fi18++ {
            v3 = v3
            if (18446744073709551615 < 4294967295) {
                v19 := GenS4{F0: -354, F1: 2147483648, F2: 0}
                _ = v19
            }
            _ = fi18
        }
        if false {
            if (42 > 42) {
                fmt.Println(v2)
                v3 -= v3
            } else {
                v3 -= v3
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 uint64 = 18446744073709551615
                    switch "" {
                    case "":
                        var v22 string = "a"
                        _ = v22
                    case "hello":
                        if (true && true) {
                            if len(v1) > 0 {
                                v1[0] = ("a" + ("\n\t" + ""))
                            }
                            fmt.Println(v2, v2, v3)
                            v3 += (v3 + float32(1.5))
                        }
                    }
                    var v23 uint64 = v21
                    _ = fi20
                    _ = v21
                    _ = v23
                }
            }
        }
    }
    var v24 uint32 = 2
    v24 += v24
    fmt.Println(v24)
    _ = len(v1)
    fmt.Println(v3, v2)
    _ = len(v1)
    fmt.Println(v2, v2, v2)
    var v25 int16 = 1000
    v24 = (^(v24))
    if (42 <= -1) {
        var v26 int = 127
        _ = v26
    }
    var v27 int64 = v2
    for fi28 := 0; fi28 < 12; fi28++ {
        if (!((-128 == -128) || true)) {
            v24 -= (v24 | v24)
        }
        var v29 int16 = v25
        _ = fi28
        _ = v29
    }
    if (65535 <= 0) {
        for fi30 := 0; fi30 < 1; fi30++ {
            for fi31 := 0; fi31 < 8; fi31++ {
                if ((^(v24)) != (v24 << ((1) & 63))) {
                    v24 -= v24
                    var v32 float32 = (1.5 + v3)
                    _ = v32
                }
                _ = fi31
            }
            var v33 uint = 0
            _ = fi30
            _ = v33
        }
        switch -176 {
        case 127:
            for fi34 := 0; fi34 < 20; fi34++ {
                var v35 bool = true
                _ = fi34
                _ = v35
            }
        case 2147483647:
            v2 = (-(v27))
        }
        v2 += v2
    }
    var v36 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v36
    _ = v4
}
