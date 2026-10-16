package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint; F2 int8 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int32 }

func main() {
    var v1 int = -2147483648
    v2 := []int{(v1 - v1), 767}
    var v3 int16 = -32768
    var v4 float64 = 0.0
    var v5 int16 = v3
    fmt.Println(v5)
    switch 9223372036854775807 {
    case -112:
        v4 = v4
    case 32768:
        fmt.Println(v4, v1)
    default:
        v6 := GenS1{F0: 1000}
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 bool = true
            v10 := GenS1{F0: (v5 % ((42) | 1))}
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := GenS2{F0: (!true), F1: 65535, F2: 0}
            if true {
                if false {
                    var v13 uint32 = 147
                    fmt.Println(fi7, fi11, v13)
                    _ = v13
                } else {
                    if true {
                        v1 -= 32768
                        fmt.Println(v5, fi7)
                    } else {
                        if (42 > 65535) {
                            var v14 bool = false
                            var v15 int16 = -228
                            v14 = (!v14)
                            _ = v14
                            _ = v15
                        }
                        v5 = v5
                    }
                    v1 = fi7
                    v16 := GenS1{F0: 535}
                    if true {
                        if true {
                            v17 := GenS3{F0: 1, F1: 0}
                            var v18 uint64 = 18446744073709551615
                            _ = v17
                            _ = v18
                        } else {
                            v2 = append(v2, 661)
                            v1 -= fi7
                        }
                        _ = len(v2)
                    }
                    fmt.Println(v3, fi11)
                    if (-1 >= -95) {
                        v1 += int(v3)
                        var v19 int32 = 0
                        _ = v19
                    }
                    v5 -= v3
                    _ = v16
                }
                var v20 uint8 = 127
                fmt.Println(fi7, fi7)
                _ = v20
            } else {
                switch "" {
                case "\n\t":
                    switch fi7 {
                    case 440:
                        v4 -= 0.0
                    case 127:
                        v1 -= fi7
                        switch "hello" {
                        case "\n\t":
                            _ = len(v2)
                        case "hello":
                            v21 := GenS4{F0: 42, F1: 32767}
                            _ = v21
                        default:
                            if false {
                                fmt.Println(v4)
                            } else {
                                fmt.Println(v3, fi7)
                                var v22 int32 = 2147483647
                                var v23 int8 = -48
                                _ = v22
                                _ = v23
                            }
                        }
                    default:
                        v1 = 672
                    }
                default:
                    v1 -= (v1 + -2147483648)
                    v5 = (v3 - 1000)
                }
                if len(v2) > 0 {
                    v2[0] = (-(v1))
                }
            }
            if false {
                fmt.Println(v3)
                v5 = 1
                if false {
                    var v24 int32 = -1
                    _ = v24
                }
            } else {
                v3 += v3
            }
            _ = fi11
            _ = v12
        }
        _ = fi7
    }
    v2 = append(v2, v1)
    var v25 uint32 = 65535
    if true {
        var v26 uint8 = 2
        v26 = v26
        if (!false) {
            switch "hello" {
            case "foo bar":
                var v27 int = -1
                _ = v27
            default:
                for fi28 := 0; fi28 < 2; fi28++ {
                    if (false || false) {
                        v1 -= (v1 | 0)
                        var v29 uint8 = 0
                        fmt.Println(v1)
                        _ = v29
                    }
                    _ = fi28
                }
            }
            _ = len(v2)
            var v30 uint = 42
            _ = v30
        }
        _ = v26
    } else {
        for fi31 := 0; fi31 < 13; fi31++ {
            v32 := GenS4{F0: 0, F1: 42}
            for fi33 := 0; fi33 < 15; fi33++ {
                var v34 string = (("a" + (("\n\t" + ("foo bar" + (("\n\t" + "a") + "a"))) + "hello")) + "")
                _ = fi33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        var v35 int16 = v5
        _ = v35
    }
    v3 = 42
    fmt.Println(v5, v5, v3)
    for fi36 := 0; fi36 < 5; fi36++ {
        v2 = append(v2, (v1 - v1))
        var v37 int32 = -382
        switch ("foo bar" + ("\n\t" + "a")) {
        case "\n\t":
            var v38 float32 = 1e10
            _ = v38
        case "":
            switch "hello" {
            case "a":
                var v39 float64 = float64(1.5)
                _ = v39
            case "\n\t":
                switch v1 {
                case -2147483648:
                    v40 := GenS5{F0: v25, F1: v37}
                    _ = v40
                case 127:
                    if false {
                        if true {
                            for fi41 := 0; fi41 < 13; fi41++ {
                                v42 := GenS5{F0: 79, F1: v37}
                                var v43 int = v1
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                            switch fi36 {
                            case -2147483648:
                                switch int(v5) {
                                case 42:
                                    v1 += fi36
                                case 32767:
                                    switch (v1 ^ (v1 / (((v1 << ((4294967295) & 63))) | 1))) {
                                    case -420:
                                        for fi44 := 0; fi44 < 14; fi44++ {
                                            var v45 string = "\n\t"
                                            var v46 string = "\n\t"
                                            _ = fi44
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                }
                            case 15:
                                var v47 int32 = 1
                                _ = v47
                            default:
                                var v48 float64 = 0.1
                                _ = v48
                            }
                        } else {
                            _ = len(v2)
                            fmt.Println(v1, v37)
                            v5 = v5
                        }
                    }
                default:
                    var v49 string = "a"
                    _ = v49
                }
            default:
                v25 += 0
            }
        }
        _ = fi36
        _ = v37
    }
    v25 += 65535
    v50 := GenS3{F0: 18446744073709551615, F1: 2}
    switch (v1 + v1) {
    case -393:
        var v51 uint8 = 1
        _ = v51
    case 32767:
        switch "foo bar" {
        case "a":
            var v52 int8 = 100
            _ = v52
        }
    default:
        var v53 int64 = 9223372036854775807
        _ = v53
    }
    v5 -= (v5 | v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}
