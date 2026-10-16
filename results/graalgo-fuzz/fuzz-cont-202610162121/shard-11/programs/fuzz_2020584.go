package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 string; F1 int32; F2 int64 }
type GenS3 struct { F0 int16 }
func genFn0() int32 {
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: fi5, F1: 0}
        switch "a" {
        case "":
            if (((false && (356 >= 18446744073709551615)) || false) && false) {
                for fi7 := 0; fi7 < 1; fi7++ {
                    if (!(-704 < 42)) {
                        if (((1e10 <= (1.5 / 1.0)) && false) || true) {
                            if (!true) {
                                var v8 uint64 = 571
                                var v9 uint16 = 2
                                v10 := GenS1{F0: fi5, F1: -1}
                                _ = v10
                                _ = v8
                                _ = v9
                            } else {
                                v11 := []int{fi5}
                                _ = v11
                            }
                            var v12 float32 = 0.1
                            var v13 int64 = 9223372036854775807
                            _ = v12
                            _ = v13
                        } else {
                            switch 42 {
                            case 32767:
                                var v14 uint64 = 1
                                _ = v14
                            default:
                                var v15 int16 = -156
                                _ = v15
                            }
                            v16 := GenS2{F0: "\n\t", F1: 0, F2: -1}
                            switch "a" {
                            case "hello":
                                fmt.Println(fi5)
                            default:
                                var v17 uint8 = 125
                                _ = v17
                            }
                            _ = v16
                        }
                    }
                    _ = fi7
                }
                var v18 uint = 65535
                var v19 uint = (v18 / ((0) | 1))
                _ = v18
                _ = v19
            }
        }
        _ = fi5
        _ = v6
    }
    v20 := GenS2{F0: ("a" + ("a" + "foo bar")), F1: 0, F2: -9223372036854775808}
    v21 := []int{127}
    for fi22 := 0; fi22 < 20; fi22++ {
        fmt.Println(fi22)
        _ = fi22
    }
    _ = v20
    _ = v21
    return 1
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[uint32]int)
    var v3 uint8 = v1
    v4 := make(map[string]int)
    if ((false && (87 != 524)) || false) {
        v23 := v4[("\n\t" + "hello")]
        v3 = (^(v3))
        fmt.Println(v3)
        _ = v23
    }
    var v24 int8 = -1
    v24 += -1
    var v25 int32 = -2147483648
    var v26 int16 = 0
    fmt.Println(v26, v26)
    if (!true) {
        delete(v2, 940)
        v26 = -32768
    }
    var v27 uint8 = v1
    v28 := GenS2{F0: ("foo bar" + "hello"), F1: v25, F2: 9223372036854775807}
    v1 -= (^(v1))
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 float64 = 1e10
        if ((v30 * ((v30 * (v30 / float64((((((1.5 / ((3.14159 * 0.0) + (((3.14159 - 0.0) + -2.25) + (1.5 - ((-2.25 * 1.5) / (((0.0 * 1.5) * float32((0.0 * v30))) + float32(v30))))))) / float32((v30 * v30))) / 1e10) * -2.25) / 0.1)))) * 1e10)) == (v30 - v30)) {
            v27 = 0
            if true {
                switch fi29 {
                case 9223372036854775807:
                    v26 -= (v26 * (v26 / ((int16(v1)) | 1)))
                case 32768:
                    var v31 int16 = (v26 - v26)
                    _ = v31
                case -2147483648:
                    if (true && true) {
                        v32 := GenS1{F0: -743, F1: 0}
                        _ = v32
                    }
                default:
                    switch "foo bar" {
                    case "":
                        var v33 int16 = v26
                        _ = v33
                    case "a":
                        var v34 int16 = 0
                        _ = v34
                    case "\n\t":
                        var v35 float64 = (float64(1.5) * ((0.0 / float64(((float32(v30) * (((-2.25 + 3.14159) * (1e10 + 0.0)) / (0.0 + float32(float64((3.14159 / (1.5 * (float32((0.1 - v30)) + (1.5 + 1.5))))))))) / -2.25))) - float64((0.0 - 0.1))))
                        _ = v35
                    default:
                        v36 := GenS1{F0: fi29, F1: -972}
                        _ = v36
                    }
                }
                for fi37 := 0; fi37 < 17; fi37++ {
                    var v38 int64 = 42
                    var v39 float64 = v30
                    var v40 uint16 = 65535
                    v41 := []int{4294967295, fi37, fi29}
                    v42 := GenS2{F0: (("hello" + "") + "a"), F1: 1, F2: (v38 ^ 9223372036854775807)}
                    _ = len(v4)
                    switch "foo bar" {
                    case "a":
                        delete(v4, ("" + "a"))
                    }
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                }
            } else {
                if ((1.5 / float32((v30 / 3.14159))) == 1e10) {
                    var v43 uint32 = 4294967295
                    _ = v43
                }
                var v44 int64 = 42
                _ = v44
            }
        } else {
            var v45 uint32 = 2
            fmt.Println(v26, v26, v30)
            _ = v45
        }
        _ = fi29
        _ = v30
    }
    v46 := GenS2{F0: "a", F1: genFn0(), F2: 9223372036854775807}
    v47 := GenS3{F0: (^(v26))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v47
}
