package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS5 struct { F0 int; F1 uint16; F2 int16 }
type GenS6 struct { F0 bool }
func genFn0(p0 int8, p1 int64, p2 int) int8 {
    fmt.Println(p2)
    var v4 int8 = (p0 - p0)
    v5 := []int{p2, p2, (p2 / ((128) | 1)), (-(p2))}
    if len(v5) > 0 {
        v5[0] = 169
    }
    if (2 <= 170) {
        v6 := GenS1{F0: v4, F1: "hello"}
        _ = v6
    } else {
        p2 -= (p2 ^ (p2 * (p2 | p2)))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if false {
            v7 := GenS1{F0: p0, F1: "\n\t"}
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 int16) uint64 {
    var v8 int16 = p0
    v9 := []int{0, 42, 0, 32767}
    p0 = (v8 | p0)
    _ = p0
    _ = v8
    _ = v9
    return 803
}

func main() {
    v1 := []int{-2147483648, 9223372036854775807}
    var v2 uint64 = 18446744073709551615
    var v3 bool = ((1 >= 882) || (true && ((false || (127 != 127)) && (v2 > (v2 + v2)))))
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int64 = 195
        _ = fi10
        _ = v11
    }
    if len(v1) > 0 {
        v1[0] = 32768
    }
    v12 := GenS1{F0: genFn0(1, 1, 9223372036854775807), F1: "hello"}
    for fi13 := 0; fi13 < 2; fi13++ {
        v2 -= genFn1(919)
        var v14 int = fi13
        fmt.Println(fi13, v3, v3)
        _ = fi13
        _ = v14
    }
    v3 = true
    switch 127 {
    case 9223372036854775807:
        for fi15 := 0; fi15 < 6; fi15++ {
            v2 = (v2 << ((845) & 63))
            _ = fi15
        }
    case 823:
        switch 254 {
        case 0:
            switch 637 {
            case -905:
                v16 := GenS2{F0: 0}
                _ = v16
            case -1:
                switch 1 {
                case 0:
                    v3 = (503 != 174)
                case 2147483647:
                    var v17 int = 128
                    _ = v17
                case 212:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if (!(v3 && v3)) {
                            v19 := GenS2{F0: 32767}
                            fmt.Println(fi18)
                            _ = v19
                        } else {
                            switch (fi18 * fi18) {
                            case 42:
                                for fi20 := 0; fi20 < 2; fi20++ {
                                    var v21 int = 4294967295
                                    v22 := GenS1{F0: -128, F1: "hello"}
                                    v23 := GenS3{F0: (1 > 0), F1: 42}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v24 int8 = -14
                                    var v25 uint32 = 42
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            case 32768:
                                if false {
                                    var v26 int = (fi18 % ((fi18) | 1))
                                    _ = v26
                                }
                            case -1:
                                v1 = append(v1, fi18)
                            default:
                                var v27 int64 = 1
                                _ = v27
                            }
                            var v28 int64 = 42
                            var v29 uint16 = 626
                            _ = v28
                            _ = v29
                        }
                        _ = fi18
                    }
                }
            case -2147483648:
                for fi30 := 0; fi30 < 12; fi30++ {
                    fmt.Println(v2, v3, v2)
                    var v31 uint64 = (v2 % ((v2) | 1))
                    _ = fi30
                    _ = v31
                }
            }
        case 9223372036854775807:
            var v32 uint16 = 0
            _ = v32
        case 462:
            v33 := GenS1{F0: int8(v2), F1: (("hello" + (("" + ("\n\t" + ("a" + (("\n\t" + (((("a" + "hello") + "\n\t") + "foo bar") + "\n\t")) + "foo bar")))) + ("a" + "\n\t"))) + "hello")}
            _ = v33
        }
    case 42:
        v34 := GenS3{F0: false, F1: 32767}
        _ = v34
    default:
        var v35 float32 = 0.0
        _ = v35
    }
    var v36 float64 = 3.14159
    if (v3 && true) {
        var v37 int16 = 153
        _ = v37
    }
    v38 := GenS4{F0: 0, F1: 0, F2: 937}
    for fi39 := 0; fi39 < 17; fi39++ {
        if len(v1) > 0 {
            v1[0] = fi39
        }
        switch "foo bar" {
        case "a":
            v40 := GenS5{F0: 907, F1: 604, F2: 1000}
            _ = v40
        default:
            fmt.Println(v36, fi39)
        }
        _ = fi39
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, -2147483648)
    v3 = (42 > 333)
    for fi41 := 0; fi41 < 5; fi41++ {
        for fi42 := 0; fi42 < 14; fi42++ {
            if (!(-818 >= 148)) {
                if true {
                    if v3 {
                        v3 = (42 >= 119)
                        v3 = false
                        var v43 uint16 = 42
                        if v3 {
                            fmt.Println(v3, v43, v43)
                        }
                        _ = v43
                    }
                    v44 := GenS6{F0: ((v3 || (32767 < -893)) || v3)}
                    _ = v44
                } else {
                    v45 := GenS2{F0: 906}
                    var v46 bool = v3
                    _ = v45
                    _ = v46
                }
                var v47 uint = 635
                var v48 float64 = float64((1e10 * 0.1))
                _ = v47
                _ = v48
            }
            var v49 uint = 4294967295
            _ = len(v1)
            _ = fi42
            _ = v49
        }
        _ = fi41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v38
}
