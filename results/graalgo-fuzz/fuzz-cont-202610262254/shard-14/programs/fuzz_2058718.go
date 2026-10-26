package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int64 }

func main() {
    var v1 uint8 = 127
    var v2 uint32 = 2
    v3 := make(map[int]string)
    var v4 float32 = 3.14159
    v5 := []int{32768, 127, 254}
    var v6 int64 = 42
    v7 := GenS1{F0: 608, F1: -2147483648, F2: 1}
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := v3[fi8]
            fmt.Println(v2, v4, v6)
            _ = fi9
            _ = v10
        }
        v2 -= 4294967295
        v4 -= -2.25
        _ = fi8
    }
    v11 := GenS1{F0: 2, F1: 42, F2: 32767}
    for fi12 := 0; fi12 < 3; fi12++ {
        v4 = float32(float64(v4))
        _ = fi12
    }
    v13 := GenS2{F0: -46}
    if ((!true) || (true || true)) {
        v14 := GenS2{F0: 100}
        _ = v14
    }
    fmt.Println(v4)
    switch "foo bar" {
    case "\n\t":
        v15 := GenS2{F0: 1}
        _ = v15
    case "hello":
        var v16 int32 = 260
        _ = v16
    case "foo bar":
        switch -1 {
        case 42:
            for fi17 := 0; fi17 < 14; fi17++ {
                for fi18 := 0; fi18 < 8; fi18++ {
                    v4 += v4
                    var v19 float64 = 3.14159
                    _ = fi18
                    _ = v19
                }
                var v20 bool = false
                for fi21 := 0; fi21 < 5; fi21++ {
                    delete(v3, fi21)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi21
                }
                _ = fi17
                _ = v20
            }
        case -103:
            for fi22 := 0; fi22 < 3; fi22++ {
                v1 += v1
                if (42 > 617) {
                    v4 += 1e10
                } else {
                    for fi23 := 0; fi23 < 2; fi23++ {
                        v3[fi22] = ""
                        _ = fi23
                    }
                    v3[32767] = ("" + ("foo bar" + ("" + "foo bar")))
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi22
            }
        case -519:
            if len(v5) > 0 {
                v5[0] = 630
            }
        }
    default:
        if true {
            v1 = 2
            fmt.Println(v1, v4)
            var v24 uint = 42
            _ = v24
        } else {
            v3[-1] = "foo bar"
            var v25 float32 = v4
            if (!true) {
                if false {
                    v2 += v2
                } else {
                    var v26 int = 32768
                    for fi27 := 0; fi27 < 18; fi27++ {
                        if (!false) {
                            switch "hello" {
                            case "":
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(fi27, v2, v1)
                                }
                            case "\n\t":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case "a":
                                if (true || (!false)) {
                                    var v28 uint64 = 4294967295
                                    switch -2147483648 {
                                    case 610:
                                        if false {
                                            switch "hello" {
                                            case "\n\t":
                                                v29 := GenS3{F0: (v26 & (fi27 + fi27)), F1: v6}
                                                _ = v29
                                            case "foo bar":
                                                if (1 >= 18446744073709551615) {
                                                    switch "hello" {
                                                    case "hello":
                                                        for fi30 := 0; fi30 < 9; fi30++ {
                                                            if false {
                                                                fmt.Println(v4)
                                                                delete(v3, -421)
                                                                v1 = v1
                                                            } else {
                                                                v1 = (v1 << ((255) & 63))
                                                                v31 := GenS3{F0: v26, F1: v6}
                                                                var v32 string = ""
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                            _ = fi30
                                                        }
                                                    case "foo bar":
                                                        v5 = append(v5, 9223372036854775807)
                                                    case "a":
                                                        for fi33 := 0; fi33 < 8; fi33++ {
                                                            v5 = append(v5, fi33)
                                                            if ((v4 >= v25) || (!true)) {
                                                                v34 := GenS1{F0: 283, F1: -2147483648, F2: 60}
                                                                v1 = v1
                                                                delete(v3, (v26 + -272))
                                                                _ = v34
                                                            } else {
                                                                _ = len(v5)
                                                            }
                                                            _ = fi33
                                                        }
                                                    }
                                                    fmt.Println(v2, v2)
                                                } else {
                                                    _ = len(v3)
                                                    var v35 int64 = v6
                                                    v36 := GenS3{F0: (fi27 | v26), F1: v6}
                                                    _ = v35
                                                    _ = v36
                                                }
                                            }
                                            if true {
                                                v1 += v1
                                                fmt.Println(v6, v26)
                                            }
                                        }
                                    case 127:
                                        v3[(v26 / ((32767) | 1))] = ("a" + "a")
                                    case -592:
                                        var v37 int16 = -1
                                        _ = v37
                                    }
                                    v1 += (v1 * (v1 - (v1 / (((v1 / ((v1) | 1))) | 1))))
                                    _ = v28
                                }
                            }
                        } else {
                            v5 = append(v5, -1)
                        }
                        _ = fi27
                    }
                    v1 -= v1
                    _ = v26
                }
            } else {
                switch ("foo bar" + "hello") {
                case "\n\t":
                    switch "a" {
                    case "hello":
                        var v38 int = 2147483647
                        _ = v38
                    }
                case "foo bar":
                    var v39 bool = (v4 < float32(0.1))
                    _ = v39
                }
                v1 = v1
            }
            _ = v25
        }
    }
    var v40 float32 = 1e10
    var v41 int64 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
