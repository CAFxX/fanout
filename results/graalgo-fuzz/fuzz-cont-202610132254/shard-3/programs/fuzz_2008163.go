package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 bool; F2 uint }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint8; F2 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 bool }
func genFn0(p0 uint8, p1 bool, p2 uint) int32 {
    v6 := GenS1{F0: 1000}
    var v7 uint16 = 124
    v8 := GenS2{F0: (p2 << ((p2) & 63)), F1: (1 < 32767), F2: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint8 = 164
    var v2 int64 = 1
    var v3 uint32 = 2147483647
    var v4 int = 632
    var v5 uint = 993
    v9 := GenS3{F0: (-(v4)), F1: 171}
    var v10 int8 = 0
    var v11 bool = (42 > (v1 ^ (v1 + v1)))
    v10 = -19
    switch "hello" {
    case "hello":
        v11 = v11
    case "":
        switch ("a" + ("hello" + "a")) {
        case "hello":
            var v12 uint = (v5 + (^(v5)))
            _ = v12
        case "\n\t":
            var v13 uint64 = 967
            _ = v13
        default:
            v14 := []int{-856, v4, (v4 / (((v4 >> (((v5 / ((65535) | 1))) & 63))) | 1)), v4}
            _ = v14
        }
    case "foo bar":
        v15 := GenS3{F0: v4, F1: 18446744073709551615}
        _ = v15
    default:
        if (((0.0 * 1.5) + (float64(0.0) + (float64(0.1) / -2.25))) < ((1.5 + ((((((float64(((float32((0.0 / 3.14159)) - (0.1 + 3.14159)) / ((-2.25 / 1.0) / -2.25))) * 3.14159) / 3.14159) - (3.14159 + -2.25)) / 1.0) * 0.0) - (3.14159 / 1e10))) / 1.0)) {
            v2 = -1
            for fi16 := 0; fi16 < 16; fi16++ {
                if v11 {
                    fmt.Println(v4)
                    var v17 uint = 868
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        var v19 uint32 = v3
                        var v20 int32 = 42
                        var v21 int = fi18
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v2 = (v2 + (v2 * (v2 >> ((v5) & 63))))
                }
                var v22 int64 = v2
                _ = fi16
                _ = v22
            }
            for fi23 := 0; fi23 < 5; fi23++ {
                if (!("\n\t" >= ("foo bar" + "foo bar"))) {
                    v4 = (v4 - (fi23 * fi23))
                    for fi24 := 0; fi24 < 10; fi24++ {
                        var v25 int = (-(fi23))
                        if true {
                            var v26 float32 = ((1e10 + -2.25) - float32(0.1))
                            switch "hello" {
                            case "hello":
                                v4 += (fi24 << (((v5 ^ uint(v1))) & 63))
                            }
                            if (float32(-2.25) != v26) {
                                switch fi23 {
                                case 1:
                                    switch "" {
                                    case "":
                                        var v27 uint64 = 18446744073709551615
                                        _ = v27
                                    case "a":
                                        v26 = (v26 / v26)
                                        var v28 int = 2147483647
                                        v29 := GenS4{F0: -405}
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                var v30 int8 = (v10 / ((v10) | 1))
                                if v11 {
                                    fmt.Println(v26, fi24, v11)
                                    v31 := GenS5{F0: 1, F1: (v1 >> ((787) & 63)), F2: true}
                                    if (!false) {
                                        if true {
                                            v32 := GenS6{F0: (-(fi23))}
                                            fmt.Println(v26)
                                            var v33 int16 = 1000
                                            _ = v32
                                            _ = v33
                                        }
                                        if (("a" + ("a" + "a")) >= ("\n\t" + ("\n\t" + ""))) {
                                            var v34 uint32 = 2
                                            _ = v34
                                        } else {
                                            v35 := GenS7{F0: fi23, F1: (v10 == int8(v5))}
                                            if (0.0 < 3.14159) {
                                                v4 = 127
                                            }
                                            v1 = v1
                                            v36 := []int{-236}
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                    _ = v31
                                } else {
                                    v30 += 70
                                    v30 = (v10 ^ (v30 >> (((v5 << ((421) & 63))) & 63)))
                                    v25 = (fi24 << ((v5) & 63))
                                }
                                _ = v30
                            }
                            var v37 int = 0
                            if ("hello" <= (("foo bar" + "") + "a")) {
                                var v38 int64 = -854
                                var v39 bool = v11
                                v3 -= (v3 << ((v5) & 63))
                                _ = v38
                                _ = v39
                            }
                            v40 := []int{v37, 2147483647, (fi24 << (((v5 / ((v5) | 1))) & 63)), v4}
                            _ = v26
                            _ = v37
                            _ = v40
                        } else {
                            if false {
                                v5 = v5
                                v41 := GenS7{F0: (fi23 - v4), F1: (-2.25 > ((float64(-2.25) / 1.0) * ((float64((-2.25 / (1e10 - 0.1))) * 0.0) * float64((float32(0.1) + -2.25)))))}
                                _ = v41
                            }
                            switch v4 {
                            case -1:
                                v42 := []int{v25, (fi23 - fi24)}
                                _ = v42
                            }
                            v4 = v25
                        }
                        fmt.Println(v25)
                        _ = fi24
                        _ = v25
                    }
                    var v43 uint32 = uint32(v1)
                    _ = v43
                } else {
                    v44 := GenS2{F0: v5, F1: (v11 && (1 > 265)), F2: 612}
                    v1 = v1
                    if v11 {
                        fmt.Println(v3, fi23)
                    }
                    _ = v44
                }
                var v45 int = (v4 << ((v5) & 63))
                _ = fi23
                _ = v45
            }
        }
    }
    v1 -= (v1 - (v1 * (v1 << (((v5 / (((v5 + (v5 / (((v5 * uint(v1))) | 1)))) | 1))) & 63))))
    for fi46 := 0; fi46 < 11; fi46++ {
        var v47 uint16 = 2
        var v48 bool = (v5 >= (v5 >> (((^(v5))) & 63)))
        switch "foo bar" {
        case "":
            v49 := GenS3{F0: (-(v4)), F1: 18446744073709551615}
            _ = v49
        }
        _ = fi46
        _ = v47
        _ = v48
    }
    if ((^(v4)) != v4) {
        v50 := GenS3{F0: (v4 + (v4 >> ((v5) & 63))), F1: 409}
        var v51 bool = (!(232 >= 2147483647))
        _ = v50
        _ = v51
    } else {
        fmt.Println(v2, v11, v1)
        switch ("a" + "") {
        case "a":
            var v52 float32 = 0.0
            _ = v52
        }
        v53 := []int{}
        _ = v53
    }
    v10 -= (v10 + (v10 / (((^(v10))) | 1)))
    switch (-(v4)) {
    case 0:
        var v54 uint16 = 0
        _ = v54
    }
    if (42 > 1) {
        var v55 int = v4
        switch 128 {
        case 127:
            v56 := GenS7{F0: (v55 % (((v4 | (v4 * (v4 | (v55 >> ((v5) & 63)))))) | 1)), F1: (v1 > 137)}
            _ = v56
        default:
            if ((!v11) && (1.5 <= (3.14159 * 3.14159))) {
                v57 := GenS4{F0: 551}
                var v58 int32 = 1
                v2 += (v2 / (((^(v2))) | 1))
                _ = v57
                _ = v58
            }
        }
        for fi59 := 0; fi59 < 6; fi59++ {
            if v11 {
                v11 = (0.1 > 0.1)
                switch (v4 & v55) {
                case 127:
                    var v60 uint64 = 2
                    _ = v60
                case 42:
                    fmt.Println(v55, v3)
                case -2147483648:
                    for fi61 := 0; fi61 < 18; fi61++ {
                        v5 -= 42
                        v62 := GenS5{F0: 4294967295, F1: 1, F2: v11}
                        _ = fi61
                        _ = v62
                    }
                default:
                    v4 = (v4 | v55)
                }
            }
            v10 -= (v10 & v10)
            v5 = v5
            _ = fi59
        }
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
