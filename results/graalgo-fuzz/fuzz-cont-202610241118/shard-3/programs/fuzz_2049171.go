package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint }
type GenS5 struct { F0 bool; F1 int }
func genFn0(p0 uint16, p1 int16, p2 uint32) int16 {
    fmt.Println(p0)
    p2 = (p2 / (((p2 * 42)) | 1))
    p1 = p1
    var v5 int32 = -266
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 / ((-239) | 1))
}
func genFn1(p0 int16, p1 bool, p2 uint64) uint8 {
    p0 -= p0
    if (!false) {
        var v6 uint = 45
        _ = v6
    }
    var v7 int8 = -58
    v7 += v7
    v8 := []int{32767, -2147483648, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 43
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = 6
    var v3 string = ("" + "a")
    var v4 uint = 18446744073709551615
    fmt.Println(v4, v4, v1)
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(fi9, v2, v1)
        _ = fi9
    }
    var v10 uint = 42
    var v11 float32 = -2.25
    switch (v1 + (v1 ^ (v1 ^ v1))) {
    case -2147483648:
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 int32 = 718
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    case 127:
        var v15 uint64 = 263
        _ = v15
    default:
        if false {
            switch (v1 >> (((v10 % (((v4 >> (((^(v4))) & 63))) | 1))) & 63)) {
            case 9223372036854775807:
                for fi16 := 0; fi16 < 3; fi16++ {
                    if false {
                        v17 := []int{v1, (v1 + (fi16 + 4294967295)), fi16}
                        if (3.14159 != ((((v11 / ((0.0 / (v11 - -2.25)) - float32(-2.25))) - v11) / (3.14159 - 1e10)) / v11)) {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                        } else {
                            var v18 float32 = float32(0.1)
                            _ = v18
                        }
                        v19 := GenS1{F0: genFn0(32767, -1, 626)}
                        _ = v17
                        _ = v19
                    }
                    _ = fi16
                }
            case -2147483648:
                fmt.Println(v1)
            }
            v20 := GenS1{F0: -32768}
            _ = v20
        }
    }
    if false {
        v21 := GenS1{F0: 1000}
        v22 := GenS2{F0: true}
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v10 += (^(v4))
        if false {
            v24 := GenS2{F0: false}
            var v25 int = (v1 + -238)
            fmt.Println(v11, fi23, v3)
            _ = v24
            _ = v25
        }
        _ = fi23
    }
    v26 := []int{int(v4), v1}
    v11 = v11
    v4 -= 4294967295
    if (1 > -9223372036854775808) {
        if true {
            v11 = v11
            if false {
                fmt.Println(v10)
                switch -2147483648 {
                case -650:
                    var v27 uint = (^(v4))
                    _ = v27
                }
            }
            v28 := GenS3{F0: 18446744073709551615, F1: 49}
            _ = v28
        } else {
            if len(v26) > 0 {
                v26[0] = v1
            }
            v29 := GenS4{F0: 0, F1: 2147483648, F2: (v10 | v4)}
            _ = v29
        }
        if (v1 != v1) {
            v30 := GenS4{F0: 922, F1: 153, F2: 862}
            if (((-1 > int32(v2)) || false) || (v2 <= (^(v2)))) {
                var v31 string = v3
                var v32 int32 = 2147483647
                _ = v31
                _ = v32
            } else {
                fmt.Println(v4, v4, v11)
            }
            if (v3 <= ("hello" + ((v3 + (((v3 + (((v3 + ("foo bar" + ("" + "a"))) + "hello") + v3)) + (v3 + v3)) + v3)) + v3))) {
                var v33 uint8 = 164
                if len(v26) > 0 {
                    v26[0] = 2147483647
                }
                _ = v33
            } else {
                v4 -= (v10 ^ v4)
            }
            _ = v30
        } else {
            for fi34 := 0; fi34 < 8; fi34++ {
                v1 = 646
                switch 948 {
                case 4294967295:
                    var v35 string = (v3 + (v3 + v3))
                    _ = v35
                }
                if ((((v3 + "\n\t") + "hello") + "") == ((("" + v3) + "") + ("foo bar" + ""))) {
                    switch "a" {
                    case "":
                        v3 = ("foo bar" + (v3 + ("foo bar" + "foo bar")))
                    }
                    var v36 int = v1
                    _ = v36
                } else {
                    v26 = append(v26, (v1 - fi34))
                    var v37 float32 = (float32(((0.0 * 3.14159) * 0.1)) + v11)
                    if len(v26) > 0 {
                        v26[0] = (v1 % ((42) | 1))
                    }
                    _ = v37
                }
                _ = fi34
            }
        }
    }
    v38 := GenS2{F0: ((!((((v3 + v3) + v3) + "\n\t") < "hello")) && (v10 < (v10 >> ((1) & 63))))}
    for fi39 := 0; fi39 < 9; fi39++ {
        var v40 uint8 = genFn1(42, ((v3 + ((v3 + v3) + v3)) > (v3 + (v3 + "a"))), 4294967295)
        switch (^(fi39)) {
        case -83:
            switch fi39 {
            case -105:
                v41 := GenS2{F0: (265 > v4)}
                _ = v41
            case 127:
                for fi42 := 0; fi42 < 7; fi42++ {
                    v2 = int8(fi39)
                    v4 += v4
                    fmt.Println(v3, v2, v40)
                    _ = fi42
                }
            }
        case 127:
            for fi43 := 0; fi43 < 12; fi43++ {
                if len(v26) > 0 {
                    v26[0] = -9223372036854775808
                }
                _ = fi43
            }
        default:
            v44 := GenS3{F0: 954, F1: genFn1(1000, (0 > -32768), 1)}
            _ = v44
        }
        fmt.Println(fi39, v4)
        _ = fi39
        _ = v40
    }
    var v45 uint = v10
    var v46 uint64 = 1
    for fi47 := 0; fi47 < 6; fi47++ {
        var v48 string = v3
        if true {
            if len(v26) > 0 {
                v26[0] = fi47
            }
            var v49 int = 42
            _ = v49
        } else {
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            v50 := GenS5{F0: ((-(fi47)) <= v1), F1: fi47}
            fmt.Println(v48)
            _ = v50
        }
        v11 = (v11 / 1.0)
        _ = fi47
        _ = v48
    }
    var v51 int = (v1 * v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v38
    _ = v4
    _ = v45
    _ = v46
    _ = v51
}
