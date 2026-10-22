package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int8; F1 int8 }
func genFn0(p0 bool, p1 uint32) uint {
    fmt.Println(p0, p0)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 int) string {
    var v8 int16 = 247
    var v9 int32 = -10
    v9 = v9
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return ((("hello" + "hello") + ("foo bar" + "hello")) + "")
}

func main() {
    var v1 int64 = 42
    var v2 float64 = ((float64(0.0) - 0.0) / 1.0)
    v3 := []int{32767, 9223372036854775807, 1, 519}
    v4 := []string{"hello", "\n\t", (("" + "a") + "foo bar"), "a"}
    var v5 uint16 = 42
    v6 := []uint32{67, 2147483647, 0}
    if (!((65535 > 0) && true)) {
        v5 = (v5 / ((v5) | 1))
        v10 := GenS1{F0: "foo bar", F1: 42, F2: 65535}
        _ = v10
    }
    var v11 bool = (!true)
    v11 = (!v11)
    v11 = ((0.1 > ((float64(1.5) * (0.0 * (((v2 / ((((float64(1e10) / v2) - ((float64(float32(float64((1e10 + ((0.1 / (float32(1e10) + -2.25)) + 0.0))))) - 0.1) / (v2 + v2))) * v2) + 1e10)) * (3.14159 - ((((v2 / v2) + (v2 + (float64((-2.25 + (float32((((0.1 - 3.14159) + v2) * 0.0)) - float32(v2)))) + 1.5))) * float64(3.14159)) - v2))) / v2))) - -2.25)) && ((315 < 4294967295) && (v11 || ((0.1 - 0.0) == (0.0 / -2.25)))))
    v5 -= (v5 & (v5 >> ((157) & 63)))
    var v12 bool = (true || v11)
    v3 = append(v3, 498)
    fmt.Println(v1, v2, v11)
    if false {
        switch "\n\t" {
        case "hello":
            fmt.Println(v1)
        case "foo bar":
            v1 = (v1 ^ v1)
        case "\n\t":
            var v13 string = ("\n\t" + "\n\t")
            _ = v13
        default:
            if v11 {
                v1 += (^(v1))
                if len(v6) > 0 {
                    v6[0] = 42
                }
                for fi14 := 0; fi14 < 3; fi14++ {
                    v12 = v12
                    switch "\n\t" {
                    case "":
                        switch (("foo bar" + "foo bar") + "\n\t") {
                        case "a":
                            switch fi14 {
                            case -821:
                                _ = len(v4)
                            default:
                                v12 = v12
                            }
                        }
                    case "foo bar":
                        for fi15 := 0; fi15 < 2; fi15++ {
                            fmt.Println(v1)
                            fmt.Println(v1, fi14, v5)
                            var v16 float32 = (3.14159 + 1.5)
                            _ = fi15
                            _ = v16
                        }
                    }
                    _ = fi14
                }
            }
        }
        v17 := GenS1{F0: "a", F1: 0, F2: 552}
        v18 := GenS2{F0: v11}
        _ = v17
        _ = v18
    }
    var v19 int32 = 42
    var v20 uint16 = v5
    var v21 uint64 = 612
    if v11 {
        v22 := GenS1{F0: ("foo bar" + (("a" + ("\n\t" + ("a" + "foo bar"))) + ("a" + ("" + "")))), F1: 65535, F2: 2}
        switch (("\n\t" + "") + ((("a" + "") + "\n\t") + "a")) {
        case "":
            v1 += v1
        default:
            switch "a" {
            case "hello":
                for fi23 := 0; fi23 < 11; fi23++ {
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                    var v24 float32 = 0.1
                    _ = fi23
                    _ = v24
                }
            case "foo bar":
                var v25 uint8 = 38
                _ = v25
            default:
                if true {
                    if v12 {
                        v1 = v1
                    }
                } else {
                    for fi26 := 0; fi26 < 2; fi26++ {
                        fmt.Println(v5)
                        switch (^(fi26)) {
                        case 32767:
                            for fi27 := 0; fi27 < 8; fi27++ {
                                if v12 {
                                    v5 -= (v20 ^ (v20 & v20))
                                    switch (fi27 >> ((2) & 63)) {
                                    case 32768:
                                        v4 = append(v4, "hello")
                                    case -863:
                                        var v28 uint16 = v20
                                        _ = v28
                                    case -981:
                                        if true {
                                            if (-2.25 > 1.5) {
                                                var v29 uint32 = 654
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                switch fi26 {
                                                case 166:
                                                    fmt.Println(v29, fi26)
                                                case -2147483648:
                                                    if v12 {
                                                        if len(v6) > 0 {
                                                            v6[0] = (v29 << ((1) & 63))
                                                        }
                                                    }
                                                    var v30 int64 = v1
                                                    _ = v30
                                                default:
                                                    v31 := GenS1{F0: ("\n\t" + ((("" + "") + "\n\t") + "foo bar")), F1: uint32(v21), F2: 42}
                                                    switch fi27 {
                                                    case 42:
                                                        v32 := GenS1{F0: "", F1: (^(v29)), F2: v29}
                                                        _ = v32
                                                    }
                                                    _ = v31
                                                }
                                                _ = v29
                                            } else {
                                                var v33 uint32 = 2147483648
                                                v11 = false
                                                if v11 {
                                                    if len(v6) > 0 {
                                                        v6[0] = v33
                                                    }
                                                    v1 = (v1 - v1)
                                                    fmt.Println(v21)
                                                } else {
                                                    var v34 string = ""
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    if (250 != 42) {
                                                        var v35 float32 = ((-2.25 + ((0.0 / (((-2.25 - 0.0) - 1e10) - (3.14159 - -2.25))) + (3.14159 - float32(3.14159)))) / -2.25)
                                                        switch int(v20) {
                                                        case -9223372036854775808:
                                                            var v36 int16 = 1000
                                                            _ = v36
                                                        case 2147483647:
                                                            var v37 uint = 255
                                                            _ = v37
                                                        default:
                                                            v12 = v11
                                                        }
                                                        _ = v35
                                                    } else {
                                                        var v38 int16 = -530
                                                        var v39 int32 = v19
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                    _ = v34
                                                }
                                                fmt.Println(fi26)
                                                _ = v33
                                            }
                                        } else {
                                            var v40 uint8 = 2
                                            var v41 int32 = v19
                                            _ = v40
                                            _ = v41
                                        }
                                    default:
                                        if v11 {
                                            var v42 string = "hello"
                                            v1 -= (v1 | v1)
                                            fmt.Println(v12, v5)
                                            v43 := GenS3{F0: (v19 / ((v19) | 1)), F1: ((43 != 18) || true)}
                                            if true {
                                                v20 += v5
                                                v44 := GenS1{F0: (v42 + ("\n\t" + "foo bar")), F1: 474, F2: 2147483648}
                                                _ = v44
                                            } else {
                                                _ = len(v4)
                                            }
                                            _ = v42
                                            _ = v43
                                        } else {
                                            v5 -= 32767
                                        }
                                    }
                                }
                                switch (fi26 >> ((0) & 63)) {
                                case 2147483647:
                                    if (113 == 42) {
                                        switch (fi27 / (((fi27 ^ 2147483647)) | 1)) {
                                        case 9223372036854775807:
                                            v45 := GenS3{F0: v19, F1: v12}
                                            _ = v45
                                        case -605:
                                            v12 = v12
                                        case -2147483648:
                                            fmt.Println(v11)
                                        }
                                        v46 := GenS3{F0: (^(v19)), F1: v11}
                                        _ = v46
                                    }
                                }
                                v47 := GenS4{F0: 100, F1: -128}
                                _ = fi27
                                _ = v47
                            }
                        case 2147483647:
                            v48 := GenS3{F0: (v19 & v19), F1: v11}
                            _ = v48
                        }
                        var v49 int = (fi26 * fi26)
                        _ = fi26
                        _ = v49
                    }
                    var v50 int = 32768
                    _ = v50
                }
            }
        }
        if (927 <= 65535) {
            fmt.Println(v19, v11, v19)
            v20 -= (^(v20))
        }
        _ = v22
    }
    fmt.Println(v12, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
