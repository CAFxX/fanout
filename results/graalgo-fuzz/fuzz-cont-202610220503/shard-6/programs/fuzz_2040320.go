package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int32; F1 int64; F2 int }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint64; F1 int8; F2 string }
type GenS5 struct { F0 uint64; F1 int16; F2 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 int; F2 int16 }

func main() {
    var v1 float64 = (float64(((1e10 + (3.14159 - float32(1.5))) + float32(1e10))) + (3.14159 + (1.5 + 1.5)))
    var v2 uint8 = 42
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 uint = 0
    v6 := GenS1{F0: v2, F1: v5}
    if false {
        if true {
            if true {
                var v7 int = 1
                fmt.Println(v7)
                v8 := v3[(v7 ^ v7)]
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 11; fi9++ {
                    if ((fi9 << ((v5) & 63)) >= (-(fi9))) {
                        var v10 float64 = ((v1 - ((v1 - (((v1 * 1.5) / v1) - 0.1)) / v1)) / float64(0.1))
                        if false {
                            var v11 int32 = 42
                            v1 -= 0.1
                            v2 -= v2
                            _ = v11
                        } else {
                            var v12 bool = (v2 == (v2 * v2))
                            _ = v12
                        }
                        _ = v10
                    } else {
                        if ((^(v2)) > v2) {
                            switch 32767 {
                            case 32768:
                                var v13 int32 = 42
                                _ = v13
                            case 9223372036854775807:
                                if true {
                                    v5 = 18446744073709551615
                                    v1 += (v1 * v1)
                                    for fi14 := 0; fi14 < 17; fi14++ {
                                        v15 := []int{-1, (fi9 % ((32767) | 1)), (fi14 % ((fi9) | 1)), fi14}
                                        if true {
                                            v16 := GenS1{F0: v2, F1: 1}
                                            _ = v16
                                        } else {
                                            if (365 > -335) {
                                                var v17 bool = false
                                                v18 := GenS2{F0: -1, F1: -9223372036854775808, F2: 2147483647}
                                                _ = v17
                                                _ = v18
                                            } else {
                                                fmt.Println(fi9, v2)
                                            }
                                            fmt.Println(fi9, fi14, v1)
                                        }
                                        _ = fi14
                                        _ = v15
                                    }
                                }
                            case 1:
                                fmt.Println(fi9, fi9)
                            }
                        } else {
                            v2 += v2
                            for fi19 := 0; fi19 < 15; fi19++ {
                                var v20 uint16 = 2
                                switch "" {
                                case "a":
                                    if ((-374 < -1) && true) {
                                        v21 := v4["foo bar"]
                                        switch (fi9 | fi19) {
                                        case 128:
                                            v22 := GenS2{F0: 2147483647, F1: 42, F2: (fi9 ^ (-(fi9)))}
                                            _ = v22
                                        default:
                                            var v23 int32 = 2147483647
                                            _ = v23
                                        }
                                        _ = v21
                                    }
                                case "":
                                    switch (fi9 % ((-2147483648) | 1)) {
                                    case 1:
                                        v24 := []int{}
                                        _ = v24
                                    default:
                                        var v25 bool = true
                                        _ = v25
                                    }
                                }
                                _ = fi19
                                _ = v20
                            }
                        }
                    }
                    _ = fi9
                }
                var v26 string = "\n\t"
                _ = v26
            }
            var v27 uint = (v5 << ((v5) & 63))
            _ = v27
        }
    } else {
        if true {
            fmt.Println(v1, v1, v5)
            switch 820 {
            case 32768:
                v28 := v3[-1]
                _ = v28
            case 42:
                v1 += ((float64((1e10 * 0.0)) - v1) - 0.0)
            case 128:
                if (false && (false && false)) {
                    for fi29 := 0; fi29 < 5; fi29++ {
                        if true {
                            if (true && (0 == 255)) {
                                v2 = v2
                                var v30 int16 = 0
                                _ = v30
                            } else {
                                v31 := GenS3{F0: 0, F1: fi29}
                                _ = v31
                            }
                            v2 -= (^(v2))
                        }
                        var v32 string = (("hello" + (("foo bar" + "\n\t") + "foo bar")) + "")
                        _ = fi29
                        _ = v32
                    }
                    if true {
                        var v33 string = "foo bar"
                        delete(v3, 42)
                        switch 2147483647 {
                        case 42:
                            v2 += v2
                        }
                        _ = v33
                    } else {
                        v5 = (^(v5))
                    }
                } else {
                    fmt.Println(v5, v2, v2)
                    v34 := []int{4294967295, 62, 2147483647}
                    var v35 int16 = 1
                    _ = v34
                    _ = v35
                }
            default:
                switch "\n\t" {
                case "a":
                    switch "a" {
                    case "":
                        v5 = 2
                    case "\n\t":
                        for fi36 := 0; fi36 < 11; fi36++ {
                            var v37 uint64 = 1
                            v38 := GenS4{F0: uint64(v2), F1: 100, F2: ((("\n\t" + (("a" + "\n\t") + "a")) + "") + "\n\t")}
                            var v39 float64 = v1
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    case "foo bar":
                        if true {
                            v40 := GenS4{F0: 832, F1: -1, F2: "\n\t"}
                            for fi41 := 0; fi41 < 2; fi41++ {
                                v3[fi41] = false
                                if false {
                                    if false {
                                        if (("a" + "hello") > "foo bar") {
                                            if false {
                                                v5 -= 65535
                                            } else {
                                                var v42 float64 = v1
                                                _ = v42
                                            }
                                            v43 := GenS4{F0: 25, F1: 127, F2: "\n\t"}
                                            _ = v43
                                        } else {
                                            v5 += 151
                                            var v44 uint16 = 0
                                            _ = v44
                                        }
                                    }
                                }
                                _ = fi41
                            }
                            _ = v40
                        }
                    }
                case "foo bar":
                    var v45 int64 = 9223372036854775807
                    _ = v45
                case "hello":
                    v46 := GenS1{F0: 99, F1: (v5 - (^(v5)))}
                    _ = v46
                }
            }
            var v47 uint16 = 2
            _ = v47
        }
        switch 2147483647 {
        case 4294967295:
            v48 := GenS5{F0: 18446744073709551615, F1: -32768, F2: -983}
            _ = v48
        case 564:
            v5 -= (v5 >> ((v5) & 63))
        }
    }
    _ = len(v3)
    fmt.Println(v2, v5)
    for fi49 := 0; fi49 < 1; fi49++ {
        for fi50 := 0; fi50 < 1; fi50++ {
            var v51 uint8 = (^(v2))
            switch (fi50 >> (((^(v5))) & 63)) {
            case -2147483648:
                var v52 string = "foo bar"
                _ = v52
            case 877:
                switch ("" + "hello") {
                case "foo bar":
                    v2 = (v51 & v51)
                case "a":
                    var v53 int64 = -9223372036854775808
                    _ = v53
                }
            }
            _ = fi50
            _ = v51
        }
        v54 := GenS6{F0: 42}
        _ = fi49
        _ = v54
    }
    fmt.Println(v5, v5)
    fmt.Println(v5, v2, v1)
    if false {
        var v55 int8 = -101
        _ = v55
    }
    if true {
        v4["foo bar"] = "foo bar"
    }
    var v56 float32 = 1e10
    v1 += ((((v1 - (float64(v56) * v1)) / v1) / v1) / (0.0 / v1))
    switch 32767 {
    case 2147483647:
        var v57 float64 = 1.5
        _ = v57
    case -147:
        v58 := GenS7{F0: 42}
        _ = v58
    }
    v59 := GenS8{F0: 42, F1: 451, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v59
    _ = v6
}
