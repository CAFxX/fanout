package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64; F1 bool; F2 uint8 }
func genFn0() uint8 {
    v7 := GenS1{F0: false, F1: 123}
    var v8 bool = (0.0 < 0.1)
    v9 := GenS1{F0: ((645 <= -179) && (4294967295 != 2147483647)), F1: 127}
    v10 := []int{0, -1, 0, 32767}
    if (32767 > 1) {
        v11 := GenS2{F0: 787}
        v12 := GenS1{F0: v8, F1: 100}
        _ = len(v10)
        _ = v11
        _ = v12
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1() string {
    v13 := GenS2{F0: 188}
    switch -2147483648 {
    case 0:
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint64 = 18446744073709551615
            _ = fi14
            _ = v15
        }
    }
    _ = v13
    return "a"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 204
    var v3 uint32 = 949
    var v4 uint16 = (v2 ^ (^(v2)))
    v5 := make(map[string]int)
    var v6 int32 = 0
    var v16 uint8 = 128
    v17 := GenS3{F0: (v6 * (v6 - (v6 ^ 42)))}
    v18 := []int{-9223372036854775808, -1, 533, 510}
    _ = len(v18)
    switch -423 {
    case 0:
        v19 := GenS4{F0: (^(v3)), F1: 16, F2: 42}
        _ = v19
    case 32767:
        if true {
            v20 := GenS5{F0: 40}
            v21 := GenS4{F0: 2, F1: 42, F2: 0}
            switch "foo bar" {
            case "foo bar":
                v22 := GenS6{F0: 927, F1: (!(false || (-32768 == 1))), F2: 2}
                _ = v22
            }
            _ = v20
            _ = v21
        }
    default:
        var v23 float64 = (1.5 - (((-2.25 - (((0.1 + float64(-2.25)) - 1e10) - 1e10)) / 1.5) / 1.0))
        _ = v23
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        if (!true) {
            v16 -= (v16 * (v16 & v16))
        }
        _ = fi24
    }
    switch 1 {
    case -9223372036854775808:
        v25 := GenS2{F0: (v2 | (v2 | v4))}
        _ = v25
    case 32767:
        var v26 uint = 65535
        _ = v26
    default:
        switch -1 {
        case -823:
            var v27 int32 = (^(v6))
            _ = v27
        case 127:
            v3 = (v3 - 834)
        case 128:
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 uint16 = v2
                _ = fi28
                _ = v29
            }
        default:
            v4 = 255
        }
    }
    if false {
        v30 := GenS2{F0: (v4 % (((^(v4))) | 1))}
        var v31 int64 = 9223372036854775807
        var v32 int32 = -918
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 string = "foo bar"
    if false {
        v34 := GenS1{F0: true, F1: 91}
        fmt.Println(v33, v16)
        switch 1 {
        case 32768:
            delete(v5, ("foo bar" + "a"))
        case 436:
            switch 4294967295 {
            case -745:
                if true {
                    fmt.Println(v4, v6)
                    v3 = (v3 ^ v3)
                }
            case 32767:
                v2 += v2
            case 0:
                var v35 int16 = int16(v2)
                _ = v35
            }
        }
        _ = v34
    }
    if (!false) {
        for fi36 := 0; fi36 < 5; fi36++ {
            fmt.Println(v4, v33)
            v16 = v16
            _ = fi36
        }
        var v37 int64 = int64(v3)
        _ = v37
    }
    v38 := GenS5{F0: 255}
    if (!(74 >= -128)) {
        v4 += (v4 - (v4 * 1))
    }
    var v39 int = -482
    switch v39 {
    case 9223372036854775807:
        v33 = v33
    case -9223372036854775808:
        if true {
            switch 4294967295 {
            case 32767:
                var v40 uint = 248
                _ = v40
            case 904:
                switch v33 {
                case "foo bar":
                    switch v39 {
                    case -155:
                        v2 = v4
                    default:
                        var v41 bool = (genFn0() < (v16 + uint8(v3)))
                        _ = v41
                    }
                case "hello":
                    var v42 string = ""
                    _ = v42
                case "\n\t":
                    for fi43 := 0; fi43 < 6; fi43++ {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        _ = fi43
                    }
                }
            case -326:
                switch ("\n\t" + ((("hello" + v33) + v33) + "")) {
                case "hello":
                    var v44 int64 = 9223372036854775807
                    _ = v44
                case "":
                    if (0 == 770) {
                        for fi45 := 0; fi45 < 6; fi45++ {
                            fmt.Println(v2, v6, fi45)
                            _ = fi45
                        }
                        v4 = v2
                    }
                case "\n\t":
                    v6 = v6
                default:
                    v46 := GenS3{F0: (v6 - -383)}
                    _ = v46
                }
            default:
                switch 641 {
                case -9223372036854775808:
                    delete(v5, v33)
                default:
                    switch -9223372036854775808 {
                    case 0:
                        var v47 float32 = 1e10
                        _ = v47
                    default:
                        v3 += v3
                    }
                }
            }
        } else {
            v39 -= v39
        }
    default:
        for fi48 := 0; fi48 < 1; fi48++ {
            switch fi48 {
            case -131:
                v5[("a" + v33)] = v39
            default:
                if (-795 < (-(fi48))) {
                    v39 += (v39 ^ v39)
                }
            }
            _ = fi48
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
