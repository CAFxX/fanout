package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 string = ("a" + "foo bar")
    var v2 string = ((v1 + v1) + (v1 + (v1 + v1)))
    v3 := make(map[string]string)
    var v4 int32 = 1
    var v5 bool = true
    if (false || (-128 >= 0)) {
        v4 = 460
        var v6 uint = 1
        switch v2 {
        case "a":
            var v7 int16 = -312
            _ = v7
        case "":
            var v8 string = ((v2 + (v1 + v1)) + (v1 + v2))
            _ = v8
        case "\n\t":
            switch 939 {
            case -939:
                var v9 int16 = 1
                _ = v9
            }
        default:
            v6 = v6
        }
        _ = v6
    } else {
        switch ("a" + v2) {
        case "\n\t":
            switch v1 {
            case "foo bar":
                v10 := GenS1{F0: 4294967295, F1: 1}
                _ = v10
            default:
                var v11 uint = 0
                _ = v11
            }
        case "foo bar":
            v12 := GenS1{F0: 2147483648, F1: -621}
            _ = v12
        default:
            v5 = (4294967295 != 1)
        }
        if ((960 <= 18446744073709551615) && (66 == 0)) {
            switch 2147483647 {
            case 1:
                switch 32767 {
                case -2147483648:
                    for fi13 := 0; fi13 < 11; fi13++ {
                        var v14 int32 = v4
                        fmt.Println(v5)
                        switch "hello" {
                        case "\n\t":
                            v4 += v4
                        case "foo bar":
                            for fi15 := 0; fi15 < 5; fi15++ {
                                v16 := []int{(fi15 * (fi15 * fi13)), fi15, (fi15 & 1)}
                                _ = fi15
                                _ = v16
                            }
                        case "hello":
                            for fi17 := 0; fi17 < 19; fi17++ {
                                switch (("" + "") + ((("\n\t" + v1) + v1) + v1)) {
                                case "hello":
                                    v18 := GenS2{F0: 2147483647, F1: 479}
                                    _ = v18
                                case "\n\t":
                                    if v5 {
                                        if false {
                                            var v19 uint16 = 65535
                                            _ = v19
                                        }
                                    } else {
                                        var v20 uint8 = 255
                                        _ = v20
                                    }
                                default:
                                    var v21 float64 = 0.0
                                    v22 := GenS1{F0: 847, F1: -1}
                                    v14 = (v14 + (v4 << ((629) & 63)))
                                    _ = v21
                                    _ = v22
                                }
                                switch "\n\t" {
                                case "hello":
                                    v23 := GenS2{F0: 4294967295, F1: 0}
                                    _ = v23
                                case "a":
                                    if ((!(false || v5)) || (v5 && (fi17 > fi17))) {
                                        v1 = (("" + "foo bar") + v1)
                                        v4 = v4
                                        v24 := []int{(fi13 & fi17)}
                                        _ = v24
                                    }
                                case "":
                                    if (!v5) {
                                        var v25 float32 = -2.25
                                        _ = v25
                                    }
                                default:
                                    var v26 uint32 = 42
                                    _ = v26
                                }
                                _ = fi17
                            }
                        default:
                            switch (fi13 ^ 32768) {
                            case 0:
                                v27 := GenS1{F0: 42, F1: -1}
                                _ = v27
                            case 1:
                                v5 = true
                            }
                        }
                        _ = fi13
                        _ = v14
                    }
                case 128:
                    var v28 uint64 = 42
                    _ = v28
                case 32767:
                    for fi29 := 0; fi29 < 7; fi29++ {
                        var v30 float32 = (1.5 - (0.1 * ((-2.25 + -2.25) / 1e10)))
                        _ = fi29
                        _ = v30
                    }
                }
            case 42:
                var v31 float32 = 0.0
                _ = v31
            }
            var v32 uint = 2
            _ = v32
        } else {
            v5 = (v1 >= (v1 + "a"))
            var v33 uint16 = 255
            v34 := v3["hello"]
            _ = v33
            _ = v34
        }
        v35 := []int{}
        _ = v35
    }
    fmt.Println(v4, v2, v1)
    for fi36 := 0; fi36 < 1; fi36++ {
        var v37 float32 = (-2.25 + ((float32(3.14159) / 3.14159) + float32((0.1 / 1.0))))
        _ = fi36
        _ = v37
    }
    switch 32767 {
    case 1:
        var v38 uint = 255
        _ = v38
    default:
        v39 := GenS2{F0: 522, F1: 0}
        _ = v39
    }
    switch "hello" {
    case "":
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 float64 = (1.5 - float64(-2.25))
            _ = fi40
            _ = v41
        }
    default:
        for fi42 := 0; fi42 < 5; fi42++ {
            var v43 int32 = 784
            for fi44 := 0; fi44 < 16; fi44++ {
                if true {
                    fmt.Println(v2)
                    if (65535 >= 315) {
                        v3[v1] = v1
                        v5 = true
                        var v45 int16 = 582
                        _ = v45
                    }
                }
                fmt.Println(v5, v1)
                v46 := []int{fi42, (-(fi42))}
                _ = fi44
                _ = v46
            }
            v47 := GenS2{F0: 0, F1: 0}
            _ = fi42
            _ = v43
            _ = v47
        }
    }
    var v48 int8 = 42
    v48 += (v48 ^ 1)
    v2 = ((("\n\t" + "hello") + v2) + "\n\t")
    var v49 uint = 79
    var v50 int64 = 266
    fmt.Println(v48, v5, v49)
    if v5 {
        var v51 uint16 = 0
        var v52 uint64 = 18446744073709551615
        _ = v51
        _ = v52
    } else {
        switch "\n\t" {
        case "hello":
            v4 = v4
        case "a":
            v4 += (v4 & (v4 - v4))
        default:
            v3[v2] = ("hello" + v2)
        }
        var v53 int32 = v4
        if (2 == 1) {
            var v54 uint64 = 1
            var v55 int8 = v48
            _ = v54
            _ = v55
        }
        _ = v53
    }
    v56 := GenS2{F0: 2, F1: 155}
    v2 = "hello"
    var v57 float64 = 1.5
    if v5 {
        v49 += 2
    }
    v58 := GenS3{F0: v50}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v56
    _ = v57
    _ = v58
}
