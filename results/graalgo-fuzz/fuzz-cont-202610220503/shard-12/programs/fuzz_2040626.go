package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 string; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int }
func genFn0(p0 int64, p1 int64) uint32 {
    var v7 int8 = 1
    v8 := GenS1{F0: 2147483647, F1: p1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483648
}
func genFn1() int32 {
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            if true {
                if (42 == 2147483648) {
                    if true {
                        fmt.Println(fi9, fi10)
                        if true {
                            v11 := GenS1{F0: fi9, F1: 9223372036854775807}
                            if false {
                                switch fi10 {
                                case 42:
                                    fmt.Println(fi10)
                                case 1:
                                    if (!(42 < -976)) {
                                        var v12 float32 = 1e10
                                        var v13 int32 = 1
                                        v14 := GenS1{F0: -599, F1: -9223372036854775808}
                                        _ = v12
                                        _ = v13
                                        _ = v14
                                    } else {
                                        v15 := GenS2{F0: "foo bar", F1: 1, F2: 1}
                                        v16 := []int{fi10, fi9, fi9}
                                        var v17 bool = true
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                    }
                                default:
                                    v18 := GenS1{F0: (fi9 >> ((4294967295) & 63)), F1: 1}
                                    _ = v18
                                }
                                var v19 uint = 255
                                _ = v19
                            }
                            var v20 uint8 = 255
                            _ = v11
                            _ = v20
                        }
                        var v21 float64 = 3.14159
                        fmt.Println(fi9, fi9)
                        v21 -= float64(0.1)
                        v22 := GenS2{F0: ("" + "a"), F1: 496, F2: 2}
                        if (127 < 42) {
                            v23 := GenS2{F0: "foo bar", F1: 32767, F2: 0}
                            _ = v23
                        }
                        if (((!(false || true)) || false) || true) {
                            var v24 string = "\n\t"
                            _ = v24
                        } else {
                            v25 := []int{}
                            v26 := GenS2{F0: "a", F1: 0, F2: 2147483647}
                            if false {
                                _ = len(v25)
                                v27 := GenS1{F0: (fi10 * fi10), F1: 0}
                                v21 -= v21
                                _ = v27
                            }
                            var v28 int64 = 510
                            _ = v25
                            _ = v26
                            _ = v28
                        }
                        var v29 string = ("" + ("\n\t" + "a"))
                        _ = v21
                        _ = v22
                        _ = v29
                    } else {
                        fmt.Println(fi10)
                        switch fi10 {
                        case 2147483647:
                            if (-880 >= 0) {
                                var v30 int = 9223372036854775807
                                v31 := GenS3{F0: -781}
                                var v32 uint16 = 65535
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                        case -655:
                            fmt.Println(fi10, fi10)
                        default:
                            var v33 uint = 4294967295
                            _ = v33
                        }
                        v34 := []int{}
                        _ = v34
                    }
                    var v35 int64 = 9223372036854775807
                    _ = v35
                }
                var v36 int64 = 485
                v36 += (v36 % ((v36) | 1))
                _ = v36
            } else {
                var v37 string = ""
                if true {
                    switch (fi10 & fi10) {
                    case -2147483648:
                        var v38 string = v37
                        _ = v38
                    case 1:
                        v37 = "foo bar"
                        v39 := []int{fi9, (fi10 >> ((42) & 63)), fi10}
                        if false {
                            v39 = append(v39, (fi10 | (fi9 * (fi9 / (((fi10 << ((4294967295) & 63))) | 1)))))
                            switch (fi9 * 32767) {
                            case -1:
                                v37 = v37
                            case 42:
                                switch (fi10 & fi9) {
                                case 9223372036854775807:
                                    if (false || (genFn0(-234, 63) != 87)) {
                                        var v40 uint64 = 4294967295
                                        _ = v40
                                    }
                                case 128:
                                    var v41 int16 = -314
                                    _ = v41
                                case -803:
                                    v42 := GenS1{F0: -752, F1: 9223372036854775807}
                                    _ = v42
                                default:
                                    switch "a" {
                                    case "hello":
                                        v37 = "foo bar"
                                    case "a":
                                        v43 := GenS2{F0: v37, F1: 155, F2: genFn0(1, 1)}
                                        _ = v43
                                    }
                                }
                            }
                            if (true && ((v37 + v37) != (v37 + (("hello" + ((v37 + "a") + ((v37 + v37) + ("\n\t" + v37)))) + (v37 + (v37 + v37)))))) {
                                v37 = v37
                            }
                            v37 = (v37 + v37)
                            v44 := GenS3{F0: -652}
                            _ = v44
                        }
                        _ = v39
                    case -9223372036854775808:
                        v37 = (v37 + v37)
                    }
                } else {
                    var v45 string = v37
                    var v46 uint16 = 2
                    _ = v45
                    _ = v46
                }
                v37 = ("foo bar" + v37)
                _ = v37
            }
            switch 32767 {
            case 1:
                fmt.Println(fi9)
            }
            _ = fi10
        }
        if true {
            switch ("hello" + "hello") {
            case "":
                switch "foo bar" {
                case "hello":
                    var v47 float32 = (1e10 / ((1e10 / 3.14159) / float32((1e10 - (3.14159 - -2.25)))))
                    _ = v47
                }
            case "a":
                var v48 uint8 = 105
                _ = v48
            case "hello":
                var v49 bool = (2 != 127)
                _ = v49
            }
            for fi50 := 0; fi50 < 3; fi50++ {
                var v51 uint32 = 42
                var v52 uint32 = v51
                fmt.Println(v51)
                _ = fi50
                _ = v51
                _ = v52
            }
        }
        var v53 int = 4294967295
        _ = fi9
        _ = v53
    }
    var v54 uint8 = 129
    var v55 int64 = 1
    v54 = (^(v54))
    for fi56 := 0; fi56 < 3; fi56++ {
        v57 := []int{}
        _ = fi56
        _ = v57
    }
    _ = v54
    _ = v55
    return 42
}

func main() {
    var v1 int64 = -1
    var v2 uint16 = 32767
    var v3 float64 = float64(1.5)
    v4 := make(map[string]bool)
    var v5 uint8 = uint8(v2)
    v6 := make(map[uint32]uint32)
    var v58 uint = 257
    fmt.Println(v3, v1, v2)
    for fi59 := 0; fi59 < 13; fi59++ {
        v60 := []int{32768, -9223372036854775808, (fi59 >> (((v58 % ((v58) | 1))) & 63))}
        _ = fi59
        _ = v60
    }
    v61 := GenS4{F0: 8}
    for fi62 := 0; fi62 < 5; fi62++ {
        var v63 int16 = 989
        v64 := GenS5{F0: fi62, F1: fi62}
        _ = fi62
        _ = v63
        _ = v64
    }
    for fi65 := 0; fi65 < 4; fi65++ {
        var v66 int8 = -96
        _ = fi65
        _ = v66
    }
    var v67 uint64 = 0
    if true {
        var v68 int8 = 100
        _ = v68
    }
    var v69 int64 = (v1 ^ v1)
    var v70 uint64 = 18446744073709551615
    v3 += v3
    switch 4294967295 {
    case -1:
        var v71 uint16 = v2
        _ = v71
    case 447:
        for fi72 := 0; fi72 < 13; fi72++ {
            v73 := GenS1{F0: fi72, F1: v1}
            var v74 uint = (v58 % ((v58) | 1))
            _ = fi72
            _ = v73
            _ = v74
        }
    default:
        v75 := v4["hello"]
        _ = v75
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v61
    _ = v67
    _ = v69
    _ = v70
}
