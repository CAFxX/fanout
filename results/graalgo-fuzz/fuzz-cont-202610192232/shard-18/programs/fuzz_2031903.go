package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int; F2 uint8 }
type GenS3 struct { F0 uint32; F1 string; F2 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
func genFn0() uint8 {
    switch 960 {
    case 4294967295:
        var v4 int64 = -545
        _ = v4
    default:
        var v5 uint32 = 997
        _ = v5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            if true {
                var v8 float64 = (3.14159 + ((0.0 / 1.0) + -2.25))
                var v9 float32 = 0.0
                _ = v8
                _ = v9
            }
            fmt.Println(fi6, fi6)
            _ = fi7
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 uint16 = 556
            _ = fi10
            _ = v11
        }
        _ = fi6
    }
    if ((float64(1e10) * (-2.25 / 1.0)) < 3.14159) {
        var v12 int64 = -9223372036854775808
        _ = v12
    } else {
        v13 := GenS1{F0: (false && false)}
        v14 := GenS2{F0: 240, F1: -9223372036854775808, F2: 213}
        _ = v13
        _ = v14
    }
    var v15 float32 = (float32((0.1 - -2.25)) + float32(float64(-2.25)))
    v15 += (((float32((0.1 + 3.14159)) - (float32((0.1 / 1.0)) / 1.0)) * ((((v15 / v15) - float32(1e10)) * v15) / v15)) + v15)
    _ = v15
    return 111
}
func genFn1(p0 uint16, p1 int32) uint16 {
    var v16 uint = 1
    p0 = 0
    p1 -= (p1 / (((p1 & -1)) | 1))
    for fi17 := 0; fi17 < 10; fi17++ {
        v16 -= v16
        switch ("a" + "a") {
        case "a":
            var v18 uint = 0
            _ = v18
        case "":
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 uint = v16
                v21 := GenS2{F0: p0, F1: fi19, F2: 41}
                _ = fi19
                _ = v20
                _ = v21
            }
        case "foo bar":
            v22 := GenS3{F0: 4294967295, F1: "hello", F2: 445}
            _ = v22
        }
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = v16
    return p0
}

func main() {
    var v1 int32 = 141
    var v2 uint = 255
    v3 := []string{"a", "hello", "\n\t", ("hello" + "")}
    switch ("foo bar" + (("hello" + ("hello" + "a")) + ("a" + ""))) {
    case "hello":
        if (uint32(v2) > 2147483647) {
            if (2 == 0) {
                for fi23 := 0; fi23 < 1; fi23++ {
                    var v24 uint32 = 2
                    _ = fi23
                    _ = v24
                }
                switch ("\n\t" + "foo bar") {
                case "a":
                    v2 = 822
                case "foo bar":
                    var v25 int8 = 100
                    _ = v25
                case "":
                    var v26 int = 42
                    _ = v26
                default:
                    _ = len(v3)
                }
            } else {
                fmt.Println(v2, v2)
                if (!false) {
                    v1 += 1
                    var v27 string = "foo bar"
                    v27 = v27
                    _ = v27
                } else {
                    switch 1 {
                    case 4294967295:
                        if (458 <= 42) {
                            switch -648 {
                            case -9223372036854775808:
                                var v28 int64 = 372
                                _ = v28
                            }
                            var v29 uint16 = 255
                            v29 = genFn1(v29, int32(v29))
                            _ = v29
                        }
                    }
                    v1 = v1
                }
                var v30 int = -523
                _ = v30
            }
        } else {
            v31 := GenS3{F0: 2, F1: "foo bar", F2: 509}
            switch -477 {
            case 669:
                fmt.Println(v2)
            case -361:
                var v32 uint8 = genFn0()
                _ = v32
            }
            for fi33 := 0; fi33 < 14; fi33++ {
                if ((42 <= 507) && ((-(fi33)) <= fi33)) {
                    v2 -= 18446744073709551615
                    v34 := GenS4{F0: (v2 + 18446744073709551615)}
                    v35 := GenS2{F0: 32767, F1: (fi33 >> (((v2 % (((^(v2))) | 1))) & 63)), F2: 128}
                    _ = v34
                    _ = v35
                } else {
                    for fi36 := 0; fi36 < 3; fi36++ {
                        var v37 int = fi36
                        v38 := GenS5{F0: 52}
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                }
                _ = fi33
            }
            _ = v31
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        var v40 uint16 = 310
        _ = fi39
        _ = v40
    }
    var v41 bool = (v1 >= (v1 ^ (v1 % (((v1 & v1)) | 1))))
    var v42 int16 = 0
    fmt.Println(v1)
    var v43 uint16 = genFn1(706, 1)
    var v44 uint16 = (v43 | (^(v43)))
    var v45 uint16 = v44
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
