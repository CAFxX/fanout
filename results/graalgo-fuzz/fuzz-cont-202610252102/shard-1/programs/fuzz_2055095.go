package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint }
type GenS5 struct { F0 uint }
func genFn0() string {
    switch 127 {
    case -1:
        var v6 float64 = 0.0
        _ = v6
    case 868:
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: 42}
            _ = fi7
            _ = v8
        }
    case -9223372036854775808:
        var v9 string = ((("" + "foo bar") + "\n\t") + ("a" + "a"))
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := []int{(fi10 + (fi10 | (^(fi10)))), 128, -1}
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 float64 = (3.14159 / 1.5)
            v13 = (((v13 / v13) * ((-2.25 - (float64(float32(-2.25)) / 1.0)) + (v13 * float64((1.5 * -2.25))))) / (v13 - -2.25))
            v14 := GenS2{F0: 255, F1: ((0.0 * 0.1) == 0.1)}
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS3{F0: 88, F1: 0}
        _ = fi10
        _ = v11
        _ = v15
    }
    if (0 <= 1) {
        for fi16 := 0; fi16 < 15; fi16++ {
            if ((float32((-2.25 - 1.5)) - 0.1) == 0.1) {
                v17 := []int{(fi16 << ((331) & 63)), fi16, -1, 42}
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = v17
            } else {
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(fi16)
                }
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS1{F0: 1}
                    _ = fi18
                    _ = v19
                }
            }
            v20 := []int{(fi16 % ((9223372036854775807) | 1)), fi16, fi16}
            _ = fi16
            _ = v20
        }
        v21 := []int{42, 127, -9223372036854775808}
        var v22 int = 2147483647
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: 6}
    _ = v23
    return ""
}
func genFn1(p0 int, p1 bool, p2 uint16) bool {
    if false {
        p1 = p1
        var v24 uint8 = 2
        _ = v24
    } else {
        var v25 int32 = 560
        var v26 int32 = v25
        _ = v25
        _ = v26
    }
    var v27 uint32 = 136
    var v28 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    return (v27 < v27)
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 170
    var v3 int32 = -1
    var v4 uint16 = 166
    var v5 int64 = 1
    v2 = v2
    for fi29 := 0; fi29 < 18; fi29++ {
        if true {
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 int8 = 100
                var v32 int64 = int64(fi30)
                switch ("" + "") {
                case "":
                    var v33 uint8 = 127
                    _ = v33
                case "hello":
                    v4 += (v4 & (v4 - v4))
                case "\n\t":
                    fmt.Println(fi29, v5)
                default:
                    var v34 int = (^(v2))
                    _ = v34
                }
                v35 := v1["foo bar"]
                _ = fi30
                _ = v31
                _ = v32
                _ = v35
            }
            v36 := GenS4{F0: 18446744073709551615, F1: 18446744073709551615, F2: 18446744073709551615}
            _ = v36
        } else {
            v4 = v4
        }
        v37 := GenS4{F0: 2, F1: 0, F2: 0}
        _ = fi29
        _ = v37
    }
    v2 = v2
    for fi38 := 0; fi38 < 11; fi38++ {
        delete(v1, "a")
        fmt.Println(v3, v5)
        _ = fi38
    }
    v39 := GenS5{F0: 18446744073709551615}
    var v40 int = v2
    var v41 int64 = -9223372036854775808
    var v42 string = "foo bar"
    v41 = v5
    var v43 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
