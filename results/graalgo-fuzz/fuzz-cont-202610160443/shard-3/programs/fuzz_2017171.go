package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int8; F2 uint }
type GenS3 struct { F0 uint64; F1 int; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 string; F2 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 int8 }

func main() {
    var v1 int16 = 0
    var v2 uint8 = 2
    v3 := make(map[int]string)
    var v4 int16 = 42
    v5 := GenS1{F0: 0}
    v6 := GenS2{F0: true, F1: -1, F2: 18446744073709551615}
    var v7 bool = (128 <= v2)
    switch "\n\t" {
    case "hello":
        var v8 string = ("a" + "a")
        _ = v8
    default:
        for fi9 := 0; fi9 < 4; fi9++ {
            _ = len(v3)
            if v7 {
                var v10 uint8 = 100
                for fi11 := 0; fi11 < 10; fi11++ {
                    if v7 {
                        v12 := GenS3{F0: 18446744073709551615, F1: (fi11 % ((817) | 1)), F2: -1}
                        switch ("foo bar" + "foo bar") {
                        case "a":
                            v13 := GenS1{F0: 634}
                            _ = v13
                        }
                        _ = v12
                    } else {
                        var v14 int64 = 1
                        v15 := []int{}
                        if len(v15) > 0 {
                            v15[0] = fi9
                        }
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS4{F0: v7}
                    v10 -= (v10 * v10)
                    _ = fi11
                    _ = v16
                }
                v7 = v7
                _ = v10
            } else {
                v17 := []int{fi9}
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    v18 := GenS2{F0: v7, F1: 100, F2: 2}
                    _ = v18
                case "":
                    var v19 string = "foo bar"
                    _ = v19
                case "hello":
                    v20 := GenS5{F0: (((("foo bar" + "foo bar") + ("hello" + "a")) + "hello") + ""), F1: "foo bar", F2: (v1 / ((v1) | 1))}
                    _ = v20
                default:
                    var v21 uint32 = 4294967295
                    _ = v21
                }
                v1 = (v1 % (((v1 | v4)) | 1))
                _ = v17
            }
            _ = fi9
        }
    }
    v3[128] = ("foo bar" + ("a" + "hello"))
    var v22 int16 = (v4 & (v1 + 0))
    var v23 int32 = 1
    v24 := []int{128, -545, 127}
    var v25 uint = 2
    delete(v3, 4294967295)
    if (((922 >= 65535) && (!v7)) && v7) {
        v22 = v1
        for fi26 := 0; fi26 < 12; fi26++ {
            for fi27 := 0; fi27 < 1; fi27++ {
                _ = len(v3)
                var v28 bool = true
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    } else {
        var v29 float32 = (-2.25 + 1.5)
        if (32767 > 91) {
            for fi30 := 0; fi30 < 20; fi30++ {
                switch ("a" + "") {
                case "":
                    var v31 bool = v7
                    _ = v31
                case "\n\t":
                    fmt.Println(v25)
                default:
                    var v32 uint = (v25 / ((v25) | 1))
                    _ = v32
                }
                _ = len(v3)
                var v33 int = fi30
                _ = fi30
                _ = v33
            }
        }
        _ = v29
    }
    v34 := GenS6{F0: 255}
    var v35 bool = (v2 < (v2 >> ((v25) & 63)))
    for fi36 := 0; fi36 < 16; fi36++ {
        if (!("" <= (("foo bar" + ("\n\t" + "hello")) + ((("hello" + "") + "") + "foo bar")))) {
            if v35 {
                v25 += (v25 * v25)
                var v37 string = ("a" + "\n\t")
                _ = v37
            } else {
                v38 := GenS6{F0: 1}
                var v39 uint32 = 65535
                v2 += 1
                _ = v38
                _ = v39
            }
        } else {
            for fi40 := 0; fi40 < 13; fi40++ {
                var v41 int16 = v1
                _ = fi40
                _ = v41
            }
        }
        _ = fi36
    }
    var v42 uint8 = v2
    for fi43 := 0; fi43 < 4; fi43++ {
        if true {
            if (347 != 2147483647) {
                if v35 {
                    if true {
                        var v44 float64 = 3.14159
                        v45 := GenS1{F0: 588}
                        if (2 > 357) {
                            v46 := GenS7{F0: -67, F1: -128}
                            _ = v46
                        }
                        _ = v44
                        _ = v45
                    } else {
                        var v47 int8 = 100
                        var v48 float64 = 1e10
                        _ = v47
                        _ = v48
                    }
                    if (0.1 != -2.25) {
                        var v49 int8 = -68
                        _ = v49
                    }
                    if false {
                        var v50 float64 = (-2.25 + 1e10)
                        _ = v50
                    } else {
                        switch ("" + "") {
                        case "":
                            for fi51 := 0; fi51 < 13; fi51++ {
                                fmt.Println(v4)
                                _ = fi51
                            }
                        case "hello":
                            v52 := GenS6{F0: 570}
                            _ = v52
                        case "\n\t":
                            var v53 uint64 = 2
                            _ = v53
                        }
                        for fi54 := 0; fi54 < 12; fi54++ {
                            var v55 int32 = (-(v23))
                            _ = fi54
                            _ = v55
                        }
                        for fi56 := 0; fi56 < 14; fi56++ {
                            switch "hello" {
                            case "a":
                                var v57 bool = (1 > 1)
                                _ = v57
                            case "hello":
                                var v58 uint8 = v42
                                _ = v58
                            }
                            _ = fi56
                        }
                    }
                }
                var v59 int = 128
                v1 = v22
                _ = v59
            } else {
                switch "foo bar" {
                case "hello":
                    _ = len(v24)
                case "foo bar":
                    for fi60 := 0; fi60 < 20; fi60++ {
                        fmt.Println(v42, v42)
                        var v61 int16 = (v1 + v1)
                        v62 := v3[(fi43 ^ (fi60 - fi60))]
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                default:
                    if (!false) {
                        if true {
                            var v63 bool = v7
                            _ = v63
                        }
                        for fi64 := 0; fi64 < 11; fi64++ {
                            v22 = v22
                            _ = fi64
                        }
                    }
                }
            }
            var v65 uint16 = 255
            _ = v65
        }
        _ = fi43
    }
    switch "foo bar" {
    case "foo bar":
        v66 := GenS2{F0: v7, F1: int8(v23), F2: v25}
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
