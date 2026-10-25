package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint16; F1 int16; F2 int8 }
type GenS3 struct { F0 uint; F1 int8; F2 string }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"a", (("foo bar" + ("foo bar" + ("" + "foo bar"))) + ""), ("" + ((("" + (("a" + ("a" + "")) + "hello")) + "") + ("\n\t" + "foo bar"))), "hello"}
    var v3 int64 = 9223372036854775807
    v4 := GenS1{F0: (4294967295 == 2147483648), F1: 1, F2: 42}
    if (!false) {
        var v5 uint8 = 210
        v6 := GenS1{F0: true, F1: 42, F2: -97}
        _ = v5
        _ = v6
    } else {
        if (1 != 1) {
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := GenS1{F0: false, F1: 1, F2: 127}
                switch "hello" {
                case "":
                    v3 += (v3 - v3)
                default:
                    if (902 >= 42) {
                        fmt.Println(fi7, v3)
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                    } else {
                        v9 := GenS2{F0: 814, F1: 42, F2: 0}
                        _ = v9
                    }
                }
                _ = fi7
                _ = v8
            }
            var v10 int16 = -1
            _ = v10
        }
        delete(v1, 65535)
        v11 := GenS3{F0: 255, F1: 117, F2: "hello"}
        _ = v11
    }
    v3 = -1
    switch "hello" {
    case "\n\t":
        if true {
            v12 := GenS3{F0: 65535, F1: 0, F2: ("a" + ("\n\t" + ("\n\t" + "\n\t")))}
            v13 := GenS2{F0: 851, F1: -32768, F2: -128}
            _ = v12
            _ = v13
        }
    case "hello":
        var v14 int = 350
        _ = v14
    }
    if (1 > -1) {
        var v15 float64 = 0.0
        _ = v15
    } else {
        v16 := GenS1{F0: (((float64(((1e10 * 1e10) - 0.1)) + 3.14159) - float64(float32(float64(((0.0 * (((float32(0.1) * 3.14159) / (0.1 * float32((0.0 - (1.5 / 3.14159))))) / 1.5)) + 0.1))))) <= (0.0 - 1e10)), F1: 4294967295, F2: 97}
        if false {
            if false {
                v17 := GenS3{F0: 1, F1: 1, F2: ((("a" + ("a" + "foo bar")) + "hello") + "foo bar")}
                v3 += (v3 & v3)
                _ = v17
            }
            var v18 string = "\n\t"
            _ = v18
        } else {
            v19 := GenS1{F0: ((!(0 == 42)) && false), F1: 42, F2: 1}
            var v20 int = 0
            _ = v19
            _ = v20
        }
        fmt.Println(v3)
        _ = v16
    }
    var v21 uint = 65535
    if (!(!false)) {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 int = -2147483648
            _ = fi22
            _ = v23
        }
    } else {
        v24 := GenS2{F0: 42, F1: 32767, F2: 127}
        _ = v24
    }
    if (65535 > 42) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v2)
    }
    v25 := GenS1{F0: (false && (-2147483648 <= 349)), F1: 42, F2: 42}
    var v26 int16 = 647
    fmt.Println(v21, v21, v21)
    fmt.Println(v26, v3, v21)
    if ((^(v3)) <= (v3 << (((v21 & v21)) & 63))) {
        switch "foo bar" {
        case "hello":
            var v27 float64 = 1.5
            _ = v27
        default:
            for fi28 := 0; fi28 < 15; fi28++ {
                v21 -= (v21 * v21)
                var v29 uint64 = 1
                v30 := GenS4{F0: 514, F1: true}
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        v31 := GenS5{F0: -2147483648, F1: v21}
        _ = v31
    }
    if (2147483647 < 127) {
        var v32 string = "a"
        var v33 bool = false
        _ = v32
        _ = v33
    } else {
        v34 := GenS2{F0: 255, F1: (v26 | v26), F2: 1}
        _ = v34
    }
    v3 = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
