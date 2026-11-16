package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 bool; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 int8 }

func main() {
    var v1 bool = (0.0 != (0.0 * ((0.1 + -2.25) - (3.14159 - 1.5))))
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    v4[("" + ("\n\t" + "a"))] = ""
    var v5 uint32 = 2147483648
    if (!false) {
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 bool = (!(((^(fi6)) < fi6) && v1))
            v1 = (v7 && (33 != 656))
            for fi8 := 0; fi8 < 2; fi8++ {
                switch "\n\t" {
                case "\n\t":
                    if (!(float64((1e10 * 0.0)) <= (1e10 * 1e10))) {
                        var v9 uint = 4294967295
                        if (((v5 % ((v5) | 1)) <= (v5 * 4294967295)) || v1) {
                            v5 -= v5
                            fmt.Println(v5, v5)
                            switch ("" + "") {
                            case "a":
                                v7 = ("hello" > "")
                            }
                        }
                        _ = v9
                    }
                default:
                    v10 := GenS1{F0: "a"}
                    _ = v10
                }
                v1 = v1
                fmt.Println(fi8)
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
        var v11 int16 = 32767
        _ = v11
    } else {
        v12 := []int{0, 128, 128, 127}
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint16 = 0
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 int32 = -1
                v18 := GenS1{F0: "hello"}
                var v19 bool = ((((false || v1) || v1) && true) || false)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi15
        }
        _ = v12
    }
    if ((0.0 * (float64(3.14159) / 3.14159)) > 0.0) {
        if (uint(v5) > 42) {
            var v20 int8 = 88
            v21 := GenS1{F0: "\n\t"}
            v5 -= v5
            _ = v20
            _ = v21
        } else {
            v22 := []int{}
            _ = v22
        }
        v23 := GenS2{F0: 255, F1: 0}
        v24 := GenS3{F0: 2147483647}
        _ = v23
        _ = v24
    }
    v25 := v2[(v5 & v5)]
    var v26 uint64 = 916
    v27 := GenS1{F0: ("" + ("\n\t" + ("\n\t" + (("a" + "\n\t") + "\n\t"))))}
    if (v25 || (((v1 || v25) || (0 < 1)) && (929 == 4294967295))) {
        v28 := GenS1{F0: "\n\t"}
        _ = v28
    } else {
        switch "" {
        case "\n\t":
            fmt.Println(v25, v1)
        case "":
            fmt.Println(v25, v5, v5)
        case "foo bar":
            if (!v1) {
                fmt.Println(v25)
                var v29 uint16 = 54
                v30 := GenS4{F0: v1, F1: (255 <= uint16(v5)), F2: (v26 << ((817) & 63))}
                _ = v29
                _ = v30
            }
        default:
            var v31 uint32 = (v5 | v5)
            _ = v31
        }
        var v32 float32 = (3.14159 + float32(1e10))
        _ = v32
    }
    switch "" {
    case "hello":
        v33 := v4[("\n\t" + "")]
        _ = v33
    }
    var v34 float64 = 1.5
    var v35 uint8 = 1
    if v1 {
        for fi36 := 0; fi36 < 9; fi36++ {
            var v37 float64 = v34
            _ = fi36
            _ = v37
        }
        fmt.Println(v34, v35)
    }
    if v25 {
        for fi38 := 0; fi38 < 6; fi38++ {
            v1 = false
            _ = fi38
        }
        fmt.Println(v34, v26, v35)
    }
    v39 := GenS5{F0: ((v35 >> ((816) & 63)) < (v35 | v35))}
    for fi40 := 0; fi40 < 18; fi40++ {
        v5 = (v5 ^ v5)
        v41 := GenS6{F0: 1, F1: 100}
        _ = fi40
        _ = v41
    }
    v42 := GenS6{F0: 848, F1: -63}
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v42
    _ = v5
}
