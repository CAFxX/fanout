package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 bool) uint {
    var v6 float32 = 0.1
    p0 = (((1e10 * (float64(0.0) + (3.14159 - 1.5))) / 1.0) <= float64((float32((1.5 - ((3.14159 + -2.25) - (-2.25 + (0.1 / 0.1))))) + v6)))
    _ = p0
    _ = v6
    return 4294967295
}
func genFn1(p0 string) string {
    switch p0 {
    case "a":
        p0 = (((p0 + (p0 + p0)) + "\n\t") + p0)
    case "hello":
        var v7 string = "hello"
        _ = v7
    }
    switch "foo bar" {
    case "foo bar":
        var v8 float64 = (0.0 - -2.25)
        _ = v8
    default:
        for fi9 := 0; fi9 < 8; fi9++ {
            p0 = p0
            var v10 int32 = -1
            _ = fi9
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        if false {
            var v12 uint8 = 231
            var v13 float64 = 3.14159
            v13 = v13
            _ = v12
            _ = v13
        } else {
            var v14 float64 = 1e10
            var v15 int64 = 515
            _ = v14
            _ = v15
        }
        var v16 uint32 = 65535
        for fi17 := 0; fi17 < 11; fi17++ {
            fmt.Println(fi17)
            v18 := GenS1{F0: p0}
            var v19 uint16 = 22
            if (596 < 18446744073709551615) {
                switch "hello" {
                case "foo bar":
                    v16 -= (v16 + v16)
                case "":
                    var v20 int8 = int8(fi11)
                    _ = v20
                default:
                    v19 = v19
                }
                v16 -= 2147483648
                v21 := GenS1{F0: "foo bar"}
                _ = v21
            }
            v19 = 2
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi11
        _ = v16
    }
    _ = p0
    return (p0 + p0)
}

func main() {
    v1 := []string{"a", "hello"}
    v2 := []int{32768, 9223372036854775807}
    var v3 int64 = -909
    v4 := []uint32{}
    var v5 uint32 = 65535
    v22 := GenS1{F0: ("foo bar" + ("a" + "\n\t"))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v23 uint = 42
    var v24 int32 = 345
    for fi25 := 0; fi25 < 18; fi25++ {
        if (!((42 < 0) && true)) {
            var v26 float32 = -2.25
            _ = v26
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := GenS1{F0: ""}
            if ((0 == v23) || (!false)) {
                v29 := GenS1{F0: "foo bar"}
                fmt.Println(fi25)
                fmt.Println(v5, v5, v3)
                var v30 uint64 = 794
                _ = v29
                _ = v30
            }
            if (v24 <= (^(v24))) {
                v3 = (v3 << (((^(v23))) & 63))
                _ = len(v4)
                if true {
                    v24 = (v24 ^ (v24 + -865))
                } else {
                    var v31 float64 = -2.25
                    fmt.Println(v5, v3, v3)
                    _ = v31
                }
            } else {
                switch 127 {
                case 0:
                    switch "" {
                    case "\n\t":
                        if len(v4) > 0 {
                            v4[0] = v5
                        }
                    case "hello":
                        if ((uint64(v5) == 1) || false) {
                            v3 += v3
                            fmt.Println(fi25, v23, fi27)
                        }
                    default:
                        v2 = append(v2, fi25)
                    }
                }
                v5 -= 2147483647
                fmt.Println(v24, fi25, fi27)
                v24 = 2147483647
            }
            _ = fi27
            _ = v28
        }
        fmt.Println(v23)
        _ = fi25
    }
    fmt.Println(v5, v5, v3)
    if len(v4) > 0 {
        v4[0] = (v5 / ((v5) | 1))
    }
    var v32 uint64 = 87
    if false {
        if false {
            var v33 uint32 = 4294967295
            _ = v33
        }
        v24 -= (v24 / (((v24 & (v24 >> ((v23) & 63)))) | 1))
        v32 = v32
    }
    v34 := GenS2{F0: false, F1: 0}
    _ = len(v2)
    if true {
        switch "\n\t" {
        case "a":
            if len(v1) > 0 {
                v1[0] = ("" + "hello")
            }
        }
    }
    fmt.Println(v3)
    v35 := GenS1{F0: ""}
    v3 += v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        _ = len(v2)
        if (673 <= 32767) {
            for fi37 := 0; fi37 < 11; fi37++ {
                v38 := GenS1{F0: ""}
                v23 = (v23 ^ genFn0(((!((v3 ^ -1) > v3)) && false)))
                v23 = (^(v23))
                _ = fi37
                _ = v38
            }
            var v39 int64 = v3
            v39 += v39
            _ = v39
        }
        if false {
            v24 += v24
            fmt.Println(v24, v23)
        }
        _ = fi36
    }
    v2 = append(v2, 127)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
