package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint32; F1 int; F2 uint16 }
type GenS6 struct { F0 uint64; F1 int16; F2 int16 }
type GenS7 struct { F0 string }
func genFn0(p0 uint8) int16 {
    p0 += p0
    v7 := GenS1{F0: 2}
    p0 += (p0 | 92)
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: 65535}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return 1000
}

func main() {
    var v1 string = "\n\t"
    var v2 int64 = 1
    var v3 int16 = 857
    v4 := []uint32{2, 2147483648, 361}
    var v5 float64 = 0.1
    var v6 float64 = v5
    if true {
        fmt.Println(v6)
    } else {
        switch 0 {
        case 4294967295:
            var v10 int = -921
            _ = v10
        default:
            v11 := GenS1{F0: 42}
            _ = v11
        }
        v12 := GenS1{F0: 540}
        if true {
            _ = len(v4)
            v1 = ""
        }
        _ = v12
    }
    switch (v1 + (("foo bar" + ("" + v1)) + v1)) {
    case "":
        for fi13 := 0; fi13 < 17; fi13++ {
            v3 = (v3 | v3)
            _ = fi13
        }
    case "foo bar":
        v1 = (v1 + "hello")
    default:
        var v14 uint16 = 32767
        _ = v14
    }
    switch 127 {
    case -260:
        v15 := GenS2{F0: 127}
        _ = v15
    default:
        for fi16 := 0; fi16 < 16; fi16++ {
            v3 = 554
            v6 += (v5 * float64((3.14159 / float32(((-2.25 - 1.5) / 1e10)))))
            _ = fi16
        }
    }
    v17 := GenS1{F0: 42}
    v3 -= v3
    if (!true) {
        v1 = (v1 + v1)
        v18 := GenS2{F0: 127}
        if len(v4) > 0 {
            v4[0] = 350
        }
        _ = v18
    }
    var v19 int16 = (v3 >> ((119) & 63))
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 uint32 = 65535
            v23 := GenS3{F0: 0, F1: (v22 << ((4294967295) & 63)), F2: 127}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        v1 = "foo bar"
        _ = fi24
    }
    var v25 string = "a"
    for fi26 := 0; fi26 < 3; fi26++ {
        switch 128 {
        case 42:
            if (((0.1 / v6) != 3.14159) && true) {
                for fi27 := 0; fi27 < 11; fi27++ {
                    if (889 < 743) {
                        if (!true) {
                            var v28 int = (fi26 / (((fi27 % (((fi27 * (^(fi27)))) | 1))) | 1))
                            v25 = v1
                            _ = v28
                        } else {
                            if true {
                                fmt.Println(fi27)
                                v6 = 0.1
                                if len(v4) > 0 {
                                    v4[0] = 2147483647
                                }
                            }
                            v4 = append(v4, 902)
                            switch "hello" {
                            case "":
                                fmt.Println(v1, v3, v5)
                            }
                        }
                        v29 := GenS2{F0: 105}
                        v30 := GenS3{F0: 161, F1: 65535, F2: 245}
                        _ = v29
                        _ = v30
                    }
                    v31 := GenS4{F0: v19, F1: ("foo bar" + (v25 + v1))}
                    _ = fi27
                    _ = v31
                }
            }
        default:
            if true {
                for fi32 := 0; fi32 < 20; fi32++ {
                    v33 := GenS5{F0: 0, F1: (fi26 | fi32), F2: 65535}
                    if (42 != 2) {
                        var v34 int16 = v19
                        v3 = genFn0(0)
                        v35 := GenS6{F0: 112, F1: v19, F2: v34}
                        v36 := GenS7{F0: ("a" + v1)}
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    var v37 uint8 = 155
                    var v38 bool = false
                    _ = fi32
                    _ = v33
                    _ = v37
                    _ = v38
                }
            } else {
                v6 = (v5 + (v5 * (1.5 / float64(-2.25))))
            }
        }
        for fi39 := 0; fi39 < 16; fi39++ {
            v40 := GenS1{F0: 2}
            fmt.Println(v25)
            _ = fi39
            _ = v40
        }
        _ = fi26
    }
    if ((3.14159 * 1e10) > (1.5 - -2.25)) {
        var v41 float64 = (3.14159 - 0.0)
        _ = v41
    }
    v4 = append(v4, 2147483648)
    v42 := GenS6{F0: 18446744073709551615, F1: (v3 * genFn0(1)), F2: (v3 & (v19 << ((65535) & 63)))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
