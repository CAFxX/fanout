package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 uint16 }
func genFn0() int {
    v6 := []int{-1}
    if false {
        if false {
            var v7 int32 = -2147483648
            if (false && true) {
                if true {
                    v8 := GenS1{F0: 318, F1: ((false && false) || true), F2: 157}
                    _ = v8
                } else {
                    fmt.Println(v7)
                    var v9 uint64 = 905
                    v6 = append(v6, 1)
                    _ = v9
                }
                var v10 uint = 42
                _ = v10
            }
            v7 = v7
            _ = v7
        }
    } else {
        v6 = append(v6, 9223372036854775807)
        var v11 int32 = 42
        if false {
            if (!((!false) || false)) {
                v11 += (v11 + (v11 >> ((1) & 63)))
                for fi12 := 0; fi12 < 1; fi12++ {
                    if (1 != 678) {
                        if (9223372036854775807 > -9223372036854775808) {
                            for fi13 := 0; fi13 < 12; fi13++ {
                                v6 = append(v6, 563)
                                _ = fi13
                            }
                        } else {
                            for fi14 := 0; fi14 < 8; fi14++ {
                                switch (fi12 + (fi14 ^ 2147483647)) {
                                case -9223372036854775808:
                                    if len(v6) > 0 {
                                        v6[0] = fi14
                                    }
                                }
                                if true {
                                    var v15 int8 = 127
                                    var v16 int8 = v15
                                    _ = v15
                                    _ = v16
                                } else {
                                    var v17 int = 2147483647
                                    _ = v17
                                }
                                _ = fi14
                            }
                            fmt.Println(fi12)
                            switch fi12 {
                            case 9223372036854775807:
                                switch "foo bar" {
                                case "\n\t":
                                    for fi18 := 0; fi18 < 7; fi18++ {
                                        var v19 uint = 1
                                        _ = fi18
                                        _ = v19
                                    }
                                }
                            case 127:
                                v6 = append(v6, (^(fi12)))
                            case 2147483647:
                                var v20 float64 = 3.14159
                                _ = v20
                            }
                        }
                    }
                    if len(v6) > 0 {
                        v6[0] = 746
                    }
                    _ = fi12
                }
            } else {
                fmt.Println(v11)
                v11 = (v11 - (v11 << ((18446744073709551615) & 63)))
            }
        }
        _ = v11
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        switch ("a" + "hello") {
        case "hello":
            v22 := GenS2{F0: 65535, F1: 1}
            _ = v22
        case "":
            var v23 uint8 = 102
            _ = v23
        case "\n\t":
            var v24 uint8 = 128
            _ = v24
        }
        _ = fi21
    }
    _ = v6
    return 127
}

func main() {
    var v1 float32 = 0.0
    var v2 int16 = -32768
    var v3 bool = true
    var v4 int8 = 1
    v5 := []string{"\n\t", ("a" + ""), "hello"}
    v4 = v4
    v25 := GenS3{F0: genFn0(), F1: 32767, F2: 713}
    var v26 int8 = (v4 / (((v4 + (v4 * v4))) | 1))
    v1 -= v1
    var v27 int32 = 42
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    v1 = 1e10
    for fi28 := 0; fi28 < 5; fi28++ {
        v4 = 42
        if v3 {
            v27 += (^(v27))
            var v29 uint = 0
            for fi30 := 0; fi30 < 4; fi30++ {
                v26 += v26
                _ = fi30
            }
            _ = v29
        } else {
            v3 = v3
            var v31 float64 = 0.1
            var v32 int64 = -451
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    if (106 > v2) {
        for fi33 := 0; fi33 < 9; fi33++ {
            v34 := GenS3{F0: fi33, F1: 348, F2: 0}
            v26 += (v4 - -55)
            var v35 string = ((("a" + ("" + "hello")) + "hello") + "foo bar")
            _ = fi33
            _ = v34
            _ = v35
        }
        switch -903 {
        case -5:
            if (100 > -1) {
                for fi36 := 0; fi36 < 14; fi36++ {
                    v26 += v4
                    v3 = false
                    _ = len(v5)
                    _ = fi36
                }
            }
        }
        _ = len(v5)
    }
    var v37 bool = v3
    v26 = (v26 >> ((1000) & 63))
    var v38 int8 = 127
    if (0.0 < (1e10 / ((1.5 / -2.25) / v1))) {
        for fi39 := 0; fi39 < 18; fi39++ {
            var v40 float32 = 1e10
            _ = fi39
            _ = v40
        }
    }
    v27 -= (v27 * (v27 | v27))
    var v41 int64 = 0
    fmt.Println(v27)
    var v42 uint = 2
    var v43 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
