package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int8; F1 uint64; F2 string }
type GenS5 struct { F0 uint16; F1 uint }

func main() {
    var v1 uint = 42
    var v2 uint64 = 798
    var v3 int = 127
    var v4 int8 = 0
    var v5 uint32 = 65535
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 255, F1: -721, F2: 42}
        v8 := GenS1{F0: 101, F1: -9223372036854775808, F2: 1}
        _ = fi6
        _ = v7
        _ = v8
    }
    v3 = (v3 * -560)
    v4 += (v4 & 1)
    v5 += (v5 + v5)
    if true {
        for fi9 := 0; fi9 < 7; fi9++ {
            if false {
                if false {
                    var v10 uint = v1
                    fmt.Println(v4)
                    _ = v10
                }
                for fi11 := 0; fi11 < 9; fi11++ {
                    if false {
                        if (true || (9223372036854775807 < 461)) {
                            v5 = v5
                        } else {
                            switch fi9 {
                            case 194:
                                if (-1 >= 1) {
                                    v3 += (v3 + fi11)
                                    if (v4 >= (v4 ^ 1)) {
                                        if false {
                                            v5 -= v5
                                            v4 = int8(v3)
                                        } else {
                                            v3 += fi9
                                        }
                                        v1 -= (^(v1))
                                        v2 = (v2 - v2)
                                    } else {
                                        v12 := GenS2{F0: 337, F1: "foo bar"}
                                        if true {
                                            var v13 uint16 = 42
                                            var v14 string = "a"
                                            v1 = v1
                                            if true {
                                                v2 += (v2 | 266)
                                                var v15 float32 = float32((3.14159 + 1.5))
                                                _ = v15
                                            } else {
                                                var v16 int = -849
                                                _ = v16
                                            }
                                            var v17 string = v14
                                            _ = v13
                                            _ = v14
                                            _ = v17
                                        } else {
                                            switch (-(v3)) {
                                            case -1:
                                                v18 := []int{fi11, 32768, 1, v3}
                                                _ = v18
                                            case -2147483648:
                                                if (true || (!true)) {
                                                    fmt.Println(fi11, v2, fi9)
                                                    v19 := GenS1{F0: 129, F1: 1, F2: 0}
                                                    fmt.Println(fi11, v1, v4)
                                                    _ = v19
                                                }
                                            case 2147483647:
                                                v3 = 2147483647
                                            default:
                                                v20 := GenS3{F0: -32768, F1: (v5 ^ 2), F2: 101}
                                                _ = v20
                                            }
                                            if true {
                                                v1 = 4294967295
                                                var v21 uint32 = 65535
                                                _ = v21
                                            } else {
                                                if false {
                                                    v22 := GenS1{F0: 175, F1: -834, F2: 32767}
                                                    var v23 int8 = v4
                                                    var v24 int8 = v23
                                                    _ = v22
                                                    _ = v23
                                                    _ = v24
                                                } else {
                                                    v1 -= 255
                                                }
                                            }
                                        }
                                        _ = v12
                                    }
                                } else {
                                    var v25 int = fi9
                                    if (42 >= 718) {
                                        v5 = v5
                                    } else {
                                        v3 = (v3 ^ (v25 / (((v3 * v3)) | 1)))
                                        v25 = (fi9 >> ((v1) & 63))
                                        var v26 int16 = 0
                                        _ = v26
                                    }
                                    v27 := []int{v3, v25, v25}
                                    _ = v25
                                    _ = v27
                                }
                            case 127:
                                if (!(2 > 2)) {
                                    var v28 float64 = (((3.14159 + float64(0.1)) / float64((((1.5 + (0.1 + (3.14159 + ((1e10 / (1.5 - (1e10 + (-2.25 * 1e10)))) * -2.25)))) / (0.1 + -2.25)) + -2.25))) * 1e10)
                                    _ = v28
                                }
                            default:
                                fmt.Println(v1)
                            }
                            v2 -= v2
                        }
                        switch (fi11 - (fi9 - 2147483647)) {
                        case 42:
                            switch 32768 {
                            case 128:
                                var v29 int32 = 42
                                _ = v29
                            case 1:
                                var v30 int16 = 944
                                _ = v30
                            case 4294967295:
                                v1 -= 2
                            }
                        case -736:
                            var v31 uint64 = v2
                            _ = v31
                        case 127:
                            if ((v2 > v2) && false) {
                                var v32 bool = (true || true)
                                _ = v32
                            } else {
                                if (false && false) {
                                    v4 += v4
                                    if true {
                                        switch "" {
                                        case "hello":
                                            v2 = (v2 | (v2 - v2))
                                        case "":
                                            v3 = v3
                                        case "\n\t":
                                            fmt.Println(v1, v3, fi11)
                                        default:
                                            switch "\n\t" {
                                            case "a":
                                                fmt.Println(v4)
                                            case "foo bar":
                                                fmt.Println(v2)
                                            }
                                        }
                                        v33 := []int{359}
                                        _ = v33
                                    }
                                }
                                var v34 int16 = 32767
                                v35 := GenS4{F0: (-(v4)), F1: v2, F2: ("hello" + ("foo bar" + "a"))}
                                _ = v34
                                _ = v35
                            }
                        default:
                            v2 -= (v2 >> ((v1) & 63))
                        }
                        fmt.Println(fi11, v5, v5)
                    } else {
                        switch ("a" + ("foo bar" + "foo bar")) {
                        case "hello":
                            if true {
                                v4 = 42
                                switch (-(fi11)) {
                                case 128:
                                    if (3.14159 > 3.14159) {
                                        if true {
                                            var v36 uint16 = 1
                                            _ = v36
                                        }
                                        v37 := GenS5{F0: 42, F1: v1}
                                        _ = v37
                                    }
                                case 127:
                                    v38 := []int{2147483647}
                                    _ = v38
                                }
                                v39 := []int{fi11, v3, (-(fi9)), fi9}
                                _ = v39
                            }
                        case "\n\t":
                            v5 -= (v5 % ((2147483647) | 1))
                        case "a":
                            switch ("" + ("foo bar" + "\n\t")) {
                            case "a":
                                var v40 int16 = 1
                                _ = v40
                            case "foo bar":
                                fmt.Println(v4, v1, fi9)
                            default:
                                v5 -= (^(v5))
                            }
                        default:
                            if false {
                                var v41 int = (v3 * (fi9 - fi9))
                                v3 -= 2147483647
                                _ = v41
                            }
                        }
                    }
                    _ = fi11
                }
                v42 := []int{-840, (fi9 | 717), 127, 9223372036854775807}
                _ = v42
            } else {
                var v43 int8 = v4
                v3 -= (v3 % ((2147483647) | 1))
                _ = v43
            }
            var v44 uint8 = 34
            _ = fi9
            _ = v44
        }
        var v45 uint32 = (^(v5))
        _ = v45
    }
    fmt.Println(v4)
    var v46 uint16 = 2
    v5 -= (v5 % ((v5) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
