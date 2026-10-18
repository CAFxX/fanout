package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 int64; F1 int64; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 bool; F2 uint16 }
func genFn0(p0 int) uint8 {
    var v7 uint8 = 6
    switch ((((("\n\t" + ("\n\t" + "hello")) + ("\n\t" + "hello")) + (("" + "") + "")) + ("" + ("" + "a"))) + "a") {
    case "a":
        var v8 uint16 = 205
        _ = v8
    case "hello":
        var v9 uint8 = 127
        _ = v9
    }
    v7 += v7
    _ = p0
    _ = v7
    return 32
}
func genFn1(p0 uint16, p1 bool, p2 bool) uint {
    var v10 float64 = 1.5
    var v11 bool = (-128 < -128)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 65535
    var v3 int16 = -1
    var v4 bool = (((-2.25 * (1.5 / -2.25)) * 0.0) == 1.5)
    var v5 string = ("hello" + "hello")
    var v6 bool = (v2 != v2)
    v3 = v3
    var v12 int16 = v3
    fmt.Println(v3, v12)
    v13 := []int{-556, 128, 260}
    for fi14 := 0; fi14 < 8; fi14++ {
        v5 = "\n\t"
        _ = fi14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS1{F0: 5, F1: 18446744073709551615, F2: v3}
        _ = fi15
        _ = v16
    }
    if true {
        delete(v1, (v5 + ""))
    }
    delete(v1, v5)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    fmt.Println(v3)
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v12, v2, v4)
        v5 = v5
        switch (fi17 % ((fi17) | 1)) {
        case -711:
            var v18 float32 = -2.25
            _ = v18
        case -1:
            if len(v13) > 0 {
                v13[0] = 498
            }
        }
        _ = fi17
    }
    v6 = (v4 && false)
    var v19 bool = (18446744073709551615 <= 1)
    v12 -= v12
    for fi20 := 0; fi20 < 9; fi20++ {
        _ = len(v13)
        if (v4 && v6) {
            switch 566 {
            case 42:
                switch "" {
                case "a":
                    v21 := GenS2{F0: 1, F1: 0, F2: -538}
                    _ = v21
                case "foo bar":
                    if v19 {
                        var v22 float64 = (((3.14159 - 0.0) - 1.5) / (0.0 + -2.25))
                        var v23 int = 4294967295
                        for fi24 := 0; fi24 < 5; fi24++ {
                            v23 = (fi24 << ((4294967295) & 63))
                            v25 := GenS3{F0: (v3 % ((-32768) | 1))}
                            _ = fi24
                            _ = v25
                        }
                        _ = v22
                        _ = v23
                    } else {
                        switch ("foo bar" + v5) {
                        case "foo bar":
                            for fi26 := 0; fi26 < 9; fi26++ {
                                v3 = 42
                                var v27 bool = false
                                _ = fi26
                                _ = v27
                            }
                        case "\n\t":
                            if ((fi20 * (fi20 >> ((2) & 63))) < fi20) {
                                v4 = true
                                fmt.Println(v6, v3, fi20)
                            }
                        case "a":
                            for fi28 := 0; fi28 < 7; fi28++ {
                                var v29 float32 = 3.14159
                                _ = fi28
                                _ = v29
                            }
                        }
                    }
                case "":
                    v30 := GenS2{F0: -1, F1: 0, F2: (fi20 * 32767)}
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 7; fi31++ {
                        v6 = false
                        if true {
                            v32 := GenS4{F0: v5, F1: (803 != 2), F2: (v2 / ((42) | 1))}
                            _ = v32
                        } else {
                            v33 := GenS2{F0: 0, F1: 42, F2: fi31}
                            var v34 uint16 = (v2 + v2)
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            _ = v33
                            _ = v34
                        }
                        _ = fi31
                    }
                }
            }
        } else {
            v3 += v3
            var v35 float32 = 1e10
            v13 = append(v13, 0)
            _ = v35
        }
        _ = fi20
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        for fi37 := 0; fi37 < 9; fi37++ {
            _ = len(v13)
            v38 := GenS1{F0: -128, F1: genFn1((v2 ^ 2), (!true), ((18446744073709551615 < uint64(v2)) && v19)), F2: (v12 >> ((735) & 63))}
            _ = fi37
            _ = v38
        }
        var v39 int8 = 100
        v12 -= v12
        _ = fi36
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
