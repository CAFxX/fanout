package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 bool; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int16; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1e10
    var v3 uint64 = 0
    var v4 uint32 = 42
    var v5 uint16 = 0
    v6 := make(map[string]uint32)
    if true {
        var v7 float64 = 0.1
        fmt.Println(v7)
        _ = v7
    } else {
        switch 32767 {
        case 2147483647:
            v8 := GenS1{F0: 2}
            _ = v8
        default:
            v9 := GenS2{F0: 65535, F1: false, F2: 153}
            _ = v9
        }
        switch -2147483648 {
        case 0:
            v10 := GenS2{F0: 42, F1: false, F2: 241}
            _ = v10
        case 32767:
            for fi11 := 0; fi11 < 14; fi11++ {
                fmt.Println(v2, fi11, v2)
                v1[v4] = fi11
                _ = fi11
            }
        case 225:
            var v12 string = ""
            _ = v12
        default:
            v3 = (^(v3))
        }
    }
    if (1 == (v5 << ((18446744073709551615) & 63))) {
        switch 128 {
        case 9223372036854775807:
            v13 := GenS3{F0: 114}
            _ = v13
        default:
            fmt.Println(v4, v2)
        }
        if true {
            if true {
                v14 := GenS4{F0: -2147483648, F1: 1}
                _ = v14
            } else {
                var v15 uint16 = v5
                fmt.Println(v5, v15, v4)
                var v16 int8 = -1
                _ = v15
                _ = v16
            }
            fmt.Println(v5, v5, v5)
            var v17 float32 = v2
            _ = v17
        } else {
            for fi18 := 0; fi18 < 8; fi18++ {
                v4 -= v4
                var v19 uint = 2
                var v20 string = "\n\t"
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 uint = 1
            v5 += v5
            _ = v21
        }
        var v22 uint64 = (^(v3))
        _ = v22
    } else {
        if (float64(float32((1e10 / 1.0))) > float64(v2)) {
            var v23 int = 2147483647
            fmt.Println(v23)
            _ = v23
        } else {
            _ = len(v1)
            var v24 int32 = 0
            _ = v24
        }
    }
    var v25 int8 = -61
    fmt.Println(v4)
    var v26 int8 = (v25 - int8(v3))
    if (1 < 233) {
        v5 = 2
    } else {
        switch "foo bar" {
        case "foo bar":
            v4 += v4
        case "hello":
            v27 := GenS5{F0: 18446744073709551615}
            _ = v27
        default:
            fmt.Println(v4)
        }
        delete(v1, (v4 >> ((255) & 63)))
    }
    switch -776 {
    case 128:
        if true {
            for fi28 := 0; fi28 < 5; fi28++ {
                fmt.Println(v25, v25)
                _ = fi28
            }
            v29 := []int{32768, -1, 0, 0}
            _ = v29
        } else {
            v26 = 0
            switch "foo bar" {
            case "":
                if ((v3 != (^(v3))) || true) {
                    v30 := []int{}
                    if (true || false) {
                        v26 = (^(v25))
                        switch -2147483648 {
                        case -424:
                            v26 = v25
                        default:
                            var v31 float64 = -2.25
                            _ = v31
                        }
                    }
                    _ = v30
                }
            case "\n\t":
                switch "" {
                case "foo bar":
                    v3 = 2
                case "hello":
                    v32 := GenS4{F0: 0, F1: (v26 / ((int8(v3)) | 1))}
                    _ = v32
                case "\n\t":
                    v33 := []int{-224, -2147483648}
                    _ = v33
                default:
                    for fi34 := 0; fi34 < 14; fi34++ {
                        var v35 uint = 539
                        _ = fi34
                        _ = v35
                    }
                }
            }
            switch 32767 {
            case -1:
                fmt.Println(v2, v4)
            }
        }
    case 42:
        fmt.Println(v2, v25, v4)
    case 720:
        fmt.Println(v26)
    }
    v3 -= (v3 & v3)
    for fi36 := 0; fi36 < 10; fi36++ {
        for fi37 := 0; fi37 < 15; fi37++ {
            switch "\n\t" {
            case "foo bar":
                delete(v1, v4)
            case "\n\t":
                v38 := []int{fi37, fi37}
                _ = v38
            default:
                switch 907 {
                case 4294967295:
                    var v39 int = -554
                    _ = v39
                default:
                    var v40 int64 = -974
                    _ = v40
                }
            }
            var v41 uint8 = 127
            _ = fi37
            _ = v41
        }
        _ = fi36
    }
    _ = len(v6)
    var v42 uint64 = (v3 % (((v3 - (v3 << ((2) & 63)))) | 1))
    fmt.Println(v2, v4, v2)
    if true {
        fmt.Println(v3, v42)
    }
    switch 32767 {
    case 2147483647:
        var v43 int8 = (-(v25))
        _ = v43
    case 9223372036854775807:
        v26 = (v25 / ((42) | 1))
    default:
        for fi44 := 0; fi44 < 4; fi44++ {
            switch (^(fi44)) {
            case -5:
                if (!(!false)) {
                    var v45 uint = 0
                    switch (-(fi44)) {
                    case -609:
                        var v46 int8 = -10
                        _ = v46
                    default:
                        for fi47 := 0; fi47 < 8; fi47++ {
                            var v48 uint = v45
                            fmt.Println(v3, v2)
                            if true {
                                var v49 uint32 = uint32(v45)
                                fmt.Println(v3, v4, v48)
                                var v50 uint32 = v4
                                var v51 bool = (499 >= 1)
                                var v52 int16 = 1000
                                v51 = ((v4 <= 414) || (255 < (v5 % ((v5) | 1))))
                                var v53 int8 = -128
                                if v51 {
                                    if v51 {
                                        if (uint8(v5) >= 255) {
                                            v5 = 502
                                            fmt.Println(v26, v3)
                                            delete(v6, "hello")
                                        }
                                    }
                                    v5 -= v5
                                } else {
                                    if (-1 >= 0) {
                                        var v54 int16 = v52
                                        _ = len(v1)
                                        switch ("hello" + "hello") {
                                        case "\n\t":
                                            v5 = v5
                                        case "foo bar":
                                            _ = len(v1)
                                        default:
                                            if v51 {
                                                v54 = v52
                                                var v55 uint32 = (v4 - v49)
                                                _ = v55
                                            }
                                        }
                                        _ = v54
                                    } else {
                                        fmt.Println(v4, v5)
                                        var v56 uint8 = 1
                                        _ = v56
                                    }
                                }
                                _ = v49
                                _ = v50
                                _ = v51
                                _ = v52
                                _ = v53
                            } else {
                                v57 := v6[((("a" + "foo bar") + "foo bar") + "a")]
                                _ = v57
                            }
                            _ = fi47
                            _ = v48
                        }
                    }
                    _ = v45
                }
            case 127:
                var v58 uint64 = v3
                _ = v58
            default:
                var v59 float64 = -2.25
                _ = v59
            }
            switch (fi44 & int(v25)) {
            case -1:
                for fi60 := 0; fi60 < 18; fi60++ {
                    v2 -= 0.1
                    _ = fi60
                }
            default:
                if (!false) {
                    v61 := GenS6{F0: false, F1: 0, F2: 2147483647}
                    v62 := []int{fi44, fi44}
                    var v63 int32 = -294
                    _ = v61
                    _ = v62
                    _ = v63
                }
            }
            var v64 int32 = 425
            _ = fi44
            _ = v64
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
