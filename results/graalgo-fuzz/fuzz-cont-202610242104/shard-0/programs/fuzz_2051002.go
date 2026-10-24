package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int64; F2 int16 }
type GenS3 struct { F0 int8; F1 string; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int32; F2 bool }

func main() {
    var v1 bool = (false || (!true))
    v2 := []int{}
    var v3 int16 = -897
    var v4 int16 = (v3 & (v3 + v3))
    var v5 int16 = 1000
    fmt.Println(v4)
    v3 = (v4 | v3)
    v4 -= v3
    if v1 {
        v6 := GenS1{F0: 1}
        var v7 float64 = 3.14159
        v4 += (v3 * (v4 + -32768))
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: -1}
        var v9 float32 = 3.14159
        var v10 int = 2147483647
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        switch "a" {
        case "a":
            if v1 {
                var v12 uint16 = 255
                _ = v12
            }
        case "":
            var v13 bool = (274 != 65535)
            _ = v13
        default:
            v3 = (-(v4))
        }
        switch -594 {
        case 215:
            v4 -= v5
        }
        var v14 uint64 = 614
        _ = fi11
        _ = v14
    }
    v3 += v5
    var v15 bool = false
    v16 := GenS2{F0: 2147483648, F1: -9223372036854775808, F2: -32768}
    fmt.Println(v5)
    var v17 uint = 18446744073709551615
    var v18 int16 = (v4 & v5)
    for fi19 := 0; fi19 < 13; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            v2 = append(v2, -9223372036854775808)
            v21 := GenS3{F0: 42, F1: "", F2: 4294967295}
            if (4294967295 >= 1) {
                switch ("a" + "foo bar") {
                case "\n\t":
                    var v22 uint = 1
                    _ = v22
                case "a":
                    switch fi20 {
                    case -601:
                        if ((65535 == 912) || v15) {
                            v1 = (-2.25 > 1.5)
                            v15 = v15
                            var v23 int16 = (v18 << (((v17 << (((v17 ^ (v17 / (((^(v17))) | 1)))) & 63))) & 63))
                            _ = v23
                        }
                    case 930:
                        _ = len(v2)
                        var v24 int32 = 2147483647
                        if true {
                            v15 = v1
                            v17 = (v17 | 672)
                            var v25 string = ("hello" + "hello")
                            _ = v25
                        }
                        _ = v24
                    }
                default:
                    switch (fi19 >> ((v17) & 63)) {
                    case -1:
                        var v26 int8 = 127
                        _ = v26
                    case -347:
                        var v27 bool = ((-128 >= 127) && (!((v17 ^ (v17 >> ((v17) & 63))) < v17)))
                        _ = v27
                    case -2147483648:
                        var v28 float64 = 1.5
                        _ = v28
                    default:
                        var v29 bool = (!v1)
                        _ = v29
                    }
                }
                if (-83 > -23) {
                    v30 := GenS1{F0: -108}
                    fmt.Println(v17, v1, fi19)
                    v31 := GenS2{F0: 609, F1: -9223372036854775808, F2: -1}
                    _ = v30
                    _ = v31
                }
            } else {
                v2 = append(v2, fi20)
                v18 = (v3 % ((154) | 1))
                if ((true && v15) || v15) {
                    if len(v2) > 0 {
                        v2[0] = fi19
                    }
                    _ = len(v2)
                }
                v2 = append(v2, (fi19 - 785))
                fmt.Println(fi20, v4, fi19)
            }
            _ = fi20
            _ = v21
        }
        if ("" >= ("a" + ("\n\t" + "a"))) {
            var v32 bool = (42 != uint16(v17))
            v33 := GenS1{F0: -48}
            for fi34 := 0; fi34 < 14; fi34++ {
                v1 = true
                v35 := GenS4{F0: 255, F1: 42, F2: v32}
                _ = fi34
                _ = v35
            }
            _ = v32
            _ = v33
        } else {
            var v36 float64 = 1e10
            v37 := GenS4{F0: 65535, F1: 650, F2: (((v18 % ((v4) | 1)) < (-(v5))) || (100 < 42))}
            if ((v17 % ((v17) | 1)) != v17) {
                if len(v2) > 0 {
                    v2[0] = (fi19 >> ((v17) & 63))
                }
                fmt.Println(fi19)
                v17 = 247
            }
            _ = v36
            _ = v37
        }
        _ = fi19
    }
    if v1 {
        v15 = (true && false)
        var v38 uint16 = 255
        v17 -= v17
        _ = v38
    }
    v15 = false
    var v39 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
