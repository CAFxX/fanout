package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint64; F2 int32 }
type GenS3 struct { F0 bool; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint }
func genFn0() int8 {
    if true {
        v5 := GenS1{F0: 0}
        if true {
            var v6 int64 = -1
            if (("a" <= "") && false) {
                switch "a" {
                case "foo bar":
                    for fi7 := 0; fi7 < 17; fi7++ {
                        v8 := GenS2{F0: v6, F1: 0, F2: -2147483648}
                        if true {
                            var v9 float32 = ((((1e10 * 0.1) / float32(3.14159)) / -2.25) - 1.5)
                            var v10 uint16 = 0
                            _ = v10
                            _ = v9
                        }
                        _ = fi7
                        _ = v8
                    }
                case "a":
                    var v11 int = 1
                    _ = v11
                default:
                    fmt.Println(v6)
                }
            } else {
                v6 = (-(v6))
                v12 := GenS2{F0: (v6 % (((v6 % ((v6) | 1))) | 1)), F1: 178, F2: 1}
                _ = v12
            }
            var v13 uint16 = 255
            _ = v13
            _ = v6
        }
        _ = v5
    }
    v14 := GenS1{F0: 32767}
    _ = v14
    return -128
}
func genFn1() uint {
    switch "" {
    case "hello":
        var v15 uint16 = 1
        _ = v15
    default:
        var v16 uint32 = 0
        _ = v16
    }
    if (!false) {
        var v17 string = "\n\t"
        switch "" {
        case "":
            v17 = v17
        case "a":
            var v18 int64 = 0
            _ = v18
        case "foo bar":
            v19 := GenS3{F0: (!true), F1: 42, F2: 225}
            _ = v19
        }
        v17 = ("" + "hello")
        _ = v17
    }
    v20 := []int{9223372036854775807, 9223372036854775807}
    v21 := GenS4{F0: 42}
    v22 := GenS4{F0: 4294967295}
    _ = v20
    _ = v21
    _ = v22
    return 25
}

func main() {
    var v1 bool = true
    v2 := []int{-454}
    var v3 uint32 = 890
    var v4 uint = 2
    for fi23 := 0; fi23 < 16; fi23++ {
        if true {
            v3 = uint32(v4)
            v3 -= 2147483648
        } else {
            v1 = v1
            var v24 int32 = 2147483647
            _ = v24
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        if v1 {
            var v26 uint = (v4 | v4)
            var v27 int16 = 42
            _ = v26
            _ = v27
        } else {
            if v1 {
                v3 = (^(v3))
            } else {
                v28 := GenS5{F0: 128, F1: 453}
                var v29 float64 = 0.0
                _ = v28
                _ = v29
            }
            switch "" {
            case "\n\t":
                var v30 float64 = 1.5
                _ = v30
            case "foo bar":
                for fi31 := 0; fi31 < 19; fi31++ {
                    v3 += (v3 << ((v4) & 63))
                    var v32 bool = (v1 && (!v1))
                    if (1e10 == 3.14159) {
                        v1 = v32
                        var v33 int = (fi25 ^ fi25)
                        fmt.Println(v1)
                        _ = v33
                    }
                    v34 := GenS1{F0: -955}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi31
                    _ = v32
                    _ = v34
                }
            }
        }
        fmt.Println(v3)
        _ = fi25
    }
    v1 = v1
    v35 := GenS1{F0: 827}
    switch "" {
    case "hello":
        v4 -= 18446744073709551615
    default:
        var v36 string = "hello"
        _ = v36
    }
    v37 := GenS4{F0: (v4 | (v4 | v4))}
    v3 = v3
    if (!(872 <= -881)) {
        fmt.Println(v3, v4, v4)
        var v38 int16 = 1000
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v38
    } else {
        var v39 bool = false
        _ = v39
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
}
