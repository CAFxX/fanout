package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint; F2 uint64 }
func genFn0() uint16 {
    var v7 float32 = -2.25
    var v8 float64 = 3.14159
    var v9 bool = false
    v10 := GenS1{F0: 8, F1: 9223372036854775807}
    var v11 uint64 = 18446744073709551615
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 911
}
func genFn1(p0 int8, p1 string, p2 string) uint8 {
    fmt.Println(p1, p2)
    v12 := []int{32768}
    if false {
        var v13 float32 = -2.25
        var v14 uint64 = 2
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 247
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]bool)
    var v3 int16 = 1000
    v4 := []uint32{2147483648, 2147483647}
    v5 := make(map[string]uint32)
    var v6 int16 = v3
    v15 := GenS1{F0: 0, F1: -9223372036854775808}
    v3 = (v6 % (((-(v3))) | 1))
    _ = len(v5)
    for fi16 := 0; fi16 < 18; fi16++ {
        _ = len(v5)
        v1 = (-(v1))
        _ = fi16
    }
    v17 := GenS1{F0: 1, F1: -9223372036854775808}
    var v18 uint8 = 255
    fmt.Println(v6, v1, v3)
    for fi19 := 0; fi19 < 14; fi19++ {
        v6 -= -429
        _ = fi19
    }
    switch "" {
    case "a":
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS2{F0: 4294967295, F1: (768 < 985), F2: 60}
            switch "foo bar" {
            case "foo bar":
                if (!(true && (0 >= v1))) {
                    v1 += (v1 | v1)
                    if (2147483647 != 1) {
                        v4 = append(v4, 2)
                        var v22 int64 = -1
                        v23 := GenS2{F0: 45, F1: (!false), F2: -41}
                        _ = v22
                        _ = v23
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v24 uint8 = 129
                        if false {
                            v2[int(v1)] = false
                        }
                        _ = v24
                    }
                }
            case "":
                v25 := GenS2{F0: 327, F1: (false || (fi20 >= fi20)), F2: 1}
                _ = v25
            case "a":
                fmt.Println(v3, fi20)
            }
            _ = fi20
            _ = v21
        }
    case "foo bar":
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            v18 -= (v18 ^ genFn1(-1, ("a" + ("a" + ("a" + ""))), (("\n\t" + "") + ("foo bar" + "\n\t"))))
            fmt.Println(v18, v3, v1)
            for fi26 := 0; fi26 < 16; fi26++ {
                if ((v18 + genFn1(99, "", "hello")) <= (v18 / (((v18 >> ((15) & 63))) | 1))) {
                    switch "a" {
                    case "":
                        switch "\n\t" {
                        case "\n\t":
                            var v27 bool = false
                            _ = v27
                        case "":
                            delete(v5, (("a" + (("hello" + (("\n\t" + "a") + ((("foo bar" + ("a" + "hello")) + ("\n\t" + "hello")) + ("" + "")))) + "a")) + ""))
                        case "hello":
                            for fi28 := 0; fi28 < 6; fi28++ {
                                v6 = v6
                                _ = fi28
                            }
                        }
                    case "\n\t":
                        var v29 uint64 = 810
                        _ = v29
                    }
                    v30 := GenS3{F0: (fi26 * (fi26 & (fi26 << ((255) & 63))))}
                    _ = v30
                } else {
                    var v31 int32 = v1
                    if true {
                        var v32 uint32 = 65535
                        v1 = 2147483647
                        delete(v2, (fi26 + int(v6)))
                        _ = v32
                    }
                    _ = v31
                }
                v6 = (v3 - (v3 << ((18446744073709551615) & 63)))
                _ = fi26
            }
        }
    case "\n\t":
        var v33 uint16 = 558
        _ = v33
    default:
        fmt.Println(v3, v6, v18)
    }
    var v34 string = "foo bar"
    var v35 int64 = 9223372036854775807
    for fi36 := 0; fi36 < 18; fi36++ {
        var v37 float32 = 1e10
        v38 := GenS1{F0: 100, F1: (v35 - (^(v35)))}
        var v39 float32 = (v37 * 0.0)
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    if (!false) {
        v40 := GenS4{F0: 1}
        v6 -= v6
        _ = v40
    } else {
        v41 := GenS5{F0: 58, F1: 2, F2: 684}
        _ = len(v5)
        delete(v2, 127)
        _ = v41
    }
    var v42 uint32 = 4294967295
    for fi43 := 0; fi43 < 11; fi43++ {
        delete(v2, fi43)
        var v44 uint = 18446744073709551615
        _ = fi43
        _ = v44
    }
    fmt.Println(v6, v34, v18)
    _ = len(v5)
    var v45 uint = 169
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v45
    _ = v5
    _ = v6
}
