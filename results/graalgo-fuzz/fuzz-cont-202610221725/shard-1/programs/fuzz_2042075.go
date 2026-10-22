package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint8; F1 uint16; F2 int8 }
func genFn0() string {
    var v7 uint16 = 65535
    var v8 uint = 222
    v8 -= v8
    v8 = 255
    v8 -= 581
    _ = v7
    _ = v8
    return (("a" + "\n\t") + "")
}

func main() {
    var v1 int16 = 42
    var v2 uint64 = 0
    var v3 int8 = -1
    var v4 bool = (32767 != 959)
    var v5 string = ("" + "hello")
    v6 := []int{-9223372036854775808, -2147483648, 0}
    v1 -= (v1 >> ((255) & 63))
    if len(v6) > 0 {
        v6[0] = -78
    }
    switch 32768 {
    case -1:
        var v9 uint64 = 2
        _ = v9
    }
    var v10 string = v5
    switch v10 {
    case "hello":
        v3 = (-(v3))
    default:
        var v11 uint = 273
        _ = v11
    }
    _ = len(v6)
    fmt.Println(v2, v10)
    _ = len(v6)
    v1 += (v1 ^ v1)
    v6 = append(v6, 128)
    switch (v5 + v10) {
    case "hello":
        v6 = append(v6, -9223372036854775808)
    case "":
        v1 -= -32768
    case "a":
        for fi12 := 0; fi12 < 4; fi12++ {
            if v4 {
                var v13 uint32 = 1
                _ = v13
            } else {
                v5 = (v5 + (("foo bar" + (("" + v10) + v5)) + v10))
                v1 = 32767
            }
            _ = fi12
        }
    }
    v1 = v1
    v2 = 4294967295
    if v4 {
        var v14 int = 0
        var v15 int16 = -1
        _ = v14
        _ = v15
    } else {
        v16 := GenS1{F0: (^(v2)), F1: false, F2: 42}
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := GenS1{F0: (v2 >> ((65535) & 63)), F1: v4, F2: 254}
            if v4 {
                fmt.Println(v1, v3, v2)
                fmt.Println(v4)
            } else {
                switch v5 {
                case "foo bar":
                    v6 = append(v6, fi17)
                default:
                    _ = len(v6)
                }
            }
            _ = fi17
            _ = v18
        }
        if false {
            if v4 {
                switch v10 {
                case "hello":
                    var v19 string = v5
                    _ = v19
                case "\n\t":
                    switch 32767 {
                    case 2147483647:
                        var v20 int16 = (v1 | 1)
                        _ = v20
                    case 32767:
                        for fi21 := 0; fi21 < 1; fi21++ {
                            var v22 uint16 = 1
                            _ = len(v6)
                            var v23 float32 = 0.1
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    default:
                        _ = len(v6)
                    }
                case "":
                    var v24 uint64 = (^(v2))
                    _ = v24
                }
                if true {
                    var v25 float64 = ((0.1 + (float64(-2.25) / 3.14159)) - (3.14159 - 0.1))
                    _ = v25
                } else {
                    v26 := GenS1{F0: (v2 << ((18446744073709551615) & 63)), F1: (((((((3.14159 - 0.1) * 1.5) / -2.25) - 0.0) + (0.0 + (0.1 * (1.5 / 1.0)))) >= (0.0 + 0.1)) || v4), F2: 2147483648}
                    if len(v6) > 0 {
                        v6[0] = 32768
                    }
                    if (!true) {
                        v27 := GenS1{F0: 413, F1: ((v2 >> ((65535) & 63)) == 42), F2: 65535}
                        var v28 int = -9223372036854775808
                        var v29 uint = 255
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = v26
                }
            } else {
                if v4 {
                    if (1 >= 560) {
                        var v30 string = v5
                        _ = v30
                    } else {
                        var v31 bool = false
                        if (177 >= 132) {
                            for fi32 := 0; fi32 < 12; fi32++ {
                                if len(v6) > 0 {
                                    v6[0] = fi32
                                }
                                for fi33 := 0; fi33 < 18; fi33++ {
                                    v4 = (!v4)
                                    v34 := GenS1{F0: v2, F1: ((196 < 298) && v4), F2: 888}
                                    v4 = v4
                                    _ = fi33
                                    _ = v34
                                }
                                v31 = (255 == 225)
                                _ = fi32
                            }
                            _ = len(v6)
                        } else {
                            var v35 uint8 = uint8(v2)
                            _ = v35
                        }
                        fmt.Println(v1, v3)
                        _ = v31
                    }
                    for fi36 := 0; fi36 < 5; fi36++ {
                        v4 = (!(int64(v1) < -25))
                        if (983 >= 0) {
                            v1 += (v1 - (-(v1)))
                            v37 := GenS2{F0: v1, F1: true}
                            _ = v37
                        }
                        var v38 int8 = (v3 + int8(v2))
                        _ = fi36
                        _ = v38
                    }
                }
            }
        } else {
            v39 := GenS3{F0: 8, F1: 0, F2: (v3 + v3)}
            var v40 uint64 = v2
            for fi41 := 0; fi41 < 18; fi41++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi41
            }
            _ = v39
            _ = v40
        }
        _ = v16
    }
    var v42 int32 = -1
    v5 = v5
    var v43 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
