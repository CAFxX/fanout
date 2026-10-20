package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8; F1 int; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int8 }

func main() {
    var v1 string = "hello"
    var v2 uint64 = 386
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int8 = 100
        _ = fi5
        _ = v6
    }
    if (1 <= -75) {
        switch v1 {
        case "\n\t":
            v1 = ("foo bar" + v1)
        default:
            v2 -= (^(v2))
        }
        var v7 int16 = 1
        for fi8 := 0; fi8 < 1; fi8++ {
            if (true || true) {
                v3[v1] = 42
                var v9 uint32 = 4294967295
                _ = len(v3)
                _ = v9
            } else {
                v1 = (v1 + "")
            }
            _ = fi8
        }
        _ = v7
    }
    v2 -= v2
    v10 := []int{526, 9223372036854775807}
    v2 -= v2
    if (4294967295 <= 4294967295) {
        if (127 != 2147483647) {
            v2 += v2
            if (255 >= 1) {
                v2 += (v2 * v2)
            } else {
                var v11 int = -444
                _ = v11
            }
            switch v1 {
            case "hello":
                switch v1 {
                case "foo bar":
                    var v12 bool = (false || true)
                    _ = v12
                case "a":
                    _ = len(v3)
                default:
                    for fi13 := 0; fi13 < 18; fi13++ {
                        _ = len(v10)
                        if (79 > 4294967295) {
                            for fi14 := 0; fi14 < 15; fi14++ {
                                if (!false) {
                                    v15 := GenS1{F0: -124, F1: 1}
                                    var v16 bool = false
                                    v2 += (v2 | v2)
                                    _ = v15
                                    _ = v16
                                }
                                v1 = v1
                                delete(v4, fi13)
                                _ = fi14
                            }
                        }
                        for fi17 := 0; fi17 < 10; fi17++ {
                            v1 = v1
                            v18 := GenS2{F0: 42, F1: fi13, F2: 42}
                            _ = fi17
                            _ = v18
                        }
                        _ = fi13
                    }
                }
            case "foo bar":
                fmt.Println(v1, v1)
            case "a":
                switch "" {
                case "hello":
                    v2 += v2
                case "\n\t":
                    var v19 uint32 = 4294967295
                    _ = v19
                default:
                    v2 -= v2
                }
            default:
                for fi20 := 0; fi20 < 5; fi20++ {
                    if (!false) {
                        for fi21 := 0; fi21 < 6; fi21++ {
                            var v22 float64 = 3.14159
                            _ = fi21
                            _ = v22
                        }
                        v1 = v1
                        v1 = "\n\t"
                    }
                    _ = fi20
                }
            }
        }
        v2 -= 2
    }
    v1 = (((v1 + v1) + (v1 + ((((v1 + v1) + v1) + "foo bar") + (v1 + v1)))) + v1)
    _ = len(v10)
    v1 = v1
    for fi23 := 0; fi23 < 10; fi23++ {
        v2 += v2
        if (0 == 2) {
            if (false || (0.1 > 3.14159)) {
                switch "foo bar" {
                case "a":
                    v24 := GenS1{F0: -1, F1: 127}
                    _ = v24
                case "foo bar":
                    v2 -= (v2 % (((v2 - v2)) | 1))
                case "hello":
                    v2 -= v2
                }
                fmt.Println(v2, fi23)
                delete(v3, ("hello" + v1))
            }
            v1 = (v1 + v1)
        }
        if len(v10) > 0 {
            v10[0] = 4294967295
        }
        _ = fi23
    }
    switch ("a" + "hello") {
    case "foo bar":
        var v25 uint = 42
        _ = v25
    case "":
        fmt.Println(v2, v1)
    default:
        v1 = v1
    }
    var v26 int32 = 1
    v26 = v26
    if true {
        if false {
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 float32 = 1e10
                _ = fi27
                _ = v28
            }
            fmt.Println(v26, v2)
            for fi29 := 0; fi29 < 11; fi29++ {
                if (57 > 128) {
                    var v30 bool = true
                    if ((!v30) || v30) {
                        v31 := GenS1{F0: -78, F1: 84}
                        var v32 bool = ((4294967295 == 2147483648) || v30)
                        _ = v31
                        _ = v32
                    }
                    delete(v4, fi29)
                    _ = v30
                }
                for fi33 := 0; fi33 < 4; fi33++ {
                    v1 = v1
                    delete(v3, (("hello" + v1) + "\n\t"))
                    v3[v1] = fi29
                    if (true || (!true)) {
                        v34 := GenS2{F0: 2, F1: (fi33 << ((1) & 63)), F2: 307}
                        if len(v10) > 0 {
                            v10[0] = (^(fi29))
                        }
                        v26 = 42
                        _ = v34
                    }
                    var v35 uint8 = 255
                    _ = fi33
                    _ = v35
                }
                v36 := GenS3{F0: v2, F1: 127, F2: 94}
                _ = fi29
                _ = v36
            }
        } else {
            for fi37 := 0; fi37 < 9; fi37++ {
                v1 = v1
                v10 = append(v10, (^(fi37)))
                _ = fi37
            }
            if len(v10) > 0 {
                v10[0] = 432
            }
            v26 = (v26 % (((v26 >> ((4294967295) & 63))) | 1))
        }
        _ = len(v3)
    }
    var v38 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v38
    _ = v4
}
