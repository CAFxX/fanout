package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 uint16; F1 string; F2 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 string }

func main() {
    v1 := []int{}
    var v2 int64 = 89
    var v3 float32 = (float32(1.5) - 0.0)
    var v4 uint = 458
    var v5 int64 = (v2 / (((v2 % ((-1) | 1))) | 1))
    v6 := GenS1{F0: -1, F1: 65535, F2: ("a" + "foo bar")}
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(v4, v5)
        fmt.Println(v3, v4, fi7)
        var v8 int16 = -721
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "a":
        var v9 uint64 = 2
        _ = v9
    }
    fmt.Println(v5)
    for fi10 := 0; fi10 < 3; fi10++ {
        if (-571 >= -1) {
            var v11 bool = false
            _ = v11
        } else {
            var v12 float32 = v3
            _ = v12
        }
        _ = fi10
    }
    v1 = append(v1, 128)
    if (true && ((!true) || false)) {
        v13 := GenS1{F0: 127, F1: 288, F2: (("\n\t" + "foo bar") + ("a" + "a"))}
        v3 = v3
        var v14 int16 = 42
        _ = v13
        _ = v14
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v3)
    case "foo bar":
        v4 -= v4
    case "\n\t":
        if ((v4 >= v4) && true) {
            if false {
                if true {
                    fmt.Println(v5, v3)
                } else {
                    v15 := GenS1{F0: 0, F1: 2, F2: "hello"}
                    v16 := GenS1{F0: -128, F1: 362, F2: ("" + "foo bar")}
                    var v17 float64 = (0.0 / 0.1)
                    _ = v15
                    _ = v16
                    _ = v17
                }
                v3 = (v3 - (v3 + v3))
            }
            for fi18 := 0; fi18 < 14; fi18++ {
                if len(v1) > 0 {
                    v1[0] = fi18
                }
                if (((128 < 220) || false) && true) {
                    if true {
                        _ = len(v1)
                        v19 := GenS2{F0: 255, F1: "hello", F2: false}
                        v4 = v4
                        _ = v19
                    }
                    v20 := GenS2{F0: 358, F1: ("foo bar" + "hello"), F2: (!true)}
                    _ = v20
                }
                _ = fi18
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 uint = (v4 | (v4 / ((v4) | 1)))
                v23 := GenS1{F0: 92, F1: 789, F2: (("" + (((("\n\t" + "") + "foo bar") + "foo bar") + "\n\t")) + "a")}
                v2 += -622
                _ = fi21
                _ = v22
                _ = v23
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v4 = v4
        }
    default:
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 int = fi24
            for fi26 := 0; fi26 < 2; fi26++ {
                var v27 uint = (v4 >> (((v4 + (v4 - (^(v4))))) & 63))
                switch fi24 {
                case -1:
                    switch 32767 {
                    case -373:
                        var v28 uint64 = 362
                        fmt.Println(v28, v4)
                        if false {
                            v2 += (v5 ^ v5)
                        } else {
                            switch fi24 {
                            case 32767:
                                switch "foo bar" {
                                case "":
                                    v3 -= 3.14159
                                default:
                                    var v29 int = (v25 + fi24)
                                    _ = v29
                                }
                            case 0:
                                v30 := GenS2{F0: 407, F1: "\n\t", F2: false}
                                _ = v30
                            }
                            v31 := GenS1{F0: 107, F1: 0, F2: ("a" + "a")}
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                            _ = len(v1)
                            if len(v1) > 0 {
                                v1[0] = (fi24 / ((fi24) | 1))
                            }
                            var v32 int = (v25 ^ v25)
                            _ = len(v1)
                            var v33 int16 = -32768
                            if ((v2 >> ((v27) & 63)) != v2) {
                                _ = len(v1)
                                if false {
                                    var v34 int8 = 0
                                    v2 = v5
                                    v4 -= 573
                                    _ = v34
                                } else {
                                    var v35 float32 = v3
                                    _ = v35
                                }
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v36 int64 = (v5 | -135)
                                _ = v36
                            }
                            v37 := GenS3{F0: 1}
                            if false {
                                v2 = (v2 << ((v4) & 63))
                            } else {
                                switch "foo bar" {
                                case "foo bar":
                                    v2 = -1
                                case "\n\t":
                                    if (!false) {
                                        var v38 int32 = -442
                                        if (v28 < 18446744073709551615) {
                                            v39 := GenS2{F0: 495, F1: "hello", F2: (true && true)}
                                            _ = v39
                                        } else {
                                            var v40 string = ""
                                            var v41 uint32 = 2147483647
                                            if (v41 > (^(v41))) {
                                                var v42 uint8 = 127
                                                v43 := GenS3{F0: (v42 & 0)}
                                                _ = v42
                                                _ = v43
                                            } else {
                                                v4 = 18446744073709551615
                                            }
                                            v1 = append(v1, (fi24 * (^(fi24))))
                                            _ = v40
                                            _ = v41
                                        }
                                        _ = v38
                                    }
                                case "hello":
                                    v3 = 1e10
                                    switch "a" {
                                    case "\n\t":
                                        var v44 uint = 0
                                        _ = v44
                                    case "hello":
                                        if len(v1) > 0 {
                                            v1[0] = (fi24 * (fi26 >> (((v4 | 2)) & 63)))
                                        }
                                    }
                                default:
                                    v45 := GenS2{F0: 65535, F1: "a", F2: true}
                                    _ = v45
                                }
                                fmt.Println(v4)
                                v46 := GenS3{F0: uint8(v4)}
                                _ = v46
                            }
                            _ = v31
                            _ = v32
                            _ = v33
                            _ = v37
                        }
                        _ = v28
                    case 42:
                        v25 -= fi24
                    }
                case 972:
                    v5 += (v2 & (^(v5)))
                case -632:
                    if (true || (-1 == 42)) {
                        v5 -= (v2 % ((-9223372036854775808) | 1))
                        fmt.Println(v3)
                        var v47 uint16 = 985
                        _ = v47
                    } else {
                        if true {
                            var v48 uint64 = 2
                            switch "a" {
                            case "foo bar":
                                if true {
                                    var v49 uint16 = 55
                                    v25 = (v25 - (fi26 ^ 127))
                                    _ = v49
                                }
                            default:
                                v5 = v2
                                v5 = int64(v48)
                                var v50 uint64 = uint64(v4)
                                if (v3 <= v3) {
                                    v3 = (float32((-2.25 + 1e10)) - v3)
                                } else {
                                    var v51 uint8 = 128
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch "" {
                                    case "hello":
                                        if (42 >= 42) {
                                            if true {
                                                var v52 float64 = 0.0
                                                var v53 uint32 = 0
                                                if (float32((((v52 - v52) + v52) * v52)) > (v3 * (1.5 / float32((((1e10 - ((v52 + v52) - (v52 + 0.0))) * v52) / (float64(v3) + 1.5)))))) {
                                                    if false {
                                                        if ((v4 << (((v27 | 4294967295)) & 63)) <= 65535) {
                                                            switch v25 {
                                                            case -2147483648:
                                                                switch (-(fi26)) {
                                                                case 4294967295:
                                                                    fmt.Println(v27, fi26, v3)
                                                                }
                                                            case -1:
                                                                var v54 int32 = 0
                                                                _ = v54
                                                            case 127:
                                                                if true {
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        var v55 uint = v4
                                                                        _ = v55
                                                                    default:
                                                                        var v56 uint16 = 856
                                                                        _ = v56
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        var v57 int32 = 42
                                                        v3 += 3.14159
                                                        _ = v57
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = fi24
                                                    }
                                                    var v58 int = (-(v25))
                                                    v1 = append(v1, v25)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v58
                                                }
                                                switch ("" + ("" + ((("" + ("a" + "foo bar")) + "") + "hello"))) {
                                                case "":
                                                    var v59 uint64 = (v48 - 4294967295)
                                                    _ = v59
                                                }
                                                _ = v52
                                                _ = v53
                                            }
                                        }
                                    case "\n\t":
                                        if (!(v2 >= 42)) {
                                            v60 := GenS2{F0: 32767, F1: ("a" + "foo bar"), F2: true}
                                            v2 = (v2 ^ (v5 ^ (^(v2))))
                                            _ = len(v1)
                                            _ = v60
                                        } else {
                                            var v61 string = ("" + "")
                                            v25 = 35
                                            _ = v61
                                        }
                                        var v62 uint64 = v50
                                        _ = v62
                                    }
                                    _ = v51
                                }
                                _ = v50
                            }
                            v63 := GenS1{F0: 0, F1: 154, F2: "\n\t"}
                            _ = v48
                            _ = v63
                        }
                        var v64 uint8 = 127
                        var v65 uint64 = 18446744073709551615
                        _ = v64
                        _ = v65
                    }
                    switch ("a" + "hello") {
                    case "hello":
                        fmt.Println(v27, v2)
                    case "\n\t":
                        v1 = append(v1, fi26)
                    }
                }
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
    }
    v1 = append(v1, 42)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi66 := 0; fi66 < 2; fi66++ {
        if (!true) {
            v67 := GenS4{F0: 0}
            _ = v67
        }
        v68 := GenS5{F0: 990, F1: "a"}
        for fi69 := 0; fi69 < 11; fi69++ {
            fmt.Println(v5, v3)
            var v70 int16 = 1
            fmt.Println(fi69, v70, fi66)
            _ = fi69
            _ = v70
        }
        _ = fi66
        _ = v68
    }
    var v71 bool = (2 < 18446744073709551615)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
}
