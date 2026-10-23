package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 string; F2 string }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 int16 = -584
    var v3 uint8 = 1
    v4 := []uint32{65535, 384, 579, 2}
    v5 := []string{("foo bar" + ""), "hello"}
    v6 := GenS1{F0: 2}
    var v7 int = 128
    v8 := GenS2{F0: false, F1: 127}
    v9 := GenS3{F0: uint16(v3)}
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            switch "hello" {
            case "foo bar":
                if false {
                    fmt.Println(fi10)
                    v1["foo bar"] = (fi11 * fi11)
                }
                var v12 uint8 = v3
                _ = v12
            case "hello":
                var v13 bool = ((3.14159 == ((-2.25 * float32(0.0)) - 1e10)) && false)
                _ = v13
            case "a":
                var v14 uint64 = 832
                _ = v14
            }
            _ = fi11
        }
        fmt.Println(fi10, fi10)
        _ = fi10
    }
    fmt.Println(v3, v3, v2)
    var v15 uint8 = v3
    if (!(!true)) {
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v1)
        } else {
            fmt.Println(v7, v7, v15)
        }
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    } else {
        if true {
            v2 += (v2 << ((74) & 63))
            if true {
                if (v7 < -9223372036854775808) {
                    delete(v1, "\n\t")
                    v7 -= (v7 * v7)
                } else {
                    var v16 int = (v7 + (-(v7)))
                    if (!false) {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            var v18 uint16 = 2
                            if ((v15 + v3) <= v15) {
                                var v19 uint64 = 114
                                for fi20 := 0; fi20 < 6; fi20++ {
                                    fmt.Println(v7, fi20, v7)
                                    v18 -= (v18 * v18)
                                    v21 := GenS1{F0: 1}
                                    _ = fi20
                                    _ = v21
                                }
                                var v22 int8 = -1
                                _ = v19
                                _ = v22
                            } else {
                                for fi23 := 0; fi23 < 2; fi23++ {
                                    if true {
                                        _ = len(v4)
                                    } else {
                                        v24 := GenS4{F0: int8(v2), F1: "hello", F2: "hello"}
                                        var v25 uint8 = 127
                                        _ = v24
                                        _ = v25
                                    }
                                    _ = fi23
                                }
                                v3 -= 126
                            }
                            _ = fi17
                            _ = v18
                        }
                        for fi26 := 0; fi26 < 15; fi26++ {
                            v16 -= fi26
                            if false {
                                if (!(((1.5 + 3.14159) * (1e10 - 0.0)) < ((1e10 / 3.14159) / 1.0))) {
                                    var v27 int64 = 1
                                    switch fi26 {
                                    case 2147483647:
                                        switch (-(v7)) {
                                        case -9223372036854775808:
                                            if (false && false) {
                                                v28 := GenS5{F0: v3, F1: (v2 & v2)}
                                                v3 -= 0
                                                _ = v28
                                            }
                                        default:
                                            v29 := GenS2{F0: true, F1: v15}
                                            _ = v29
                                        }
                                    default:
                                        if (true || false) {
                                            var v30 string = ("a" + ((("a" + "a") + ("foo bar" + "")) + "\n\t"))
                                            v30 = (v30 + ("a" + "hello"))
                                            _ = v30
                                        }
                                    }
                                    var v31 uint8 = (v3 + v3)
                                    _ = v27
                                    _ = v31
                                }
                                v32 := GenS2{F0: true, F1: (v3 % ((255) | 1))}
                                _ = v32
                            } else {
                                for fi33 := 0; fi33 < 7; fi33++ {
                                    if (1e10 > -2.25) {
                                        switch fi26 {
                                        case 32767:
                                            v2 -= -128
                                        }
                                        if (0.1 == float64(((0.0 / 1e10) * (3.14159 - 1.5)))) {
                                            fmt.Println(v2, v3)
                                            v34 := GenS1{F0: 2}
                                            _ = v34
                                        } else {
                                            v7 -= -2147483648
                                        }
                                    } else {
                                        var v35 uint64 = 2
                                        v1["a"] = (fi33 + -47)
                                        _ = v35
                                    }
                                    v5 = append(v5, "foo bar")
                                    _ = fi33
                                }
                            }
                            if (0.1 == (1.5 + 0.1)) {
                                v15 = v15
                                v7 = v16
                                for fi36 := 0; fi36 < 10; fi36++ {
                                    v37 := GenS6{F0: 578}
                                    v3 -= (v15 * v15)
                                    _ = fi36
                                    _ = v37
                                }
                            } else {
                                var v38 string = "a"
                                _ = v38
                            }
                            _ = fi26
                        }
                    }
                    fmt.Println(v7, v16)
                    _ = v16
                }
                if false {
                    delete(v1, ("\n\t" + (("\n\t" + "foo bar") + "foo bar")))
                    v15 = 117
                    for fi39 := 0; fi39 < 16; fi39++ {
                        _ = len(v5)
                        _ = fi39
                    }
                } else {
                    for fi40 := 0; fi40 < 19; fi40++ {
                        if (true || false) {
                            v3 = (v15 / (((v3 / (((v15 | (^(v15)))) | 1))) | 1))
                            v7 = (fi40 ^ -845)
                        } else {
                            var v41 float32 = float32((1.5 + (0.1 * (1.5 - -2.25))))
                            _ = v41
                        }
                        v2 = (v2 % ((-86) | 1))
                        delete(v1, (("\n\t" + "a") + "a"))
                        _ = fi40
                    }
                }
            }
        } else {
            var v42 int64 = 9223372036854775807
            var v43 int64 = (v42 & v42)
            _ = v42
            _ = v43
        }
    }
    v3 -= (v15 & v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
