package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int32 }
func genFn0() string {
    var v6 string = "a"
    var v7 uint64 = 1
    var v8 int = 548
    switch ("a" + ((v6 + v6) + "\n\t")) {
    case "hello":
        if true {
            v6 = (((v6 + "") + (("\n\t" + (v6 + "")) + v6)) + "a")
        } else {
            v6 = v6
            v8 += (v8 / ((v8) | 1))
        }
    case "\n\t":
        v9 := []int{(v8 << ((560) & 63))}
        _ = v9
    default:
        v10 := []int{(v8 + v8), (v8 & v8)}
        _ = v10
    }
    v11 := []int{(-(v8)), v8, (v8 - v8)}
    _ = v11
    _ = v6
    _ = v7
    _ = v8
    return ((v6 + "\n\t") + v6)
}

func main() {
    var v1 uint = 354
    var v2 int8 = 100
    v3 := make(map[int]uint32)
    var v4 bool = (true && (false || true))
    v5 := []string{"\n\t", ("a" + (("" + ("foo bar" + ("\n\t" + "a"))) + (("" + (("foo bar" + "") + "hello")) + ("\n\t" + "\n\t")))), ("foo bar" + ("hello" + "\n\t")), "hello"}
    switch "hello" {
    case "foo bar":
        var v12 uint = (v1 / ((v1) | 1))
        _ = v12
    case "":
        switch -532 {
        case 4294967295:
            if false {
                v13 := v3[32768]
                var v14 float64 = (3.14159 * 1e10)
                fmt.Println(v13, v14, v2)
                _ = v13
                _ = v14
            } else {
                switch 12 {
                case 2147483647:
                    fmt.Println(v2, v4, v2)
                default:
                    fmt.Println(v1)
                }
                var v15 uint16 = 107
                _ = v15
            }
        case -9223372036854775808:
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS1{F0: 65535, F1: 0, F2: 0}
                if v4 {
                    switch (fi16 / ((fi16) | 1)) {
                    case 32767:
                        v18 := GenS1{F0: 42, F1: 255, F2: 9223372036854775807}
                        _ = v18
                    case -863:
                        var v19 uint64 = 1
                        _ = v19
                    case 42:
                        v20 := GenS1{F0: 42, F1: 127, F2: 779}
                        _ = v20
                    default:
                        if (1e10 < 1.5) {
                            if true {
                                v2 -= 42
                                switch 128 {
                                case 2147483647:
                                    v4 = (2147483647 > 93)
                                case 1:
                                    var v21 uint8 = 128
                                    _ = v21
                                case -2147483648:
                                    switch (fi16 ^ fi16) {
                                    case 42:
                                        v1 += (v1 + v1)
                                    case 0:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                default:
                                    for fi22 := 0; fi22 < 12; fi22++ {
                                        v5 = append(v5, "\n\t")
                                        var v23 string = "\n\t"
                                        if true {
                                            v1 = 54
                                            if (128 >= 1) {
                                                var v24 int32 = 42
                                                _ = v24
                                            }
                                            switch fi22 {
                                            case 4294967295:
                                                var v25 uint16 = 130
                                                _ = v25
                                            }
                                        }
                                        var v26 bool = (999 >= fi22)
                                        _ = fi22
                                        _ = v23
                                        _ = v26
                                    }
                                }
                                v27 := GenS2{F0: -434}
                                _ = v27
                            }
                        } else {
                            var v28 float64 = 1.5
                            v29 := GenS2{F0: 0}
                            _ = v28
                            _ = v29
                        }
                    }
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v1)
                    case "a":
                        var v30 uint8 = 2
                        _ = v30
                    case "hello":
                        for fi31 := 0; fi31 < 18; fi31++ {
                            v2 += (v2 - 85)
                            fmt.Println(fi31)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi31
                        }
                    }
                    v2 -= v2
                } else {
                    var v32 int = 128
                    delete(v3, v32)
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v34 := GenS3{F0: v2, F1: v2}
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                }
                switch 873 {
                case 1:
                    v4 = v4
                case -113:
                    _ = len(v5)
                }
                _ = fi16
                _ = v17
            }
        default:
            v35 := GenS4{F0: 801}
            _ = v35
        }
    default:
        var v36 bool = v4
        _ = v36
    }
    var v37 uint16 = 1
    fmt.Println(v1, v4, v4)
    var v38 uint64 = 1
    switch ("foo bar" + "") {
    case "hello":
        var v39 uint16 = (v37 / (((v37 / (((v37 % ((v37) | 1))) | 1))) | 1))
        _ = v39
    case "a":
        for fi40 := 0; fi40 < 11; fi40++ {
            v41 := GenS5{F0: v1, F1: 42}
            if (-1 < 2147483647) {
                v1 = (v1 >> (((^(v1))) & 63))
            } else {
                if len(v5) > 0 {
                    v5[0] = ("\n\t" + "hello")
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v42 := GenS3{F0: 1, F1: (v2 >> (((v1 >> ((v1) & 63))) & 63))}
                _ = v42
            }
            _ = fi40
            _ = v41
        }
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if v4 {
        fmt.Println(v38, v37, v37)
    }
    switch 128 {
    case -2147483648:
        fmt.Println(v1, v2, v1)
    case 2147483647:
        v37 -= (^(v37))
    case 1:
        var v43 int = 1
        _ = v43
    default:
        v44 := GenS4{F0: (v37 | v37)}
        _ = v44
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi45 := 0; fi45 < 11; fi45++ {
        var v46 uint32 = 310
        var v47 int = -746
        _ = fi45
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
