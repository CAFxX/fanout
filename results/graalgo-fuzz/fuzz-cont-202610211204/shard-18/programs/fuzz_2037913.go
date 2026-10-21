package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 string; F1 uint8; F2 int32 }
type GenS4 struct { F0 int; F1 int8; F2 int32 }
type GenS5 struct { F0 uint64; F1 int16; F2 int32 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 uint8; F1 int16; F2 int }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 uint64; F1 int32; F2 bool }

func main() {
    var v1 int = 128
    var v2 float32 = (3.14159 / 1.5)
    var v3 uint64 = 810
    v4 := GenS1{F0: 0, F1: -1, F2: (v1 & -9223372036854775808)}
    v5 := GenS2{F0: (("a" + "hello") + (("a" + "hello") + "hello")), F1: 42}
    if true {
        var v6 float32 = -2.25
        v7 := GenS3{F0: (("a" + "") + "foo bar"), F1: 127, F2: -1}
        _ = v6
        _ = v7
    }
    v8 := []int{v1, v1}
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 float64 = ((0.0 * (0.1 * ((0.1 + (float64((v2 * (v2 - -2.25))) / (0.1 * (float64(v2) + 3.14159)))) + -2.25))) + (1e10 + 1.5))
            if false {
                v8 = append(v8, (fi10 ^ fi9))
                if false {
                    fmt.Println(v3, fi9, v2)
                    switch 4294967295 {
                    case 0:
                        v1 -= (v1 | v1)
                        v1 = (fi9 - fi10)
                        v12 := GenS4{F0: (fi9 >> ((2) & 63)), F1: 1, F2: 2147483647}
                        v11 = v11
                        _ = v12
                    case -52:
                        if (1 != 18446744073709551615) {
                            v13 := GenS3{F0: (("" + ("a" + "foo bar")) + (("foo bar" + "foo bar") + "a")), F1: 0, F2: -577}
                            _ = v13
                        } else {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        }
                    case 44:
                        switch "" {
                        case "foo bar":
                            v14 := GenS5{F0: (v3 | (v3 >> ((65535) & 63))), F1: -480, F2: 2147483647}
                            _ = v14
                        default:
                            fmt.Println(v11, v11)
                        }
                    }
                    v1 += v1
                }
                v2 = (v2 + 0.1)
                var v15 uint8 = 223
                _ = v15
            }
            _ = fi10
            _ = v11
        }
        if (float32((1.5 + 0.1)) != (float32(3.14159) / v2)) {
            var v16 string = ""
            v1 -= 1
            _ = v16
        } else {
            var v17 string = "hello"
            v2 -= float32(float64(((v2 - -2.25) - 0.0)))
            _ = v17
        }
        _ = fi9
    }
    v3 = v3
    switch "a" {
    case "":
        v18 := GenS6{F0: 496, F1: 2}
        _ = v18
    }
    fmt.Println(v1)
    switch "" {
    case "hello":
        v1 = (v1 % (((v1 ^ (^(v1)))) | 1))
    case "\n\t":
        v1 = -129
    case "foo bar":
        if true {
            for fi19 := 0; fi19 < 3; fi19++ {
                if len(v8) > 0 {
                    v8[0] = v1
                }
                _ = fi19
            }
            for fi20 := 0; fi20 < 12; fi20++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                v21 := GenS7{F0: 18446744073709551615, F1: (255 != 255)}
                if len(v8) > 0 {
                    v8[0] = fi20
                }
                _ = fi20
                _ = v21
            }
        }
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        if len(v8) > 0 {
            v8[0] = 32767
        }
        _ = fi22
    }
    switch "" {
    case "a":
        for fi23 := 0; fi23 < 6; fi23++ {
            for fi24 := 0; fi24 < 6; fi24++ {
                if (!false) {
                    var v25 int32 = -185
                    var v26 bool = true
                    v27 := GenS5{F0: v3, F1: 1, F2: -2147483648}
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS4{F0: 1, F1: -1, F2: 0}
                    var v29 uint = 65535
                    _ = v28
                    _ = v29
                }
                var v30 int16 = -793
                _ = fi24
                _ = v30
            }
            v31 := GenS2{F0: "hello", F1: 2147483647}
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 uint8 = 42
                var v34 bool = (65535 <= 812)
                if (!((((uint(v33) >= 420) && (true && v34)) || v34) || (true && v34))) {
                    if (252 <= v33) {
                        if len(v8) > 0 {
                            v8[0] = (-(fi32))
                        }
                        fmt.Println(v1)
                    }
                } else {
                    v35 := GenS8{F0: (v33 & v33), F1: -32768, F2: fi32}
                    v8 = append(v8, 1)
                    if (1000 != 42) {
                        fmt.Println(v1, v1, fi32)
                    } else {
                        var v36 int32 = -1
                        fmt.Println(v3, v33)
                        switch ("\n\t" + "") {
                        case "foo bar":
                            var v37 int16 = 0
                            _ = v37
                        case "a":
                            switch (fi23 / ((0) | 1)) {
                            case 32768:
                                if true {
                                    v2 = v2
                                } else {
                                    v8 = append(v8, 32767)
                                    v3 = 0
                                    if (v34 && v34) {
                                        if ((fi23 >> ((65535) & 63)) < -360) {
                                            v3 += 217
                                            fmt.Println(v1)
                                        } else {
                                            v38 := GenS7{F0: 1, F1: true}
                                            v39 := GenS9{F0: (-(fi23)), F1: 380}
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                }
                            case 181:
                                v40 := GenS8{F0: 2, F1: 440, F2: v1}
                                _ = v40
                            case 128:
                                v2 += (float32((3.14159 * 3.14159)) / 1e10)
                            default:
                                if v34 {
                                    var v41 int8 = 42
                                    _ = v41
                                } else {
                                    v2 = ((v2 * v2) * (v2 / 0.1))
                                    var v42 int16 = 33
                                    _ = v42
                                }
                            }
                        default:
                            switch (v1 / ((fi32) | 1)) {
                            case 32767:
                                if ((("foo bar" + ("\n\t" + "")) + "foo bar") > "hello") {
                                    var v43 uint8 = v33
                                    var v44 string = "\n\t"
                                    v43 = 128
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    _ = v43
                                    _ = v44
                                }
                            case 32768:
                                v2 -= (v2 / 1.5)
                            }
                        }
                        _ = v36
                    }
                    _ = v35
                }
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi23
            _ = v31
        }
    case "hello":
        if len(v8) > 0 {
            v8[0] = 127
        }
    case "":
        v45 := GenS10{F0: (v3 ^ (v3 | v3)), F1: 1, F2: true}
        _ = v45
    }
    var v46 string = ("hello" + "a")
    v3 = (v3 * (^(v3)))
    var v47 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v8
}
