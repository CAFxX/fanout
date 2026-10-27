package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint64; F2 string }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 uint8; F2 int32 }
type GenS8 struct { F0 bool; F1 bool }
type GenS9 struct { F0 int16; F1 uint16 }
func genFn0() uint32 {
    var v6 float64 = -2.25
    var v7 int16 = 42
    if false {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{fi8}
            v6 = v6
            switch ("hello" + "") {
            case "":
                switch (fi8 << ((2) & 63)) {
                case 128:
                    v7 -= (-(v7))
                }
            case "\n\t":
                if false {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        _ = len(v9)
                        _ = fi10
                    }
                    v11 := GenS1{F0: false, F1: 65535}
                    _ = v11
                }
            case "foo bar":
                switch 42 {
                case 0:
                    for fi12 := 0; fi12 < 1; fi12++ {
                        if (65535 < 4294967295) {
                            fmt.Println(fi8)
                            if len(v9) > 0 {
                                v9[0] = fi12
                            }
                            if (true && (0 != 25)) {
                                v13 := GenS1{F0: (2 == 4294967295), F1: 185}
                                v14 := GenS1{F0: (false || false), F1: 42}
                                _ = v13
                                _ = v14
                            } else {
                                v15 := GenS2{F0: 956}
                                _ = v15
                            }
                        }
                        v7 -= (-(v7))
                        if false {
                            switch "foo bar" {
                            case "\n\t":
                                switch fi12 {
                                case -1:
                                    fmt.Println(v6)
                                default:
                                    var v16 bool = true
                                    v17 := GenS3{F0: 9223372036854775807, F1: 1, F2: ("hello" + "\n\t")}
                                    _ = v16
                                    _ = v17
                                }
                            case "a":
                                var v18 uint = 65535
                                _ = v18
                            case "":
                                v19 := GenS4{F0: -643, F1: 684}
                                v7 = v7
                                switch fi12 {
                                case -88:
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    switch ((("" + "a") + "hello") + "foo bar") {
                                    case "foo bar":
                                        if (-1 > 816) {
                                            var v20 int = (fi8 + 930)
                                            var v21 uint = 844
                                            _ = v20
                                            _ = v21
                                        } else {
                                            switch (fi8 << ((812) & 63)) {
                                            case 127:
                                                switch (fi8 ^ fi12) {
                                                case 0:
                                                    v7 = v7
                                                default:
                                                    var v22 float64 = ((v6 * (v6 + ((float64(-2.25) + float64(3.14159)) - v6))) / v6)
                                                    _ = v22
                                                }
                                            case 0:
                                                if ((^(v7)) < v7) {
                                                    v7 = v7
                                                    v6 += 3.14159
                                                } else {
                                                    var v23 string = "a"
                                                    var v24 uint8 = 57
                                                    _ = v23
                                                    _ = v24
                                                }
                                            case 16:
                                                fmt.Println(v6, fi12)
                                            default:
                                                v7 = v7
                                            }
                                        }
                                    default:
                                        var v25 uint = 18446744073709551615
                                        _ = v25
                                    }
                                case 654:
                                    fmt.Println(fi12, fi8)
                                default:
                                    v6 += v6
                                }
                                if false {
                                    var v26 uint16 = 0
                                    _ = v26
                                }
                                _ = v19
                            default:
                                v7 -= v7
                            }
                            v27 := GenS4{F0: -9223372036854775808, F1: 2147483647}
                            v28 := GenS2{F0: 1}
                            _ = v27
                            _ = v28
                        }
                        _ = fi12
                    }
                }
            default:
                for fi29 := 0; fi29 < 10; fi29++ {
                    v9 = append(v9, (fi8 << ((1) & 63)))
                    v30 := GenS1{F0: false, F1: 2147483647}
                    v6 += 3.14159
                    _ = fi29
                    _ = v30
                }
            }
            _ = fi8
            _ = v9
        }
        v31 := GenS5{F0: 42, F1: 452}
        var v32 string = ("foo bar" + "")
        _ = v31
        _ = v32
    }
    _ = v6
    _ = v7
    return 617
}

