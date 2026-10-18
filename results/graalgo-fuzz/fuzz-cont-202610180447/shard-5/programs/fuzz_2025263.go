package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int8; F1 int32; F2 int8 }
type GenS4 struct { F0 int8; F1 int8; F2 bool }
type GenS5 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]string)
    var v3 int64 = (v1 + v1)
    var v4 int32 = 42
    if (!true) {
        var v5 float64 = (0.0 + -2.25)
        _ = v5
    } else {
        if false {
            switch 42 {
            case -2147483648:
                switch 429 {
                case -9223372036854775808:
                    v6 := GenS1{F0: ("\n\t" + "a")}
                    _ = v6
                case 9223372036854775807:
                    var v7 float64 = (((0.1 - float64(1e10)) + 1.5) / 1.5)
                    _ = v7
                case 2147483647:
                    var v8 int64 = (-(v3))
                    _ = v8
                }
            }
        } else {
            var v9 int8 = 42
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 uint = 612
                _ = fi10
                _ = v11
            }
            var v12 uint64 = 288
            _ = v12
            _ = v9
        }
    }
    v13 := GenS1{F0: "hello"}
    v3 -= (v3 | (v3 + (v3 >> ((552) & 63))))
    switch "\n\t" {
    case "":
        switch "\n\t" {
        case "":
            fmt.Println(v3, v3, v3)
        }
    case "hello":
        if (0 > 127) {
            v1 = v1
            var v14 int8 = 42
            v14 = (v14 & v14)
            _ = v14
        }
    case "a":
        v15 := []int{}
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        switch (fi16 | (fi16 << ((502) & 63))) {
        case 1:
            fmt.Println(v1)
        }
        fmt.Println(fi16, v3)
        _ = fi16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        fmt.Println(v4)
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS1{F0: "\n\t"}
            if (!((!true) || false)) {
                v4 = (v4 / (((v4 & v4)) | 1))
                switch ("foo bar" + ("" + ("hello" + (("hello" + "\n\t") + "foo bar")))) {
                case "\n\t":
                    v1 = (v1 ^ v3)
                default:
                    v4 = -507
                }
            }
            _ = fi18
            _ = v19
        }
        switch "a" {
        case "foo bar":
            switch -9223372036854775808 {
            case 127:
                v4 -= v4
            case 1:
                var v20 uint32 = 0
                _ = v20
            }
        case "hello":
            v3 -= (v3 % ((v1) | 1))
        default:
            v21 := GenS2{F0: 255, F1: -32768}
            _ = v21
        }
        _ = fi17
    }
    v22 := GenS2{F0: 255, F1: 42}
    var v23 bool = false
    if false {
        switch 4294967295 {
        case 2147483647:
            if (2147483647 > -872) {
                switch "foo bar" {
                case "\n\t":
                    for fi24 := 0; fi24 < 7; fi24++ {
                        var v25 uint16 = 1
                        v26 := []int{(fi24 & fi24), (fi24 % ((-1) | 1)), (fi24 + (fi24 & fi24)), (^(fi24))}
                        var v27 uint16 = 65535
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                default:
                    v4 = 236
                }
                switch 42 {
                case 0:
                    var v28 uint = 0
                    _ = v28
                }
            }
        case 0:
            v1 = (v3 | (-(v3)))
        default:
            if (950 == 1) {
                v29 := []int{}
                v29 = append(v29, -2147483648)
                if (500 != 18446744073709551615) {
                    v30 := GenS2{F0: 255, F1: int16(v3)}
                    _ = v30
                } else {
                    var v31 float32 = ((1e10 / 1.0) - 1.5)
                    v32 := GenS2{F0: 22, F1: -182}
                    v33 := v2[2147483648]
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = v29
            } else {
                switch 32768 {
                case 1:
                    var v34 uint16 = 42
                    _ = v34
                }
                fmt.Println(v1, v4)
                switch 1 {
                case 32767:
                    var v35 int64 = v1
                    _ = v35
                }
            }
        }
        for fi36 := 0; fi36 < 16; fi36++ {
            v3 = v1
            _ = fi36
        }
    } else {
        if true {
            fmt.Println(v4)
            if (0.1 > ((-2.25 + -2.25) + 0.1)) {
                for fi37 := 0; fi37 < 15; fi37++ {
                    v23 = (float64((1.5 / (0.0 - 3.14159))) < 0.0)
                    if v23 {
                        if v23 {
                            if v23 {
                                var v38 float64 = 3.14159
                                for fi39 := 0; fi39 < 8; fi39++ {
                                    v4 -= 721
                                    if v23 {
                                        var v40 float64 = 0.1
                                        delete(v2, 42)
                                        v41 := []int{fi37, fi39, (fi37 & fi37)}
                                        v42 := GenS3{F0: 42, F1: int32(fi37), F2: -51}
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                    v2[806] = "a"
                                    var v43 uint = 65535
                                    if false {
                                        var v44 bool = (v23 && (!(635 <= 695)))
                                        _ = v44
                                    }
                                    if (v23 || (!v23)) {
                                        if false {
                                            var v45 int64 = v3
                                            if (v23 || (v43 != (v43 % ((v43) | 1)))) {
                                                switch ("a" + "a") {
                                                case "hello":
                                                    switch fi39 {
                                                    case -355:
                                                        v46 := GenS3{F0: 0, F1: v4, F2: 91}
                                                        _ = v46
                                                    }
                                                case "\n\t":
                                                    v3 = v3
                                                default:
                                                    switch "a" {
                                                    case "hello":
                                                        var v47 string = "foo bar"
                                                        _ = v47
                                                    case "a":
                                                        if (-128 > -1) {
                                                            v45 = (^(v3))
                                                            v45 = v3
                                                        }
                                                    case "foo bar":
                                                        var v48 float64 = v38
                                                        _ = v48
                                                    }
                                                    v23 = (1.5 <= 3.14159)
                                                    if v23 {
                                                        v49 := []int{127, fi39, -2147483648}
                                                        fmt.Println(v3, fi37)
                                                        _ = v49
                                                    } else {
                                                        v1 -= (v1 * (v1 << ((42) & 63)))
                                                        var v50 int8 = int8(fi37)
                                                        var v51 string = "hello"
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    var v52 bool = true
                                                    _ = v52
                                                }
                                                var v53 uint8 = 1
                                                v54 := GenS4{F0: -128, F1: 42, F2: (false && (v23 || (!("hello" > "\n\t"))))}
                                                _ = v53
                                                _ = v54
                                            }
                                            v4 = (v4 ^ 1)
                                            v55 := []int{(fi39 - fi39), fi37}
                                            _ = v45
                                            _ = v55
                                        }
                                        delete(v2, 42)
                                    }
                                    v38 += v38
                                    _ = fi39
                                    _ = v43
                                }
                                _ = v38
                            } else {
                                var v56 int64 = v1
                                _ = v56
                            }
                            for fi57 := 0; fi57 < 19; fi57++ {
                                v58 := []int{fi37}
                                var v59 float32 = (1.5 + 3.14159)
                                fmt.Println(v3, v1, v4)
                                v23 = ((0 >= 697) && v23)
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                            v2[974] = "hello"
                        } else {
                            v60 := GenS2{F0: 570, F1: -522}
                            v3 += (v1 << ((65535) & 63))
                            _ = v60
                        }
                    } else {
                        v4 = (v4 ^ -1)
                        v61 := GenS5{F0: 2147483647, F1: 255}
                        _ = v61
                    }
                    var v62 string = "a"
                    _ = fi37
                    _ = v62
                }
                var v63 int = 638
                _ = v63
            } else {
                v4 += v4
                if (true || (!(true && (352 >= 42)))) {
                    v2[2147483647] = "hello"
                } else {
                    v64 := GenS3{F0: 70, F1: (v4 | -934), F2: 0}
                    _ = v64
                }
            }
        }
        fmt.Println(v23)
    }
    v65 := []int{128, -9223372036854775808, 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v65:", len(v65), v65)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v65
}
