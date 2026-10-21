package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 45
    var v3 int16 = 42
    var v4 int32 = 2147483647
    var v5 uint32 = 2147483648
    var v6 int16 = -18
    for fi7 := 0; fi7 < 20; fi7++ {
        switch (fi7 + (fi7 & (fi7 * (fi7 ^ 32767)))) {
        case -2147483648:
            v8 := GenS1{F0: 42, F1: (fi7 | (fi7 | (fi7 ^ int(v5))))}
            _ = v8
        case 32768:
            fmt.Println(v4, fi7)
        }
        _ = fi7
    }
    v4 -= -2147483648
    v2 -= v2
    v9 := GenS2{F0: 65535}
    fmt.Println(v4, v3)
    v6 = (v3 % ((1000) | 1))
    fmt.Println(v6)
    var v10 int8 = 1
    switch 128 {
    case 1:
        if (v2 >= (v2 ^ 1)) {
            if false {
                v5 += 0
                v1[("" + "a")] = (v5 / ((v5) | 1))
                v10 += -42
            }
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int = fi11
                v13 := []int{-977, int(v2)}
                _ = fi11
                _ = v12
                _ = v13
            }
        } else {
            var v14 int = 42
            _ = v14
        }
    case -2147483648:
        fmt.Println(v10)
    case 2147483647:
        switch "foo bar" {
        case "":
            if true {
                v6 = (^(v3))
                v15 := []int{32767}
                _ = v15
            }
        case "hello":
            for fi16 := 0; fi16 < 12; fi16++ {
                delete(v1, "a")
                fmt.Println(v4)
                switch "foo bar" {
                case "hello":
                    v17 := []int{(fi16 + 127)}
                    _ = v17
                case "a":
                    v1["foo bar"] = v5
                }
                _ = fi16
            }
        case "a":
            v6 += 32767
        default:
            v6 -= 1000
        }
    default:
        switch "hello" {
        case "hello":
            var v18 bool = (0.0 == 0.1)
            _ = v18
        default:
            if (!true) {
                if true {
                    var v19 uint16 = 148
                    for fi20 := 0; fi20 < 3; fi20++ {
                        switch fi20 {
                        case -407:
                            if (1 >= 42) {
                                v3 = (v3 + (v6 % ((v6) | 1)))
                                for fi21 := 0; fi21 < 8; fi21++ {
                                    fmt.Println(v2, v4)
                                    v22 := GenS2{F0: 4294967295}
                                    var v23 bool = ((v10 % ((-1) | 1)) < (v10 & v10))
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                            } else {
                                switch (fi20 ^ fi20) {
                                case 128:
                                    v19 = (v19 / (((v19 & v19)) | 1))
                                case -461:
                                    fmt.Println(v4, v2, v5)
                                case 1:
                                    v24 := []int{(fi20 << ((42) & 63)), int(v5), (fi20 | (-(fi20)))}
                                    _ = v24
                                }
                            }
                        case 954:
                            switch "foo bar" {
                            case "":
                                v4 = 2147483647
                            case "a":
                                var v25 string = ""
                                _ = v25
                            case "foo bar":
                                var v26 uint32 = 1
                                _ = v26
                            }
                        }
                        v3 = (^(v3))
                        v10 = -128
                        _ = fi20
                    }
                    var v27 float64 = -2.25
                    _ = v19
                    _ = v27
                } else {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v3)
                    }
                }
                for fi28 := 0; fi28 < 15; fi28++ {
                    if false {
                        if (((^(v10)) == (v10 % (((^(v10))) | 1))) || true) {
                            v6 = v3
                            v29 := GenS1{F0: 65535, F1: fi28}
                            _ = v29
                        } else {
                            v30 := GenS1{F0: 62, F1: fi28}
                            _ = v30
                        }
                        if (true && (fi28 <= fi28)) {
                            v6 = -734
                        }
                    }
                    _ = fi28
                }
            }
        }
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        v3 += v3
        _ = fi31
    }
    if false {
        v32 := GenS2{F0: 2}
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
