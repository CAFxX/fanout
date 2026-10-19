package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = false
    var v3 float32 = float32(1.5)
    var v4 uint = 18446744073709551615
    fmt.Println(v3, v4)
    var v5 string = "\n\t"
    v2 = v2
    switch -758 {
    case -238:
        switch "a" {
        case "hello":
            switch 4294967295 {
            case 549:
                if (!(!(42 >= 397))) {
                    for fi6 := 0; fi6 < 8; fi6++ {
                        if (int8(v4) >= 83) {
                            if v2 {
                                v7 := []int{fi6, 0}
                                var v8 int = fi6
                                for fi9 := 0; fi9 < 2; fi9++ {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    var v10 uint16 = 390
                                    _ = fi9
                                    _ = v10
                                }
                                _ = v7
                                _ = v8
                            }
                            var v11 int = (fi6 & fi6)
                            _ = v11
                        } else {
                            var v12 int = 32768
                            var v13 int8 = 127
                            v14 := []int{(v12 ^ v12), fi6}
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        _ = fi6
                    }
                }
            case 970:
                if (42 >= 0) {
                    for fi15 := 0; fi15 < 6; fi15++ {
                        for fi16 := 0; fi16 < 9; fi16++ {
                            _ = len(v1)
                            _ = fi16
                        }
                        v3 = (1e10 * 1.5)
                        _ = fi15
                    }
                } else {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v5 = "a"
                        var v18 int = -2147483648
                        _ = fi17
                        _ = v18
                    }
                    var v19 string = ((((v5 + v5) + "") + "hello") + (v5 + (v5 + v5)))
                    _ = v19
                }
            case 9223372036854775807:
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 uint8 = 127
                    var v22 int8 = -1
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            default:
                switch 2147483647 {
                case -1:
                    if (false || (v2 || true)) {
                        switch 2147483647 {
                        case -9223372036854775808:
                            v23 := []int{9223372036854775807, -9223372036854775808}
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 12; fi24++ {
                            fmt.Println(v3, v5)
                            _ = fi24
                        }
                        v25 := []int{0, 1}
                        _ = v25
                    }
                case 0:
                    if (v2 || v2) {
                        var v26 int = 32767
                        _ = v26
                    }
                case 32768:
                    v5 = (v5 + "hello")
                default:
                    for fi27 := 0; fi27 < 15; fi27++ {
                        var v28 int = 42
                        v29 := GenS1{F0: 42, F1: 1, F2: v4}
                        v4 = 673
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
            }
        case "a":
            v1[2147483647] = (("" + "foo bar") + "foo bar")
        case "\n\t":
            var v30 uint8 = 170
            _ = v30
        default:
            switch "\n\t" {
            case "hello":
                var v31 int16 = -32768
                _ = v31
            case "\n\t":
                v4 += (v4 - v4)
            case "a":
                fmt.Println(v5, v2, v5)
            }
        }
    case 127:
        v5 = "hello"
    }
    v32 := GenS2{F0: 42, F1: 181}
    v2 = (!true)
    v33 := GenS2{F0: 252, F1: 0}
    var v34 float64 = 3.14159
    if (!(65535 >= 2147483648)) {
        var v35 int32 = -616
        if v2 {
            var v36 int16 = -32768
            for fi37 := 0; fi37 < 4; fi37++ {
                v34 = (v34 * -2.25)
                _ = fi37
            }
            _ = v36
        }
        _ = v35
    }
    v5 = (v5 + "\n\t")
    v34 = v34
    var v38 uint = (v4 / (((v4 ^ 65535)) | 1))
    v39 := []int{1, 9223372036854775807, 4294967295}
    for fi40 := 0; fi40 < 6; fi40++ {
        v39 = append(v39, fi40)
        v34 = ((v34 * ((v34 + float64(1.5)) * (v34 / 0.1))) * (1.5 - v34))
        _ = fi40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
