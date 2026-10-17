package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 string; F2 int32 }
type GenS4 struct { F0 string; F1 uint8; F2 int32 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint8) uint8 {
    v5 := GenS1{F0: 32767, F1: false}
    p0 = p0
    _ = p0
    _ = v5
    return p0
}
func genFn1() int32 {
    var v6 bool = true
    var v7 int = 292
    var v8 uint64 = 42
    v7 = -2147483648
    _ = v6
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    var v1 uint16 = 0
    var v2 int64 = 9223372036854775807
    var v3 float32 = (0.1 + 1.5)
    var v4 int32 = 137
    v9 := []int{0}
    v10 := GenS1{F0: (v1 * (v1 - v1)), F1: false}
    v11 := GenS2{F0: genFn0(1)}
    var v12 int64 = (v2 << ((18446744073709551615) & 63))
    v1 += v1
    switch (("hello" + "") + (("a" + "hello") + ((((("foo bar" + ("" + "")) + ("hello" + "\n\t")) + "a") + "foo bar") + "a"))) {
    case "\n\t":
        var v13 int8 = 1
        _ = v13
    case "hello":
        switch "a" {
        case "\n\t":
            if true {
                var v14 uint = 4294967295
                _ = v14
            } else {
                v4 += 0
            }
        case "foo bar":
            v15 := GenS3{F0: ((v4 * (v4 + 1)) != genFn1()), F1: ("hello" + ("" + "")), F2: genFn1()}
            _ = v15
        case "a":
            for fi16 := 0; fi16 < 3; fi16++ {
                v1 += (v1 - (v1 | (v1 ^ (v1 >> ((1) & 63)))))
                if (!false) {
                    var v17 float64 = (-2.25 + 0.0)
                    if false {
                        switch 2147483647 {
                        case 389:
                            var v18 bool = true
                            _ = v18
                        case 32768:
                            if false {
                                switch "a" {
                                case "\n\t":
                                    switch "" {
                                    case "\n\t":
                                        for fi19 := 0; fi19 < 2; fi19++ {
                                            v2 -= v12
                                            _ = fi19
                                        }
                                    case "":
                                        v20 := GenS3{F0: true, F1: "a", F2: genFn1()}
                                        _ = v20
                                    case "hello":
                                        v21 := GenS4{F0: "foo bar", F1: genFn0(genFn0(255)), F2: v4}
                                        _ = v21
                                    default:
                                        switch fi16 {
                                        case -124:
                                            var v22 uint32 = 344
                                            _ = v22
                                        case 0:
                                            var v23 uint8 = 210
                                            _ = v23
                                        default:
                                            fmt.Println(v4)
                                        }
                                    }
                                case "":
                                    var v24 float64 = v17
                                    _ = v24
                                case "foo bar":
                                    var v25 uint16 = (v1 * 864)
                                    _ = v25
                                default:
                                    var v26 int16 = 42
                                    _ = v26
                                }
                                fmt.Println(v4)
                                for fi27 := 0; fi27 < 7; fi27++ {
                                    if (true && (100 != 42)) {
                                        if (117 <= 2147483647) {
                                            v12 = v12
                                            v28 := GenS5{F0: genFn1()}
                                            _ = v28
                                        } else {
                                            v12 = (v12 & -9223372036854775808)
                                            fmt.Println(v2, fi27, v17)
                                            fmt.Println(v17, v3, v1)
                                        }
                                        var v29 bool = (!true)
                                        _ = v29
                                    }
                                    switch fi16 {
                                    case 42:
                                        var v30 float32 = 0.0
                                        _ = v30
                                    }
                                    if false {
                                        if true {
                                            var v31 int64 = (v12 & 394)
                                            v3 = v3
                                            var v32 uint8 = 128
                                            _ = v31
                                            _ = v32
                                        } else {
                                            v9 = append(v9, (-(fi16)))
                                            v1 += 0
                                        }
                                    } else {
                                        if (!true) {
                                            v2 = v12
                                        } else {
                                            var v33 int8 = -128
                                            v2 -= (v2 / ((v2) | 1))
                                            v12 += v12
                                            _ = v33
                                        }
                                        if len(v9) > 0 {
                                            v9[0] = (fi27 % ((fi16) | 1))
                                        }
                                    }
                                    _ = fi27
                                }
                            } else {
                                switch (fi16 - (fi16 << ((2) & 63))) {
                                case 4294967295:
                                    v34 := GenS2{F0: 255}
                                    _ = v34
                                case 1:
                                    v12 -= (v12 / ((v12) | 1))
                                case -153:
                                    if false {
                                        var v35 int = (-(fi16))
                                        switch "\n\t" {
                                        case "":
                                            switch 42 {
                                            case 128:
                                                var v36 int = v35
                                                _ = v36
                                            case 1:
                                                switch "" {
                                                case "\n\t":
                                                    var v37 string = "a"
                                                    _ = v37
                                                case "foo bar":
                                                    v38 := GenS2{F0: 127}
                                                    _ = v38
                                                case "hello":
                                                    switch (("\n\t" + "foo bar") + "\n\t") {
                                                    case "foo bar":
                                                        fmt.Println(v2, v4)
                                                    case "hello":
                                                        var v39 uint = 4294967295
                                                        _ = v39
                                                    default:
                                                        if false {
                                                            fmt.Println(v4)
                                                            v2 = (^(v12))
                                                        }
                                                    }
                                                }
                                            case 32767:
                                                v17 = 0.1
                                            }
                                        case "a":
                                            v40 := GenS2{F0: 2}
                                            _ = v40
                                        case "\n\t":
                                            fmt.Println(v3)
                                        default:
                                            v41 := GenS4{F0: "\n\t", F1: 2, F2: v4}
                                            _ = v41
                                        }
                                        _ = v35
                                    } else {
                                        v12 += v2
                                        v4 -= v4
                                        for fi42 := 0; fi42 < 18; fi42++ {
                                            _ = len(v9)
                                            _ = fi42
                                        }
                                    }
                                }
                            }
                        }
                    }
                    v17 -= (v17 / v17)
                    _ = v17
                } else {
                    var v43 float64 = (1e10 - ((float64((v3 * (((((float32(float64(((v3 - 0.1) / 1.0))) - v3) - v3) * (v3 / v3)) - float32(((3.14159 * float64(3.14159)) - float64(v3)))) - 0.0))) * 3.14159) / 1.0))
                    _ = v43
                }
                _ = fi16
            }
        default:
            v2 = v2
        }
    case "a":
        var v44 int16 = 1000
        _ = v44
    default:
        fmt.Println(v12, v2)
    }
    v45 := GenS3{F0: (4294967295 > 42), F1: ("a" + "hello"), F2: -834}
    _ = len(v9)
    switch ("\n\t" + "a") {
    case "a":
        for fi46 := 0; fi46 < 10; fi46++ {
            v12 = v2
            _ = fi46
        }
    case "foo bar":
        v47 := GenS4{F0: ("" + ""), F1: 203, F2: v4}
        _ = v47
    case "":
        v4 -= genFn1()
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v9
}
