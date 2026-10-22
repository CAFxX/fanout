package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 int64; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 string; F2 bool }
type GenS6 struct { F0 bool; F1 uint; F2 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 uint8) string {
    p0 -= (p0 ^ p0)
    var v5 int8 = -113
    v6 := GenS1{F0: -1, F1: (p0 * 42)}
    _ = p0
    _ = v5
    _ = v6
    return "a"
}
func genFn1(p0 uint, p1 uint16, p2 string) int32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 float64 = 0.1
        v9 := []int{fi7, -2147483648, (fi7 - (fi7 / (((fi7 >> ((2) & 63))) | 1)))}
        var v10 float64 = 0.1
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if true {
        var v11 uint8 = 244
        _ = v11
    } else {
        p2 = "hello"
        for fi12 := 0; fi12 < 12; fi12++ {
            if ("foo bar" == p2) {
                p1 += (p1 % ((p1) | 1))
            }
            _ = fi12
        }
    }
    var v13 uint8 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -513
}

func main() {
    var v1 string = "hello"
    v2 := []uint32{2147483647}
    var v3 uint32 = 42
    v4 := []string{v1}
    switch "foo bar" {
    case "a":
        switch 70 {
        case 249:
            switch -1 {
            case 9223372036854775807:
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
            case 1:
                for fi14 := 0; fi14 < 15; fi14++ {
                    v1 = ""
                    if (true || false) {
                        var v15 float32 = 3.14159
                        v1 = ("\n\t" + v1)
                        v16 := GenS1{F0: genFn1(521, 2, (v1 + v1)), F1: 0}
                        _ = v15
                        _ = v16
                    } else {
                        v1 = v1
                    }
                    for fi17 := 0; fi17 < 2; fi17++ {
                        var v18 int8 = 42
                        if false {
                            if false {
                                switch ("\n\t" + v1) {
                                case "foo bar":
                                    if true {
                                        v1 = v1
                                    }
                                case "\n\t":
                                    v3 = (v3 | 0)
                                }
                            } else {
                                v4 = append(v4, (v1 + v1))
                                v18 -= (v18 % ((int8(fi17)) | 1))
                                var v19 int32 = 2147483647
                                _ = v19
                            }
                            switch 32767 {
                            case 764:
                                v3 = 774
                            case 127:
                                v20 := GenS2{F0: 224, F1: (v18 >> ((18446744073709551615) & 63))}
                                _ = v20
                            }
                            v21 := GenS2{F0: 65535, F1: v18}
                            _ = v21
                        } else {
                            fmt.Println(v18)
                            var v22 int16 = -1
                            var v23 float64 = 0.0
                            _ = v22
                            _ = v23
                        }
                        var v24 int8 = (v18 - (v18 / ((-128) | 1)))
                        _ = fi17
                        _ = v18
                        _ = v24
                    }
                    _ = fi14
                }
            case -665:
                var v25 uint8 = 255
                _ = v25
            }
        }
    case "foo bar":
        v26 := GenS3{F0: v1, F1: -1, F2: -812}
        _ = v26
    default:
        v27 := GenS2{F0: 255, F1: -45}
        _ = v27
    }
    var v28 int16 = 42
    var v29 float32 = 0.0
    if true {
        v28 = (-(v28))
        if (!true) {
            v28 += -32768
            var v30 int16 = (v28 << ((4294967295) & 63))
            _ = v30
        }
    }
    if false {
        switch ("\n\t" + ("hello" + v1)) {
        case "\n\t":
            if (133 <= 492) {
                v31 := GenS4{F0: 255}
                var v32 uint32 = (v3 / (((v3 / ((1) | 1))) | 1))
                _ = v31
                _ = v32
            }
        case "foo bar":
            fmt.Println(v28)
        case "a":
            switch v1 {
            case "":
                v33 := GenS5{F0: (!(v1 >= ("hello" + v1))), F1: v1, F2: (!true)}
                _ = v33
            case "hello":
                v3 = 0
            case "foo bar":
                if true {
                    var v34 int8 = 100
                    switch int(v34) {
                    case 767:
                        v35 := GenS3{F0: "foo bar", F1: -1, F2: (v28 & 42)}
                        _ = v35
                    case 637:
                        v34 += v34
                    case -853:
                        fmt.Println(v1)
                    default:
                        if (224 > 20) {
                            switch "foo bar" {
                            case "foo bar":
                                switch v1 {
                                case "":
                                    fmt.Println(v28, v28, v1)
                                case "\n\t":
                                    v36 := GenS6{F0: false, F1: 1, F2: 426}
                                    _ = v36
                                case "a":
                                    var v37 uint64 = 83
                                    _ = v37
                                }
                            case "\n\t":
                                var v38 uint8 = 128
                                _ = v38
                            case "":
                                v39 := GenS5{F0: (((2 > 384) && (v34 == (v34 | v34))) && false), F1: v1, F2: (true && true)}
                                _ = v39
                            default:
                                for fi40 := 0; fi40 < 19; fi40++ {
                                    v3 += v3
                                    var v41 int16 = (v28 * v28)
                                    _ = fi40
                                    _ = v41
                                }
                            }
                            v42 := GenS7{F0: 42}
                            _ = v42
                        }
                    }
                    _ = len(v2)
                    _ = v34
                }
            }
        }
    } else {
        var v43 int32 = genFn1(1, 255, v1)
        var v44 int32 = v43
        _ = v43
        _ = v44
    }
    if true {
        if true {
            v3 += (v3 * v3)
            v29 -= (v29 * (0.1 - (0.0 * v29)))
            var v45 float32 = float32(0.1)
            _ = v45
        }
    } else {
        var v46 float32 = 1e10
        _ = v46
    }
    var v47 int16 = v28
    if false {
        v47 += 0
        switch v1 {
        case "a":
            v4 = append(v4, v1)
        case "":
            if (!(((v1 + ("\n\t" + "\n\t")) + v1) != v1)) {
                if (!(!true)) {
                    if (42 <= 0) {
                        v48 := GenS8{F0: 1}
                        var v49 float64 = 0.1
                        _ = v48
                        _ = v49
                    } else {
                        v28 = v28
                    }
                    fmt.Println(v28, v47)
                }
                v29 = float32((((1e10 * 1.5) - 1.5) * 0.0))
                var v50 int = 281
                _ = v50
            } else {
                var v51 float32 = v29
                _ = v51
            }
        }
    } else {
        _ = len(v4)
    }
    v29 += v29
    for fi52 := 0; fi52 < 14; fi52++ {
        switch fi52 {
        case 0:
            v53 := GenS5{F0: (32767 >= 69), F1: "a", F2: false}
            _ = v53
        case -9223372036854775808:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -92:
            v3 = 2147483647
        }
        if (true || false) {
            var v54 int = (fi52 - fi52)
            fmt.Println(v28, v3, fi52)
            v1 = (((v1 + ((v1 + ("foo bar" + (("hello" + v1) + "\n\t"))) + v1)) + (v1 + v1)) + "\n\t")
            _ = v54
        }
        var v55 uint8 = 2
        _ = fi52
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v47
}
