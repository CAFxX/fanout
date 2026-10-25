package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int32 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int32; F1 int32; F2 uint64 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint64; F1 int64; F2 bool }
type GenS6 struct { F0 int16; F1 string }
func genFn0(p0 int32, p1 bool) int8 {
    v7 := []int{582, 1}
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return -13
}

func main() {
    var v1 int8 = 57
    var v2 uint = 694
    var v3 int = -193
    var v4 float32 = float32(0.0)
    var v5 uint32 = 2147483647
    var v6 uint16 = 903
    switch (("\n\t" + "") + "foo bar") {
    case "\n\t":
        var v8 uint = v2
        _ = v8
    case "":
        switch 2147483647 {
        case 128:
            if false {
                var v9 uint32 = v5
                _ = v9
            }
        case 561:
            v2 = (v2 - 712)
        }
    case "hello":
        if false {
            if false {
                fmt.Println(v5, v2)
                if true {
                    v6 += v6
                    var v10 int16 = 0
                    v11 := GenS1{F0: 1, F1: 475, F2: -456}
                    _ = v10
                    _ = v11
                }
            }
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 float64 = float64(-2.25)
                _ = fi12
                _ = v13
            }
            var v14 uint8 = 255
            _ = v14
        } else {
            fmt.Println(v1, v4, v4)
            if false {
                fmt.Println(v3, v2)
            } else {
                var v15 uint16 = uint16(v2)
                _ = v15
            }
            if (!(0 > -507)) {
                fmt.Println(v2)
            } else {
                for fi16 := 0; fi16 < 19; fi16++ {
                    var v17 int8 = -121
                    var v18 uint16 = (v6 * (v6 ^ 2))
                    v19 := GenS2{F0: -737, F1: (^(v5)), F2: 143}
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
            }
        }
    }
    v2 -= v2
    switch "a" {
    case "foo bar":
        var v20 uint64 = 261
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := GenS3{F0: -743, F1: -873, F2: 42}
        v23 := []int{(fi21 % ((fi21) | 1)), fi21, (^(v3)), fi21}
        v23 = append(v23, (^(fi21)))
        _ = fi21
        _ = v22
        _ = v23
    }
    if false {
        if (false || true) {
            v3 -= v3
            fmt.Println(v4, v4)
            var v24 uint = v2
            _ = v24
        } else {
            v1 -= -1
            v6 += (v6 % ((v6) | 1))
            var v25 float64 = -2.25
            _ = v25
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            v27 := []int{-44, v3, fi26, v3}
            var v28 float32 = 1e10
            _ = fi26
            _ = v27
            _ = v28
        }
        if (18446744073709551615 != 679) {
            switch ("foo bar" + "foo bar") {
            case "hello":
                v29 := []int{(-(v3))}
                _ = v29
            case "":
                if false {
                    v1 += 0
                    v2 += (v2 ^ (v2 + 502))
                }
            case "\n\t":
                fmt.Println(v1)
            default:
                for fi30 := 0; fi30 < 7; fi30++ {
                    v6 -= (v6 ^ 32767)
                    for fi31 := 0; fi31 < 14; fi31++ {
                        var v32 int8 = (^(v1))
                        _ = fi31
                        _ = v32
                    }
                    _ = fi30
                }
            }
            for fi33 := 0; fi33 < 18; fi33++ {
                if (4294967295 < 804) {
                    var v34 uint16 = (^(v6))
                    _ = v34
                }
                v35 := []int{v3, (v3 / ((v3) | 1)), v3, (fi33 >> ((v2) & 63))}
                v36 := GenS4{F0: "hello", F1: 18446744073709551615}
                _ = fi33
                _ = v35
                _ = v36
            }
        }
    }
    v37 := GenS5{F0: 2, F1: 9223372036854775807, F2: true}
    var v38 float32 = v4
    v2 -= v2
    for fi39 := 0; fi39 < 1; fi39++ {
        for fi40 := 0; fi40 < 14; fi40++ {
            v38 = v4
            v41 := []int{-170, (fi40 * fi39), 2147483647, (v3 & (v3 - (fi40 + -9223372036854775808)))}
            var v42 uint16 = v6
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = fi39
    }
    v43 := []int{}
    v44 := GenS1{F0: -1, F1: 307, F2: 0}
    v45 := GenS6{F0: -32768, F1: "hello"}
    var v46 string = "a"
    var v47 int = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
