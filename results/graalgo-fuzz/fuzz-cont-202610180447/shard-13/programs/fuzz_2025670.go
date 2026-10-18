package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 int64; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint64; F2 int32 }
type GenS4 struct { F0 int32; F1 int64; F2 int64 }
type GenS5 struct { F0 string; F1 int }
func genFn0(p0 int16) int8 {
    p0 += p0
    var v4 int64 = 507
    var v5 string = ("foo bar" + "a")
    _ = p0
    _ = v4
    _ = v5
    return -128
}
func genFn1(p0 uint64) int {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{fi6, 434}
        p0 -= (p0 << ((65535) & 63))
        p0 += (p0 + (p0 << ((0) & 63)))
        _ = fi6
        _ = v7
    }
    var v8 string = ("hello" + "")
    _ = p0
    _ = v8
    return -1
}

func main() {
    var v1 int16 = 1000
    var v2 uint32 = 823
    var v3 int = -801
    v9 := []int{269}
    if true {
        v3 += v3
        v10 := GenS1{F0: 18446744073709551615, F1: -341, F2: (v2 - (v2 / (((v2 + v2)) | 1)))}
        if (128 > 166) {
            switch genFn1(2) {
            case 887:
                var v11 uint32 = 2147483648
                _ = v11
            case -1:
                if ((v2 << ((146) & 63)) >= (v2 % (((v2 << ((4294967295) & 63))) | 1))) {
                    v1 = 0
                }
            case 32768:
                if (int(v2) >= (v3 >> ((705) & 63))) {
                    var v12 int8 = 42
                    switch "" {
                    case "hello":
                        switch (v3 >> ((4294967295) & 63)) {
                        case 9223372036854775807:
                            switch v3 {
                            case 266:
                                switch (v3 % ((9223372036854775807) | 1)) {
                                case 9223372036854775807:
                                    for fi13 := 0; fi13 < 15; fi13++ {
                                        v1 = 0
                                        if true {
                                            var v14 float64 = 0.1
                                            _ = v14
                                        }
                                        _ = fi13
                                    }
                                default:
                                    var v15 uint32 = (^(v2))
                                    _ = v15
                                }
                            }
                        case 827:
                            for fi16 := 0; fi16 < 16; fi16++ {
                                v1 = (v1 * v1)
                                var v17 int8 = v12
                                var v18 string = "foo bar"
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        }
                    default:
                        v19 := GenS2{F0: 42, F1: 42, F2: 2}
                        _ = v19
                    }
                    v20 := GenS2{F0: -9223372036854775808, F1: -9223372036854775808, F2: 3}
                    _ = v12
                    _ = v20
                } else {
                    v3 = v3
                    switch "" {
                    case "foo bar":
                        var v21 int8 = 127
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v2 -= (v2 * (v2 >> ((65535) & 63)))
                            _ = fi22
                        }
                    }
                }
            }
        }
        _ = v10
    } else {
        if ("foo bar" != "hello") {
            switch ("foo bar" + (("a" + "foo bar") + ("\n\t" + ""))) {
            case "\n\t":
                v23 := GenS3{F0: 988, F1: 1, F2: 441}
                _ = v23
            case "a":
                v9 = append(v9, v3)
            default:
                v24 := GenS4{F0: 30, F1: 1, F2: -324}
                _ = v24
            }
            var v25 uint = 529
            v1 = v1
            _ = v25
        } else {
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 float32 = 3.14159
                _ = fi26
                _ = v27
            }
        }
        var v28 string = ("hello" + "hello")
        var v29 int = v3
        _ = v28
        _ = v29
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v30 int64 = -1
    for fi31 := 0; fi31 < 5; fi31++ {
        switch (fi31 + -2147483648) {
        case -1:
            if len(v9) > 0 {
                v9[0] = v3
            }
        }
        _ = fi31
    }
    switch ("hello" + ("a" + ("a" + "a"))) {
    case "hello":
        fmt.Println(v30)
    case "foo bar":
        fmt.Println(v1, v2)
    default:
        v32 := GenS1{F0: 1, F1: 42, F2: (v2 + v2)}
        _ = v32
    }
    var v33 uint64 = 351
    v30 += -811
    switch "\n\t" {
    case "a":
        var v34 uint8 = 128
        _ = v34
    case "foo bar":
        v35 := GenS5{F0: ("\n\t" + "\n\t"), F1: (v3 + genFn1(v33))}
        _ = v35
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        var v37 float64 = (3.14159 / 1.0)
        _ = fi36
        _ = v37
    }
    var v38 int8 = -128
    switch v3 {
    case 1:
        var v39 int64 = -9223372036854775808
        _ = v39
    default:
        v38 += (v38 & v38)
    }
    if ((^(v2)) > v2) {
        for fi40 := 0; fi40 < 15; fi40++ {
            fmt.Println(v3, fi40)
            var v41 int64 = v30
            _ = fi40
            _ = v41
        }
        var v42 int8 = (-(v38))
        _ = v42
    } else {
        var v43 float64 = 3.14159
        v43 += float64(1.5)
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v38
    _ = v9
}
