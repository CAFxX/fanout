package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 string }
type GenS4 struct { F0 uint; F1 uint8; F2 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 uint64; F2 bool }
type GenS7 struct { F0 int }
func genFn0(p0 uint64) uint32 {
    fmt.Println(p0)
    fmt.Println(p0)
    var v4 int32 = 1
    _ = p0
    _ = v4
    return 2
}
func genFn1(p0 bool, p1 int8, p2 int) int16 {
    v5 := []int{p2}
    fmt.Println(p1, p1)
    switch p2 {
    case 128:
        var v6 int64 = -9223372036854775808
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -350
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]bool)
    var v3 float32 = float32(-2.25)
    v7 := GenS1{F0: ("" + "a"), F1: 42}
    switch "a" {
    case "foo bar":
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: "\n\t", F1: 42}
            _ = fi8
            _ = v9
        }
    case "\n\t":
        if false {
            v10 := GenS2{F0: 42, F1: 0, F2: 4294967295}
            var v11 uint16 = 255
            var v12 bool = true
            _ = v10
            _ = v11
            _ = v12
        }
    case "hello":
        v13 := []int{}
        _ = v13
    }
    v3 = -2.25
    fmt.Println(v1)
    var v14 uint8 = 42
    var v15 bool = (!true)
    fmt.Println(v14)
    var v16 uint64 = 863
    var v17 float64 = 0.0
    switch "hello" {
    case "":
        v18 := GenS3{F0: "foo bar", F1: 76, F2: ("\n\t" + "foo bar")}
        _ = v18
    case "a":
        var v19 uint64 = v16
        _ = v19
    default:
        var v20 int64 = -766
        _ = v20
    }
    v16 += v16
    if (!v15) {
        v21 := GenS1{F0: "", F1: 2}
        v22 := GenS4{F0: 0, F1: (v14 >> ((675) & 63)), F2: 9223372036854775807}
        v1 = v1
        _ = v21
        _ = v22
    } else {
        if v15 {
            switch 4294967295 {
            case 32768:
                switch "" {
                case "":
                    var v23 int8 = 127
                    _ = v23
                case "\n\t":
                    v24 := GenS5{F0: 16}
                    _ = v24
                case "a":
                    switch ("hello" + "hello") {
                    case "":
                        v2["foo bar"] = (true || true)
                    default:
                        for fi25 := 0; fi25 < 5; fi25++ {
                            v26 := GenS4{F0: 647, F1: v14, F2: 42}
                            v27 := GenS6{F0: (v16 & v16), F1: v16, F2: ((!("\n\t" == "\n\t")) || (v15 || v15))}
                            var v28 float32 = ((((v3 - (0.1 + v3)) * (float32(v17) - float32(-2.25))) - v3) * 0.0)
                            _ = fi25
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    }
                }
            default:
                for fi29 := 0; fi29 < 3; fi29++ {
                    v2[("hello" + (("\n\t" + "hello") + (("hello" + "a") + "foo bar")))] = false
                    v30 := GenS7{F0: fi29}
                    var v31 uint = 22
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
        }
    }
    var v32 uint64 = v16
    switch 32767 {
    case 856:
        v15 = v15
    case 4294967295:
        var v33 float64 = v17
        _ = v33
    case -702:
        if (18446744073709551615 <= 1) {
            v34 := GenS4{F0: 65535, F1: (^(v14)), F2: 9223372036854775807}
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v7
}