func main() {
    var v1 float64 = ((3.14159 * (-2.25 + 0.0)) / 3.14159)
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    var v4 int64 = -64
    var v5 int8 = 42
    var v33 bool = (346 == 603)
    v34 := GenS6{F0: "a"}
    v5 -= 1
    v33 = (float32(v1) <= (float32(0.1) - -2.25))
    if ((int(v5) <= -2147483648) && (false && ("hello" >= ("\n\t" + "a")))) {
        v35 := GenS7{F0: 409, F1: 42, F2: 665}
        _ = v35
    }
    var v36 uint32 = 820
    for fi37 := 0; fi37 < 8; fi37++ {
        for fi38 := 0; fi38 < 2; fi38++ {
            v39 := GenS7{F0: 355, F1: 0, F2: 631}
            var v40 string = "foo bar"
            var v41 int32 = 1
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
        fmt.Println(fi37, v5)
        if false {
            if true {
                switch ("a" + "hello") {
                case "\n\t":
                    v5 -= (v5 / (((v5 + v5)) | 1))
                case "hello":
                    v42 := []int{}
                    _ = v42
                }
                v43 := GenS7{F0: 265, F1: 1, F2: int32(v5)}
                _ = v43
            } else {
                v44 := GenS4{F0: v4, F1: -1}
                _ = v44
            }
            fmt.Println(v1, v1)
            v36 = 2
        } else {
            var v45 float32 = 1.5
            v3["a"] = v36
            if (true || (127 < 127)) {
                fmt.Println(fi37, v5)
            }
            _ = v45
        }
        _ = fi37
    }
    var v46 int64 = v4
    v47 := GenS2{F0: 704}
    for fi48 := 0; fi48 < 10; fi48++ {
        for fi49 := 0; fi49 < 10; fi49++ {
            var v50 int = fi49
            switch (("\n\t" + "a") + "hello") {
            case "a":
                if (!(!v33)) {
                    var v51 uint16 = 485
                    _ = v51
                } else {
                    v52 := GenS8{F0: ((^(v46)) != (-(v4))), F1: v33}
                    v53 := []int{}
                    if ((18446744073709551615 >= 42) && (int(v5) > fi48)) {
                        if (true && v33) {
                            switch v50 {
                            case 863:
                                v54 := GenS9{F0: -497, F1: 1}
                                _ = v54
                            default:
                                fmt.Println(v46, fi48)
                            }
                            v1 += (v1 - ((-2.25 + v1) * float64(((1.5 - 0.0) + 0.1))))
                            delete(v3, ("hello" + ("foo bar" + "hello")))
                            if (v1 == (((v1 - v1) * 1.5) - 1e10)) {
                                v55 := GenS3{F0: v46, F1: 42, F2: "\n\t"}
                                _ = v55
                            }
                            v56 := GenS3{F0: (v46 / (((v4 % (((v4 % ((v46) | 1))) | 1))) | 1)), F1: 159, F2: "a"}
                            v53 = append(v53, (fi49 | fi49))
                            _ = v56
                        } else {
                            _ = len(v53)
                            var v57 int32 = 194
                            var v58 uint64 = 167
                            _ = v57
                            _ = v58
                        }
                        fmt.Println(v46)
                        var v59 float32 = (float32((v1 + v1)) - 1e10)
                        _ = v59
                    }
                    _ = v52
                    _ = v53
                }
            case "foo bar":
                fmt.Println(v36, fi48, fi49)
            default:
                var v60 float32 = 1.5
                _ = v60
            }
            _ = fi49
            _ = v50
        }
        _ = fi48
    }
    if v33 {
        v36 = 29
        v61 := []int{0}
        if len(v61) >= 2 {
            v61 = v61[1:len(v61)]
        }
        _ = v61
    } else {
        var v62 int16 = 42
        _ = v62
    }
    if v33 {
        var v63 int16 = 42
        switch ("hello" + "a") {
        case "\n\t":
            for fi64 := 0; fi64 < 17; fi64++ {
                fmt.Println(v36, fi64)
                var v65 uint = 0
                _ = fi64
                _ = v65
            }
        case "a":
            fmt.Println(v4, v4)
        }
        v33 = ((-9223372036854775808 == 4294967295) || v33)
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
