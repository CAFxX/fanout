package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint; F2 uint64 }
type GenS5 struct { F0 string; F1 int8; F2 int32 }
func genFn0() int16 {
    var v6 uint64 = 450
    if false {
        v7 := []int{701, 32767, 42, 544}
        var v8 int8 = 79
        _ = v7
        _ = v8
    }
    _ = v6
    return 551
}
func genFn1(p0 int, p1 int16) string {
    p1 += (p1 & -207)
    p0 = 4294967295
    v9 := GenS1{F0: 2, F1: (p0 + (p0 ^ 521)), F2: 42}
    switch (p0 << ((631) & 63)) {
    case -1:
        fmt.Println(p0, p1)
    case 293:
        switch "a" {
        case "":
            if (!false) {
                if false {
                    var v10 uint64 = 910
                    _ = v10
                } else {
                    v11 := GenS1{F0: 0, F1: (p0 | 1), F2: 0}
                    p1 = p1
                    v12 := []int{}
                    _ = v11
                    _ = v12
                }
                v13 := GenS2{F0: (18446744073709551615 >= 4294967295), F1: 2, F2: 255}
                _ = v13
            } else {
                v14 := []int{9223372036854775807, p0}
                _ = v14
            }
        case "hello":
            p1 = 245
        case "a":
            if (1 <= 42) {
                var v15 int16 = 0
                _ = v15
            } else {
                v16 := GenS1{F0: 328, F1: (p0 / ((p0) | 1)), F2: 65535}
                var v17 uint16 = 42
                _ = v16
                _ = v17
            }
        }
    case 360:
        var v18 bool = false
        _ = v18
    default:
        var v19 string = (((("\n\t" + ("" + "\n\t")) + "\n\t") + "") + (("" + "foo bar") + "\n\t"))
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v9
    return ""
}

func main() {
    var v1 float64 = (0.0 - ((-2.25 - 1e10) * (3.14159 * float64(float32(((0.0 - float64((0.1 - (0.0 / -2.25)))) + 1e10))))))
    var v2 uint16 = 634
    var v3 uint16 = (v2 + (v2 ^ v2))
    v4 := []uint32{2147483648, 0, 2147483647}
    var v5 int = -2147483648
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 int32 = -1
        switch ((("\n\t" + ("a" + "a")) + ("\n\t" + ("" + "foo bar"))) + "hello") {
        case "hello":
            fmt.Println(fi20, fi20)
        case "\n\t":
            if false {
                v21 = v21
                v5 = (v5 >> ((4294967295) & 63))
            } else {
                for fi22 := 0; fi22 < 15; fi22++ {
                    var v23 uint64 = 42
                    var v24 uint64 = 42
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
        case "a":
            fmt.Println(v3)
        }
        if false {
            v25 := GenS3{F0: ("\n\t" + "")}
            _ = v25
        } else {
            v2 += v3
            if false {
                for fi26 := 0; fi26 < 4; fi26++ {
                    var v27 uint8 = 128
                    var v28 float32 = 1.5
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 5; fi29++ {
                    var v30 int64 = 1
                    var v31 uint16 = (v3 / ((v3) | 1))
                    v5 = (v5 << ((593) & 63))
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                v2 -= (v3 & 754)
            } else {
                var v32 int16 = 350
                fmt.Println(v3, v1)
                _ = v32
            }
        }
        _ = fi20
        _ = v21
    }
    v33 := GenS1{F0: 4294967295, F1: -855, F2: 16}
    var v34 uint32 = 42
    fmt.Println(v34)
    v2 = (v3 >> ((1) & 63))
    if (!true) {
        v1 += v1
    } else {
        for fi35 := 0; fi35 < 12; fi35++ {
            for fi36 := 0; fi36 < 19; fi36++ {
                v34 += (v34 + (^(v34)))
                switch "a" {
                case "":
                    v3 = v2
                }
                if false {
                    if ((1.5 >= 1.5) || false) {
                        v37 := GenS1{F0: 1, F1: (fi35 * (v5 % ((42) | 1))), F2: (^(v34))}
                        fmt.Println(v3)
                        if (!false) {
                            fmt.Println(v2, v34)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if len(v4) > 0 {
                                v4[0] = v34
                            }
                        }
                        v38 := GenS2{F0: (genFn0() >= int16(v34)), F1: (^(v3)), F2: 82}
                        v3 -= 65535
                        _ = v37
                        _ = v38
                    }
                    v3 -= 900
                    _ = len(v4)
                    switch "foo bar" {
                    case "hello":
                        var v39 int64 = 0
                        _ = v39
                    case "":
                        v40 := GenS2{F0: false, F1: v2, F2: 2}
                        _ = v40
                    }
                }
                v2 += (^(v3))
                fmt.Println(v3)
                v41 := GenS4{F0: (v1 <= 0.0), F1: 1, F2: 18446744073709551615}
                _ = fi36
                _ = v41
            }
            for fi42 := 0; fi42 < 5; fi42++ {
                var v43 uint8 = 0
                v5 += (fi42 - (v5 - fi35))
                _ = fi42
                _ = v43
            }
            _ = fi35
        }
    }
    v1 = 0.0
    var v44 bool = true
    fmt.Println(v2, v2)
    fmt.Println(v3)
    var v45 int64 = 0
    switch "" {
    case "a":
        for fi46 := 0; fi46 < 7; fi46++ {
            v2 += v2
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v3 += (v2 / ((2) | 1))
            _ = fi46
        }
    }
    var v47 int64 = v45
    v45 = (v45 << ((18446744073709551615) & 63))
    v48 := GenS5{F0: "hello", F1: -72, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v5
}
