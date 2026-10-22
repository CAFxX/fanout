package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int64; F1 string }

func main() {
    var v1 int = 292
    v2 := []uint32{2}
    var v3 int = 594
    var v4 int32 = 2147483647
    var v5 uint64 = 42
    var v6 float64 = ((0.1 / 3.14159) / 1e10)
    if (!false) {
        switch ((("hello" + "") + "a") + "foo bar") {
        case "\n\t":
            if true {
                if true {
                    if false {
                        v7 := GenS1{F0: 127}
                        v8 := GenS1{F0: 100}
                        _ = v7
                        _ = v8
                    } else {
                        switch "a" {
                        case "a":
                            if len(v2) > 0 {
                                v2[0] = 863
                            }
                        }
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                }
                v9 := GenS2{F0: 1000, F1: ((!(((-2.25 / (v6 + v6)) * v6) < float64(0.0))) && (true && (int8(v5) < -128)))}
                var v10 int64 = -1
                _ = v10
                _ = v9
            }
        case "foo bar":
            fmt.Println(v6)
        case "hello":
            switch (v3 * int(v4)) {
            case 9223372036854775807:
                if true {
                    v2 = append(v2, 4294967295)
                    v4 = v4
                    _ = len(v2)
                } else {
                    if (false && true) {
                        var v11 bool = (((v5 == (v5 << ((42) & 63))) && true) || true)
                        var v12 int64 = 90
                        v13 := GenS2{F0: 1000, F1: (v4 == (v4 - v4))}
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        if false {
                            var v14 uint = 65535
                            _ = v14
                        }
                    }
                }
            case -312:
                v2 = append(v2, 1)
            }
        default:
            v1 += v3
        }
    } else {
        _ = len(v2)
        switch "foo bar" {
        case "a":
            var v15 int = -9223372036854775808
            _ = v15
        default:
            v16 := GenS2{F0: -1, F1: true}
            _ = v16
        }
        if true {
            var v17 uint16 = 65535
            for fi18 := 0; fi18 < 12; fi18++ {
                v6 = v6
                var v19 int8 = -1
                _ = fi18
                _ = v19
            }
            _ = v17
        }
    }
    v1 = v1
    v20 := GenS1{F0: -99}
    var v21 int = -23
    switch -585 {
    case -2147483648:
        for fi22 := 0; fi22 < 9; fi22++ {
            v6 += (v6 - ((((-2.25 / ((float64(0.1) - ((v6 * float64(0.0)) + v6)) / (v6 - ((v6 * (v6 - (v6 / v6))) / (float64(((-2.25 + 3.14159) - 1.5)) + v6))))) / float64(0.1)) + float64(1.5)) - 3.14159))
            for fi23 := 0; fi23 < 15; fi23++ {
                if len(v2) > 0 {
                    v2[0] = 312
                }
                var v24 int32 = (v4 / ((-805) | 1))
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    case 128:
        var v25 bool = false
        _ = v25
    }
    fmt.Println(v4, v6, v5)
    var v26 float64 = -2.25
    v6 = 3.14159
    for fi27 := 0; fi27 < 11; fi27++ {
        v5 += v5
        v28 := GenS1{F0: 1}
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v4, v1, v6)
    v2 = append(v2, 37)
    for fi29 := 0; fi29 < 10; fi29++ {
        var v30 uint32 = 681
        if true {
            for fi31 := 0; fi31 < 12; fi31++ {
                v4 = v4
                var v32 uint16 = 65535
                if false {
                    _ = len(v2)
                } else {
                    fmt.Println(v6)
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        var v33 string = "hello"
                        _ = v33
                    case "foo bar":
                        if (!true) {
                            var v34 int64 = 42
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v34
                        } else {
                            fmt.Println(fi31, v1)
                            fmt.Println(v32, v30, v32)
                            if true {
                                switch v1 {
                                case 42:
                                    if false {
                                        _ = len(v2)
                                        var v35 uint64 = uint64(v32)
                                        v21 += (fi31 & (^(v21)))
                                        var v36 string = (("hello" + (("\n\t" + (("hello" + "\n\t") + ("hello" + "foo bar"))) + "")) + (("" + "a") + "\n\t"))
                                        switch -9223372036854775808 {
                                        case -1:
                                            var v37 uint32 = v30
                                            _ = v37
                                        case 4294967295:
                                            switch v1 {
                                            case 4294967295:
                                                v2 = append(v2, (v30 >> ((42) & 63)))
                                            case 1:
                                                v32 += v32
                                            }
                                        case 1:
                                            v4 = v4
                                        }
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v21 = fi29
                                        v38 := GenS3{F0: -9223372036854775808, F1: "hello"}
                                        _ = v38
                                    }
                                default:
                                    _ = len(v2)
                                }
                            }
                        }
                    case "":
                        var v39 int64 = 9223372036854775807
                        _ = v39
                    }
                    v26 -= v6
                }
                _ = fi31
                _ = v32
            }
            for fi40 := 0; fi40 < 5; fi40++ {
                var v41 int8 = 1
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                _ = fi40
                _ = v41
            }
        }
        fmt.Println(v5, v3)
        _ = fi29
        _ = v30
    }
    var v42 int = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
