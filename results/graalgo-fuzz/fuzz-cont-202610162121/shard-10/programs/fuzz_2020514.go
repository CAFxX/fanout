package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint16; F1 int8; F2 int32 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int; F1 int32 }
func genFn0(p0 uint32) bool {
    v5 := []int{}
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 uint64 = 2
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint32 = p0
            var v10 int32 = 42
            var v11 int16 = 599
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    }
    var v12 bool = false
    if true {
        p0 = (p0 | 2147483647)
    }
    _ = p0
    _ = v12
    _ = v5
    return (9223372036854775807 >= 42)
}
func genFn1(p0 uint64) uint8 {
    var v13 bool = ((((float64(float32(float64(0.1))) + (-2.25 + 1.5)) - (1.5 / -2.25)) == -2.25) && (0.0 >= float32(((0.0 * 0.0) * 1.5))))
    p0 = p0
    p0 -= p0
    if (!v13) {
        var v14 float32 = (((-2.25 / -2.25) / (1.5 * 1e10)) + 0.0)
        _ = v14
    } else {
        v15 := []int{4294967295, 4294967295, 1}
        var v16 uint8 = 40
        _ = v15
        _ = v16
    }
    _ = p0
    _ = v13
    return 0
}

func main() {
    var v1 int8 = 0
    v2 := []int{}
    var v3 uint64 = 4294967295
    var v4 string = ("\n\t" + "foo bar")
    fmt.Println(v3, v4)
    var v17 uint = 0
    if true {
        v1 = v1
        var v18 int64 = -12
        _ = v18
    } else {
        for fi19 := 0; fi19 < 13; fi19++ {
            switch "hello" {
            case "a":
                if (v1 == -128) {
                    v3 -= v3
                }
            default:
                var v20 int32 = 2147483647
                _ = v20
            }
            switch ((v4 + ("a" + v4)) + (("hello" + v4) + "hello")) {
            case "\n\t":
                var v21 uint16 = 65535
                _ = v21
            case "a":
                fmt.Println(v17, v3)
            default:
                fmt.Println(v3, fi19)
            }
            v3 -= (v3 >> (((v17 & 257)) & 63))
            _ = fi19
        }
        if false {
            var v22 bool = (!(-630 > -9223372036854775808))
            switch ("foo bar" + ("a" + "")) {
            case "\n\t":
                var v23 uint32 = 1
                _ = v23
            case "":
                for fi24 := 0; fi24 < 9; fi24++ {
                    if ((v3 ^ 2) <= 1) {
                        v4 = v4
                        var v25 uint8 = 252
                        _ = v25
                    } else {
                        v17 += (v17 - (v17 / (((^(v17))) | 1)))
                        for fi26 := 0; fi26 < 17; fi26++ {
                            v27 := GenS1{F0: v4, F1: "foo bar"}
                            v2 = append(v2, -1)
                            _ = fi26
                            _ = v27
                        }
                    }
                    switch (fi24 >> (((v17 & v17)) & 63)) {
                    case -9223372036854775808:
                        v28 := GenS1{F0: (v4 + ""), F1: v4}
                        _ = v28
                    case 32767:
                        for fi29 := 0; fi29 < 6; fi29++ {
                            var v30 uint = v17
                            _ = fi29
                            _ = v30
                        }
                    default:
                        v4 = v4
                    }
                    _ = fi24
                }
            case "hello":
                if (v22 && v22) {
                    for fi31 := 0; fi31 < 11; fi31++ {
                        v1 = v1
                        _ = fi31
                    }
                    fmt.Println(v22, v1)
                } else {
                    v32 := GenS2{F0: 0, F1: -1, F2: 2147483647}
                    v33 := GenS3{F0: 32767, F1: v4}
                    var v34 int16 = 1000
                    _ = v32
                    _ = v33
                    _ = v34
                }
            default:
                v35 := GenS4{F0: 392, F1: -169}
                _ = v35
            }
            v1 = 37
            _ = v22
        }
    }
    var v36 string = "a"
    if (177 >= 915) {
        v4 = v4
        _ = len(v2)
    }
    var v37 int = -2147483648
    var v38 uint32 = 2
    switch (v37 ^ v37) {
    case 2147483647:
        fmt.Println(v17, v4, v1)
    case -983:
        var v39 int64 = -698
        _ = v39
    case 483:
        var v40 uint8 = genFn1((v3 % (((^(v3))) | 1)))
        _ = v40
    }
    v41 := GenS1{F0: (v4 + v4), F1: "foo bar"}
    var v42 float64 = (0.0 * 3.14159)
    fmt.Println(v36)
    var v43 string = v4
    var v44 uint64 = (v3 >> (((v17 << ((4294967295) & 63))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
