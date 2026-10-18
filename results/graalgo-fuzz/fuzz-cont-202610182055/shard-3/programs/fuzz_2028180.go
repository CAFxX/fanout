package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int8 }
type GenS2 struct { F0 uint16; F1 bool; F2 int }
type GenS3 struct { F0 int16; F1 int16; F2 string }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint8 }
func genFn0() int {
    var v5 string = ("" + ("\n\t" + ""))
    v6 := GenS1{F0: (!(0 != -1)), F1: 42, F2: 0}
    for fi7 := 0; fi7 < 6; fi7++ {
        switch (v5 + v5) {
        case "a":
            v8 := GenS1{F0: true, F1: 2147483647, F2: 42}
            _ = v8
        case "foo bar":
            v5 = v5
        case "\n\t":
            switch (((v5 + v5) + v5) + v5) {
            case "\n\t":
                var v9 uint16 = 255
                _ = v9
            case "":
                v5 = (v5 + (v5 + (v5 + v5)))
            }
        }
        _ = fi7
    }
    v10 := GenS1{F0: true, F1: -29, F2: 1}
    _ = v10
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 uint32, p1 uint64, p2 uint) int8 {
    p2 = (p2 / ((p2) | 1))
    var v11 bool = ((0.1 + float32(0.0)) >= (3.14159 + (float32((float64((0.1 / 3.14159)) / 1e10)) / 1.0)))
    v12 := GenS2{F0: 496, F1: (0 < 255), F2: int(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -5
}

func main() {
    var v1 string = ("hello" + "\n\t")
    v2 := []int{2147483647, -858}
    var v3 float64 = -2.25
    var v4 int8 = 1
    if (true && (!false)) {
        v13 := GenS1{F0: (!(false || true)), F1: -2147483648, F2: (v4 / ((genFn1(42, 42, 18446744073709551615)) | 1))}
        v3 -= v3
        _ = v13
    }
    v3 = ((v3 + (v3 + float64(0.1))) - 1.5)
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 bool = (false && (true || (!true)))
        _ = fi14
        _ = v15
    }
    switch 759 {
    case -633:
        v16 := GenS3{F0: 42, F1: 0, F2: "a"}
        _ = v16
    case 2147483647:
        v17 := GenS2{F0: 1, F1: false, F2: 4294967295}
        _ = v17
    }
    _ = len(v2)
    var v18 string = ""
    switch -541 {
    case -625:
        v18 = v1
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        if true {
            v20 := GenS1{F0: (362 != 1), F1: 1, F2: genFn1(323, 18446744073709551615, 0)}
            v3 = v3
            var v21 uint64 = 4294967295
            _ = v20
            _ = v21
        }
        fmt.Println(v1, v3, v18)
        _ = fi19
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if ((!(245 >= 128)) || true) {
            fmt.Println(v1, v4)
            v1 = v1
            if len(v2) > 0 {
                v2[0] = 127
            }
        } else {
            switch (^(fi22)) {
            case 378:
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 int16 = -1
                    switch fi23 {
                    case 9223372036854775807:
                        var v25 uint8 = 1
                        _ = v25
                    case 29:
                        if (v3 != (v3 - (v3 / v3))) {
                            v1 = ""
                            v26 := GenS4{F0: -2147483648, F1: (^(v4))}
                            fmt.Println(v3)
                            _ = v26
                        } else {
                            fmt.Println(v4, v1, fi22)
                            v2 = append(v2, (fi23 & (fi23 / (((fi22 | fi22)) | 1))))
                            switch ("\n\t" + v18) {
                            case "":
                                var v27 uint32 = 89
                                _ = v27
                            }
                        }
                    }
                    var v28 uint8 = 169
                    _ = fi23
                    _ = v24
                    _ = v28
                }
            default:
                for fi29 := 0; fi29 < 8; fi29++ {
                    var v30 float32 = float32(v3)
                    _ = fi29
                    _ = v30
                }
            }
            var v31 uint64 = 0
            for fi32 := 0; fi32 < 2; fi32++ {
                v18 = ("a" + "hello")
                switch "\n\t" {
                case "a":
                    var v33 bool = (!true)
                    switch v1 {
                    case "":
                        switch (fi32 << ((987) & 63)) {
                        case 32767:
                            v33 = false
                        case 102:
                            if len(v2) > 0 {
                                v2[0] = 32767
                            }
                        case 4294967295:
                            switch "a" {
                            case "\n\t":
                                v34 := GenS5{F0: 1}
                                _ = v34
                            case "":
                                v35 := GenS4{F0: 675, F1: (v4 % ((v4) | 1))}
                                _ = v35
                            }
                        default:
                            v2 = append(v2, fi22)
                        }
                    default:
                        v4 -= genFn1(uint32(v31), (v31 % ((v31) | 1)), 267)
                    }
                    fmt.Println(v31, fi22, v4)
                    v2 = append(v2, (^(fi22)))
                    _ = v33
                default:
                    if (false && true) {
                        var v36 int8 = (v4 >> ((190) & 63))
                        switch fi32 {
                        case 127:
                            var v37 uint = 42
                            v18 = "foo bar"
                            _ = v37
                        case 9223372036854775807:
                            v4 -= genFn1(1, 320, 255)
                        case 0:
                            var v38 int = fi32
                            _ = v38
                        default:
                            if true {
                                var v39 float64 = 0.1
                                _ = v39
                            } else {
                                v18 = "hello"
                                v4 = v36
                                v18 = ("foo bar" + v1)
                            }
                        }
                        _ = v36
                    }
                    var v40 string = "hello"
                    _ = v40
                }
                var v41 int64 = 9223372036854775807
                _ = fi32
                _ = v41
            }
            _ = v31
        }
        v42 := GenS5{F0: 150}
        if len(v2) > 0 {
            v2[0] = fi22
        }
        _ = fi22
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
