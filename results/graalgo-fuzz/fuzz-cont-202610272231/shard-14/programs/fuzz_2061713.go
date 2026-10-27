package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 bool; F2 bool }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int32 }
func genFn0(p0 int) int64 {
    v6 := GenS1{F0: 0}
    p0 += 0
    v7 := GenS2{F0: 2147483647, F1: (p0 / (((p0 | 2147483647)) | 1))}
    _ = p0
    _ = v6
    _ = v7
    return -428
}
func genFn1(p0 string) int32 {
    v8 := GenS3{F0: 13}
    var v9 int = 833
    _ = p0
    _ = v8
    _ = v9
    return -528
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint8 = 128
    var v3 uint64 = 716
    var v4 uint = 662
    v5 := []int{4294967295, 9223372036854775807, -73, -9223372036854775808}
    v5 = append(v5, 128)
    v10 := GenS2{F0: 1, F1: -584}
    switch "a" {
    case "a":
        v1 = (v1 / ((2147483647) | 1))
    case "\n\t":
        v2 = (v2 ^ 127)
    case "hello":
        v11 := GenS4{F0: (^(v3)), F1: (2147483647 == -1), F2: false}
        _ = v11
    }
    var v12 bool = false
    v3 = (v3 + (v3 % ((v3) | 1)))
    var v13 uint8 = v2
    v14 := GenS4{F0: v3, F1: v12, F2: v12}
    switch -9223372036854775808 {
    case 32767:
        var v15 int32 = -2147483648
        _ = v15
    case 473:
        var v16 bool = v12
        _ = v16
    case 9223372036854775807:
        switch -9223372036854775808 {
        case 1:
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS4{F0: v3, F1: v12, F2: v12}
                _ = fi17
                _ = v18
            }
        case 58:
            if ((((0 > 0) && v12) && (true || v12)) || true) {
                switch 4294967295 {
                case -341:
                    for fi19 := 0; fi19 < 7; fi19++ {
                        for fi20 := 0; fi20 < 19; fi20++ {
                            if v12 {
                                var v21 uint = v4
                                if v12 {
                                    var v22 uint16 = 267
                                    v2 = 127
                                    _ = v22
                                }
                                var v23 bool = false
                                _ = v21
                                _ = v23
                            }
                            var v24 float64 = 3.14159
                            var v25 int32 = 2147483647
                            _ = fi20
                            _ = v24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 12; fi26++ {
                            var v27 int16 = 1000
                            _ = fi26
                            _ = v27
                        }
                        v4 = v4
                        _ = fi19
                    }
                default:
                    switch "a" {
                    case "a":
                        v4 = v4
                    case "foo bar":
                        var v28 uint16 = 820
                        _ = v28
                    case "":
                        v5 = append(v5, 42)
                    default:
                        for fi29 := 0; fi29 < 6; fi29++ {
                            for fi30 := 0; fi30 < 6; fi30++ {
                                fmt.Println(v4, v13)
                                _ = fi30
                            }
                            switch "" {
                            case "":
                                if v12 {
                                    v2 = uint8(v1)
                                    fmt.Println(v12, fi29, v1)
                                    if len(v5) > 0 {
                                        v5[0] = 128
                                    }
                                }
                            default:
                                var v31 int16 = 42
                                _ = v31
                            }
                            if (!(("foo bar" + ("foo bar" + (("foo bar" + "a") + "foo bar"))) > "")) {
                                v13 += 1
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            } else {
                                v5 = append(v5, (fi29 | fi29))
                            }
                            _ = fi29
                        }
                    }
                }
                if (!(3.14159 > 1e10)) {
                    v32 := GenS4{F0: v3, F1: (23 != (v1 >> ((v4) & 63))), F2: (!v12)}
                    v33 := GenS5{F0: 1, F1: 127}
                    _ = v32
                    _ = v33
                } else {
                    v34 := GenS3{F0: 45}
                    _ = v34
                }
            }
        case -2147483648:
            if true {
                var v35 uint32 = v1
                v36 := GenS6{F0: -39}
                _ = v35
                _ = v36
            }
        default:
            if false {
                for fi37 := 0; fi37 < 1; fi37++ {
                    switch "hello" {
                    case "a":
                        var v38 float32 = 1e10
                        _ = v38
                    case "\n\t":
                        var v39 uint = v4
                        _ = v39
                    }
                    _ = fi37
                }
                var v40 int32 = genFn1("hello")
                _ = v40
            } else {
                fmt.Println(v2)
            }
        }
    default:
        v41 := GenS1{F0: uint32(v4)}
        _ = v41
    }
    fmt.Println(v12)
    var v42 string = "foo bar"
    var v43 string = ("\n\t" + "\n\t")
    v43 = (v42 + ((v43 + v42) + (("" + ((v43 + "a") + "")) + v42)))
    var v44 float64 = 3.14159
    var v45 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
