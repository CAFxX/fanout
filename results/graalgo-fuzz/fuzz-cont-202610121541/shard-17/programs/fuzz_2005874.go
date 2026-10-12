package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 bool; F2 int8 }
type GenS5 struct { F0 int8; F1 uint; F2 bool }

func main() {
    var v1 float64 = 1e10
    var v2 int = 32767
    var v3 uint = 1
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    v6 := GenS1{F0: (0 < 2), F1: 127, F2: v3}
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 float32 = ((0.1 * (1e10 + 0.0)) * (1.5 * -2.25))
        v9 := GenS1{F0: false, F1: 42, F2: v3}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v3 = (v3 & (v3 ^ (v3 % (((v3 ^ (^(v3)))) | 1))))
        _ = fi10
    }
    v11 := GenS2{F0: 42, F1: 128, F2: 65535}
    fmt.Println(v3)
    v3 += (^(v3))
    switch "hello" {
    case "foo bar":
        for fi12 := 0; fi12 < 5; fi12++ {
            switch "hello" {
            case "foo bar":
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 uint = v3
                    v15 := GenS2{F0: 321, F1: 0, F2: uint32(v3)}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case "\n\t":
                var v16 int32 = 1
                _ = v16
            }
            _ = fi12
        }
    case "a":
        switch 288 {
        case 4294967295:
            if (103 == 42) {
                if (!(false || (249 <= 826))) {
                    v3 += v3
                    switch ("" + "\n\t") {
                    case "a":
                        v3 -= (v3 | (v3 & v3))
                    }
                } else {
                    v17 := GenS1{F0: false, F1: 0, F2: v3}
                    _ = v17
                }
                v1 = -2.25
                v18 := GenS2{F0: 94, F1: 104, F2: 42}
                _ = v18
            } else {
                switch (v2 | (^(v2))) {
                case 42:
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if true {
                            if (51 == 128) {
                                if true {
                                    switch fi19 {
                                    case 9223372036854775807:
                                        if false {
                                            v2 += (fi19 / ((v2) | 1))
                                            v1 = 1.5
                                            for fi20 := 0; fi20 < 5; fi20++ {
                                                var v21 uint32 = 4294967295
                                                var v22 uint32 = v21
                                                delete(v5, fi19)
                                                delete(v4, v2)
                                                _ = fi20
                                                _ = v21
                                                _ = v22
                                            }
                                        } else {
                                            v1 = (1.5 - float64(1.5))
                                        }
                                    case 128:
                                        switch ("foo bar" + ("foo bar" + ("" + "\n\t"))) {
                                        case "foo bar":
                                            fmt.Println(fi19, v1, fi19)
                                        case "hello":
                                            fmt.Println(v2, v2, v1)
                                        }
                                    default:
                                        var v23 uint8 = 128
                                        _ = v23
                                    }
                                }
                                if (1 > 128) {
                                    if (!((v2 - (fi19 | 327)) != fi19)) {
                                        var v24 string = "hello"
                                        switch (v24 + (v24 + "hello")) {
                                        case "\n\t":
                                            switch (v24 + ((v24 + v24) + (((v24 + (("foo bar" + (v24 + v24)) + v24)) + "") + ""))) {
                                            case "a":
                                                fmt.Println(v1, v24)
                                            case "hello":
                                                v1 += v1
                                            case "\n\t":
                                                var v25 float64 = 0.0
                                                _ = v25
                                            }
                                        default:
                                            var v26 uint8 = 0
                                            _ = v26
                                        }
                                        _ = v24
                                    } else {
                                        switch fi19 {
                                        case 9223372036854775807:
                                            if true {
                                                var v27 int = v2
                                                if (2 > 1) {
                                                    var v28 bool = ((v3 == v3) || false)
                                                    v29 := []int{9223372036854775807, (v2 | v2), -628}
                                                    v30 := GenS2{F0: 18446744073709551615, F1: 255, F2: 0}
                                                    _ = v28
                                                    _ = v29
                                                    _ = v30
                                                }
                                                v31 := v5[v2]
                                                _ = v27
                                                _ = v31
                                            } else {
                                                if (true && false) {
                                                    switch 32767 {
                                                    case 883:
                                                        fmt.Println(v2)
                                                    case -2147483648:
                                                        var v32 int = -391
                                                        _ = v32
                                                    case 32767:
                                                        v33 := []int{1}
                                                        _ = v33
                                                    default:
                                                        v2 = 32768
                                                    }
                                                }
                                            }
                                        case -1:
                                            v1 += ((v1 / float64((3.14159 * (1.5 * -2.25)))) - (1.5 - -2.25))
                                        case 0:
                                            if true {
                                                v34 := GenS1{F0: false, F1: 42, F2: (v3 << ((v3) & 63))}
                                                fmt.Println(v3)
                                                if (-284 >= -1) {
                                                    if ("a" == "\n\t") {
                                                        var v35 float64 = (((float64(((0.0 + -2.25) + 0.1)) * v1) + ((3.14159 / v1) * -2.25)) + v1)
                                                        v36 := GenS3{F0: -1}
                                                        _ = v35
                                                        _ = v36
                                                    } else {
                                                        v3 = v3
                                                        v37 := GenS4{F0: v3, F1: ((0.1 >= (3.14159 * (1.5 / 1.0))) && (v3 > v3)), F2: 0}
                                                        _ = v37
                                                    }
                                                    fmt.Println(v3)
                                                    var v38 int = v2
                                                    _ = v38
                                                }
                                                _ = v34
                                            }
                                        }
                                    }
                                } else {
                                    v3 = v3
                                    var v39 uint16 = 2
                                    _ = v39
                                }
                            }
                        }
                        _ = fi19
                    }
                case -342:
                    if (!true) {
                        v2 = (^(v2))
                        for fi40 := 0; fi40 < 7; fi40++ {
                            if (fi40 == -1) {
                                delete(v5, 32768)
                                var v41 int16 = 1
                                _ = v41
                            }
                            _ = fi40
                        }
                    }
                case 32767:
                    v2 = (v2 / ((v2) | 1))
                default:
                    v2 = v2
                }
            }
        case -395:
            if (((((-2.25 / (0.0 + 1e10)) - 0.1) * 1.5) * 1.5) <= (1.5 + 0.0)) {
                var v42 int32 = 134
                for fi43 := 0; fi43 < 19; fi43++ {
                    var v44 float64 = v1
                    switch (v2 * fi43) {
                    case 854:
                        v45 := GenS5{F0: 100, F1: (v3 & v3), F2: false}
                        _ = v45
                    case 0:
                        for fi46 := 0; fi46 < 16; fi46++ {
                            v1 += (v44 * -2.25)
                            _ = fi46
                        }
                    default:
                        var v47 string = "hello"
                        _ = v47
                    }
                    _ = fi43
                    _ = v44
                }
                switch 32767 {
                case 128:
                    var v48 int64 = -421
                    _ = v48
                case 9223372036854775807:
                    for fi49 := 0; fi49 < 4; fi49++ {
                        var v50 float64 = 3.14159
                        var v51 uint64 = uint64(v3)
                        var v52 bool = true
                        _ = fi49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                case 42:
                    var v53 int8 = 100
                    _ = v53
                }
                _ = v42
            } else {
                var v54 uint64 = 1
                var v55 uint = (v3 << ((v3) & 63))
                var v56 int = (v2 * 32768)
                _ = v54
                _ = v55
                _ = v56
            }
        case -9223372036854775808:
            var v57 int16 = -32768
            _ = v57
        default:
            for fi58 := 0; fi58 < 19; fi58++ {
                v59 := []int{fi58}
                _ = fi58
                _ = v59
            }
        }
    case "":
        v60 := GenS4{F0: v3, F1: true, F2: 127}
        _ = v60
    default:
        fmt.Println(v2)
    }
    v2 = v2
    for fi61 := 0; fi61 < 11; fi61++ {
        var v62 uint8 = 1
        fmt.Println(v62)
        fmt.Println(v3)
        _ = fi61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
