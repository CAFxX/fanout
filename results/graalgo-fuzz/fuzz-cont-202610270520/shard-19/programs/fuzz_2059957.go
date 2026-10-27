package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 uint32; F2 string }
func genFn0(p0 uint) int16 {
    v5 := GenS1{F0: -32768}
    p0 = (p0 ^ 870)
    var v6 uint = p0
    var v7 float32 = 1e10
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 float64 = (0.1 + 1.5)
    var v2 int8 = 100
    var v3 uint16 = 42
    v4 := []int{2147483647, 32768}
    if false {
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                v1 -= (v1 + -2.25)
                _ = fi9
            }
            fmt.Println(v1, v3, v2)
            if len(v4) > 0 {
                v4[0] = fi8
            }
            _ = fi8
        }
        fmt.Println(v3, v2, v3)
    }
    _ = len(v4)
    v10 := GenS1{F0: genFn0(255)}
    switch "a" {
    case "a":
        _ = len(v4)
    case "foo bar":
        var v11 string = ("hello" + "")
        _ = v11
    }
    switch "a" {
    case "foo bar":
        _ = len(v4)
    case "\n\t":
        var v12 int8 = (v2 + (^(v2)))
        _ = v12
    case "hello":
        var v13 uint32 = 653
        _ = v13
    default:
        if false {
            if false {
                var v14 float32 = (1.5 + (3.14159 - 0.1))
                var v15 int16 = 1000
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 int = fi16
                    _ = fi16
                    _ = v17
                }
                _ = v14
                _ = v15
            } else {
                fmt.Println(v3, v3, v3)
                var v18 float64 = 1e10
                _ = len(v4)
                _ = v18
            }
            switch -353 {
            case 32768:
                fmt.Println(v3, v1, v3)
            default:
                var v19 bool = true
                _ = v19
            }
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        if (v3 <= (^(v3))) {
            switch fi20 {
            case 42:
                for fi21 := 0; fi21 < 12; fi21++ {
                    var v22 int8 = v2
                    if len(v4) > 0 {
                        v4[0] = (fi21 ^ (fi21 >> ((323) & 63)))
                    }
                    _ = fi21
                    _ = v22
                }
            case 105:
                _ = len(v4)
            case 2147483647:
                for fi23 := 0; fi23 < 2; fi23++ {
                    switch -9223372036854775808 {
                    case 0:
                        switch (fi23 & fi20) {
                        case 9223372036854775807:
                            v24 := GenS2{F0: ((1 < 1000) || false)}
                            _ = v24
                        case 42:
                            switch fi23 {
                            case -2147483648:
                                switch -1 {
                                case -923:
                                    var v25 int8 = v2
                                    _ = v25
                                case 2147483647:
                                    v1 = v1
                                case 4294967295:
                                    if true {
                                        v26 := GenS2{F0: ((v3 % (((v3 * 2)) | 1)) >= v3)}
                                        v2 = (-(v2))
                                        if (42 >= 1) {
                                            var v27 float64 = v1
                                            var v28 uint32 = 339
                                            v4 = append(v4, 1)
                                            var v29 uint8 = 255
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                        _ = v26
                                    }
                                default:
                                    var v30 float64 = v1
                                    _ = v30
                                }
                            default:
                                var v31 string = "\n\t"
                                _ = len(v4)
                                v32 := GenS3{F0: 0}
                                _ = v31
                                _ = v32
                            }
                        case 4294967295:
                            fmt.Println(v1, fi23)
                        }
                    default:
                        switch (fi23 | -2147483648) {
                        case -39:
                            var v33 int16 = 0
                            _ = v33
                        default:
                            fmt.Println(v1)
                        }
                    }
                    v2 -= -128
                    fmt.Println(fi20, fi20)
                    _ = fi23
                }
            }
            v34 := GenS3{F0: 0}
            var v35 int32 = -2147483648
            _ = v34
            _ = v35
        } else {
            v36 := GenS4{F0: 553}
            fmt.Println(v1, fi20, v1)
            _ = v36
        }
        v37 := GenS2{F0: false}
        _ = fi20
        _ = v37
    }
    v3 += 72
    if (0 >= 327) {
        v2 = v2
        if len(v4) > 0 {
            v4[0] = 32767
        }
        var v38 bool = true
        _ = v38
    }
    fmt.Println(v2)
    v3 = v3
    v39 := GenS3{F0: int32(v3)}
    switch "hello" {
    case "hello":
        v1 = v1
    case "":
        switch -1 {
        case -1:
            _ = len(v4)
        default:
            for fi40 := 0; fi40 < 4; fi40++ {
                var v41 int = fi40
                switch (v41 << ((uint(v3)) & 63)) {
                case 42:
                    v1 = v1
                case 0:
                    if true {
                        v42 := GenS5{F0: (v3 / (((^(v3))) | 1)), F1: 42, F2: "hello"}
                        if (!true) {
                            var v43 int32 = -1
                            _ = v43
                        }
                        _ = v42
                    }
                default:
                    v44 := GenS4{F0: 255}
                    _ = v44
                }
                var v45 uint16 = v3
                _ = fi40
                _ = v41
                _ = v45
            }
        }
    case "\n\t":
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v39
    _ = v4
}
