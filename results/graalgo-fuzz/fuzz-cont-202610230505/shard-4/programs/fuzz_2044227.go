package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int64; F2 bool }
type GenS3 struct { F0 int; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint }
type GenS5 struct { F0 int64 }

func main() {
    v1 := []int{128}
    var v2 uint8 = 42
    v3 := make(map[string]uint32)
    v2 += (v2 & v2)
    if (true && true) {
        if (2147483647 < -1) {
            v4 := GenS1{F0: 32767}
            v5 := GenS1{F0: 9223372036854775807}
            _ = v4
            _ = v5
        } else {
            if (!false) {
                for fi6 := 0; fi6 < 13; fi6++ {
                    v2 = 2
                    v2 = 127
                    _ = fi6
                }
            }
            v2 = 255
            delete(v3, "hello")
        }
        var v7 uint32 = 2147483648
        _ = v7
    } else {
        switch "" {
        case "":
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS2{F0: 4294967295, F1: 9223372036854775807, F2: (("hello" <= "hello") || true)}
                _ = fi8
                _ = v9
            }
        default:
            var v10 string = ("foo bar" + ("a" + ("\n\t" + "foo bar")))
            _ = v10
        }
        v2 = (v2 & v2)
    }
    v11 := v3["a"]
    if (((42 < 2) || false) || false) {
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                if true {
                    switch (fi13 & 127) {
                    case 4294967295:
                        var v14 bool = false
                        _ = v14
                    }
                    v15 := GenS2{F0: 42, F1: -145, F2: true}
                    var v16 uint32 = v11
                    v16 = (v11 & (v16 ^ v16))
                    v16 = (^(v16))
                    var v17 uint16 = 65535
                    v16 += v16
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    var v18 bool = true
                    if v18 {
                        v19 := GenS3{F0: 32767, F1: v18, F2: 18446744073709551615}
                        _ = v19
                    } else {
                        v2 = (^(v2))
                    }
                    var v20 int = -1
                    _ = v18
                    _ = v20
                }
                v1 = append(v1, -2147483648)
                _ = fi13
            }
            _ = fi12
        }
        switch "\n\t" {
        case "foo bar":
            v2 += 2
        case "hello":
            fmt.Println(v11, v11)
        default:
            v21 := GenS2{F0: 42, F1: -9223372036854775808, F2: false}
            _ = v21
        }
        v22 := GenS4{F0: 1, F1: 2, F2: 4294967295}
        _ = v22
    } else {
        if (!false) {
            v2 = (v2 + (v2 % ((2) | 1)))
            fmt.Println(v11)
        }
        var v23 string = "a"
        _ = v23
    }
    if (v11 > v11) {
        _ = len(v3)
        for fi24 := 0; fi24 < 12; fi24++ {
            v2 -= (v2 / ((127) | 1))
            var v25 float64 = 0.0
            _ = fi24
            _ = v25
        }
    } else {
        if false {
            v2 = v2
            v11 = v11
        } else {
            v1 = append(v1, 2147483647)
            v26 := GenS3{F0: -1, F1: (((!(542 >= 458)) && (0.1 >= (1.5 / (0.1 / float64(-2.25))))) && ((v2 > v2) && false)), F2: 751}
            _ = v26
        }
        if (9223372036854775807 > 886) {
            fmt.Println(v11, v11)
            v2 = v2
        } else {
            fmt.Println(v11, v11)
            if true {
                v27 := GenS1{F0: 0}
                for fi28 := 0; fi28 < 10; fi28++ {
                    _ = len(v1)
                    v29 := GenS4{F0: 0, F1: 18446744073709551615, F2: 65535}
                    _ = fi28
                    _ = v29
                }
                if (627 <= 1) {
                    var v30 bool = false
                    if false {
                        if v30 {
                            v31 := GenS3{F0: 0, F1: v30, F2: 1}
                            _ = v31
                        } else {
                            fmt.Println(v2, v2, v11)
                            v32 := v3[(((("" + "") + "a") + "hello") + (("hello" + "") + "hello"))]
                            _ = v32
                        }
                        if (!false) {
                            _ = len(v1)
                            switch "" {
                            case "":
                                v30 = (42 <= 206)
                            case "hello":
                                var v33 bool = (504 <= 4294967295)
                                _ = v33
                            default:
                                var v34 int64 = -781
                                _ = v34
                            }
                        }
                    } else {
                        if true {
                            fmt.Println(v11, v2, v11)
                            v1 = append(v1, 855)
                        }
                        if v30 {
                            v2 -= v2
                            var v35 string = ("hello" + "\n\t")
                            switch 2147483647 {
                            case -2147483648:
                                v36 := GenS2{F0: 803, F1: 9223372036854775807, F2: ((v11 / ((v11) | 1)) == v11)}
                                _ = v36
                            default:
                                if false {
                                    v37 := GenS1{F0: 2147483647}
                                    var v38 int16 = 32767
                                    _ = v37
                                    _ = v38
                                } else {
                                    var v39 uint = 4294967295
                                    _ = v39
                                }
                            }
                            _ = v35
                        } else {
                            var v40 int = -897
                            _ = v40
                        }
                        fmt.Println(v2, v2)
                    }
                    v2 = v2
                    _ = v30
                } else {
                    var v41 int8 = -128
                    var v42 float64 = (0.0 + 1e10)
                    _ = v41
                    _ = v42
                }
                _ = v27
            } else {
                var v43 uint = 18446744073709551615
                switch -9223372036854775808 {
                case 361:
                    var v44 uint64 = 0
                    _ = v44
                }
                var v45 bool = ((^(v11)) != 726)
                _ = v43
                _ = v45
            }
        }
        v46 := GenS2{F0: 18446744073709551615, F1: 1, F2: false}
        _ = v46
    }
    var v47 uint32 = v11
    if false {
        v47 = v47
        v48 := GenS5{F0: 0}
        v47 = (v11 + v11)
        _ = v48
    }
    v1 = append(v1, -925)
    var v49 float32 = (3.14159 * float32(1e10))
    v47 -= (v47 | (v47 - (v47 & 42)))
    v2 -= (v2 ^ (v2 - (^(v2))))
    if (!true) {
        fmt.Println(v49, v49)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v47
    _ = v49
}
