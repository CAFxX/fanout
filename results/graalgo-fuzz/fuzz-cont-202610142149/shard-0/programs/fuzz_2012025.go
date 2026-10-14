package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int16 }
func genFn0() int64 {
    var v5 bool = ("a" < ("a" + "foo bar"))
    var v6 int64 = 0
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint64 = 18446744073709551615
        var v9 uint16 = 42
        if v5 {
            if ((((!false) || v5) || v5) || ((65535 == 65535) || false)) {
                for fi10 := 0; fi10 < 16; fi10++ {
                    v11 := GenS1{F0: -2147483648, F1: 255}
                    var v12 uint32 = 1
                    switch fi7 {
                    case 1:
                        v13 := []int{42}
                        _ = v13
                    case -2147483648:
                        fmt.Println(v12)
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            } else {
                switch "a" {
                case "a":
                    v5 = false
                default:
                    if false {
                        v14 := GenS2{F0: 100, F1: (^(v9)), F2: 42}
                        _ = v14
                    }
                }
                var v15 uint64 = 0
                var v16 int8 = 42
                _ = v15
                _ = v16
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 825
    var v2 uint64 = 1
    v3 := []string{"a", "", ("hello" + "hello"), ("a" + ("a" + "\n\t"))}
    v4 := make(map[uint32]int)
    v17 := GenS2{F0: 127, F1: 42, F2: 0}
    if (v2 < 0) {
        _ = len(v3)
        if true {
            fmt.Println(v1, v2)
            switch -1 {
            case 9223372036854775807:
                v2 -= 0
            case -2147483648:
                v2 = v2
            default:
                if false {
                    v18 := GenS2{F0: 42, F1: (v1 + (^(v1))), F2: 2147483647}
                    v2 = 4294967295
                    v19 := v4[2147483648]
                    _ = v18
                    _ = v19
                }
            }
        }
    }
    if true {
        var v20 int32 = 42
        v21 := GenS3{F0: 42, F1: genFn0()}
        _ = v20
        _ = v21
    } else {
        if (true || ((v1 & 429) > (v1 | 172))) {
            v22 := GenS3{F0: 159, F1: 121}
            v1 += v1
            _ = v22
        }
        switch ("a" + ("" + ("foo bar" + "hello"))) {
        case "hello":
            var v23 bool = true
            _ = v23
        default:
            var v24 uint = 1
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 int64 = -1
        if true {
            switch "hello" {
            case "a":
                v26 = (v26 - (^(v26)))
            case "":
                v1 = (v1 - (v1 % ((2) | 1)))
            default:
                v26 = genFn0()
            }
            v3 = append(v3, "\n\t")
        } else {
            v1 += (v1 - 255)
        }
        for fi27 := 0; fi27 < 9; fi27++ {
            switch fi25 {
            case 32767:
                switch "\n\t" {
                case "\n\t":
                    v28 := v4[2]
                    _ = v28
                default:
                    v26 += 0
                }
            case 0:
                var v29 uint32 = 42
                _ = v29
            default:
                var v30 float64 = ((1e10 - ((1e10 - float64((1.5 / (0.1 * 1e10)))) * (1.5 - 3.14159))) / -2.25)
                _ = v30
            }
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    fmt.Println(v2, v1)
    fmt.Println(v1, v2)
    var v31 float32 = -2.25
    if true {
        if (((("hello" + ("" + "a")) == ("hello" + ((("foo bar" + "") + ("foo bar" + "foo bar")) + ("foo bar" + "hello")))) || (2 > 42)) && (112 < 42)) {
            v32 := GenS1{F0: -2147483648, F1: 255}
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            for fi33 := 0; fi33 < 2; fi33++ {
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 string = "a"
                    v35 = (v35 + "hello")
                    v1 += 32767
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
            _ = v32
        }
        fmt.Println(v2)
        if false {
            v4[1] = -2147483648
        } else {
            var v36 uint64 = v2
            _ = v36
        }
    } else {
        var v37 float64 = ((float64(0.1) / 1.5) + 0.1)
        v2 += v2
        _ = v37
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        v3 = append(v3, "")
        if true {
            _ = len(v4)
            switch "foo bar" {
            case "":
                for fi39 := 0; fi39 < 5; fi39++ {
                    v2 = 42
                    fmt.Println(fi38, v31, v2)
                    _ = fi39
                }
            case "hello":
                for fi40 := 0; fi40 < 11; fi40++ {
                    var v41 uint16 = v1
                    _ = fi40
                    _ = v41
                }
            default:
                v42 := GenS1{F0: 2147483647, F1: 1}
                _ = v42
            }
        } else {
            v2 += v2
            var v43 uint16 = (v1 % ((v1) | 1))
            _ = v43
        }
        _ = fi38
    }
    if true {
        var v44 uint8 = 255
        for fi45 := 0; fi45 < 19; fi45++ {
            if (false && true) {
                v44 -= 93
                var v46 int8 = 0
                var v47 bool = false
                _ = v46
                _ = v47
            }
            for fi48 := 0; fi48 < 13; fi48++ {
                switch "" {
                case "foo bar":
                    switch fi48 {
                    case 47:
                        fmt.Println(v1, v2)
                    default:
                        if false {
                            if (-9223372036854775808 < 42) {
                                var v49 uint64 = v2
                                v4[uint32(v1)] = (fi45 & fi48)
                                var v50 int8 = 0
                                _ = v49
                                _ = v50
                            }
                        } else {
                            v51 := GenS4{F0: 0}
                            var v52 uint64 = (v2 % ((4294967295) | 1))
                            v53 := GenS3{F0: uint(v1), F1: genFn0()}
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                    }
                    var v54 bool = false
                    _ = v54
                }
                if false {
                    var v55 float32 = 0.1
                    _ = v55
                }
                _ = fi48
            }
            var v56 int8 = -121
            _ = fi45
            _ = v56
        }
        var v57 string = "a"
        _ = v44
        _ = v57
    } else {
        switch -9223372036854775808 {
        case -2147483648:
            for fi58 := 0; fi58 < 16; fi58++ {
                for fi59 := 0; fi59 < 20; fi59++ {
                    fmt.Println(v31, v2)
                    _ = fi59
                }
                _ = fi58
            }
        case -1:
            fmt.Println(v31, v31)
        default:
            v2 = (v2 - v2)
        }
        var v60 int8 = -1
        for fi61 := 0; fi61 < 14; fi61++ {
            v1 = v1
            if true {
                v1 = v1
                var v62 int64 = genFn0()
                _ = v62
            } else {
                var v63 uint32 = 270
                for fi64 := 0; fi64 < 7; fi64++ {
                    var v65 string = "hello"
                    _ = fi64
                    _ = v65
                }
                var v66 float64 = 1e10
                _ = v63
                _ = v66
            }
            _ = fi61
        }
        _ = v60
    }
    for fi67 := 0; fi67 < 8; fi67++ {
        for fi68 := 0; fi68 < 19; fi68++ {
            v1 = v1
            fmt.Println(v2)
            _ = fi68
        }
        v4[0] = -983
        _ = fi67
    }
    switch 128 {
    case 42:
        var v69 float32 = ((0.1 + 1.5) + -2.25)
        _ = v69
    default:
        v70 := GenS2{F0: -1, F1: 32767, F2: 42}
        _ = v70
    }
    var v71 uint8 = 255
    fmt.Println(v71, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v71:", v71)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v71
}
