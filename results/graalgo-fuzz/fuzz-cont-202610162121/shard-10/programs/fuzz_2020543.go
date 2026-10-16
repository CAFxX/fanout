package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int32 }
func genFn0() int64 {
    var v7 float32 = 1.5
    v8 := []int{130, 127}
    _ = v7
    _ = v8
    return 1
}
func genFn1() uint32 {
    var v9 int = -948
    v9 -= -9223372036854775808
    var v10 int32 = 42
    _ = v10
    _ = v9
    return 707
}

func main() {
    v1 := []string{"\n\t", ((("" + ("" + "a")) + "hello") + "a"), "a", ("" + ("" + "\n\t"))}
    var v2 uint16 = 1
    v3 := []uint32{360, 1}
    var v4 float64 = ((1.5 - 1.5) + float64(((0.0 * -2.25) + (0.1 - ((1e10 * float32(3.14159)) * 0.1)))))
    var v5 int32 = -2147483648
    v6 := []int{-2147483648, 2147483647, 755}
    v11 := GenS1{F0: 330, F1: 4294967295, F2: v2}
    if false {
        for fi12 := 0; fi12 < 1; fi12++ {
            _ = len(v6)
            if true {
                v5 -= 42
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v13 uint16 = v2
            _ = fi12
            _ = v13
        }
        fmt.Println(v4, v5, v5)
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 float32 = 0.0
                _ = fi15
                _ = v16
            }
            if ("a" < "") {
                v2 -= 2
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi14
        }
    } else {
        v4 -= v4
        fmt.Println(v4)
    }
    v4 -= v4
    var v17 uint64 = 2
    fmt.Println(v2, v4)
    v18 := GenS1{F0: -32768, F1: v17, F2: v2}
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "foo bar")
    }
    switch "foo bar" {
    case "hello":
        if true {
            var v19 int = int(v5)
            var v20 bool = true
            _ = v19
            _ = v20
        } else {
            _ = len(v3)
            if ((v2 ^ (v2 << ((42) & 63))) < v2) {
                if (!false) {
                    if ((v4 / v4) == -2.25) {
                        var v21 uint64 = 4294967295
                        switch "hello" {
                        case "a":
                            var v22 uint16 = v2
                            _ = v22
                        case "hello":
                            if false {
                                var v23 int8 = -128
                                fmt.Println(v5, v4)
                                var v24 float64 = ((1.5 + v4) / v4)
                                _ = v23
                                _ = v24
                            }
                        }
                        _ = v21
                    }
                    v3 = append(v3, 2147483648)
                }
            } else {
                var v25 int8 = 71
                for fi26 := 0; fi26 < 1; fi26++ {
                    v27 := GenS1{F0: int16(v25), F1: v17, F2: 73}
                    v28 := GenS1{F0: 1, F1: 271, F2: (^(v2))}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v25
            }
        }
    case "":
        if len(v3) > 0 {
            v3[0] = 1
        }
    }
    var v29 uint64 = (v17 & 2)
    fmt.Println(v5, v2)
    switch 4294967295 {
    case 9:
        var v30 int8 = 0
        _ = v30
    case -638:
        fmt.Println(v29, v2, v17)
    case 128:
        switch ("a" + "foo bar") {
        case "":
            var v31 uint16 = 846
            _ = v31
        case "a":
            v17 = 2
        case "\n\t":
            for fi32 := 0; fi32 < 8; fi32++ {
                if false {
                    for fi33 := 0; fi33 < 12; fi33++ {
                        v4 = (v4 + v4)
                        if false {
                            var v34 int = fi32
                            v4 = ((-2.25 / (float64(((((0.1 + 0.0) * 3.14159) / 1e10) - (0.1 * 1.5))) / ((((-2.25 - (((v4 * 3.14159) / (1e10 * v4)) * (v4 * -2.25))) * v4) * ((1.5 / ((v4 - float64(0.0)) + -2.25)) / ((v4 * float64(1e10)) + (float64((0.0 / 1e10)) / 1.0)))) - -2.25))) + ((v4 - v4) - 3.14159))
                            v35 := GenS1{F0: 32767, F1: (v29 >> ((716) & 63)), F2: uint16(v17)}
                            _ = v34
                            _ = v35
                        }
                        v4 = ((-2.25 - v4) + (3.14159 / float64(1e10)))
                        _ = fi33
                    }
                } else {
                    v36 := GenS2{F0: 338, F1: v17, F2: v17}
                    var v37 int64 = -849
                    v38 := GenS3{F0: (v5 - v5)}
                    _ = v36
                    _ = v37
                    _ = v38
                }
                _ = fi32
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
