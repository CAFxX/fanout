package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int16; F1 int; F2 int32 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 int16; F1 int64; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 bool; F2 int }
func genFn0() int16 {
    var v6 string = "\n\t"
    v7 := []int{128, -916, 311, -414}
    var v8 uint8 = 2
    if false {
        var v9 uint16 = 42
        _ = v9
    }
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 bool) int64 {
    v10 := []int{2147483647, 2147483647, 875, 42}
    for fi11 := 0; fi11 < 15; fi11++ {
        if p0 {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: 744, F1: 380, F2: p0}
                switch "a" {
                case "":
                    p0 = (350 != 0)
                }
                p0 = false
                _ = fi12
                _ = v13
            }
        }
        var v14 string = ("a" + ("foo bar" + "\n\t"))
        _ = fi11
        _ = v14
    }
    var v15 float32 = 1e10
    _ = p0
    _ = v10
    _ = v15
    return 9223372036854775807
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]string)
    var v3 float32 = 1e10
    v4 := []string{"", ("a" + "a"), "foo bar", "hello"}
    var v5 uint64 = 230
    var v16 uint8 = 0
    switch "foo bar" {
    case "\n\t":
        v5 = v5
    case "foo bar":
        v17 := GenS1{F0: 203, F1: 2, F2: ((v16 << ((0) & 63)) < 128)}
        _ = v17
    case "":
        if true {
            if true {
                v1 += v1
            }
            v16 -= (v16 & v16)
        }
    }
    v18 := GenS2{F0: v5, F1: "hello"}
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS3{F0: 0, F1: (fi20 * (fi19 ^ (fi20 % ((fi19) | 1)))), F2: v1}
            if (true || true) {
                v22 := GenS1{F0: 42, F1: 0, F2: ((1 == 4294967295) || true)}
                var v23 uint16 = 42
                switch "a" {
                case "hello":
                    var v24 uint = 0
                    var v25 int16 = int16(v16)
                    v26 := GenS4{F0: genFn1((((v24 > 2) && true) || false)), F1: v5, F2: ((^(v16)) != 127)}
                    _ = v24
                    _ = v25
                    _ = v26
                case "":
                    v27 := GenS2{F0: v5, F1: ("a" + "hello")}
                    _ = v27
                }
                var v28 int8 = 0
                fmt.Println(v1)
                _ = v22
                _ = v23
                _ = v28
            }
            var v29 uint64 = 4294967295
            if (((-2.25 / (((((1e10 - float64(3.14159)) * 0.1) - 3.14159) - 1.5) - -2.25)) == 0.0) && false) {
                fmt.Println(v29, v29, fi20)
                var v30 uint32 = 2147483647
                _ = v30
            }
            _ = fi20
            _ = v21
            _ = v29
        }
        _ = fi19
    }
    v3 += v3
    switch 32768 {
    case -568:
        if true {
            v3 -= 1e10
            v31 := GenS5{F0: 0, F1: 763, F2: 0}
            _ = v31
        }
    default:
        v3 += 3.14159
    }
    if (!false) {
        var v32 uint16 = 845
        _ = v32
    } else {
        v16 = 127
    }
    switch "foo bar" {
    case "\n\t":
        v5 += 2
    default:
        fmt.Println(v16, v16)
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        if ((!false) && ((0 <= 0) && ((!(!("a" == ("a" + "foo bar")))) && true))) {
            v34 := GenS2{F0: (^(v5)), F1: ("hello" + ("" + ((("" + "a") + "foo bar") + "hello")))}
            _ = v34
        } else {
            switch fi33 {
            case 64:
                v16 = v16
            case -1:
                var v35 bool = true
                _ = v35
            default:
                v3 = (float32((0.0 / 1e10)) / 0.1)
            }
        }
        for fi36 := 0; fi36 < 12; fi36++ {
            delete(v2, 2147483647)
            v37 := GenS6{F0: 42}
            _ = fi36
            _ = v37
        }
        v38 := GenS1{F0: 166, F1: 679, F2: true}
        _ = fi33
        _ = v38
    }
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    if (!false) {
        v39 := GenS7{F0: "", F1: (32767 < 752), F2: 1}
        _ = v39
    }
    v40 := GenS7{F0: ("foo bar" + ("a" + "a")), F1: false, F2: 1}
    var v41 int64 = 894
    var v42 bool = true
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v43 uint16 = 397
    if v42 {
        var v44 string = "foo bar"
        _ = v44
    }
    v41 += v41
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
