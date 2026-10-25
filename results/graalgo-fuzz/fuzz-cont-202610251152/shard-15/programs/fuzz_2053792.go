package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 string }
type GenS2 struct { F0 uint; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int64; F2 uint8 }
type GenS7 struct { F0 string; F1 uint64 }
func genFn0(p0 uint16, p1 int8) int64 {
    var v6 string = ""
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint64 = 0
            v10 := GenS1{F0: false, F1: (p1 * p1), F2: v6}
            v11 := GenS1{F0: true, F1: p1, F2: v6}
            var v12 uint = 776
            _ = fi8
            _ = v10
            _ = v11
            _ = v12
            _ = v9
        }
        var v13 float32 = -2.25
        if ((-(p1)) >= (p1 >> ((255) & 63))) {
            var v14 uint32 = 4294967295
            fmt.Println(v14, v13, fi7)
            v15 := []int{}
            _ = v14
            _ = v15
        }
        _ = fi7
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint = 0
    var v4 uint64 = 566
    var v5 bool = false
    var v16 int64 = 1
    var v17 int32 = 2147483647
    v5 = (!((!v5) || v5))
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := []int{(fi19 ^ (fi19 | (fi19 << ((v3) & 63))))}
            v4 += v4
            _ = fi19
            _ = v20
        }
        if (((1.5 / (1e10 / 3.14159)) * ((1e10 + 1e10) + (3.14159 + 1e10))) == 3.14159) {
            var v21 string = (("foo bar" + ("hello" + "foo bar")) + ("hello" + "hello"))
            _ = v21
        }
        fmt.Println(v3, v4)
        _ = fi18
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        v5 = (((("foo bar" + "\n\t") + "hello") + "") <= ("\n\t" + ((("hello" + "\n\t") + "\n\t") + "hello")))
        v4 = v4
        v23 := GenS2{F0: (v3 / (((v3 * v3)) | 1)), F1: 1000, F2: 199}
        _ = fi22
        _ = v23
    }
    if v5 {
        var v24 int32 = v17
        if v5 {
            v16 += (v16 / ((v16) | 1))
            v25 := GenS3{F0: 127}
            v26 := GenS3{F0: 128}
            _ = v25
            _ = v26
        }
        _ = v24
    }
    if false {
        v3 = (v3 / (((v3 / ((961) | 1))) | 1))
    } else {
        _ = len(v2)
    }
    v3 = v3
    if (3.14159 != 0.0) {
        v17 = (v17 % (((^(v17))) | 1))
    }
    v17 -= v17
    v3 = v3
    if true {
        v27 := []int{-571, 4294967295, -1}
        v5 = (!true)
        var v28 int8 = 1
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 bool = false
        if ("\n\t" >= "hello") {
            fmt.Println(v3, v5)
            for fi31 := 0; fi31 < 11; fi31++ {
                switch ("\n\t" + "hello") {
                case "hello":
                    if v5 {
                        var v32 uint64 = 18446744073709551615
                        switch "a" {
                        case "\n\t":
                            v5 = (!v30)
                        case "hello":
                            fmt.Println(v30, v4)
                        default:
                            if true {
                                var v33 uint64 = 647
                                _ = v33
                            }
                        }
                        fmt.Println(v17)
                        if ("hello" < ("" + "")) {
                            var v34 int8 = 42
                            if (false && v5) {
                                var v35 float32 = (3.14159 * ((1.5 - 1.5) / (0.0 - 1e10)))
                                v34 = -44
                                v36 := GenS4{F0: 127, F1: 1}
                                var v37 uint32 = 695
                                _ = len(v2)
                                _ = v35
                                _ = v36
                                _ = v37
                            } else {
                                v3 += v3
                                switch (fi31 ^ fi29) {
                                case 1:
                                    v38 := GenS5{F0: 255}
                                    _ = v38
                                case -2147483648:
                                    var v39 int8 = (v34 & (v34 | v34))
                                    _ = v39
                                case -9223372036854775808:
                                    switch fi29 {
                                    case -1:
                                        fmt.Println(v4)
                                    case 42:
                                        if false {
                                            v40 := GenS6{F0: 1, F1: (^(v16)), F2: 128}
                                            _ = v40
                                        } else {
                                            v1["hello"] = 58
                                            v34 = v34
                                        }
                                    default:
                                        var v41 uint8 = 142
                                        _ = v41
                                    }
                                }
                                delete(v1, "\n\t")
                            }
                            v3 += v3
                            _ = v34
                        }
                        _ = v32
                    } else {
                        var v42 int16 = -1
                        switch (fi29 ^ -2147483648) {
                        case -1:
                            fmt.Println(v42, v4)
                        case 1:
                            v5 = false
                        case 32768:
                            v43 := GenS7{F0: "\n\t", F1: v4}
                            _ = v43
                        default:
                            fmt.Println(v42)
                        }
                        v2[42] = v30
                        if v30 {
                            var v44 uint = 409
                            v16 = (v16 & v16)
                            _ = v44
                        } else {
                            var v45 float32 = 1.5
                            v46 := GenS7{F0: "foo bar", F1: (v4 ^ (^(v4)))}
                            v47 := GenS1{F0: v5, F1: -128, F2: ""}
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        _ = v42
                    }
                case "\n\t":
                    v48 := []int{(fi29 * fi31), 42, fi31}
                    _ = v48
                default:
                    v3 = 255
                }
                switch 127 {
                case 127:
                    var v49 int32 = -2147483648
                    _ = v49
                }
                var v50 int32 = v17
                _ = fi31
                _ = v50
            }
            switch "hello" {
            case "":
                if (-2.25 >= (0.1 / ((float32(-2.25) - ((float32(0.1) / 1.5) + ((1e10 / 1.0) * 0.1))) * (0.1 * 0.1)))) {
                    switch (fi29 | (-(fi29))) {
                    case 534:
                        var v51 int32 = 1
                        _ = v51
                    case -1:
                        fmt.Println(v4)
                    case -405:
                        switch "" {
                        case "a":
                            v5 = false
                        case "":
                            if v5 {
                                v52 := []int{fi29, (fi29 | fi29), fi29}
                                _ = v52
                            } else {
                                for fi53 := 0; fi53 < 19; fi53++ {
                                    if (!(4294967295 == v4)) {
                                        v3 = v3
                                        var v54 uint8 = 124
                                        _ = v54
                                    }
                                    fmt.Println(v16, v3, fi53)
                                    _ = fi53
                                }
                                v5 = ((v16 << ((65535) & 63)) >= (v16 & (v16 % ((0) | 1))))
                            }
                        }
                    }
                }
            }
        }
        v55 := []int{(fi29 << (((v3 ^ (v3 % ((v3) | 1)))) & 63)), fi29, 32768, -1}
        _ = fi29
        _ = v30
        _ = v55
    }
    if v5 {
        v16 = v16
        for fi56 := 0; fi56 < 10; fi56++ {
            var v57 int16 = 1000
            var v58 string = ""
            fmt.Println(fi56, v3)
            _ = fi56
            _ = v57
            _ = v58
        }
        if (42 >= 1) {
            fmt.Println(v3)
        } else {
            v59 := []int{}
            _ = v59
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
