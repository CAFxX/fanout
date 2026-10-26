package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 int16 }

func main() {
    var v1 uint16 = 795
    var v2 float64 = float64((1.5 - float32(float64((float32((0.0 / -2.25)) * ((-2.25 / 1e10) / -2.25))))))
    var v3 uint8 = 127
    if (65535 == 0) {
        var v4 float32 = (1.5 / 1.0)
        var v5 int64 = 1
        v6 := []int{}
        _ = v4
        _ = v5
        _ = v6
    }
    if (0 < v3) {
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := GenS1{F0: v1, F1: true}
                var v10 uint8 = 127
                v10 = 140
                _ = fi8
                _ = v10
                _ = v9
            }
            if (1 == 4294967295) {
                v2 += (v2 / v2)
                fmt.Println(fi7)
            } else {
                v11 := []int{}
                var v12 uint = 1
                _ = v11
                _ = v12
            }
            _ = fi7
        }
    } else {
        var v13 int64 = 42
        _ = v13
    }
    switch 32767 {
    case 9223372036854775807:
        switch "a" {
        case "hello":
            switch (("" + (("a" + (("foo bar" + "foo bar") + "\n\t")) + "\n\t")) + "\n\t") {
            case "\n\t":
                fmt.Println(v3)
            case "hello":
                switch "" {
                case "":
                    switch "hello" {
                    case "foo bar":
                        var v14 int8 = 1
                        _ = v14
                    case "\n\t":
                        var v15 int = 0
                        _ = v15
                    default:
                        v16 := []int{4294967295}
                        _ = v16
                    }
                default:
                    var v17 bool = ((18446744073709551615 == 42) || (v3 == (v3 % ((v3) | 1))))
                    _ = v17
                }
            }
        case "":
            if true {
                for fi18 := 0; fi18 < 15; fi18++ {
                    v2 = 3.14159
                    switch "" {
                    case "":
                        if (((fi18 << ((4294967295) & 63)) != -2147483648) && (339 < 18446744073709551615)) {
                            v19 := GenS1{F0: (v1 >> ((65535) & 63)), F1: true}
                            var v20 int8 = 74
                            var v21 uint32 = 65535
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            for fi22 := 0; fi22 < 16; fi22++ {
                                switch -951 {
                                case 4294967295:
                                    switch fi18 {
                                    case 0:
                                        var v23 uint64 = 2
                                        _ = v23
                                    default:
                                        switch "hello" {
                                        case "hello":
                                            var v24 uint8 = v3
                                            _ = v24
                                        }
                                    }
                                default:
                                    if true {
                                        if true {
                                            if (!((true && (v1 != (^(v1)))) && (2147483647 == 691))) {
                                                v25 := []int{fi22, fi22}
                                                v1 += v1
                                                _ = v25
                                            }
                                            v1 -= (^(v1))
                                        } else {
                                            switch "\n\t" {
                                            case "":
                                                v26 := GenS1{F0: v1, F1: true}
                                                _ = v26
                                            case "\n\t":
                                                v3 = 42
                                            default:
                                                fmt.Println(fi18, fi18)
                                            }
                                        }
                                        fmt.Println(v2, v1)
                                    }
                                }
                                _ = fi22
                            }
                        }
                    case "\n\t":
                        for fi27 := 0; fi27 < 7; fi27++ {
                            v28 := GenS1{F0: 2, F1: false}
                            v29 := []int{fi18, fi27}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    default:
                        v30 := GenS1{F0: (v1 << ((42) & 63)), F1: ("" == "\n\t")}
                        _ = v30
                    }
                    if true {
                        var v31 uint = 712
                        v31 += v31
                        _ = v31
                    }
                    _ = fi18
                }
                v2 = 1e10
            }
        }
    default:
        v32 := GenS1{F0: (v1 / ((v1) | 1)), F1: (true && (float64(1.5) >= (v2 / -2.25)))}
        _ = v32
    }
    if false {
        fmt.Println(v1, v1, v2)
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        fmt.Println(fi33, v2)
        _ = fi33
    }
    var v34 uint64 = 1
    v35 := GenS1{F0: (v1 % ((v1) | 1)), F1: (-80 > 100)}
    v36 := []int{9223372036854775807, 192, 1}
    v37 := GenS2{F0: v1, F1: 2147483647, F2: 133}
    if true {
        v1 = 134
        if len(v36) > 0 {
            v36[0] = 32768
        }
        v3 = v3
    }
    v38 := GenS2{F0: v1, F1: 4294967295, F2: -32768}
    _ = len(v36)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
