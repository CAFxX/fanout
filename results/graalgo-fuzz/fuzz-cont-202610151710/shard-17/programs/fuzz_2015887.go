package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int16; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int }
func genFn0(p0 int64, p1 int, p2 string) uint16 {
    v7 := GenS1{F0: p1, F1: 18446744073709551615}
    var v8 int8 = -76
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1() int8 {
    var v9 string = ("a" + "hello")
    if true {
        fmt.Println(v9)
        v9 = v9
    } else {
        v10 := GenS2{F0: -1, F1: -32768, F2: 127}
        v11 := []int{499}
        _ = v10
        _ = v11
    }
    var v12 float32 = (float32(3.14159) * ((3.14159 / float32((float64(3.14159) + ((((0.1 * (float64(((0.1 - float32(1.5)) + (1.5 / 1.0))) - 3.14159)) / 1.0) + ((1e10 + -2.25) - 3.14159)) - 0.0)))) - 1.5))
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            v12 = (1.5 + v12)
            var v15 int32 = 42
            _ = fi14
            _ = v15
        }
        v16 := []int{-2147483648, fi13, (fi13 >> ((1) & 63)), (fi13 << ((0) & 63))}
        _ = fi13
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int32 = 0
        var v19 float32 = v12
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = v12
    _ = v9
    return 54
}

func main() {
    var v1 uint32 = 1
    var v2 uint = 18446744073709551615
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    var v5 uint64 = 773
    v6 := []string{(("a" + "") + "foo bar"), "foo bar"}
    if false {
        if len(v6) > 0 {
            v6[0] = "foo bar"
        }
        fmt.Println(v1, v5, v5)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        if false {
            if false {
                fmt.Println(v2, v2, v1)
                v5 -= 250
            } else {
                if false {
                    v2 = 656
                } else {
                    v20 := GenS1{F0: 42, F1: (^(v5))}
                    _ = v20
                }
                var v21 int16 = 813
                v22 := GenS1{F0: 32768, F1: (v5 % ((4294967295) | 1))}
                _ = v21
                _ = v22
            }
            var v23 string = "a"
            _ = v23
        } else {
            switch "foo bar" {
            case "":
                switch ("a" + "") {
                case "\n\t":
                    if (false || (-9223372036854775808 > 9223372036854775807)) {
                        v5 = 18446744073709551615
                    }
                case "a":
                    switch "" {
                    case "":
                        for fi24 := 0; fi24 < 12; fi24++ {
                            var v25 int32 = -1
                            switch (fi24 ^ (fi24 + int(v5))) {
                            case 0:
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    delete(v4, ("" + ""))
                                    _ = fi26
                                }
                            }
                            switch fi24 {
                            case 264:
                                var v27 int64 = -1
                                _ = v27
                            case 42:
                                v28 := GenS2{F0: -859, F1: 32767, F2: 135}
                                _ = v28
                            case -9:
                                var v29 uint = (v2 & v2)
                                _ = v29
                            default:
                                v25 = int32(v2)
                            }
                            _ = fi24
                            _ = v25
                        }
                    }
                case "foo bar":
                    delete(v3, "a")
                default:
                    var v30 uint8 = 128
                    _ = v30
                }
            case "foo bar":
                if true {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        v5 -= v5
                        var v32 uint32 = 2
                        _ = fi31
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 10; fi33++ {
                        if (-2.25 >= 1e10) {
                            v1 = (v1 / ((v1) | 1))
                            if (v1 >= 700) {
                                var v34 uint64 = 2
                                fmt.Println(v34, v34, v5)
                                _ = v34
                            }
                        }
                        var v35 uint = (v2 << (((v2 << ((42) & 63))) & 63))
                        v36 := GenS3{F0: 212}
                        _ = fi33
                        _ = v35
                        _ = v36
                    }
                    v37 := GenS3{F0: 1}
                    _ = v37
                } else {
                    v38 := GenS2{F0: 32767, F1: -32768, F2: 128}
                    var v39 string = ("\n\t" + (("a" + (("foo bar" + (("hello" + "foo bar") + ("\n\t" + ("foo bar" + "hello")))) + "foo bar")) + ""))
                    fmt.Println(v5)
                    _ = v38
                    _ = v39
                }
            default:
                fmt.Println(v5)
            }
            for fi40 := 0; fi40 < 8; fi40++ {
                for fi41 := 0; fi41 < 8; fi41++ {
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v1, fi40, v5)
                    case "\n\t":
                        var v42 float64 = 0.1
                        _ = v42
                    case "a":
                        v43 := GenS1{F0: fi41, F1: v5}
                        _ = v43
                    default:
                        v1 -= (^(v1))
                    }
                    _ = fi41
                }
                _ = fi40
            }
        }
        v5 -= (v5 / ((v5) | 1))
        v5 = uint64(v1)
    }
    v44 := GenS4{F0: 1, F1: 32767}
    var v45 float64 = (1.5 + 0.0)
    var v46 int32 = -1
    var v47 int64 = -9223372036854775808
    for fi48 := 0; fi48 < 9; fi48++ {
        var v49 int = fi48
        _ = fi48
        _ = v49
    }
    var v50 uint64 = 0
    for fi51 := 0; fi51 < 4; fi51++ {
        v2 = (v2 / (((v2 * v2)) | 1))
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v50
    _ = v6
}
