package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint; F1 int; F2 uint }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 int64; F1 int16 }
func genFn0(p0 int16, p1 uint8) uint8 {
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = v5
    return 255
}
func genFn1(p0 int8, p1 uint16, p2 uint64) uint {
    switch 0 {
    case 756:
        p2 += (^(p2))
    case 673:
        v6 := []int{-268, 4294967295, 996, 9223372036854775807}
        _ = v6
    default:
        var v7 int16 = 0
        _ = v7
    }
    var v8 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 4294967295
}

func main() {
    v1 := []string{"", "", "\n\t", "hello"}
    var v2 uint16 = 2
    v3 := make(map[string]uint32)
    var v4 int8 = 127
    if false {
        var v9 float64 = 0.1
        var v10 bool = true
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v2 -= (v2 / ((896) | 1))
        if true {
            if ((!true) || ("" < "")) {
                v12 := GenS1{F0: 42, F1: "foo bar"}
                _ = v12
            }
            if ((fi11 + (fi11 % ((int(v4)) | 1))) != (-(fi11))) {
                var v13 uint = genFn1((-(v4)), (v2 - 65535), 42)
                for fi14 := 0; fi14 < 1; fi14++ {
                    v15 := GenS2{F0: genFn0(42, 172), F1: "foo bar"}
                    v1 = append(v1, ("hello" + "foo bar"))
                    v13 = (v13 << (((v13 << (((v13 & (v13 << ((genFn1(v4, 65535, 636)) & 63)))) & 63))) & 63))
                    v16 := GenS3{F0: genFn1(v4, 2, uint64(v13)), F1: 127, F2: v13}
                    if (0.1 == 1e10) {
                        v1 = append(v1, ("hello" + ("" + ("\n\t" + ""))))
                        var v17 uint8 = 136
                        _ = v17
                    } else {
                        var v18 string = "\n\t"
                        _ = v18
                    }
                    var v19 uint16 = v2
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v19
                }
                _ = v13
            }
        } else {
            v20 := GenS3{F0: genFn1(v4, v2, 18446744073709551615), F1: fi11, F2: 0}
            v2 = v2
            _ = v20
        }
        _ = fi11
    }
    v2 = (^(v2))
    fmt.Println(v4)
    for fi21 := 0; fi21 < 15; fi21++ {
        v2 = (^(v2))
        switch "hello" {
        case "hello":
            var v22 int64 = 1
            _ = v22
        case "a":
            v23 := GenS4{F0: (v4 << ((0) & 63)), F1: 2147483647}
            _ = v23
        }
        if false {
            _ = len(v1)
        } else {
            v24 := GenS4{F0: (v4 & (v4 << ((1) & 63))), F1: 0}
            v2 = (^(v2))
            fmt.Println(v4)
            _ = v24
        }
        _ = fi21
    }
    var v25 float32 = 1.5
    v2 = 405
    switch "foo bar" {
    case "a":
        var v26 uint32 = 65535
        _ = v26
    case "foo bar":
        v27 := GenS1{F0: 255, F1: "\n\t"}
        _ = v27
    case "hello":
        for fi28 := 0; fi28 < 3; fi28++ {
            if false {
                if false {
                    v2 = v2
                    v2 = 307
                } else {
                    switch ("a" + "hello") {
                    case "\n\t":
                        var v29 int = fi28
                        _ = v29
                    case "hello":
                        _ = len(v1)
                    case "":
                        fmt.Println(v25, fi28)
                    default:
                        if (v2 <= v2) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch fi28 {
                            case 996:
                                for fi30 := 0; fi30 < 13; fi30++ {
                                    v4 = v4
                                    var v31 string = "foo bar"
                                    _ = fi30
                                    _ = v31
                                }
                            case 1:
                                v32 := GenS5{F0: -667, F1: 0}
                                _ = v32
                            }
                        }
                    }
                    for fi33 := 0; fi33 < 10; fi33++ {
                        var v34 bool = (!((!(true || false)) && (127 >= 223)))
                        _ = len(v1)
                        v4 -= (v4 % (((v4 >> ((23) & 63))) | 1))
                        _ = fi33
                        _ = v34
                    }
                }
                switch "" {
                case "hello":
                    fmt.Println(v2, fi28)
                case "foo bar":
                    var v35 string = "a"
                    _ = v35
                }
            } else {
                for fi36 := 0; fi36 < 10; fi36++ {
                    v4 += (v4 + -121)
                    v25 += (1e10 / v25)
                    _ = fi36
                }
            }
            for fi37 := 0; fi37 < 16; fi37++ {
                var v38 uint64 = 848
                var v39 int32 = 238
                v40 := GenS5{F0: -9223372036854775808, F1: -211}
                v4 = 100
                var v41 bool = (2 > 18446744073709551615)
                v42 := GenS1{F0: genFn0(-32768, 42), F1: ("hello" + "a")}
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
                _ = v41
                _ = v42
            }
            var v43 int16 = 884
            _ = fi28
            _ = v43
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v44 float64 = (-2.25 + float64(1.5))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v44 -= (1e10 / v44)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v44
}
