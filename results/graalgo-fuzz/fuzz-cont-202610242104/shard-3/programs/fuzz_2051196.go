package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int64; F1 uint16 }

func main() {
    var v1 uint = 35
    var v2 float32 = 0.0
    var v3 uint32 = 42
    var v4 float64 = 3.14159
    var v5 uint64 = 159
    if (1 == 255) {
        v6 := GenS1{F0: false}
        _ = v6
    } else {
        v5 = v5
    }
    v7 := []int{766, 32768, 128, 127}
    v8 := GenS2{F0: -327, F1: 65535}
    for fi9 := 0; fi9 < 4; fi9++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if (!true) {
            var v10 float64 = -2.25
            if true {
                if (!true) {
                    _ = len(v7)
                    var v11 int64 = 0
                    _ = v11
                } else {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    var v12 int64 = int64(v1)
                    var v13 uint32 = uint32(v5)
                    _ = v12
                    _ = v13
                }
                v5 -= (v5 - (v5 ^ (v5 << ((255) & 63))))
            }
            _ = v10
        } else {
            _ = len(v7)
            var v14 float32 = (((v2 * float32(((v4 + 0.1) + v4))) / v2) + ((v2 * (1e10 - ((float32(0.1) + 3.14159) * 0.0))) * (v2 / -2.25)))
            _ = v14
        }
        if len(v7) > 0 {
            v7[0] = 2147483647
        }
        _ = fi9
    }
    switch "hello" {
    case "hello":
        var v15 uint = 1
        _ = v15
    }
    fmt.Println(v5)
    v5 -= (^(v5))
    var v16 int16 = -782
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 int16 = v16
            if (1 == -9223372036854775808) {
                var v20 int64 = -9223372036854775808
                _ = v20
            } else {
                switch 0 {
                case -2147483648:
                    if false {
                        v21 := GenS2{F0: -2147483648, F1: 65535}
                        var v22 string = "\n\t"
                        _ = v21
                        _ = v22
                    }
                case 128:
                    v3 += (v3 >> ((157) & 63))
                case 2147483647:
                    var v23 int64 = 96
                    _ = v23
                }
            }
            _ = fi18
            _ = v19
        }
        if false {
            switch "foo bar" {
            case "hello":
                fmt.Println(v2, v2, v4)
            }
            if (!(false && (!true))) {
                if len(v7) > 0 {
                    v7[0] = int(v16)
                }
                switch (fi17 * fi17) {
                case 9223372036854775807:
                    v7 = append(v7, (fi17 + (fi17 % ((fi17) | 1))))
                case -2147483648:
                    if ((true || ((true || true) && false)) || (v16 <= -805)) {
                        for fi24 := 0; fi24 < 10; fi24++ {
                            var v25 int16 = v16
                            _ = fi24
                            _ = v25
                        }
                    } else {
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v27 := GenS1{F0: true}
                            _ = fi26
                            _ = v27
                        }
                        if (408 != 881) {
                            switch "\n\t" {
                            case "a":
                                for fi28 := 0; fi28 < 16; fi28++ {
                                    v3 = (v3 + (v3 * v3))
                                    switch (-(fi17)) {
                                    case -9223372036854775808:
                                        v1 += v1
                                    case -2147483648:
                                        fmt.Println(v1, v3, v4)
                                    }
                                    _ = fi28
                                }
                            default:
                                var v29 uint16 = 32767
                                _ = v29
                            }
                            var v30 uint8 = 1
                            v4 = (((1e10 + (((-2.25 - float64(v2)) - v4) / float64((v2 + (1e10 + (v2 - v2)))))) + 0.1) * float64((v2 - v2)))
                            _ = v30
                        } else {
                            v31 := GenS3{F0: "a", F1: (v16 * v16)}
                            var v32 uint32 = v3
                            var v33 float64 = ((((v4 / float64(1.5)) * (1.5 - (((v4 / v4) * ((v4 + 1e10) + 0.1)) + 0.0))) / v4) / ((v4 * (v4 * v4)) + v4))
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    }
                case 127:
                    _ = len(v7)
                }
            }
            for fi34 := 0; fi34 < 14; fi34++ {
                if false {
                    var v35 float64 = v4
                    var v36 uint64 = v5
                    v36 -= (v36 << (((v1 >> ((255) & 63))) & 63))
                    _ = v35
                    _ = v36
                } else {
                    v2 -= 1.5
                    if true {
                        if len(v7) > 0 {
                            v7[0] = 202
                        }
                        switch ("hello" + ("foo bar" + (("hello" + (("a" + "") + (("a" + "\n\t") + ""))) + "a"))) {
                        case "hello":
                            v7 = append(v7, (^(fi17)))
                        case "\n\t":
                            var v37 bool = true
                            _ = v37
                        }
                        var v38 uint = (v1 + v1)
                        _ = v38
                    } else {
                        var v39 uint = uint(v5)
                        _ = v39
                    }
                    fmt.Println(v4, v2, v2)
                    v7 = append(v7, (fi34 ^ (-(fi34))))
                }
                _ = fi34
            }
        } else {
            if true {
                _ = len(v7)
                var v40 uint16 = 68
                _ = v40
            } else {
                if len(v7) > 0 {
                    v7[0] = (fi17 + fi17)
                }
                var v41 uint16 = 65535
                _ = v41
            }
            fmt.Println(v2)
            var v42 uint32 = (^(v3))
            _ = v42
        }
        if (false || false) {
            v43 := GenS1{F0: (!((255 >= 255) && true))}
            v2 -= 1.5
            switch "\n\t" {
            case "":
                var v44 int8 = 42
                _ = v44
            case "a":
                v45 := GenS3{F0: "a", F1: 769}
                _ = v45
            case "\n\t":
                if false {
                    for fi46 := 0; fi46 < 9; fi46++ {
                        _ = len(v7)
                        if ((!(!(!true))) || false) {
                            var v47 int = fi46
                            v2 -= (v2 + v2)
                            _ = v47
                        } else {
                            var v48 int64 = 9223372036854775807
                            v1 = v1
                            v7 = append(v7, (-(fi17)))
                            _ = v48
                        }
                        switch (-(fi46)) {
                        case 670:
                            var v49 int16 = int16(v5)
                            _ = v49
                        case -878:
                            if ((!false) && (-957 == int16(fi46))) {
                                v16 = v16
                                if (!(!true)) {
                                    var v50 int8 = 64
                                    v3 += 42
                                    _ = v50
                                } else {
                                    var v51 float32 = (1.5 + v2)
                                    if true {
                                        if len(v7) > 0 {
                                            v7[0] = fi46
                                        }
                                        v52 := GenS3{F0: "", F1: (^(v16))}
                                        v3 += v3
                                        v3 -= v3
                                        if false {
                                            v53 := GenS3{F0: "hello", F1: 197}
                                            _ = v53
                                        }
                                        _ = v52
                                    } else {
                                        var v54 string = (("foo bar" + "a") + ("hello" + ""))
                                        var v55 uint32 = v3
                                        _ = v54
                                        _ = v55
                                    }
                                    var v56 uint32 = (^(v3))
                                    _ = v51
                                    _ = v56
                                }
                            } else {
                                if true {
                                    var v57 int = (^(fi17))
                                    _ = v57
                                } else {
                                    if false {
                                        v4 = v4
                                    } else {
                                        fmt.Println(v5, fi46)
                                        fmt.Println(v2)
                                        var v58 int8 = 91
                                        v3 -= v3
                                        _ = v58
                                    }
                                    var v59 float64 = ((v4 / v4) + -2.25)
                                    v4 -= (v59 - ((v59 * v4) + (1.5 + 1.5)))
                                    _ = v59
                                }
                                fmt.Println(v3, v5)
                                var v60 bool = (143 == 0)
                                _ = v60
                            }
                        }
                        v2 += v2
                        _ = fi46
                    }
                }
            }
            _ = v43
        } else {
            for fi61 := 0; fi61 < 1; fi61++ {
                var v62 bool = false
                v63 := GenS4{F0: int64(v16), F1: 1}
                var v64 int16 = (v16 | 0)
                fmt.Println(v3, v2)
                _ = fi61
                _ = v62
                _ = v63
                _ = v64
            }
            v16 = (v16 ^ (^(v16)))
            v16 += -32768
        }
        _ = fi17
    }
    for fi65 := 0; fi65 < 15; fi65++ {
        var v66 int64 = -1
        var v67 string = "foo bar"
        v68 := GenS2{F0: fi65, F1: 62}
        _ = fi65
        _ = v66
        _ = v67
        _ = v68
    }
    if (162 == (v5 | 18446744073709551615)) {
        v7 = append(v7, -582)
        v2 += -2.25
    } else {
        var v69 int32 = -678
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
